package factorial_power

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Base case: 0! = 1", 0, 1},
		{"Base case: 1! = 1", 1, 1},
		{"Small: 2!", 2, 2},
		{"Small: 3!", 3, 6},
		{"Medium: 5!", 5, 120},
		{"Medium: 7!", 7, 5040},
		{"Large: 10!", 10, 3628800},
		{"Edge: 12!", 12, 479001600},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Factorial(tt.input)
			if got != tt.want {
				t.Errorf("Factorial(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		want     int
	}{
		{"Any number to power 0", 2, 0, 1},
		{"Different base to power 0", 99, 0, 1},
		{"Base to power 1", 5, 1, 5},
		{"2^5", 2, 5, 32},
		{"3^4", 3, 4, 81},
		{"5^3", 5, 3, 125},
		{"10^2", 10, 2, 100},
		{"2^10", 2, 10, 1024},
		{"Edge: 1^100", 1, 100, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Power(tt.base, tt.exponent)
			if got != tt.want {
				t.Errorf("Power(%d, %d) = %d, want %d", tt.base, tt.exponent, got, tt.want)
			}
		})
	}
}

// Benchmark to compare with iterative versions (if you implement them)
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10)
	}
}

func BenchmarkPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Power(2, 10)
	}
}
