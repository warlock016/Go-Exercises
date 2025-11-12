package types

// This file shows the IDIOMATIC Go version of your working code.
// Compare this with your types.go to see the improvements.

import (
	"strings"
	"unicode"
)

// ============================================================================
// IDIOMATIC VERSIONS - What the IDE wants you to write
// ============================================================================

// CountLetters_Idiomatic - The Go-idiomatic way
func CountLetters_Idiomatic(s string) int {
	numLetters := 0

	// KEY INSIGHT: range over string directly gives you runes!
	// No need for []rune(s) conversion
	for _, r := range s { // ← Just "range s", not "range []rune(s)"
		if unicode.IsLetter(r) {
			numLetters++
		}
	}

	return numLetters
}

// CountDigits_Idiomatic - The Go-idiomatic way
func CountDigits_Idiomatic(s string) int {
	numDigits := 0

	for _, r := range s { // ← Direct iteration over string
		if unicode.IsDigit(r) {
			numDigits++
		}
	}

	return numDigits
}

// FilterLettersOnly_Idiomatic - The Go-idiomatic way
func FilterLettersOnly_Idiomatic(s string) string {
	var b strings.Builder

	for _, r := range s { // ← Direct iteration
		if unicode.IsLetter(r) {
			b.WriteRune(r)
		}
	}

	return b.String() // ← Direct return, no intermediate variable
}

// ============================================================================
// EXPLANATION: When DO you need []rune(s)?
// ============================================================================

// Example 1: When you need to MODIFY individual characters
func ReplaceCharAt(s string, index int, newChar rune) string {
	// Here you NEED []rune because strings are immutable
	runes := []rune(s) // ← Necessary here!
	runes[index] = newChar
	return string(runes)
}

// Example 2: When you need to ACCESS by index
func GetCharAt(s string, index int) rune {
	// Here you NEED []rune to safely index by character position
	runes := []rune(s) // ← Necessary here!
	return runes[index]
}

// Example 3: When you need to REVERSE (we covered this in Ex 05!)
func ReverseSimple(s string) string {
	// Here you NEED []rune to reverse the characters
	runes := []rune(s) // ← Necessary here!

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// ============================================================================
// PERFORMANCE COMPARISON
// ============================================================================

// Your way (works but allocates memory):
func YourWay(s string) int {
	col := []rune(s) // Allocates new slice, copies all runes
	count := 0
	for _, r := range col {
		if unicode.IsLetter(r) {
			count++
		}
	}
	return count
}

// Idiomatic way (no allocation):
func IdiomaticWay(s string) int {
	count := 0
	for _, r := range s { // No allocation, iterates directly
		if unicode.IsLetter(r) {
			count++
		}
	}
	return count
}

// For a 1000-character string:
// YourWay:       Allocates 4000 bytes (1000 runes × 4 bytes each)
// IdiomaticWay:  Allocates 0 bytes
//
// Both produce the same result, but idiomatic way is faster and uses less memory!

// ============================================================================
// RULE OF THUMB
// ============================================================================
//
// Convert to []rune when you need to:
// ✓ Modify individual characters
// ✓ Access characters by index
// ✓ Reverse or reorder characters
// ✓ Work with a mutable slice of characters
//
// DON'T convert to []rune when you're just:
// ✗ Iterating to read/check characters
// ✗ Counting characters with properties
// ✗ Filtering or transforming one-by-one
//
// The "range" keyword over a string already gives you runes!
