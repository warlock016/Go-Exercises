package merge_sort

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		left  []int
		right []int
		want  []int
	}{
		{"Both empty", []int{}, []int{}, []int{}},
		{"Left empty", []int{}, []int{1, 2}, []int{1, 2}},
		{"Right empty", []int{1, 2}, []int{}, []int{1, 2}},
		{"Single elements", []int{1}, []int{2}, []int{1, 2}},
		{"Interleaved", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"All left first", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"All right first", []int{4, 5, 6}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.left, tt.right)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge(%v, %v) = %v, want %v", tt.left, tt.right, got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"Empty", []int{}, []int{}},
		{"Single element", []int{5}, []int{5}},
		{"Two elements - sorted", []int{1, 2}, []int{1, 2}},
		{"Two elements - unsorted", []int{2, 1}, []int{1, 2}},
		{"Three elements", []int{3, 1, 2}, []int{1, 2, 3}},
		{"Already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"Random order", []int{5, 2, 8, 1, 9}, []int{1, 2, 5, 8, 9}},
		{"With duplicates", []int{3, 3, 1, 2, 1}, []int{1, 1, 2, 3, 3}},
		{"Larger array", []int{9, 2, 5, 1, 7, 6, 8, 3, 4}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSort(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	nums := []int{9, 2, 5, 1, 7, 6, 8, 3, 4, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSort(nums)
	}
}
