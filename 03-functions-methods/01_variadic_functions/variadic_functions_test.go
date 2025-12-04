package variadic_functions

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"empty", []int{}, 0},
		{"single value", []int{5}, 5},
		{"multiple values", []int{1, 2, 3}, 6},
		{"negative values", []int{-1, -2, -3}, -6},
		{"mixed values", []int{10, -5, 3}, 8},
		{"large numbers", []int{1000, 2000, 3000}, 6000},
		{"many values", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.nums...)
			if got != tt.want {
				t.Errorf("Sum(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		want      int
		wantError bool
	}{
		{"empty", []int{}, 0, true},
		{"single value", []int{5}, 5, false},
		{"multiple values", []int{1, 9, 3, 7}, 9, false},
		{"all negative", []int{-5, -10, -2}, -2, false},
		{"first is max", []int{10, 5, 3}, 10, false},
		{"last is max", []int{1, 5, 10}, 10, false},
		{"all same", []int{5, 5, 5}, 5, false},
		{"large numbers", []int{1000, 5000, 3000}, 5000, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Max(tt.nums...)
			if tt.wantError {
				if err == nil {
					t.Errorf("Max(%v) expected error, got nil", tt.nums)
				}
			} else {
				if err != nil {
					t.Errorf("Max(%v) unexpected error: %v", tt.nums, err)
				}
				if got != tt.want {
					t.Errorf("Max(%v) = %d, want %d", tt.nums, got, tt.want)
				}
			}
		})
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		name      string
		separator string
		parts     []string
		want      string
	}{
		{"empty", ", ", []string{}, ""},
		{"single part", ", ", []string{"apple"}, "apple"},
		{"two parts", ", ", []string{"a", "b"}, "a, b"},
		{"three parts", ", ", []string{"a", "b", "c"}, "a, b, c"},
		{"space separator", " - ", []string{"Go", "is", "fun"}, "Go - is - fun"},
		{"path separator", "/", []string{"usr", "local", "bin"}, "usr/local/bin"},
		{"empty separator", "", []string{"a", "b", "c"}, "abc"},
		{"many parts", ", ", []string{"one", "two", "three", "four", "five"}, "one, two, three, four, five"},
		{"parts with spaces", "|", []string{"hello world", "foo bar"}, "hello world|foo bar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Concat(tt.separator, tt.parts...)
			if got != tt.want {
				t.Errorf("Concat(%q, %v) = %q, want %q", tt.separator, tt.parts, got, tt.want)
			}
		})
	}
}

func BenchmarkSum(b *testing.B) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(nums...)
	}
}

func BenchmarkConcat(b *testing.B) {
	parts := []string{"one", "two", "three", "four", "five"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Concat(", ", parts...)
	}
}
