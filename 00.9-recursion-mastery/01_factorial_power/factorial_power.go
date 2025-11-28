package factorial_power

// Factorial calculates n! = n × (n-1) × ... × 1
func Factorial(n int) int {
	// TODO(human): Implement

	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}

// 3: -> 3*Factorial(2) -> 3*2*Factorial(1) -> 3*2*1 = 6

// Power calculates base^exponent using recursion
func Power(base, exponent int) int {
	// TODO(human): Implement

	if exponent < 0 { // base^-1 == 1 / (base^1)
		return base * Power(base, exponent-1)
	}

	if exponent == 0 { // base^0 = 1
		return 1
	}

	if exponent == 1 { // base^1 = base
		return base
	}

	return base * Power(base, exponent-1)
}
