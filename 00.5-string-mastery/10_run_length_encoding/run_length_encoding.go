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
//   Encode("aaabbc") → "3a2b1c"
//   Encode("hello") → "1h1e2l1o"
//   Encode("aaa") → "3a"
//   Encode("") → ""
func Encode(s string) string {
	// TODO(human): Implement run-length encoding
	//
	// ALGORITHM:
	// 1. Handle empty string edge case
	// 2. Convert to []rune (for Unicode correctness)
	// 3. Track current character and its count
	// 4. Iterate through runes:
	//    - If same as current: increment count
	//    - If different: write count+char, start new run
	// 5. After loop: write the last run
	//
	// DETAILED STEPS:
	//
	// STEP 1: Edge case
	//   if s == "" {
	//       return ""
	//   }
	//
	// STEP 2: Convert to runes
	//   runes := []rune(s)
	//   var builder strings.Builder
	//
	// STEP 3: Initialize tracking
	//   count := 1
	//   currentRune := runes[0]  // First character
	//
	// STEP 4: Iterate and count runs
	//   for i := 1; i < len(runes); i++ {
	//       if runes[i] == currentRune {
	//           // Same character - extend current run
	//           count++
	//       } else {
	//           // Different character - write current run
	//           builder.WriteString(strconv.Itoa(count))
	//           builder.WriteRune(currentRune)
	//
	//           // Start new run
	//           currentRune = runes[i]
	//           count = 1
	//       }
	//   }
	//
	// STEP 5: Write last run (IMPORTANT - don't forget!)
	//   builder.WriteString(strconv.Itoa(count))
	//   builder.WriteRune(currentRune)
	//
	// STEP 6: Return result
	//   return builder.String()
	//
	// EXAMPLE TRACE:
	// Input: "aaabbc"
	// Runes: ['a', 'a', 'a', 'b', 'b', 'c']
	//
	// i=0: currentRune='a', count=1
	// i=1: runes[1]='a' == 'a' → count=2
	// i=2: runes[2]='a' == 'a' → count=3
	// i=3: runes[3]='b' != 'a' → write "3a", currentRune='b', count=1
	// i=4: runes[4]='b' == 'b' → count=2
	// i=5: runes[5]='c' != 'b' → write "2b", currentRune='c', count=1
	// End: write "1c"
	// Result: "3a2b1c"

	_ = strconv.Itoa     // Hint: convert count to string
	_ = strings.Builder{} // Hint: build result efficiently

	// Your code here
	return ""
}

// Decode decompresses a run-length encoded string back to original.
// The encoded format is: count + character (e.g., "3a2b1c" → "aaabbc")
//
// Examples:
//   Decode("3a2b1c") → "aaabbc"
//   Decode("1h1e2l1o") → "hello"
//   Decode("") → ""
func Decode(s string) string {
	// TODO(human): Implement run-length decoding
	//
	// ALGORITHM:
	// 1. Handle empty string edge case
	// 2. Convert to []rune (for Unicode correctness)
	// 3. Parse the encoded string:
	//    - Read digits to get count
	//    - Read next character
	//    - Write character 'count' times
	// 4. Repeat until end of string
	//
	// DETAILED STEPS:
	//
	// STEP 1: Edge case
	//   if s == "" {
	//       return ""
	//   }
	//
	// STEP 2: Setup
	//   runes := []rune(s)
	//   var builder strings.Builder
	//   i := 0  // Current position
	//
	// STEP 3: Parse loop
	//   for i < len(runes) {
	//       // Read the count (one or more digits)
	//       countStr := ""
	//       for i < len(runes) && unicode.IsDigit(runes[i]) {
	//           countStr += string(runes[i])
	//           i++
	//       }
	//
	//       // Convert to integer
	//       count, _ := strconv.Atoi(countStr)
	//
	//       // Read the character to repeat
	//       if i < len(runes) {
	//           char := runes[i]
	//           i++
	//
	//           // Write 'count' copies of 'char'
	//           for j := 0; j < count; j++ {
	//               builder.WriteRune(char)
	//           }
	//       }
	//   }
	//
	// STEP 4: Return result
	//   return builder.String()
	//
	// EXAMPLE TRACE:
	// Input: "3a2b1c"
	// Runes: ['3', 'a', '2', 'b', '1', 'c']
	//
	// i=0: Read '3' → countStr="3", count=3
	//      Read 'a' → write 'a' 3 times → "aaa"
	// i=2: Read '2' → countStr="2", count=2
	//      Read 'b' → write 'b' 2 times → "aaabb"
	// i=4: Read '1' → countStr="1", count=1
	//      Read 'c' → write 'c' 1 time → "aaabbc"
	// Result: "aaabbc"
	//
	// CHALLENGE: Multi-digit counts
	// Input: "13a"
	// i=0: Read '1', '3' → countStr="13", count=13
	//      Read 'a' → write 'a' 13 times
	//
	// This is why we use a loop to read ALL consecutive digits!

	_ = unicode.IsDigit  // Hint: check if rune is a digit
	_ = strconv.Atoi     // Hint: convert string to int

	// Your code here
	return ""
}
