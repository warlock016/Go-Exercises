package divide

import "fmt"

// Divide returns a/b or an error if b is zero
func Divide(a, b float64) (float64, error) {
	// TODO(human): Implement safe division

	if b == 0 {
		return 0, fmt.Errorf("invalid division by zero")
	}

	return (a / b), nil
}
