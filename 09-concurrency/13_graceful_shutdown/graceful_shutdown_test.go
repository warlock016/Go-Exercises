package graceful_shutdown

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("processes requests", func(t *testing.T) {
		var processed int64
		server := NewServer(func(req int) error {
			atomic.AddInt64(&processed, 1)
			return nil
		})
		if server == nil {
			t.Fatal("NewServer() returned nil")
		}

		requests := make(chan int, 5)
		server.Start(requests)

		for i := 0; i < 5; i++ {
			requests <- i
		}
		close(requests)

		time.Sleep(100 * time.Millisecond)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		server.Shutdown(ctx)

		if processed != 5 {
			t.Errorf("Processed %d requests, want 5", processed)
		}
	})

	t.Run("completes in-flight on shutdown", func(t *testing.T) {
		var completed int64
		server := NewServer(func(req int) error {
			time.Sleep(100 * time.Millisecond)
			atomic.AddInt64(&completed, 1)
			return nil
		})

		requests := make(chan int, 3)
		server.Start(requests)

		// Send requests
		requests <- 1
		requests <- 2
		requests <- 3
		close(requests)

		// Wait a bit for processing to start
		time.Sleep(50 * time.Millisecond)

		// Shutdown with enough time
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			t.Errorf("Shutdown() error = %v", err)
		}

		if completed != 3 {
			t.Errorf("Completed %d requests, want 3", completed)
		}
	})

	t.Run("shutdown times out", func(t *testing.T) {
		server := NewServer(func(req int) error {
			time.Sleep(1 * time.Second)
			return nil
		})

		requests := make(chan int, 1)
		server.Start(requests)
		requests <- 1
		close(requests)

		time.Sleep(50 * time.Millisecond)

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		err := server.Shutdown(ctx)
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("Shutdown() error = %v, want DeadlineExceeded", err)
		}
	})
}

func TestGracefulShutdown(t *testing.T) {
	t.Run("shuts down all components", func(t *testing.T) {
		var shutdown1, shutdown2 bool

		comp1 := &mockShutdownable{
			shutdownFn: func(ctx context.Context) error {
				shutdown1 = true
				return nil
			},
		}
		comp2 := &mockShutdownable{
			shutdownFn: func(ctx context.Context) error {
				shutdown2 = true
				return nil
			},
		}

		ctx := context.Background()
		err := GracefulShutdown(ctx, comp1, comp2)

		if err != nil {
			t.Errorf("GracefulShutdown() error = %v", err)
		}
		if !shutdown1 || !shutdown2 {
			t.Error("Not all components were shutdown")
		}
	})

	t.Run("returns first error", func(t *testing.T) {
		expectedErr := errors.New("shutdown failed")

		comp1 := &mockShutdownable{
			shutdownFn: func(ctx context.Context) error {
				return nil
			},
		}
		comp2 := &mockShutdownable{
			shutdownFn: func(ctx context.Context) error {
				return expectedErr
			},
		}

		ctx := context.Background()
		err := GracefulShutdown(ctx, comp1, comp2)

		if err == nil {
			t.Error("GracefulShutdown() should return error")
		}
	})

	t.Run("respects context timeout", func(t *testing.T) {
		comp := &mockShutdownable{
			shutdownFn: func(ctx context.Context) error {
				<-ctx.Done()
				return ctx.Err()
			},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()

		start := time.Now()
		err := GracefulShutdown(ctx, comp)
		elapsed := time.Since(start)

		if elapsed > 100*time.Millisecond {
			t.Errorf("GracefulShutdown() took %v, should respect timeout", elapsed)
		}
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("GracefulShutdown() error = %v, want DeadlineExceeded", err)
		}
	})
}

func TestDrainer(t *testing.T) {
	t.Run("drains all items", func(t *testing.T) {
		ch := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)

		var sum int
		ctx := context.Background()
		err := Drainer(ctx, ch, func(v int) error {
			sum += v
			return nil
		})

		if err != nil {
			t.Errorf("Drainer() error = %v", err)
		}
		if sum != 15 {
			t.Errorf("sum = %d, want 15", sum)
		}
	})

	t.Run("stops on context cancel", func(t *testing.T) {
		ch := make(chan int)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := Drainer(ctx, ch, func(v int) error {
			return nil
		})

		if !errors.Is(err, context.Canceled) {
			t.Errorf("Drainer() error = %v, want Canceled", err)
		}
	})

	t.Run("propagates process error", func(t *testing.T) {
		ch := make(chan int, 3)
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)

		expectedErr := errors.New("process failed")
		ctx := context.Background()
		err := Drainer(ctx, ch, func(v int) error {
			if v == 2 {
				return expectedErr
			}
			return nil
		})

		if err != expectedErr {
			t.Errorf("Drainer() error = %v, want %v", err, expectedErr)
		}
	})
}

func TestCoordinator(t *testing.T) {
	t.Run("runs all workers", func(t *testing.T) {
		coord := NewCoordinator()
		if coord == nil {
			t.Fatal("NewCoordinator() returned nil")
		}

		var ran1, ran2 bool
		var mu sync.Mutex

		coord.Add("worker1", func(ctx context.Context) error {
			mu.Lock()
			ran1 = true
			mu.Unlock()
			<-ctx.Done()
			return nil
		})
		coord.Add("worker2", func(ctx context.Context) error {
			mu.Lock()
			ran2 = true
			mu.Unlock()
			<-ctx.Done()
			return nil
		})

		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			time.Sleep(100 * time.Millisecond)
			cancel()
		}()

		coord.Run(ctx)

		mu.Lock()
		if !ran1 || !ran2 {
			t.Error("Not all workers ran")
		}
		mu.Unlock()
	})

	t.Run("shuts down gracefully", func(t *testing.T) {
		coord := NewCoordinator()
		var cleanedUp int64

		coord.Add("worker", func(ctx context.Context) error {
			<-ctx.Done()
			atomic.AddInt64(&cleanedUp, 1)
			return nil
		})

		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			time.Sleep(50 * time.Millisecond)
			coord.Shutdown(context.Background())
		}()

		go func() {
			time.Sleep(200 * time.Millisecond)
			cancel()
		}()

		coord.Run(ctx)

		if cleanedUp != 1 {
			t.Errorf("cleanedUp = %d, want 1", cleanedUp)
		}
	})

	t.Run("handles worker errors", func(t *testing.T) {
		coord := NewCoordinator()
		expectedErr := errors.New("worker failed")

		coord.Add("failing-worker", func(ctx context.Context) error {
			return expectedErr
		})

		ctx := context.Background()
		err := coord.Run(ctx)

		if err != expectedErr {
			t.Errorf("Coordinator.Run() error = %v, want %v", err, expectedErr)
		}
	})
}

// Mock implementation for testing
type mockShutdownable struct {
	shutdownFn func(ctx context.Context) error
}

func (m *mockShutdownable) Shutdown(ctx context.Context) error {
	return m.shutdownFn(ctx)
}

// Benchmarks
func BenchmarkDrainer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 100)
		for j := 0; j < 100; j++ {
			ch <- j
		}
		close(ch)

		ctx := context.Background()
		Drainer(ctx, ch, func(v int) error { return nil })
	}
}
