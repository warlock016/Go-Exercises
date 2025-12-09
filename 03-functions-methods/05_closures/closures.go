package closures

import (
	"fmt"
	"time"
)

// Counter returns a function that increments and returns a count
func Counter() func() int {
	// TODO(human): Implement
	count := 0
	return func() int {
		count++
		return count
	}
}

// Accumulator returns a function that adds to a running sum
func Accumulator() func(int) int {
	// TODO(human): Implement
	count := 0

	return func(i int) int {
		count += i
		return count
	}
}

// Multiplier returns a function that multiplies by the given factor
func Multiplier(factor int) func(int) int {
	// TODO(human): Implement
	return func(i int) int {
		return factor * i
	}

}

// RateLimiter returns a function that allows maxCalls within the time window
func RateLimiter(maxCalls int, window time.Duration) func() bool {
	// TODO(human): Implement

	var calls []time.Time

	return func() bool {
		now := time.Now()
		cutoff := now.Add(-window)
		validCalls := []time.Time{}

		fmt.Printf("now: %v, cutoff: %v, time slice: %v\n", now, cutoff, validCalls)

		for _, t := range calls {
			if t.After(cutoff) {
				validCalls = append(validCalls, t)
			}
		}
		calls = validCalls

		if len(calls) < maxCalls {
			calls = append(calls, now)
			return true
		}
		return false
	}
}
