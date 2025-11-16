package slice_internals

import (
	"testing"
)

func TestSharesBackingArray(t *testing.T) {
	tests := []struct {
		name  string
		setup func() ([]int, []int)
		want  bool
	}{
		{
			"sliced from same array",
			func() ([]int, []int) {
				s1 := []int{1, 2, 3, 4, 5}
				s2 := s1[1:4]
				return s1, s2
			},
			true,
		},
		{
			"completely separate slices",
			func() ([]int, []int) {
				s1 := []int{1, 2, 3}
				s2 := []int{1, 2, 3}
				return s1, s2
			},
			false,
		},
		{
			"deep copied slice",
			func() ([]int, []int) {
				s1 := []int{1, 2, 3}
				s2 := make([]int, len(s1))
				copy(s2, s1)
				return s1, s2
			},
			false,
		},
		{
			"empty slices",
			func() ([]int, []int) {
				return []int{}, []int{}
			},
			false,
		},
		{
			"one empty slice",
			func() ([]int, []int) {
				s1 := []int{1, 2, 3}
				s2 := []int{}
				return s1, s2
			},
			false,
		},
		{
			"same slice variable",
			func() ([]int, []int) {
				s := []int{1, 2, 3}
				return s, s
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s1, s2 := tt.setup()
			got := SharesBackingArray(s1, s2)
			if got != tt.want {
				t.Errorf("SharesBackingArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDemonstrateCapacityGrowth(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"small growth", 5},
		{"medium growth", 10},
		{"larger growth", 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			capacities := DemonstrateCapacityGrowth(tt.n)

			if len(capacities) != tt.n {
				t.Fatalf("DemonstrateCapacityGrowth(%d) length = %d, want %d", tt.n, len(capacities), tt.n)
			}

			// Check capacities are non-decreasing
			for i := 1; i < len(capacities); i++ {
				if capacities[i] < capacities[i-1] {
					t.Errorf("DemonstrateCapacityGrowth(%d)[%d] = %d < [%d] = %d (capacity decreased)",
						tt.n, i, capacities[i], i-1, capacities[i-1])
				}
			}

			// First append should create capacity of at least 1
			if capacities[0] < 1 {
				t.Errorf("DemonstrateCapacityGrowth(%d)[0] = %d, want >= 1", tt.n, capacities[0])
			}
		})
	}
}

func TestSafeAppend(t *testing.T) {
	original := []int{1, 2, 3}
	result := SafeAppend(original, 4)

	// Check result
	want := []int{1, 2, 3, 4}
	if len(result) != len(want) {
		t.Fatalf("SafeAppend() length = %d, want %d", len(result), len(want))
	}

	for i := range want {
		if result[i] != want[i] {
			t.Errorf("SafeAppend()[%d] = %d, want %d", i, result[i], want[i])
		}
	}

	// Check original unchanged
	if len(original) != 3 {
		t.Errorf("SafeAppend() modified original length = %d, want 3", len(original))
	}

	// Check no backing array sharing
	if SharesBackingArray(original, result) {
		t.Error("SafeAppend() result shares backing array with original")
	}
}

func TestDeepCopy(t *testing.T) {
	tests := []struct {
		name     string
		original []int
	}{
		{"small slice", []int{1, 2, 3}},
		{"single element", []int{42}},
		{"empty slice", []int{}},
		{"larger slice", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copied := DeepCopy(tt.original)

			// Check contents match
			if len(copied) != len(tt.original) {
				t.Fatalf("DeepCopy() length = %d, want %d", len(copied), len(tt.original))
			}

			for i := range tt.original {
				if copied[i] != tt.original[i] {
					t.Errorf("DeepCopy()[%d] = %d, want %d", i, copied[i], tt.original[i])
				}
			}

			// Check independence (if non-empty)
			if len(tt.original) > 0 {
				if SharesBackingArray(tt.original, copied) {
					t.Error("DeepCopy() shares backing array with original")
				}

				// Modify copy, ensure original unchanged
				copied[0] = 9999
				if tt.original[0] == 9999 {
					t.Error("DeepCopy() modification affected original")
				}
			}
		})
	}
}

func TestFullSliceExpression(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	tests := []struct {
		name    string
		low     int
		high    int
		max     int
		wantLen int
		wantCap int
	}{
		{"middle section", 2, 5, 5, 3, 3},
		{"with extra capacity", 1, 4, 6, 3, 5},
		{"from start", 0, 3, 3, 3, 3},
		{"single element", 5, 6, 6, 1, 1},
		{"empty slice", 3, 3, 3, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FullSliceExpression(s, tt.low, tt.high, tt.max)

			if len(result) != tt.wantLen {
				t.Errorf("FullSliceExpression(%v, %d, %d, %d) length = %d, want %d",
					s, tt.low, tt.high, tt.max, len(result), tt.wantLen)
			}

			if cap(result) != tt.wantCap {
				t.Errorf("FullSliceExpression(%v, %d, %d, %d) capacity = %d, want %d",
					s, tt.low, tt.high, tt.max, cap(result), tt.wantCap)
			}

			// Check values
			for i := 0; i < len(result); i++ {
				expected := s[tt.low+i]
				if result[i] != expected {
					t.Errorf("FullSliceExpression(%v, %d, %d, %d)[%d] = %d, want %d",
						s, tt.low, tt.high, tt.max, i, result[i], expected)
				}
			}
		})
	}
}

func TestTrimSlice(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		index   int
		wantLen int
		wantCap int
	}{
		{"trim half", []int{1, 2, 3, 4, 5}, 3, 3, 3},
		{"trim to one", []int{1, 2, 3}, 1, 1, 1},
		{"trim to empty", []int{1, 2, 3}, 0, 0, 0},
		{"trim nothing", []int{1, 2, 3}, 3, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TrimSlice(tt.slice, tt.index)

			if len(result) != tt.wantLen {
				t.Errorf("TrimSlice() length = %d, want %d", len(result), tt.wantLen)
			}

			if cap(result) != tt.wantCap {
				t.Errorf("TrimSlice() capacity = %d, want %d", cap(result), tt.wantCap)
			}

			// Check values
			for i := 0; i < len(result); i++ {
				if result[i] != tt.slice[i] {
					t.Errorf("TrimSlice()[%d] = %d, want %d", i, result[i], tt.slice[i])
				}
			}
		})
	}
}

func TestExtendSlice(t *testing.T) {
	tests := []struct {
		name              string
		initial           []int
		values            []int
		expectRealloc     bool
		wantLen           int
	}{
		{
			"within capacity",
			make([]int, 2, 10),
			[]int{1, 2, 3},
			false,
			5,
		},
		{
			"exceeds capacity",
			make([]int, 3, 3),
			[]int{1, 2, 3},
			true,
			6,
		},
		{
			"exactly at capacity",
			make([]int, 5, 5),
			[]int{1},
			true,
			6,
		},
		{
			"empty to non-empty",
			[]int{},
			[]int{1, 2},
			true,
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, reallocated := ExtendSlice(tt.initial, tt.values...)

			if len(result) != tt.wantLen {
				t.Errorf("ExtendSlice() length = %d, want %d", len(result), tt.wantLen)
			}

			if reallocated != tt.expectRealloc {
				t.Errorf("ExtendSlice() reallocated = %v, want %v", reallocated, tt.expectRealloc)
			}
		})
	}
}

func TestCountSliceAllocations(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		maxExpected int
	}{
		{"small append", 5, 5},
		{"medium append", 10, 10},
		{"larger append", 100, 20}, // Should be much less than 100 due to growth
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := CountSliceAllocations(tt.n)

			if count < 0 {
				t.Errorf("CountSliceAllocations(%d) = %d, should be non-negative", tt.n, count)
			}

			if count > tt.maxExpected {
				t.Errorf("CountSliceAllocations(%d) = %d, expected <= %d (inefficient growth)",
					tt.n, count, tt.maxExpected)
			}

			// Should have at least one allocation (unless n=0)
			if tt.n > 0 && count == 0 {
				t.Errorf("CountSliceAllocations(%d) = 0, expected > 0", tt.n)
			}
		})
	}
}

func TestPreallocatedSlice(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
	}{
		{"small capacity", 5},
		{"medium capacity", 100},
		{"large capacity", 1000},
		{"zero capacity", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PreallocatedSlice(tt.capacity)

			if len(s) != 0 {
				t.Errorf("PreallocatedSlice(%d) length = %d, want 0", tt.capacity, len(s))
			}

			if cap(s) != tt.capacity {
				t.Errorf("PreallocatedSlice(%d) capacity = %d, want %d", tt.capacity, cap(s), tt.capacity)
			}

			// Test that we can append without reallocation
			if tt.capacity > 0 {
				for i := 0; i < tt.capacity; i++ {
					oldCap := cap(s)
					s = append(s, i)
					if cap(s) != oldCap {
						t.Errorf("PreallocatedSlice(%d) reallocated at append %d", tt.capacity, i)
						break
					}
				}
			}
		})
	}
}

func TestSliceModificationWithSharing(t *testing.T) {
	// Test that demonstrates the danger of shared backing arrays
	original := []int{1, 2, 3, 4, 5}
	slice1 := original[0:3]  // [1, 2, 3]
	slice2 := original[2:5]  // [3, 4, 5]

	// Verify they share backing array
	if !SharesBackingArray(original, slice1) {
		t.Error("slice1 should share backing array with original")
	}

	if !SharesBackingArray(original, slice2) {
		t.Error("slice2 should share backing array with original")
	}

	// Modifying slice1 affects original and slice2 (where they overlap)
	slice1[2] = 99

	if original[2] != 99 {
		t.Error("Modifying slice1[2] should affect original[2]")
	}

	if slice2[0] != 99 {
		t.Error("Modifying slice1[2] should affect slice2[0] (same element)")
	}
}
