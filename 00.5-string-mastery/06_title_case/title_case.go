package titlecase

import (
	"strings"
	"unicode"
)

// ToTitleCase converts a string to title case where the first letter of each
// word is capitalized and all other letters are lowercase.
//
// A "word" is defined as a sequence of letters. Any non-letter character
// (space, punctuation, digit, etc.) acts as a word boundary.
//
// Examples:
//
//	ToTitleCase("hello world") → "Hello World"
//	ToTitleCase("STOP SHOUTING") → "Stop Shouting"
//	ToTitleCase("it's nice") → "It's Nice"
//	ToTitleCase("hello-world") → "Hello-World"
func ToTitleCase(s string) string {
	// TODO(human): Implement title case conversion
	//
	// APPROACH:
	// 1. Use a boolean flag to track if we're at the start of a word
	// 2. Iterate through each rune in the string
	// 3. For letters:
	//    - If at word start: capitalize (unicode.ToUpper)
	//    - Otherwise: lowercase (unicode.ToLower)
	// 4. For non-letters:
	//    - Keep unchanged
	//    - Set flag to true (next letter starts a new word)
	//
	// TIPS:
	// - Use strings.Builder for efficient string construction
	// - Use unicode.IsLetter(r) to check if rune is a letter
	// - Use unicode.ToUpper(r) and unicode.ToLower(r) for case conversion
	// - Remember: 'range' over a string gives you runes, not bytes!
	//
	// PSEUDOCODE:
	//   var builder strings.Builder
	//   atWordStart := true  // First character is at word start
	//
	//   for _, r := range s {
	//       if unicode.IsLetter(r) {
	//           if atWordStart {
	//               builder.WriteRune(unicode.ToUpper(r))
	//               atWordStart = false
	//           } else {
	//               builder.WriteRune(unicode.ToLower(r))
	//           }
	//       } else {
	//           builder.WriteRune(r)
	//           atWordStart = true
	//       }
	//   }
	//
	//   return builder.String()

	_ = unicode.IsLetter // Hint: you'll need this function
	_ = unicode.ToUpper  // Hint: and this one
	_ = unicode.ToLower  // Hint: and this one

	var builder strings.Builder
	// Your code here

	// col := []rune(s) // [h,e,l,l,o, ,w,o,r,l,d]
	// var pos int = 0       // keep count of IsLetter sequences, resets to 0 when a non-letter character is encountered
	flag := false // flag IsLetter

	for _, v := range s {
		if !flag && unicode.IsLetter(v) {
			flag = true
			builder.WriteRune(unicode.ToUpper(v))
		} else if unicode.IsLetter(v) {
			builder.WriteRune(unicode.ToLower(v))
		} else {
			flag = false
			builder.WriteRune(v)
		}
	}

	return builder.String()
}
