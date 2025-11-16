package performance_optimization

// AppendNaive appends n integers (0 to n-1) to a slice without preallocation
func AppendNaive(n int) []int {
	// TODO(human): Implement without preallocation
	return nil
}

// AppendOptimized appends n integers (0 to n-1) with proper preallocation
func AppendOptimized(n int) []int {
	// TODO(human): Implement with preallocation
	return nil
}

// ConcatStringsNaive concatenates n copies of str using the + operator
func ConcatStringsNaive(str string, n int) string {
	// TODO(human): Implement using + operator
	return ""
}

// ConcatStringsOptimized concatenates n copies of str using strings.Builder
func ConcatStringsOptimized(str string, n int) string {
	// TODO(human): Implement using strings.Builder
	return ""
}

// FilterNaive filters a slice, keeping only even numbers (no capacity hint)
func FilterNaive(numbers []int) []int {
	// TODO(human): Implement without capacity hint
	return nil
}

// FilterOptimized filters a slice, keeping only even numbers (with capacity hint)
func FilterOptimized(numbers []int) []int {
	// TODO(human): Implement with capacity hint
	return nil
}
