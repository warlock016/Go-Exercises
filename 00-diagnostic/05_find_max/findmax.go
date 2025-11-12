package findmax

import "fmt"

// FindMax returns the maximum value in the slice
// Returns an error if the slice is empty
func FindMax(numbers []int) (int, error) {
	// TODO(human): Implement finding maximum value

	if len(numbers) == 0 {
		return 0, fmt.Errorf("empty slice")
	}

	maxVal := numbers[0]

	for idx := range numbers {
		if numbers[idx] > maxVal {
			maxVal = numbers[idx]
		}
	}

	return maxVal, nil
}
