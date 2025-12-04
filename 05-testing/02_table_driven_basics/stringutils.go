package stringutils

import (
	"strings"
	"unicode"
)

// Reverse returns the reversed string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome returns true if the string is a palindrome (case-sensitive)
func IsPalindrome(s string) bool {
	return s == Reverse(s)
}

// CountVowels returns the number of vowels (a, e, i, o, u) in the string (case-insensitive)
func CountVowels(s string) int {
	count := 0
	lower := strings.ToLower(s)
	for _, r := range lower {
		if unicode.IsLetter(r) && strings.ContainsRune("aeiou", r) {
			count++
		}
	}
	return count
}
