package slice_operations

import (
	"testing"
)

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name string
		src  []int
	}{
		{"empty slice", []int{}},
		{"single element", []int{42}},
		{"multiple elements", []int{1, 2, 3, 4, 5}},
		{"large slice", make([]int, 100)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CopySlice(tt.src)

			// Check length
			if len(got) != len(tt.src) {
				t.Errorf("CopySlice() length = %d, want %d", len(got), len(tt.src))
			}

			// Check values
			for i := range tt.src {
				if got[i] != tt.src[i] {
					t.Errorf("CopySlice()[%d] = %d, want %d", i, got[i], tt.src[i])
				}
			}

			// Check independence (modify copy shouldn't affect original)
			if len(got) > 0 {
				originalValue := tt.src[0]
				got[0] = 9999
				if tt.src[0] != originalValue {
					t.Errorf("CopySlice() created shared backing array, modifying copy affected original")
				}
			}
		})
	}
}

func TestAppendWithCapacityCheck(t *testing.T) {
	tests := []struct {
		name       string
		slice      []int
		value      int
		wantOk     bool
		wantLength int
	}{
		{"has capacity", make([]int, 3, 5), 99, true, 4},
		{"at capacity", make([]int, 5, 5), 99, false, 5},
		{"empty with capacity", make([]int, 0, 10), 1, true, 1},
		{"nil slice", nil, 1, false, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := AppendWithCapacityCheck(tt.slice, tt.value)
			if ok != tt.wantOk {
				t.Errorf("AppendWithCapacityCheck() ok = %v, want %v", ok, tt.wantOk)
			}
			if len(got) != tt.wantLength {
				t.Errorf("AppendWithCapacityCheck() length = %d, want %d", len(got), tt.wantLength)
			}
			if ok && got[len(got)-1] != tt.value {
				t.Errorf("AppendWithCapacityCheck() last element = %d, want %d", got[len(got)-1], tt.value)
			}
		})
	}
}

func TestExtendSlice(t *testing.T) {
	tests := []struct {
		name string
		slice []int
		wantLen int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 2},
		{"multiple elements", []int{1, 2, 3}, 6},
		{"even count", []int{10, 20, 30, 40}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.slice))
			copy(original, tt.slice)

			got := ExtendSlice(tt.slice)

			if len(got) != tt.wantLen {
				t.Errorf("ExtendSlice(%v) length = %d, want %d", tt.slice, len(got), tt.wantLen)
			}

			// Check original elements are preserved
			for i := range original {
				if got[i] != original[i] {
					t.Errorf("ExtendSlice(%v)[%d] = %d, want %d", tt.slice, i, got[i], original[i])
				}
			}

			// Check new elements are zero
			for i := len(original); i < len(got); i++ {
				if got[i] != 0 {
					t.Errorf("ExtendSlice(%v)[%d] = %d, want 0 (zero value)", tt.slice, i, got[i])
				}
			}
		})
	}
}

func TestTruncateSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		wantLen int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 0},
		{"two elements", []int{1, 2}, 1},
		{"odd count", []int{1, 2, 3, 4, 5}, 2},
		{"even count", []int{1, 2, 3, 4}, 2},
		{"six elements", []int{10, 20, 30, 40, 50, 60}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TruncateSlice(tt.slice)
			if len(got) != tt.wantLen {
				t.Errorf("TruncateSlice(%v) length = %d, want %d", tt.slice, len(got), tt.wantLen)
			}
			// Check elements match first half
			for i := 0; i < tt.wantLen; i++ {
				if got[i] != tt.slice[i] {
					t.Errorf("TruncateSlice(%v)[%d] = %d, want %d", tt.slice, i, got[i], tt.slice[i])
				}
			}
		})
	}
}

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"odd count", []int{1, 2, 3}, []int{3, 2, 1}},
		{"even count", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"five elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy since ReverseSlice modifies in place
			slice := make([]int, len(tt.slice))
			copy(slice, tt.slice)

			got := ReverseSlice(slice)

			if len(got) != len(tt.want) {
				t.Fatalf("ReverseSlice(%v) length = %d, want %d", tt.slice, len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("ReverseSlice(%v)[%d] = %d, want %d", tt.slice, i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestFilterEven(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"empty slice", []int{}, []int{}},
		{"all odd", []int{1, 3, 5}, []int{}},
		{"all even", []int{2, 4, 6}, []int{2, 4, 6}},
		{"mixed", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"single even", []int{1, 3, 4, 7}, []int{4}},
		{"with zero", []int{0, 1, 2, 3}, []int{0, 2}},
		{"negative numbers", []int{-2, -1, 0, 1, 2}, []int{-2, 0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterEven(tt.slice)
			if len(got) != len(tt.want) {
				t.Fatalf("FilterEven(%v) length = %d, want %d", tt.slice, len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("FilterEven(%v)[%d] = %d, want %d", tt.slice, i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestInsertAt(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		index int
		value int
		want  []int
	}{
		{"insert at beginning", []int{2, 3, 4}, 0, 1, []int{1, 2, 3, 4}},
		{"insert in middle", []int{1, 2, 4, 5}, 2, 3, []int{1, 2, 3, 4, 5}},
		{"insert at end", []int{1, 2, 3}, 3, 4, []int{1, 2, 3, 4}},
		{"insert beyond end", []int{1, 2}, 10, 99, []int{1, 2, 99}},
		{"insert in empty", []int{}, 0, 1, []int{1}},
		{"insert at index 1", []int{1, 3}, 1, 2, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertAt(tt.slice, tt.index, tt.value)
			if len(got) != len(tt.want) {
				t.Fatalf("InsertAt(%v, %d, %d) length = %d, want %d", tt.slice, tt.index, tt.value, len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("InsertAt(%v, %d, %d)[%d] = %d, want %d", tt.slice, tt.index, tt.value, i, got[i], tt.want[i])
				}
			}
		})
	}
}
