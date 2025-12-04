package variadic_functions

import (
	"fmt"
	"strings"
)

// Sum returns the sum of all provided integers
func Sum(nums ...int) int {
	// TODO(human): Implement

	var result int

	for _, v := range nums {
		result += v
	}
	return result
}

// Max returns the maximum value from the provided integers
func Max(nums ...int) (int, error) {
	// TODO(human): Implement

	// 1. split the input array into two parts
	// if length(part) == 2 -> compare first against second slice value and return the max val
	// else if length(part) == 1 -> return slice val
	// else split each part.

	// if len(nums) == 0 {
	// 	return 0, fmt.Errorf("function argument: empty slice")
	// } else if len(nums) == 1 {
	// 	return nums[0], nil
	// } else if len(nums) == 2 {

	// }

	// var result int, nil

	// [1,2,3] // odd slice length unique
	// [1,1,3] // odd slice length repeated
	// [0,1,2,2] // even slice length repeated
	// [1,2,3,4] // even slice length unique

	var result int

	switch len(nums) {
	case 0:
		return result, fmt.Errorf("function argument: empty slice")

	case 1:
		result = nums[0]

	case 2:
		result = max(nums[0], nums[1])
	default:
		mid := len(nums) / 2
		partA := nums[:mid]
		partB := nums[mid:]

		resA, erra := Max(partA...)
		resB, errb := Max(partB...)

		if erra == nil && errb == nil {
			result = max(resA, resB)
		}
	}

	return result, nil
}

// Concat joins strings with a separator
func Concat(separator string, parts ...string) string {
	// TODO(human): Implement

	// [apple, orange, kiwi]; len = 3
	// [orange, kiwi]; len = 2
	// [apple]; len = 1
	// []; len = 0
	// ...

	var result strings.Builder

	for i, part := range parts {
		result.WriteString(part)

		// we write the separator only when the slice contains more than one element
		// AND
		// the current index is not pointing to the last item of the slice.

		if len(parts) > 1 && i != len(parts)-1 {
			result.WriteString(separator)
		}
	}

	return result.String()
}
