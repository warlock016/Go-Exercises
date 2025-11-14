package regexmatcher

// Token represents a parsed pattern element
// Examples: {char: 'a', op: "*"} for "a*"
//           {char: 'b', op: "+"} for "b+"
//           {char: 'c', op: ""} for "c" (literal)
type Token struct {
	char rune
	op   string // "", "*", or "+"
}

// Match returns true if the text matches the pattern exactly
//
// Pattern syntax:
//   - 'a*' matches zero or more 'a's
//   - 'a+' matches one or more 'a's
//   - 'a'  matches exactly one 'a'
//
// Examples:
//   Match("a*b", "aaab")  → true
//   Match("a+b", "b")     → false (needs at least one 'a')
//   Match("ab*c", "ac")   → true (zero b's)
func Match(pattern string, text string) bool {
	// TODO(human): Implement regex matcher using finite automaton
	//
	// Algorithm overview:
	// 1. Parse pattern into tokens (handle *, +, literals)
	// 2. Process text character-by-character against tokens
	// 3. For each token type, apply appropriate state machine logic:
	//    - Literal: must match exactly one character
	//    - *: match zero or more (greedy - consume as many as possible)
	//    - +: match one or more (required first match, then greedy)
	// 4. Verify all text was consumed after processing all tokens
	//
	// Hints:
	// - Use parsePattern(pattern) helper function
	// - Track position in text with textIdx
	// - For *, use: for textIdx < len(text) && text[textIdx] == char { textIdx++ }
	// - For +, first check required match, then loop like *
	// - Return textIdx == len(text) at the end
	//
	// Debugging tips:
	// - Print tokens after parsing to verify structure
	// - Add log statements showing textIdx and current character
	// - Test each operator type in isolation first

	return false
}

// parsePattern converts a pattern string into a list of tokens
//
// Examples:
//   "a*b"  → [{char: 'a', op: "*"}, {char: 'b', op: ""}]
//   "a+b*" → [{char: 'a', op: "+"}, {char: 'b', op: "*"}]
//   "abc"  → [{char: 'a', op: ""}, {char: 'b', op: ""}, {char: 'c', op: ""}]
func parsePattern(pattern string) []Token {
	// TODO(human): Parse pattern into tokens
	//
	// Algorithm:
	// 1. Convert pattern to []rune for Unicode correctness
	// 2. Loop through runes with index i
	// 3. Look ahead: if runes[i+1] is '*' or '+', create token with operator
	//    - Increment i by 2 (skip char and operator)
	// 4. Otherwise, create token without operator (literal)
	//    - Increment i by 1
	// 5. Return list of tokens
	//
	// Example trace for "a*b":
	// i=0: runes[0]='a', runes[1]='*' → Token{char: 'a', op: "*"}, i=2
	// i=2: runes[2]='b', no lookahead → Token{char: 'b', op: ""}, i=3
	// Done
	//
	// Bounds checking:
	// - Before checking runes[i+1], ensure i+1 < len(runes)
	//
	// Hint: Use []Token{} to initialize empty slice
	//       Use append(tokens, Token{...}) to add tokens

	tokens := []Token{}

	// YOUR CODE HERE

	return tokens
}
