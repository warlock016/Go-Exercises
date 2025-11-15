package primechecker

// IsPrime returns true if n is a prime number, false otherwise.
// Numbers less than 2 are not prime.
func IsPrime(n int) bool {
	// TODO(human): Implement prime checking with square root optimization
	// 1. Handle edge cases: n < 2 returns false
	// 2. Handle special case: n == 2 returns true (only even prime)
	// 3. Check if n is even: if n % 2 == 0, return false
	// 4. Loop from i = 3 to √n (use i * i <= n), incrementing by 2 (odd numbers only)
	//    - If n % i == 0, return false (found a divisor)
	// 5. If no divisors found, return true
	return false
}

// FirstNPrimes returns a slice containing the first n prime numbers.
// If n <= 0, return an empty slice.
func FirstNPrimes(n int) []int {
	// TODO(human): Generate the first n prime numbers
	// 1. Handle edge case: if n <= 0, return []int{}
	// 2. Create a slice to store primes: primes := make([]int, 0, n)
	// 3. Start with candidate := 2
	// 4. Loop while len(primes) < n:
	//    - If IsPrime(candidate), append candidate to primes
	//    - Increment candidate
	// 5. Return primes
	return nil
}

// PrimesInRange returns all prime numbers in the range [start, end] inclusive.
// If start > end or no primes exist in the range, return an empty slice.
func PrimesInRange(start, end int) []int {
	// TODO(human): Find all primes in the given range
	// 1. Handle edge case: if start > end, return []int{}
	// 2. Adjust start: if start < 2, set start = 2 (primes must be >= 2)
	// 3. Create a slice to store primes: primes := []int{}
	// 4. Loop from num = start to num <= end:
	//    - If IsPrime(num), append num to primes
	// 5. Return primes
	return nil
}
