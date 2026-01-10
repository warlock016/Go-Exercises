package channel_reinforcement

import (
	"testing"
	"testing/synctest"
	"time"
)

// =============================================================================
// Exercise 1: Echo
// =============================================================================

func TestEcho(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  int
	}{
		{"positive", 5, 10},
		{"zero", 0, 0},
		{"negative", -3, -6},
		{"large", 1000000, 2000000},
		{"one", 1, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Echo(tt.value)
			if got != tt.want {
				t.Errorf("Echo(%d) = %d, want %d", tt.value, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Exercise 2: Countdown
// =============================================================================

func TestCountdown(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"three", 3, []int{3, 2, 1}},
		{"one", 1, []int{1}},
		{"zero", 0, []int{}},
		{"five", 5, []int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := Countdown(tt.n)
			if ch == nil {
				t.Fatal("Countdown returned nil channel")
			}

			var got []int
			for v := range ch {
				got = append(got, v)
			}

			if len(got) != len(tt.want) {
				t.Errorf("Countdown(%d) sent %d values, want %d", tt.n, len(got), len(tt.want))
				return
			}

			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("Countdown(%d)[%d] = %d, want %d", tt.n, i, v, tt.want[i])
				}
			}
		})
	}
}

func TestCountdown_ChannelCloses(t *testing.T) {
	ch := Countdown(2)
	if ch == nil {
		t.Fatal("Countdown returned nil channel")
	}

	// Drain the channel
	for range ch {
	}

	// Channel should be closed, so receiving should return zero value and false
	select {
	case _, ok := <-ch:
		if ok {
			t.Error("Channel should be closed after draining")
		}
	default:
		// This is also acceptable - channel is closed
	}
}

// =============================================================================
// Exercise 3: Relay
// =============================================================================

func TestRelay(t *testing.T) {
	tests := []struct {
		name   string
		value  int
		stages int
		want   int
	}{
		{"three_stages", 10, 3, 13},
		{"five_stages", 0, 5, 5},
		{"zero_stages", 5, 0, 5},
		{"one_stage", 5, 1, 6},
		{"large_value", 1000, 10, 1010},
		{"negative_value", -5, 3, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Relay(tt.value, tt.stages)
			if got != tt.want {
				t.Errorf("Relay(%d, %d) = %d, want %d", tt.value, tt.stages, got, tt.want)
			}
		})
	}
}

func TestRelay_ManyStages(t *testing.T) {
	// Test with many stages to verify goroutine coordination
	got := Relay(0, 100)
	if got != 100 {
		t.Errorf("Relay(0, 100) = %d, want 100", got)
	}
}

// =============================================================================
// Exercise 4: FanIn
// =============================================================================

func TestFanIn(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{"three_values", []int{1, 2, 3}, 6},
		{"single_value", []int{10}, 10},
		{"empty", []int{}, 0},
		{"with_negative", []int{-1, 0, 1}, 0},
		{"larger", []int{10, 20, 30, 40}, 100},
		{"all_same", []int{5, 5, 5, 5, 5}, 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FanIn(tt.values)
			if got != tt.want {
				t.Errorf("FanIn(%v) = %d, want %d", tt.values, got, tt.want)
			}
		})
	}
}

func TestFanIn_Concurrent(t *testing.T) {
	// Test with many values to verify concurrent behavior
	values := make([]int, 100)
	want := 0
	for i := range values {
		values[i] = i + 1
		want += i + 1
	}

	got := FanIn(values)
	if got != want {
		t.Errorf("FanIn(1..100) = %d, want %d", got, want)
	}
}

// =============================================================================
// Exercise 5: Ticker (with synctest)
// =============================================================================

func TestTicker(t *testing.T) {
	// Note: t.Run cannot be called inside synctest.Test bubble
	// So each test case gets its own synctest.Test call

	tests := []struct {
		name     string
		n        int
		interval time.Duration
		want     []int
	}{
		{"three_values", 3, time.Second, []int{0, 1, 2}},
		{"one_value", 1, time.Millisecond, []int{0}},
		{"zero_values", 0, time.Second, []int{}},
		{"five_values", 5, 100 * time.Millisecond, []int{0, 1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Each subtest gets its own synctest bubble
			synctest.Test(t, func(t *testing.T) {
				ch := Ticker(tt.n, tt.interval)
				if ch == nil && tt.n > 0 {
					t.Fatal("Ticker returned nil channel")
				}
				if ch == nil && tt.n == 0 {
					// nil channel for zero values is acceptable
					return
				}

				var got []int
				for v := range ch {
					got = append(got, v)
				}

				if len(got) != len(tt.want) {
					t.Errorf("Ticker(%d, %v) sent %d values, want %d", tt.n, tt.interval, len(got), len(tt.want))
					return
				}

				for i, v := range got {
					if v != tt.want[i] {
						t.Errorf("Ticker(%d, %v)[%d] = %d, want %d", tt.n, tt.interval, i, v, tt.want[i])
					}
				}
			})
		})
	}
}

func TestTicker_ChannelCloses(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ch := Ticker(2, time.Second)
		if ch == nil {
			t.Fatal("Ticker returned nil channel")
		}

		// Drain the channel
		for range ch {
		}

		// Verify channel is closed
		select {
		case _, ok := <-ch:
			if ok {
				t.Error("Channel should be closed after draining")
			}
		default:
			// Closed channel is acceptable
		}
	})
}

// TestTicker_TimingWithSynctest demonstrates fake time behavior
func TestTicker_TimingWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// With synctest, this test runs instantly despite the 1-hour interval!
		ch := Ticker(3, time.Hour)
		if ch == nil {
			t.Fatal("Ticker returned nil channel")
		}

		count := 0
		for range ch {
			count++
		}

		if count != 3 {
			t.Errorf("Got %d values, want 3", count)
		}
	})
}

// =============================================================================
// Exercise 6: Pipeline
// =============================================================================

func TestPipeline(t *testing.T) {
	// Pipeline: Generate → FilterEven → Double → AddOne → FilterOdd → Sum
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"three", 3, 5},   // 1,2,3 → 2 → 4 → 5 → 5 → 5
		{"five", 5, 14},   // 1,2,3,4,5 → 2,4 → 4,8 → 5,9 → 5,9 → 14
		{"zero", 0, 0},    // no values generated
		{"one", 1, 0},     // 1 → (empty, 1 is odd) → 0
		{"ten", 10, 65},   // 2,4,6,8,10 → 4,8,12,16,20 → 5,9,13,17,21 → sum=65
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Pipeline(tt.n)
			if got != tt.want {
				t.Errorf("Pipeline(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestPipeline_LargeInput(t *testing.T) {
	// Pipeline(100): 50 even numbers → doubled → +1 → sum of 5,9,13,...,201
	// Arithmetic sequence: a=5, d=4, n=50, sum = 50/2 * (5+201) = 5150
	got := Pipeline(100)
	want := 5150
	if got != want {
		t.Errorf("Pipeline(100) = %d, want %d", got, want)
	}
}

// =============================================================================
// Benchmarks
// =============================================================================

func BenchmarkEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo(42)
	}
}

func BenchmarkCountdown(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := Countdown(10)
		for range ch {
		}
	}
}

func BenchmarkRelay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Relay(0, 10)
	}
}

func BenchmarkFanIn(b *testing.B) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FanIn(values)
	}
}

func BenchmarkPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pipeline(100)
	}
}
