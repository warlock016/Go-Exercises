package calc

import "fmt"

// Add returns the sum of two integers
func Add(a, b int) int {
	// TODO(human): Implement addition
	return a + b
}

// Subtract returns the difference (a - b)
func Subtract(a, b int) int {
	// TODO(human): Implement subtraction
	return a - b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	// TODO(human): Implement multiplication
	return a * b
}

// Divide returns the quotient and remainder of a/b
// Returns an error if b is zero
func Divide(a, b int) (quotient, remainder int, err error) {
	// TODO(human): Implement division with error handling for division by zero

	if b == 0 {
		return 0, 0, fmt.Errorf("invalid division by zero")
	}
	quotient = a / b
	remainder = a % b
	return quotient, remainder, nil
}

// Modulo returns the remainder of a/b
// Returns an error if b is zero
func Modulo(a, b int) (int, error) {
	// TODO(human): Implement modulo with error handling for division by zero
	if b == 0 {
		return 0, fmt.Errorf("invalid division by zero")
	}

	return a % b, nil
}
