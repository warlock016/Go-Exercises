package stringutils

import "testing"

// TODO(human): Create a table-driven test for the Reverse function
// Structure:
//   func TestReverse(t *testing.T) {
//       tests := []struct {
//           name  string
//           input string
//           want  string
//       }{
//           {"test case name", "input", "expected"},
//           // Add 5-6 test cases
//       }
//
//       for _, tt := range tests {
//           got := Reverse(tt.input)
//           if got != tt.want {
//               t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
//           }
//       }
//   }
//
// Test cases to include:
// - Simple word: "hello" → "olleh"
// - Empty string: "" → ""
// - Single character: "a" → "a"
// - Palindrome: "racecar" → "racecar"
// - With spaces: "hello world" → "dlrow olleh"
// - Your choice of additional cases

// TODO(human): Create a table-driven test for the IsPalindrome function
// Structure:
//   tests := []struct {
//       name  string
//       input string
//       want  bool
//   }{
//       {"test case name", "input", true/false},
//   }
//
// Test cases to include:
// - True palindromes: "racecar", "noon"
// - Not palindromes: "hello", "golang"
// - Edge cases: "", "a"
// - Two characters: "aa" (true), "ab" (false)

// TODO(human): Create a table-driven test for the CountVowels function
// Structure:
//   tests := []struct {
//       name  string
//       input string
//       want  int
//   }{
//       {"test case name", "input", expectedCount},
//   }
//
// Test cases to include:
// - "hello" → 2 (e, o)
// - "xyz" → 0 (no vowels)
// - "aeiou" → 5 (all vowels)
// - "" → 0 (empty)
// - "AEIOU" → 5 (uppercase vowels)
// - "Hello World" → 3 (mixed case)
