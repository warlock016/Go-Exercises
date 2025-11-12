# Exercise 11: Update Value

**Concept:** Pointers, reference vs value
**Difficulty:** Medium
**Estimated Time:** 5-8 minutes

## Problem

Write a function that updates an integer value using a pointer.

## Function Signature

```go
func UpdateValue(n *int, newValue int)
```

## Example

```go
num := 10
UpdateValue(&num, 42)
// num is now 42
```

## Instructions

1. Open `pointer.go`
2. Implement the `UpdateValue` function
3. Run `go test -v` to check your solution

## Hints

- The `*int` parameter is a pointer to an integer
- Use `*n` to dereference and update the value
- The function doesn't return anything; it modifies the value at the pointer

## What This Tests

- Understanding of pointers
- Dereferencing pointers
- Modifying values through pointers
- Pass by reference vs pass by value
