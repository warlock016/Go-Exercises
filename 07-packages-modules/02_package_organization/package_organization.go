package organization

import (
	"errors"

	"github.com/warlock016/go-exercises/07-packages-modules/02_package_organization/calculator"
)

// Calculate performs the specified operation on two integers
func Calculate(a, b int, operation string) (int, error) {
	// TODO(human): Implement using calculator sub-package
	var res int
	var err error = nil

	switch operation {
	case "add":
		res = calculator.Add(a, b)
	case "subtract":
		res = calculator.Subtract(a, b)
	case "multiply":
		res = calculator.Multiply(a, b)
	case "divide":
		res, err = calculator.Divide(a, b)
	default:
		err = errors.New("unknown operation")
	}

	return res, err
}
