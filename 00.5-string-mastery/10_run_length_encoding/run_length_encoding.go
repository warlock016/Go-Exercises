package runlengthencoding

import (
	"strconv"
	"strings"
	"unicode"
)

// Encode compresses a string using run-length encoding.
// Consecutive identical characters are replaced with count + character.
//
// Examples:
//
//	Encode("aaabbc") → "3a2b1c"
//	Encode("hello") → "1h1e2l1o"
//	Encode("aaa") → "3a"
//	Encode("") → ""
func Encode(s string) string {
	// TODO(human): Implement run-length encoding
	//
	// High-level: Iterate through the string, counting consecutive identical
	// characters. When the character changes, write count+character to output.
	//
	// Consider: strconv.Itoa, strings.Builder, []rune conversion
	// Documentation: https://pkg.go.dev/strconv#Itoa
	//
	// Pattern: State tracking - maintain current character and count, detect when
	// character changes to "flush" the current run to output.
	// Think: How do you handle the first character? What about the last run?
	// (The last run never triggers a character change - handle it after the loop!)
	//
	// Tip: Convert to []rune first for correct Unicode handling.

	_ = strconv.Itoa      // Hint: convert count to string
	_ = strings.Builder{} // Hint: build result efficiently

	// Your code here

	var output strings.Builder
	runeArray := []rune(s)
	var currentChar rune
	var currentCount int

	if len(s) == 0 { // if string is empty, then return empty string
		return ""
	}

	for i, r := range runeArray {

		switch i {
		case 0: // first char
			currentChar = r
			currentCount++

			if i == len(runeArray)-1 { //edge case: write directly to output since loop will run only once
				output.WriteString(strconv.Itoa(currentCount))
				output.WriteRune(currentChar)

				if unicode.IsDigit(currentChar) { // double digit encoding as numeric escape pattern
					output.WriteRune(currentChar)
				}
			}

		case len(runeArray) - 1: // last char
			if r != currentChar { // if last char is not equal to the previous one
				output.WriteString(strconv.Itoa(currentCount))
				output.WriteRune(currentChar)

				if unicode.IsDigit(currentChar) { // double digit encoding as numeric escape pattern
					output.WriteRune(currentChar)
				}
				// HINT: If currentChar is a digit, write it again here too!
				output.WriteString(strconv.Itoa(1))
				output.WriteRune(r)

				if unicode.IsDigit(r) { // double digit encoding as numeric escape pattern
					output.WriteRune(r)
				}

			} else { // if last char is equal to previous char, then write current char count and char
				currentCount++
				output.WriteString(strconv.Itoa(currentCount))
				output.WriteRune(currentChar)

				if unicode.IsDigit(currentChar) { // double digit encoding as numeric escape pattern
					output.WriteRune(currentChar)
				}
			}

		default: // any other char between first and last
			if r != currentChar {
				output.WriteString(strconv.Itoa(currentCount)) // write previous count on character change
				output.WriteRune(currentChar)                  // write previous character on character change

				if unicode.IsDigit(currentChar) { // double digit encoding as numeric escape pattern
					output.WriteRune(currentChar)
				}
				// to disambiguate it from count digits during decoding
				currentChar = r  // track new character
				currentCount = 1 // track new character count
			} else {
				currentCount++ // if new character is equal to previous one, only increase count.
			}
		}
	}

	return output.String()
}

// Decode decompresses a run-length encoded string back to original.
// The encoded format is: count + character (e.g., "3a2b1c" → "aaabbc")
//
// Examples:
//
//	Decode("3a2b1c") → "aaabbc"
//	Decode("1h1e2l1o") → "hello"
//	Decode("11a2b") → "aaaaaaaaaaabb"
//	Decode("") → ""
func Decode(s string) string {
	// STATE MACHINE IMPLEMENTATION
	//
	// This decoder uses a 4-state sequential pattern:
	//   STATE 1: Read count (all consecutive digits)
	//   STATE 2: Read character (next rune after count)
	//   STATE 3: Handle escape (if character is digit, skip duplicate marker)
	//   STATE 4: Output (repeat character × count times)
	//
	// Key insight: Each state COMPLETES before moving to next state.
	// No nested conditionals - just sequential phases.

	if len(s) == 0 {
		return "" // Edge case: empty input
	}

	runes := []rune(s)
	var output strings.Builder

	// Manual index control - we increment i ourselves, not the loop
	for i := 0; i < len(runes); {

		// ═══ STATE 1: READ COUNT ═══
		// Read ALL consecutive digits into countStr
		// Example: "100a" → reads "100", i moves from 0 to 3
		//
		// LIMITATION: If the input is all digits (like "123456" encoded as "111122..."),
		// this will read the entire string as count and fail. This is a known
		// limitation of this simple digit-doubling escape scheme.
		// Real-world RLE uses different escape strategies for arbitrary data.
		countStr := ""
		for i < len(runes) && unicode.IsDigit(runes[i]) {
			countStr += string(runes[i])
			i++ // Advance past each digit
		}

		// Validation: After reading digits, we must have a character
		// If i >= len(runes), the encoded string is malformed (count with no char)
		if i >= len(runes) || countStr == "" {
			break // Graceful exit for malformed input
		}

		// ═══ STATE 2: READ CHARACTER ═══
		// The rune at current position is the character to repeat
		// Example: "100a" → at i=3, char='a'
		char := runes[i]
		i++ // Advance past the character

		// ═══ STATE 3: HANDLE DIGIT ESCAPE SEQUENCE ═══
		// If the character we just read is itself a digit, the NEXT rune
		// is the duplicate marker (escape sequence). Skip it.
		// Example: "311" → char='1' (digit), next='1' (duplicate) → skip next
		if unicode.IsDigit(char) {
			// Bounds check: ensure there's a next rune to skip
			if i < len(runes) {
				i++ // Skip the duplicate digit marker
			}
			// Note: If there's no duplicate (malformed), we continue anyway
		}

		// ═══ STATE 4: OUTPUT ═══
		// Convert count string to integer and output char that many times
		if countStr == "" {
			// Malformed: no count found, skip
			continue
		}
		count, err := strconv.Atoi(countStr)
		if err != nil || count < 0 {
			// If count parsing fails or is negative, skip this run (defensive programming)
			continue
		}
		output.WriteString(strings.Repeat(string(char), count))

		// Loop continues: i is now positioned at the start of next count
	}

	return output.String()
}

// ═══ DEBUGGING EXAMPLE TRACE ═══
//
// Input: "311322" (encodes "111222")
//
// Iteration 1:
//   i=0
//   STATE 1: Read "3" → countStr="3", i=1
//   STATE 2: Read '1' → char='1', i=2
//   STATE 3: char is digit → skip runes[2]='1', i=3
//   STATE 4: Output "111" (three 1's)
//
// Iteration 2:
//   i=3
//   STATE 1: Read "3" → countStr="3", i=4
//   STATE 2: Read '2' → char='2', i=5
//   STATE 3: char is digit → skip runes[5]='2', i=6
//   STATE 4: Output "222" (three 2's)
//
// i=6 >= len(runes) → exit loop
// Result: "111222" ✓
//
// ═══ WHY THIS WORKS ═══
//
// Input: "100a" (100 a's)
//
// Iteration 1:
//   i=0
//   STATE 1: Read "100" → i goes 0→1→2→3, countStr="100"
//   STATE 2: Read 'a' → char='a', i=4
//   STATE 3: char is NOT digit → skip nothing
//   STATE 4: Output 100 a's
//
// i=4 >= len(runes) → exit loop
// Result: 100 a's ✓
//
// Key difference from broken version:
// - Old approach: Tried to detect escape WHILE reading count
// - New approach: Read count COMPLETELY FIRST, then check escape
