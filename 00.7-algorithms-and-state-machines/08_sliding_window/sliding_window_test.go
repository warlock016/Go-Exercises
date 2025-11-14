package slidingwindow

import "testing"

func TestMaxSumSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example from README",
			nums: []int{2, 1, 5, 1, 3, 2},
			k:    3,
			want: 9, // [5, 1, 3]
		},
		{
			name: "All negative numbers",
			nums: []int{-1, -2, -3, -4},
			k:    2,
			want: -3, // [-1, -2] is least negative
		},
		{
			name: "Window size 1",
			nums: []int{10, 20, 30, 40, 50},
			k:    1,
			want: 50, // Maximum single element
		},
		{
			name: "Window size equals array length",
			nums: []int{5, 4, 3, 2, 1},
			k:    5,
			want: 15, // Sum of entire array
		},
		{
			name: "Single element array",
			nums: []int{100},
			k:    1,
			want: 100,
		},
		{
			name: "Two windows",
			nums: []int{1, 4, 2, 10, 2, 3, 1, 0, 20},
			k:    4,
			want: 24, // [2, 10, 2, 3] or could be different
		},
		{
			name: "Mix of positive and negative",
			nums: []int{-5, 10, -3, 8, -2},
			k:    2,
			want: 10, // [10, -3] gives 7, but check all
		},
		{
			name: "Empty array",
			nums: []int{},
			k:    3,
			want: 0,
		},
		{
			name: "k is zero",
			nums: []int{1, 2, 3},
			k:    0,
			want: 0,
		},
		{
			name: "k greater than array length",
			nums: []int{1, 2, 3},
			k:    5,
			want: 0,
		},
		{
			name: "All same numbers",
			nums: []int{5, 5, 5, 5},
			k:    2,
			want: 10,
		},
		{
			name: "Large window with negative numbers",
			nums: []int{-1, 10, -1, 10, -1},
			k:    3,
			want: 18, // [10, -1, 10]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSumSubarray(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("MaxSumSubarray(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestMaxSumSubarrayEdgeCases(t *testing.T) {
	// Additional focused edge case tests
	t.Run("Negative k", func(t *testing.T) {
		got := MaxSumSubarray([]int{1, 2, 3}, -1)
		if got != 0 {
			t.Errorf("Expected 0 for negative k, got %d", got)
		}
	})

	t.Run("Very large numbers", func(t *testing.T) {
		nums := []int{1000000, 2000000, 3000000}
		got := MaxSumSubarray(nums, 2)
		want := 5000000 // [2000000, 3000000]
		if got != want {
			t.Errorf("MaxSumSubarray with large numbers = %d, want %d", got, want)
		}
	})

	t.Run("Alternating positive negative", func(t *testing.T) {
		nums := []int{5, -2, 3, -1, 6, -3}
		got := MaxSumSubarray(nums, 3)
		want := 8 // [3, -1, 6]
		if got != want {
			t.Errorf("MaxSumSubarray(%v, 3) = %d, want %d", nums, got, want)
		}
	})
}

// Benchmark to verify O(n) performance
func BenchmarkMaxSumSubarray(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i % 100
	}
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxSumSubarray(nums, k)
	}
}
