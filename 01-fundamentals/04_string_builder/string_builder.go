package stringbuilder

import (
	"fmt"
	"strings"
)

// Concat joins any number of strings efficiently.
// Returns an empty string if no arguments are provided.
//
// Example:
//
//	Concat("Hello", " ", "World") => "Hello World"
//	Concat("Go", "lang") => "Golang"
func Concat(parts ...string) string {
	// TODO(human): Implement efficient string concatenation

	var output strings.Builder
	var totalChars int = 0

	for _, part := range parts {
		for range part {
			totalChars++
		}
	}

	output.Grow(totalChars)

	for _, part := range parts {
		for _, r := range part {
			output.WriteRune(r)
		}
	}

	return output.String()
}

// BuildGreeting creates a greeting string with optional title.
// If title is empty, format is "Hello, Name!"
// If title is provided, format is "Hello, Title Name!"
//
// Example:
//
//	BuildGreeting("Smith", "Dr.") => "Hello, Dr. Smith!"
//	BuildGreeting("Alice", "") => "Hello, Alice!"
func BuildGreeting(name, title string) string {
	// TODO(human): Implement greeting builder

	var output strings.Builder

	output.WriteString("Hello,")
	if title != "" {
		output.WriteString(" " + title)
	}
	output.WriteString(" " + name + "!")

	return output.String() //fmt.Sprintf("Hello, %s %s!", title, name)
}

// RepeatWithSeparator repeats a string n times with a separator between occurrences.
// Returns empty string if n <= 0.
// Does not add separator after the last occurrence.
//
// Example:
//
//	RepeatWithSeparator("Go", 3, "-") => "Go-Go-Go"
//	RepeatWithSeparator("*", 5, "") => "*****"
//	RepeatWithSeparator("ha", 1, ",") => "ha"
func RepeatWithSeparator(s string, n int, sep string) string {
	// TODO(human): Implement repeat with separator

	var output strings.Builder

	for i := 0; i < n; i++ {
		output.WriteString(s)

		if i+1 < n {
			output.WriteString(sep)
		}
	}
	return output.String()
}

// FormatTable creates a simple two-row table from headers and values.
// First row contains headers joined by " | "
// Second row contains values joined by " | "
// Rows are separated by a newline.
//
// Example:
//
//	headers := []string{"Name", "Age"}
//	values := []string{"Alice", "30"}
//	FormatTable(headers, values) =>
//	  "Name | Age
//	   Alice | 30"
func FormatTable(headers, values []string) string {
	// TODO(human): Implement table formatter
	// headers and values are both slices

	var output strings.Builder

	if len(headers) != len(values) || len(headers) == 0 || len(values) == 0 { // edge cases where headers len does not match values len or headers len, values len == 0
		fmt.Println("invalid inputs: headers len does not match values len")
		return ""
	}

	for i := range headers {
		if i+1 < len(headers) {
			output.WriteString(headers[i] + " | ")
		} else if i+1 == len(headers) {
			output.WriteString(headers[i] + "\n")
		}
	}

	for i := range values {
		if i+1 < len(values) {
			output.WriteString(values[i] + " | ")
		} else if i+1 == len(values) {
			output.WriteString(values[i])
		}
	}

	return output.String()
}
