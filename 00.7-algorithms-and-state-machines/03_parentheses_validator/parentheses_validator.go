package parenthesesvalidator

// IsValid checks if parentheses in a string are balanced and properly nested
func IsValid(s string) bool {
	// TODO(human): Implement the validator using a stack
	//
	// Algorithm:
	// 1. Create a stack (slice of runes) to track opening parentheses
	// 2. Loop through each character:
	//    a. If '(' → push onto stack
	//    b. If ')' → check if stack is empty:
	//       - If empty: return false (no matching opening)
	//       - If not empty: pop from stack (matched a pair)
	// 3. After loop: return whether stack is empty
	//    - Empty = all pairs matched ✓
	//    - Not empty = some openings never closed ✗
	//
	// Stack operations in Go:
	// - Push: stack = append(stack, item)
	// - Pop: stack = stack[:len(stack)-1]
	// - Check empty: len(stack) == 0
	//
	// Hint: Start with `stack := []rune{}`

	return false // TODO(human): Replace with your implementation
}
