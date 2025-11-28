package binary_search_recursive

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"Empty slice", []int{}, 5, -1},
		{"Single element - found", []int{5}, 5, 0},
		{"Single element - not found", []int{5}, 3, -1},
		{"Found at beginning", []int{1, 2, 3, 4, 5}, 1, 0},
		{"Found at end", []int{1, 2, 3, 4, 5}, 5, 4},
		{"Found in middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"Not found - too small", []int{1, 2, 3, 4, 5}, 0, -1},
		{"Not found - too large", []int{1, 2, 3, 4, 5}, 6, -1},
		{"Not found - gap", []int{1, 3, 5, 7, 9}, 4, -1},
		{"Larger array", []int{1, 3, 5, 7, 9, 11, 13, 15}, 11, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("BinarySearch(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestFindFirst(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"Not found", []int{1, 2, 3}, 4, -1},
		{"Single occurrence", []int{1, 2, 3}, 2, 1},
		{"Multiple - first", []int{1, 2, 2, 2, 3}, 2, 1},
		{"Multiple - all same", []int{1, 1, 1, 1}, 1, 0},
		{"Multiple - at beginning", []int{2, 2, 3, 4}, 2, 0},
		{"Multiple - at end", []int{1, 2, 3, 3, 3}, 3, 2},
		{"Two occurrences", []int{1, 2, 2, 3}, 2, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindFirst(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("FindFirst(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestFindLast(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"Not found", []int{1, 2, 3}, 4, -1},
		{"Single occurrence", []int{1, 2, 3}, 2, 1},
		{"Multiple - last", []int{1, 2, 2, 2, 3}, 2, 3},
		{"Multiple - all same", []int{1, 1, 1, 1}, 1, 3},
		{"Multiple - at beginning", []int{2, 2, 3, 4}, 2, 1},
		{"Multiple - at end", []int{1, 2, 3, 3, 3}, 3, 4},
		{"Two occurrences", []int{1, 2, 2, 3}, 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindLast(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("FindLast(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestSearchInsertPosition(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"Insert at beginning", []int{1, 3, 5, 7}, 0, 0},
		{"Insert at end", []int{1, 3, 5, 7}, 8, 4},
		{"Insert in middle", []int{1, 3, 5, 7}, 4, 2},
		{"Element exists - beginning", []int{1, 3, 5, 7}, 1, 0},
		{"Element exists - middle", []int{1, 3, 5, 7}, 5, 2},
		{"Element exists - end", []int{1, 3, 5, 7}, 7, 3},
		{"Empty array", []int{}, 5, 0},
		{"Single element - before", []int{5}, 3, 0},
		{"Single element - after", []int{5}, 7, 1},
		{"Single element - equal", []int{5}, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchInsertPosition(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("SearchInsertPosition(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

// Benchmarks - compare with linear search
func BenchmarkBinarySearch(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i * 2 // Even numbers
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(nums, 5000)
	}
}

func BenchmarkFindFirst(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i / 100 // Lots of duplicates
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindFirst(nums, 50)
	}
}

func BenchmarkFindLast(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i / 100
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindLast(nums, 50)
	}
}
