# Vowel Counter

## Learning Goal

Practice string iteration and character classification while working with case-insensitive matching and Unicode text.

## Problem Description

Create functions to analyze vowels and consonants in strings. You'll need to iterate through strings, classify characters as vowels or consonants, and count their frequencies. This exercise builds on your string mastery skills, particularly working with runes and handling case-insensitive comparisons.

## Function Signatures

```go
func CountVowels(s string) int
func CountConsonants(s string) int
func VowelFrequency(s string) map[rune]int
```

## Examples

```go
CountVowels("hello")           // returns 2 (e, o)
CountVowels("AEIOU")           // returns 5
CountVowels("rhythm")          // returns 0
CountVowels("café")            // returns 2 (a, é)
CountVowels("")                // returns 0

CountConsonants("hello")       // returns 3 (h, l, l)
CountConsonants("AEIOU")       // returns 0
CountConsonants("123 abc!")    // returns 3 (b, c - ignores numbers/punctuation)
CountConsonants("")            // returns 0

VowelFrequency("hello")        // returns map[rune]int{'e': 1, 'o': 1}
VowelFrequency("AEIOU")        // returns map[rune]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
VowelFrequency("Programming")  // returns map[rune]int{'a': 1, 'i': 1, 'o': 1}
VowelFrequency("")             // returns map[rune]int{}
```

## Instructions

1. Implement `CountVowels`
2. Implement `CountConsonants`
3. Implement `VowelFrequency`
4. Run tests with `go test -v`

## Think About

1. Why do we use `unicode.IsLetter()` instead of checking if a character is between 'a'-'z'?
2. How does `range` over a string give you runes instead of bytes?
3. What happens to accented vowels like 'é' or 'ñ' in your implementation?
4. Could you make a helper function `isVowel(r rune) bool` to avoid repeating code?

## What This Teaches

- **String iteration:** Using `range` to iterate over runes in a string
- **Character classification:** Using `unicode` package for proper letter detection
- **Case-insensitive comparison:** Converting to lowercase for matching
- **Map operations:** Building frequency maps dynamically
- **Unicode awareness:** Understanding that not all vowels are a/e/i/o/u in all languages
