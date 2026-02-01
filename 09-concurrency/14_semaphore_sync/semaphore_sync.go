package semaphore_sync

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

// Semaphore limits concurrent access
type Semaphore struct {
	// TODO(human): Define fields
	sem chan struct{}
}

// NewSemaphore creates a semaphore with the given number of permits
func NewSemaphore(permits int) *Semaphore {
	// TODO(human): Implement
	return &Semaphore{
		sem: make(chan struct{}, permits),
	}
}

// Acquire blocks until a permit is available
func (s *Semaphore) Acquire() {
	// TODO(human): Implement
	s.sem <- struct{}{}
}

// TryAcquire attempts to acquire without blocking
func (s *Semaphore) TryAcquire() bool {
	// TODO(human): Implement
	select {
	case s.sem <- struct{}{}:
		return true
	default:
		return false
	}
}

// Release returns a permit to the semaphore
func (s *Semaphore) Release() {
	// TODO(human): Implement
	<-s.sem
}

// Available returns the number of available permits
func (s *Semaphore) Available() int {
	// TODO(human): Implement
	return cap(s.sem) - len(s.sem)
}

// WeightedSemaphore allows different weights per acquisition
type WeightedSemaphore struct {
	// TODO(human): Define fields
	availableWeight int64
	maxWeight       int64
	mu              *sync.Cond
}

// NewWeightedSemaphore creates a weighted semaphore
func NewWeightedSemaphore(maxWeight int64) *WeightedSemaphore {
	// TODO(human): Implement

	return &WeightedSemaphore{
		availableWeight: maxWeight,
		maxWeight:       maxWeight,
		mu:              sync.NewCond(&sync.Mutex{}),
	}
}

// Acquire blocks until weight is available or context is cancelled
func (ws *WeightedSemaphore) Acquire(ctx context.Context, weight int64) error {
	// TODO(human): Implement
	if weight > ws.maxWeight {
		return errors.New("weight exceeds max weight")
	}

	ws.mu.L.Lock()
	defer ws.mu.L.Unlock()

	for weight > ws.availableWeight {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			ws.mu.Wait()
		}
	}

	ws.availableWeight -= weight
	return nil
}

// TryAcquire attempts to acquire weight without blocking
func (ws *WeightedSemaphore) TryAcquire(weight int64) bool {
	// TODO(human): Implement
	if weight > ws.maxWeight {
		return false
	}

	ws.mu.L.Lock()
	defer ws.mu.L.Unlock()
	if weight > ws.availableWeight {
		return false
	}
	ws.availableWeight -= weight
	return true
}

// Release returns weight to the semaphore
func (ws *WeightedSemaphore) Release(weight int64) {
	// TODO(human): Implement
	ws.mu.L.Lock()
	defer ws.mu.L.Unlock()
	ws.availableWeight = min(ws.availableWeight+weight, ws.maxWeight) // clamps available weight
	ws.mu.Broadcast()
}

// Barrier waits for N goroutines to arrive before proceeding
type Barrier struct {
	// TODO(human): Define fields
	size      int
	waitCount int
	cond      *sync.Cond
}

// NewBarrier creates a barrier for count participants
func NewBarrier(count int) *Barrier {
	// TODO(human): Implement
	return &Barrier{
		size: count,
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

// Wait blocks until all participants arrive
func (b *Barrier) Wait() {
	// TODO(human): Implement
	b.cond.L.Lock()
	defer b.cond.L.Unlock()
	b.waitCount++
	if b.waitCount == b.size {
		b.waitCount = 0
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
}

// SingleFlight deduplicates concurrent calls for same key

type call struct {
	result any
	err    error
	done   chan struct{}
}
type SingleFlight struct {
	// TODO(human): Define fields
	mu   sync.Mutex
	call map[string]*call
}

// NewSingleFlight creates a new SingleFlight
func NewSingleFlight() *SingleFlight {
	// TODO(human): Implement

	return &SingleFlight{
		call: make(map[string]*call),
	}
}

// Do executes fn for key, deduplicating concurrent calls.
//
// If multiple goroutines call Do("same-key", fn) concurrently:
//   - The FIRST caller becomes the "executor" and runs fn()
//   - All OTHER callers become "waiters" and block until executor finishes
//   - ALL callers receive the SAME result (from the single fn() execution)
func (sf *SingleFlight) Do(key string, fn func() (any, error)) (any, error) {
	var res any
	var err error

	// ============================================================
	// PHASE 1: REGISTRATION (under lock)
	// Determine our role: executor (first) or waiter (subsequent)
	// ============================================================
	sf.mu.Lock()

	// Check if someone is already working on this key.
	// c = reference to the call struct (nil if not exists)
	// exists = true if key is in map (someone already started)
	c, exists := sf.call[key]

	if !exists {
		// We're FIRST for this key → we become the executor.
		// Create a new call struct to track this in-flight operation:
		// - done channel: will be closed when fn() completes (broadcasts to waiters)
		// - result/err: will hold fn()'s return values for waiters to read
		c = &call{
			done: make(chan struct{}),
		}
		// Register in map so subsequent callers find us and become waiters.
		sf.call[key] = c
	}
	// If exists == true, c now points to the existing call struct
	// created by the executor. We'll wait on c.done.

	sf.mu.Unlock()
	// IMPORTANT: We release the lock BEFORE doing slow work (fn() or waiting).
	// This allows other goroutines to check the map and become waiters.

	// ============================================================
	// PHASE 2: EXECUTION OR WAITING (no lock held)
	// ============================================================
	if !exists {
		// ----------------------------------------------------
		// EXECUTOR PATH: We run fn() and share results
		// ----------------------------------------------------

		// Run the expensive function (this is the whole point of SingleFlight:
		// only ONE goroutine runs this, even if 100 called Do() concurrently)
		res, err = fn()

		// Store results in the shared call struct.
		// Waiters will read these AFTER we close the done channel.
		c.result = res
		c.err = err

		// Signal completion to ALL waiters by closing the channel.
		// close() is a broadcast: every goroutine blocked on <-c.done wakes up.
		close(c.done)

		// Clean up: remove from map so future calls start fresh.
		// Need lock because we're modifying the map.
		sf.mu.Lock()
		delete(sf.call, key)
		sf.mu.Unlock()

	} else {
		// ----------------------------------------------------
		// WAITER PATH: Block until executor finishes, then read results
		// ----------------------------------------------------

		// Block here until the executor closes c.done.
		// When close(c.done) happens, this receive completes immediately.
		<-c.done

		// Executor has finished and stored results. Read them.
		// Safe to read: executor wrote BEFORE close(), we read AFTER close().
		// This is the "happens-before" guarantee of channel close.
		res = c.result
		err = c.err
	}

	// ============================================================
	// PHASE 3: RETURN (both paths converge here)
	// ============================================================
	return res, err
}

// LazyInit provides lazy initialization with sync.Once semantics
type LazyInit[T any] struct {
	// TODO(human): Define fields
	once sync.Once
	val  T
	init atomic.Bool
}

// Get returns the initialized value, calling init once if needed
func (l *LazyInit[T]) Get(init func() T) T {
	// TODO(human): Implement
	// var zero T
	l.once.Do(func() {
		l.val = init()
		l.init.Store(true)
	})
	return l.val
}

// IsInitialized returns whether the value has been initialized
func (l *LazyInit[T]) IsInitialized() bool {
	// TODO(human): Implement
	return l.init.Load()
}
