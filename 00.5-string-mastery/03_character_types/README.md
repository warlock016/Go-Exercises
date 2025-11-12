# Exercise 03: Character Types

**Concept:** Identifying and classifying different types of runes
**Difficulty:** Medium
**Estimated Time:** 25 minutes

## 🎯 Learning Goal

Learn to identify different types of characters (letters, digits, spaces, etc.) using Go's `unicode` package. Understand that runes have properties beyond just their numeric value.

## The Problem

Not all runes are created equal! Some are letters, some are digits, some are spaces, and some are symbols. Go provides the `unicode` package to check these properties.

```go
'a' → letter
'5' → digit
' ' → space
'!' → punctuation
'世' → letter (Chinese character)
'👍' → symbol (emoji)
```

## Your Task

Implement functions that classify and filter runes based on their types:

### 1. IsLetter
Check if a rune is a letter (any alphabet, any language)

### 2. IsDigit
Check if a rune is a digit (0-9)

### 3. CountLetters
Count how many letters are in a string

### 4. CountDigits
Count how many digits are in a string

### 5. FilterLettersOnly
Return a new string containing only the letters from the input

## Function Signatures

```go
func IsLetter(r rune) bool
func IsDigit(r rune) bool
func CountLetters(s string) int
func CountDigits(s string) int
func FilterLettersOnly(s string) string
```

## Examples

```go
IsLetter('a')     // → true
IsLetter('5')     // → false
IsLetter('世')    // → true (Chinese character is a letter!)

IsDigit('5')      // → true
IsDigit('a')      // → false

CountLetters("Hello123")      // → 5
CountLetters("café")          // → 4
CountLetters("Hello, 世界!")  // → 7 (includes Chinese chars)

CountDigits("Hello123")       // → 3
CountDigits("Year: 2024")     // → 4

FilterLettersOnly("Hello123")     // → "Hello"
FilterLettersOnly("café!")        // → "café"
FilterLettersOnly("Hello, 世界!") // → "Hello世界"
```

## Instructions

1. Open `types.go`
2. Implement all five functions
3. Run `go test -v`
4. Append your learnings to `../EXPLANATION.md`

## Hints

### The unicode Package

```go
import "unicode"

// Check if rune is a letter
unicode.IsLetter('a')  // true
unicode.IsLetter('世') // true
unicode.IsLetter('5')  // false

// Check if rune is a digit
unicode.IsDigit('5')   // true
unicode.IsDigit('a')   // false

// Other useful functions:
unicode.IsSpace(' ')   // true
unicode.IsPunct('!')   // true
unicode.IsSymbol('$')  // true
```

### Counting Letters

```go
func CountLetters(s string) int {
    count := 0
    for _, r := range s {  // Iterate as runes!
        if unicode.IsLetter(r) {
            count++
        }
    }
    return count
}
```

### Filtering Letters

```go
import "strings"

func FilterLettersOnly(s string) string {
    var builder strings.Builder
    for _, r := range s {
        if unicode.IsLetter(r) {
            builder.WriteRune(r)  // Add this rune to result
        }
    }
    return builder.String()
}
```

## 🧠 Think About

1. Why does `unicode.IsLetter('世')` return true? What makes it a "letter"?
2. How would you check for uppercase vs lowercase?
3. What about numbers like '²' or '½'? Are they digits?
4. Why do we use `strings.Builder` for FilterLettersOnly instead of string concatenation?

## What This Teaches

- Runes have properties (letter, digit, space, etc.)
- The `unicode` package for character classification
- Works with ALL languages (not just English!)
- Efficient string building with `strings.Builder`
- Practical string filtering patterns

## Extra Challenge (Optional)

After completing the basic functions, try these:

```go
// Count letters, digits, and others separately
func CharacterStats(s string) (letters, digits, others int)

// Convert string to only alphanumeric (remove punctuation, spaces, etc.)
func ToAlphanumeric(s string) string

// Check if string contains only ASCII characters
func IsASCII(s string) bool
```

## After Completing

Update your `EXPLANATION.md` with:
- What makes a rune a "letter" vs other types
- Why `unicode` package is better than checking ranges like `'a' <= r <= 'z'`
- One example where character classification is useful in real programs
