package slice_internals

import (
	"fmt"
)

// SharesBackingArray returns true if s1 and s2 share the same backing array
func SharesBackingArray(s1, s2 []int) bool {
	// TODO(human): Compare backing array addresses
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}

	for i := range s1 { // if we use values, we are comparing value copies instead of the underlying value addresses
		for j := range s2 { // that's why we use indexes instead of values! (i instead of _, v)
			if &s1[i] == &s2[j] {
				return true
			}
		}
	}
	return false
}

// DemonstrateCapacityGrowth returns a slice showing how capacity grows during appends
func DemonstrateCapacityGrowth(n int) []int {
	// TODO(human): Track capacity changes during appends
	new := []int{}
	capacities := []int{}

	for i := range n {
		new = append(new, i)
		capacities = append(capacities, cap(new))
		// new = slices.Grow(new, n)
		fmt.Printf("slice: %v, len: %d, cap: %d\n", capacities, len(capacities), cap(capacities))
	}

	return capacities
}

// SafeAppend creates a new slice that doesn't share backing array with original
// This prevents append from affecting the original slice
func SafeAppend(original []int, value int) []int {
	// TODO(human): Create independent copy and append
	result := make([]int, cap(original))
	copy(result, original)
	result = append(result, value)

	return result
}

// DeepCopy creates a completely independent copy of the slice
// No shared backing array
func DeepCopy(s []int) []int {
	// TODO(human): Create independent copy
	result := make([]int, cap(s))
	copy(result, s)

	return result
}

// FullSliceExpression returns slice with restricted capacity using three-index syntax
func FullSliceExpression(s []int, low, high, max int) []int {
	// TODO(human): Use three-index slice expression
	return s[low:high:max]
}

// TrimSlice removes elements from index i onwards with restricted capacity
// [1, 2, 3, 4, 5]
// i = 3, len = 5, cap = 5
// result -> [1, 2, 3]
// len = 3, cap = 3
func TrimSlice(s []int, i int) []int {
	// TODO(human): Trim slice and restrict capacity
	if i < 0 || i >= len(s) {
		return s
	}

	return s[:i:i]
}

// ExtendSlice appends values and reports if reallocation occurred
func ExtendSlice(s []int, values ...int) ([]int, bool) {
	// TODO(human): Append and detect reallocation

	// oldCap := cap(s)
	// oldCap := &s[0]

	var oldPrt *int

	if len(s) > 0 {
		oldPrt = &s[0]
	}

	s = append(s, values...)
	newPtr := &s[0]

	if oldPrt != newPtr {
		return s, true
	}

	return s, false
}

// CountSliceAllocations returns number of reallocations when appending n elements
func CountSliceAllocations(n int) int {
	// TODO(human): Count reallocations during n appends

	cap := 1
	count := 1

	for range n {
		if cap < n {
			cap *= 2
			count++
		}
	}
	return count
}

// PreallocatedSlice creates a slice with capacity to hold n elements without reallocation
func PreallocatedSlice(n int) []int {
	// TODO(human): Create slice with preallocated capacity
	result := make([]int, 0, n)

	// fmt.Printf("%v, len: %d, cap: %d\n", result, len(result), cap(result))
	return result
}
