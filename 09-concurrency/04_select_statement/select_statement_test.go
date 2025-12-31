package select_statement

import (
	"sort"
	"testing"
	"time"
)

func TestFirstResponse(t *testing.T) {
	t.Run("ch1 ready first", func(t *testing.T) {
		ch1 := make(chan string, 1)
		ch2 := make(chan string, 1)
		ch1 <- "first"

		got := FirstResponse(ch1, ch2)
		if got != "first" {
			t.Errorf("FirstResponse() = %q, want %q", got, "first")
		}
	})

	t.Run("ch2 ready first", func(t *testing.T) {
		ch1 := make(chan string, 1)
		ch2 := make(chan string, 1)
		ch2 <- "second"

		got := FirstResponse(ch1, ch2)
		if got != "second" {
			t.Errorf("FirstResponse() = %q, want %q", got, "second")
		}
	})

	t.Run("both ready", func(t *testing.T) {
		ch1 := make(chan string, 1)
		ch2 := make(chan string, 1)
		ch1 <- "one"
		ch2 <- "two"

		got := FirstResponse(ch1, ch2)
		if got != "one" && got != "two" {
			t.Errorf("FirstResponse() = %q, want 'one' or 'two'", got)
		}
	})
}

func TestWithTimeout(t *testing.T) {
	t.Run("value received", func(t *testing.T) {
		ch := make(chan string, 1)
		ch <- "hello"

		got, err := WithTimeout(ch, 100*time.Millisecond)
		if err != nil {
			t.Errorf("WithTimeout() error = %v, want nil", err)
		}
		if got != "hello" {
			t.Errorf("WithTimeout() = %q, want %q", got, "hello")
		}
	})

	t.Run("timeout occurs", func(t *testing.T) {
		ch := make(chan string)

		start := time.Now()
		_, err := WithTimeout(ch, 50*time.Millisecond)
		elapsed := time.Since(start)

		if err == nil {
			t.Error("WithTimeout() error = nil, want timeout error")
		}
		if elapsed < 40*time.Millisecond {
			t.Errorf("WithTimeout() returned too fast: %v", elapsed)
		}
		if elapsed > 100*time.Millisecond {
			t.Errorf("WithTimeout() took too long: %v", elapsed)
		}
	})

	t.Run("value arrives before timeout", func(t *testing.T) {
		ch := make(chan string)
		go func() {
			time.Sleep(20 * time.Millisecond)
			ch <- "delayed"
		}()

		got, err := WithTimeout(ch, 100*time.Millisecond)
		if err != nil {
			t.Errorf("WithTimeout() error = %v, want nil", err)
		}
		if got != "delayed" {
			t.Errorf("WithTimeout() = %q, want %q", got, "delayed")
		}
	})
}

func TestTryReceive(t *testing.T) {
	t.Run("empty channel", func(t *testing.T) {
		ch := make(chan int, 1)
		_, ok := TryReceive(ch)
		if ok {
			t.Error("TryReceive() on empty channel returned ok=true")
		}
	})

	t.Run("value available", func(t *testing.T) {
		ch := make(chan int, 1)
		ch <- 42

		v, ok := TryReceive(ch)
		if !ok {
			t.Error("TryReceive() with value returned ok=false")
		}
		if v != 42 {
			t.Errorf("TryReceive() = %d, want 42", v)
		}
	})

	t.Run("multiple values", func(t *testing.T) {
		ch := make(chan int, 3)
		ch <- 1
		ch <- 2
		ch <- 3

		for i := 1; i <= 3; i++ {
			v, ok := TryReceive(ch)
			if !ok {
				t.Errorf("TryReceive() %d: ok=false, want true", i)
			}
			if v != i {
				t.Errorf("TryReceive() %d = %d, want %d", i, v, i)
			}
		}

		// Now should be empty
		_, ok := TryReceive(ch)
		if ok {
			t.Error("TryReceive() on emptied channel returned ok=true")
		}
	})
}

func TestTrySend(t *testing.T) {
	t.Run("buffer available", func(t *testing.T) {
		ch := make(chan int, 1)
		ok := TrySend(ch, 42)
		if !ok {
			t.Error("TrySend() to empty buffered channel returned false")
		}

		v := <-ch
		if v != 42 {
			t.Errorf("Sent value = %d, want 42", v)
		}
	})

	t.Run("buffer full", func(t *testing.T) {
		ch := make(chan int, 1)
		ch <- 1 // Fill buffer

		ok := TrySend(ch, 2)
		if ok {
			t.Error("TrySend() to full channel returned true")
		}
	})

	t.Run("unbuffered no receiver", func(t *testing.T) {
		ch := make(chan int)
		ok := TrySend(ch, 1)
		if ok {
			t.Error("TrySend() to unbuffered channel with no receiver returned true")
		}
	})
}

func TestMerge(t *testing.T) {
	t.Run("both channels have values", func(t *testing.T) {
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

		merged := Merge(ch1, ch2)
		if merged == nil {
			t.Fatal("Merge() returned nil")
		}

		var got []int
		for v := range merged {
			got = append(got, v)
		}

		sort.Ints(got)
		want := []int{1, 2, 3, 4, 5, 6}
		if len(got) != len(want) {
			t.Errorf("Merge() produced %d values, want %d", len(got), len(want))
		}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("Merge() sorted[%d] = %d, want %d", i, v, want[i])
			}
		}
	})

	t.Run("one channel empty", func(t *testing.T) {
		ch1 := make(chan int, 2)
		ch2 := make(chan int)

		ch1 <- 1
		ch1 <- 2
		close(ch1)
		close(ch2)

		merged := Merge(ch1, ch2)
		var got []int
		for v := range merged {
			got = append(got, v)
		}

		if len(got) != 2 {
			t.Errorf("Merge() produced %d values, want 2", len(got))
		}
	})

	t.Run("both empty", func(t *testing.T) {
		ch1 := make(chan int)
		ch2 := make(chan int)
		close(ch1)
		close(ch2)

		merged := Merge(ch1, ch2)
		count := 0
		for range merged {
			count++
		}

		if count != 0 {
			t.Errorf("Merge() of empty channels produced %d values", count)
		}
	})
}

func TestMultiplex(t *testing.T) {
	t.Run("multiple channels", func(t *testing.T) {
		ch1 := make(chan int, 2)
		ch2 := make(chan int, 2)
		ch3 := make(chan int, 2)

		ch1 <- 10
		ch1 <- 11
		close(ch1)

		ch2 <- 20
		close(ch2)

		ch3 <- 30
		ch3 <- 31
		close(ch3)

		out := Multiplex(ch1, ch2, ch3)
		if out == nil {
			t.Fatal("Multiplex() returned nil")
		}

		results := make(map[int][]int)
		for iv := range out {
			results[iv.Index] = append(results[iv.Index], iv.Value)
		}

		// Check channel 0 (ch1)
		if len(results[0]) != 2 {
			t.Errorf("Channel 0 produced %d values, want 2", len(results[0]))
		}

		// Check channel 1 (ch2)
		if len(results[1]) != 1 {
			t.Errorf("Channel 1 produced %d values, want 1", len(results[1]))
		}

		// Check channel 2 (ch3)
		if len(results[2]) != 2 {
			t.Errorf("Channel 2 produced %d values, want 2", len(results[2]))
		}
	})

	t.Run("no channels", func(t *testing.T) {
		out := Multiplex()
		if out == nil {
			t.Fatal("Multiplex() with no channels returned nil")
		}

		count := 0
		for range out {
			count++
		}
		if count != 0 {
			t.Errorf("Multiplex() with no channels produced %d values", count)
		}
	})
}

// Benchmarks
func BenchmarkWithTimeout(b *testing.B) {
	ch := make(chan string, 1)
	for i := 0; i < b.N; i++ {
		ch <- "test"
		WithTimeout(ch, time.Second)
	}
}

func BenchmarkTryReceive(b *testing.B) {
	ch := make(chan int, 1)
	for i := 0; i < b.N; i++ {
		ch <- i
		TryReceive(ch)
	}
}
