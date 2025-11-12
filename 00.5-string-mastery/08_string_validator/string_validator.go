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
//   IsValidEmail("user@example.com") → true
//   IsValidEmail("invalid") → false (no @)
//   IsValidEmail("@example.com") → false (no local part)
//   IsValidEmail("user@nodot") → false (domain has no .)
func IsValidEmail(s string) bool {
	// TODO(human): Implement email validation
	//
	// APPROACH:
	// 1. Count '@' symbols (must be exactly 1)
	// 2. Split on '@' to get local and domain parts
	// 3. Validate local part is not empty
	// 4. Validate domain contains '.'
	// 5. Validate domain doesn't end with '.'
	//
	// USEFUL FUNCTIONS:
	// - strings.Count(s, "@") → count occurrences of @
	// - strings.Split(s, "@") → split into [local, domain]
	// - strings.Contains(domain, ".") → check for dot
	// - strings.LastIndex(domain, ".") → find position of last dot
	//
	// PSEUDOCODE:
	//   atCount := strings.Count(s, "@")
	//   if atCount != 1 {
	//       return false
	//   }
	//
	//   parts := strings.Split(s, "@")
	//   local := parts[0]
	//   domain := parts[1]
	//
	//   if len(local) == 0 {
	//       return false
	//   }
	//
	//   if !strings.Contains(domain, ".") {
	//       return false
	//   }
	//
	//   lastDot := strings.LastIndex(domain, ".")
	//   if lastDot == len(domain)-1 {
	//       return false
	//   }
	//
	//   return true

	_ = strings.Count // Hint: count @ symbols
	_ = strings.Split // Hint: split on @ to get parts

	// Your code here
	return false
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
//   IsValidPassword("Strong123") → true
//   IsValidPassword("weak") → false (too short, no uppercase, no digit)
//   IsValidPassword("NoDigits") → false (no digit)
func IsValidPassword(s string) bool {
	// TODO(human): Implement password validation
	//
	// APPROACH:
	// 1. Check length (>= 8)
	// 2. Use boolean flags to track: hasUpper, hasLower, hasDigit
	// 3. Iterate through runes, set flags when conditions met
	// 4. Return true only if ALL flags are true
	//
	// USEFUL FUNCTIONS:
	// - unicode.IsUpper(r) → true if rune is uppercase letter
	// - unicode.IsLower(r) → true if rune is lowercase letter
	// - unicode.IsDigit(r) → true if rune is digit
	//
	// PSEUDOCODE:
	//   if len(s) < 8 {
	//       return false
	//   }
	//
	//   hasUpper := false
	//   hasLower := false
	//   hasDigit := false
	//
	//   for _, r := range s {
	//       if unicode.IsUpper(r) {
	//           hasUpper = true
	//       }
	//       if unicode.IsLower(r) {
	//           hasLower = true
	//       }
	//       if unicode.IsDigit(r) {
	//           hasDigit = true
	//       }
	//   }
	//
	//   return hasUpper && hasLower && hasDigit

	_ = unicode.IsUpper // Hint: check for uppercase
	_ = unicode.IsLower // Hint: check for lowercase
	_ = unicode.IsDigit // Hint: check for digits

	// Your code here
	return false
}

// ContainsOnly checks if string s contains only characters from the allowed set.
// Returns true if s is empty.
// Returns false if s contains any character not in allowed.
//
// Examples:
//   ContainsOnly("abc", "abcdef") → true
//   ContainsOnly("abc!", "abc") → false (! not allowed)
//   ContainsOnly("", "abc") → true (empty string OK)
func ContainsOnly(s string, allowed string) bool {
	// TODO(human): Implement character set validation
	//
	// APPROACH 1 (Efficient): Use a map as a set
	// 1. Build a map of allowed runes (map[rune]bool)
	// 2. For each rune in s, check if it's in the map
	// 3. Return false if any rune is not in the map
	//
	// APPROACH 2 (Simple): Use strings.ContainsRune
	// 1. For each rune in s, use strings.ContainsRune(allowed, r)
	// 2. Return false if any rune is not found
	//
	// Approach 1 is faster for long strings!
	//
	// PSEUDOCODE (Approach 1):
	//   // Build set of allowed characters
	//   allowedSet := make(map[rune]bool)
	//   for _, r := range allowed {
	//       allowedSet[r] = true
	//   }
	//
	//   // Check each character in s
	//   for _, r := range s {
	//       if !allowedSet[r] {
	//           return false
	//       }
	//   }
	//
	//   return true
	//
	// PSEUDOCODE (Approach 2):
	//   for _, r := range s {
	//       if !strings.ContainsRune(allowed, r) {
	//           return false
	//       }
	//   }
	//   return true

	_ = strings.ContainsRune // Hint: one approach uses this

	// Your code here
	return false
}
