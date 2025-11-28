package binary_search_recursive

// BinarySearch finds the index of target in sorted slice nums, returns -1 if not found
func BinarySearch(nums []int, target int) int {
	// TODO(human): Implement with helper function

	if len(nums) == 0 {
		return -1
	}

	return binaryHelper(nums, 0, len(nums)-1, target)
}

func binaryHelper(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return binaryHelper(nums, mid+1, right, target)
	} else {
		return binaryHelper(nums, left, mid-1, target)
	}
}

// FindFirst finds the first (leftmost) occurrence of target in sorted slice with duplicates
func FindFirst(nums []int, target int) int {
	// TODO(human): Implement with helper function

	if len(nums) == 0 {
		return 0
	}
	return -1
}

// FindLast finds the last (rightmost) occurrence of target in sorted slice with duplicates
func FindLast(nums []int, target int) int {
	// TODO(human): Implement with helper function
	return -1
}

// SearchInsertPosition finds the index where target would be inserted to maintain sorted order
func SearchInsertPosition(nums []int, target int) int {
	// TODO(human): Implement with helper function
	return 0
}
