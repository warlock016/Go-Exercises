package slice_internals

// SharesBackingArray returns true if s1 and s2 share the same backing array
func SharesBackingArray(s1, s2 []int) bool {
	// TODO(human): Compare backing array addresses
	return false
}

// DemonstrateCapacityGrowth returns a slice showing how capacity grows during appends
func DemonstrateCapacityGrowth(n int) []int {
	// TODO(human): Track capacity changes during appends
	return nil
}

// SafeAppend creates a new slice that doesn't share backing array with original
// This prevents append from affecting the original slice
func SafeAppend(original []int, value int) []int {
	// TODO(human): Create independent copy and append
	return nil
}

// DeepCopy creates a completely independent copy of the slice
// No shared backing array
func DeepCopy(s []int) []int {
	// TODO(human): Create independent copy
	return nil
}

// FullSliceExpression returns slice with restricted capacity using three-index syntax
func FullSliceExpression(s []int, low, high, max int) []int {
	// TODO(human): Use three-index slice expression
	return nil
}

// TrimSlice removes elements from index i onwards with restricted capacity
func TrimSlice(s []int, i int) []int {
	// TODO(human): Trim slice and restrict capacity
	return nil
}

// ExtendSlice appends values and reports if reallocation occurred
func ExtendSlice(s []int, values ...int) ([]int, bool) {
	// TODO(human): Append and detect reallocation
	return nil, false
}

// CountSliceAllocations returns number of reallocations when appending n elements
func CountSliceAllocations(n int) int {
	// TODO(human): Count reallocations during n appends
	return 0
}

// PreallocatedSlice creates a slice with capacity to hold n elements without reallocation
func PreallocatedSlice(n int) []int {
	// TODO(human): Create slice with preallocated capacity
	return nil
}
