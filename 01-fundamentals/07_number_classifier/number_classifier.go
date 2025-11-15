package numberclassifier

// ClassifyParity returns "even" if n is divisible by 2, otherwise "odd"
func ClassifyParity(n int) string {
	// TODO(human): Use a switch statement on n % 2
	// Case 0 means even, default means odd
	return ""
}

// ClassifySign returns "positive", "negative", or "zero" based on the number's sign
func ClassifySign(n int) string {
	// TODO(human): Use a tagless switch with boolean conditions
	// switch {
	// case n > 0:
	//     return "positive"
	// case n < 0:
	//     return "negative"
	// default:
	//     return "zero"
	// }
	return ""
}

// ClassifyMagnitude classifies numbers into size ranges
// Small: 0-10, Medium: 11-100, Large: >100
// Uses absolute value for negative numbers
func ClassifyMagnitude(n int) string {
	// TODO(human): Get absolute value first
	// absN := n
	// if n < 0 {
	//     absN = -n
	// }
	//
	// Then switch on boolean conditions:
	// case absN <= 10: return "small"
	// case absN <= 100: return "medium"
	// default: return "large"
	return ""
}

// DayName converts day numbers (1-7) to day names
// 1=Monday, 2=Tuesday, ..., 7=Sunday
// Returns "Invalid day" for numbers outside 1-7
func DayName(n int) string {
	// TODO(human): Switch on n with cases 1-7
	// case 1: return "Monday"
	// case 2: return "Tuesday"
	// ...
	// default: return "Invalid day"
	return ""
}
