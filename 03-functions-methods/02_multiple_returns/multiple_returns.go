package multiple_returns

import (
	"fmt"
	"strings"
)

// Divide returns the quotient of a divided by b
func Divide(a, b float64) (float64, error) {
	// TODO(human): Implement

	if b == 0 {
		return 0, fmt.Errorf("illegal: invalid zero division")
	}
	return a / b, nil
}

// ParseName extracts first and last name from a full name
func ParseName(fullName string) (firstName string, lastName string, err error) {
	// TODO(human): Implement

	result := strings.Fields(fullName)
	err = nil

	switch len(result) {
	case 0:
		return firstName, lastName, fmt.Errorf("error: zero-length field array extraction")
	case 1:
		return firstName, lastName, fmt.Errorf("error: extracted single field, failing output condition")
	case 2:
		firstName = result[0]
		lastName = result[1]
		return firstName, lastName, nil
	default:
		var fName strings.Builder

		for i, v := range result {
			if i != len(result)-1 {
				fName.WriteString(v)
				if i != len(result)-2 {
					fName.WriteString(" ")
				}
			}
		}

		firstName = fName.String()
		lastName = result[len(result)-1]

		return firstName, lastName, err
	}

	// return result[0], result[1], err
}

// Stats calculates min, max, and average from a slice
func Stats(numbers []int) (min int, max int, avg float64, err error) {
	// TODO(human): Implement

	switch len(numbers) {
	case 0:
		return 0, 0, 0, fmt.Errorf("function argument: empty slice")
	case 1:
		min, max, avg = numbers[0], numbers[0], float64(numbers[0])
		return min, max, avg, nil
	default:
		sum := 0
		min, max = numbers[0], numbers[0]
		for _, v := range numbers {
			sum += v

			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}

		avg = float64(sum) / float64(len(numbers))
	}
	return min, max, avg, nil
}

// FindIndex returns the index and whether the target was found
func FindIndex(slice []int, target int) (int, bool) {
	// TODO(human): Implement
	switch len(slice) {
	case 0:
		return 0, false
	default:
		for i, v := range slice {
			if v == target {
				return i, true
			}
		}
	}
	return 0, false
}
