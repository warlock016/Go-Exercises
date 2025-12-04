package math

// Abs returns the absolute value of x
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Max returns the larger of a and b
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Clamp restricts value to be within min and max (inclusive)
func Clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
