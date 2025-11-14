package wordwrapper

import (
	"strings"
)

// WrapText wraps text to fit within the specified maximum width.
// It breaks lines at word boundaries (spaces) and never breaks within a word.
//
// Rules:
//   - Lines should not exceed maxWidth characters
//   - Break only at word boundaries (spaces)
//   - Words longer than maxWidth are kept on a single line (not broken)
//   - Multiple consecutive spaces are normalized to single spaces
//   - Lines are joined with \n
//
// Examples:
//
//	WrapText("hello world", 20) → "hello world"
//	WrapText("hello world", 5) → "hello\nworld"
//	WrapText("The quick brown fox", 10) → "The quick\nbrown fox"
func WrapText(text string, maxWidth int) string {
	// TODO(human): Implement text wrapping
	//
	// High-level: Split text into words, then build lines by adding words until
	// the next word would exceed maxWidth. When that happens, start a new line.
	//
	// Consider: strings.Fields, strings.Join, strings.Builder
	// Documentation: https://pkg.go.dev/strings#Fields
	//
	// Pattern: Look-ahead calculation - before adding a word, calculate what the
	// length would be if you add it (current length + space + word length).
	// Think: When should you wrap? What about words longer than maxWidth?
	// Don't forget to handle the last line after the loop ends!
	//
	// Key insight: Calculate the "future length" before committing to add a word.
	// If it would exceed maxWidth AND the current line isn't empty, wrap.

	_ = strings.Fields    // Hint: split text into words
	_ = strings.Join      // Hint: join lines with \n
	_ = strings.Builder{} // Hint: build each line efficiently

	// Your code here
	// 1. Split string into words
	// 2. Count characters per word
	// 3. If len(word) < maxWidth, then append word to string and increase charCount.
	// 4. If l

	var output strings.Builder
	charCount := 0

	// strings.Join()

	fields := strings.Fields(text)

	if maxWidth <= 0 {
		for i, word := range fields {
			output.WriteString(word)

			if i != len(fields)-1 {
				output.WriteString(" ")
			}
		}
		return output.String()
	}

	for i, w := range fields {

		switch i {
		case 0: // first word
			output.WriteString(w)
			charCount += len(w)

		case len(fields) - 1: // last word
			if charCount+len(w)+1 <= maxWidth {
				output.WriteString(" ")
				output.WriteString(w)
			} else {
				output.WriteString("\n")
				output.WriteString(w)
			}

		default: // any word between first and last
			if charCount+len(w)+1 <= maxWidth {
				output.WriteString(" ")
				output.WriteString(w)
				charCount += len(w) + 1
			} else {
				output.WriteString("\n")
				output.WriteString(w)
				charCount = len(w)
			}
		}

	}

	return output.String()
}
