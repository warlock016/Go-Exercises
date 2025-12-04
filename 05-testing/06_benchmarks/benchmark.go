package benchmark

import "strings"

// ConcatStrings concatenates strings using the + operator
func ConcatStrings(strs []string) string {
	result := ""
	for _, s := range strs {
		result += s
	}
	return result
}

// ConcatBuilder concatenates strings using strings.Builder
func ConcatBuilder(strs []string) string {
	var builder strings.Builder
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}

// FibonacciRecursive calculates Fibonacci number recursively (inefficient)
func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

// FibonacciIterative calculates Fibonacci number iteratively (efficient)
func FibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
