# Exercise 06: Reverse String

**Concept:** Strings, runes, slice manipulation
**Difficulty:** Medium
**Estimated Time:** 7-10 minutes

## Problem

Write a function that reverses a string, properly handling Unicode characters.

## Function Signature

```go
func Reverse(s string) string
```

## Examples

- `Reverse("hello")` → `"olleh"`
- `Reverse("Go!")` → `"!oG"`
- `Reverse("café")` → `"éfac"` (handles accented characters)
- `Reverse("")` → `""`

## Instructions

1. Open `reverse.go`
2. Implement the `Reverse` function
3. Run `go test -v` to check your solution

## Hints

- Convert string to []rune to handle Unicode properly: `runes := []rune(s)`
- Reverse the slice by swapping elements
- Convert back to string: `string(runes)`
- Remember: strings are immutable, but slices are not

## What This Tests

- Understanding runes vs bytes
- Slice manipulation
- String/rune conversion
- Handling Unicode correctly
