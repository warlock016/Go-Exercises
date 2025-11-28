package integer_algorithms

import "testing"

func TestGCD(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"Same numbers", 17, 17, 17},
		{"One is multiple of other", 24, 8, 8},
		{"Coprime numbers", 7, 3, 1},
		{"Classic example", 48, 18, 6},
		{"Different order", 18, 48, 6},
		{"Large numbers", 100, 35, 5},
		{"With zero", 15, 0, 15},
		{"Both factors", 12, 8, 4},
		{"Prime numbers", 13, 17, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GCD(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("GCD(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSumDigits(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Zero", 0, 0},
		{"Single digit", 5, 5},
		{"Two digits", 42, 6},
		{"Three digits", 123, 6},
		{"All nines", 999, 27},
		{"Larger number", 1234, 10},
		{"With zeros", 1001, 2},
		{"Large", 98765, 35},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumDigits(tt.input)
			if got != tt.want {
				t.Errorf("SumDigits(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestCountDigits(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Zero", 0, 1},
		{"Single digit", 5, 1},
		{"Two digits", 42, 2},
		{"Three digits", 999, 3},
		{"Four digits", 1234, 4},
		{"Five digits", 12345, 5},
		{"Power of 10", 10000, 5},
		{"Large", 987654321, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountDigits(tt.input)
			if got != tt.want {
				t.Errorf("CountDigits(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"Zero", 0, false},
		{"One (2^0)", 1, true},
		{"Two (2^1)", 2, true},
		{"Three", 3, false},
		{"Four (2^2)", 4, true},
		{"Eight (2^3)", 8, true},
		{"Ten", 10, false},
		{"Sixteen (2^4)", 16, true},
		{"Thirty-two (2^5)", 32, true},
		{"Sixty-four (2^6)", 64, true},
		{"Hundred", 100, false},
		{"128 (2^7)", 128, true},
		{"1024 (2^10)", 1024, true},
		{"1023", 1023, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPowerOfTwo(tt.input)
			if got != tt.want {
				t.Errorf("IsPowerOfTwo(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(48, 18)
	}
}

func BenchmarkSumDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumDigits(123456789)
	}
}

func BenchmarkCountDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountDigits(123456789)
	}
}

func BenchmarkIsPowerOfTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPowerOfTwo(1024)
	}
}
