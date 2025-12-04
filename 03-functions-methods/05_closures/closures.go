package closures

import "time"

// Counter returns a function that increments and returns a count
func Counter() func() int {
	// TODO(human): Implement
	return nil
}

// Accumulator returns a function that adds to a running sum
func Accumulator() func(int) int {
	// TODO(human): Implement
	return nil
}

// Multiplier returns a function that multiplies by the given factor
func Multiplier(factor int) func(int) int {
	// TODO(human): Implement
	return nil
}

// RateLimiter returns a function that allows maxCalls within the time window
func RateLimiter(maxCalls int, window time.Duration) func() bool {
	// TODO(human): Implement
	return nil
}
