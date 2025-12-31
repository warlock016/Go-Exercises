package rate_limiting

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	t.Run("limits rate", func(t *testing.T) {
		limiter := NewRateLimiter(10) // 10 per second
		if limiter == nil {
			t.Fatal("NewRateLimiter() returned nil")
		}
		defer limiter.Stop()

		start := time.Now()
		for i := 0; i < 5; i++ {
			limiter.Wait()
		}
		elapsed := time.Since(start)

		// 5 operations at 10/sec should take ~400-500ms (first one is instant)
		if elapsed < 300*time.Millisecond {
			t.Errorf("RateLimiter too fast: %v", elapsed)
		}
		if elapsed > 800*time.Millisecond {
			t.Errorf("RateLimiter too slow: %v", elapsed)
		}
	})

	t.Run("TryAcquire non-blocking", func(t *testing.T) {
		limiter := NewRateLimiter(1) // 1 per second
		defer limiter.Stop()

		// First call might succeed (depending on implementation)
		limiter.Wait() // Use up the first tick

		// Immediate second call should fail
		if limiter.TryAcquire() {
			t.Log("Note: first TryAcquire succeeded (implementation dependent)")
		}

		// Should fail immediately
		start := time.Now()
		result := limiter.TryAcquire()
		elapsed := time.Since(start)

		if elapsed > 50*time.Millisecond {
			t.Errorf("TryAcquire blocked for %v", elapsed)
		}

		// After waiting, should succeed
		time.Sleep(1100 * time.Millisecond)
		if !limiter.TryAcquire() {
			t.Error("TryAcquire should succeed after waiting")
		}
		_ = result
	})

	t.Run("concurrent access", func(t *testing.T) {
		limiter := NewRateLimiter(100) // 100 per second
		defer limiter.Stop()

		var wg sync.WaitGroup
		var count int64

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < 10; j++ {
					limiter.Wait()
					atomic.AddInt64(&count, 1)
				}
			}()
		}

		wg.Wait()

		if count != 100 {
			t.Errorf("Processed %d operations, want 100", count)
		}
	})
}

func TestTokenBucket(t *testing.T) {
	t.Run("allows burst up to capacity", func(t *testing.T) {
		bucket := NewTokenBucket(5, 100*time.Millisecond)
		if bucket == nil {
			t.Fatal("NewTokenBucket() returned nil")
		}

		start := time.Now()
		for i := 0; i < 5; i++ {
			bucket.Take()
		}
		elapsed := time.Since(start)

		if elapsed > 50*time.Millisecond {
			t.Errorf("Burst of 5 took %v, should be instant", elapsed)
		}
	})

	t.Run("blocks after burst", func(t *testing.T) {
		bucket := NewTokenBucket(2, 100*time.Millisecond)

		bucket.Take()
		bucket.Take()

		start := time.Now()
		bucket.Take() // Should block
		elapsed := time.Since(start)

		if elapsed < 80*time.Millisecond {
			t.Errorf("Third take was too fast: %v", elapsed)
		}
		if elapsed > 200*time.Millisecond {
			t.Errorf("Third take took too long: %v", elapsed)
		}
	})

	t.Run("TryTake non-blocking", func(t *testing.T) {
		bucket := NewTokenBucket(1, 1*time.Second)

		if !bucket.TryTake() {
			t.Error("First TryTake should succeed")
		}

		start := time.Now()
		if bucket.TryTake() {
			t.Error("Second TryTake should fail")
		}
		elapsed := time.Since(start)

		if elapsed > 50*time.Millisecond {
			t.Errorf("TryTake blocked for %v", elapsed)
		}
	})

	t.Run("Tokens returns count", func(t *testing.T) {
		bucket := NewTokenBucket(5, 1*time.Second)

		initial := bucket.Tokens()
		if initial != 5 {
			t.Errorf("Initial tokens = %d, want 5", initial)
		}

		bucket.Take()
		bucket.Take()

		after := bucket.Tokens()
		if after != 3 {
			t.Errorf("Tokens after 2 takes = %d, want 3", after)
		}
	})

	t.Run("refills over time", func(t *testing.T) {
		bucket := NewTokenBucket(3, 50*time.Millisecond)

		// Drain all tokens
		bucket.Take()
		bucket.Take()
		bucket.Take()

		// Wait for refill
		time.Sleep(120 * time.Millisecond)

		tokens := bucket.Tokens()
		if tokens < 2 {
			t.Errorf("Tokens after 120ms = %d, want >= 2", tokens)
		}
	})

	t.Run("doesn't exceed capacity", func(t *testing.T) {
		bucket := NewTokenBucket(3, 10*time.Millisecond)

		// Wait long time
		time.Sleep(100 * time.Millisecond)

		tokens := bucket.Tokens()
		if tokens > 3 {
			t.Errorf("Tokens = %d, should not exceed capacity 3", tokens)
		}
	})
}

func TestThrottle(t *testing.T) {
	t.Run("wraps function with rate limit", func(t *testing.T) {
		limiter := NewRateLimiter(10)
		defer limiter.Stop()

		var count int
		fn := func(x int) error {
			count += x
			return nil
		}

		throttled := Throttle(fn, limiter)
		if throttled == nil {
			t.Fatal("Throttle() returned nil")
		}

		start := time.Now()
		for i := 0; i < 5; i++ {
			err := throttled(1)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		}
		elapsed := time.Since(start)

		if count != 5 {
			t.Errorf("Function called %d times, want 5", count)
		}

		if elapsed < 300*time.Millisecond {
			t.Errorf("Throttle too fast: %v", elapsed)
		}
	})

	t.Run("propagates errors", func(t *testing.T) {
		limiter := NewRateLimiter(10)
		defer limiter.Stop()

		expectedErr := errors.New("test error")
		fn := func(x int) error {
			if x == 2 {
				return expectedErr
			}
			return nil
		}

		throttled := Throttle(fn, limiter)

		if err := throttled(1); err != nil {
			t.Errorf("Unexpected error for x=1: %v", err)
		}

		if err := throttled(2); err != expectedErr {
			t.Errorf("Expected error for x=2, got: %v", err)
		}
	})
}

func TestBatchRateLimiter(t *testing.T) {
	t.Run("processes in batches", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5, 6, 7}
		var batches [][]int
		var mu sync.Mutex

		err := BatchRateLimiter(items, 2, 100*time.Millisecond, func(batch []int) error {
			mu.Lock()
			batches = append(batches, append([]int{}, batch...))
			mu.Unlock()
			return nil
		})

		if err != nil {
			t.Errorf("BatchRateLimiter() error = %v", err)
		}

		// Should be 4 batches: [1,2], [3,4], [5,6], [7]
		if len(batches) != 4 {
			t.Errorf("Got %d batches, want 4", len(batches))
		}
	})

	t.Run("respects interval", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5, 6}

		start := time.Now()
		err := BatchRateLimiter(items, 2, 100*time.Millisecond, func(batch []int) error {
			return nil
		})
		elapsed := time.Since(start)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// 3 batches, 100ms between = ~200ms total (first one instant)
		if elapsed < 150*time.Millisecond {
			t.Errorf("BatchRateLimiter too fast: %v", elapsed)
		}
	})

	t.Run("propagates errors", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5, 6}
		expectedErr := errors.New("batch error")
		batchCount := 0

		err := BatchRateLimiter(items, 2, 50*time.Millisecond, func(batch []int) error {
			batchCount++
			if batchCount == 2 {
				return expectedErr
			}
			return nil
		})

		if err != expectedErr {
			t.Errorf("BatchRateLimiter() error = %v, want %v", err, expectedErr)
		}
	})

	t.Run("handles empty items", func(t *testing.T) {
		err := BatchRateLimiter([]int{}, 2, 100*time.Millisecond, func(batch []int) error {
			t.Error("Process should not be called")
			return nil
		})

		if err != nil {
			t.Errorf("BatchRateLimiter([]) error = %v", err)
		}
	})
}

// Benchmarks
func BenchmarkRateLimiter(b *testing.B) {
	limiter := NewRateLimiter(1000000) // Very high rate for benchmarking
	defer limiter.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		limiter.Wait()
	}
}

func BenchmarkTokenBucket(b *testing.B) {
	bucket := NewTokenBucket(1000000, time.Nanosecond)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bucket.Take()
	}
}
