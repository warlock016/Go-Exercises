package stringvalidator

import (
	"strings"
	"unicode"
)

// IsValidEmail performs basic email validation.
// This is a simplified version - real email validation is extremely complex!
//
// Rules:
//   - Must contain exactly one '@'
//   - Must have at least one character before '@' (local part)
//   - Must have at least one character after '@' (domain part)
//   - Domain must contain at least one '.'
//   - Domain must have at least one character after the last '.'
//
// Examples:
//
//	IsValidEmail("user@example.com") → true
//	IsValidEmail("invalid") → false (no @)
//	IsValidEmail("@example.com") → false (no local part)
//	IsValidEmail("user@nodot") → false (domain has no .)
func IsValidEmail(s string) bool {
	// TODO(human): Implement basic email validation
	//
	// High-level: Verify the @ symbol exists exactly once, and check structure
	// around it (local and domain parts, domain must have a dot, etc.)
	//
	// Consider: strings.Count, strings.Split, strings.Contains, strings.LastIndex
	// Documentation: https://pkg.go.dev/strings
	//
	// Pattern: Validation involves checking multiple conditions sequentially
	// Think: What makes an email valid structurally? Break the problem into
	// smaller checks (@ presence, parts exist, domain has dot, etc.)

	_ = strings.Count // Hint: count @ symbols
	_ = strings.Split // Hint: split on @ to get parts

	// Your code here
	// 1. Split string into three slices (prefix, suffix), by using the rune "@" as separator
	// 2. Check if prefix, suffix exist
	// 3. Check if prefix, suffix and domain are in the right order [prefix, suffix, domain]

	if strings.Count(s, "@") != 1 { // if either no "@" or multiple "@", then invalid email string -> early return
		return false
	}

	parts := strings.Split(s, "@") // split string into prefix and suffix via "@" separator
	// pos := strings.IndexRune(s, '@')
	prefix := parts[0] // contains any string before the "@"
	suffix := parts[1] // contains any string after the "@"

	if strings.Count(suffix, ".") == 0 { // check if suffix contains at least one "."
		return false
	}

	if len(prefix) == 0 || len(suffix) == 0 { // if prefix is empty, then invalid email address (e.g. @hotmail.com, john.doe@ )
		return false
	}

	// prefCheck := true

	for i, r := range prefix {
		if i == len(prefix)-1 && !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	// failing edge cases: domain..com, b.c, sub.domain.example.com, example.com, domain..com
	for i, r := range suffix {
		if i == 0 && !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			return false
		}
		if i == len(suffix)-1 && !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

// IsValidPassword validates password strength.
//
// Requirements:
//   - At least 8 characters long
//   - Contains at least one uppercase letter
//   - Contains at least one lowercase letter
//   - Contains at least one digit
//
// Examples:
//
//	IsValidPassword("Strong123") → true
//	IsValidPassword("weak") → false (too short, no uppercase, no digit)
//	IsValidPassword("NoDigits") → false (no digit)
func IsValidPassword(s string) bool {
	// TODO(human): Implement password validation
	//
	// High-level: Check length requirement, then iterate to verify presence of
	// required character types (uppercase, lowercase, digit)
	//
	// Consider: unicode.IsUpper, unicode.IsLower, unicode.IsDigit
	// Documentation: https://pkg.go.dev/unicode
	//
	// Pattern: Use boolean flags to track whether requirements have been met
	// Think: How can you track multiple conditions during a single iteration?
	// All conditions must be true for password to be valid.

	_ = unicode.IsUpper // Hint: check for uppercase
	_ = unicode.IsLower // Hint: check for lowercase
	_ = unicode.IsDigit // Hint: check for digits

	// Your code here

	var chars []rune
	containsDigit := false
	containsUpper := false
	containsLower := false
	matchLen := false

	for _, r := range s {
		chars = append(chars, r)

		if unicode.IsDigit(r) {
			containsDigit = true
		}
		if unicode.IsUpper(r) {
			containsUpper = true
		}

		if unicode.IsLower(r) {
			containsLower = true
		}
	}

	matchLen = len(chars) >= 8

	return (containsDigit && containsLower && containsUpper && matchLen)
}

// ContainsOnly checks if string s contains only characters from the allowed set.
// Returns true if s is empty.
// Returns false if s contains any character not in allowed.
//
// Examples:
//
//	ContainsOnly("abc", "abcdef") → true
//	ContainsOnly("abc!", "abc") → false (! not allowed)
//	ContainsOnly("", "abc") → true (empty string OK)
func ContainsOnly(s string, allowed string) bool {
	// TODO(human): Implement character set validation
	//
	// High-level: Check if every character in s exists in the allowed set
	//
	// Consider: strings.ContainsRune OR use a map[rune]bool as a set
	// Documentation: https://pkg.go.dev/strings#ContainsRune
	//
	// Pattern: Membership testing - is each element in the allowed collection?
	// Think: Two approaches possible - simple iteration with ContainsRune, or
	// build a set first for O(1) lookups. Which is better for long strings?

	_ = strings.ContainsRune // Hint: one approach uses this

	// Your code here

	chars := []rune{}
	refChars := make(map[rune]bool)

	for _, r := range s {
		chars = append(chars, r)
	}

	for _, r := range allowed {
		refChars[r] = true
	}

	for _, r := range s {
		if !refChars[r] {
			return false
		}
	}

	return true
}
