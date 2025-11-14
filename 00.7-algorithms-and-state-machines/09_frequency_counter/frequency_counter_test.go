package frequencycounter

import (
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int // Expected elements (order doesn't matter)
	}{
		{
			name: "Example from README",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "Single element",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "Tied frequencies",
			nums: []int{4, 1, -1, 2, -1, 2, 3},
			k:    2,
			want: []int{-1, 2}, // Both appear twice
		},
		{
			name: "All same element",
			nums: []int{5, 5, 5, 5, 5},
			k:    1,
			want: []int{5},
		},
		{
			name: "All unique elements",
			nums: []int{1, 2, 3, 4, 5},
			k:    3,
			want: []int{1, 2, 3}, // Any 3 are valid since all freq=1
		},
		{
			name: "k equals number of unique elements",
			nums: []int{1, 1, 2, 2, 3, 3},
			k:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "k is 1",
			nums: []int{1, 2, 3, 3, 3, 2},
			k:    1,
			want: []int{3}, // 3 appears 3 times
		},
		{
			name: "Empty array",
			nums: []int{},
			k:    1,
			want: []int{},
		},
		{
			name: "k is 0",
			nums: []int{1, 2, 3},
			k:    0,
			want: []int{},
		},
		{
			name: "k greater than unique elements",
			nums: []int{1, 1, 2},
			k:    5,
			want: []int{1, 2}, // Only 2 unique elements
		},
		{
			name: "Negative numbers",
			nums: []int{-1, -1, -2, -2, -2, -3},
			k:    2,
			want: []int{-2, -1},
		},
		{
			name: "Large array with clear winners",
			nums: []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 4},
			k:    2,
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TopKFrequent(tt.nums, tt.k)

			// Check length
			if len(got) != len(tt.want) {
				t.Errorf("TopKFrequent() returned %d elements, want %d", len(got), len(tt.want))
				return
			}

			// For tests with unique frequencies, check exact match
			// For tests with tied frequencies, check if result is valid subset
			if !containsSameElements(got, tt.want) {
				// Special case: all elements have same frequency
				if allSameFrequency(tt.nums) {
					// Any k elements are valid
					if !allElementsExist(got, tt.nums) {
						t.Errorf("TopKFrequent() = %v, but elements not from input %v", got, tt.nums)
					}
				} else {
					t.Errorf("TopKFrequent() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestTopKFrequentEdgeCases(t *testing.T) {
	t.Run("Very large k", func(t *testing.T) {
		nums := []int{1, 2, 3}
		k := 1000
		got := TopKFrequent(nums, k)
		// Should return all unique elements (3), not crash or return garbage
		if len(got) > 3 {
			t.Errorf("k=1000 but only 3 unique elements, got %d results", len(got))
		}
	})

	t.Run("Repeated k=1 test with different winners", func(t *testing.T) {
		tests := []struct {
			nums []int
			want int
		}{
			{[]int{1, 1, 1, 2, 2}, 1},
			{[]int{5, 5, 5, 5, 1}, 5},
			{[]int{2, 3, 3, 3, 4, 4}, 3},
		}

		for _, tt := range tests {
			got := TopKFrequent(tt.nums, 1)
			if len(got) != 1 || got[0] != tt.want {
				t.Errorf("TopKFrequent(%v, 1) = %v, want [%d]", tt.nums, got, tt.want)
			}
		}
	})
}

// Helper functions for testing

// containsSameElements checks if two slices contain the same elements (order doesn't matter)
func containsSameElements(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	aCopy := make([]int, len(a))
	bCopy := make([]int, len(b))
	copy(aCopy, a)
	copy(bCopy, b)

	sort.Ints(aCopy)
	sort.Ints(bCopy)

	for i := range aCopy {
		if aCopy[i] != bCopy[i] {
			return false
		}
	}
	return true
}

// allSameFrequency checks if all unique elements in nums have the same frequency
func allSameFrequency(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	firstCount := -1
	for _, count := range freq {
		if firstCount == -1 {
			firstCount = count
		} else if count != firstCount {
			return false
		}
	}
	return true
}

// allElementsExist checks if all elements in subset exist in the full set
func allElementsExist(subset, fullset []int) bool {
	exists := make(map[int]bool)
	for _, num := range fullset {
		exists[num] = true
	}

	for _, num := range subset {
		if !exists[num] {
			return false
		}
	}
	return true
}

// Benchmark to test performance
func BenchmarkTopKFrequent(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i % 100 // 100 unique elements with varying frequencies
	}
	k := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TopKFrequent(nums, k)
	}
}
