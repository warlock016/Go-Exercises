# The Complete Guide to Strings, Runes, and Bytes in Go

**A deep dive into Go's string handling - created specifically for your learning path**

---

## 🎯 Why This Guide Exists

During your diagnostic assessment, you spent 45 minutes on the reverse string exercise - longer than any other exercise. You noted: *"extremely hard due to my lack of knowledge and experience with runes and iterations over strings."*

**This guide will fix that.** By the end, you'll understand exactly why string handling in Go works the way it does, and you'll never be confused again.

---

## 📚 Table of Contents

1. [The Three Types](#the-three-types)
2. [Why Go Does This](#why-go-does-this)
3. [Common Mistakes](#common-mistakes)
4. [How to Iterate Correctly](#how-to-iterate-correctly)
5. [Converting Between Types](#converting-between-types)
6. [String Building](#string-building)
7. [Practical Examples](#practical-examples)
8. [Quick Reference](#quick-reference)

---

## The Three Types

### 1. `string` - Immutable Byte Sequence

```go
s := "hello"  // string type
```

**Key Facts:**
- Immutable (cannot change individual characters)
- Stored internally as `[]byte`
- UTF-8 encoded
- `len(s)` returns **number of bytes**, NOT characters!

**Analogy:** Think of a string as a read-only book. You can read it, but you can't change the pages.

### 2. `rune` - A Single Unicode Character

```go
r := 'a'  // rune type (single quotes!)
```

**Key Facts:**
- Alias for `int32`
- Represents a Unicode code point
- One rune can be 1-4 bytes when encoded as UTF-8
- Use runes when working with individual characters

**Analogy:** A rune is like a single letter or symbol - 'a', 'é', '世', or '👍'.

### 3. `byte` - A Single Byte

```go
b := byte(65)  // byte type
```

**Key Facts:**
- Alias for `uint8`
- Represents a single byte (0-255)
- Use bytes for binary data or ASCII
- Multiple bytes combine to form UTF-8 encoded runes

**Analogy:** A byte is like a single ink dot. Multiple dots combine to form a letter.

---

## Why Go Does This

**Most languages hide the complexity of Unicode.**

- JavaScript: Strings behave like characters (but breaks on emoji!)
- Python 3: Strings are Unicode by default
- C: Strings are just byte arrays

**Go is explicit:** Strings are bytes. Characters are runes. This is intentional!

### The Problem Go Solves

```go
s := "café"

// In memory (UTF-8):
// Bytes: [99, 97, 102, 195, 169]
// Chars:  c   a   f    é (2 bytes!)

len(s)  // → 5 (bytes)
// NOT 4 characters!
```

**Why?**
- The character 'é' is encoded as TWO bytes in UTF-8 (195 and 169)
- The emoji '👍' is FOUR bytes (240, 159, 145, 141)

If `len()` pretended to count characters, Go would have to scan the entire string every time (slow!). Instead, Go is honest: "This string occupies 5 bytes of memory."

---

## Common Mistakes

### Mistake #1: Using `len()` to Count Characters

```go
❌ WRONG
s := "café"
numChars := len(s)  // 5, but there are only 4 characters!

✅ CORRECT
import "unicode/utf8"
numChars := utf8.RuneCountInString(s)  // 4
```

### Mistake #2: Indexing Directly

```go
❌ WRONG
s := "café"
lastChar := s[3]  // 195 (byte!), not 'é'

✅ CORRECT
runes := []rune(s)
lastChar := runes[3]  // 'é'
```

### Mistake #3: Byte Iteration Instead of Rune Iteration

```go
❌ WRONG - Breaks on Unicode
s := "café"
for i := 0; i < len(s); i++ {
    fmt.Printf("%c", s[i])  // c a f Ã © (BROKEN!)
}

✅ CORRECT - Handles Unicode
for _, r := range s {
    fmt.Printf("%c", r)  // c a f é (CORRECT!)
}
```

### Mistake #4: Concatenating Strings in a Loop

```go
❌ WRONG - Inefficient (creates new string each time)
result := ""
for _, word := range words {
    result += word  // Slow! Each += creates a new string
}

✅ CORRECT - Use strings.Builder
import "strings"

var builder strings.Builder
for _, word := range words {
    builder.WriteString(word)  // Fast! Amortized O(1)
}
result := builder.String()
```

---

## How to Iterate Correctly

### Option 1: Range (For Characters/Runes)

```go
s := "Hello, 世界"

for i, r := range s {
    fmt.Printf("At byte %d: character %c (rune value: %d)\n", i, r, r)
}

// Output:
// At byte 0: character H (rune value: 72)
// At byte 1: character e (rune value: 101)
// ...
// At byte 7: character 世 (rune value: 19990)
// At byte 10: character 界 (rune value: 30028)
```

**Note:** The index `i` is the BYTE position, not character position!

**Use when:** You need to process each character

### Option 2: Index (For Bytes)

```go
s := "hello"

for i := 0; i < len(s); i++ {
    fmt.Printf("Byte %d: %d\n", i, s[i])
}
```

**Use when:** You're working with ASCII or binary data

---

## Converting Between Types

### String → []rune

```go
s := "café"
runes := []rune(s)  // []rune{'c', 'a', 'f', 'é'}
// Now you can modify: runes[3] = 'é'
```

### String → []byte

```go
s := "hello"
bytes := []byte(s)  // []byte{104, 101, 108, 108, 111}
// Now you can modify: bytes[0] = 'H'
```

### []rune → String

```go
runes := []rune{'h', 'i', '!'}
s := string(runes)  // "hi!"
```

### []byte → String

```go
bytes := []byte{72, 105}
s := string(bytes)  // "Hi"
```

**Important:** These conversions CREATE COPIES. They're not free!

---

## String Building

### The Problem

```go
// DON'T DO THIS (slow for many concatenations)
result := ""
for i := 0; i < 1000; i++ {
    result += "word"  // Creates 1000 intermediate strings!
}
```

### The Solution: strings.Builder

```go
import "strings"

var builder strings.Builder

builder.WriteString("Hello")
builder.WriteString(" ")
builder.WriteString("World")
builder.WriteRune('!')

result := builder.String()  // "Hello World!"
```

**Why it's faster:**
- Grows internal buffer as needed
- Only creates final string once
- Amortized O(1) for each write

---

## Practical Examples

### Example 1: Count Vowels

```go
func CountVowels(s string) int {
    vowels := "aeiouAEIOU"
    count := 0

    for _, r := range s {  // Iterate as runes
        if strings.ContainsRune(vowels, r) {
            count++
        }
    }

    return count
}

CountVowels("café")  // → 2 (a, é)
```

### Example 2: Reverse a String (Correctly!)

```go
func Reverse(s string) string {
    runes := []rune(s)  // Convert to runes

    // Reverse the slice
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)
}

Reverse("café")  // → "éfac" (CORRECT!)
// NOT "éfac" with broken Unicode
```

### Example 3: Check if String Contains Only ASCII

```go
func IsASCII(s string) bool {
    for _, r := range s {
        if r > 127 {  // ASCII is 0-127
            return false
        }
    }
    return true
}

IsASCII("hello")  // → true
IsASCII("café")   // → false (é is beyond ASCII)
```

### Example 4: Truncate to N Characters

```go
func TruncateRunes(s string, n int) string {
    runes := []rune(s)
    if len(runes) <= n {
        return s
    }
    return string(runes[:n])
}

TruncateRunes("Hello, World!", 5)  // → "Hello"
TruncateRunes("café", 3)            // → "caf"
```

---

## Quick Reference

| Task | Code |
|------|------|
| Count bytes | `len(s)` |
| Count characters | `utf8.RuneCountInString(s)` |
| Iterate characters | `for _, r := range s { ... }` |
| Iterate bytes | `for i := 0; i < len(s); i++ { b := s[i] }` |
| Get Nth character | `runes := []rune(s); char := runes[n]` |
| Convert to runes | `runes := []rune(s)` |
| Convert to bytes | `bytes := []byte(s)` |
| Build string efficiently | `var b strings.Builder; b.WriteString(...)` |
| Reverse string | `runes := []rune(s); reverse slice; string(runes)` |

---

## Decision Tree

```
Need to work with a string?
│
├─ Need to count/iterate characters?
│  └─ Use: range iteration or []rune
│
├─ Need to modify individual characters?
│  └─ Use: Convert to []rune, modify, convert back
│
├─ Building a string in a loop?
│  └─ Use: strings.Builder
│
├─ Working with ASCII or binary data?
│  └─ Use: []byte or byte iteration
│
└─ Just reading/passing around text?
   └─ Use: string (keep it as is)
```

---

## Mental Model

Think of strings like this:

```
String "café" in memory:
┌─────────────────────────────┐
│ 99│ 97│102│195│169│         │  ← Bytes (what Go stores)
└─────────────────────────────┘
  c    a   f   é (2 bytes)       ← Runes (what humans see)
  0    1   2        3            ← Rune indices
```

- **Byte level:** 5 bytes, indexed 0-4
- **Rune level:** 4 characters, indexed 0-3
- **The 'é' character:** Starts at byte 3, spans bytes 3-4

---

## Key Takeaways

1. ✅ **Strings are byte sequences** (UTF-8 encoded)
2. ✅ **`len(s)` counts bytes**, not characters
3. ✅ **Use `range` to iterate characters** correctly
4. ✅ **Use `[]rune(s)` when you need to modify** individual characters
5. ✅ **Use `strings.Builder`** for efficient string construction
6. ✅ **Direct indexing `s[i]`** gives you a byte, not a character

---

## Practice Exercises

After reading this guide, you should:

1. ✅ Be able to explain why `len("café")` returns 5
2. ✅ Know when to use `range` vs index iteration
3. ✅ Understand why `s[3]` doesn't give you the 4th character
4. ✅ Be confident reversing a Unicode string
5. ✅ Know how to count vowels including accented ones

**Ready to practice?** Head to Module 00.5: String Mastery!

---

**Remember:** Go's explicit handling of bytes vs characters prevents bugs. Other languages hide this complexity, which leads to subtle Unicode bugs in production. Go's approach takes a bit more learning up front, but you'll write more correct code!
