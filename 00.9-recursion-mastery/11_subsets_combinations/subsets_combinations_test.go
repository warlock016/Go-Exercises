package subsets_combinations

import (
	"testing"
)

func TestGenerateSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int // count of subsets
	}{
		{"Empty", []int{}, 1},           // [[]]
		{"Single", []int{1}, 2},         // [[], [1]]
		{"Two elements", []int{1, 2}, 4}, // [[], [1], [2], [1,2]]
		{"Three elements", []int{1, 2, 3}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateSubsets(tt.nums)
			if len(got) != tt.want {
				t.Errorf("GenerateSubsets(%v) returned %d subsets, want %d", tt.nums, len(got), tt.want)
			}
		})
	}
}

func TestCombinations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"C(3,2)", []int{1, 2, 3}, 2, 3}, // [[1,2], [1,3], [2,3]]
		{"C(4,2)", []int{1, 2, 3, 4}, 2, 6},
		{"C(3,1)", []int{1, 2, 3}, 1, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Combinations(tt.nums, tt.k)
			if len(got) != tt.want {
				t.Errorf("Combinations(%v, %d) returned %d combinations, want %d", tt.nums, tt.k, len(got), tt.want)
			}
		})
	}
}

func TestPermutations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int // factorial(n)
	}{
		{"1 element", []int{1}, 1},
		{"2 elements", []int{1, 2}, 2},
		{"3 elements", []int{1, 2, 3}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Permutations(tt.nums)
			if len(got) != tt.want {
				t.Errorf("Permutations(%v) returned %d permutations, want %d", tt.nums, len(got), tt.want)
			}
		})
	}
}
