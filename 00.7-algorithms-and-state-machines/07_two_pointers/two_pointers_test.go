package twopointers

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "basic duplicates",
			nums: []int{1, 1, 2, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "no duplicates",
			nums: []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "all duplicates",
			nums: []int{1, 1, 1, 1},
			want: []int{1},
		},
		{
			name: "empty array",
			nums: []int{},
			want: []int{},
		},
		{
			name: "single element",
			nums: []int{5},
			want: []int{5},
		},
		{
			name: "two same elements",
			nums: []int{2, 2},
			want: []int{2},
		},
		{
			name: "two different elements",
			nums: []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "multiple duplicates",
			nums: []int{1, 1, 2, 3, 3, 3, 4, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "many consecutive duplicates",
			nums: []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 3},
			want: []int{0, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy since we might modify in-place
			input := make([]int, len(tt.nums))
			copy(input, tt.nums)

			got := RemoveDuplicates(input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicates(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicatesNegativeNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "negative numbers",
			nums: []int{-3, -3, -2, -1, -1, 0, 0, 1},
			want: []int{-3, -2, -1, 0, 1},
		},
		{
			name: "all negative",
			nums: []int{-5, -5, -4, -4, -3},
			want: []int{-5, -4, -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.nums))
			copy(input, tt.nums)

			got := RemoveDuplicates(input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicates(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicatesLargeArray(t *testing.T) {
	// Test with larger input
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i / 10 // Each number appears 10 times
	}

	want := make([]int, 100)
	for i := 0; i < 100; i++ {
		want[i] = i
	}

	got := RemoveDuplicates(nums)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveDuplicates(large array) length = %d, want %d", len(got), len(want))
	}
}

// BONUS TESTS - Uncomment when ready to attempt bonus challenges!

/*
func TestRemoveDuplicatesKeepTwo(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "keep two duplicates",
			nums: []int{1, 1, 1, 2, 2, 3},
			want: []int{1, 1, 2, 2, 3},
		},
		{
			name: "all same",
			nums: []int{1, 1, 1, 1, 1},
			want: []int{1, 1},
		},
		{
			name: "no duplicates",
			nums: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "exactly two each",
			nums: []int{1, 1, 2, 2, 3, 3},
			want: []int{1, 1, 2, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.nums))
			copy(input, tt.nums)

			got := RemoveDuplicatesKeepTwo(input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicatesKeepTwo(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "zeros in middle",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "no zeros",
			nums: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "zeros at start",
			nums: []int{0, 0, 1, 2},
			want: []int{1, 2, 0, 0},
		},
		{
			name: "zeros at end",
			nums: []int{1, 2, 0, 0},
			want: []int{1, 2, 0, 0},
		},
		{
			name: "single zero",
			nums: []int{0},
			want: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.nums))
			copy(input, tt.nums)

			got := MoveZeroes(input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveZeroes(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
*/
