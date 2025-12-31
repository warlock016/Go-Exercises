package fan_out_fan_in

import (
	"sort"
	"sync"
	"testing"
	"time"
)

func TestFanOut(t *testing.T) {
	t.Run("distributes to multiple channels", func(t *testing.T) {
		input := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)

		outputs := FanOut(input, 2)
		if outputs == nil {
			t.Fatal("FanOut() returned nil")
		}
		if len(outputs) != 2 {
			t.Fatalf("FanOut() returned %d channels, want 2", len(outputs))
		}

		// Collect all values
		var all []int
		var wg sync.WaitGroup
		var mu sync.Mutex

		for _, ch := range outputs {
			wg.Add(1)
			go func(c <-chan int) {
				defer wg.Done()
				for v := range c {
					mu.Lock()
					all = append(all, v)
					mu.Unlock()
				}
			}(ch)
		}
		wg.Wait()

		sort.Ints(all)
		if len(all) != 5 {
			t.Errorf("Got %d values, want 5", len(all))
		}
		for i, v := range all {
			if v != i+1 {
				t.Errorf("Missing value %d", i+1)
			}
		}
	})

	t.Run("handles empty input", func(t *testing.T) {
		input := make(chan int)
		close(input)

		outputs := FanOut(input, 3)
		if len(outputs) != 3 {
			t.Fatalf("FanOut() returned %d channels, want 3", len(outputs))
		}

		// All outputs should close
		for i, ch := range outputs {
			select {
			case _, ok := <-ch:
				if ok {
					t.Errorf("Output %d should be closed", i)
				}
			case <-time.After(100 * time.Millisecond):
				t.Errorf("Output %d did not close", i)
			}
		}
	})
}

func TestFanIn(t *testing.T) {
	t.Run("merges multiple channels", func(t *testing.T) {
		ch1 := make(chan int, 3)
		ch2 := make(chan int, 3)

		ch1 <- 1
		ch1 <- 3
		ch1 <- 5
		close(ch1)

		ch2 <- 2
		ch2 <- 4
		ch2 <- 6
		close(ch2)

		merged := FanIn(ch1, ch2)
		if merged == nil {
			t.Fatal("FanIn() returned nil")
		}

		var all []int
		for v := range merged {
			all = append(all, v)
		}

		sort.Ints(all)
		if len(all) != 6 {
			t.Errorf("Got %d values, want 6", len(all))
		}
	})

	t.Run("handles empty channels", func(t *testing.T) {
		ch1 := make(chan int)
		ch2 := make(chan int)
		close(ch1)
		close(ch2)

		merged := FanIn(ch1, ch2)
		count := 0
		for range merged {
			count++
		}

		if count != 0 {
			t.Errorf("Got %d values from empty channels", count)
		}
	})

	t.Run("handles no channels", func(t *testing.T) {
		merged := FanIn[int]()
		if merged == nil {
			t.Fatal("FanIn() with no inputs returned nil")
		}

		select {
		case _, ok := <-merged:
			if ok {
				t.Error("Expected channel to be closed")
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Channel should close immediately")
		}
	})
}

func TestParallelMap(t *testing.T) {
	t.Run("transforms all items", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		results := ParallelMap(items, 3, func(x int) int {
			return x * 2
		})

		if len(results) != 5 {
			t.Fatalf("ParallelMap() returned %d results, want 5", len(results))
		}

		expected := []int{2, 4, 6, 8, 10}
		for i, r := range results {
			if r != expected[i] {
				t.Errorf("results[%d] = %d, want %d", i, r, expected[i])
			}
		}
	})

	t.Run("preserves order", func(t *testing.T) {
		items := []int{5, 4, 3, 2, 1}
		results := ParallelMap(items, 3, func(x int) int {
			// Longer delay for smaller numbers to test order
			time.Sleep(time.Duration(6-x) * 10 * time.Millisecond)
			return x * 10
		})

		expected := []int{50, 40, 30, 20, 10}
		for i, r := range results {
			if r != expected[i] {
				t.Errorf("Order not preserved: results[%d] = %d, want %d", i, r, expected[i])
			}
		}
	})

	t.Run("runs concurrently", func(t *testing.T) {
		items := make([]int, 10)
		for i := range items {
			items[i] = i
		}

		start := time.Now()
		ParallelMap(items, 5, func(x int) int {
			time.Sleep(50 * time.Millisecond)
			return x
		})
		elapsed := time.Since(start)

		// 10 items, 5 workers, 50ms each = ~100ms
		if elapsed > 200*time.Millisecond {
			t.Errorf("ParallelMap() took %v, expected concurrent execution", elapsed)
		}
	})

	t.Run("handles empty slice", func(t *testing.T) {
		results := ParallelMap([]int{}, 3, func(x int) int {
			return x
		})

		if results == nil {
			// nil is acceptable
		} else if len(results) != 0 {
			t.Errorf("ParallelMap([]) returned %d results", len(results))
		}
	})
}

func TestParallelMapStream(t *testing.T) {
	t.Run("processes all items", func(t *testing.T) {
		input := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)

		output := ParallelMapStream(input, 2, func(x int) int {
			return x * 2
		})
		if output == nil {
			t.Fatal("ParallelMapStream() returned nil")
		}

		var results []int
		for v := range output {
			results = append(results, v)
		}

		sort.Ints(results)
		expected := []int{2, 4, 6, 8, 10}
		for i, r := range results {
			if r != expected[i] {
				t.Errorf("results contain %d, want %d", r, expected[i])
			}
		}
	})

	t.Run("closes output when done", func(t *testing.T) {
		input := make(chan int, 2)
		input <- 1
		input <- 2
		close(input)

		output := ParallelMapStream(input, 2, func(x int) int { return x })

		// Drain and verify close
		for range output {
		}

		// Should be closed now
		select {
		case _, ok := <-output:
			if ok {
				t.Error("Channel should be closed")
			}
		default:
			// Already closed - ok
		}
	})
}

func TestProcessWithFanOut(t *testing.T) {
	t.Run("processes with fan-out/fan-in", func(t *testing.T) {
		input := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)

		output := ProcessWithFanOut(input, 3, func(x int) string {
			return string(rune('A' + x - 1))
		})
		if output == nil {
			t.Fatal("ProcessWithFanOut() returned nil")
		}

		var results []string
		for v := range output {
			results = append(results, v)
		}

		sort.Strings(results)
		expected := []string{"A", "B", "C", "D", "E"}
		for i, r := range results {
			if r != expected[i] {
				t.Errorf("results contain %q, want %q", r, expected[i])
			}
		}
	})

	t.Run("runs in parallel", func(t *testing.T) {
		input := make(chan int, 10)
		for i := 0; i < 10; i++ {
			input <- i
		}
		close(input)

		start := time.Now()
		output := ProcessWithFanOut(input, 5, func(x int) int {
			time.Sleep(50 * time.Millisecond)
			return x
		})

		// Drain
		for range output {
		}
		elapsed := time.Since(start)

		if elapsed > 200*time.Millisecond {
			t.Errorf("ProcessWithFanOut() took %v, expected parallel execution", elapsed)
		}
	})
}

// Benchmarks
func BenchmarkFanIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch1 := make(chan int, 100)
		ch2 := make(chan int, 100)

		go func() {
			for j := 0; j < 100; j++ {
				ch1 <- j
			}
			close(ch1)
		}()
		go func() {
			for j := 0; j < 100; j++ {
				ch2 <- j
			}
			close(ch2)
		}()

		merged := FanIn(ch1, ch2)
		for range merged {
		}
	}
}

func BenchmarkParallelMap(b *testing.B) {
	items := make([]int, 100)
	for i := range items {
		items[i] = i
	}

	b.Run("workers=1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ParallelMap(items, 1, func(x int) int { return x * 2 })
		}
	})

	b.Run("workers=4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ParallelMap(items, 4, func(x int) int { return x * 2 })
		}
	})
}
