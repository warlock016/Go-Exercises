package slice_algorithms

// Map applies function to each element, returning new slice
func Map(slice []int, fn func(int) int) []int {
	// TODO(human): Apply fn to each element and return new slice
	return nil
}

// Filter keeps elements matching predicate, returning new slice
func Filter(slice []int, predicate func(int) bool) []int {
	// TODO(human): Return new slice containing only elements where predicate returns true
	return nil
}

// Reduce reduces slice to single value using accumulator function
func Reduce(slice []int, initial int, fn func(int, int) int) int {
	// TODO(human): Apply fn to accumulate all elements into a single value
	return 0
}

// Find returns first matching element and true, or zero value and false
func Find(slice []int, predicate func(int) bool) (int, bool) {
	// TODO(human): Return first element where predicate returns true
	return 0, false
}

// Contains checks if value exists in slice
func Contains(slice []int, value int) bool {
	// TODO(human): Return true if value exists in slice
	return false
}

// Unique returns slice with duplicates removed (maintains order)
func Unique(slice []int) []int {
	// TODO(human): Return new slice with duplicates removed, preserving first occurrence order
	return nil
}

// Partition splits slice into two slices based on predicate
// First slice contains elements matching predicate, second contains rest
func Partition(slice []int, predicate func(int) bool) ([]int, []int) {
	// TODO(human): Split slice into two based on predicate result
	return nil, nil
}
