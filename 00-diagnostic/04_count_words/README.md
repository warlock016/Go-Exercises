# Exercise 04: Count Words

**Concept:** Maps, string splitting, iteration
**Difficulty:** Easy-Medium
**Estimated Time:** 5-10 minutes

## Problem

Write a function that counts how many times each word appears in a string.

## Function Signature

```go
func CountWords(text string) map[string]int
```

## Examples

`CountWords("hello world hello")` → `map[string]int{"hello": 2, "world": 1}`

`CountWords("go go go")` → `map[string]int{"go": 3}`

## Instructions

1. Open `wordcount.go`
2. Implement the `CountWords` function
3. Run `go test -v` to check your solution

## Hints

- Use `strings.Fields(text)` to split text into words
- Initialize a map: `m := make(map[string]int)`
- Increment the count for each word
- Don't forget to import "strings"

## What This Tests

- Creating and using maps
- String splitting
- Map initialization and updates
- Iterating over slices
