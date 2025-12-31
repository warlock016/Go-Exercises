package done_channel

import (
	"sync"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	t.Run("processes all items", func(t *testing.T) {
		items := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			items <- i
		}
		close(items)

		done := make(chan struct{})
		// Don't close done - let worker process all items

		count := Worker(items, done)
		if count != 5 {
			t.Errorf("Worker() processed %d items, want 5", count)
		}
	})

	t.Run("stops when done closed", func(t *testing.T) {
		items := make(chan int, 10)
		for i := 1; i <= 10; i++ {
			items <- i
		}

		done := make(chan struct{})
		close(done) // Signal done immediately

		count := Worker(items, done)
		if count != 0 {
			t.Errorf("Worker() processed %d items after done, want 0", count)
		}
	})

	t.Run("handles empty channel", func(t *testing.T) {
		items := make(chan int)
		close(items)

		done := make(chan struct{})
		count := Worker(items, done)
		if count != 0 {
			t.Errorf("Worker() on empty channel = %d, want 0", count)
		}
	})
}

func TestCancellableLoop(t *testing.T) {
	t.Run("runs until done", func(t *testing.T) {
		var counter int
		done := make(chan struct{})

		go func() {
			time.Sleep(50 * time.Millisecond)
			close(done)
		}()

		count := CancellableLoop(func() {
			counter++
			time.Sleep(5 * time.Millisecond)
		}, done)

		if count < 5 {
			t.Errorf("CancellableLoop() ran %d times, expected at least 5", count)
		}
		if counter != count {
			t.Errorf("Counter mismatch: counter=%d, returned=%d", counter, count)
		}
	})

	t.Run("immediate cancel", func(t *testing.T) {
		done := make(chan struct{})
		close(done)

		count := CancellableLoop(func() {
			t.Error("Work should not run after done closed")
		}, done)

		if count != 0 {
			t.Errorf("CancellableLoop() with immediate done = %d, want 0", count)
		}
	})
}

func TestBroadcaster(t *testing.T) {
	t.Run("sends to all channels", func(t *testing.T) {
		outputs := make([]chan<- int, 3)
		receivers := make([]chan int, 3)
		for i := 0; i < 3; i++ {
			ch := make(chan int, 1)
			outputs[i] = ch
			receivers[i] = ch
		}

		done := make(chan struct{})
		count := Broadcaster(42, outputs, done)

		if count != 3 {
			t.Errorf("Broadcaster() sent to %d channels, want 3", count)
		}

		for i, ch := range receivers {
			select {
			case v := <-ch:
				if v != 42 {
					t.Errorf("Channel %d received %d, want 42", i, v)
				}
			default:
				t.Errorf("Channel %d did not receive value", i)
			}
		}
	})

	t.Run("respects done signal", func(t *testing.T) {
		// Use unbuffered channels that will block
		outputs := make([]chan<- int, 3)
		for i := 0; i < 3; i++ {
			outputs[i] = make(chan int) // unbuffered, will block
		}

		done := make(chan struct{})
		close(done)

		count := Broadcaster(42, outputs, done)
		if count != 0 {
			t.Errorf("Broadcaster() with done sent to %d channels, want 0", count)
		}
	})

	t.Run("empty outputs", func(t *testing.T) {
		done := make(chan struct{})
		count := Broadcaster(42, nil, done)
		if count != 0 {
			t.Errorf("Broadcaster() with nil outputs = %d, want 0", count)
		}
	})
}

func TestGenerator(t *testing.T) {
	t.Run("produces sequential integers", func(t *testing.T) {
		done := make(chan struct{})
		gen := Generator(done)
		if gen == nil {
			t.Fatal("Generator() returned nil")
		}

		var got []int
		for i := 0; i < 5; i++ {
			got = append(got, <-gen)
		}
		close(done)

		want := []int{0, 1, 2, 3, 4}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("Generator()[%d] = %d, want %d", i, v, want[i])
			}
		}
	})

	t.Run("closes channel when done", func(t *testing.T) {
		done := make(chan struct{})
		gen := Generator(done)

		// Read a few values
		<-gen
		<-gen

		// Close done
		close(done)

		// Channel should eventually close
		timeout := time.After(100 * time.Millisecond)
		for {
			select {
			case _, ok := <-gen:
				if !ok {
					return // Success - channel closed
				}
				// Keep draining
			case <-timeout:
				t.Error("Generator channel did not close after done signal")
				return
			}
		}
	})
}

func TestTimeout(t *testing.T) {
	t.Run("closes after duration", func(t *testing.T) {
		start := time.Now()
		done := Timeout(50 * time.Millisecond)
		if done == nil {
			t.Fatal("Timeout() returned nil")
		}

		<-done // Wait for close
		elapsed := time.Since(start)

		if elapsed < 40*time.Millisecond {
			t.Errorf("Timeout closed too early: %v", elapsed)
		}
		if elapsed > 100*time.Millisecond {
			t.Errorf("Timeout closed too late: %v", elapsed)
		}
	})

	t.Run("can be used with select", func(t *testing.T) {
		done := Timeout(50 * time.Millisecond)
		work := make(chan int)

		select {
		case <-work:
			t.Error("Received from work channel unexpectedly")
		case <-done:
			// Expected
		}
	})
}

func TestMergeCancel(t *testing.T) {
	t.Run("closes when first closes", func(t *testing.T) {
		done1 := make(chan struct{})
		done2 := make(chan struct{})

		merged := MergeCancel(done1, done2)
		if merged == nil {
			t.Fatal("MergeCancel() returned nil")
		}

		close(done1)

		select {
		case <-merged:
			// Success
		case <-time.After(100 * time.Millisecond):
			t.Error("MergeCancel did not close when done1 closed")
		}
	})

	t.Run("closes when second closes", func(t *testing.T) {
		done1 := make(chan struct{})
		done2 := make(chan struct{})

		merged := MergeCancel(done1, done2)

		close(done2)

		select {
		case <-merged:
			// Success
		case <-time.After(100 * time.Millisecond):
			t.Error("MergeCancel did not close when done2 closed")
		}
	})

	t.Run("handles both closed", func(t *testing.T) {
		done1 := make(chan struct{})
		done2 := make(chan struct{})
		close(done1)
		close(done2)

		merged := MergeCancel(done1, done2)

		select {
		case <-merged:
			// Success
		case <-time.After(100 * time.Millisecond):
			t.Error("MergeCancel did not close when both inputs closed")
		}
	})
}

func TestWorkerConcurrency(t *testing.T) {
	t.Run("multiple workers share done", func(t *testing.T) {
		items := make(chan int, 100)
		for i := 0; i < 100; i++ {
			items <- i
		}
		close(items)

		done := make(chan struct{})
		var wg sync.WaitGroup
		var total int
		var mu sync.Mutex

		// Start 3 workers
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				count := Worker(items, done)
				mu.Lock()
				total += count
				mu.Unlock()
			}()
		}

		wg.Wait()

		if total != 100 {
			t.Errorf("Workers processed %d items total, want 100", total)
		}
	})
}

// Benchmarks
func BenchmarkGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		done := make(chan struct{})
		gen := Generator(done)
		for j := 0; j < 100; j++ {
			<-gen
		}
		close(done)
	}
}
