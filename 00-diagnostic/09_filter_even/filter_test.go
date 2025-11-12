package filter

import (
	"reflect"
	"testing"
)

func TestFilterEven(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    []int
	}{
		{"Mixed numbers", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"No evens", []int{1, 3, 5}, []int{}},
		{"Empty input", []int{}, []int{}},
		{"All even", []int{2, 4, 6}, []int{2, 4, 6}},
		{"With zero", []int{0, 1, 2}, []int{0, 2}},
		{"Negative numbers", []int{-4, -3, -2, -1, 0, 1, 2}, []int{-4, -2, 0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterEven(tt.numbers)
			// Handle nil vs empty slice comparison
			if got == nil && tt.want != nil && len(tt.want) == 0 {
				got = []int{}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterEven(%v) = %v, want %v", tt.numbers, got, tt.want)
			}
		})
	}
}
