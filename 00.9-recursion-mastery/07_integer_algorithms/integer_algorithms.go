package integer_algorithms

// GCD calculates the greatest common divisor using Euclid's algorithm
func GCD(a, b int) int {
	// TODO(human): Implement

	// var remainder int

	if a > b && b != 0 {
		return gcdHelper(a, b)
	} else if b > a && a != 0 {
		return gcdHelper(b, a)
	} else if a == b && a != 0 {
		return a
	} else {
		return 0
	}
}

func gcdHelper(a, b int) int {

	remainder := a % b
	quotient := (a - remainder) / b
	bnext := (a - remainder) / quotient

	if remainder == 0 {
		return bnext
	}

	return gcdHelper(bnext, remainder)
}

// SumDigits calculates the sum of all digits in n
func SumDigits(n int) int {
	// TODO(human): Implement

	if n == 0 {
		return 0
	}

	result := n%10 + SumDigits(n/10)
	// fmt.Printf("Input: %d, Result: %d\n", n, result)

	return result
}

// CountDigits counts the number of digits in n
func CountDigits(n int) int {
	// TODO(human): Implement

	if n < 10 {
		return 1
	}

	result := 1
	result += CountDigits(n / 10)
	return result
}

// IsPowerOfTwo checks if n is a power of 2
func IsPowerOfTwo(n int) bool {
	// TODO(human): Implement

	if n == 1 {
		return true
	}

	if n <= 0 || n%2 != 0 {
		return false
	}

	return IsPowerOfTwo(n / 2)
}
