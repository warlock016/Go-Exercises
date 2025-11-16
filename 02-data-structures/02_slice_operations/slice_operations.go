package slice_operations

// CopySlice creates a true copy of the slice (not shared backing array)
func CopySlice(src []int) []int {
	// TODO(human): Create a new slice with make([]int, len(src))
	// Then use copy(dest, src) to copy elements
	// Return the new independent slice
	new := make([]int, len(src))

	if len(src) == 0 {
		return new
	}

	copy(new, src)

	return new
}

// AppendWithCapacityCheck appends value to slice only if there's capacity
// Returns the slice and a bool indicating if append happened
func AppendWithCapacityCheck(slice []int, value int) ([]int, bool) {
	// TODO(human): Check if len(slice) < cap(slice)
	// If true: append and return (newSlice, true)
	// If false: return (slice, false) without appending

	if len(slice) < cap(slice) {
		slice = append(slice, value)
		return slice, true
	}

	return slice, false
}

// ExtendSlice extends the slice to double its current length
// New elements should be zero-initialized
func ExtendSlice(slice []int) []int {
	// TODO(human): Create a new slice with make([]int, len(slice)*2)
	// Use copy() to copy original elements to the first half
	// Return the extended slice (second half will be zero-initialized)
	new := make([]int, 2*len(slice))
	copy(new, slice)

	return new
}

// TruncateSlice returns a new slice with only the first half of elements
func TruncateSlice(slice []int) []int {
	// TODO(human): Calculate midpoint: mid := len(slice) / 2
	// Return slice[:mid]
	center := len(slice) / 2
	return slice[0:center]
}

// ReverseSlice reverses the slice in place and returns it
func ReverseSlice(slice []int) []int {
	// TODO(human): Use two pointers: i starting at 0, j starting at len(slice)-1
	// While i < j: swap slice[i] and slice[j], then i++, j--
	// Return the slice

	i := 0
	j := len(slice) - 1

	var a int
	// var b int

	for i < j {
		a = slice[i]
		slice[i] = slice[j]
		slice[j] = a
		i++
		j--
	}

	return slice
}

// FilterEven returns a new slice containing only even numbers
func FilterEven(slice []int) []int {
	// TODO(human): Create an empty result slice: result := []int{}
	// Iterate through slice with range
	// If val % 2 == 0, append it to result
	// Return result

	new := []int{}

	for _, v := range slice {
		if v%2 == 0 {
			new = append(new, v)
		}
	}

	return new
}

// InsertAt inserts value at the specified index (shifting elements right)
// If index is out of bounds, append to the end
func InsertAt(slice []int, index, value int) []int {
	// TODO(human): If index >= len(slice), return append(slice, value)
	// Otherwise:
	//   1. Append 0 to make room: slice = append(slice, 0)
	//   2. Shift elements right: copy(slice[index+1:], slice[index:])
	//   3. Insert value: slice[index] = value
	//   4. Return slice

	if index < 0 {
		return slice
	}

	if index >= len(slice) {
		slice = append(slice, value)
		return slice
	}

	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])
	slice[index] = value

	return slice
}
