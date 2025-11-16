# Unicode Normalization Guide

**Purpose:** Understand how Unicode represents text, why normalization is necessary, and how to process international characters properly in Go.

**Prerequisites:** Basic understanding of strings and runes in Go

**Reading Time:** 30-40 minutes

---

## Table of Contents

1. [The Unicode Problem](#the-unicode-problem)
2. [Understanding Unicode Basics](#understanding-unicode-basics)
3. [Normalization Forms](#normalization-forms)
4. [Using Unicode Normalization in Go](#using-unicode-normalization-in-go)
5. [Practical Examples](#practical-examples)
6. [Common Pitfalls](#common-pitfalls)
7. [When to Normalize vs When Not To](#when-to-normalize-vs-when-not-to)
8. [Real-World Use Cases](#real-world-use-cases)

---

## The Unicode Problem

### The Same Character, Multiple Representations

Consider the character `é` (e with acute accent). In Unicode, this can be represented in **two different ways**:

**Method 1: Precomposed Character**
```
é = U+00E9 (single code point: LATIN SMALL LETTER E WITH ACUTE)
```

**Method 2: Decomposed Character**
```
é = U+0065 (e: LATIN SMALL LETTER E)
  + U+0301 (combining acute accent)
```

Both look identical when displayed, but they are **different byte sequences**:

```go
precomposed := "é"   // \u00e9
decomposed := "é"    // e\u0301

fmt.Println(precomposed == decomposed)  // false!
fmt.Println(len(precomposed))           // 2 bytes
fmt.Println(len(decomposed))            // 3 bytes
```

### Why This Matters

If you're searching for "café" in a document, you might miss it if:
- Your search string uses precomposed `é` (U+00E9)
- The document uses decomposed `e + ´` (U+0065 + U+0301)

This is why **normalization** exists: to convert different representations into a canonical form.

---

## Understanding Unicode Basics

### Code Points vs Code Units vs Graphemes

**Code Point:** A numeric value assigned to a character
```
'A' = U+0041
'é' = U+00E9
'🎉' = U+1F389
```

**Rune in Go:** A `rune` is an alias for `int32` and represents a Unicode code point.

**Combining Characters:** Special characters that modify the previous character
```
U+0301 = Combining Acute Accent (´)
U+0308 = Combining Diaeresis (¨)
U+0327 = Combining Cedilla (¸)
```

**Grapheme Cluster:** What a human perceives as a single character (may be multiple code points)
```
"é" could be:
  - 1 code point (U+00E9)
  - 2 code points (U+0065 + U+0301)
```

### How Strings Work in Go

```go
s := "café"  // Precomposed é

// Iteration gives you RUNES (code points)
for i, r := range s {
    fmt.Printf("Index %d: rune %c (U+%04X)\n", i, r, r)
}
// Output:
// Index 0: rune c (U+0063)
// Index 1: rune a (U+0061)
// Index 2: rune f (U+0066)
// Index 3: rune é (U+00E9)

// Byte length vs rune count
fmt.Println(len(s))                    // 5 bytes
fmt.Println(utf8.RuneCountInString(s)) // 4 runes
```

---

## Normalization Forms

Unicode defines **four normalization forms**. Each has a specific purpose.

### NFC (Canonical Composition)

**NFC = Normalization Form Canonical Composition**

- Composes characters into their precomposed form whenever possible
- **Most common form** for storage and transmission
- Used by macOS file systems

```go
input:  e + ´    (2 code points: U+0065 + U+0301)
output: é        (1 code point: U+00E9)
```

### NFD (Canonical Decomposition)

**NFD = Normalization Form Canonical Decomposition**

- Decomposes precomposed characters into base + combining marks
- Useful for **analysis and processing**
- Separates "what letter" from "what accent"

```go
input:  é        (1 code point: U+00E9)
output: e + ´    (2 code points: U+0065 + U+0301)
```

### NFKC (Compatibility Composition)

**NFKC = Normalization Form Compatibility Composition**

- Like NFC, but also replaces compatibility characters
- **Loses formatting information**
- Useful for searching/matching

```go
input:  ﬁ        (ligature: U+FB01)
output: fi       (2 letters: U+0066 + U+0069)

input:  ②        (circled 2: U+2461)
output: 2        (digit 2: U+0032)

input:  ℃        (degree celsius: U+2103)
output: °C       (degree sign + C: U+00B0 + U+0043)
```

### NFKD (Compatibility Decomposition)

**NFKD = Normalization Form Compatibility Decomposition**

- Like NFD, but also decomposes compatibility characters
- **Most aggressive normalization**
- Useful for text analysis where you want base forms

```go
input:  ﬁ        (ligature)
output: f + i    (2 separate letters)
```

### Summary Table

| Form | Composition | Compatibility | Use Case |
|------|-------------|---------------|----------|
| **NFC**  | Yes | No  | Storage, transmission, display |
| **NFD**  | No  | No  | Analysis, base character extraction |
| **NFKC** | Yes | Yes | Search, matching (loses formatting) |
| **NFKD** | No  | Yes | Deep analysis (most destructive) |

---

## Using Unicode Normalization in Go

### Installing the Package

The `golang.org/x/text` package is not in the standard library:

```bash
go get golang.org/x/text/unicode/norm
```

Add to `go.mod`:
```
require golang.org/x/text v0.14.0
```

### Basic Usage

```go
import (
    "fmt"
    "golang.org/x/text/unicode/norm"
)

func main() {
    // Two different representations of "café"
    precomposed := "café"  // é = U+00E9
    decomposed := "cafe\u0301"  // e + combining acute = U+0065 + U+0301

    fmt.Println(precomposed == decomposed)  // false

    // Normalize both to NFC
    nfc1 := norm.NFC.String(precomposed)
    nfc2 := norm.NFC.String(decomposed)

    fmt.Println(nfc1 == nfc2)  // true!

    // Normalize both to NFD
    nfd1 := norm.NFD.String(precomposed)
    nfd2 := norm.NFD.String(decomposed)

    fmt.Println(nfd1 == nfd2)  // true!
}
```

### API Methods

```go
// String normalization
normalized := norm.NFC.String(input)
normalized := norm.NFD.String(input)
normalized := norm.NFKC.String(input)
normalized := norm.NFKD.String(input)

// Byte slice normalization
normalized := norm.NFC.Bytes([]byte(input))

// Append to existing slice (efficient)
buf := make([]byte, 0, len(input)*2)
buf = norm.NFC.Append(buf, []byte(input)...)

// Check if already normalized (fast path)
isNormalized := norm.NFC.IsNormalString(input)
```

---

## Practical Examples

### Example 1: Case-Insensitive String Comparison

```go
import (
    "strings"
    "golang.org/x/text/unicode/norm"
)

func EqualFold(a, b string) bool {
    // Normalize to NFC, then compare case-insensitively
    a = norm.NFC.String(a)
    b = norm.NFC.String(b)
    return strings.EqualFold(a, b)
}

func main() {
    s1 := "café"           // Precomposed é
    s2 := "CAFE\u0301"     // Uppercase E + combining accent

    fmt.Println(s1 == s2)              // false
    fmt.Println(EqualFold(s1, s2))     // true!
}
```

### Example 2: Extracting Base Characters (Vowel Detection)

**Problem:** Detect if `é`, `è`, `ê`, `ë` are all vowels without hardcoding every variant.

**Solution:** Use NFD to decompose, then check the base character.

```go
import (
    "unicode"
    "golang.org/x/text/unicode/norm"
)

func isVowel(r rune) bool {
    // Decompose the rune (e.g., é → e + ´)
    decomposed := norm.NFD.String(string(r))

    // Get the first rune (base character)
    runes := []rune(decomposed)
    if len(runes) == 0 {
        return false
    }

    base := unicode.ToLower(runes[0])

    // Check if base is a, e, i, o, u
    return base == 'a' || base == 'e' || base == 'i' ||
           base == 'o' || base == 'u'
}

func CountVowels(s string) int {
    count := 0
    for _, r := range s {
        if isVowel(r) {
            count++
        }
    }
    return count
}

func main() {
    fmt.Println(CountVowels("café"))      // 2 (a, é)
    fmt.Println(CountVowels("naïve"))     // 3 (a, ï, e)
    fmt.Println(CountVowels("Zürich"))    // 2 (ü, i)
    fmt.Println(CountVowels("SEÑOR"))     // 2 (e, o)
}
```

### Example 3: Search with Normalization

```go
func Contains(haystack, needle string) bool {
    // Normalize both to NFC for consistent comparison
    haystack = norm.NFC.String(haystack)
    needle = norm.NFC.String(needle)
    return strings.Contains(haystack, needle)
}

func main() {
    text := "The café is open"     // Precomposed é
    query := "cafe\u0301"           // Decomposed é

    fmt.Println(strings.Contains(text, query))  // false (different bytes)
    fmt.Println(Contains(text, query))          // true (normalized)
}
```

### Example 4: Removing Accents (for Search Indexing)

```go
import (
    "unicode"
    "golang.org/x/text/runes"
    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
)

func RemoveAccents(s string) string {
    // 1. Decompose (NFD): é → e + ´
    // 2. Remove combining marks
    // 3. Recompose (NFC): e → e

    t := transform.Chain(
        norm.NFD,
        runes.Remove(runes.In(unicode.Mn)), // Mn = nonspacing marks
        norm.NFC,
    )

    result, _, _ := transform.String(t, s)
    return result
}

func main() {
    fmt.Println(RemoveAccents("café"))      // "cafe"
    fmt.Println(RemoveAccents("naïve"))     // "naive"
    fmt.Println(RemoveAccents("Zürich"))    // "Zurich"
    fmt.Println(RemoveAccents("SEÑOR"))     // "SENOR"
}
```

---

## Common Pitfalls

### Pitfall 1: Not All "Same Looking" Characters Normalize

```go
// LATIN SMALL LETTER A
a1 := "a"  // U+0061

// CYRILLIC SMALL LETTER A (looks identical!)
a2 := "а"  // U+0430

fmt.Println(a1 == a2)  // false
fmt.Println(norm.NFC.String(a1) == norm.NFC.String(a2))  // still false!
```

**Why:** These are fundamentally different letters from different scripts. Normalization doesn't convert Cyrillic to Latin.

### Pitfall 2: Normalization Changes Byte Length

```go
s := "é"  // Precomposed (2 bytes)
nfd := norm.NFD.String(s)  // Decomposed (3 bytes)

// Can't use fixed-size buffers!
buf := make([]byte, len(s))  // TOO SMALL after NFD
```

**Solution:** Always allocate extra space or use `norm.NFC.Append()`.

### Pitfall 3: Performance Cost

```go
// SLOW: Normalizing in a tight loop
for _, str := range millionStrings {
    if norm.NFC.String(str) == norm.NFC.String(query) {
        // ...
    }
}

// BETTER: Normalize query once, check if strings are already normalized
query = norm.NFC.String(query)
for _, str := range millionStrings {
    if norm.NFC.IsNormalString(str) && str == query {
        // Fast path: already normalized
    } else if norm.NFC.String(str) == query {
        // Slow path: needs normalization
    }
}
```

### Pitfall 4: NFD Creates Multiple Runes

```go
s := "é"  // Precomposed: 1 rune
nfd := norm.NFD.String(s)  // Decomposed: 2 runes (e + ´)

fmt.Println(utf8.RuneCountInString(s))    // 1
fmt.Println(utf8.RuneCountInString(nfd))  // 2

// Iterating gives you combining marks!
for _, r := range nfd {
    fmt.Printf("%c ", r)
}
// Output: e ´
```

---

## When to Normalize vs When Not To

### ✅ NORMALIZE when:

1. **Comparing user input**
   - Login usernames: "José" vs "Jose\u0301"
   - Search queries: "café" should match "cafe\u0301"

2. **Storing in databases**
   - Use NFC for consistent storage
   - Prevents duplicate entries with different encodings

3. **Hashing/checksums**
   - Normalize before hashing to get consistent results

4. **File paths** (on some systems)
   - macOS automatically uses NFD for filenames
   - Windows uses NFC
   - Normalize for cross-platform consistency

5. **Text analysis**
   - Use NFD to separate base characters from accents
   - Vowel detection, stemming, etc.

### ❌ DON'T NORMALIZE when:

1. **Displaying to users**
   - Show text exactly as entered
   - Users expect to see what they typed

2. **Preserving original data**
   - Archival systems
   - Legal documents
   - Digital signatures

3. **Working with fixed encodings**
   - If your entire system uses UTF-8 consistently
   - And you never compare strings from different sources

4. **Performance-critical code**
   - Real-time processing
   - Unless you've proven normalization is necessary

---

## Real-World Use Cases

### Use Case 1: Email Address Validation

```go
func NormalizeEmail(email string) string {
    // Emails should be NFC normalized per RFC 6531
    return norm.NFC.String(strings.ToLower(email))
}

func EmailsEqual(a, b string) bool {
    return NormalizeEmail(a) == NormalizeEmail(b)
}
```

### Use Case 2: Username Uniqueness

```go
func UsernameExists(username string, existingUsers []string) bool {
    // Normalize to NFC and lowercase for case-insensitive comparison
    normalized := strings.ToLower(norm.NFC.String(username))

    for _, existing := range existingUsers {
        if strings.ToLower(norm.NFC.String(existing)) == normalized {
            return true
        }
    }
    return false
}
```

### Use Case 3: Full-Text Search

```go
type SearchIndex struct {
    normalized map[string][]int  // normalized word → document IDs
}

func (idx *SearchIndex) Index(docID int, text string) {
    // Remove accents and normalize for searching
    normalized := RemoveAccents(strings.ToLower(text))
    words := strings.Fields(normalized)

    for _, word := range words {
        idx.normalized[word] = append(idx.normalized[word], docID)
    }
}

func (idx *SearchIndex) Search(query string) []int {
    normalized := RemoveAccents(strings.ToLower(query))
    return idx.normalized[normalized]
}

// Now "café", "cafe", "CAFÉ" all match the same documents!
```

### Use Case 4: Sorting International Names

```go
import "golang.org/x/text/collate"
import "golang.org/x/text/language"

func SortNames(names []string) {
    // Collation handles language-specific sorting rules
    cl := collate.New(language.English)

    sort.Slice(names, func(i, j int) bool {
        return cl.CompareString(names[i], names[j]) < 0
    })
}

// Correctly sorts: "Åse", "Bjørn", "Zürich"
```

---

## Summary & Best Practices

### Key Takeaways

1. **Unicode is complex**: The same visual character can have multiple encodings
2. **Normalization creates canonical forms**: Use NFC for storage, NFD for analysis
3. **Go's unicode/norm package**: Essential for international text processing
4. **NFD + base character extraction**: Solves the "accented vowel" problem elegantly

### Recommended Workflow

```go
// For COMPARISON (equality)
func Equal(a, b string) bool {
    return norm.NFC.String(a) == norm.NFC.String(b)
}

// For STORAGE (databases, files)
func Store(s string) string {
    return norm.NFC.String(s)
}

// For ANALYSIS (text processing)
func Analyze(s string) {
    nfd := norm.NFD.String(s)
    // Process decomposed form
}

// For SEARCHING (fuzzy matching)
func Search(query, text string) bool {
    query = RemoveAccents(strings.ToLower(query))
    text = RemoveAccents(strings.ToLower(text))
    return strings.Contains(text, query)
}
```

### Performance Tips

1. **Check if already normalized** before normalizing
2. **Normalize once** (at input time), not repeatedly
3. **Use NFC for storage** (more compact than NFD)
4. **Profile before optimizing** - normalization isn't always the bottleneck

---

## Further Reading

- [Unicode Normalization Forms (Unicode Standard)](https://unicode.org/reports/tr15/)
- [Go x/text package documentation](https://pkg.go.dev/golang.org/x/text)
- [Unicode Character Database](https://www.unicode.org/ucd/)
- [RFC 5198: Unicode Format for Network Interchange](https://www.rfc-editor.org/rfc/rfc5198)

---

**Created:** 2025-11-14
**Last Updated:** 2025-11-14
**Related Guides:** STRING_GUIDE.md, STRING_MASTERY_CONCEPTS.md
