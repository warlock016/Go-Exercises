# Exercise 05: Reverse String Redux

**Concept:** Applying rune knowledge to solve the diagnostic challenge
**Difficulty:** Medium
**Estimated Time:** 25 minutes

## 🎯 Learning Goal

Revisit the string reversal problem from your diagnostic assessment. This time, you'll understand EXACTLY why it works, and you'll implement it correctly on your own.

## Reflection

**From your diagnostic:**
- Exercise 06 (Reverse String): **45 minutes**
- Your note: *"extremely hard due to my lack of knowledge and experience with runes"*
- You had to copy from documentation

**Now you know:**
- ✅ Strings are byte sequences
- ✅ Runes are characters (can be multiple bytes)
- ✅ How to iterate correctly with `range`
- ✅ How to convert between string/[]rune/[]byte
- ✅ How to use strings.Builder

**This exercise completes your journey.** You're going to solve it independently this time.

## The Problem

Reversing a string seems simple, but it's NOT just reversing bytes:

```go
s := "café"

// WRONG - Reverse bytes (breaks UTF-8!)
bytes := []byte(s)
// reverse bytes...
// Result: "éfac" (BROKEN! é's bytes are in wrong order)

// RIGHT - Reverse runes (characters)
runes := []rune(s)
// reverse runes...
// Result: "éfac" (CORRECT! é stays intact)
```

## Your Task

Implement three different approaches to string reversal:

### 1. ReverseSimple
Reverse a string by converting to []rune, reversing, converting back

### 2. ReverseBuilder
Reverse a string using strings.Builder (iterate backwards)

### 3. ReverseUTF8
Reverse a string using utf8.DecodeLastRune (advanced - like your diagnostic solution)

## Function Signatures

```go
func ReverseSimple(s string) string
func ReverseBuilder(s string) string
func ReverseUTF8(s string) string
```

## Examples

```go
ReverseSimple("hello")      // → "olleh"
ReverseSimple("café")       // → "éfac"
ReverseSimple("Hello, 世界") // → "界世 ,olleH"
ReverseSimple("👍🎉")        // → "🎉👍"

// All three functions should produce identical results
ReverseBuilder("hello")     // → "olleh"
ReverseUTF8("hello")        // → "olleh"
```

## Instructions

1. Open `reverse.go`
2. Implement all three functions
3. Run `go test -v`
4. **Compare with your diagnostic solution** - understand the differences
5. Update `../EXPLANATION.md` with your insights

## Hints

### Approach 1: ReverseSimple (Recommended - Easiest)

```go
func ReverseSimple(s string) string {
    // Convert to runes
    runes := []rune(s)

    // Reverse the slice (swap first with last, etc.)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    // Convert back to string
    return string(runes)
}
```

### Approach 2: ReverseBuilder (Efficient)

```go
import "strings"

func ReverseBuilder(s string) string {
    runes := []rune(s)
    var builder strings.Builder

    // Write runes in reverse order
    for i := len(runes) - 1; i >= 0; i-- {
        builder.WriteRune(runes[i])
    }

    return builder.String()
}
```

### Approach 3: ReverseUTF8 (Advanced - Similar to Diagnostic)

```go
import "unicode/utf8"

func ReverseUTF8(s string) string {
    // This approach processes the string from end to beginning
    // without creating a []rune slice first

    var builder strings.Builder
    bytes := []byte(s)

    // Decode runes from the end
    for len(bytes) > 0 {
        r, size := utf8.DecodeLastRune(bytes)
        builder.WriteRune(r)
        bytes = bytes[:len(bytes)-size]
    }

    return builder.String()
}
```

## 🧠 Think About

1. **Why does reversing bytes break UTF-8?**
   - Hint: Multi-byte runes need their bytes in order

2. **Which approach is most readable?**
   - Simplicity vs performance tradeoffs

3. **Which approach is most efficient?**
   - Memory allocations: ReverseSimple creates []rune, ReverseUTF8 doesn't

4. **Why did the diagnostic version work?**
   - Now you understand `utf8.DecodeLastRune`!

## What This Teaches

- Why byte reversal breaks Unicode strings
- Three valid approaches to the same problem
- Tradeoffs between readability and performance
- How your diagnostic solution actually worked
- Confidence in string manipulation

## Comparison with Diagnostic

**Your diagnostic code:**
```go
func Reverse(s string) string {
    reverse := ""
    b := []byte(s)
    for len(b) > 0 {
        r, size := utf8.DecodeLastRune(b)
        reverse += string(r)  // This is slow!
        b = b[:len(b)-size]
    }
    return reverse
}
```

**What you can improve:**
1. ❌ Used `reverse += string(r)` (slow concatenation)
2. ✅ Used `utf8.DecodeLastRune` (correct approach)

**Your improved version (ReverseUTF8):**
- Use `strings.Builder` instead of `+=`
- Same algorithm, better performance!

## Challenge Questions

After implementing all three, answer in your EXPLANATION.md:

1. Which implementation do you find most readable? Why?
2. Which would you use in production code? Why?
3. How does your understanding now compare to the diagnostic?
4. Could you explain this problem to someone else?

## After Completing

This is a milestone! You went from:
- ❌ 45 minutes, needed documentation, confused
- ✅ Understanding three approaches, knowing why each works

Update your JOURNAL.md with a reflection on this journey.

---

**This is your redemption arc!** The problem that took 45 minutes during the diagnostic is now a problem you fully understand. Well done! 🎉
