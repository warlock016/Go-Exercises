package wordwrapper

import "strings"

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
//   WrapText("hello world", 20) → "hello world"
//   WrapText("hello world", 5) → "hello\nworld"
//   WrapText("The quick brown fox", 10) → "The quick\nbrown fox"
func WrapText(text string, maxWidth int) string {
	// TODO(human): Implement text wrapping
	//
	// ALGORITHM:
	// 1. Handle edge cases (empty text, maxWidth <= 0)
	// 2. Split text into words using strings.Fields
	// 3. Build lines by adding words until maxWidth would be exceeded
	// 4. When adding a word would exceed maxWidth, start a new line
	// 5. Join all lines with \n
	//
	// KEY INSIGHT:
	// Before adding each word, calculate what the length WOULD BE:
	//   - Current line length
	//   - + 1 for space (if line isn't empty)
	//   - + word length
	// If this exceeds maxWidth, wrap.
	//
	// DETAILED ALGORITHM:
	//
	// STEP 1: Edge cases
	//   if text == "" {
	//       return ""
	//   }
	//   if maxWidth <= 0 {
	//       return text  // Can't wrap with no width limit
	//   }
	//
	// STEP 2: Split into words
	//   words := strings.Fields(text)
	//   // Fields splits on whitespace, removing empty strings
	//   // "hello  world" → ["hello", "world"]
	//
	//   if len(words) == 0 {
	//       return ""  // Text was only whitespace
	//   }
	//
	// STEP 3: Build lines
	//   var lines []string
	//   var currentLine strings.Builder
	//
	//   for _, word := range words {
	//       // Calculate length if we add this word
	//       futureLen := currentLine.Len()
	//       if currentLine.Len() > 0 {
	//           futureLen++  // +1 for space before word
	//       }
	//       futureLen += len(word)
	//
	//       // Check if we need to wrap
	//       if futureLen > maxWidth && currentLine.Len() > 0 {
	//           // Save current line and start new one
	//           lines = append(lines, currentLine.String())
	//           currentLine.Reset()
	//           currentLine.WriteString(word)
	//       } else {
	//           // Add word to current line
	//           if currentLine.Len() > 0 {
	//               currentLine.WriteRune(' ')
	//           }
	//           currentLine.WriteString(word)
	//       }
	//   }
	//
	// STEP 4: Add the last line
	//   if currentLine.Len() > 0 {
	//       lines = append(lines, currentLine.String())
	//   }
	//
	// STEP 5: Join lines
	//   return strings.Join(lines, "\n")
	//
	// TRICKY PARTS:
	// 1. The condition "futureLen > maxWidth && currentLine.Len() > 0"
	//    - The "currentLine.Len() > 0" prevents wrapping before first word
	//    - If a word is longer than maxWidth, it still gets a line
	//
	// 2. Adding spaces
	//    - Only add space if currentLine is not empty
	//    - This avoids leading spaces on new lines
	//
	// 3. Saving last line
	//    - The loop ends with words still in currentLine
	//    - Must append it after the loop!

	_ = strings.Fields      // Hint: split text into words
	_ = strings.Join        // Hint: join lines with \n
	_ = strings.Builder{}   // Hint: build each line efficiently

	// Your code here
	return ""
}
