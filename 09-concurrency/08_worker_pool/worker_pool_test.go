package worker_pool

import (
	"errors"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestProcessJobs(t *testing.T) {
	t.Run("basic processing", func(t *testing.T) {
		jobs := []Job{
			{ID: 1, Input: 10},
			{ID: 2, Input: 20},
			{ID: 3, Input: 30},
		}

		results := ProcessJobs(jobs, 2, func(j Job) Result {
			return Result{JobID: j.ID, Output: j.Input * 2}
		})

		if len(results) != 3 {
			t.Fatalf("ProcessJobs() returned %d results, want 3", len(results))
		}

		// Sort by JobID to check results
		sort.Slice(results, func(i, j int) bool {
			return results[i].JobID < results[j].JobID
		})

		expected := []int{20, 40, 60}
		for i, r := range results {
			if r.Output != expected[i] {
				t.Errorf("Result for job %d: Output = %d, want %d", r.JobID, r.Output, expected[i])
			}
		}
	})

	t.Run("empty jobs", func(t *testing.T) {
		results := ProcessJobs(nil, 2, func(j Job) Result {
			return Result{}
		})

		if results == nil {
			// nil is acceptable
		} else if len(results) != 0 {
			t.Errorf("ProcessJobs(nil) returned %d results, want 0", len(results))
		}
	})

	t.Run("with errors", func(t *testing.T) {
		jobs := []Job{
			{ID: 1, Input: 10},
			{ID: 2, Input: -1}, // Will cause error
			{ID: 3, Input: 30},
		}

		results := ProcessJobs(jobs, 2, func(j Job) Result {
			if j.Input < 0 {
				return Result{JobID: j.ID, Err: errors.New("negative input")}
			}
			return Result{JobID: j.ID, Output: j.Input * 2}
		})

		if len(results) != 3 {
			t.Fatalf("ProcessJobs() returned %d results, want 3", len(results))
		}

		var errorCount int
		for _, r := range results {
			if r.Err != nil {
				errorCount++
			}
		}
		if errorCount != 1 {
			t.Errorf("Got %d errors, want 1", errorCount)
		}
	})

	t.Run("concurrent execution", func(t *testing.T) {
		jobs := make([]Job, 20)
		for i := range jobs {
			jobs[i] = Job{ID: i, Input: i}
		}

		start := time.Now()
		ProcessJobs(jobs, 10, func(j Job) Result {
			time.Sleep(50 * time.Millisecond)
			return Result{JobID: j.ID, Output: j.Input}
		})
		elapsed := time.Since(start)

		// With 10 workers, 20 jobs at 50ms each should take ~100ms
		// Sequential would take 1000ms
		if elapsed > 300*time.Millisecond {
			t.Errorf("ProcessJobs() took %v, expected concurrent execution", elapsed)
		}
	})
}

func TestWorkerPool(t *testing.T) {
	t.Run("basic usage", func(t *testing.T) {
		pool := NewWorkerPool(3, func(j Job) Result {
			return Result{JobID: j.ID, Output: j.Input * 2}
		})
		if pool == nil {
			t.Fatal("NewWorkerPool() returned nil")
		}

		// Submit jobs
		for i := 0; i < 5; i++ {
			pool.Submit(Job{ID: i, Input: i * 10})
		}
		pool.Shutdown()

		// Collect results
		var results []Result
		for r := range pool.Results() {
			results = append(results, r)
		}

		if len(results) != 5 {
			t.Errorf("Got %d results, want 5", len(results))
		}
	})

	t.Run("results channel closes after shutdown", func(t *testing.T) {
		pool := NewWorkerPool(2, func(j Job) Result {
			return Result{JobID: j.ID}
		})

		pool.Submit(Job{ID: 1})
		pool.Shutdown()

		// Results channel should close
		timeout := time.After(1 * time.Second)
		for {
			select {
			case _, ok := <-pool.Results():
				if !ok {
					return // Success
				}
			case <-timeout:
				t.Fatal("Results channel did not close after shutdown")
			}
		}
	})

	t.Run("concurrent submits", func(t *testing.T) {
		pool := NewWorkerPool(5, func(j Job) Result {
			return Result{JobID: j.ID, Output: j.Input}
		})

		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				pool.Submit(Job{ID: id, Input: id})
			}(i)
		}

		go func() {
			wg.Wait()
			pool.Shutdown()
		}()

		var count int
		for range pool.Results() {
			count++
		}

		if count != 10 {
			t.Errorf("Got %d results, want 10", count)
		}
	})

	t.Run("worker count limits concurrency", func(t *testing.T) {
		var concurrent int64
		var maxConcurrent int64

		pool := NewWorkerPool(3, func(j Job) Result {
			c := atomic.AddInt64(&concurrent, 1)
			// Track max
			for {
				max := atomic.LoadInt64(&maxConcurrent)
				if c <= max || atomic.CompareAndSwapInt64(&maxConcurrent, max, c) {
					break
				}
			}
			time.Sleep(50 * time.Millisecond)
			atomic.AddInt64(&concurrent, -1)
			return Result{JobID: j.ID}
		})

		for i := 0; i < 10; i++ {
			pool.Submit(Job{ID: i})
		}
		pool.Shutdown()

		// Drain results
		for range pool.Results() {
		}

		if maxConcurrent > 3 {
			t.Errorf("Max concurrent = %d, want <= 3", maxConcurrent)
		}
	})
}

func TestWorkerPoolErrors(t *testing.T) {
	t.Run("handles job errors", func(t *testing.T) {
		pool := NewWorkerPool(2, func(j Job) Result {
			if j.Input < 0 {
				return Result{JobID: j.ID, Err: errors.New("negative")}
			}
			return Result{JobID: j.ID, Output: j.Input}
		})

		pool.Submit(Job{ID: 1, Input: 10})
		pool.Submit(Job{ID: 2, Input: -5})
		pool.Submit(Job{ID: 3, Input: 30})
		pool.Shutdown()

		var errorCount int
		for r := range pool.Results() {
			if r.Err != nil {
				errorCount++
			}
		}

		if errorCount != 1 {
			t.Errorf("Got %d errors, want 1", errorCount)
		}
	})
}

// Benchmarks
func BenchmarkProcessJobs(b *testing.B) {
	jobs := make([]Job, 100)
	for i := range jobs {
		jobs[i] = Job{ID: i, Input: i}
	}

	b.Run("workers=1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ProcessJobs(jobs, 1, func(j Job) Result {
				return Result{JobID: j.ID, Output: j.Input * 2}
			})
		}
	})

	b.Run("workers=4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ProcessJobs(jobs, 4, func(j Job) Result {
				return Result{JobID: j.ID, Output: j.Input * 2}
			})
		}
	})

	b.Run("workers=10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ProcessJobs(jobs, 10, func(j Job) Result {
				return Result{JobID: j.ID, Output: j.Input * 2}
			})
		}
	})
}
