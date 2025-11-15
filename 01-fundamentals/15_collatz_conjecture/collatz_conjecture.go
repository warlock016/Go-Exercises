package collatzconjecture

import "errors"

// CollatzSequence returns the complete Collatz sequence starting from n.
// The sequence follows these rules:
//   - If n is even: divide by 2
//   - If n is odd: multiply by 3 and add 1
//   - Continue until reaching 1
//
// Returns an error if n <= 0.
func CollatzSequence(n int) ([]int, error) {
	// TODO(human): Validate that n is positive
	// TODO(human): Create a slice to hold the sequence, starting with n
	// TODO(human): Loop while n is not 1:
	//   - If n is even (n % 2 == 0), divide by 2
	//   - If n is odd, compute 3*n + 1
	//   - Append the new value to the sequence
	// TODO(human): Return the sequence and nil error
	return nil, errors.New("not implemented")
}

// CollatzLength returns the length of the Collatz sequence for n.
// This is more efficient than CollatzSequence when you only need the count.
//
// Returns an error if n <= 0.
func CollatzLength(n int) (int, error) {
	// TODO(human): Validate that n is positive
	// TODO(human): Initialize a counter to 1 (for the starting number)
	// TODO(human): Loop while n is not 1:
	//   - Apply Collatz rules to update n
	//   - Increment counter
	// TODO(human): Return the counter and nil error
	return 0, errors.New("not implemented")
}

// MaxCollatzInRange finds the number in [start, end] with the longest Collatz sequence.
// Returns the number, the length of its sequence, and an error if inputs are invalid.
//
// If multiple numbers have the same maximum length, returns the smallest number.
// Returns an error if start > end or start <= 0.
func MaxCollatzInRange(start, end int) (num, length int, err error) {
	// TODO(human): Validate that start <= end and start > 0
	// TODO(human): Initialize variables to track max length and corresponding number
	// TODO(human): Loop from start to end:
	//   - Get the Collatz length for current number
	//   - If it's longer than current max, update max and num
	//   - (The loop naturally handles ties by keeping the first/smallest)
	// TODO(human): Return the number with max length, the length, and nil error
	return 0, 0, errors.New("not implemented")
}
