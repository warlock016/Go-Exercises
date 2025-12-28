package calculator

import "errors"

// validate checks if the operation is valid
func validate(a, b int, op string) error {
	// TODO(human): Implement division by zero check

	if op == "divide" && b == 0 {
		return errors.New("invalid division by zero")
	}

	return nil
}
