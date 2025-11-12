package builder

import (
	"strconv"
	"strings"
)

// JoinWords joins a slice of words with the given separator
// Similar to strings.Join, but you're implementing it!
func JoinWords(words []string, separator string) string {
	// TODO(human): Use strings.Builder to join words with separator
	// Hint: Handle empty slice, add first word, then add remaining with separator
	var b strings.Builder

	for i, word := range words {
		b.WriteString(word)
		if i < len(words)-1 {
			b.WriteString(separator)
		}
	}

	return b.String()
}

// Repeat returns a string consisting of s repeated n times
// Similar to strings.Repeat, but you're implementing it!
func Repeat(s string, n int) string {
	// TODO(human): Use strings.Builder to repeat the string n times
	// Hint: Simple loop that WriteString n times

	var b strings.Builder

	for range n {
		b.WriteString(s)
	}

	return b.String()
}

// BuildList creates a numbered list from the items
// Format: "1. Item\n2. Item\n3. Item" (no trailing newline)
func BuildList(items []string) string {
	// TODO(human): Use strings.Builder to create numbered list
	// Hint: Use strconv.Itoa to convert numbers to strings
	// Don't add newline after the last item

	var b strings.Builder

	for i, r := range items {
		b.WriteString(strconv.Itoa(i + 1))
		b.WriteString(". ")
		b.WriteString(r)
		if i < len(items)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

// BuildCSV converts a 2D slice into CSV format
// Each row is on a new line, values separated by commas
// No trailing newline after last row
func BuildCSV(rows [][]string) string {
	// TODO(human): Use strings.Builder to create CSV
	// Hint: Iterate over rows, join each row with commas, separate rows with newlines
	var b strings.Builder

	for i, row := range rows {
		for j, col := range row {
			b.WriteString(col)
			if j < len(row)-1 {
				b.WriteString(",")
			}
		}
		if i < len(rows)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}
