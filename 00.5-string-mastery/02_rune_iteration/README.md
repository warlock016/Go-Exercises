# Exercise 02: Rune Iteration

**Concept:** Correct ways to iterate over strings
**Difficulty:** Easy
**Estimated Time:** 20 minutes

## 🎯 Learning Goal

Learn the TWO ways to iterate over strings and when to use each.

## The Problem

There are two ways to iterate over a string in Go:

### Method 1: Index Iteration (Bytes)
```go
for i := 0; i < len(s); i++ {
    b := s[i]  // b is a byte
}
```
❌ **Problem:** Breaks on multi-byte characters like "café" or "👍"

### Method 2: Range Iteration (Runes)
```go
for i, r := range s {
    // r is a rune (character)
    // i is the byte index where this rune starts
}
```
✅ **Correct:** Handles Unicode properly

## Your Task

Implement three functions that demonstrate proper string iteration:

### 1. CollectBytes
Return a slice of bytes from the string

### 2. CollectRunes
Return a slice of runes from the string

### 3. GetRuneAt
Get the Nth rune (character) from a string (0-indexed)
Return an error if index is out of bounds

## Function Signatures

```go
func CollectBytes(s string) []byte
func CollectRunes(s string) []rune
func GetRuneAt(s string, index int) (rune, error)
```

## Examples

```go
CollectBytes("hi")    // → []byte{104, 105}
CollectRunes("hi")    // → []rune{'h', 'i'}

CollectBytes("café")  // → []byte{99, 97, 102, 195, 169} (5 bytes)
CollectRunes("café")  // → []rune{'c', 'a', 'f', 'é'} (4 runes)

GetRuneAt("café", 3)  // → 'é', nil
GetRuneAt("café", 4)  // → 0, error (out of bounds)
GetRuneAt("👍🎉", 1)  // → '🎉', nil
```

## Instructions

1. Open `iterate.go`
2. Implement all three functions
3. Run `go test -v`
4. Create `EXPLANATION.md`

## Hints

### CollectBytes
```go
// Method 1: Direct conversion
return []byte(s)

// Method 2: Manual iteration
bytes := make([]byte, 0, len(s))
for i := 0; i < len(s); i++ {
    bytes = append(bytes, s[i])
}
return bytes
```

### CollectRunes
```go
// Method 1: Direct conversion
return []rune(s)

// Method 2: Using range
runes := make([]rune, 0)
for _, r := range s {
    runes = append(runes, r)
}
return runes
```

### GetRuneAt
```go
// Convert to []rune and check bounds
runes := []rune(s)
if index < 0 || index >= len(runes) {
    return 0, errors.New("index out of bounds")
}
return runes[index], nil

// OR iterate with range tracking position
currentIndex := 0
for _, r := range s {
    if currentIndex == index {
        return r, nil
    }
    currentIndex++
}
return 0, errors.New("index out of bounds")
```

## 🧠 Think About

1. Why does `s[3]` NOT give you the 4th character in "café"?
2. What does the index in `for i, r := range s` represent?
3. When would you want byte iteration vs rune iteration?

## What This Teaches

- Two ways to iterate: index (bytes) vs range (runes)
- Why `range` is usually the correct choice
- How to safely access the Nth character
- Converting between strings, []byte, and []rune

## After Completing

Write in your `EXPLANATION.md`:
- When to use index iteration vs range iteration
- What the index in `range` represents (hint: byte position!)
- Why accessing `s[n]` directly is dangerous for Unicode
