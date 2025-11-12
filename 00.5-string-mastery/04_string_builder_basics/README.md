# Exercise 04: String Builder Basics

**Concept:** Efficient string construction with strings.Builder
**Difficulty:** Medium
**Estimated Time:** 20 minutes

## 🎯 Learning Goal

Learn why string concatenation in loops is slow and how to build strings efficiently using `strings.Builder`.

## The Problem

Strings in Go are **immutable**. Every time you concatenate with `+=`, Go creates a NEW string:

```go
// SLOW - Creates N intermediate strings!
result := ""
for i := 0; i < 1000; i++ {
    result += "word"  // Creates 1000 new strings!
}
// O(N²) time complexity!
```

```go
// FAST - Builds once at the end
var builder strings.Builder
for i := 0; i < 1000; i++ {
    builder.WriteString("word")  // Appends to buffer
}
result := builder.String()  // Creates string once
// O(N) time complexity!
```

**Performance difference:** For 10,000 concatenations:
- String concatenation: ~500ms
- strings.Builder: ~1ms

**500x faster!**

## Your Task

Implement functions using `strings.Builder` for efficient string building:

### 1. JoinWords
Join a slice of strings with a separator (like strings.Join, but you implement it)

### 2. Repeat
Repeat a string N times (like strings.Repeat, but you implement it)

### 3. BuildList
Build a numbered list from a slice of items:
```
1. Apple
2. Banana
3. Cherry
```

### 4. BuildCSV
Convert a 2D slice into CSV format (comma-separated values)

## Function Signatures

```go
func JoinWords(words []string, separator string) string
func Repeat(s string, n int) string
func BuildList(items []string) string
func BuildCSV(rows [][]string) string
```

## Examples

```go
JoinWords([]string{"Hello", "World"}, " ")  // → "Hello World"
JoinWords([]string{"a", "b", "c"}, "-")     // → "a-b-c"
JoinWords([]string{"one"}, ",")             // → "one"
JoinWords([]string{}, ",")                  // → ""

Repeat("Go", 3)      // → "GoGoGo"
Repeat("Hi!", 2)     // → "Hi!Hi!"
Repeat("test", 0)    // → ""

items := []string{"Apple", "Banana", "Cherry"}
BuildList(items)
// →
// 1. Apple
// 2. Banana
// 3. Cherry

rows := [][]string{
    {"Name", "Age", "City"},
    {"Alice", "25", "NYC"},
    {"Bob", "30", "LA"},
}
BuildCSV(rows)
// →
// Name,Age,City
// Alice,25,NYC
// Bob,30,LA
```

## Instructions

1. Open `builder.go`
2. Implement all four functions using `strings.Builder`
3. Run `go test -v`
4. Append learnings to `../EXPLANATION.md`

## Hints

### Basic strings.Builder Usage

```go
import "strings"

var builder strings.Builder

// Add strings
builder.WriteString("Hello")
builder.WriteString(" ")
builder.WriteString("World")

// Add individual runes
builder.WriteRune('!')

// Get the final string
result := builder.String()  // "Hello World!"
```

### JoinWords Pattern

```go
func JoinWords(words []string, separator string) string {
    if len(words) == 0 {
        return ""
    }

    var builder strings.Builder

    // Add first word
    builder.WriteString(words[0])

    // Add remaining words with separator
    for i := 1; i < len(words); i++ {
        builder.WriteString(separator)
        builder.WriteString(words[i])
    }

    return builder.String()
}
```

### Repeat Pattern

```go
func Repeat(s string, n int) string {
    var builder strings.Builder

    for i := 0; i < n; i++ {
        builder.WriteString(s)
    }

    return builder.String()
}
```

### BuildList Pattern

```go
import (
    "strings"
    "strconv"
)

func BuildList(items []string) string {
    var builder strings.Builder

    for i, item := range items {
        // Write number
        builder.WriteString(strconv.Itoa(i + 1))
        builder.WriteString(". ")

        // Write item
        builder.WriteString(item)

        // Add newline (except after last item)
        if i < len(items)-1 {
            builder.WriteRune('\n')
        }
    }

    return builder.String()
}
```

## 🧠 Think About

1. Why is `result += str` slow in a loop?
2. How does `strings.Builder` avoid creating intermediate strings?
3. When should you use `WriteString` vs `WriteRune`?
4. Is there ever a time when `+=` is acceptable?

## What This Teaches

- Why strings are immutable in Go
- Performance implications of string concatenation
- How to use `strings.Builder` efficiently
- Common string building patterns
- When optimization matters

## Performance Insight

```go
// For small, fixed concatenations, += is fine:
s := "Hello" + " " + "World"  // OK - only 3 parts

// For loops, use Builder:
var builder strings.Builder
for _, word := range manyWords {
    builder.WriteString(word)  // Much faster
}
```

**Rule of thumb:** More than 3-4 concatenations? Use `strings.Builder`.

## After Completing

Update your `EXPLANATION.md` with:
- Why string immutability affects performance
- When to use `strings.Builder` vs simple concatenation
- How `strings.Builder` achieves better performance
- One real-world example where this optimization matters
