package filter

// FilterEven returns a new slice containing only the even numbers
func FilterEven(numbers []int) []int {
	// TODO(human): Implement filtering for even numbers
	var filtered []int

	for _, val := range numbers {
		if val%2 == 0 {
			filtered = append(filtered, val)
		}
	}

	return filtered
}
