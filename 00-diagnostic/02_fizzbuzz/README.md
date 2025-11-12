# Exercise 02: FizzBuzz

**Concept:** Control structures, conditionals, string operations
**Difficulty:** Easy
**Estimated Time:** 5-8 minutes

## Problem

Implement the classic FizzBuzz function that returns a string for numbers 1 to n:
- Return "Fizz" for multiples of 3
- Return "Buzz" for multiples of 5
- Return "FizzBuzz" for multiples of both 3 and 5
- Return the number as a string otherwise

## Function Signature

```go
func FizzBuzz(n int) []string
```

## Example

`FizzBuzz(15)` returns:
```
["1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"]
```

## Instructions

1. Open `fizzbuzz.go`
2. Implement the `FizzBuzz` function
3. Run `go test -v` to check your solution

## Hints

- Use the modulo operator `%` to check for multiples
- Check for FizzBuzz (divisible by both) FIRST
- Use `strconv.Itoa()` to convert int to string
- Remember to import "strconv" if you use it

## What This Tests

- For loops
- Conditional logic (if/else)
- Working with slices
- String conversion
