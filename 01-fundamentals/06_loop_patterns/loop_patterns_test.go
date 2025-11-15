package looppatterns

import (
	"reflect"
	"testing"
)

func TestCountUp(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"Count to 5", 5, []int{1, 2, 3, 4, 5}},
		{"Count to 1", 1, []int{1}},
		{"Count to 10", 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Zero returns empty", 0, []int{}},
		{"Negative returns empty", -3, []int{}},
		{"Count to 3", 3, []int{1, 2, 3}},
		{"Large negative", -100, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountUp(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountUp(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestCountDown(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"Count down from 5", 5, []int{5, 4, 3, 2, 1}},
		{"Count down from 1", 1, []int{1}},
		{"Count down from 10", 10, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{"Zero returns empty", 0, []int{}},
		{"Negative returns empty", -2, []int{}},
		{"Count down from 3", 3, []int{3, 2, 1}},
		{"Large negative", -50, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountDown(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountDown(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestSumSlice(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Sum 1 to 5", []int{1, 2, 3, 4, 5}, 15},
		{"Sum tens", []int{10, 20, 30}, 60},
		{"Negative sum", []int{-5, 5}, 0},
		{"Empty slice", []int{}, 0},
		{"Single element", []int{42}, 42},
		{"All negatives", []int{-1, -2, -3}, -6},
		{"Mixed numbers", []int{-10, 5, 15, -5, 25}, 30},
		{"Large numbers", []int{1000, 2000, 3000}, 6000},
		{"Zeros", []int{0, 0, 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumSlice(tt.nums)
			if got != tt.want {
				t.Errorf("SumSlice(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestFirstNEvens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"First 5 evens", 5, []int{2, 4, 6, 8, 10}},
		{"First 3 evens", 3, []int{2, 4, 6}},
		{"First even", 1, []int{2}},
		{"Zero returns empty", 0, []int{}},
		{"Negative returns empty", -5, []int{}},
		{"First 10 evens", 10, []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}},
		{"First 2 evens", 2, []int{2, 4}},
		{"Large negative", -100, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstNEvens(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstNEvens(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

// Benchmark tests to demonstrate loop performance characteristics
func BenchmarkCountUp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountUp(100)
	}
}

func BenchmarkCountDown(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountDown(100)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumSlice(nums)
	}
}

func BenchmarkFirstNEvens(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstNEvens(100)
	}
}
