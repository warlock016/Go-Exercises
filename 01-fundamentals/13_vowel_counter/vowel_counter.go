package vowelcounter

import (
	"strings"
	"unicode"
)

// CountVowels counts the number of vowels (a, e, i, o, u) in the string.
// The comparison is case-insensitive.
func CountVowels(s string) int {
	// TODO(human): Implement vowel counting

	var count int
	lowerCase := strings.ToLower(s)

	for _, v := range lowerCase {
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		default:
			continue
		}
	}
	return count
}

// CountConsonants counts the number of consonants (letters that are not vowels) in the string.
// Non-letter characters (numbers, spaces, punctuation) are ignored.
func CountConsonants(s string) int {
	// TODO(human): Implement consonant counting
	var count int
	lowerCase := strings.ToLower(s)

	for _, v := range lowerCase {
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			continue
		default:
			if unicode.IsLetter(v) {
				count++
			}
		}
	}
	return count
}

// VowelFrequency returns a map counting the frequency of each vowel in the string.
// The map keys are lowercase vowel runes.
func VowelFrequency(s string) map[rune]int {
	// TODO(human): Implement vowel frequency counting

	freq := make(map[rune]int, 0)
	lowerCase := strings.ToLower(s)

	for _, v := range lowerCase {
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			freq[v]++
		default:
			continue
		}
	}
	return freq
}
