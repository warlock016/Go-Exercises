package closure_practice

import (
	"sync"
	"time"
)

// Memoize wraps a function to cache its results
// Subsequent calls with the same input return cached result
func Memoize(fn func(int) int) func(int) int {
	// TODO(human): Implement
	cache := map[int]int{}

	return func(i int) int {
		if _, ok := cache[i]; ok {
			return cache[i]
		} else {
			cache[i] = fn(i)
			return cache[i]
		}
		// return fn(i)
	}
}

// Toggle returns a function that alternates between a and b
func Toggle[T any](a, b T) func() T {
	// TODO(human): Implement

	currState := true

	return func() T {
		if currState {
			currState = false
			return a
		} else {
			currState = true
			return b
		}
	}
}

// Once returns a function that calls fn only on the first invocation
// Subsequent calls return the first result without calling fn again
func Once[T any](fn func() T) func() T {
	// TODO(human): Implement

	var cache T
	toggle := false

	return func() T {
		if !toggle {
			toggle = true
			cache = fn()
			return cache
		} else {
			return cache

		}
	}
}

// Debounce returns a function that delays calling fn until
// 'delay' duration passes without another call
// Returns the debounced function and a cancel function
func Debounce(fn func(), delay time.Duration) (debounced func(), cancel func()) {
	// TODO(human): Implement
	var (
		mu    sync.Mutex
		timer *time.Timer
		seq   int64
	)

	debounced = func() {
		mu.Lock()    // acquire lock
		seq++        // increment counter
		mySeq := seq // set local counter to current closure counter

		if timer != nil { // check if timer expired
			timer.Stop()
		}

		timer = time.AfterFunc(delay, func() {

			mu.Lock()
			ok := (mySeq == seq)
			mu.Unlock()

			if ok {
				fn()
			}
		})

		mu.Unlock()
	}

	cancel = func() {
		mu.Lock()
		seq++
		if timer != nil {
			timer.Stop()
			timer = nil
		}

		mu.Unlock()
	}

	return debounced, cancel
}

// Sequence creates a generator starting at 'start'
// Each call applies 'next' function to get the following value
func Sequence(start int, next func(int) int) func() int {
	// TODO(human): Implement

	nextVal := start

	return func() int {
		val := nextVal
		nextVal = next(nextVal)
		return val
	}
}

// Retry wraps a function to retry on error up to maxAttempts times
// Uses exponential backoff starting with initialDelay
func Retry(fn func() error, maxAttempts int, initialDelay time.Duration) func() error {
	// TODO(human): Implement

	// start := time.Time{}
	delay := initialDelay
	attempts := 0
	var lastError error

	// if "fn" succeeds it returns nil, else it returns some error
	// while "fn" != nil, we check for time passed and retry

	return func() error {

		for attempts < maxAttempts {

			lastError = fn()
			if lastError == nil {
				return nil
			}
			time.Sleep(delay)
			delay *= 2
			attempts++
		}

		return lastError
	}
}

// BONUS: Fibonacci returns a generator for the Fibonacci sequence
func Fibonacci() func() int {
	// TODO(human): Implement
	init := []int{}

	return func() int {
		var next int

		switch len(init) {
		case 0:
			next = 0
			init = append(init, next)
			return next
		case 1:
			next = 1
			init = append(init, next)
			return next
		default:
			next = init[len(init)-2] + init[len(init)-1]
			init = append(init, next)
			return next
		}
	}
}
