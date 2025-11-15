package stringtokenizer

import "strings"

// Tokenize splits a string into words, treating any whitespace as delimiter
// Multiple consecutive whitespace characters are treated as a single delimiter
func Tokenize(s string) []string {

	tokens := make([]string, 0)
	var word strings.Builder
	inWord := false

	// "hello "
	for i, r := range s {
		switch isWhitespace(r) {
		case true:
			if inWord {
				inWord = false
				tokens = append(tokens, word.String())
				word.Reset()
			}
		case false:
			if !inWord {
				inWord = true
			}
			word.WriteRune(r)

			if i == len(s)-1 {
				tokens = append(tokens, word.String())
				word.Reset()
			}
		}
	}

	// TODO(human): Implement the tokenizer
	//
	// High-level algorithm:
	// 1. Create result slice and buffer (strings.Builder)
	// 2. Track state: are we currently inside a word?
	// 3. Loop through each character:
	//    - If whitespace and in word: flush buffer, transition to skipping
	//    - If whitespace and not in word: continue skipping
	//    - If non-whitespace: add to buffer, mark as in word
	// 4. After loop: flush final word if buffer has content
	//
	// Hint: Use a boolean flag `inWord` to track state
	// Hint: Use isWhitespace() helper function below
	// Hint: You'll need to import "strings" package for strings.Builder

	return tokens // TODO(human): Replace with your implementation
}

// isWhitespace checks if a rune is whitespace (space, tab, newline, carriage return)
func isWhitespace(r rune) bool {
	// TODO(human): Implement whitespace check
	// Hint: Check if r equals ' ', '\t', '\n', or '\r'

	if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
		return true
	}

	return false // TODO(human): Replace with your implementation
}
