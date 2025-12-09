package closure_practice

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestMemoize(t *testing.T) {
	tests := []struct {
		name       string
		inputs     []int
		wantCalls  int
		wantResult int
	}{
		{
			name:       "caches repeated calls",
			inputs:     []int{5, 5, 5},
			wantCalls:  1,
			wantResult: 25,
		},
		{
			name:       "computes for different inputs",
			inputs:     []int{2, 3, 4},
			wantCalls:  3,
			wantResult: 16,
		},
		{
			name:       "mixed repeated and new",
			inputs:     []int{2, 2, 3, 2, 3},
			wantCalls:  2,
			wantResult: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			square := func(n int) int {
				callCount++
				return n * n
			}

			memoized := Memoize(square)
			var result int
			for _, input := range tt.inputs {
				result = memoized(input)
			}

			if callCount != tt.wantCalls {
				t.Errorf("call count = %d, want %d", callCount, tt.wantCalls)
			}
			if result != tt.wantResult {
				t.Errorf("last result = %d, want %d", result, tt.wantResult)
			}
		})
	}
}

func TestToggle(t *testing.T) {
	t.Run("string toggle", func(t *testing.T) {
		toggle := Toggle("on", "off")
		wants := []string{"on", "off", "on", "off", "on"}

		for i, want := range wants {
			if got := toggle(); got != want {
				t.Errorf("call %d = %q, want %q", i+1, got, want)
			}
		}
	})

	t.Run("bool toggle", func(t *testing.T) {
		toggle := Toggle(true, false)
		wants := []bool{true, false, true, false}

		for i, want := range wants {
			if got := toggle(); got != want {
				t.Errorf("call %d = %v, want %v", i+1, got, want)
			}
		}
	})

	t.Run("int toggle", func(t *testing.T) {
		toggle := Toggle(1, 0)
		wants := []int{1, 0, 1, 0}

		for i, want := range wants {
			if got := toggle(); got != want {
				t.Errorf("call %d = %d, want %d", i+1, got, want)
			}
		}
	})

	t.Run("independent toggles", func(t *testing.T) {
		t1 := Toggle("a", "b")
		t2 := Toggle("x", "y")

		t1() // a
		t1() // b

		if got := t2(); got != "x" {
			t.Errorf("t2 first call = %q, want 'x' (should be independent)", got)
		}
	})
}

func TestOnce(t *testing.T) {
	t.Run("calls function only once", func(t *testing.T) {
		callCount := 0
		init := func() string {
			callCount++
			return "initialized"
		}

		once := Once(init)

		for i := 0; i < 5; i++ {
			result := once()
			if result != "initialized" {
				t.Errorf("call %d = %q, want 'initialized'", i+1, result)
			}
		}

		if callCount != 1 {
			t.Errorf("function called %d times, want 1", callCount)
		}
	})

	t.Run("returns same value", func(t *testing.T) {
		counter := 0
		fn := func() int {
			counter++
			return counter
		}

		once := Once(fn)
		first := once()
		second := once()
		third := once()

		if first != 1 || second != 1 || third != 1 {
			t.Errorf("got %d, %d, %d - want all to be 1", first, second, third)
		}
	})

	t.Run("independent once functions", func(t *testing.T) {
		fn := func() int { return 42 }

		once1 := Once(fn)
		once2 := Once(fn)

		r1 := once1()
		r2 := once2()

		if r1 != 42 || r2 != 42 {
			t.Errorf("independent once functions should both work")
		}
	})
}

func TestDebounce(t *testing.T) {
	t.Run("executes after delay", func(t *testing.T) {
		executed := false
		fn := func() { executed = true }

		debounced, _ := Debounce(fn, 50*time.Millisecond)
		debounced()

		if executed {
			t.Error("should not execute immediately")
		}

		time.Sleep(100 * time.Millisecond)

		if !executed {
			t.Error("should execute after delay")
		}
	})

	t.Run("resets on subsequent calls", func(t *testing.T) {
		callCount := 0
		fn := func() { callCount++ }

		debounced, _ := Debounce(fn, 50*time.Millisecond)

		debounced()
		time.Sleep(25 * time.Millisecond)
		debounced() // Reset timer
		time.Sleep(25 * time.Millisecond)
		debounced() // Reset timer again

		if callCount != 0 {
			t.Error("should not have executed yet")
		}

		time.Sleep(100 * time.Millisecond)

		if callCount != 1 {
			t.Errorf("call count = %d, want 1", callCount)
		}
	})

	t.Run("cancel prevents execution", func(t *testing.T) {
		executed := false
		fn := func() { executed = true }

		debounced, cancel := Debounce(fn, 50*time.Millisecond)
		debounced()
		cancel()

		time.Sleep(100 * time.Millisecond)

		if executed {
			t.Error("cancel should prevent execution")
		}
	})
}

func TestSequence(t *testing.T) {
	t.Run("counting sequence", func(t *testing.T) {
		counter := Sequence(1, func(n int) int { return n + 1 })
		wants := []int{1, 2, 3, 4, 5}

		for i, want := range wants {
			if got := counter(); got != want {
				t.Errorf("call %d = %d, want %d", i+1, got, want)
			}
		}
	})

	t.Run("powers of 2", func(t *testing.T) {
		powers := Sequence(1, func(n int) int { return n * 2 })
		wants := []int{1, 2, 4, 8, 16}

		for i, want := range wants {
			if got := powers(); got != want {
				t.Errorf("call %d = %d, want %d", i+1, got, want)
			}
		}
	})

	t.Run("countdown", func(t *testing.T) {
		countdown := Sequence(10, func(n int) int { return n - 1 })
		wants := []int{10, 9, 8, 7, 6}

		for i, want := range wants {
			if got := countdown(); got != want {
				t.Errorf("call %d = %d, want %d", i+1, got, want)
			}
		}
	})

	t.Run("independent sequences", func(t *testing.T) {
		s1 := Sequence(0, func(n int) int { return n + 1 })
		s2 := Sequence(100, func(n int) int { return n + 1 })

		s1()
		s1()
		s1() // s1 is at 3

		if got := s2(); got != 100 {
			t.Errorf("s2 first call = %d, want 100 (should be independent)", got)
		}
	})
}

func TestRetry(t *testing.T) {
	t.Run("succeeds immediately", func(t *testing.T) {
		attempts := 0
		fn := func() error {
			attempts++
			return nil
		}

		retry := Retry(fn, 3, time.Millisecond)
		err := retry()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if attempts != 1 {
			t.Errorf("attempts = %d, want 1", attempts)
		}
	})

	t.Run("succeeds after retries", func(t *testing.T) {
		attempts := 0
		fn := func() error {
			attempts++
			if attempts < 3 {
				return errors.New("temporary failure")
			}
			return nil
		}

		retry := Retry(fn, 5, time.Millisecond)
		err := retry()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if attempts != 3 {
			t.Errorf("attempts = %d, want 3", attempts)
		}
	})

	t.Run("fails after max attempts", func(t *testing.T) {
		attempts := 0
		fn := func() error {
			attempts++
			return errors.New("permanent failure")
		}

		retry := Retry(fn, 3, time.Millisecond)
		err := retry()

		if err == nil {
			t.Error("expected error after max attempts")
		}
		if attempts != 3 {
			t.Errorf("attempts = %d, want 3", attempts)
		}
	})

	t.Run("uses exponential backoff", func(t *testing.T) {
		var timestamps []time.Time
		fn := func() error {
			timestamps = append(timestamps, time.Now())
			if len(timestamps) < 3 {
				return errors.New("fail")
			}
			return nil
		}

		retry := Retry(fn, 5, 10*time.Millisecond)
		retry()

		if len(timestamps) < 3 {
			t.Fatal("need at least 3 timestamps to check backoff")
		}

		delay1 := timestamps[1].Sub(timestamps[0])
		delay2 := timestamps[2].Sub(timestamps[1])

		// delay2 should be roughly 2x delay1 (with some tolerance)
		if delay2 < delay1 {
			t.Errorf("expected exponential backoff: delay1=%v, delay2=%v", delay1, delay2)
		}
	})
}

func TestFibonacci(t *testing.T) {
	fib := Fibonacci()
	wants := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	for i, want := range wants {
		if got := fib(); got != want {
			t.Errorf("fib() call %d = %d, want %d", i+1, got, want)
		}
	}
}

func TestFibonacci_Independent(t *testing.T) {
	fib1 := Fibonacci()
	fib2 := Fibonacci()

	// Advance fib1
	for i := 0; i < 5; i++ {
		fib1()
	}

	// fib2 should start from beginning
	if got := fib2(); got != 0 {
		t.Errorf("fib2 first call = %d, want 0 (should be independent)", got)
	}
}

// Benchmark tests
func BenchmarkMemoize(b *testing.B) {
	expensive := func(n int) int {
		// Simulate expensive computation
		result := 0
		for i := 0; i < 1000; i++ {
			result += n
		}
		return result
	}

	memoized := Memoize(expensive)

	// Warm up cache
	memoized(42)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoized(42) // Should hit cache
	}
}

func BenchmarkMemoize_Miss(b *testing.B) {
	expensive := func(n int) int {
		return n * n
	}

	memoized := Memoize(expensive)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoized(i) // Always miss
	}
}

// Test for potential race conditions (run with -race flag)
func TestMemoize_Concurrent(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping concurrent test in short mode")
	}

	callCount := 0
	var mu sync.Mutex
	fn := func(n int) int {
		mu.Lock()
		callCount++
		mu.Unlock()
		return n * n
	}

	memoized := Memoize(fn)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			memoized(n % 10) // Only 10 unique values
		}(i)
	}
	wg.Wait()

	// Note: basic memoize is not thread-safe
	// This test documents expected behavior, not correctness
	t.Logf("Call count with concurrent access: %d (may vary due to races)", callCount)
}
