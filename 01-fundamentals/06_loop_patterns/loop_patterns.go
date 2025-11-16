package looppatterns

// CountUp returns a slice containing [1, 2, 3, ..., n]
// Edge case: if n <= 0, return empty slice
func CountUp(n int) []int {
	// TODO(human): Implement CountUp

	result := []int{}
	if n <= 0 {
		return result
	}

	for i := range n {
		result = append(result, i+1)
	}
	return result
}

// CountDown returns a slice containing [n, n-1, n-2, ..., 1]
// Edge case: if n <= 0, return empty slice
func CountDown(n int) []int {
	// TODO(human): Implement CountDown
	result := []int{}
	if n <= 0 {
		return result
	}

	for i := n; i > 0; i-- {
		result = append(result, i)
	}

	return result
}

// SumSlice returns the sum of all numbers in the slice
// Edge case: empty slice returns 0
func SumSlice(nums []int) int {
	// TODO(human): Implement SumSlice

	result := 0

	if len(nums) == 0 {
		return result
	}

	for _, v := range nums {
		result += v
	}

	return result
}

// FirstNEvens returns the first n even positive numbers [2, 4, 6, ...]
// Edge case: if n <= 0, return empty slice
func FirstNEvens(n int) []int {
	// TODO(human): Implement FirstNEvens
	result := []int{}
	if n <= 0 {
		return result
	}

	for i := 1; i <= n; i++ {
		result = append(result, i*2)
	}

	return result
}
