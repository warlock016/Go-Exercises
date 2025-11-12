package sum

// Sum returns the sum of all numbers in the slice
func Sum(numbers []int) int {
	// TODO(human): Implement sum calculation
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
