package slice_algorithms

import "slices"

// Map applies function to each element, returning new slice
func Map(slice []int, fn func(int) int) []int {
	// TODO(human): Apply fn to each element and return new slice

	result := []int{}

	for _, v := range slice {
		result = append(result, fn(v))
	}
	return result
}

// Filter keeps elements matching predicate, returning new slice
func Filter(slice []int, predicate func(int) bool) []int {
	// TODO(human): Return new slice containing only elements where predicate returns true
	result := []int{}

	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce reduces slice to single value using accumulator function
func Reduce(slice []int, initial int, fn func(int, int) int) int {
	// TODO(human): Apply fn to accumulate all elements into a single value

	for _, v := range slice {
		initial = fn(initial, v)
	}

	return initial
}

// Find returns first matching element and true, or zero value and false
func Find(slice []int, predicate func(int) bool) (int, bool) {
	// TODO(human): Return first element where predicate returns true

	for _, v := range slice {
		if predicate(v) {
			return v, true
		}
	}

	return 0, false
}

// Contains checks if value exists in slice
func Contains(slice []int, value int) bool {
	// TODO(human): Return true if value exists in slice

	return slices.Contains(slice, value)
}

// Unique returns slice with duplicates removed (maintains order)
func Unique(slice []int) []int {
	// TODO(human): Return new slice with duplicates removed, preserving first occurrence order

	if len(slice) == 0 {
		return nil
	}

	seen := map[int]bool{}
	result := []int{}

	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

// Partition splits slice into two slices based on predicate
// First slice contains elements matching predicate, second contains rest
func Partition(slice []int, predicate func(int) bool) ([]int, []int) {
	// TODO(human): Split slice into two based on predicate result

	sliceA := []int{}
	sliceB := []int{}

	if len(slice) == 0 {
		return nil, nil
	}

	for _, v := range slice {
		if predicate(v) {
			sliceA = append(sliceA, v)
		} else {
			sliceB = append(sliceB, v)
		}
	}
	return sliceA, sliceB
}
