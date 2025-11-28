package helper_functions

import (
	"reflect"
	"testing"
)

func TestFactorialAccum(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Base case: 0!", 0, 1},
		{"Base case: 1!", 1, 1},
		{"Small: 5!", 5, 120},
		{"Medium: 7!", 7, 5040},
		{"Large: 10!", 10, 3628800},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FactorialAccum(tt.input)
			if got != tt.want {
				t.Errorf("FactorialAccum(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"Empty list", []int{}, []int{}},
		{"Single element", []int{1}, []int{1}},
		{"Two elements", []int{1, 2}, []int{2, 1}},
		{"Three elements", []int{1, 2, 3}, []int{3, 2, 1}},
		{"Multiple elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"Already reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseList(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestRangeSum(t *testing.T) {
	tests := []struct {
		name  string
		start int
		end   int
		want  int
	}{
		{"Single number", 5, 5, 5},
		{"Small range", 1, 5, 15},        // 1+2+3+4+5 = 15
		{"Different range", 3, 7, 25},    // 3+4+5+6+7 = 25
		{"Larger range", 1, 10, 55},      // 1+2+...+10 = 55
		{"Starting from 0", 0, 4, 10},    // 0+1+2+3+4 = 10
		{"Invalid range", 10, 1, 0},      // start > end
		{"Negative range", -3, 3, 0},     // -3+-2+-1+0+1+2+3 = 0
		{"All negative", -5, -1, -15},    // -5+-4+-3+-2+-1 = -15
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RangeSum(tt.start, tt.end)
			if got != tt.want {
				t.Errorf("RangeSum(%d, %d) = %d, want %d", tt.start, tt.end, got, tt.want)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Zero", "0", 0},
		{"Single digit", "5", 5},
		{"Two digits", "42", 42},
		{"Three digits", "123", 123},
		{"Four digits", "4567", 4567},
		{"Large number", "98765", 98765},
		{"Negative single", "-5", -5},
		{"Negative two digits", "-42", -42},
		{"Negative three digits", "-123", -123},
		{"Negative large", "-9876", -9876},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringToInt(tt.input)
			if got != tt.want {
				t.Errorf("StringToInt(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// Benchmark accumulator vs regular recursion
func BenchmarkFactorialAccum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialAccum(10)
	}
}

func BenchmarkReverseList(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseList(nums)
	}
}

func BenchmarkRangeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeSum(1, 100)
	}
}

func BenchmarkStringToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToInt("123456")
	}
}
