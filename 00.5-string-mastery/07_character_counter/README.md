# Exercise 07: Character Counter

**Concept:** Maps, rune frequency analysis, and character classification
**Difficulty:** Medium
**Estimated Time:** 25-30 minutes

## Learning Goal

Master using maps to count and analyze character frequencies. Learn to combine rune iteration with map operations and discover how to find extremes (most common character). This exercise introduces you to frequency analysis, a fundamental technique in text processing.

## The Problem

Counting character frequencies is essential for:
- Text analysis and statistics
- Compression algorithms (like the one in Exercise 10)
- Detecting anagrams
- Finding patterns in text
- Building frequency tables for encryption

```go
s := "hello"
// Count: h:1, e:1, l:2, o:1

s := "aabbcc"
// Most common: could be a, b, or c (all appear twice)

IsAnagram("listen", "silent")  // → true (same letters, same counts)
IsAnagram("hello", "world")     // → false (different letters/counts)
```

## Your Tasks

Implement three related functions:

### 1. CountRunes
Count the frequency of each rune in a string.

```go
func CountRunes(s string) map[rune]int
```

### 2. MostCommon
Find the most frequently occurring rune. If there's a tie, return the one that appears first in the string.

```go
func MostCommon(s string) rune
```

### 3. IsAnagram
Check if two strings are anagrams (contain the same characters with the same frequencies, ignoring case and spaces).

```go
func IsAnagram(s1, s2 string) bool
```

## Examples

### CountRunes Examples
```go
CountRunes("hello")
// → map[h:1 e:1 l:2 o:1]

CountRunes("aabbcc")
// → map[a:2 b:2 c:2]

CountRunes("Hello, World!")
// → map[H:1 e:1 l:3 o:2 ,:1  :1 W:1 r:1 d:1 !:1]

CountRunes("café")
// → map[c:1 a:1 f:1 é:1]

CountRunes("")
// → map[] (empty map)
```

### MostCommon Examples
```go
MostCommon("hello")      // → 'l' (appears 2 times)
MostCommon("aabbcc")     // → 'a' (tie: all appear twice, 'a' is first)
MostCommon("mississippi") // → 'i' (appears 4 times)
MostCommon("a")          // → 'a'
MostCommon("")           // → '\x00' (or 0, the zero value for rune)
```

### IsAnagram Examples
```go
IsAnagram("listen", "silent")           // → true
IsAnagram("Hello", "hello")             // → true (case insensitive)
IsAnagram("Astronomer", "Moon starer")  // → true (ignore spaces)
IsAnagram("hello", "world")             // → false
IsAnagram("aab", "abb")                 // → false
IsAnagram("", "")                       // → true
IsAnagram("a", "")                      // → false
```

## Instructions

1. Open `character_counter.go`
2. Implement all three functions
3. Run `go test -v` to verify your solution
4. Understand how maps work with rune keys

## Hints

### Hint 1: CountRunes Structure

```go
func CountRunes(s string) map[rune]int {
    counts := make(map[rune]int)

    for _, r := range s {
        counts[r]++  // Increment count for this rune
    }

    return counts
}
```

**Key insight:** In Go, accessing a non-existent map key returns the zero value (0 for int). So `counts[r]++` works even the first time!

### Hint 2: MostCommon Logic

```go
func MostCommon(s string) rune {
    counts := CountRunes(s)

    var mostCommon rune
    maxCount := 0

    // Need to iterate in order of appearance...
    // How do we do that?

    return mostCommon
}
```

**Problem:** Maps are unordered! To find the "first" in case of ties, iterate the original string:

```go
for _, r := range s {
    if counts[r] > maxCount {
        maxCount = counts[r]
        mostCommon = r
    }
}
```

But this has duplicates... You need to track what you've seen:

```go
seen := make(map[rune]bool)
for _, r := range s {
    if !seen[r] {
        if counts[r] > maxCount {
            maxCount = counts[r]
            mostCommon = r
        }
        seen[r] = true
    }
}
```

### Hint 3: IsAnagram Approach

Two strings are anagrams if they have identical character frequency maps (ignoring case and spaces).

```go
func IsAnagram(s1, s2 string) bool {
    // 1. Normalize both strings (lowercase, remove spaces)
    // 2. Count runes in both
    // 3. Compare the maps
}
```

**Normalizing a string:**
```go
import "unicode"

func normalize(s string) string {
    var builder strings.Builder
    for _, r := range s {
        if !unicode.IsSpace(r) {  // Skip spaces
            builder.WriteRune(unicode.ToLower(r))
        }
    }
    return builder.String()
}
```

**Comparing maps:**
```go
// Check if maps have same length
if len(map1) != len(map2) {
    return false
}

// Check if all keys and values match
for key, val1 := range map1 {
    val2, exists := map2[key]
    if !exists || val1 != val2 {
        return false
    }
}

return true
```

### Hint 4: Full IsAnagram Solution

```go
func IsAnagram(s1, s2 string) bool {
    // Normalize strings
    n1 := normalize(s1)
    n2 := normalize(s2)

    // Count runes
    counts1 := CountRunes(n1)
    counts2 := CountRunes(n2)

    // Compare maps
    if len(counts1) != len(counts2) {
        return false
    }

    for r, count := range counts1 {
        if counts2[r] != count {
            return false
        }
    }

    return true
}
```

## Think About

1. **Why do map accesses return zero values for missing keys?**
   - This makes counting patterns very convenient!
   - No need to check `if _, exists := map[key]; !exists`

2. **Why is MostCommon more complex than CountRunes?**
   - Finding the maximum requires iteration
   - Handling ties requires tracking order

3. **What's the time complexity of these functions?**
   - CountRunes: O(n) where n is string length
   - MostCommon: O(n) for counting + O(n) for finding max = O(n)
   - IsAnagram: O(n + m) where n, m are string lengths

4. **Why normalize for anagrams?**
   - "Listen" and "Silent" should match despite case
   - "Moon starer" and "Astronomer" should match despite spaces

## What This Teaches

- **Map operations** - creating, updating, and iterating maps
- **Rune keys** - using runes (not strings) as map keys
- **Frequency counting** - a fundamental algorithm pattern
- **Finding extremes** - tracking maximum while iterating
- **Map comparison** - checking if two maps are equal
- **String normalization** - preprocessing for meaningful comparisons
- **Zero values** - leveraging Go's default values for convenience

## Common Mistakes to Avoid

1. **Iterating bytes instead of runes**
   ```go
   // WRONG - breaks on multi-byte characters
   for i := 0; i < len(s); i++ {
       counts[s[i]]++  // s[i] is a byte!
   }

   // RIGHT
   for _, r := range s {
       counts[r]++  // r is a rune
   }
   ```

2. **Not handling empty strings**
   ```go
   func MostCommon(s string) rune {
       // Need to handle s == ""
       if s == "" {
           return 0  // or '\x00', the zero value for rune
       }
       // ...
   }
   ```

3. **Forgetting maps are unordered**
   ```go
   // WRONG - map iteration order is random
   for r := range counts {
       if counts[r] > maxCount {
           // This might not respect "first occurrence"
       }
   }

   // RIGHT - iterate original string
   for _, r := range s {
       if !seen[r] && counts[r] > maxCount {
           // Respects order of first occurrence
       }
   }
   ```

4. **Incorrect map comparison**
   ```go
   // WRONG - can't compare maps with ==
   if counts1 == counts2 {  // Compile error!
   }

   // RIGHT - compare manually
   if len(counts1) != len(counts2) {
       return false
   }
   for k, v := range counts1 {
       if counts2[k] != v {
           return false
       }
   }
   ```

5. **Not normalizing anagram inputs**
   ```go
   // WRONG - case sensitive
   IsAnagram("Listen", "Silent")  // → false (should be true)

   // RIGHT - normalize first
   normalize := func(s string) string {
       // Convert to lowercase, remove spaces
   }
   ```

## Challenge Extensions (Optional)

After completing the basic version, try these:

1. **MostCommonLetter**: Like MostCommon but only count letters (ignore punctuation, spaces, digits)
2. **LeastCommon**: Find the least frequent rune
3. **FrequencyList**: Return a sorted list of runes by frequency (descending)
4. **CountOnlyLetters**: Like CountRunes but only count letters (case-insensitive)

## Real-World Applications

- **Compression**: Huffman coding uses character frequencies
- **Cryptography**: Frequency analysis breaks simple ciphers
- **Spell checkers**: Anagram detection finds similar words
- **Text analysis**: Word/character frequency in documents
- **Data deduplication**: Finding repeated patterns

## After Completing

You now understand:
- How to use maps for counting and frequency analysis
- How to work with runes as map keys
- How to find extremes (max/min) in collections
- How to compare complex data structures
- How to normalize data for meaningful comparisons

Maps are one of Go's most powerful built-in types. Mastering them opens up many algorithmic possibilities!

---

**Next up:** Exercise 08 - String Validator (validation logic and complex conditions)
