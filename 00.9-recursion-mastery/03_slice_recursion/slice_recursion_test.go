package slice_recursion

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{5}, 5},
		{"Two elements", []int{3, 7}, 10},
		{"Multiple positive", []int{1, 2, 3, 4, 5}, 15},
		{"With negatives", []int{10, -5, 3, -2}, 6},
		{"All negative", []int{-1, -2, -3}, -6},
		{"Large slice", []int{100, 200, 300, 400}, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.input)
			if got != tt.want {
				t.Errorf("Sum(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"Single element", []int{5}, 5},
		{"Two elements", []int{3, 7}, 7},
		{"First is max", []int{10, 5, 3}, 10},
		{"Last is max", []int{1, 5, 9}, 9},
		{"Middle is max", []int{3, 10, 5}, 10},
		{"All negative", []int{-5, -2, -10}, -2},
		{"Duplicates", []int{5, 9, 9, 3}, 9},
		{"Large numbers", []int{1000, 500, 2000, 1500}, 2000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.input)
			if got != tt.want {
				t.Errorf("Max(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{"Empty slice", []int{}, 5, false},
		{"Single element - found", []int{5}, 5, true},
		{"Single element - not found", []int{5}, 3, false},
		{"Found at beginning", []int{1, 2, 3, 4}, 1, true},
		{"Found at end", []int{1, 2, 3, 4}, 4, true},
		{"Found in middle", []int{1, 2, 3, 4}, 3, true},
		{"Not found", []int{1, 2, 3, 4}, 5, false},
		{"Negative numbers", []int{-5, -2, 0, 3}, -2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("Contains(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestCountOccurrences(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"Empty slice", []int{}, 5, 0},
		{"Not found", []int{1, 2, 3}, 5, 0},
		{"Found once", []int{1, 2, 3}, 2, 1},
		{"Found multiple times", []int{1, 2, 1, 3, 1}, 1, 3},
		{"All elements match", []int{5, 5, 5, 5}, 5, 4},
		{"Found at boundaries", []int{7, 2, 3, 7}, 7, 2},
		{"Large slice with repeats", []int{1, 2, 3, 2, 4, 2, 5}, 2, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountOccurrences(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("CountOccurrences(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkSum(b *testing.B) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(nums)
	}
}

func BenchmarkMax(b *testing.B) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Max(nums)
	}
}

func BenchmarkContains(b *testing.B) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(nums, 50)
	}
}
