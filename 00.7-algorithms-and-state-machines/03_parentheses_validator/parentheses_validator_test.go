package parenthesesvalidator

import "testing"

func TestIsValidEmpty(t *testing.T) {
	if !IsValid("") {
		t.Error("Empty string should be valid")
	}
}

func TestIsValidSinglePair(t *testing.T) {
	if !IsValid("()") {
		t.Error("Single pair '()' should be valid")
	}
}

func TestIsValidNested(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Two levels", "(())"},
		{"Three levels", "((()))"},
		{"Four levels", "(((())))"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !IsValid(tt.input) {
				t.Errorf("IsValid(%q) = false, want true", tt.input)
			}
		})
	}
}

func TestIsValidMultiplePairs(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Two pairs", "()()"},
		{"Three pairs", "()()()"},
		{"Mixed", "(())()"},
		{"Complex", "()(())()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !IsValid(tt.input) {
				t.Errorf("IsValid(%q) = false, want true", tt.input)
			}
		})
	}
}

func TestIsValidUnbalanced(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Missing closing", "(("},
		{"Missing closing nested", "(()"},
		{"Extra closing", "())"},
		{"Only opening", "("},
		{"Only closing", ")"},
		{"Multiple only opening", "(((("},
		{"Multiple only closing", "))))"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if IsValid(tt.input) {
				t.Errorf("IsValid(%q) = true, want false", tt.input)
			}
		})
	}
}

func TestIsValidWrongOrder(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Simple reverse", ")("},
		{"Complex reverse", ")()("},
		{"Closing before opening", ")()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if IsValid(tt.input) {
				t.Errorf("IsValid(%q) = true, want false", tt.input)
			}
		})
	}
}

func TestIsValidWithOtherCharacters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid with letters", "a(b)c", true},
		{"Valid with numbers", "1(2)3", true},
		{"Invalid with letters", "a(b(c", false},
		{"Complex valid", "(a(b)c)", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValid(tt.input)
			if result != tt.expected {
				t.Errorf("IsValid(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
