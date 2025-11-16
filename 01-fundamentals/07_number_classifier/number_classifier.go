package numberclassifier

// ClassifyParity returns "even" if n is divisible by 2, otherwise "odd"
func ClassifyParity(n int) string {
	// TODO(human): Implement ClassifyParity
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}

// ClassifySign returns "positive", "negative", or "zero" based on the number's sign
func ClassifySign(n int) string {
	// TODO(human): Implement ClassifySign

	if n > 0 {
		return "positive"
	} else if n < 0 {
		return "negative"
	}

	return "zero"
}

// ClassifyMagnitude classifies numbers into size ranges
// Small: 0-10, Medium: 11-100, Large: >100
// Uses absolute value for negative numbers
func ClassifyMagnitude(n int) string {
	// TODO(human): Implement ClassifyMagnitude

	if n < 0 {
		n *= -1
	}

	if n <= 10 {
		return "small"
	} else if n <= 100 {
		return "medium"
	}

	return "large"
}

// DayName converts day numbers (1-7) to day names
// 1=Monday, 2=Tuesday, ..., 7=Sunday
// Returns "Invalid day" for numbers outside 1-7
func DayName(n int) string {
	// TODO(human): Implement DayName
	switch n {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}
