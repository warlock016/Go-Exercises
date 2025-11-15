package asciiart

// DrawBox creates a rectangular box using the specified character.
// The box will have 'width' characters horizontally and 'height' rows vertically.
// If width or height is less than 2, returns an empty string.
//
// Example: DrawBox(5, 3, '*') produces:
//
//	*****
//	*   *
//	*****
func DrawBox(width, height int, char rune) string {
	// TODO(human): Implement box drawing
	// 1. Check if width < 2 or height < 2, return "" if true
	// 2. Import and create a strings.Builder
	// 3. First row: repeat char width times, add newline
	// 4. Middle rows (height-2 times): char + (width-2) spaces + char + newline
	// 5. Last row: same as first row
	// 6. Return builder.String()
	//
	// Hint: You'll need to import "strings" package
	// Hint: Use strings.Repeat(string(char), count) for repeating characters
	// Hint: Use builder.WriteString() and builder.WriteRune()
	return ""
}

// DrawDiamond creates a diamond pattern with n rows in the top half.
// The diamond uses '*' characters and is centered with spaces.
// If n is less than 1, returns an empty string.
//
// Example: DrawDiamond(3) produces:
//
//	  *
//	 ***
//	*****
//	 ***
//	  *
func DrawDiamond(n int) string {
	// TODO(human): Implement diamond drawing
	// 1. Check if n < 1, return "" if true
	// 2. Create a strings.Builder
	// 3. Top half (i from 0 to n-1):
	//    - spaces = n - i - 1
	//    - stars = 2*i + 1
	//    - Write: spaces + stars + newline
	// 4. Bottom half (i from n-2 down to 0):
	//    - Same calculation as top half
	//    - This creates the mirror effect
	// 5. Return builder.String()
	//
	// Hint: Use strings.Repeat(" ", spaces) for leading spaces
	// Hint: Use strings.Repeat("*", stars) for the stars
	return ""
}

// DrawChessboard creates an n×n chessboard pattern using filled (█) and empty (░) blocks.
// The pattern alternates based on position: if (row+col) is even, use █, otherwise use ░.
// If n is less than 1, returns an empty string.
//
// Example: DrawChessboard(4) produces:
//
//	█░█░
//	░█░█
//	█░█░
//	░█░█
func DrawChessboard(n int) string {
	// TODO(human): Implement chessboard drawing
	// 1. Check if n < 1, return "" if true
	// 2. Create a strings.Builder
	// 3. Outer loop for rows (i from 0 to n-1):
	//    - Inner loop for columns (j from 0 to n-1):
	//      - If (i+j)%2 == 0: write '█'
	//      - Else: write '░'
	//    - After inner loop: write newline
	// 4. Return builder.String()
	//
	// Hint: Use builder.WriteRune() to write individual runes
	// Hint: The modulo operator % checks if a number is even/odd
	return ""
}
