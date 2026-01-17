package worker_pool_patterns

import (
	"context"
	"errors"
	"sort"
	"sync/atomic"
	"testing"
	"time"
)

// =============================================================================
// StreamingPool Tests
// =============================================================================

func TestStreamingPool(t *testing.T) {
	t.Run("basic processing", func(t *testing.T) {
		// Setup: Create jobs channel and send jobs concurrently
		jobs := make(chan Job)
		go func() {
			for i := 1; i <= 5; i++ {
				jobs <- Job{ID: i, Data: i * 10}
			}
			close(jobs)
		}()

		// Process
		results := StreamingPool(jobs, 2, func(j Job) Result {
			return Result{JobID: j.ID, Output: j.Data * 2}
		})

		// Collect results (consumer runs concurrently with workers)
		var collected []Result
		for r := range results {
			collected = append(collected, r)
		}

		if len(collected) != 5 {
			t.Errorf("got %d results, want 5", len(collected))
		}

		// Verify all jobs processed (order may vary)
		sort.Slice(collected, func(i, j int) bool {
			return collected[i].JobID < collected[j].JobID
		})
		expected := []int{20, 40, 60, 80, 100}
		for i, r := range collected {
			if r.Output != expected[i] {
				t.Errorf("job %d: got output %d, want %d", r.JobID, r.Output, expected[i])
			}
		}
	})

	t.Run("empty input", func(t *testing.T) {
		jobs := make(chan Job)
		close(jobs) // Immediately closed

		results := StreamingPool(jobs, 2, func(j Job) Result {
			return Result{JobID: j.ID}
		})

		count := 0
		for range results {
			count++
		}

		if count != 0 {
			t.Errorf("got %d results for empty input, want 0", count)
		}
	})

	t.Run("single worker", func(t *testing.T) {
		jobs := make(chan Job)
		go func() {
			for i := 0; i < 10; i++ {
				jobs <- Job{ID: i, Data: i}
			}
			close(jobs)
		}()

		results := StreamingPool(jobs, 1, func(j Job) Result {
			return Result{JobID: j.ID, Output: j.Data + 1}
		})

		count := 0
		for range results {
			count++
		}

		if count != 10 {
			t.Errorf("got %d results, want 10", count)
		}
	})

	t.Run("backpressure self-regulates", func(t *testing.T) {
		// KEY TEST: Demonstrates that streaming pattern works with SMALL buffers
		// because producer and consumer run concurrently.
		//
		// This test processes 100 jobs with only 3 workers.
		// With proper backpressure, this works without large buffers.

		jobs := make(chan Job)
		go func() {
			for i := 0; i < 100; i++ {
				jobs <- Job{ID: i, Data: i}
			}
			close(jobs)
		}()

		var processed int64
		results := StreamingPool(jobs, 3, func(j Job) Result {
			atomic.AddInt64(&processed, 1)
			return Result{JobID: j.ID, Output: j.Data}
		})

		// Consumer drains results (this is what enables backpressure to work)
		count := 0
		for range results {
			count++
		}

		if count != 100 {
			t.Errorf("got %d results, want 100", count)
		}
		if atomic.LoadInt64(&processed) != 100 {
			t.Errorf("processed %d jobs, want 100", processed)
		}
	})

	t.Run("concurrent execution", func(t *testing.T) {
		// Verify workers run concurrently by checking timing
		jobs := make(chan Job)
		go func() {
			for i := 0; i < 10; i++ {
				jobs <- Job{ID: i}
			}
			close(jobs)
		}()

		start := time.Now()
		results := StreamingPool(jobs, 5, func(j Job) Result {
			time.Sleep(50 * time.Millisecond)
			return Result{JobID: j.ID}
		})

		for range results {
		}
		elapsed := time.Since(start)

		// 10 jobs, 5 workers, 50ms each = ~100ms concurrent, 500ms sequential
		if elapsed > 250*time.Millisecond {
			t.Errorf("took %v, expected concurrent execution (~100ms)", elapsed)
		}
	})
}

// =============================================================================
// ProcessBatch Tests
// =============================================================================

func TestProcessBatch(t *testing.T) {
	t.Run("basic batch", func(t *testing.T) {
		jobs := []Job{
			{ID: 1, Data: 10},
			{ID: 2, Data: 20},
			{ID: 3, Data: 30},
		}

		results := ProcessBatch(jobs, 2, func(j Job) Result {
			return Result{JobID: j.ID, Output: j.Data * 2}
		})

		if len(results) != 3 {
			t.Fatalf("got %d results, want 3", len(results))
		}

		// Sort for comparison
		sort.Slice(results, func(i, j int) bool {
			return results[i].JobID < results[j].JobID
		})

		expected := []int{20, 40, 60}
		for i, r := range results {
			if r.Output != expected[i] {
				t.Errorf("job %d: got %d, want %d", r.JobID, r.Output, expected[i])
			}
		}
	})

	t.Run("empty batch", func(t *testing.T) {
		results := ProcessBatch(nil, 2, func(j Job) Result {
			return Result{}
		})

		if results != nil && len(results) != 0 {
			t.Errorf("got %d results for empty batch, want 0", len(results))
		}
	})

	t.Run("large batch", func(t *testing.T) {
		// ProcessBatch should handle large batches correctly
		// (internally it uses the streaming pattern)
		jobs := make([]Job, 100)
		for i := range jobs {
			jobs[i] = Job{ID: i, Data: i}
		}

		results := ProcessBatch(jobs, 4, func(j Job) Result {
			return Result{JobID: j.ID, Output: j.Data + 1}
		})

		if len(results) != 100 {
			t.Errorf("got %d results, want 100", len(results))
		}
	})

	t.Run("with errors in results", func(t *testing.T) {
		jobs := []Job{
			{ID: 1, Data: 10},
			{ID: 2, Data: -1}, // Will cause "error"
			{ID: 3, Data: 30},
		}

		results := ProcessBatch(jobs, 2, func(j Job) Result {
			if j.Data < 0 {
				return Result{JobID: j.ID, Err: errors.New("negative data")}
			}
			return Result{JobID: j.ID, Output: j.Data * 2}
		})

		if len(results) != 3 {
			t.Fatalf("got %d results, want 3", len(results))
		}

		errorCount := 0
		for _, r := range results {
			if r.Err != nil {
				errorCount++
			}
		}
		if errorCount != 1 {
			t.Errorf("got %d errors, want 1", errorCount)
		}
	})
}

// =============================================================================
// PoolWithCancel Tests
// =============================================================================

func TestPoolWithCancel(t *testing.T) {
	t.Run("normal completion", func(t *testing.T) {
		ctx := context.Background()
		jobs := make(chan Job)
		go func() {
			for i := 0; i < 5; i++ {
				jobs <- Job{ID: i}
			}
			close(jobs)
		}()

		results := PoolWithCancel(ctx, jobs, 2, func(j Job) Result {
			return Result{JobID: j.ID}
		})

		count := 0
		for range results {
			count++
		}

		if count != 5 {
			t.Errorf("got %d results, want 5", count)
		}
	})

	t.Run("cancellation stops accepting new jobs", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		jobs := make(chan Job, 100) // Buffered so we can send without blocking
		for i := 0; i < 20; i++ {
			jobs <- Job{ID: i}
		}
		// Note: NOT closing jobs channel

		results := PoolWithCancel(ctx, jobs, 2, func(j Job) Result {
			time.Sleep(10 * time.Millisecond)
			return Result{JobID: j.ID}
		})

		// Let some jobs process
		time.Sleep(50 * time.Millisecond)
		cancel() // Cancel context

		// Collect remaining results
		var collected []Result
		for r := range results {
			collected = append(collected, r)
		}

		// Should have processed SOME but not ALL jobs
		if len(collected) >= 20 {
			t.Errorf("processed %d jobs, expected fewer due to cancellation", len(collected))
		}
		if len(collected) == 0 {
			t.Error("processed 0 jobs, expected some before cancellation")
		}
	})

	t.Run("in-flight jobs complete", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		jobs := make(chan Job, 5)
		for i := 0; i < 5; i++ {
			jobs <- Job{ID: i}
		}
		close(jobs)

		var completedAfterCancel int64
		results := PoolWithCancel(ctx, jobs, 2, func(j Job) Result {
			// Cancel while job is processing
			if j.ID == 0 {
				cancel()
			}
			time.Sleep(20 * time.Millisecond)
			atomic.AddInt64(&completedAfterCancel, 1)
			return Result{JobID: j.ID}
		})

		for range results {
		}

		// Jobs in-flight when cancelled should still complete
		if atomic.LoadInt64(&completedAfterCancel) == 0 {
			t.Error("no jobs completed, expected in-flight jobs to finish")
		}
	})
}

// =============================================================================
// PoolWithErrors Tests
// =============================================================================

func TestPoolWithErrors(t *testing.T) {
	t.Run("separates results and errors", func(t *testing.T) {
		jobs := make(chan Job)
		go func() {
			jobs <- Job{ID: 1, Data: 10}
			jobs <- Job{ID: 2, Data: -1} // Error
			jobs <- Job{ID: 3, Data: 30}
			jobs <- Job{ID: 4, Data: -2} // Error
			jobs <- Job{ID: 5, Data: 50}
			close(jobs)
		}()

		results, errs := PoolWithErrors(jobs, 2, func(j Job) (Result, error) {
			if j.Data < 0 {
				return Result{}, errors.New("negative data")
			}
			return Result{JobID: j.ID, Output: j.Data * 2}, nil
		})

		// Must drain both channels
		var resultCount, errorCount int
		resultsDone := make(chan bool)
		errorsDone := make(chan bool)

		go func() {
			for range results {
				resultCount++
			}
			resultsDone <- true
		}()

		go func() {
			for range errs {
				errorCount++
			}
			errorsDone <- true
		}()

		<-resultsDone
		<-errorsDone

		if resultCount != 3 {
			t.Errorf("got %d results, want 3", resultCount)
		}
		if errorCount != 2 {
			t.Errorf("got %d errors, want 2", errorCount)
		}
	})

	t.Run("all success", func(t *testing.T) {
		jobs := make(chan Job)
		go func() {
			for i := 0; i < 5; i++ {
				jobs <- Job{ID: i, Data: i}
			}
			close(jobs)
		}()

		results, errs := PoolWithErrors(jobs, 2, func(j Job) (Result, error) {
			return Result{JobID: j.ID, Output: j.Data + 1}, nil
		})

		resultCount := 0
		errorCount := 0

		done := make(chan bool)
		go func() {
			for range errs {
				errorCount++
			}
			done <- true
		}()

		for range results {
			resultCount++
		}
		<-done

		if resultCount != 5 {
			t.Errorf("got %d results, want 5", resultCount)
		}
		if errorCount != 0 {
			t.Errorf("got %d errors, want 0", errorCount)
		}
	})

	t.Run("all errors", func(t *testing.T) {
		jobs := make(chan Job)
		go func() {
			for i := 0; i < 3; i++ {
				jobs <- Job{ID: i}
			}
			close(jobs)
		}()

		results, errs := PoolWithErrors(jobs, 2, func(j Job) (Result, error) {
			return Result{}, errors.New("always fails")
		})

		resultCount := 0
		errorCount := 0

		done := make(chan bool)
		go func() {
			for range results {
				resultCount++
			}
			done <- true
		}()

		for range errs {
			errorCount++
		}
		<-done

		if resultCount != 0 {
			t.Errorf("got %d results, want 0", resultCount)
		}
		if errorCount != 3 {
			t.Errorf("got %d errors, want 3", errorCount)
		}
	})

	t.Run("empty input", func(t *testing.T) {
		jobs := make(chan Job)
		close(jobs)

		results, errs := PoolWithErrors(jobs, 2, func(j Job) (Result, error) {
			return Result{}, nil
		})

		resultCount := 0
		errorCount := 0

		done := make(chan bool)
		go func() {
			for range errs {
				errorCount++
			}
			done <- true
		}()

		for range results {
			resultCount++
		}
		<-done

		if resultCount != 0 || errorCount != 0 {
			t.Errorf("got results=%d, errors=%d; want both 0", resultCount, errorCount)
		}
	})
}

// =============================================================================
// Benchmarks
// =============================================================================

func BenchmarkStreamingPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jobs := make(chan Job)
		go func() {
			for j := 0; j < 100; j++ {
				jobs <- Job{ID: j, Data: j}
			}
			close(jobs)
		}()

		results := StreamingPool(jobs, 4, func(j Job) Result {
			return Result{JobID: j.ID, Output: j.Data * 2}
		})

		for range results {
		}
	}
}

func BenchmarkProcessBatch(b *testing.B) {
	jobs := make([]Job, 100)
	for i := range jobs {
		jobs[i] = Job{ID: i, Data: i}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ProcessBatch(jobs, 4, func(j Job) Result {
			return Result{JobID: j.ID, Output: j.Data * 2}
		})
	}
}
