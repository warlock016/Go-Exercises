package patternprinter

import (
	"strconv"
	"strings"
)

// PrintSquare generates an n×n square pattern of asterisks.
// Each row contains n asterisks followed by a newline.
// Returns an empty string if n <= 0.
//
// Example:
//
//	PrintSquare(3) returns "***\n***\n***\n"
func PrintSquare(n int) string {
	// TODO(human): Implement PrintSquare
	var square strings.Builder

	if n <= 0 {
		return ""
	}

	for range n {
		square.WriteString(strings.Repeat("*", n))
		square.WriteString("\n")
	}

	return square.String()
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

	var triangle strings.Builder

	if n <= 0 {
		return ""
	}

	for i := 1; i <= n; i++ {
		triangle.WriteString(strings.Repeat("*", i))
		triangle.WriteString("\n")
	}

	return triangle.String()
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
	var pyramid strings.Builder

	if n <= 0 {
		return ""
	}

	for i := 1; i <= n; i++ {
		pyramid.WriteString(strings.Repeat(" ", n-i))
		pyramid.WriteString(strings.Repeat("*", 2*i-1))
		pyramid.WriteString("\n")
	}

	return pyramid.String()
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
	var square strings.Builder

	if n <= 0 {
		return ""
	}

	for i := 1; i <= n; i++ {
		square.WriteString(strings.Repeat(strconv.Itoa(i), n))
		square.WriteString("\n")
	}

	return square.String()
}
