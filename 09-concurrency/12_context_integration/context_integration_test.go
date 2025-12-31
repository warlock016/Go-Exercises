package context_integration

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestWorkerWithContext(t *testing.T) {
	t.Run("processes items", func(t *testing.T) {
		ctx := context.Background()
		items := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			items <- i
		}
		close(items)

		var sum int
		err := WorkerWithContext(ctx, items, func(i int) error {
			sum += i
			return nil
		})

		if err != nil {
			t.Errorf("WorkerWithContext() error = %v", err)
		}
		if sum != 15 {
			t.Errorf("sum = %d, want 15", sum)
		}
	})

	t.Run("stops on context cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		items := make(chan int)

		go func() {
			for i := 0; ; i++ {
				select {
				case items <- i:
				case <-ctx.Done():
					return
				}
			}
		}()

		var count int
		go func() {
			time.Sleep(50 * time.Millisecond)
			cancel()
		}()

		err := WorkerWithContext(ctx, items, func(i int) error {
			count++
			time.Sleep(10 * time.Millisecond)
			return nil
		})

		if !errors.Is(err, context.Canceled) {
			t.Errorf("WorkerWithContext() error = %v, want context.Canceled", err)
		}
	})

	t.Run("propagates process error", func(t *testing.T) {
		ctx := context.Background()
		items := make(chan int, 3)
		items <- 1
		items <- 2
		items <- 3
		close(items)

		expectedErr := errors.New("process failed")
		err := WorkerWithContext(ctx, items, func(i int) error {
			if i == 2 {
				return expectedErr
			}
			return nil
		})

		if err != expectedErr {
			t.Errorf("WorkerWithContext() error = %v, want %v", err, expectedErr)
		}
	})
}

func TestFetchAll(t *testing.T) {
	t.Run("fetches all successfully", func(t *testing.T) {
		ctx := context.Background()
		urls := []string{"http://a.com", "http://b.com"}

		results, err := FetchAll(ctx, urls, func(ctx context.Context, url string) (string, error) {
			return "result-" + url, nil
		})

		if err != nil {
			t.Errorf("FetchAll() error = %v", err)
		}
		if results == nil {
			t.Fatal("FetchAll() results is nil")
		}
		if len(results) != 2 {
			t.Errorf("FetchAll() got %d results, want 2", len(results))
		}
	})

	t.Run("cancels all on context cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		urls := []string{"http://a.com", "http://b.com"}

		var fetchCount int64
		cancel() // Cancel immediately

		_, err := FetchAll(ctx, urls, func(ctx context.Context, url string) (string, error) {
			atomic.AddInt64(&fetchCount, 1)
			select {
			case <-ctx.Done():
				return "", ctx.Err()
			case <-time.After(1 * time.Second):
				return "result", nil
			}
		})

		if !errors.Is(err, context.Canceled) {
			t.Errorf("FetchAll() error = %v, want context.Canceled", err)
		}
	})

	t.Run("runs in parallel", func(t *testing.T) {
		ctx := context.Background()
		urls := []string{"http://a.com", "http://b.com", "http://c.com"}

		start := time.Now()
		FetchAll(ctx, urls, func(ctx context.Context, url string) (string, error) {
			time.Sleep(50 * time.Millisecond)
			return "result", nil
		})
		elapsed := time.Since(start)

		if elapsed > 100*time.Millisecond {
			t.Errorf("FetchAll() took %v, expected parallel execution", elapsed)
		}
	})
}

func TestFetchRace(t *testing.T) {
	t.Run("returns first result", func(t *testing.T) {
		ctx := context.Background()
		urls := []string{"http://slow.com", "http://fast.com", "http://medium.com"}

		start := time.Now()
		result, err := FetchRace(ctx, urls, func(ctx context.Context, url string) (string, error) {
			delay := 100 * time.Millisecond
			if url == "http://fast.com" {
				delay = 10 * time.Millisecond
			}
			select {
			case <-time.After(delay):
				return "result-" + url, nil
			case <-ctx.Done():
				return "", ctx.Err()
			}
		})
		elapsed := time.Since(start)

		if err != nil {
			t.Errorf("FetchRace() error = %v", err)
		}
		if result != "result-http://fast.com" {
			t.Errorf("FetchRace() = %q, want result from fast.com", result)
		}
		if elapsed > 50*time.Millisecond {
			t.Errorf("FetchRace() took %v, should return early", elapsed)
		}
	})

	t.Run("handles all errors", func(t *testing.T) {
		ctx := context.Background()
		urls := []string{"http://a.com", "http://b.com"}

		_, err := FetchRace(ctx, urls, func(ctx context.Context, url string) (string, error) {
			return "", errors.New("fetch failed")
		})

		if err == nil {
			t.Error("FetchRace() should return error when all fail")
		}
	})

	t.Run("empty urls", func(t *testing.T) {
		ctx := context.Background()

		_, err := FetchRace(ctx, nil, func(ctx context.Context, url string) (string, error) {
			return "result", nil
		})

		if err == nil {
			t.Error("FetchRace() with no urls should return error")
		}
	})
}

func TestTimeoutPipeline(t *testing.T) {
	t.Run("processes items", func(t *testing.T) {
		ctx := context.Background()
		input := make(chan int, 3)
		input <- 1
		input <- 2
		input <- 3
		close(input)

		output := TimeoutPipeline(ctx, input, 100*time.Millisecond, func(ctx context.Context, i int) (int, error) {
			return i * 2, nil
		})
		if output == nil {
			t.Fatal("TimeoutPipeline() returned nil")
		}

		var results []int
		for r := range output {
			if r.Err != nil {
				t.Errorf("Unexpected error: %v", r.Err)
			}
			results = append(results, r.Value)
		}

		if len(results) != 3 {
			t.Errorf("Got %d results, want 3", len(results))
		}
	})

	t.Run("times out slow items", func(t *testing.T) {
		ctx := context.Background()
		input := make(chan int, 2)
		input <- 1 // Fast
		input <- 2 // Slow
		close(input)

		output := TimeoutPipeline(ctx, input, 50*time.Millisecond, func(ctx context.Context, i int) (int, error) {
			if i == 2 {
				select {
				case <-time.After(200 * time.Millisecond):
					return i, nil
				case <-ctx.Done():
					return 0, ctx.Err()
				}
			}
			return i * 2, nil
		})

		var errors int
		for r := range output {
			if r.Err != nil {
				errors++
			}
		}

		if errors != 1 {
			t.Errorf("Got %d timeout errors, want 1", errors)
		}
	})
}

func TestRetryWithBackoff(t *testing.T) {
	t.Run("succeeds on first try", func(t *testing.T) {
		ctx := context.Background()
		attempts := 0

		err := RetryWithBackoff(ctx, 3, 10*time.Millisecond, func(ctx context.Context) error {
			attempts++
			return nil
		})

		if err != nil {
			t.Errorf("RetryWithBackoff() error = %v", err)
		}
		if attempts != 1 {
			t.Errorf("attempts = %d, want 1", attempts)
		}
	})

	t.Run("retries on failure", func(t *testing.T) {
		ctx := context.Background()
		attempts := 0

		err := RetryWithBackoff(ctx, 3, 10*time.Millisecond, func(ctx context.Context) error {
			attempts++
			if attempts < 3 {
				return errors.New("temporary error")
			}
			return nil
		})

		if err != nil {
			t.Errorf("RetryWithBackoff() error = %v", err)
		}
		if attempts != 3 {
			t.Errorf("attempts = %d, want 3", attempts)
		}
	})

	t.Run("returns error after max retries", func(t *testing.T) {
		ctx := context.Background()
		expectedErr := errors.New("persistent error")

		err := RetryWithBackoff(ctx, 3, 10*time.Millisecond, func(ctx context.Context) error {
			return expectedErr
		})

		if err != expectedErr {
			t.Errorf("RetryWithBackoff() error = %v, want %v", err, expectedErr)
		}
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		err := RetryWithBackoff(ctx, 10, 1*time.Second, func(ctx context.Context) error {
			return errors.New("should not reach here multiple times")
		})

		if !errors.Is(err, context.Canceled) {
			t.Errorf("RetryWithBackoff() error = %v, want context.Canceled", err)
		}
	})

	t.Run("uses exponential backoff", func(t *testing.T) {
		ctx := context.Background()
		attempts := 0

		start := time.Now()
		RetryWithBackoff(ctx, 3, 50*time.Millisecond, func(ctx context.Context) error {
			attempts++
			if attempts < 3 {
				return errors.New("fail")
			}
			return nil
		})
		elapsed := time.Since(start)

		// Delays: 50ms + 100ms = 150ms minimum
		if elapsed < 100*time.Millisecond {
			t.Errorf("Backoff too fast: %v", elapsed)
		}
	})
}

func TestParallel(t *testing.T) {
	t.Run("runs all functions", func(t *testing.T) {
		ctx := context.Background()
		var count int64

		err := Parallel(ctx,
			func(ctx context.Context) error { atomic.AddInt64(&count, 1); return nil },
			func(ctx context.Context) error { atomic.AddInt64(&count, 1); return nil },
			func(ctx context.Context) error { atomic.AddInt64(&count, 1); return nil },
		)

		if err != nil {
			t.Errorf("Parallel() error = %v", err)
		}
		if count != 3 {
			t.Errorf("count = %d, want 3", count)
		}
	})

	t.Run("returns first error", func(t *testing.T) {
		ctx := context.Background()
		expectedErr := errors.New("function failed")

		err := Parallel(ctx,
			func(ctx context.Context) error { time.Sleep(100 * time.Millisecond); return nil },
			func(ctx context.Context) error { return expectedErr },
			func(ctx context.Context) error { time.Sleep(100 * time.Millisecond); return nil },
		)

		if err != expectedErr {
			t.Errorf("Parallel() error = %v, want %v", err, expectedErr)
		}
	})

	t.Run("cancels others on error", func(t *testing.T) {
		ctx := context.Background()
		var cancelled int64

		Parallel(ctx,
			func(ctx context.Context) error {
				select {
				case <-ctx.Done():
					atomic.AddInt64(&cancelled, 1)
					return ctx.Err()
				case <-time.After(1 * time.Second):
					return nil
				}
			},
			func(ctx context.Context) error {
				return errors.New("fast failure")
			},
			func(ctx context.Context) error {
				select {
				case <-ctx.Done():
					atomic.AddInt64(&cancelled, 1)
					return ctx.Err()
				case <-time.After(1 * time.Second):
					return nil
				}
			},
		)

		time.Sleep(100 * time.Millisecond) // Give goroutines time to notice cancellation

		if cancelled == 0 {
			t.Error("Other functions should be cancelled on error")
		}
	})

	t.Run("runs in parallel", func(t *testing.T) {
		ctx := context.Background()

		start := time.Now()
		Parallel(ctx,
			func(ctx context.Context) error { time.Sleep(50 * time.Millisecond); return nil },
			func(ctx context.Context) error { time.Sleep(50 * time.Millisecond); return nil },
			func(ctx context.Context) error { time.Sleep(50 * time.Millisecond); return nil },
		)
		elapsed := time.Since(start)

		if elapsed > 100*time.Millisecond {
			t.Errorf("Parallel() took %v, expected parallel execution", elapsed)
		}
	})

	t.Run("empty functions", func(t *testing.T) {
		ctx := context.Background()
		err := Parallel(ctx)

		if err != nil {
			t.Errorf("Parallel() with no functions error = %v", err)
		}
	})
}

// Benchmarks
func BenchmarkRetryWithBackoff(b *testing.B) {
	ctx := context.Background()

	b.Run("success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RetryWithBackoff(ctx, 3, time.Nanosecond, func(ctx context.Context) error {
				return nil
			})
		}
	})
}

func BenchmarkParallel(b *testing.B) {
	ctx := context.Background()
	fn := func(ctx context.Context) error { return nil }

	b.Run("3-functions", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Parallel(ctx, fn, fn, fn)
		}
	})
}
