package patternprinter

// PrintSquare generates an n×n square pattern of asterisks.
// Each row contains n asterisks followed by a newline.
// Returns an empty string if n <= 0.
//
// Example:
//
//	PrintSquare(3) returns "***\n***\n***\n"
func PrintSquare(n int) string {
	// TODO(human): Implement PrintSquare
	// 1. Handle edge case: if n <= 0, return ""
	// 2. Create a strings.Builder for efficient string building
	// 3. Use outer loop for rows (iterate n times)
	// 4. Use inner loop for columns (iterate n times, write "*")
	// 5. After each row's inner loop, write "\n"
	// 6. Return builder.String()
	return ""
}

// PrintTriangle generates a right triangle pattern with n rows.
// Row i has i asterisks. Each row ends with a newline.
// Returns an empty string if n <= 0.
//
// Example:
//
//	PrintTriangle(4) returns "*\n**\n***\n****\n"
func PrintTriangle(n int) string {
	// TODO(human): Implement PrintTriangle
	// 1. Handle edge case: if n <= 0, return ""
	// 2. Create a strings.Builder
	// 3. Outer loop: for row from 0 to n-1
	// 4. Inner loop: iterate (row + 1) times, write "*"
	// 5. After inner loop, write "\n"
	// 6. Return builder.String()
	return ""
}

// PrintPyramid generates a centered pyramid pattern with n rows.
// Row i (1-indexed) has (2*i - 1) asterisks centered with leading spaces.
// Returns an empty string if n <= 0.
//
// Example:
//
//	PrintPyramid(3) returns "  *\n ***\n*****\n"
func PrintPyramid(n int) string {
	// TODO(human): Implement PrintPyramid
	// 1. Handle edge case: if n <= 0, return ""
	// 2. Create a strings.Builder
	// 3. Outer loop: for row from 0 to n-1
	// 4. Write leading spaces: loop (n - row - 1) times, write " "
	// 5. Write asterisks: loop (2*row + 1) times, write "*"
	// 6. Write "\n"
	// 7. Return builder.String()
	return ""
}

// PrintNumberSquare generates an n×n grid where each row displays its row number.
// Row 1 contains n "1"s, row 2 contains n "2"s, etc.
// Returns an empty string if n <= 0.
//
// Example:
//
//	PrintNumberSquare(3) returns "111\n222\n333\n"
func PrintNumberSquare(n int) string {
	// TODO(human): Implement PrintNumberSquare
	// 1. Handle edge case: if n <= 0, return ""
	// 2. Create a strings.Builder
	// 3. Outer loop: for row from 0 to n-1
	// 4. Convert row number to string: strconv.Itoa(row + 1)
	// 5. Inner loop: iterate n times, write the digit string
	// 6. After inner loop, write "\n"
	// 7. Return builder.String()
	return ""
}
