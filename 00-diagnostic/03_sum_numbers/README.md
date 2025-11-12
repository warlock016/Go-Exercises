# Exercise 03: Sum Numbers

**Concept:** Slices, loops, accumulation
**Difficulty:** Easy
**Estimated Time:** 3-5 minutes

## Problem

Write a function that takes a slice of integers and returns their sum.

## Function Signature

```go
func Sum(numbers []int) int
```

## Examples

- `Sum([]int{1, 2, 3, 4, 5})` → `15`
- `Sum([]int{10, 20})` → `30`
- `Sum([]int{})` → `0` (empty slice)
- `Sum([]int{-5, 5})` → `0`

## Instructions

1. Open `sum.go`
2. Implement the `Sum` function
3. Run `go test -v` to check your solution

## Hints

- Use a for loop with `range` to iterate over the slice
- Initialize a variable to accumulate the sum
- Empty slices should return 0

## What This Tests

- Working with slices
- For loop with range
- Accumulator pattern
- Handling edge cases (empty slice)
