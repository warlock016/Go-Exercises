package csvparser

import (
	"reflect"
	"testing"
)

func TestParseCSVSimple(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Three fields",
			input:    "Alice,Bob,Charlie",
			expected: []string{"Alice", "Bob", "Charlie"},
		},
		{
			name:     "Numbers",
			input:    "1,2,3",
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "Mixed",
			input:    "Alice,30,Engineer",
			expected: []string{"Alice", "30", "Engineer"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseCSV(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseCSV(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseCSVQuoted(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "All quoted",
			input:    `"Alice","Bob","Charlie"`,
			expected: []string{"Alice", "Bob", "Charlie"},
		},
		{
			name:     "Quoted with comma inside",
			input:    `Alice,"New York, NY",30`,
			expected: []string{"Alice", "New York, NY", "30"},
		},
		{
			name:     "Multiple commas inside quotes",
			input:    `Name,"City, State, Country",Age`,
			expected: []string{"Name", "City, State, Country", "Age"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseCSV(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseCSV(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseCSVDoubledQuotes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Simple doubled quote",
			input:    `"Say ""hello"""`,
			expected: []string{`Say "hello"`},
		},
		{
			name:     "Doubled quote with comma",
			input:    `"Say ""hello""",world`,
			expected: []string{`Say "hello"`, "world"},
		},
		{
			name:     "Multiple doubled quotes",
			input:    `"""Hello"" and ""Goodbye"""`,
			expected: []string{`"Hello" and "Goodbye"`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseCSV(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseCSV(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseCSVEmptyFields(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "One empty field at start",
			input:    ",Bob,Charlie",
			expected: []string{"", "Bob", "Charlie"},
		},
		{
			name:     "One empty field in middle",
			input:    "Alice,,Charlie",
			expected: []string{"Alice", "", "Charlie"},
		},
		{
			name:     "One empty field at end",
			input:    "Alice,Bob,",
			expected: []string{"Alice", "Bob", ""},
		},
		{
			name:     "All empty fields",
			input:    ",,",
			expected: []string{"", "", ""},
		},
		{
			name:     "Empty quoted field",
			input:    `Alice,"",Charlie`,
			expected: []string{"Alice", "", "Charlie"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseCSV(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseCSV(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseCSVEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: []string{""},
		},
		{
			name:     "Single field",
			input:    "Alice",
			expected: []string{"Alice"},
		},
		{
			name:     "Single quoted field",
			input:    `"Alice"`,
			expected: []string{"Alice"},
		},
		{
			name:     "Only quotes",
			input:    `""`,
			expected: []string{""},
		},
		{
			name:     "Mixed quoted and unquoted",
			input:    `Alice,"Bob",Charlie,"Dave"`,
			expected: []string{"Alice", "Bob", "Charlie", "Dave"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseCSV(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseCSV(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
