package calculator

// Multiply returns the product of a and b
func Multiply(a, b int) int {
	// TODO(human): Implement
	return a * b
}

// Divide returns the quotient of a and b
func Divide(a, b int) (int, error) {
	// TODO(human): Implement using validate()
	if err := validate(a, b, "divide"); err != nil {
		return 0, err
	}
	return a / b, nil
}
