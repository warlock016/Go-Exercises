package slice_basics

import (
	"testing"
)

func TestCreateSliceWithMake(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		capacity int
		wantLen  int
		wantCap  int
	}{
		{"zero length and capacity", 0, 0, 0, 0},
		{"zero length, non-zero capacity", 0, 10, 0, 10},
		{"equal length and capacity", 5, 5, 5, 5},
		{"length less than capacity", 3, 10, 3, 10},
		{"large capacity", 0, 1000, 0, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateSliceWithMake(tt.length, tt.capacity)
			if len(got) != tt.wantLen {
				t.Errorf("CreateSliceWithMake(%d, %d) length = %d, want %d", tt.length, tt.capacity, len(got), tt.wantLen)
			}
			if cap(got) != tt.wantCap {
				t.Errorf("CreateSliceWithMake(%d, %d) capacity = %d, want %d", tt.length, tt.capacity, cap(got), tt.wantCap)
			}
			// Check zero values
			for i, val := range got {
				if val != 0 {
					t.Errorf("CreateSliceWithMake(%d, %d)[%d] = %d, want 0 (zero value)", tt.length, tt.capacity, i, val)
				}
			}
		})
	}
}

func TestCreateSliceLiteral(t *testing.T) {
	got := CreateSliceLiteral()
	want := []int{1, 2, 3, 4, 5}

	if len(got) != len(want) {
		t.Fatalf("CreateSliceLiteral() length = %d, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("CreateSliceLiteral()[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestAppendToSlice(t *testing.T) {
	tests := []struct {
		name   string
		slice  []int
		values []int
		want   []int
	}{
		{"append to empty slice", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"append single value", []int{1, 2}, []int{3}, []int{1, 2, 3}},
		{"append multiple values", []int{1}, []int{2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"append to nil slice", nil, []int{1, 2}, []int{1, 2}},
		{"append empty to slice", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AppendToSlice(tt.slice, tt.values...)
			if len(got) != len(tt.want) {
				t.Fatalf("AppendToSlice(%v, %v) length = %d, want %d", tt.slice, tt.values, len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("AppendToSlice(%v, %v)[%d] = %d, want %d", tt.slice, tt.values, i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestGetLengthAndCapacity(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		wantLen int
		wantCap int
	}{
		{"nil slice", nil, 0, 0},
		{"empty slice", []int{}, 0, 0},
		{"slice with elements", []int{1, 2, 3}, 3, 3},
		{"slice with capacity", make([]int, 5, 10), 5, 10},
		{"large slice", make([]int, 100), 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLen, gotCap := GetLengthAndCapacity(tt.slice)
			if gotLen != tt.wantLen {
				t.Errorf("GetLengthAndCapacity(%v) length = %d, want %d", tt.slice, gotLen, tt.wantLen)
			}
			if gotCap != tt.wantCap {
				t.Errorf("GetLengthAndCapacity(%v) capacity = %d, want %d", tt.slice, gotCap, tt.wantCap)
			}
		})
	}
}

func TestIterateAndSum(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{5}, 5},
		{"multiple elements", []int{1, 2, 3, 4, 5}, 15},
		{"negative numbers", []int{-1, -2, -3}, -6},
		{"mixed positive and negative", []int{10, -5, 3, -2}, 6},
		{"zeros", []int{0, 0, 0}, 0},
		{"large numbers", []int{1000, 2000, 3000}, 6000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IterateAndSum(tt.slice)
			if got != tt.want {
				t.Errorf("IterateAndSum(%v) = %d, want %d", tt.slice, got, tt.want)
			}
		})
	}
}

func TestIterateWithIndex(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{5}, []int{10}},
		{"multiple elements", []int{1, 2, 3}, []int{2, 4, 6}},
		{"zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"negative numbers", []int{-1, -2, -3}, []int{-2, -4, -6}},
		{"large slice", []int{10, 20, 30, 40, 50}, []int{20, 40, 60, 80, 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IterateWithIndex(tt.slice)
			if len(got) != len(tt.want) {
				t.Fatalf("IterateWithIndex(%v) length = %d, want %d", tt.slice, len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("IterateWithIndex(%v)[%d] = %d, want %d", tt.slice, i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestSliceFirst3(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"empty slice", []int{}, []int{}},
		{"one element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{1, 2}},
		{"exactly three elements", []int{1, 2, 3}, []int{1, 2, 3}},
		{"more than three elements", []int{1, 2, 3, 4, 5}, []int{1, 2, 3}},
		{"ten elements", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceFirst3(tt.slice)
			if len(got) != len(tt.want) {
				t.Fatalf("SliceFirst3(%v) length = %d, want %d", tt.slice, len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("SliceFirst3(%v)[%d] = %d, want %d", tt.slice, i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestSliceFrom2ToEnd(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"empty slice", []int{}, []int{}},
		{"one element", []int{1}, []int{}},
		{"two elements", []int{1, 2}, []int{}},
		{"three elements", []int{1, 2, 3}, []int{3}},
		{"five elements", []int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
		{"ten elements", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceFrom2ToEnd(tt.slice)
			if len(got) != len(tt.want) {
				t.Fatalf("SliceFrom2ToEnd(%v) length = %d, want %d", tt.slice, len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("SliceFrom2ToEnd(%v)[%d] = %d, want %d", tt.slice, i, got[i], tt.want[i])
				}
			}
		})
	}
}
