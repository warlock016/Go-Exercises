package helper_functions

// FactorialAccum calculates factorial using accumulator pattern
func FactorialAccum(n int) int {
	// TODO(human): Implement with helper function
	return factorialHelper(n, 1)
}

// 3,1 -> 2, 3*1 -> 1, 2*3*1 -> 0, 2*3*1
func factorialHelper(n, accum int) int {
	if n == 0 {
		return accum
	}
	return factorialHelper(n-1, n*accum)
}

// ReverseList reverses a slice using accumulator pattern
func ReverseList(nums []int) []int {
	// TODO(human): Implement with helper function
	return reverseHelper(nums, []int{})
}

// [0, 1, 2, 3]
func reverseHelper(nums, acc []int) []int {
	if len(nums) == 0 {
		return acc
	}

	acc = append(acc, nums[len(nums)-1])
	nums = nums[:len(nums)-1]

	return reverseHelper(nums, acc)
}

// RangeSum calculates sum from start to end using accumulator
func RangeSum(start, end int) int {
	// TODO(human): Implement with helper function
	return 0 + rangeHelper(start, end, 0)
}

// 0, 1, 2, 3 : start = 0, end = 3
func rangeHelper(start, end, acc int) int {

	// fmt.Printf("start: %d, end: %d, acc: %d\n", start, end, acc)
	if start > end {
		return acc
	}

	// acc += start
	// start += 1

	return rangeHelper(start+1, end, acc+start)
}

// StringToInt converts a numeric string to integer recursively
func StringToInt(s string) int {
	// TODO(human): Implement with helper function

	if len(s) == 0 {
		return 0
	}

	negative := false
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}

	result := intHelper(s, 0)

	if negative {
		return -result
	}

	return result
}

// 1990 == [1,9,9,0]
func intHelper(s string, acc int) int {

	if len(s) == 0 {
		return acc
	}

	digit := int(s[0] - '0')
	return intHelper(s[1:], acc*10+digit)
}
