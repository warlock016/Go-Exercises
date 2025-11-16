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
	var builder strings.Builder

	_ = builder      // Using builder
	_ = strconv.Itoa // For converting int to string

	if n <= 0 {
		return ""
	}

	for i := 1; i <= n; i++ {
		// fmt.Println(i)
		for j := 1; j <= i; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		builder.WriteString("\n")
	}

	return builder.String()
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
	var builder strings.Builder

	_ = builder      // Using builder
	_ = strconv.Itoa // For converting int to string

	if n <= 0 {
		return ""
	}

	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		builder.WriteString("\n")
	}

	return builder.String()
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
	var builder strings.Builder

	_ = builder      // Using builder
	_ = strconv.Itoa // For converting int to string

	if n < 0 {
		return ""
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			builder.WriteString(strconv.Itoa(i * j))
			if j < n {
				builder.WriteString(" ")
			} else {
				builder.WriteString("\n")
			}
		}

	}

	return builder.String()
}
