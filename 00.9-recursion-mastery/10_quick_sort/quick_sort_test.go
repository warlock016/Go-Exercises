package quick_sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"Empty", []int{}, []int{}},
		{"Single", []int{5}, []int{5}},
		{"Two elements", []int{2, 1}, []int{1, 2}},
		{"Random", []int{5, 2, 8, 1, 9}, []int{1, 2, 5, 8, 9}},
		{"With duplicates", []int{3, 3, 1, 2}, []int{1, 2, 3, 3}},
		{"Already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Reverse", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := QuickSort(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
