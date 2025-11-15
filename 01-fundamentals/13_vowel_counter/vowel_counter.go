package vowelcounter

import (
	"strings"
	"unicode"
)

// CountVowels counts the number of vowels (a, e, i, o, u) in the string.
// The comparison is case-insensitive.
func CountVowels(s string) int {
	// TODO(human): Implement vowel counting
	// Hint: Convert to lowercase, range over runes, check against vowel set
	return 0
}

// CountConsonants counts the number of consonants (letters that are not vowels) in the string.
// Non-letter characters (numbers, spaces, punctuation) are ignored.
func CountConsonants(s string) int {
	// TODO(human): Implement consonant counting
	// Hint: Use unicode.IsLetter() to identify letters, then check if NOT a vowel
	return 0
}

// VowelFrequency returns a map counting the frequency of each vowel in the string.
// The map keys are lowercase vowel runes.
func VowelFrequency(s string) map[rune]int {
	// TODO(human): Implement vowel frequency counting
	// Hint: Create map, range over lowercase string, increment counts for vowels
	return nil
}
