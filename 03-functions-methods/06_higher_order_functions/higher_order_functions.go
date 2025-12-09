package higher_order_functions

import "slices"

// Map transforms each element using the given function
func Map(slice []int, fn func(int) int) []int {
	// TODO(human): Implement
	result := make([]int, 0, len(slice))

	for _, v := range slice {
		calc := fn(v)
		result = append(result, calc)
	}
	return result
}

// Filter returns elements that match the predicate
func Filter(slice []int, predicate func(int) bool) []int {
	// TODO(human): Implement

	if len(slice) == 0 {
		return []int{}
	}

	result := make([]int, 0, len(slice))

	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce accumulates values using the given function
func Reduce(slice []int, initial int, fn func(int, int) int) int {
	// TODO(human): Implement

	var result int = initial

	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// Compose returns a function that applies g then f
func Compose(f func(int) int, g func(int) int) func(int) int {
	// TODO(human): Implement
	return func(i int) int {
		j := g(i)
		return f(j)
	}
}

// ForEach executes a function on each element
func ForEach(slice []int, fn func(int)) {
	// TODO(human): Implement
	for _, exec := range slice {
		fn(exec)
	}
}

// Any checks if at least one element matches the predicate
func Any(slice []int, predicate func(int) bool) bool {
	// TODO(human): Implement
	return slices.ContainsFunc(slice, predicate)
}

// All checks if all elements match the predicate
func All(slice []int, predicate func(int) bool) bool {
	// TODO(human): Implement
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}
