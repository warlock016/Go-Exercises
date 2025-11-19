package performance_optimization

import "strings"

// AppendNaive appends n integers (0 to n-1) to a slice without preallocation
func AppendNaive(n int) []int {
	// TODO(human): Implement without preallocation
	result := []int{}

	for i := range n {
		result = append(result, i)
	}
	return result
} // 302.9 ns/op

// AppendOptimized appends n integers (0 to n-1) with proper preallocation
func AppendOptimized(n int) []int {
	// TODO(human): Implement with preallocation
	result := make([]int, 0, n)

	for i := range n {
		result = append(result, i)
	}
	return result
} // 47.22 ns/op -> ~85% faster

// ConcatStringsNaive concatenates n copies of str using the + operator
func ConcatStringsNaive(str string, n int) string {
	// TODO(human): Implement using + operator

	var result string

	for range n {
		result += str
	}

	return result
} // 2252 ns/op

// ConcatStringsOptimized concatenates n copies of str using strings.Builder
func ConcatStringsOptimized(str string, n int) string {
	// TODO(human): Implement using strings.Builder

	var result strings.Builder

	for range n {
		result.WriteString(str)
	}

	return result.String()
} // 605.6 ns/op -> ~73% faster

// FilterNaive filters a slice, keeping only even numbers (no capacity hint)
func FilterNaive(numbers []int) []int {
	// TODO(human): Implement without capacity hint

	result := []int{}

	for _, v := range numbers {
		if v%2 == 0 {
			result = append(result, v)
		}
	}

	return result
} // 1703 ns/op

// FilterOptimized filters a slice, keeping only even numbers (with capacity hint)
func FilterOptimized(numbers []int) []int {
	// TODO(human): Implement with capacity hint

	srcLen := len(numbers) / 2

	result := make([]int, 0, srcLen)

	for _, v := range numbers {
		if v%2 == 0 {
			result = append(result, v)
		}
	}

	return result
} // 1262 ns/op -> ~25% faster
