# Exercise 05: Find Maximum

**Concept:** Slices, comparison, edge cases
**Difficulty:** Easy
**Estimated Time:** 5-7 minutes

## Problem

Write a function that finds the maximum value in a slice of integers.

## Function Signature

```go
func FindMax(numbers []int) (int, error)
```

## Examples

- `FindMax([]int{1, 5, 3, 9, 2})` → `9, nil`
- `FindMax([]int{-10, -5, -20})` → `-5, nil`
- `FindMax([]int{42})` → `42, nil`
- `FindMax([]int{})` → `0, error` (error: "empty slice")

## Instructions

1. Open `findmax.go`
2. Implement the `FindMax` function
3. Run `go test -v` to check your solution

## Hints

- Return an error for empty slices using `errors.New("message")`
- Import "errors" package
- Initialize max to the first element, then compare with the rest
- Don't forget to handle the empty slice case

## What This Tests

- Working with slices
- Comparison operations
- Error handling (returning errors)
- Edge case handling
