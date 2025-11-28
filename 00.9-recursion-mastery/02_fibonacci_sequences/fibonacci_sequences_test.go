package fibonacci_sequences

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Base case: fib(0)", 0, 0},
		{"Base case: fib(1)", 1, 1},
		{"Small: fib(2)", 2, 1},
		{"Small: fib(3)", 3, 2},
		{"Small: fib(4)", 4, 3},
		{"Medium: fib(5)", 5, 5},
		{"Medium: fib(6)", 6, 8},
		{"Medium: fib(7)", 7, 13},
		{"Medium: fib(10)", 10, 55},
		{"Large: fib(15)", 15, 610},
		{"Large: fib(20)", 20, 6765},
		// Don't go beyond 25 with naive recursion - too slow!
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Fibonacci(tt.input)
			if got != tt.want {
				t.Errorf("Fibonacci(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestTribonacci(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Base case: trib(0)", 0, 0},
		{"Base case: trib(1)", 1, 1},
		{"Base case: trib(2)", 2, 1},
		{"Small: trib(3)", 3, 2},
		{"Small: trib(4)", 4, 4},
		{"Medium: trib(5)", 5, 7},
		{"Medium: trib(6)", 6, 13},
		{"Medium: trib(7)", 7, 24},
		{"Medium: trib(10)", 10, 149},
		{"Large: trib(15)", 15, 3136},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Tribonacci(tt.input)
			if got != tt.want {
				t.Errorf("Tribonacci(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"1 stair", 1, 1},
		{"2 stairs", 2, 2},
		{"3 stairs", 3, 3},
		{"4 stairs", 4, 5},
		{"5 stairs", 5, 8},
		{"6 stairs", 6, 13},
		{"10 stairs", 10, 89},
		// Notice: this is Fibonacci(n+1)!
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClimbStairs(tt.input)
			if got != tt.want {
				t.Errorf("ClimbStairs(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// Benchmarks to experience exponential slowdown
func BenchmarkFibonacci10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

func BenchmarkFibonacci20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}

func BenchmarkFibonacci25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(25) // Notice how much slower this is!
	}
}

func BenchmarkTribonacci10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tribonacci(10)
	}
}

func BenchmarkTribonacci15(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tribonacci(15)
	}
}
