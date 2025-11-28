package string_recursion

import "strings"

// ReverseString reverses a string recursively
func ReverseString(s string) string {
	// TODO(human): Implement
	var output strings.Builder
	runes := []rune(s)

	if len(runes) == 0 {
		return ""
	}

	last := runes[len(runes)-1]
	output.WriteRune(last)

	var next strings.Builder
	rest := runes[:len(runes)-1]
	for _, v := range rest {
		next.WriteRune(v)
	}

	return output.String() + ReverseString(next.String())
}

// IsPalindrome checks if a string is a palindrome (ignoring spaces and case)
func IsPalindrome(s string) bool {
	// TODO(human): Implement
	return false
}

// CountVowels counts the number of vowels in a string
func CountVowels(s string) int {
	chars := []rune(strings.ToLower(s))
	var count int

	if len(chars) == 0 {
		return 0
	}

	current := chars[0]
	next := chars[1:]

	switch current {
	case 'a', 'e', 'i', 'o', 'u':
		count++
	}

	// TODO(human): Implement
	return count + CountVowels(string(next))
}

// RemoveChar removes all occurrences of char from string s
// "hello", "d"
func RemoveChar(s string, char rune) string {
	// TODO(human): Implement

	var output strings.Builder
	chars := []rune(s)

	if len(chars) == 0 {
		return ""
	}

	current := chars[0]
	next := chars[1:]

	if current != char {
		output.WriteString(string(current))
	}

	return output.String() + RemoveChar(string(next), char)
}
