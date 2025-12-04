package closures

import (
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	count := Counter()

	// Test sequential calls
	for i := 1; i <= 5; i++ {
		got := count()
		if got != i {
			t.Errorf("count() = %d, want %d", got, i)
		}
	}

	// Test independent counters
	count2 := Counter()
	if got := count2(); got != 1 {
		t.Errorf("new counter() = %d, want 1", got)
	}

	// Original counter continues from where it left off
	if got := count(); got != 6 {
		t.Errorf("original counter() = %d, want 6", got)
	}
}

func TestAccumulator(t *testing.T) {
	acc := Accumulator()

	tests := []struct {
		add  int
		want int
	}{
		{5, 5},
		{3, 8},
		{-2, 6},
		{0, 6},
		{10, 16},
	}

	for _, tt := range tests {
		got := acc(tt.add)
		if got != tt.want {
			t.Errorf("acc(%d) = %d, want %d", tt.add, got, tt.want)
		}
	}

	// Test independent accumulator
	acc2 := Accumulator()
	if got := acc2(5); got != 5 {
		t.Errorf("new accumulator(5) = %d, want 5", got)
	}
}

func TestMultiplier(t *testing.T) {
	double := Multiplier(2)
	triple := Multiplier(3)

	tests := []struct {
		name string
		fn   func(int) int
		arg  int
		want int
	}{
		{"double 5", double, 5, 10},
		{"triple 5", triple, 5, 15},
		{"double 7", double, 7, 14},
		{"triple 4", triple, 4, 12},
		{"double 0", double, 0, 0},
		{"triple -3", triple, -3, -9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn(tt.arg)
			if got != tt.want {
				t.Errorf("multiplier(%d) = %d, want %d", tt.arg, got, tt.want)
			}
		})
	}

	// Test with different factors
	times10 := Multiplier(10)
	if got := times10(7); got != 70 {
		t.Errorf("Multiplier(10)(7) = %d, want 70", got)
	}
}

func TestRateLimiter(t *testing.T) {
	t.Run("allows calls within limit", func(t *testing.T) {
		limiter := RateLimiter(3, 100*time.Millisecond)

		// First 3 calls should succeed
		for i := 0; i < 3; i++ {
			if !limiter() {
				t.Errorf("call %d should be allowed", i+1)
			}
		}

		// 4th call should fail
		if limiter() {
			t.Error("4th call should be denied")
		}
	})

	t.Run("resets after window", func(t *testing.T) {
		limiter := RateLimiter(2, 50*time.Millisecond)

		// Use up limit
		limiter()
		limiter()

		// Should be denied
		if limiter() {
			t.Error("3rd call should be denied")
		}

		// Wait for window to pass
		time.Sleep(60 * time.Millisecond)

		// Should work again
		if !limiter() {
			t.Error("call after window should be allowed")
		}
	})

	t.Run("sliding window", func(t *testing.T) {
		limiter := RateLimiter(2, 100*time.Millisecond)

		// First call
		if !limiter() {
			t.Error("first call should be allowed")
		}

		// Wait 60ms
		time.Sleep(60 * time.Millisecond)

		// Second call (should work, within window)
		if !limiter() {
			t.Error("second call should be allowed")
		}

		// Third call should fail (2 calls in last 100ms)
		if limiter() {
			t.Error("third call should be denied")
		}

		// Wait another 50ms (first call now outside window)
		time.Sleep(50 * time.Millisecond)

		// Should work again
		if !limiter() {
			t.Error("call after first expires should be allowed")
		}
	})

	t.Run("high rate limiter", func(t *testing.T) {
		limiter := RateLimiter(10, 100*time.Millisecond)

		// Should allow 10 calls
		for i := 0; i < 10; i++ {
			if !limiter() {
				t.Errorf("call %d should be allowed", i+1)
			}
		}

		// 11th should fail
		if limiter() {
			t.Error("11th call should be denied")
		}
	})
}

func BenchmarkCounter(b *testing.B) {
	count := Counter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count()
	}
}

func BenchmarkAccumulator(b *testing.B) {
	acc := Accumulator()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		acc(1)
	}
}

func BenchmarkMultiplier(b *testing.B) {
	mult := Multiplier(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mult(i)
	}
}
