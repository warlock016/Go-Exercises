package primechecker

import (
	"reflect"
	"testing"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		// Edge cases: numbers less than 2
		{"negative number", -5, false},
		{"zero", 0, false},
		{"one", 1, false},

		// Small primes
		{"two (only even prime)", 2, true},
		{"three", 3, true},
		{"five", 5, true},
		{"seven", 7, true},
		{"eleven", 11, true},
		{"thirteen", 13, true},

		// Small composites
		{"four", 4, false},
		{"six", 6, false},
		{"eight", 8, false},
		{"nine", 9, false},
		{"ten", 10, false},
		{"twelve", 12, false},

		// Larger primes
		{"seventeen", 17, true},
		{"nineteen", 19, true},
		{"twenty-three", 23, true},
		{"twenty-nine", 29, true},
		{"ninety-seven", 97, true},
		{"one hundred one", 101, true},

		// Larger composites
		{"fifteen", 15, false},
		{"twenty-one", 21, false},
		{"twenty-five", 25, false},
		{"one hundred", 100, false},

		// Edge case: large prime
		{"large prime", 7919, true},
		{"large composite", 7920, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPrime(tt.input)
			if got != tt.want {
				t.Errorf("IsPrime(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestFirstNPrimes(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  []int
	}{
		// Edge cases
		{"zero primes", 0, []int{}},
		{"negative count", -5, []int{}},

		// Small counts
		{"first prime", 1, []int{2}},
		{"first two primes", 2, []int{2, 3}},
		{"first five primes", 5, []int{2, 3, 5, 7, 11}},

		// Medium count
		{"first ten primes", 10, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},

		// Larger count
		{"first fifteen primes", 15, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstNPrimes(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstNPrimes(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestPrimesInRange(t *testing.T) {
	tests := []struct {
		name  string
		start int
		end   int
		want  []int
	}{
		// Edge cases
		{"start > end", 10, 5, []int{}},
		{"negative range", -10, -5, []int{}},
		{"range with negative start", -5, 10, []int{2, 3, 5, 7}},

		// Ranges with no primes
		{"no primes (14-16)", 14, 16, []int{}},
		{"no primes (24-28)", 24, 28, []int{}},

		// Small ranges
		{"1 to 10", 1, 10, []int{2, 3, 5, 7}},
		{"2 to 10", 2, 10, []int{2, 3, 5, 7}},
		{"3 to 10", 3, 10, []int{3, 5, 7}},

		// Medium ranges
		{"10 to 20", 10, 20, []int{11, 13, 17, 19}},
		{"20 to 30", 20, 30, []int{23, 29}},
		{"11 to 20", 11, 20, []int{11, 13, 17, 19}},

		// Range with single prime
		{"single prime (29)", 29, 29, []int{29}},
		{"single prime (2)", 2, 2, []int{2}},

		// Larger range
		{"50 to 70", 50, 70, []int{53, 59, 61, 67}},
		{"1 to 30", 1, 30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrimesInRange(tt.start, tt.end)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrimesInRange(%d, %d) = %v, want %v", tt.start, tt.end, got, tt.want)
			}
		})
	}
}

// Benchmark IsPrime with a large prime number
func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(7919)
	}
}

// Benchmark FirstNPrimes with moderate count
func BenchmarkFirstNPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstNPrimes(100)
	}
}

// Benchmark PrimesInRange with moderate range
func BenchmarkPrimesInRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimesInRange(1, 1000)
	}
}
