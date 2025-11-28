package merge_sort

// MergeSort sorts a slice using the merge sort algorithm
func MergeSort(nums []int) []int {
	// TODO(human): Implement

	if len(nums) <= 1 {
		return nums
	}

	left, right := Divide(nums)

	sortedLeft := MergeSort(left)
	sortedRight := MergeSort(right)

	return Merge(sortedLeft, sortedRight)
}

func Divide(input []int) ([]int, []int) {
	if len(input) <= 1 {
		return input, nil
	}

	mid := len(input) / 2

	return input[:mid], input[mid:]
}

// Merge combines two sorted slices into one sorted slice
func Merge(left, right []int) []int {
	// TODO(human): Implement
	result := make([]int, 0, len(left)+len(right))

	if len(left) == 0 && len(right) == 0 {
		return []int{}
	}

	if len(left) == 0 {
		return right
	}

	if len(right) == 0 {
		return left
	}

	leftElement := left[0]
	rightElement := right[0]

	if leftElement > rightElement {
		result = append(result, rightElement)
		right = right[1:]
	} else {
		result = append(result, leftElement)
		left = left[1:]
	}

	result = append(result, Merge(left, right)...)

	return result
}
