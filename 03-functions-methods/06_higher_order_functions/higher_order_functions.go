package higher_order_functions

// Map transforms each element using the given function
func Map(slice []int, fn func(int) int) []int {
	// TODO(human): Implement
	return nil
}

// Filter returns elements that match the predicate
func Filter(slice []int, predicate func(int) bool) []int {
	// TODO(human): Implement
	return nil
}

// Reduce accumulates values using the given function
func Reduce(slice []int, initial int, fn func(int, int) int) int {
	// TODO(human): Implement
	return 0
}

// Compose returns a function that applies g then f
func Compose(f func(int) int, g func(int) int) func(int) int {
	// TODO(human): Implement
	return nil
}

// ForEach executes a function on each element
func ForEach(slice []int, fn func(int)) {
	// TODO(human): Implement
}

// Any checks if at least one element matches the predicate
func Any(slice []int, predicate func(int) bool) bool {
	// TODO(human): Implement
	return false
}

// All checks if all elements match the predicate
func All(slice []int, predicate func(int) bool) bool {
	// TODO(human): Implement
	return false
}
