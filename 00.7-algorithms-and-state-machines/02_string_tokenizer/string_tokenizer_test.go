package stringtokenizer

import (
	"reflect"
	"testing"
)

func TestTokenizeEmpty(t *testing.T) {
	result := Tokenize("")
	if len(result) != 0 {
		t.Errorf("Tokenize(\"\") = %v, want []", result)
	}
}

func TestTokenizeSingleWord(t *testing.T) {
	result := Tokenize("hello")
	expected := []string{"hello"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Tokenize(\"hello\") = %v, want %v", result, expected)
	}
}

func TestTokenizeMultipleWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Two words",
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "Three words",
			input:    "hello world test",
			expected: []string{"hello", "world", "test"},
		},
		{
			name:     "Multiple spaces between words",
			input:    "hello  world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "Many spaces between words",
			input:    "hello    world    test",
			expected: []string{"hello", "world", "test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Tokenize(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Tokenize(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTokenizeLeadingTrailingSpaces(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Leading space",
			input:    " hello",
			expected: []string{"hello"},
		},
		{
			name:     "Trailing space",
			input:    "hello ",
			expected: []string{"hello"},
		},
		{
			name:     "Leading and trailing spaces",
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			name:     "Leading and trailing with multiple words",
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Tokenize(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Tokenize(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTokenizeOnlyWhitespace(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Single space", " "},
		{"Multiple spaces", "   "},
		{"Tabs", "\t\t"},
		{"Mixed whitespace", " \t \n "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Tokenize(tt.input)
			if len(result) != 0 {
				t.Errorf("Tokenize(%q) = %v, want []", tt.input, result)
			}
		})
	}
}

func TestTokenizeWithTabs(t *testing.T) {
	result := Tokenize("hello\tworld")
	expected := []string{"hello", "world"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Tokenize(\"hello\\tworld\") = %v, want %v", result, expected)
	}
}

func TestTokenizeWithNewlines(t *testing.T) {
	result := Tokenize("hello\nworld\ntest")
	expected := []string{"hello", "world", "test"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Tokenize(\"hello\\nworld\\ntest\") = %v, want %v", result, expected)
	}
}
