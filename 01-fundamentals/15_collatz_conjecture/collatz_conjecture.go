package collatzconjecture

import (
	"errors"
)

// CollatzSequence returns the complete Collatz sequence starting from n.
// The sequence follows these rules:
//   - If n is even: divide by 2
//   - If n is odd: multiply by 3 and add 1
//   - Continue until reaching 1
//
// Returns an error if n <= 0.
func CollatzSequence(n int) ([]int, error) {
	// TODO(human): Implement Collatz sequence generation
	seq := []int{}

	if n <= 0 {
		return nil, errors.New("invalid input integer")
	}

	for {
		seq = append(seq, n)
		if n == 1 {
			break
		} else {
			switch n%2 == 0 {
			case true:
				n /= 2
			case false:
				n = n*3 + 1
			}
		}

	}

	return seq, nil
}

// CollatzLength returns the length of the Collatz sequence for n.
// This is more efficient than CollatzSequence when you only need the count.
//
// Returns an error if n <= 0.
func CollatzLength(n int) (int, error) {
	// TODO(human): Implement Collatz sequence length calculation
	seq := []int{}

	if n <= 0 {
		return 0, errors.New("invalid input integer")
	}

	for {
		seq = append(seq, n)
		if n == 1 {
			break
		} else {
			switch n%2 == 0 {
			case true:
				n /= 2
			case false:
				n = n*3 + 1
			}
		}
	}

	return len(seq), nil
}

// MaxCollatzInRange finds the number in [start, end] with the longest Collatz sequence.
// Returns the number, the length of its sequence, and an error if inputs are invalid.
//
// If multiple numbers have the same maximum length, returns the smallest number.
// Returns an error if start > end or start <= 0.
func MaxCollatzInRange(start, end int) (num, length int, err error) {
	// TODO(human): Find number with longest Collatz sequence in range

	if start <= 0 || end < start {
		return 0, 0, errors.New("invalid input values")
	}

	maxNum := 0
	maxLen := 0

	for i := start; i <= end; i++ {
		currLen, _ := CollatzLength(i)

		if currLen > maxLen {
			maxLen = currLen
			maxNum = i
		}
	}

	return maxNum, maxLen, nil
}
