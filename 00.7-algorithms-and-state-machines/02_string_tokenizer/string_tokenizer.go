package stringtokenizer

// Tokenize splits a string into words, treating any whitespace as delimiter
// Multiple consecutive whitespace characters are treated as a single delimiter
func Tokenize(s string) []string {
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

	return nil // TODO(human): Replace with your implementation
}

// isWhitespace checks if a rune is whitespace (space, tab, newline, carriage return)
func isWhitespace(r rune) bool {
	// TODO(human): Implement whitespace check
	// Hint: Check if r equals ' ', '\t', '\n', or '\r'

	return false // TODO(human): Replace with your implementation
}
