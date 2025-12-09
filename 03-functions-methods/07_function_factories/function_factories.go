package function_factories

import (
	"math"
	"strings"
)

// MakeAdder returns a function that adds x to its argument
func MakeAdder(x int) func(int) int {
	// TODO(human): Implement
	return func(i int) int {
		return x + i

	}
}

// MakeGreeter returns a function that greets with the given greeting
func MakeGreeter(greeting string) func(string) string {
	// TODO(human): Implement

	return func(s string) string {
		var result strings.Builder

		if len(s) == 0 {
			return result.String()
		}

		result.WriteString(greeting + ", ")
		result.WriteString(s + "!")

		return result.String()
	}
}

// MakeValidator returns a function that validates if a number is in range
func MakeValidator(min, max int) func(int) bool {
	// TODO(human): Implement
	return func(i int) bool {
		if i < min || i > max {
			return false
		} else {
			return true
		}
	}
}

// MakeFormatter returns a function that wraps strings with prefix and suffix
func MakeFormatter(prefix, suffix string) func(string) string {
	// TODO(human): Implement
	return func(s string) string {
		var result strings.Builder
		result.WriteString(prefix + s + suffix)
		return result.String()
	}
}

// MakePowerFunction returns a function that raises numbers to the given exponent
func MakePowerFunction(exponent int) func(float64) float64 {
	// TODO(human): Implement
	return func(f float64) float64 {
		return math.Pow(f, float64(exponent))
	}
}
