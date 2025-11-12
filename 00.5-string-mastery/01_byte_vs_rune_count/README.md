# Exercise 01: Byte vs Rune Count

**Concept:** Understanding the fundamental difference between bytes and runes
**Difficulty:** Easy
**Estimated Time:** 15 minutes

## 🎯 Learning Goal

Understand that `len(s)` returns **bytes**, not characters. Learn to count **runes** (characters) properly.

## The Problem

In Go, strings are sequences of bytes (UTF-8 encoded). A single character might be multiple bytes!

Examples:
- `"hello"` - 5 bytes, 5 runes (ASCII characters are 1 byte each)
- `"café"` - 5 bytes, 4 runes (`é` is 2 bytes!)
- `"👍"` - 4 bytes, 1 rune (emoji are often 4 bytes!)

## Your Task

Implement two functions:

### 1. ByteCount
Count the number of **bytes** in a string

### 2. RuneCount
Count the number of **runes** (characters) in a string

## Function Signatures

```go
func ByteCount(s string) int
func RuneCount(s string) int
```

## Examples

```go
ByteCount("hello")  // → 5 (5 ASCII chars = 5 bytes)
RuneCount("hello")  // → 5 (5 characters)

ByteCount("café")   // → 5 (c=1, a=1, f=1, é=2 bytes)
RuneCount("café")   // → 4 (4 characters)

ByteCount("👍")     // → 4 (emoji are multi-byte)
RuneCount("👍")     // → 1 (1 character)

ByteCount("Hello, 世界")  // → 13 bytes
RuneCount("Hello, 世界")  // → 9 runes
```

## Instructions

1. Open `count.go`
2. Implement both functions
3. Run `go test -v`
4. Create `EXPLANATION.md` explaining why the counts differ

## Hints

### Byte Count
```go
// Built-in len() returns byte count
byteCount := len(s)
```

### Rune Count
```go
import "unicode/utf8"

// Use utf8.RuneCountInString() to count runes
runeCount := utf8.RuneCountInString(s)

// OR iterate with range (range over string gives runes)
count := 0
for range s {
    count++
}
```

## 🧠 Think About

1. Why does `len("café")` return 5 instead of 4?
2. What happens if you try to get `s[2]` in "café"? Do you get 'f' or 'é'?
3. Why does Go make this distinction?

## What This Teaches

- Strings are byte sequences, not character sequences
- `len()` returns **bytes**, not **characters**
- How to properly count characters using `utf8.RuneCountInString()`
- Why Unicode handling is important

## After Completing

Write in your `EXPLANATION.md`:
- The difference between a byte and a rune
- When each count is useful
- An example where they differ and why
