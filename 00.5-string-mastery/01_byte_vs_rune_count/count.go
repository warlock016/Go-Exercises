package count

import "unicode/utf8"

// ByteCount returns the number of bytes in the string
func ByteCount(s string) int {
	// TODO(human): Implement byte counting
	// Hint: Use the built-in len() function
	return len(s)
}

// RuneCount returns the number of runes (characters) in the string
func RuneCount(s string) int {
	// TODO(human): Implement rune counting
	// Hint: Import "unicode/utf8" and use utf8.RuneCountInString()
	return utf8.RuneCountInString(s)
}
