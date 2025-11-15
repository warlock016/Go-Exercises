# Prime Checker

## Learning Goal
Master algorithmic thinking with loops by implementing prime number detection with optimization techniques (square root optimization).

## Problem Description

Prime numbers are natural numbers greater than 1 that have no positive divisors other than 1 and themselves. For example, 2, 3, 5, 7, 11, and 13 are prime numbers. The number 1 is not considered prime, and 2 is the only even prime number.

In this exercise, you'll implement three functions that work with prime numbers:
1. Check if a single number is prime
2. Generate the first N prime numbers
3. Find all prime numbers in a given range

The key optimization you'll learn is: **to check if a number n is prime, you only need to test divisors up to √n**. This is because if n has a divisor greater than √n, it must also have a corresponding divisor less than √n.

## Function Signatures

```go
// IsPrime returns true if n is a prime number, false otherwise.
// Numbers less than 2 are not prime.
func IsPrime(n int) bool

// FirstNPrimes returns a slice containing the first n prime numbers.
// If n <= 0, return an empty slice.
func FirstNPrimes(n int) []int

// PrimesInRange returns all prime numbers in the range [start, end] inclusive.
// If start > end or no primes exist in the range, return an empty slice.
func PrimesInRange(start, end int) []int
```

## Examples

### IsPrime
```go
IsPrime(2)   // true  (2 is prime)
IsPrime(3)   // true  (3 is prime)
IsPrime(4)   // false (4 = 2 × 2)
IsPrime(17)  // true  (17 is prime)
IsPrime(1)   // false (1 is not prime by definition)
IsPrime(0)   // false (0 is not prime)
IsPrime(-5)  // false (negative numbers are not prime)
```

### FirstNPrimes
```go
FirstNPrimes(0)  // []
FirstNPrimes(1)  // [2]
FirstNPrimes(5)  // [2, 3, 5, 7, 11]
FirstNPrimes(10) // [2, 3, 5, 7, 11, 13, 17, 19, 23, 29]
```

### PrimesInRange
```go
PrimesInRange(1, 10)   // [2, 3, 5, 7]
PrimesInRange(10, 20)  // [11, 13, 17, 19]
PrimesInRange(24, 30)  // [29]
PrimesInRange(14, 16)  // []       (no primes in this range)
PrimesInRange(5, 2)    // []       (start > end)
```

## Instructions

1. Implement `IsPrime`
2. Implement `FirstNPrimes`
3. Implement `PrimesInRange`
4. Run tests with `go test -v`

## Think About

1. **Why is checking up to √n sufficient?** Can you explain with a concrete example like n = 36?

2. **What's the performance difference** between checking all numbers up to n versus checking up to √n? For n = 10,000, how many iterations does each approach require?

3. **Why do we skip even numbers** (except 2) when checking for divisors? How does this optimization cut the work in half?

4. **Could you make FirstNPrimes more efficient?** Instead of checking every number, could you skip even numbers after 2?

## What This Teaches

- **Algorithmic optimization:** Understanding why √n is sufficient for primality testing
- **Loop patterns:** Using for loops with conditions that depend on calculations (i * i <= n)
- **Edge case handling:** Dealing with special cases (2 is the only even prime, numbers < 2, empty ranges)
- **Working with slices:** Building result slices incrementally with append
- **Function composition:** Using IsPrime as a building block for more complex functions
- **Performance awareness:** Recognizing that small optimizations (checking only odd numbers, √n limit) have significant impact
