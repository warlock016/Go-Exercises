package slice_basics

// CreateSliceWithMake creates a slice of integers using make with specified length and capacity
func CreateSliceWithMake(length, capacity int) []int {
	// TODO(human): Use make([]int, length, capacity) to create and return a slice

	return make([]int, length, capacity)
}

// CreateSliceLiteral creates a slice containing the numbers 1, 2, 3, 4, 5
func CreateSliceLiteral() []int {
	// TODO(human): Use slice literal syntax []int{...} to create and return [1, 2, 3, 4, 5]
	return []int{1, 2, 3, 4, 5}
}

// AppendToSlice appends the values to the slice and returns the result
func AppendToSlice(slice []int, values ...int) []int {
	// TODO(human): Use append(slice, values...) to append all values and return the result
	// Remember: append returns a new slice, so you must return it

	slice = append(slice, values...)
	return slice
}

// GetLengthAndCapacity returns both the length and capacity of the slice
func GetLengthAndCapacity(slice []int) (length, capacity int) {
	// TODO(human): Use len(slice) and cap(slice) to return both values
	return len(slice), cap(slice)
}

// IterateAndSum returns the sum of all elements in the slice
func IterateAndSum(slice []int) int {
	// TODO(human): Use a for-range loop to iterate through slice and sum all values
	// for _, val := range slice { ... }

	sum := 0

	for _, v := range slice {
		sum += v
	}
	return sum
}

// IterateWithIndex returns a new slice where each element is doubled
func IterateWithIndex(slice []int) []int {
	// TODO(human): Create a new slice with make([]int, len(slice))
	// Then use for i, val := range slice to iterate and double each value

	// this version is less performant due to multiple allocations during runtime when filling out the slice
	// doubled := []int{}
	// for _, v := range slice {
	// 	doubled = append(doubled, 2*v)
	// }

	// This version is better in the current context since we know in advance the size of the slice, so we skip multiple reallocations during runtime
	doubled := make([]int, len(slice)) // -> failed due to the creation of extra slots.
	for i := range slice {
		doubled[i] = slice[i] * 2
	}

	return doubled
}

// SliceFirst3 returns a new slice containing only the first 3 elements
// If the slice has fewer than 3 elements, return the entire slice
func SliceFirst3(slice []int) []int {
	// TODO(human): Check if len(slice) < 3, if so return slice
	// Otherwise return slice[:3]

	if len(slice) < 3 {
		return slice
	}

	return slice[:3]
}

// SliceFrom2ToEnd returns a new slice from index 2 to the end
// If the slice has fewer than 2 elements, return an empty slice
func SliceFrom2ToEnd(slice []int) []int {
	// TODO(human): Check if len(slice) < 2, if so return []int{}
	// Otherwise return slice[2:]

	if len(slice) < 2 {
		return []int{}
	}

	return slice[2:]
}
