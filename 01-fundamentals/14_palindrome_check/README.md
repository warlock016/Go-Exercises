# Exercise 14: Palindrome Check

## Learning Goal

Master string manipulation, rune handling, and algorithmic thinking by implementing various palindrome detection strategies with different normalization rules.

## Problem Description

A palindrome is a word, phrase, or sequence that reads the same backward as forward. In this exercise, you'll implement multiple palindrome detection functions with increasing complexity:

1. Basic case-sensitive palindrome check
2. Case-insensitive palindrome check
3. Palindrome check ignoring spaces (and optionally punctuation)
4. Finding the longest palindromic substring within a string

This exercise combines your string mastery skills with algorithmic problem-solving.

## Function Signatures

```go
package palindromecheck

// IsPalindrome checks if a string is a palindrome (case-sensitive)
func IsPalindrome(s string) bool

// IsPalindromeIgnoreCase checks if a string is a palindrome, ignoring case
func IsPalindromeIgnoreCase(s string) bool

// IsPalindromeIgnoreSpaces checks if a string is a palindrome, ignoring spaces and case
func IsPalindromeIgnoreSpaces(s string) bool

// LongestPalindromeSubstring finds the longest palindromic substring
func LongestPalindromeSubstring(s string) string
```

## Examples

### IsPalindrome (case-sensitive)
```go
IsPalindrome("racecar")    // true
IsPalindrome("Racecar")    // false (different case)
IsPalindrome("hello")      // false
IsPalindrome("")           // true (empty string is palindrome)
IsPalindrome("a")          // true
IsPalindrome("noon")       // true
```

### IsPalindromeIgnoreCase
```go
IsPalindromeIgnoreCase("Racecar")         // true
IsPalindromeIgnoreCase("RaceCar")         // true
IsPalindromeIgnoreCase("Hello")           // false
IsPalindromeIgnoreCase("A")               // true
IsPalindromeIgnoreCase("Aba")             // true
```

### IsPalindromeIgnoreSpaces
```go
IsPalindromeIgnoreSpaces("A man a plan a canal Panama")  // true
IsPalindromeIgnoreSpaces("race car")                      // false
IsPalindromeIgnoreSpaces("Was it a car or a cat I saw")   // true
IsPalindromeIgnoreSpaces("hello world")                   // false
IsPalindromeIgnoreSpaces("A Santa at NASA")               // true
```

### LongestPalindromeSubstring
```go
LongestPalindromeSubstring("babad")        // "bab" or "aba" (both valid)
LongestPalindromeSubstring("cbbd")         // "bb"
LongestPalindromeSubstring("racecar")      // "racecar"
LongestPalindromeSubstring("hello")        // "ll"
LongestPalindromeSubstring("a")            // "a"
LongestPalindromeSubstring("")             // ""
LongestPalindromeSubstring("forgeeksskeegfor")  // "geeksskeeg"
```

## Instructions

1. Implement `IsPalindrome`
2. Implement `IsPalindromeIgnoreCase`
3. Implement `IsPalindromeIgnoreSpaces`
4. Implement `LongestPalindromeSubstring`
5. Run tests with `go test -v`

## Think About

1. Why do we need to convert strings to runes before checking palindromes?
2. What's the time complexity of each function? Can you optimize any of them?
3. For `LongestPalindromeSubstring`, why do we need to check both odd and even length palindromes?
4. How would you modify `IsPalindromeIgnoreSpaces` to also ignore punctuation?
5. What happens if you don't use `strings.Builder` and instead use string concatenation with `+=`?

## What This Teaches

- **Rune manipulation:** Working with Unicode-safe string operations
- **Two-pointer technique:** Efficient comparison from both ends
- **String normalization:** Filtering and transforming strings for comparison
- **Algorithmic thinking:** Expand-around-center technique for substring problems
- **Edge case handling:** Empty strings, single characters, Unicode
- **Performance awareness:** `strings.Builder` vs concatenation
- **Code reusability:** Building complex functions from simpler ones
