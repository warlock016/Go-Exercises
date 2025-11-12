# Exercise 09: Filter Even Numbers

**Concept:** Slices, filtering, append
**Difficulty:** Medium
**Estimated Time:** 5-8 minutes

## Problem

Write a function that filters a slice to keep only even numbers.

## Function Signature

```go
func FilterEven(numbers []int) []int
```

## Examples

- `FilterEven([]int{1, 2, 3, 4, 5, 6})` → `[]int{2, 4, 6}`
- `FilterEven([]int{1, 3, 5})` → `[]int{}` (empty slice, no evens)
- `FilterEven([]int{})` → `[]int{}` (empty input)
- `FilterEven([]int{2, 4, 6})` → `[]int{2, 4, 6}` (all even)

## Instructions

1. Open `filter.go`
2. Implement the `FilterEven` function
3. Run `go test -v` to check your solution

## Hints

- Create an empty result slice: `result := []int{}`
- Use modulo operator `%` to check if even: `num % 2 == 0`
- Append even numbers: `result = append(result, num)`
- Return the result slice

## What This Tests

- Slice creation and manipulation
- Append operation
- Filtering patterns
- Modulo operator
