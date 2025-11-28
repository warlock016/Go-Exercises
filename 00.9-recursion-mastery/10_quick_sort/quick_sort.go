package quick_sort

// QuickSort sorts a slice using the quick sort algorithm
func QuickSort(nums []int) []int {
	// TODO(human): Implement

	if len(nums) <= 1 {
		return nums
	}

	quickSortHelper(nums, 0, len(nums)-1)
	return nums
}

func quickSortHelper(nums []int, low, high int) {
	if low < high {
		pivotIndex := Partition(nums, low, high)

		quickSortHelper(nums, low, pivotIndex-1)
		quickSortHelper(nums, pivotIndex+1, high)
	}
}

// Partition rearranges elements: smaller than pivot on left, larger on right
func Partition(nums []int, low, high int) int {
	// TODO(human): Implement

	pivot := nums[high]
	i := low - 1

	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	i++
	nums[i], nums[high] = nums[high], nums[i]

	return i
}
