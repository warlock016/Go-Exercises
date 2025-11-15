package looppatterns

// CountUp returns a slice containing [1, 2, 3, ..., n]
// Use traditional for loop: for i := start; i <= n; i++
// Edge case: if n <= 0, return empty slice
func CountUp(n int) []int {
	// TODO(human): Implement using traditional for loop (init; condition; post)
	// 1. Handle edge case: if n <= 0, return []int{}
	// 2. Create empty slice: result := []int{}
	// 3. Loop from i := 1; i <= n; i++
	// 4. Append i to result in each iteration
	// 5. Return result
	return nil
}

// CountDown returns a slice containing [n, n-1, n-2, ..., 1]
// Use while-style for loop: for condition { ... }
// Edge case: if n <= 0, return empty slice
func CountDown(n int) []int {
	// TODO(human): Implement using while-style for loop (only condition)
	// 1. Handle edge case: if n <= 0, return []int{}
	// 2. Create empty slice: result := []int{}
	// 3. Initialize counter: current := n
	// 4. Loop while current > 0
	// 5. Inside loop: append current, then decrement current--
	// 6. Return result
	return nil
}

// SumSlice returns the sum of all numbers in the slice
// Use range-based for loop: for _, v := range nums
// Edge case: empty slice returns 0
func SumSlice(nums []int) int {
	// TODO(human): Implement using range-based for loop
	// 1. Initialize sum := 0
	// 2. Use for _, num := range nums to iterate
	// 3. Add each num to sum
	// 4. Return sum
	// Note: Empty slice naturally returns 0 (sum starts at 0, loop never runs)
	return 0
}

// FirstNEvens returns the first n even positive numbers [2, 4, 6, ...]
// Use infinite for loop with break: for { ... if count == n { break } }
// Edge case: if n <= 0, return empty slice
func FirstNEvens(n int) []int {
	// TODO(human): Implement using infinite for loop + break
	// 1. Handle edge case: if n <= 0, return []int{}
	// 2. Create empty slice: result := []int{}
	// 3. Initialize current := 2 (first even number)
	// 4. Start infinite loop: for {
	// 5. Append current to result
	// 6. Check if len(result) == n, if so: break
	// 7. Increment current by 2 (next even)
	// 8. Return result after loop
	return nil
}
