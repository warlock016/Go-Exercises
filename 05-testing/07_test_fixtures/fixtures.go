package fixtures

import (
	"fmt"
	"strings"
)

// Record represents a data record for reports
type Record struct {
	ID    int
	Name  string
	Value float64
}

// ProcessMarkdown converts simple markdown to HTML
// Supports: # headers, **bold**, *italic*
func ProcessMarkdown(input string) string {
	result := input

	// Headers
	result = strings.ReplaceAll(result, "# ", "<h1>")
	result = strings.ReplaceAll(result, "\n", "</h1>\n")

	// Bold
	result = strings.ReplaceAll(result, "**", "<b>")
	parts := strings.Split(result, "<b>")
	result = ""
	for i, part := range parts {
		if i > 0 && i%2 == 0 {
			result += "</b>"
		}
		result += part
		if i > 0 && i%2 == 1 {
			result += "<b>"
		}
	}
	result = strings.ReplaceAll(result, "<b>", "**")

	// Simple implementation
	result = strings.ReplaceAll(result, "**", "<strong>")
	result = strings.ReplaceAll(result, "**", "</strong>")

	return result
}

// GenerateReport creates a formatted report from records
func GenerateReport(records []Record) string {
	var sb strings.Builder

	sb.WriteString("REPORT\n")
	sb.WriteString("======\n\n")

	for _, r := range records {
		sb.WriteString(fmt.Sprintf("ID: %d\n", r.ID))
		sb.WriteString(fmt.Sprintf("Name: %s\n", r.Name))
		sb.WriteString(fmt.Sprintf("Value: %.2f\n", r.Value))
		sb.WriteString("\n")
	}

	return sb.String()
}
