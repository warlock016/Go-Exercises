# Exercise 10: Run-Length Encoding

**Concept:** Compression algorithm using pattern recognition
**Difficulty:** Hard
**Estimated Time:** 40-45 minutes

## Learning Goal

Implement a classic compression algorithm that encodes repeated characters as count+character pairs. This exercise combines pattern recognition, state tracking, string building, and parsing. You'll learn how compression works and practice converting between different string representations.

## The Problem

Run-Length Encoding (RLE) is a simple compression technique:
- Replace consecutive identical characters with count + character
- "aaabbc" → "3a2b1c"
- "hello" → "1h1e2l1o"

This is useful for:
- Compressing data with many repeated characters (images, fax machines)
- Understanding basic compression concepts
- Practicing string transformation algorithms

The challenge: Handle edge cases like digits in the original string, decode back to original, and maintain correctness.

```go
Encode("aaabbc")    // → "3a2b1c"
Decode("3a2b1c")    // → "aaabbc"

// Round trip property: Decode(Encode(s)) == s
```

## Your Tasks

Implement two complementary functions:

### 1. Encode
Compress a string using run-length encoding.

```go
func Encode(s string) string
```

### 2. Decode
Decompress a run-length encoded string back to original.

```go
func Decode(s string) string
```

## Examples

### Encode Examples
```go
Encode("aaabbc")           // → "3a2b1c"
Encode("hello")            // → "1h1e2l1o"
Encode("aaa")              // → "3a"
Encode("abc")              // → "1a1b1c"
Encode("")                 // → ""
Encode("a")                // → "1a"
Encode("aaaaaaaaa")        // → "9a"
Encode("aabbccddee")       // → "2a2b2c2d2e"
Encode("mississippi")      // → "1m1i2s1i2s1i2p1i"
```

### Decode Examples
```go
Decode("3a2b1c")           // → "aaabbc"
Decode("1h1e2l1o")         // → "hello"
Decode("3a")               // → "aaa"
Decode("")                 // → ""
Decode("1a")               // → "a"
Decode("9a")               // → "aaaaaaaaa"
Decode("2a2b2c2d2e")       // → "aabbccddee"
```

### Round-Trip Property
```go
// For any string s, Decode(Encode(s)) should equal s
s := "hello world"
encoded := Encode(s)         // → "1h1e2l1o1 1w1o1r1l1d"
decoded := Decode(encoded)   // → "hello world"
// decoded == s  ✓
```

## Edge Cases to Handle

### 1. Empty Strings
```go
Encode("")  // → ""
Decode("")  // → ""
```

### 2. Single Characters
```go
Encode("a")      // → "1a"
Decode("1a")     // → "a"
```

### 3. No Repetition
```go
Encode("abcdef")  // → "1a1b1c1d1e1f"
```

### 4. All Same Character
```go
Encode("aaaaaaa")  // → "7a"
```

### 5. Digits in Original String (CHALLENGE!)
```go
Encode("aa11bb")   // → "2a2112b"
// This is ambiguous when decoding!
// Is "211" → 2 '1's and 1 '1', or 21 '1's?
// Solution: Always encode, even single chars
```

### 6. Multi-digit Counts
```go
Encode("aaaaaaaaaaaaa")  // → "13a" (count > 9)
Decode("13a")            // → "aaaaaaaaaaaaa"
```

## Instructions

1. Open `run_length_encoding.go`
2. Implement both `Encode` and `Decode` functions
3. Run `go test -v` to verify your solution
4. Make sure round-trip tests pass: `Decode(Encode(s)) == s`

## Hints

### Hint 1: Encode Algorithm

```go
func Encode(s string) string {
    if s == "" {
        return ""
    }

    var builder strings.Builder
    runes := []rune(s)

    count := 1
    currentRune := runes[0]

    for i := 1; i < len(runes); i++ {
        if runes[i] == currentRune {
            count++  // Same character, increment count
        } else {
            // Different character, write current run
            builder.WriteString(strconv.Itoa(count))
            builder.WriteRune(currentRune)

            // Start new run
            currentRune = runes[i]
            count = 1
        }
    }

    // Write the last run
    builder.WriteString(strconv.Itoa(count))
    builder.WriteRune(currentRune)

    return builder.String()
}
```

### Hint 2: Decode Algorithm

Decoding is trickier - you need to parse numbers and characters:

```go
import (
    "strconv"
    "strings"
    "unicode"
)

func Decode(s string) string {
    if s == "" {
        return ""
    }

    var builder strings.Builder
    runes := []rune(s)
    i := 0

    for i < len(runes) {
        // Read the count (one or more digits)
        countStr := ""
        for i < len(runes) && unicode.IsDigit(runes[i]) {
            countStr += string(runes[i])
            i++
        }

        // Convert count to int
        count, _ := strconv.Atoi(countStr)

        // Read the character
        if i < len(runes) {
            char := runes[i]
            // Write 'count' copies of 'char'
            for j := 0; j < count; j++ {
                builder.WriteRune(char)
            }
            i++
        }
    }

    return builder.String()
}
```

### Hint 3: Handling the Last Group

A common mistake in Encode:

```go
// BUG - forgets last group
for i := 1; i < len(runes); i++ {
    // ... process runs ...
}
return builder.String()  // Last run never written!

// FIX - write last run after loop
// ... loop ...
builder.WriteString(strconv.Itoa(count))
builder.WriteRune(currentRune)
```

### Hint 4: Parsing Multi-Digit Numbers

In Decode, counts can be multi-digit:

```go
// WRONG - only reads one digit
count := int(runes[i] - '0')

// RIGHT - read all consecutive digits
countStr := ""
for i < len(runes) && unicode.IsDigit(runes[i]) {
    countStr += string(runes[i])
    i++
}
count, _ := strconv.Atoi(countStr)
```

## Think About

1. **When is RLE effective?**
   - Good: "aaaaaabbbbbb" → "6a6b" (50% compression)
   - Bad: "abcdefgh" → "1a1b1c1d1e1f1g1h" (200% expansion!)
   - RLE works best with long runs of repeated data

2. **Why convert to []rune before encoding?**
   - To handle multi-byte Unicode characters correctly
   - "😀😀😀" should be "3😀", not corrupted bytes

3. **How would you handle the digit ambiguity problem?**
   - Current approach: always encode with counts
   - Alternative: use a delimiter (e.g., "3a,2b,1c")
   - Alternative: escape digits in original string

4. **What's the worst-case space complexity?**
   - Input: n characters, all different
   - Output: 2n characters (count + char for each)
   - So RLE can make things worse!

## What This Teaches

- **Pattern recognition** - detecting consecutive identical elements
- **State tracking** - maintaining current run and count
- **String building** - efficient construction with strings.Builder
- **Parsing** - reading numbers and characters from encoded format
- **Number conversion** - strconv.Itoa and strconv.Atoi
- **Unicode handling** - working with runes for correct character handling
- **Algorithm design** - encoding and decoding are inverse operations
- **Testing** - verifying round-trip property

## Common Mistakes to Avoid

1. **Not handling empty strings**
   ```go
   // BUG - crashes on empty input
   runes := []rune(s)
   currentRune := runes[0]  // Panic if s is empty!

   // FIX
   if s == "" {
       return ""
   }
   ```

2. **Forgetting the last run in Encode**
   ```go
   // BUG - last group not written
   for i := 1; i < len(runes); i++ {
       if runes[i] != currentRune {
           builder.WriteString(...)
           count = 1
           currentRune = runes[i]
       }
   }
   // Missing: write the last group!
   ```

3. **Only parsing single-digit counts in Decode**
   ```go
   // BUG - "13a" would be misread
   count := int(runes[i] - '0')

   // FIX - parse full number
   countStr := ""
   for i < len(runes) && unicode.IsDigit(runes[i]) {
       countStr += string(runes[i])
       i++
   }
   count, _ := strconv.Atoi(countStr)
   ```

4. **Using bytes instead of runes**
   ```go
   // WRONG - breaks on Unicode
   for i := 0; i < len(s); i++ {
       c := s[i]  // byte, not rune!
   }

   // RIGHT
   for _, r := range s {
       // r is a rune
   }
   ```

5. **Not incrementing index correctly in Decode**
   ```go
   // BUG - infinite loop
   for i < len(runes) {
       if unicode.IsDigit(runes[i]) {
           count++
           // Missing: i++
       }
   }
   ```

6. **String concatenation instead of Builder**
   ```go
   // SLOW - creates new string each time
   result := ""
   for ... {
       result += strconv.Itoa(count)
   }

   // FAST
   var builder strings.Builder
   for ... {
       builder.WriteString(strconv.Itoa(count))
   }
   ```

## Challenge Extensions (Optional)

After completing the basic version, try these:

1. **Variable-length encoding**: Use special marker for single chars
   ```go
   // "abc" → "abc" (no counts for singles)
   // "aabbc" → "2abc" (count only for runs > 1)
   ```

2. **EncodeWithDelimiter**: Use delimiter to avoid ambiguity
   ```go
   Encode("aa11bb") → "2a,211,2b"
   ```

3. **MaxRunLength**: Limit count to prevent overflow
   ```go
   // If count > 255, split into multiple runs
   Encode("aaa...aaa") → "255a255a43a" (for 553 a's)
   ```

4. **BitwiseRLE**: Encode binary strings (0s and 1s only)
   ```go
   // More efficient encoding for binary data
   ```

## Real-World Applications

- **Image compression**: BMP and PCX formats use RLE
- **Fax machines**: Compress mostly-white pages
- **Game development**: Compress tile maps and sprites
- **Network protocols**: Compress repetitive packet data
- **Data storage**: Compress log files with repeated patterns
- **Teaching compression**: Introduction to compression concepts

## After Completing

You now understand:
- How compression algorithms work at a basic level
- How to implement encoding and decoding as inverse operations
- How to parse structured text (numbers + characters)
- How to track state across consecutive elements
- Why compression effectiveness depends on data patterns
- The importance of round-trip testing

Run-length encoding is one of the simplest compression algorithms, but the concepts you learned (pattern detection, state tracking, encoding/decoding) apply to much more sophisticated algorithms like Huffman coding, LZ77, and more!

## Final Reflection

Congratulations! You've completed the String Mastery module. You've learned:

1. **Exercise 01-05**: Foundations (bytes vs runes, iteration, unicode, Builder)
2. **Exercise 06**: State tracking and transformations (title case)
3. **Exercise 07**: Maps and frequency analysis (character counter)
4. **Exercise 08**: Validation and boolean logic (validators)
5. **Exercise 09**: Complex algorithms (word wrapper)
6. **Exercise 10**: Compression and parsing (run-length encoding)

You now have a deep understanding of string manipulation in Go. These skills will serve you well in:
- Text processing and analysis
- Data validation and cleaning
- Algorithm implementation
- Real-world application development

**Well done! You've mastered string manipulation in Go!**

---

**Module complete!** Time to apply your skills to real projects.
