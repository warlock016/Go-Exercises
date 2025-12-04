package coverage

import "errors"

// Classify categorizes a number
func Classify(n int) string {
	if n < 0 {
		return "negative"
	}
	if n == 0 {
		return "zero"
	}
	if n < 10 {
		return "small positive"
	}
	if n < 100 {
		return "medium positive"
	}
	return "large positive"
}

// ProcessData processes a slice of integers
func ProcessData(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("empty data")
	}

	sum := 0
	for _, v := range data {
		if v < 0 {
			return 0, errors.New("negative value not allowed")
		}
		sum += v
	}

	if sum > 1000 {
		return 0, errors.New("sum too large")
	}

	return sum, nil
}

// Calculate performs complex calculation with many branches
func Calculate(a, b int, op string) (int, error) {
	switch op {
	case "add":
		return a + b, nil
	case "sub":
		return a - b, nil
	case "mul":
		if a > 1000 || b > 1000 {
			return 0, errors.New("values too large for multiplication")
		}
		return a * b, nil
	case "div":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("unknown operation")
	}
}
