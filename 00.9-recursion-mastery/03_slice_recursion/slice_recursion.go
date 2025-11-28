package slice_recursion

// Sum calculates the sum of all integers in the slice
func Sum(nums []int) int {
	// TODO(human): Implement
	if len(nums) == 0 {
		return 0
	}

	// if len(nums) == 1 {
	// 	return nums[0]
	// }

	current := nums[0]
	nums = nums[1:]

	return current + Sum(nums)
}

// Max finds the maximum value in a non-empty slice
func Max(nums []int) int {
	// TODO(human): Implement

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	current := nums[0]
	new := nums[1:]

	return max(current, Max(new))
}

// Contains checks if the slice contains the target value
func Contains(nums []int, target int) bool {
	// TODO(human): Implement

	if len(nums) == 0 {
		return false
	}

	// if len(nums) == 1 && nums[0] == target {
	// 	return true
	// }

	current := nums[0]
	new := nums[1:]

	return current == target || Contains(new, target)
}

// CountOccurrences counts how many times target appears in the slice
func CountOccurrences(nums []int, target int) int {
	// TODO(human): Implement

	var counter int

	if len(nums) == 0 {
		return 0
	}

	current := nums[0]
	new := nums[1:]

	counter += CountOccurrences(new, target)

	if current == target {
		counter++
	}

	return counter
}
