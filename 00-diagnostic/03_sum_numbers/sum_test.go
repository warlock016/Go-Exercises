package sum

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{"Basic sum", []int{1, 2, 3, 4, 5}, 15},
		{"Two numbers", []int{10, 20}, 30},
		{"Empty slice", []int{}, 0},
		{"Negative numbers", []int{-5, 5}, 0},
		{"All negative", []int{-1, -2, -3}, -6},
		{"Single number", []int{42}, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.numbers)
			if got != tt.want {
				t.Errorf("Sum(%v) = %v, want %v", tt.numbers, got, tt.want)
			}
		})
	}
}
