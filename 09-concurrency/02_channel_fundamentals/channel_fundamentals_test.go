package channel_fundamentals

import (
	"testing"
)

func TestSendReceive(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  int
	}{
		{"zero", 0, 0},
		{"positive", 42, 42},
		{"negative", -17, -17},
		{"large", 1000000, 1000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SendReceive(tt.value)
			if got != tt.want {
				t.Errorf("SendReceive(%d) = %d, want %d", tt.value, got, tt.want)
			}
		})
	}
}

func TestGenerator(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"zero", 0, []int{}},
		{"one", 1, []int{1}},
		{"five", 5, []int{1, 2, 3, 4, 5}},
		{"ten", 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := Generator(tt.n)
			if ch == nil {
				t.Fatal("Generator() returned nil channel")
			}

			var got []int
			for v := range ch {
				got = append(got, v)
			}

			if len(got) != len(tt.want) {
				t.Errorf("Generator(%d) produced %d values, want %d", tt.n, len(got), len(tt.want))
				return
			}

			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("Generator(%d)[%d] = %d, want %d", tt.n, i, v, tt.want[i])
				}
			}
		})
	}
}

func TestGeneratorCloses(t *testing.T) {
	ch := Generator(3)

	// Drain the channel
	for range ch {
	}

	// Verify channel is closed (receive returns zero value, ok=false)
	v, ok := <-ch
	if ok {
		t.Errorf("Generator channel not closed, received %d", v)
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{"empty", []int{}, 0},
		{"single", []int{5}, 5},
		{"multiple", []int{1, 2, 3, 4, 5}, 15},
		{"with negatives", []int{-5, 10, -3, 8}, 10},
		{"all zeros", []int{0, 0, 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create and populate channel
			ch := make(chan int, len(tt.values))
			for _, v := range tt.values {
				ch <- v
			}
			close(ch)

			got := Sum(ch)
			if got != tt.want {
				t.Errorf("Sum() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestSumWithGenerator(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"five", 5, 15},       // 1+2+3+4+5
		{"ten", 10, 55},       // 1+2+...+10
		{"hundred", 100, 5050}, // n*(n+1)/2
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(Generator(tt.n))
			if got != tt.want {
				t.Errorf("Sum(Generator(%d)) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestPing(t *testing.T) {
	tests := []struct {
		name string
		msg  string
	}{
		{"hello", "hello"},
		{"empty", ""},
		{"unicode", "你好世界"},
		{"spaces", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := Ping(tt.msg)
			if ch == nil {
				t.Fatal("Ping() returned nil channel")
			}

			got := <-ch
			if got != tt.msg {
				t.Errorf("Ping(%q) sent %q, want %q", tt.msg, got, tt.msg)
			}
		})
	}
}

func TestPingPong(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero rounds", 0, 0},
		{"one round", 1, 2},    // ping adds 1, pong adds 1
		{"three rounds", 3, 6}, // 3 round trips, each adds 2
		{"five rounds", 5, 10},
		{"ten rounds", 10, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PingPong(tt.n)
			if got != tt.want {
				t.Errorf("PingPong(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

// Benchmark channel operations
func BenchmarkSendReceive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SendReceive(42)
	}
}

func BenchmarkGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for range Generator(100) {
		}
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(Generator(100))
	}
}
