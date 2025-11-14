# Encoding/Decoding Patterns: A Comprehensive Guide

## Overview

**Encoding** transforms data from one representation to another. **Decoding** reverses the process. This document explores the fundamental patterns, challenges, and solutions that appear across all encoding schemes - from simple compression to complex serialization formats.

## Table of Contents

1. [Core Concepts](#core-concepts)
2. [The Fundamental Property: Round-Trip Invariance](#the-fundamental-property-round-trip-invariance)
3. [The Ambiguity Problem](#the-ambiguity-problem)
4. [Escape Sequences: The Solution](#escape-sequences-the-solution)
5. [Common Encoding Patterns](#common-encoding-patterns)
6. [Edge Cases in Encoding Schemes](#edge-cases-in-encoding-schemes)
7. [Design Checklist](#design-checklist)
8. [Real-World Examples](#real-world-examples)
9. [Debugging Encoding/Decoding Issues](#debugging-encodingdecoding-issues)

---

## Core Concepts

### What is Encoding?

**Encoding** is a transformation `E: D → R` where:
- `D` = Domain (original representation)
- `R` = Range (encoded representation)
- `E` = Encoding function

**Decoding** is the inverse transformation `D: R → D` where:
- `D(E(x)) = x` for all `x` in the domain

### Types of Encoding

| Type | Purpose | Example |
|------|---------|---------|
| **Compression** | Reduce size | Run-length, Huffman, LZ77 |
| **Serialization** | Convert data to transmittable format | JSON, Protocol Buffers, XML |
| **Character Encoding** | Represent text | UTF-8, ASCII, Base64 |
| **Escape Encoding** | Make data safe for transport | URL encoding, HTML entities |
| **Encryption** | Secure data | AES, RSA (not covered here) |

### Key Properties

1. **Correctness:** `Decode(Encode(x)) = x` (round-trip)
2. **Unambiguity:** Each encoded string has exactly one valid decoding
3. **Efficiency:** Encoding/decoding should be fast
4. **Size:** Encoded form should be small (for compression)
5. **Robustness:** Handle edge cases gracefully

---

## The Fundamental Property: Round-Trip Invariance

### Definition

For any encoding scheme to be **correct**, it must satisfy:

```
∀x ∈ Domain: Decode(Encode(x)) = x
```

This is called the **round-trip property** or **inverse property**.

### Why This Matters

If round-trip fails, data is **lost** or **corrupted**:

```go
// Bad encoding - loses information
Encode("hello world") → "hello"
Decode("hello") → "hello"
// ❌ Original was "hello world", not "hello"!
```

### Testing Round-Trip

**Always test this property independently:**

```go
func TestRoundTrip(t *testing.T) {
    testCases := []string{
        "simple",
        "edge cases: empty, single, unicode 😀",
        "special: 123!@#$%",
    }

    for _, original := range testCases {
        encoded := Encode(original)
        decoded := Decode(encoded)

        if decoded != original {
            t.Errorf("Round-trip failed!")
            t.Errorf("  Original: %q", original)
            t.Errorf("  Encoded:  %q", encoded)
            t.Errorf("  Decoded:  %q", decoded)
        }
    }
}
```

### When Round-Trip Can Fail

1. **Lossy encoding** (intentional): JPEG, MP3
2. **Ambiguous format** (bug): Decoder can't distinguish metadata from data
3. **Incomplete implementation** (bug): Encoder/decoder mismatch
4. **Edge case handling** (bug): Empty strings, boundary conditions

---

## The Ambiguity Problem

### What is Ambiguity?

Ambiguity occurs when a single encoded string can be decoded in **multiple ways**.

### Example: Run-Length Encoding

**Naive approach:**
```
Encode("111222") → "3132"
```

**Decoding problem:**
```
Is "3132":
  a) count=3, char='1', count=3, char='2' → "111222" ?
  b) count=31, char='3', count=2, char=??? → (incomplete)
  c) All digits, no characters → "" ?
```

The decoder **cannot determine** where counts end and characters begin when characters are also digits!

### General Form

Ambiguity arises when:
```
Metadata_Characters ∩ Data_Characters ≠ ∅
```

In other words: **When metadata and data share the same alphabet.**

### Real-World Examples

| Format | Metadata | Data | Overlap | Problem |
|--------|----------|------|---------|---------|
| Run-length encoding | Digits (counts) | Any character | Digit chars | Can't distinguish count from character |
| CSV | Commas (delimiters) | Text | Commas in text | Can't tell delimiter from data comma |
| URLs | `?, &, =` (syntax) | Text | Special chars in values | Can't tell syntax from data |
| Strings | `\` (escape), `"` (quotes) | Text | `\`, `"` in data | Can't tell string boundary |
| JSON | `{, }, :, ,` (syntax) | Strings | Special chars in values | Can't tell syntax from data |

### Recognizing Ambiguity

**Ask yourself:**
1. What characters have special meaning in my format? (metadata)
2. Can those characters also appear in the data?
3. If yes, how does the decoder distinguish them?

**Example:**
- Format: `count + character`
- Metadata: Digits (for count)
- Data: Any character (including digits!)
- Overlap: **YES** → Ambiguity problem!

---

## Escape Sequences: The Solution

### What are Escape Sequences?

An **escape sequence** is a special marker that says: "The next character is data, not metadata."

### Common Escape Patterns

#### 1. **Doubling/Repetition**

Used when: Metadata character is rare in data

```
Pattern: Write metadata character twice to mean literal character

Examples:
- Run-length encoding: "211" = count=2, char='1', marker='1'
- SQL: 'O''Brien' = literal O'Brien (doubled apostrophe)
- CSV: "Hello, ""World""" = literal Hello, "World"
```

**Advantages:**
- Simple to implement
- No extra characters needed
- Works well when special char is rare

**Disadvantages:**
- Data containing special char grows (doubled)
- Only works for single special character

#### 2. **Backslash Escaping**

Used when: Multiple special characters need escaping

```
Pattern: Prefix special characters with '\'

Examples:
- Strings: "He said \"Hi\"" = literal He said "Hi"
- Regex: "Match \." = literal dot, not "any character"
- JSON: "{\"name\": \"value\"}"
```

**Advantages:**
- Handles many special characters
- Widely understood convention
- Compact

**Disadvantages:**
- Backslash itself must be escaped: "\\"
- Can become hard to read (e.g., Windows paths: "C:\\Users\\...")

#### 3. **Delimiters/Framing**

Used when: Data is variable-length or complex

```
Pattern: Wrap data in special markers

Examples:
- HTML: <tag>data</tag>
- JSON: {"key": "value"}
- Length-prefixed: "5:hello" = length 5, then "hello"
```

**Advantages:**
- Clear boundaries
- Can nest
- Supports complex structures

**Disadvantages:**
- More verbose
- Requires parsing nested structures

#### 4. **Encoding/Translation**

Used when: Data must avoid certain characters entirely

```
Pattern: Replace problematic characters with safe equivalents

Examples:
- URL encoding: "hello world" → "hello%20world"
- Base64: Binary → printable ASCII
- HTML entities: "<" → "&lt;"
```

**Advantages:**
- Guarantees output alphabet
- Safe for transport (URLs, HTML, etc.)
- No ambiguity

**Disadvantages:**
- Data expansion (often significant)
- Requires lookup table
- Encoding/decoding overhead

#### 5. **Out-of-Band Signaling**

Used when: Can communicate metadata separately

```
Pattern: Store metadata separately from data

Examples:
- Protocol Buffers: Type info in schema, data separate
- HTTP: Headers (metadata) separate from body (data)
- Binary formats: Header with length/type, then payload
```

**Advantages:**
- No data expansion
- Efficient
- Clear separation

**Disadvantages:**
- Requires structured format
- More complex to implement
- Can't embed in simple string

### Choosing an Escape Strategy

| Use Case | Recommended Strategy |
|----------|---------------------|
| One special character, rare in data | Doubling |
| Multiple special characters | Backslash escaping |
| Structured data | Delimiters/framing |
| Must avoid certain characters | Encoding/translation |
| Binary protocol | Out-of-band signaling |

---

## Common Encoding Patterns

### Pattern 1: Count-Prefix Encoding

**Structure:** `<count><data>`

```
Examples:
- Length-prefixed strings: "5:hello"
- Run-length encoding: "3a" = 3 a's
- Pascal strings: [5, 'h', 'e', 'l', 'l', 'o']
```

**Edge Cases:**
- Empty data: "0:" or just ""?
- Count overflow: What if count > max int?
- Data contains digits: Use escape sequence!

### Pattern 2: Delimiter-Separated Values

**Structure:** `<item><delim><item><delim>...`

```
Examples:
- CSV: "a,b,c"
- TSV: "a\tb\tc"
- PATH: "/usr/local/bin"
```

**Edge Cases:**
- Empty items: "a,,c" (two empty items)
- Delimiter in data: Use escaping or quoting
- Trailing delimiter: "a,b,c," (3 or 4 items?)

### Pattern 3: Tagged/Typed Encoding

**Structure:** `<type><data>`

```
Examples:
- JSON: {"name": "John"} (object type, then key-value pairs)
- Protocol Buffers: Field number + wire type + value
- Type-length-value (TLV): <type:1byte><length:2bytes><value>
```

**Edge Cases:**
- Unknown types: How to skip?
- Type mismatch: Expected string, got number
- Nested types: Recursive parsing

### Pattern 4: State Machine Parsing

**Structure:** Multi-state decoder

```
States:
- READING_COUNT: Accumulate digits
- READING_CHAR: Read character
- READING_ESCAPE: Next char is literal
```

**Edge Cases:**
- Unexpected state transitions
- End of input in middle of sequence
- Invalid state combinations

### Pattern 5: Fixed-Width Encoding

**Structure:** Each element has predetermined size

```
Examples:
- IPv4 addresses: 4 bytes, fixed
- UUID: 16 bytes, fixed format
- Binary protocols: Header (fixed) + payload (variable)
```

**Edge Cases:**
- Alignment/padding
- Endianness (byte order)
- Version changes (breaking fixed width)

---

## Edge Cases in Encoding Schemes

### Universal Edge Cases

Every encoding scheme should handle:

#### 1. **Empty Input**

```go
Encode("") → ???
// Options:
// a) Empty string ""
// b) Special marker "EMPTY"
// c) Error

// Best practice: Empty input → empty output
Encode("") → ""
Decode("") → ""
```

#### 2. **Single Element**

```go
// Minimum valid input
Encode("a") → "1a"
Decode("1a") → "a"

// Test both minimum cases
```

#### 3. **Boundary Values**

```go
// Maximum count (for count-based encoding)
Encode(strings.Repeat("a", 999999)) → "999999a"

// Multi-byte characters
Encode("😀😀😀") → "3😀"

// Special characters
Encode("\n\t\r") → ???
```

#### 4. **Data Containing Metadata Characters**

```go
// The ambiguity problem!
Encode("111") → ??? // Digits in run-length encoding
Encode("a,b,c") → ??? // Commas in CSV
Encode("50%") → ??? // Percent in URL encoding
```

#### 5. **Maximum Length/Overflow**

```go
// What if count exceeds max int?
Encode(strings.Repeat("a", math.MaxInt64 + 1)) → ???

// Options:
// a) Split into multiple runs: "9999999a9999999a..."
// b) Use larger type: uint64
// c) Return error
```

#### 6. **Invalid Encoded Input**

```go
// What if decoder receives malformed data?
Decode("3") → ??? // Incomplete (no character)
Decode("abc") → ??? // No count
Decode("-5a") → ??? // Negative count

// Best practice: Return error or best-effort decode
```

#### 7. **Unicode and Multi-byte Characters**

```go
// Are you counting bytes or runes?
s := "café" // 4 runes, 5 bytes
len(s) // 5 (bytes) ❌
len([]rune(s)) // 4 (runes) ✓

// Always use runes for character-based encoding!
```

#### 8. **Whitespace**

```go
// Are spaces significant?
Encode("a b c") → "1a1 1b1 1c" // Spaces encoded
Decode("1a1 1b1 1c") → "a b c" // Spaces decoded

// Or are they delimiters?
```

#### 9. **Case Sensitivity**

```go
// Does case matter?
Encode("AaA") → "1A1a1A" // Case-sensitive
// vs
Encode("AaA") → "3A" // Case-insensitive (normalized)
```

#### 10. **Null Characters**

```go
// Can you encode null bytes?
Encode("\x00\x00\x00") → ???

// Common issue: C strings treat \x00 as terminator
```

---

## Design Checklist

When designing an encoding scheme, verify:

### Correctness
- [ ] Round-trip property holds for all inputs
- [ ] No ambiguity in decoding
- [ ] Handles all character types (ASCII, Unicode, control chars)

### Edge Cases
- [ ] Empty input handled
- [ ] Single element handled
- [ ] Maximum length handled (no overflow)
- [ ] Metadata characters in data handled (escape sequences)
- [ ] Invalid encoded input handled gracefully

### Efficiency
- [ ] No unnecessary copies or allocations
- [ ] Use `strings.Builder` for string construction
- [ ] Avoid repeated concatenation (`s += x` in loop)
- [ ] Consider space complexity (expansion factor)

### Robustness
- [ ] Doesn't panic on malformed input
- [ ] Clear error messages for invalid data
- [ ] Graceful degradation when possible

### Testing
- [ ] Test round-trip property
- [ ] Test all edge cases
- [ ] Test with random/fuzz data
- [ ] Benchmark performance

---

## Real-World Examples

### Example 1: URL Encoding

**Problem:** URLs can only contain certain ASCII characters. Spaces and special characters must be encoded.

**Solution:** Percent-encoding

```
Original: "hello world?foo=bar&baz"
Encoded:  "hello%20world%3Ffoo%3Dbar%26baz"

Rules:
- Space → %20
- ? → %3F
- = → %3D
- & → %26
```

**Ambiguity solved:** `%` is the escape character. Literal `%` becomes `%25`.

**Round-trip:**
```go
original := "100% correct!"
encoded := url.QueryEscape(original) // "100%25+correct%21"
decoded := url.QueryUnescape(encoded) // "100% correct!"
// ✓ Round-trip successful
```

### Example 2: JSON String Escaping

**Problem:** JSON strings are delimited by `"`. How to include `"` in the string?

**Solution:** Backslash escaping

```json
{
  "message": "He said \"Hi\" and smiled."
}
```

**Escape sequences:**
- `\"` → literal quote
- `\\` → literal backslash
- `\n` → newline
- `\t` → tab
- `\uXXXX` → Unicode character

**Ambiguity solved:** Backslash is the escape character.

### Example 3: CSV Quoting

**Problem:** CSV uses commas as delimiters. How to include commas in data?

**Solution:** Quote fields containing commas, double internal quotes

```csv
name,description
Alice,"Works at ACME, Inc."
Bob,"Said ""Hi"" to everyone"
```

**Rules:**
- If field contains `,` → wrap in `"`
- If field contains `"` → double it: `""`
- If field contains newline → wrap in `"`

**Ambiguity solved:** Doubling quotes inside quoted fields.

### Example 4: Base64 Encoding

**Problem:** Transmit binary data over text-only channels (email, JSON).

**Solution:** Encode 3 bytes → 4 ASCII characters

```
Binary: [0x48, 0x65, 0x6C] (Hel)
Base64: "SGVs"

Alphabet: A-Z, a-z, 0-9, +, /
Padding: = (when length not multiple of 3)
```

**Round-trip:**
```go
data := []byte("Hello, World!")
encoded := base64.StdEncoding.EncodeToString(data)
decoded, _ := base64.StdEncoding.DecodeString(encoded)
// string(decoded) == "Hello, World!" ✓
```

**No ambiguity:** Output alphabet is disjoint from binary input.

### Example 5: Protocol Buffers

**Problem:** Efficient binary serialization with schema evolution.

**Solution:** Tag-length-value encoding with out-of-band schema

```
Message:
  field 1 (int32): 150
  field 2 (string): "test"

Encoded:
  [tag:1, type:varint, value:150]
  [tag:2, type:length-delim, length:4, value:"test"]
```

**Advantages:**
- Compact (no field names in data)
- Fast parsing (binary format)
- Schema evolution (can add fields)

**Ambiguity solved:** Tags separate fields, types determine parsing.

---

## Debugging Encoding/Decoding Issues

### Symptom 1: Round-Trip Fails

**Check:**
1. Trace encoder output manually
2. Trace decoder input manually
3. Find where they diverge

**Example:**
```go
original := "111"
encoded := Encode(original)
fmt.Printf("Encoded: %q\n", encoded) // "31"
decoded := Decode(encoded)
fmt.Printf("Decoded: %q\n", decoded) // ""

// Trace decoder:
// Input: "31"
// - '3' is digit → count="3"
// - '1' is digit → count="31"
// - End of string, count="31", NO CHARACTER!
// - Returns empty string

// Bug: All-digit inputs don't trigger character output
```

### Symptom 2: Decoder Produces Wrong Output

**Check:**
1. Print decoder state at each step
2. Verify state transitions
3. Check for off-by-one errors

**Example:**
```go
func Decode(s string) string {
    for i := 0; i < len(s); i++ {
        fmt.Printf("Position %d: char=%c, state=%v\n", i, s[i], state)
        // ...
    }
}
```

### Symptom 3: Data Expansion Unexpectedly Large

**Check:**
1. Identify worst-case input
2. Calculate expansion factor
3. Verify escape sequences are minimal

**Example:**
```go
// Run-length encoding worst case:
input := "abcdefgh" // 8 chars
encoded := Encode(input) // "1a1b1c1d1e1f1g1h" // 16 chars
// Expansion: 200%!

// This is expected for no-repetition input
```

### Symptom 4: Unicode Characters Corrupted

**Check:**
1. Are you using `[]byte` instead of `[]rune`?
2. Are you using `len(s)` instead of `len([]rune(s))`?
3. Are you using `s[i]` instead of `range s`?

**Example:**
```go
// WRONG - byte-based
s := "café"
for i := 0; i < len(s); i++ {
    fmt.Printf("%c ", s[i]) // c a f Ã © (corrupted!)
}

// RIGHT - rune-based
for _, r := range s {
    fmt.Printf("%c ", r) // c a f é ✓
}
```

### Symptom 5: Decoder Hangs/Infinite Loop

**Check:**
1. Are you incrementing loop index correctly?
2. Are you advancing past consumed characters?
3. Are you handling end-of-input correctly?

**Example:**
```go
// BUG - infinite loop
for i := 0; i < len(s); {
    // Read count
    for unicode.IsDigit(rune(s[i])) {
        count += string(s[i])
        // BUG: forgot to increment i!
    }
}

// FIX
for i := 0; i < len(s); {
    for i < len(s) && unicode.IsDigit(rune(s[i])) {
        count += string(s[i])
        i++ // ✓
    }
}
```

---

## Key Takeaways

1. **Round-trip property is sacred:** Always test `Decode(Encode(x)) = x`

2. **Ambiguity is the enemy:** When metadata and data share the same alphabet, use escape sequences

3. **Choose escape strategy wisely:**
   - Rare special char → Doubling
   - Multiple special chars → Backslash
   - Complex structures → Delimiters
   - Transport safety → Translation (Base64, URL encoding)

4. **Edge cases matter:**
   - Empty input
   - Single element
   - Metadata chars in data
   - Unicode
   - Boundary values

5. **Test independently:** Don't rely solely on unit tests. Verify properties hold.

6. **Use runes for characters:** Go strings are UTF-8 bytes, but `range` gives you runes.

7. **Build incrementally:** Start with simple case, add edge cases one at a time.

8. **Trace when debugging:** Print state at each step to understand where logic diverges.

---

## Further Reading

### Academic Resources
- "Data Compression" - David Salomon
- "Automata Theory" - Parsing and state machines

### Practical Resources
- [Go strings package](https://pkg.go.dev/strings)
- [Unicode in Go](https://go.dev/blog/strings)
- [Base64 encoding](https://pkg.go.dev/encoding/base64)
- [JSON encoding](https://pkg.go.dev/encoding/json)

### Standards
- RFC 3986 (URL encoding)
- RFC 4648 (Base64/Base32/Base16)
- RFC 8259 (JSON)
- UTF-8 specification (RFC 3629)

---

**Version:** 1.0
**Last Updated:** 2025-11-14
**Related Exercises:** 10_run_length_encoding, future serialization exercises
