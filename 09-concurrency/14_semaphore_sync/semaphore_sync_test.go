package semaphore_sync

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	t.Run("basic acquire and release", func(t *testing.T) {
		sem := NewSemaphore(2)
		if sem == nil {
			t.Fatal("NewSemaphore() returned nil")
		}

		sem.Acquire()
		sem.Acquire()

		// Third should block, use TryAcquire to test
		if sem.TryAcquire() {
			t.Error("TryAcquire should fail when at capacity")
		}

		sem.Release()

		if !sem.TryAcquire() {
			t.Error("TryAcquire should succeed after release")
		}
	})

	t.Run("limits concurrency", func(t *testing.T) {
		sem := NewSemaphore(3)
		var concurrent int64
		var maxConcurrent int64
		var wg sync.WaitGroup

		for i := 0; i < 20; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				sem.Acquire()
				defer sem.Release()

				c := atomic.AddInt64(&concurrent, 1)
				for {
					max := atomic.LoadInt64(&maxConcurrent)
					if c <= max || atomic.CompareAndSwapInt64(&maxConcurrent, max, c) {
						break
					}
				}
				time.Sleep(10 * time.Millisecond)
				atomic.AddInt64(&concurrent, -1)
			}()
		}

		wg.Wait()

		if maxConcurrent > 3 {
			t.Errorf("Max concurrent = %d, want <= 3", maxConcurrent)
		}
	})

	t.Run("Available returns correct count", func(t *testing.T) {
		sem := NewSemaphore(5)

		if a := sem.Available(); a != 5 {
			t.Errorf("Available() = %d, want 5", a)
		}

		sem.Acquire()
		sem.Acquire()

		if a := sem.Available(); a != 3 {
			t.Errorf("Available() after 2 acquires = %d, want 3", a)
		}

		sem.Release()

		if a := sem.Available(); a != 4 {
			t.Errorf("Available() after release = %d, want 4", a)
		}
	})
}

func TestWeightedSemaphore(t *testing.T) {
	t.Run("respects weights", func(t *testing.T) {
		ws := NewWeightedSemaphore(10)
		if ws == nil {
			t.Fatal("NewWeightedSemaphore() returned nil")
		}

		ctx := context.Background()

		err := ws.Acquire(ctx, 4)
		if err != nil {
			t.Errorf("Acquire(4) error = %v", err)
		}

		err = ws.Acquire(ctx, 4)
		if err != nil {
			t.Errorf("Acquire(4) error = %v", err)
		}

		// Only 2 weight left, can't acquire 4
		if ws.TryAcquire(4) {
			t.Error("TryAcquire(4) should fail with only 2 available")
		}

		if !ws.TryAcquire(2) {
			t.Error("TryAcquire(2) should succeed with 2 available")
		}
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		ws := NewWeightedSemaphore(5)
		ctx := context.Background()

		// Use up all weight
		ws.Acquire(ctx, 5)

		cancelCtx, cancel := context.WithCancel(context.Background())
		cancel()

		err := ws.Acquire(cancelCtx, 3)
		if !errors.Is(err, context.Canceled) {
			t.Errorf("Acquire with cancelled context error = %v, want Canceled", err)
		}
	})

	t.Run("release restores weight", func(t *testing.T) {
		ws := NewWeightedSemaphore(10)
		ctx := context.Background()

		ws.Acquire(ctx, 7)

		if ws.TryAcquire(5) {
			t.Error("Should not acquire 5 with only 3 available")
		}

		ws.Release(4)

		if !ws.TryAcquire(5) {
			t.Error("Should acquire 5 with 7 available")
		}
	})
}

func TestBarrier(t *testing.T) {
	t.Run("waits for all participants", func(t *testing.T) {
		barrier := NewBarrier(3)
		if barrier == nil {
			t.Fatal("NewBarrier() returned nil")
		}

		var phase1Complete int64
		var phase2Started int64
		var wg sync.WaitGroup

		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()

				// Phase 1
				time.Sleep(time.Duration(id*10) * time.Millisecond)
				atomic.AddInt64(&phase1Complete, 1)

				barrier.Wait()

				// Phase 2 - all should see phase1Complete == 3
				if atomic.LoadInt64(&phase1Complete) != 3 {
					t.Errorf("Worker %d saw phase1Complete != 3", id)
				}
				atomic.AddInt64(&phase2Started, 1)
			}(i)
		}

		wg.Wait()

		if phase2Started != 3 {
			t.Errorf("phase2Started = %d, want 3", phase2Started)
		}
	})

	t.Run("blocks until all arrive", func(t *testing.T) {
		barrier := NewBarrier(2)
		arrived := make(chan struct{})

		go func() {
			barrier.Wait()
			close(arrived)
		}()

		// Should not complete immediately
		select {
		case <-arrived:
			t.Error("Barrier should block until all arrive")
		case <-time.After(50 * time.Millisecond):
			// Expected
		}

		// Second participant arrives
		go barrier.Wait()

		// Now should complete
		select {
		case <-arrived:
			// Success
		case <-time.After(100 * time.Millisecond):
			t.Error("Barrier should release after all arrive")
		}
	})
}

func TestSingleFlight(t *testing.T) {
	t.Run("executes function once", func(t *testing.T) {
		sf := NewSingleFlight()
		if sf == nil {
			t.Fatal("NewSingleFlight() returned nil")
		}

		var callCount int64
		var wg sync.WaitGroup

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				sf.Do("key", func() (interface{}, error) {
					atomic.AddInt64(&callCount, 1)
					time.Sleep(50 * time.Millisecond)
					return "result", nil
				})
			}()
		}

		wg.Wait()

		if callCount != 1 {
			t.Errorf("Function called %d times, want 1", callCount)
		}
	})

	t.Run("returns same result to all", func(t *testing.T) {
		sf := NewSingleFlight()
		var wg sync.WaitGroup
		results := make(chan interface{}, 5)

		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result, _ := sf.Do("key", func() (interface{}, error) {
					time.Sleep(50 * time.Millisecond)
					return 42, nil
				})
				results <- result
			}()
		}

		wg.Wait()
		close(results)

		for r := range results {
			if r != 42 {
				t.Errorf("Got result %v, want 42", r)
			}
		}
	})

	t.Run("different keys execute separately", func(t *testing.T) {
		sf := NewSingleFlight()
		var count1, count2 int64
		var wg sync.WaitGroup

		wg.Add(2)
		go func() {
			defer wg.Done()
			sf.Do("key1", func() (interface{}, error) {
				atomic.AddInt64(&count1, 1)
				return nil, nil
			})
		}()
		go func() {
			defer wg.Done()
			sf.Do("key2", func() (interface{}, error) {
				atomic.AddInt64(&count2, 1)
				return nil, nil
			})
		}()

		wg.Wait()

		if count1 != 1 || count2 != 1 {
			t.Errorf("Different keys should execute separately: count1=%d, count2=%d", count1, count2)
		}
	})

	t.Run("propagates errors", func(t *testing.T) {
		sf := NewSingleFlight()
		expectedErr := errors.New("test error")

		_, err := sf.Do("key", func() (interface{}, error) {
			return nil, expectedErr
		})

		if err != expectedErr {
			t.Errorf("Do() error = %v, want %v", err, expectedErr)
		}
	})
}

func TestLazyInit(t *testing.T) {
	t.Run("initializes once", func(t *testing.T) {
		var li LazyInit[int]
		var callCount int64

		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				li.Get(func() int {
					atomic.AddInt64(&callCount, 1)
					return 42
				})
			}()
		}

		wg.Wait()

		if callCount != 1 {
			t.Errorf("Init called %d times, want 1", callCount)
		}
	})

	t.Run("returns same value", func(t *testing.T) {
		var li LazyInit[string]

		v1 := li.Get(func() string { return "first" })
		v2 := li.Get(func() string { return "second" })

		if v1 != "first" || v2 != "first" {
			t.Errorf("Got different values: %q, %q", v1, v2)
		}
	})

	t.Run("IsInitialized", func(t *testing.T) {
		var li LazyInit[int]

		if li.IsInitialized() {
			t.Error("IsInitialized should be false before Get")
		}

		li.Get(func() int { return 1 })

		if !li.IsInitialized() {
			t.Error("IsInitialized should be true after Get")
		}
	})
}

// Benchmarks
func BenchmarkSemaphore(b *testing.B) {
	sem := NewSemaphore(10)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sem.Acquire()
			sem.Release()
		}
	})
}

func BenchmarkSingleFlight(b *testing.B) {
	sf := NewSingleFlight()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sf.Do("key", func() (interface{}, error) {
				return 42, nil
			})
		}
	})
}
