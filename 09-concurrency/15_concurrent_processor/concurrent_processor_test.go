package concurrent_processor

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func defaultConfig() ProcessorConfig {
	return ProcessorConfig{
		Workers:       3,
		QueueSize:     10,
		RateLimit:     0, // unlimited for tests
		RetryDelay:    10 * time.Millisecond,
		MaxRetryDelay: 100 * time.Millisecond,
	}
}

func TestProcessor_BasicProcessing(t *testing.T) {
	var processed int64
	handler := func(ctx context.Context, job Job) error {
		atomic.AddInt64(&processed, 1)
		return nil
	}

	processor := NewProcessor(defaultConfig(), handler)
	if processor == nil {
		t.Fatal("NewProcessor() returned nil")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go processor.Start(ctx)

	// Submit jobs
	for i := 0; i < 10; i++ {
		err := processor.Submit(Job{ID: string(rune('A' + i))})
		if err != nil {
			t.Errorf("Submit() error = %v", err)
		}
	}

	// Wait for processing
	err := processor.Wait()
	if err != nil {
		t.Errorf("Wait() error = %v", err)
	}

	if processed != 10 {
		t.Errorf("Processed %d jobs, want 10", processed)
	}
}

func TestProcessor_Results(t *testing.T) {
	handler := func(ctx context.Context, job Job) error {
		time.Sleep(10 * time.Millisecond)
		return nil
	}

	processor := NewProcessor(defaultConfig(), handler)

	ctx, cancel := context.WithCancel(context.Background())

	go processor.Start(ctx)

	processor.Submit(Job{ID: "1"})
	processor.Submit(Job{ID: "2"})
	processor.Submit(Job{ID: "3"})

	results := processor.Results()
	if results == nil {
		t.Fatal("Results() returned nil")
	}

	var count int
	timeout := time.After(2 * time.Second)
	for count < 3 {
		select {
		case r, ok := <-results:
			if !ok {
				t.Fatal("Results channel closed early")
			}
			if !r.Success {
				t.Errorf("Job %s failed: %v", r.JobID, r.Error)
			}
			if r.Duration == 0 {
				t.Errorf("Job %s has zero duration", r.JobID)
			}
			count++
		case <-timeout:
			t.Fatalf("Timeout waiting for results, got %d/3", count)
		}
	}

	cancel()
}

func TestProcessor_Retry(t *testing.T) {
	var attempts int64
	handler := func(ctx context.Context, job Job) error {
		a := atomic.AddInt64(&attempts, 1)
		if a < 3 {
			return errors.New("temporary failure")
		}
		return nil
	}

	processor := NewProcessor(defaultConfig(), handler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go processor.Start(ctx)

	processor.Submit(Job{ID: "retry-job", MaxRetry: 3})

	// Collect result
	timeout := time.After(2 * time.Second)
	select {
	case r := <-processor.Results():
		if !r.Success {
			t.Errorf("Job should succeed after retries: %v", r.Error)
		}
	case <-timeout:
		t.Fatal("Timeout waiting for result")
	}

	// Check retries in metrics
	m := processor.Metrics()
	if m.Retried < 2 {
		t.Errorf("Metrics.Retried = %d, want >= 2", m.Retried)
	}
}

func TestProcessor_MaxRetryExceeded(t *testing.T) {
	handler := func(ctx context.Context, job Job) error {
		return errors.New("always fails")
	}

	processor := NewProcessor(defaultConfig(), handler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go processor.Start(ctx)

	processor.Submit(Job{ID: "always-fail", MaxRetry: 2})

	timeout := time.After(2 * time.Second)
	select {
	case r := <-processor.Results():
		if r.Success {
			t.Error("Job should fail after max retries")
		}
		if r.Error == nil {
			t.Error("Failed job should have error")
		}
	case <-timeout:
		t.Fatal("Timeout waiting for result")
	}

	m := processor.Metrics()
	if m.Failed != 1 {
		t.Errorf("Metrics.Failed = %d, want 1", m.Failed)
	}
}

func TestProcessor_Metrics(t *testing.T) {
	var wg sync.WaitGroup
	handler := func(ctx context.Context, job Job) error {
		wg.Done()
		time.Sleep(50 * time.Millisecond)
		return nil
	}

	processor := NewProcessor(defaultConfig(), handler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go processor.Start(ctx)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		processor.Submit(Job{ID: string(rune('A' + i))})
	}

	// Wait for jobs to start
	wg.Wait()

	m := processor.Metrics()
	if m.Submitted != 5 {
		t.Errorf("Metrics.Submitted = %d, want 5", m.Submitted)
	}
	if m.InFlight == 0 {
		t.Error("Metrics.InFlight should be > 0 while processing")
	}

	// Wait for completion
	processor.Wait()

	m = processor.Metrics()
	if m.Completed != 5 {
		t.Errorf("Metrics.Completed = %d, want 5", m.Completed)
	}
	if m.InFlight != 0 {
		t.Errorf("Metrics.InFlight = %d, want 0", m.InFlight)
	}
}

func TestProcessor_GracefulShutdown(t *testing.T) {
	var completed int64
	handler := func(ctx context.Context, job Job) error {
		time.Sleep(50 * time.Millisecond)
		atomic.AddInt64(&completed, 1)
		return nil
	}

	processor := NewProcessor(defaultConfig(), handler)

	ctx, cancel := context.WithCancel(context.Background())

	go processor.Start(ctx)

	// Submit jobs
	for i := 0; i < 5; i++ {
		processor.Submit(Job{ID: string(rune('A' + i))})
	}

	// Let some jobs start
	time.Sleep(30 * time.Millisecond)

	// Shutdown with enough time to complete
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer shutdownCancel()

	err := processor.Shutdown(shutdownCtx)
	if err != nil {
		t.Errorf("Shutdown() error = %v", err)
	}
	cancel()

	if completed != 5 {
		t.Errorf("Completed %d jobs, want 5 (graceful shutdown should complete in-flight)", completed)
	}
}

func TestProcessor_ShutdownTimeout(t *testing.T) {
	handler := func(ctx context.Context, job Job) error {
		select {
		case <-time.After(1 * time.Second):
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	processor := NewProcessor(defaultConfig(), handler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go processor.Start(ctx)

	processor.Submit(Job{ID: "slow"})

	// Let job start
	time.Sleep(30 * time.Millisecond)

	// Shutdown with short timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer shutdownCancel()

	err := processor.Shutdown(shutdownCtx)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("Shutdown() error = %v, want DeadlineExceeded", err)
	}
}

func TestProcessor_SubmitAfterShutdown(t *testing.T) {
	handler := func(ctx context.Context, job Job) error {
		return nil
	}

	processor := NewProcessor(defaultConfig(), handler)

	ctx, cancel := context.WithCancel(context.Background())

	go processor.Start(ctx)

	// Shutdown immediately
	processor.Shutdown(context.Background())
	cancel()

	// Submit after shutdown should fail
	err := processor.Submit(Job{ID: "late"})
	if err == nil {
		t.Error("Submit() after shutdown should return error")
	}
}

func TestProcessor_SubmitBatch(t *testing.T) {
	var processed int64
	handler := func(ctx context.Context, job Job) error {
		atomic.AddInt64(&processed, 1)
		return nil
	}

	processor := NewProcessor(defaultConfig(), handler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go processor.Start(ctx)

	jobs := []Job{
		{ID: "1"},
		{ID: "2"},
		{ID: "3"},
		{ID: "4"},
		{ID: "5"},
	}

	err := processor.SubmitBatch(jobs)
	if err != nil {
		t.Errorf("SubmitBatch() error = %v", err)
	}

	processor.Wait()

	if processed != 5 {
		t.Errorf("Processed %d jobs, want 5", processed)
	}
}

func TestProcessor_RateLimit(t *testing.T) {
	config := ProcessorConfig{
		Workers:   5,
		QueueSize: 20,
		RateLimit: 10, // 10 per second
	}

	handler := func(ctx context.Context, job Job) error {
		return nil
	}

	processor := NewProcessor(config, handler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go processor.Start(ctx)

	start := time.Now()

	// Submit 10 jobs
	for i := 0; i < 10; i++ {
		processor.Submit(Job{ID: string(rune('A' + i))})
	}

	processor.Wait()
	elapsed := time.Since(start)

	// At 10/sec, 10 jobs should take ~1 second
	if elapsed < 800*time.Millisecond {
		t.Errorf("Rate limiting too fast: %v (expected ~1s)", elapsed)
	}
}

func TestProcessor_ConcurrentSubmit(t *testing.T) {
	var processed int64
	handler := func(ctx context.Context, job Job) error {
		atomic.AddInt64(&processed, 1)
		return nil
	}

	processor := NewProcessor(defaultConfig(), handler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go processor.Start(ctx)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				processor.Submit(Job{ID: string(rune(id*10 + j))})
			}
		}(i)
	}

	wg.Wait()
	processor.Wait()

	if processed != 100 {
		t.Errorf("Processed %d jobs, want 100", processed)
	}
}

// Benchmarks
func BenchmarkProcessor_Submit(b *testing.B) {
	handler := func(ctx context.Context, job Job) error {
		return nil
	}

	processor := NewProcessor(ProcessorConfig{
		Workers:   10,
		QueueSize: 1000,
	}, handler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go processor.Start(ctx)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		processor.Submit(Job{ID: "bench"})
	}

	b.StopTimer()
	processor.Wait()
}

func BenchmarkProcessor_Throughput(b *testing.B) {
	handler := func(ctx context.Context, job Job) error {
		return nil
	}

	processor := NewProcessor(ProcessorConfig{
		Workers:   10,
		QueueSize: 1000,
	}, handler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go processor.Start(ctx)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		processor.Submit(Job{ID: "bench"})
	}

	processor.Wait()
}
