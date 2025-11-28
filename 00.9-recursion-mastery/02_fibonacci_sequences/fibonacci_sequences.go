package fibonacci_sequences

// Fibonacci calculates the nth Fibonacci number
func Fibonacci(n int) int {
	// TODO(human): Implement
	// 0 1 1 2 3 5 8 13 21 34 55 89 144 233
	// 0 1 2 3 4 5 6  7  8  9 10 11  12  13

	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Tribonacci calculates the nth Tribonacci number
func Tribonacci(n int) int {
	// TODO(human): Implement

	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return Tribonacci(n-1) + Tribonacci(n-2) + Tribonacci(n-3)
}

// ClimbStairs calculates number of ways to climb n stairs (1 or 2 steps at a time)
func ClimbStairs(n int) int {
	// TODO(human): Implement

	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return ClimbStairs(n-1) + ClimbStairs(n-2)
}
