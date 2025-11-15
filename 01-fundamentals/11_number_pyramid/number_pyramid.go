package numberpyramid

import (
	"strconv"
	"strings"
)

// NumberTriangle generates a number triangle pattern.
// Row 1: "1\n"
// Row 2: "12\n"
// Row 3: "123\n"
// ...
//
// Example:
//
//	NumberTriangle(3) returns:
//	1
//	12
//	123
func NumberTriangle(n int) string {
	// TODO(human): Build a triangle where each row contains numbers from 1 to row number
	// Hint: Use nested loops - outer for rows (1 to n), inner for numbers (1 to current row)
	// Hint: Use strings.Builder for efficient string construction
	// Hint: Don't forget to add "\n" at the end of each row
	var builder strings.Builder

	_ = builder     // Using builder
	_ = strconv.Itoa // For converting int to string

	return ""
}

// ReversePyramid generates a reverse pyramid pattern.
// Row 1: "12345\n" (for n=5)
// Row 2: "1234\n"
// Row 3: "123\n"
// Row 4: "12\n"
// Row 5: "1\n"
//
// Example:
//
//	ReversePyramid(4) returns:
//	1234
//	123
//	12
//	1
func ReversePyramid(n int) string {
	// TODO(human): Build a reverse pyramid starting from n numbers down to 1
	// Hint: Outer loop could go from n down to 1 (or 1 to n with different inner loop)
	// Hint: Each row i prints numbers from 1 to (n - row + 1) or similar logic
	// Hint: Think about the relationship between row number and how many numbers to print
	var builder strings.Builder

	_ = builder     // Using builder
	_ = strconv.Itoa // For converting int to string

	return ""
}

// MultiplicationTable generates an n×n multiplication table.
// Each row contains products of row_number × column_number.
// Numbers are separated by single spaces, no trailing space at end of row.
//
// Example:
//
//	MultiplicationTable(3) returns:
//	1 2 3
//	2 4 6
//	3 6 9
func MultiplicationTable(n int) string {
	// TODO(human): Build an n×n multiplication table
	// Hint: Nested loops - outer for rows (1 to n), inner for columns (1 to n)
	// Hint: Calculate product = row * col
	// Hint: Add space between numbers EXCEPT after the last number in each row
	// Hint: Check if col < n before adding space
	var builder strings.Builder

	_ = builder     // Using builder
	_ = strconv.Itoa // For converting int to string

	return ""
}
