package palindromecheck

import (
	"strings"
	"unicode"
)

// IsPalindrome checks if a string is a palindrome (case-sensitive).
// A palindrome reads the same forward and backward.
//
// Examples:
//   - IsPalindrome("racecar") returns true
//   - IsPalindrome("Racecar") returns false
//   - IsPalindrome("hello") returns false
func IsPalindrome(s string) bool {
	// TODO(human): Implement palindrome check

	runes := []rune(s)

	var forward strings.Builder
	var backward strings.Builder

	forward.WriteString(s)

	for i := len(runes) - 1; i >= 0; i-- {
		backward.WriteRune(runes[i])
	}

	return forward.String() == backward.String()
}

// IsPalindromeIgnoreCase checks if a string is a palindrome, ignoring case.
// Converts the string to lowercase before checking.
//
// Examples:
//   - IsPalindromeIgnoreCase("Racecar") returns true
//   - IsPalindromeIgnoreCase("RaceCar") returns true
//   - IsPalindromeIgnoreCase("Hello") returns false
func IsPalindromeIgnoreCase(s string) bool {
	// TODO(human): Implement case-insensitive palindrome check
	runes := []rune(s)

	var forward strings.Builder
	var backward strings.Builder

	forward.WriteString(s)

	for i := len(runes) - 1; i >= 0; i-- {
		backward.WriteRune(runes[i])
	}

	return strings.EqualFold(forward.String(), backward.String())
	// return false
}

// IsPalindromeIgnoreSpaces checks if a string is a palindrome,
// ignoring spaces and case. Only letters are considered.
//
// Examples:
//   - IsPalindromeIgnoreSpaces("A man a plan a canal Panama") returns true
//   - IsPalindromeIgnoreSpaces("race car") returns true
//   - IsPalindromeIgnoreSpaces("Was it a car or a cat I saw") returns true
//   - IsPalindromeIgnoreSpaces("hello world") returns false
func IsPalindromeIgnoreSpaces(s string) bool {
	// TODO(human): Implement palindrome check ignoring spaces and case

	lowerCase := strings.ToLower(s)
	var forwardStr strings.Builder
	var backwardStr strings.Builder
	// forward := []rune{}
	// backward := []rune{}

	for _, r := range lowerCase {
		if unicode.IsLetter(r) {
			forwardStr.WriteRune(r)
		}
	}

	for i := len([]rune(forwardStr.String())) - 1; i >= 0; i-- {
		backwardStr.WriteRune([]rune(forwardStr.String())[i])
	}

	return forwardStr.String() == backwardStr.String()
}

// LongestPalindromeSubstring finds the longest palindromic substring in s.
// If multiple palindromes have the same length, returns the first one found.
//
// Examples:
//   - LongestPalindromeSubstring("babad") returns "bab" or "aba"
//   - LongestPalindromeSubstring("cbbd") returns "bb"
//   - LongestPalindromeSubstring("racecar") returns "racecar"
func LongestPalindromeSubstring(s string) string {
	// TODO(human): Find the longest palindromic substring

	runes := []rune(s)
	maxLen := 0
	start := 0

	for i := range runes {
		oddLen := expandAroundCenter(runes, i, i)
		evenLen := expandAroundCenter(runes, i, i+1)
		currLen := max(oddLen, evenLen)

		if currLen > maxLen {
			maxLen = currLen
			start = i - (currLen-1)/2
		}
	}

	return string(runes[start : start+maxLen])
}

// expandAroundCenter is a helper function for LongestPalindromeSubstring.
// It expands around the center positions (left, right) and returns the length
// of the palindrome found.
//
// For odd-length palindromes: left == right (single center)
// For even-length palindromes: left + 1 == right (two centers)
func expandAroundCenter(runes []rune, left, right int) int {
	// TODO(human): Implement helper function to expand around center

	if len(runes) == 0 {
		return 0
	}

	// ['b','a','b','a','d']
	// [ 0 , 1,  2,  3,  4 ]
	// left =1, right = 1

	for {
		if left >= 0 && right < len(runes) && runes[left] == runes[right] {
			left--
			right++
		} else {
			return right - left - 1
		}
	}

}
