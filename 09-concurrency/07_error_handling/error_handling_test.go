package error_handling

import (
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestProcessWithErrors(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		errs := ProcessWithErrors(items, func(i int) error {
			return nil
		})

		if len(errs) != 0 {
			t.Errorf("ProcessWithErrors() got %d errors, want 0", len(errs))
		}
	})

	t.Run("some errors", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		errs := ProcessWithErrors(items, func(i int) error {
			if i%2 == 0 {
				return fmt.Errorf("even: %d", i)
			}
			return nil
		})

		if len(errs) != 2 {
			t.Errorf("ProcessWithErrors() got %d errors, want 2", len(errs))
		}
	})

	t.Run("all errors", func(t *testing.T) {
		items := []int{1, 2, 3}
		errs := ProcessWithErrors(items, func(i int) error {
			return fmt.Errorf("error %d", i)
		})

		if len(errs) != 3 {
			t.Errorf("ProcessWithErrors() got %d errors, want 3", len(errs))
		}
	})

	t.Run("empty items", func(t *testing.T) {
		errs := ProcessWithErrors(nil, func(i int) error {
			return errors.New("should not be called")
		})

		if errs == nil {
			// nil is acceptable for empty
		} else if len(errs) != 0 {
			t.Errorf("ProcessWithErrors() with empty items got %d errors", len(errs))
		}
	})

	t.Run("processes concurrently", func(t *testing.T) {
		items := make([]int, 10)
		for i := range items {
			items[i] = i
		}

		start := time.Now()
		ProcessWithErrors(items, func(i int) error {
			time.Sleep(50 * time.Millisecond)
			return nil
		})
		elapsed := time.Since(start)

		// Sequential would take 500ms, concurrent should be ~50ms
		if elapsed > 200*time.Millisecond {
			t.Errorf("ProcessWithErrors() took %v, expected concurrent execution", elapsed)
		}
	})
}

func TestFirstError(t *testing.T) {
	t.Run("all succeed", func(t *testing.T) {
		tasks := []func() error{
			func() error { return nil },
			func() error { return nil },
			func() error { return nil },
		}

		err := FirstError(tasks)
		if err != nil {
			t.Errorf("FirstError() = %v, want nil", err)
		}
	})

	t.Run("first task fails", func(t *testing.T) {
		tasks := []func() error{
			func() error { return errors.New("first error") },
			func() error { time.Sleep(100 * time.Millisecond); return nil },
		}

		err := FirstError(tasks)
		if err == nil {
			t.Error("FirstError() = nil, want error")
		}
	})

	t.Run("later task fails", func(t *testing.T) {
		tasks := []func() error{
			func() error { time.Sleep(100 * time.Millisecond); return nil },
			func() error { return errors.New("quick error") },
		}

		start := time.Now()
		err := FirstError(tasks)
		elapsed := time.Since(start)

		if err == nil {
			t.Error("FirstError() = nil, want error")
		}
		// Should return quickly, not wait for slow task
		if elapsed > 50*time.Millisecond {
			t.Errorf("FirstError() took %v, should return early", elapsed)
		}
	})

	t.Run("empty tasks", func(t *testing.T) {
		err := FirstError(nil)
		if err != nil {
			t.Errorf("FirstError(nil) = %v, want nil", err)
		}
	})
}

func TestProcessResults(t *testing.T) {
	t.Run("all succeed", func(t *testing.T) {
		items := []int{1, 2, 3}
		results := ProcessResults(items, func(i int) (string, error) {
			return fmt.Sprintf("result-%d", i), nil
		})

		if len(results) != 3 {
			t.Fatalf("ProcessResults() got %d results, want 3", len(results))
		}

		for i, r := range results {
			if r.Err != nil {
				t.Errorf("results[%d].Err = %v, want nil", i, r.Err)
			}
			want := fmt.Sprintf("result-%d", items[i])
			if r.Value != want {
				t.Errorf("results[%d].Value = %q, want %q", i, r.Value, want)
			}
		}
	})

	t.Run("some fail", func(t *testing.T) {
		items := []int{1, 2, 3}
		results := ProcessResults(items, func(i int) (string, error) {
			if i == 2 {
				return "", errors.New("failed")
			}
			return fmt.Sprintf("result-%d", i), nil
		})

		if len(results) != 3 {
			t.Fatalf("ProcessResults() got %d results, want 3", len(results))
		}

		// Result for item 2 should have error
		var foundError bool
		for _, r := range results {
			if r.Err != nil {
				foundError = true
			}
		}
		if !foundError {
			t.Error("ProcessResults() should have at least one error")
		}
	})

	t.Run("preserves order", func(t *testing.T) {
		items := []int{10, 20, 30}
		results := ProcessResults(items, func(i int) (int, error) {
			time.Sleep(time.Duration(30-i) * time.Millisecond) // Reverse completion order
			return i * 2, nil
		})

		// Results should be in original order despite completion order
		expected := []int{20, 40, 60}
		for i, r := range results {
			if r.Value != expected[i] {
				t.Errorf("results[%d].Value = %d, want %d (order not preserved)", i, r.Value, expected[i])
			}
		}
	})
}

func TestRunWithTimeout(t *testing.T) {
	t.Run("completes before timeout", func(t *testing.T) {
		err := RunWithTimeout(func() error {
			time.Sleep(10 * time.Millisecond)
			return nil
		}, 100*time.Millisecond)

		if err != nil {
			t.Errorf("RunWithTimeout() = %v, want nil", err)
		}
	})

	t.Run("task error", func(t *testing.T) {
		taskErr := errors.New("task failed")
		err := RunWithTimeout(func() error {
			return taskErr
		}, 100*time.Millisecond)

		if err != taskErr {
			t.Errorf("RunWithTimeout() = %v, want %v", err, taskErr)
		}
	})

	t.Run("timeout exceeded", func(t *testing.T) {
		start := time.Now()
		err := RunWithTimeout(func() error {
			time.Sleep(500 * time.Millisecond)
			return nil
		}, 50*time.Millisecond)
		elapsed := time.Since(start)

		if err == nil {
			t.Error("RunWithTimeout() = nil, want timeout error")
		}
		if elapsed > 100*time.Millisecond {
			t.Errorf("RunWithTimeout() took %v, should timeout faster", elapsed)
		}
	})
}

func TestParallelFetch(t *testing.T) {
	t.Run("all succeed", func(t *testing.T) {
		urls := []string{"http://a.com", "http://b.com", "http://c.com"}
		results, errs := ParallelFetch(urls, func(url string) (string, error) {
			return "content-" + url, nil
		})

		if len(errs) != 0 {
			t.Errorf("ParallelFetch() got %d errors, want 0", len(errs))
		}

		if results == nil {
			t.Fatal("ParallelFetch() results is nil")
		}

		for _, url := range urls {
			want := "content-" + url
			if got := results[url]; got != want {
				t.Errorf("results[%s] = %q, want %q", url, got, want)
			}
		}
	})

	t.Run("some fail", func(t *testing.T) {
		urls := []string{"http://good.com", "http://bad.com", "http://good2.com"}
		results, errs := ParallelFetch(urls, func(url string) (string, error) {
			if strings.Contains(url, "bad") {
				return "", errors.New("fetch failed")
			}
			return "content", nil
		})

		if len(errs) != 1 {
			t.Errorf("ParallelFetch() got %d errors, want 1", len(errs))
		}

		// Good URLs should still have results
		if _, ok := results["http://good.com"]; !ok {
			t.Error("results missing http://good.com")
		}
		if _, ok := results["http://good2.com"]; !ok {
			t.Error("results missing http://good2.com")
		}
	})

	t.Run("concurrent execution", func(t *testing.T) {
		urls := make([]string, 10)
		for i := range urls {
			urls[i] = fmt.Sprintf("http://%d.com", i)
		}

		var counter int64
		start := time.Now()
		ParallelFetch(urls, func(url string) (string, error) {
			atomic.AddInt64(&counter, 1)
			time.Sleep(50 * time.Millisecond)
			return "content", nil
		})
		elapsed := time.Since(start)

		if counter != 10 {
			t.Errorf("ParallelFetch() called fetch %d times, want 10", counter)
		}

		// Sequential would take 500ms
		if elapsed > 200*time.Millisecond {
			t.Errorf("ParallelFetch() took %v, expected concurrent execution", elapsed)
		}
	})

	t.Run("empty urls", func(t *testing.T) {
		results, errs := ParallelFetch(nil, func(url string) (string, error) {
			return "", errors.New("should not be called")
		})

		if len(errs) != 0 {
			t.Errorf("ParallelFetch(nil) got %d errors, want 0", len(errs))
		}

		if results == nil {
			// nil map is acceptable
		} else if len(results) != 0 {
			t.Errorf("ParallelFetch(nil) got %d results, want 0", len(results))
		}
	})
}

// Benchmark
func BenchmarkProcessWithErrors(b *testing.B) {
	items := make([]int, 100)
	for i := range items {
		items[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ProcessWithErrors(items, func(i int) error {
			return nil
		})
	}
}

func BenchmarkFirstError(b *testing.B) {
	tasks := make([]func() error, 10)
	for i := range tasks {
		tasks[i] = func() error { return nil }
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstError(tasks)
	}
}
