package performance_optimization

import (
	"testing"
)

func TestAppendNaive(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"empty", 0, []int{}},
		{"single element", 1, []int{0}},
		{"five elements", 5, []int{0, 1, 2, 3, 4}},
		{"ten elements", 10, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AppendNaive(tt.n)
			if len(got) != len(tt.want) {
				t.Fatalf("AppendNaive(%d) length = %d, want %d", tt.n, len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("AppendNaive(%d)[%d] = %d, want %d", tt.n, i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestAppendOptimized(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"empty", 0, []int{}},
		{"single element", 1, []int{0}},
		{"five elements", 5, []int{0, 1, 2, 3, 4}},
		{"ten elements", 10, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AppendOptimized(tt.n)
			if len(got) != len(tt.want) {
				t.Fatalf("AppendOptimized(%d) length = %d, want %d", tt.n, len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("AppendOptimized(%d)[%d] = %d, want %d", tt.n, i, got[i], tt.want[i])
				}
			}
			// Check capacity is optimal (should be exactly n)
			if cap(got) > tt.n && tt.n > 0 {
				t.Logf("AppendOptimized(%d) capacity = %d (could be optimized to exactly %d)", tt.n, cap(got), tt.n)
			}
		})
	}
}

func TestConcatStringsNaive(t *testing.T) {
	tests := []struct {
		name string
		str  string
		n    int
		want string
	}{
		{"empty string", "", 5, ""},
		{"zero repetitions", "Go", 0, ""},
		{"single repetition", "Go", 1, "Go"},
		{"three repetitions", "Go", 3, "GoGoGo"},
		{"five repetitions", "Hello", 5, "HelloHelloHelloHelloHello"},
		{"unicode string", "🔥", 3, "🔥🔥🔥"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConcatStringsNaive(tt.str, tt.n)
			if got != tt.want {
				t.Errorf("ConcatStringsNaive(%q, %d) = %q, want %q", tt.str, tt.n, got, tt.want)
			}
		})
	}
}

func TestConcatStringsOptimized(t *testing.T) {
	tests := []struct {
		name string
		str  string
		n    int
		want string
	}{
		{"empty string", "", 5, ""},
		{"zero repetitions", "Go", 0, ""},
		{"single repetition", "Go", 1, "Go"},
		{"three repetitions", "Go", 3, "GoGoGo"},
		{"five repetitions", "Hello", 5, "HelloHelloHelloHelloHello"},
		{"unicode string", "🔥", 3, "🔥🔥🔥"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConcatStringsOptimized(tt.str, tt.n)
			if got != tt.want {
				t.Errorf("ConcatStringsOptimized(%q, %d) = %q, want %q", tt.str, tt.n, got, tt.want)
			}
		})
	}
}

func TestFilterNaive(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    []int
	}{
		{"empty slice", []int{}, []int{}},
		{"all odd", []int{1, 3, 5, 7}, []int{}},
		{"all even", []int{2, 4, 6, 8}, []int{2, 4, 6, 8}},
		{"mixed", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"single even", []int{2}, []int{2}},
		{"single odd", []int{1}, []int{}},
		{"zero is even", []int{0, 1, 2}, []int{0, 2}},
		{"negative numbers", []int{-4, -3, -2, -1, 0, 1, 2}, []int{-4, -2, 0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterNaive(tt.numbers)
			if len(got) != len(tt.want) {
				t.Fatalf("FilterNaive(%v) length = %d, want %d", tt.numbers, len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("FilterNaive(%v)[%d] = %d, want %d", tt.numbers, i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestFilterOptimized(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    []int
	}{
		{"empty slice", []int{}, []int{}},
		{"all odd", []int{1, 3, 5, 7}, []int{}},
		{"all even", []int{2, 4, 6, 8}, []int{2, 4, 6, 8}},
		{"mixed", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"single even", []int{2}, []int{2}},
		{"single odd", []int{1}, []int{}},
		{"zero is even", []int{0, 1, 2}, []int{0, 2}},
		{"negative numbers", []int{-4, -3, -2, -1, 0, 1, 2}, []int{-4, -2, 0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterOptimized(tt.numbers)
			if len(got) != len(tt.want) {
				t.Fatalf("FilterOptimized(%v) length = %d, want %d", tt.numbers, len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("FilterOptimized(%v)[%d] = %d, want %d", tt.numbers, i, got[i], tt.want[i])
				}
			}
		})
	}
}

// Benchmark tests - Run with: go test -bench=. -benchmem

func BenchmarkAppendNaive100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendNaive(100)
	}
}

func BenchmarkAppendOptimized100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendOptimized(100)
	}
}

func BenchmarkAppendNaive1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendNaive(1000)
	}
}

func BenchmarkAppendOptimized1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendOptimized(1000)
	}
}

func BenchmarkConcatStringsNaive10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringsNaive("Go", 10)
	}
}

func BenchmarkConcatStringsOptimized10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringsOptimized("Go", 10)
	}
}

func BenchmarkConcatStringsNaive100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringsNaive("Go", 100)
	}
}

func BenchmarkConcatStringsOptimized100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringsOptimized("Go", 100)
	}
}

func BenchmarkFilterNaive(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterNaive(numbers)
	}
}

func BenchmarkFilterOptimized(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterOptimized(numbers)
	}
}
