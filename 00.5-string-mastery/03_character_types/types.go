package types

import (
	"strings"
	"unicode"
)

// IsLetter checks if the given rune is a letter
func IsLetter(r rune) bool {
	// TODO(human): Check if r is a letter
	// Hint: import "unicode" and use unicode.IsLetter(r)
	return unicode.IsLetter(r)
}

// IsDigit checks if the given rune is a digit (0-9)
func IsDigit(r rune) bool {
	// TODO(human): Check if r is a digit
	// Hint: use unicode.IsDigit(r)
	return unicode.IsDigit(r)
}

// CountLetters returns the number of letters in the string
func CountLetters(s string) int {
	// EXPLANATION: IDE Warning "should range over string, not []rune(string)"
	//
	// Your approach: col := []rune(s); for _, r := range col
	// IDE suggests: for _, r := range s
	//
	// WHY? When you use "range" over a string DIRECTLY, Go automatically
	// iterates over RUNES (not bytes)! You don't need to convert to []rune first.
	//
	// Your way works but is LESS efficient because:
	// 1. []rune(s) allocates a NEW slice in memory
	// 2. Copies all runes from the string into that slice
	// 3. Then iterates over the slice
	//
	// Better way: range directly over the string - no allocation, no copy!
	//
	// Both approaches give you runes, but direct ranging is more idiomatic and faster.

	// col := []rune(s) // This line is unnecessary - remove it for better code
	numLetters := 0

	for _, r := range s { // IDE wants: for _, r := range s
		if unicode.IsLetter(r) {
			numLetters++
		}
	}

	return numLetters
}

// CountDigits returns the number of digits in the string
func CountDigits(s string) int {
	// SAME ISSUE: You're converting to []rune unnecessarily
	// Just do: for _, r := range s

	// col := []rune(s) // Remove this line
	numDigits := 0

	for _, r := range s { // Change to: for _, r := range s
		if unicode.IsDigit(r) {
			numDigits++
		}
	}
	return numDigits
}

// FilterLettersOnly returns a new string containing only letters
func FilterLettersOnly(s string) string {
	// EXPLANATION: Why you needed intermediate variable "out"
	//
	// Your code: out = b.String(); return out
	// Better:     return b.String()
	//
	// ANSWER: You DON'T need the intermediate variable! You can return directly.
	// Go allows returning function call results directly.
	//
	// Why you might have seen an error earlier:
	// - Perhaps the variable declaration was in wrong scope
	// - Maybe you tried b.String directly without parentheses
	// - Or had a different syntax issue that's now resolved
	//
	// The intermediate variable "out" adds no value here - it's just extra code.
	// Always prefer the simpler: return b.String()

	// col := []rune(s) // Again, unnecessary! Use: for _, r := range s

	var b strings.Builder
	// var out string // This variable is unnecessary - remove it

	for _, r := range s { // Change to: for _, r := range s
		if unicode.IsLetter(r) {
			b.WriteRune(r)
		}
	}

	// out = b.String() // Remove this line
	return b.String() // Change to: return b.String()
}
