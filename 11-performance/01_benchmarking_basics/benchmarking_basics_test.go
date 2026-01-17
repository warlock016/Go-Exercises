package benchmarking_basics

import (
	"fmt"
	"testing"
)

// =============================================================================
// Unit Tests
// =============================================================================

func TestSumLoop(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{5, 15},
		{10, 55},
		{100, 5050},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			if got := SumLoop(tt.n); got != tt.want {
				t.Errorf("SumLoop(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestSumFormula(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{5, 15},
		{10, 55},
		{100, 5050},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			if got := SumFormula(tt.n); got != tt.want {
				t.Errorf("SumFormula(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{5, 5},
		{10, 55},
		{20, 6765},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Recursive/n=%d", tt.n), func(t *testing.T) {
			if got := FibonacciRecursive(tt.n); got != tt.want {
				t.Errorf("FibonacciRecursive(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("Iterative/n=%d", tt.n), func(t *testing.T) {
			if got := FibonacciIterative(tt.n); got != tt.want {
				t.Errorf("FibonacciIterative(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"a", "a"},
		{"hello", "olleh"},
		{"Hello, World!", "!dlroW ,olleH"},
		{"日本語", "語本日"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Naive/%s", tt.input), func(t *testing.T) {
			if got := ReverseStringNaive(tt.input); got != tt.want {
				t.Errorf("ReverseStringNaive(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("Builder/%s", tt.input), func(t *testing.T) {
			if got := ReverseStringBuilder(tt.input); got != tt.want {
				t.Errorf("ReverseStringBuilder(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Benchmarks
// =============================================================================

// TODO(human): Write benchmarks for the functions above
//
// Suggested benchmarks:
//
// 1. BenchmarkSumLoop - benchmark SumLoop with different sizes
// 2. BenchmarkSumFormula - benchmark SumFormula with different sizes
// 3. BenchmarkSum - sub-benchmarks comparing Loop vs Formula at different sizes
// 4. BenchmarkFibonacciRecursive - benchmark recursive (use small n!)
// 5. BenchmarkFibonacciIterative - benchmark iterative
// 6. BenchmarkReverseString - compare Naive vs Builder with different string sizes
//
// Remember:
// - Use b.N for the loop count
// - Use b.Run() for sub-benchmarks
// - Use b.ReportAllocs() for memory stats
// - Use b.ResetTimer() after setup if needed

func BenchmarkSumLoop(b *testing.B) {
	// TODO(human): Implement
}

func BenchmarkSumFormula(b *testing.B) {
	// TODO(human): Implement
}

func BenchmarkSum(b *testing.B) {
	// TODO(human): Implement with sub-benchmarks for different sizes
}

func BenchmarkFibonacciRecursive(b *testing.B) {
	// TODO(human): Implement (use small n like 20, recursive is slow!)
}

func BenchmarkFibonacciIterative(b *testing.B) {
	// TODO(human): Implement
}

func BenchmarkReverseString(b *testing.B) {
	// TODO(human): Implement comparing Naive vs Builder
}
