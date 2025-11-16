package primechecker

// IsPrime returns true if n is a prime number, false otherwise.
// Numbers less than 2 are not prime.
func IsPrime(n int) bool {
	// TODO(human): Implement prime checking
	// Prime numbers are only divisible by 1 and themselves -> cannot be decomposed into factors!

	if n < 2 {
		return false
	}

	for i := 2; i <= n; i++ {
		if n%i == 0 && n != i {
			return false
		}
	}

	return true
}

// FirstNPrimes returns a slice containing the first n prime numbers.
// If n <= 0, return an empty slice.
func FirstNPrimes(n int) []int {
	// TODO(human): Generate the first n prime numbers

	result := []int{}

	if n <= 0 {
		return result
	}

	counter := 0

	for len(result) < n {
		if IsPrime(counter) {
			result = append(result, counter)
		}
		counter++
	}

	return result
}

// PrimesInRange returns all prime numbers in the range [start, end] inclusive.
// If start > end or no primes exist in the range, return an empty slice.
func PrimesInRange(start, end int) []int {
	// TODO(human): Find all primes in the given range

	result := []int{}

	if start > end {
		return result
	}

	for i := start; i <= end; i++ {
		if IsPrime(i) {
			result = append(result, i)
		}
	}

	return result
}
