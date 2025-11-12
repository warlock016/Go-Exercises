package types

import "testing"

func TestIsLetter(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Lowercase letter", 'a', true},
		{"Uppercase letter", 'Z', true},
		{"Digit", '5', false},
		{"Space", ' ', false},
		{"Punctuation", '!', false},
		{"Chinese character", '世', true},
		{"Accented letter", 'é', true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsLetter(tt.r)
			if got != tt.want {
				t.Errorf("IsLetter(%q) = %v, want %v", tt.r, got, tt.want)
			}
		})
	}
}

func TestIsDigit(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Zero", '0', true},
		{"Five", '5', true},
		{"Nine", '9', true},
		{"Letter", 'a', false},
		{"Space", ' ', false},
		{"Chinese number", '三', false}, // Chinese word for "three", not a digit
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsDigit(tt.r)
			if got != tt.want {
				t.Errorf("IsDigit(%q) = %v, want %v", tt.r, got, tt.want)
			}
		})
	}
}

func TestCountLetters(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Only letters", "hello", 5},
		{"Letters and digits", "Hello123", 5},
		{"Empty string", "", 0},
		{"Only digits", "12345", 0},
		{"With spaces", "Hello World", 10},
		{"With punctuation", "Hello, World!", 10},
		{"Accented letters", "café", 4},
		{"Chinese characters", "Hello世界", 7}, // 5 + 2
		{"Mixed", "Hello, 世界! 123", 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountLetters(tt.input)
			if got != tt.want {
				t.Errorf("CountLetters(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCountDigits(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Only digits", "12345", 5},
		{"Letters and digits", "Hello123", 3},
		{"Empty string", "", 0},
		{"Only letters", "hello", 0},
		{"With spaces", "Year 2024", 4},
		{"Phone number", "555-1234", 7},
		{"Mixed", "Room 101, Floor 2", 4}, // 1, 0, 1, 2
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountDigits(tt.input)
			if got != tt.want {
				t.Errorf("CountDigits(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestFilterLettersOnly(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Only letters", "hello", "hello"},
		{"Letters and digits", "Hello123", "Hello"},
		{"Empty string", "", ""},
		{"Only digits", "12345", ""},
		{"With spaces", "Hello World", "HelloWorld"},
		{"With punctuation", "Hello, World!", "HelloWorld"},
		{"Accented letters", "café!", "café"},
		{"Chinese characters", "Hello, 世界!", "Hello世界"},
		{"Complex mixed", "User@123.com", "Usercom"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterLettersOnly(tt.input)
			if got != tt.want {
				t.Errorf("FilterLettersOnly(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// Demonstrates the power of unicode package for international text
func TestInternationalCharacters(t *testing.T) {
	testCases := []struct {
		text    string
		letters int
		digits  int
	}{
		{"Hello", 5, 0},                    // English
		{"Привет", 6, 0},                   // Russian
		{"こんにちは", 5, 0},                    // Japanese
		{"مرحبا", 5, 0},                     // Arabic
		{"café", 4, 0},                     // French
		{"Straße", 6, 0},                   // German
		{"Room 101", 4, 3},                 // Mixed
		{"用户123", 2, 3},                    // Chinese + digits
	}

	for _, tc := range testCases {
		t.Run(tc.text, func(t *testing.T) {
			gotLetters := CountLetters(tc.text)
			gotDigits := CountDigits(tc.text)

			if gotLetters != tc.letters {
				t.Errorf("%q: CountLetters = %v, want %v", tc.text, gotLetters, tc.letters)
			}
			if gotDigits != tc.digits {
				t.Errorf("%q: CountDigits = %v, want %v", tc.text, gotDigits, tc.digits)
			}
		})
	}
}
