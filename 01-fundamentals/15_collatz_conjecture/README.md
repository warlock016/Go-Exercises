# Exercise 15: Collatz Conjecture

## Learning Goal

Master loop construction and sequence generation by implementing the famous Collatz Conjecture (3n+1 problem), practicing slice building, iteration control, and range operations.

## Problem Description

The Collatz Conjecture is one of mathematics' most famous unsolved problems. It states that for any positive integer, if you repeatedly apply these rules, you'll always eventually reach 1:

1. If the number is even, divide it by 2
2. If the number is odd, multiply by 3 and add 1
3. Repeat until you reach 1

For example, starting with 3:
- 3 (odd) → 3×3+1 = 10
- 10 (even) → 10÷2 = 5
- 5 (odd) → 5×3+1 = 16
- 16 (even) → 16÷2 = 8
- 8 (even) → 8÷2 = 4
- 4 (even) → 4÷2 = 2
- 2 (even) → 2÷2 = 1

The sequence is: [3, 10, 5, 16, 8, 4, 2, 1]

Despite extensive testing, no one has proven this always works, but no counterexample has ever been found!

## Function Signatures

```go
func CollatzSequence(n int) ([]int, error)
func CollatzLength(n int) (int, error)
func MaxCollatzInRange(start, end int) (num, length int, err error)
```

## Examples

### CollatzSequence
```go
CollatzSequence(1)  → [1], nil
CollatzSequence(3)  → [3, 10, 5, 16, 8, 4, 2, 1], nil
CollatzSequence(6)  → [6, 3, 10, 5, 16, 8, 4, 2, 1], nil
CollatzSequence(0)  → nil, error
CollatzSequence(-5) → nil, error
```

### CollatzLength
```go
CollatzLength(1)  → 1, nil
CollatzLength(3)  → 8, nil
CollatzLength(6)  → 9, nil
CollatzLength(27) → 112, nil
```

### MaxCollatzInRange
```go
MaxCollatzInRange(1, 10)   → num: 9, length: 20, nil
MaxCollatzInRange(1, 1)    → num: 1, length: 1, nil
MaxCollatzInRange(10, 5)   → 0, 0, error
MaxCollatzInRange(-5, 10)  → 0, 0, error
```

## Instructions

1. Implement `CollatzSequence`
2. Implement `CollatzLength`
3. Implement `MaxCollatzInRange`
4. Run tests with `go test -v`

## Think About

1. Why is it more efficient to count steps rather than build the full sequence?
2. What's the time complexity of finding the max in a range?
3. Could you optimize MaxCollatzInRange by caching previously computed sequences?
4. Why does the conjecture require n to be positive?
5. What happens to very large numbers - do sequences get longer or shorter on average?

## What This Teaches

- **Loop control**: Using conditions to control iteration (while n != 1)
- **Sequence building**: Growing slices dynamically with append
- **Conditional logic**: Branching based on number properties (even/odd)
- **Error handling**: Validating input constraints
- **Algorithm optimization**: Computing counts vs storing full sequences
- **Range operations**: Iterating through number ranges efficiently
- **Multiple return values**: Functions returning both data and errors
- **Integer arithmetic**: Division, multiplication, modulo operations

## Edge Cases to Consider

- n = 1 (sequence is just [1])
- Very small ranges in MaxCollatzInRange
- Ranges where all numbers have same length sequence
- Large numbers that produce long sequences
- Invalid inputs (zero, negative, inverted ranges)
