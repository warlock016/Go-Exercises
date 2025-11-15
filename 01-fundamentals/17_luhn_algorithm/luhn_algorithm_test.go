package luhnalgorithm

import "testing"

func TestIsValidLuhn(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		// Valid credit card numbers
		{"Valid Visa card", "4532015112830366", true},
		{"Valid Discover card", "6011111111111117", true},
		{"Valid MasterCard", "5425233430109903", true},
		{"Valid American Express", "374245455400126", true},
		{"Valid generic number", "1234567812345670", true},
		{"Valid all zeros", "0000000000000000", true},
		{"Valid single check digit", "18", true},

		// Invalid credit card numbers
		{"Invalid check digit wrong", "4532015112830360", false},
		{"Invalid number simple", "1234567812345678", false},
		{"Invalid all nines", "9999999999999999", false},
		{"Invalid single wrong digit", "4532015112830367", false},

		// Edge cases
		{"Empty string", "", false},
		{"Single digit", "0", true},
		{"Single digit non-zero", "5", false},
		{"Two digits invalid", "15", false},

		// With formatting (spaces, hyphens)
		{"Valid with hyphens", "4532-0151-1283-0366", true},
		{"Valid with spaces", "4532 0151 1283 0366", true},
		{"Valid mixed formatting", "4532-0151 1283-0366", true},
		{"Invalid with hyphens", "4532-0151-1283-0360", false},

		// Non-digit characters
		{"With letters should clean", "45X32Y015Z1128W30366", true},
		{"Only letters", "ABCDEF", false},

		// Additional test cases
		{"Valid short number", "8763", true},
		{"Invalid short number", "8764", false},
		{"Valid another pattern", "79927398713", true},
		{"Invalid transposed digits", "4532015112830636", false}, // 36 swapped to 63
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidLuhn(tt.input)
			if got != tt.want {
				t.Errorf("IsValidLuhn(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestGenerateCheckDigit(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		// Standard cases
		{"Visa card partial", "453201511283036", 6},
		{"Discover card partial", "601111111111111", 7},
		{"MasterCard partial", "542523343010990", 3},
		{"Amex partial", "37424545540012", 6},
		{"Generic number", "123456781234567", 0},

		// Edge cases
		{"Single digit 0", "0", 0},
		{"Single digit 1", "1", 8},
		{"Single digit 2", "2", 6},
		{"Single digit 5", "5", 5},
		{"Single digit 9", "9", 1},
		{"Empty string", "", 0},
		{"Two digits", "12", 5},

		// All same digits
		{"All ones", "111111111111111", 4},
		{"All fives", "555555555555555", 5},
		{"All nines", "999999999999999", 3},

		// With formatting (should be cleaned)
		{"With hyphens", "4532-0151-1283-036", 6},
		{"With spaces", "4532 0151 1283 036", 6},
		{"Mixed formatting", "453-2015 11283-036", 6},

		// Verify check digit works with IsValidLuhn
		{"Short valid sequence", "876", 3}, // Makes "8763" valid
		{"Another sequence", "7992739871", 3}, // Makes "79927398713" valid
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateCheckDigit(tt.input)
			if got != tt.want {
				t.Errorf("GenerateCheckDigit(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// TestGenerateCheckDigitIntegration verifies that generated check digits
// produce valid Luhn numbers when combined with the input
func TestGenerateCheckDigitIntegration(t *testing.T) {
	testCases := []string{
		"453201511283036",
		"601111111111111",
		"542523343010990",
		"123456781234567",
		"7992739871",
		"876",
		"12",
		"5",
	}

	for _, input := range testCases {
		t.Run(input, func(t *testing.T) {
			checkDigit := GenerateCheckDigit(input)
			fullNumber := input + string(rune('0'+checkDigit))

			if !IsValidLuhn(fullNumber) {
				t.Errorf("GenerateCheckDigit(%q) = %d, but %q is not valid according to IsValidLuhn",
					input, checkDigit, fullNumber)
			}
		})
	}
}

// TestLuhnAlgorithmProperties tests mathematical properties of the Luhn algorithm
func TestLuhnAlgorithmProperties(t *testing.T) {
	t.Run("DetectsSingleDigitError", func(t *testing.T) {
		valid := "4532015112830366"
		// Change one digit
		invalid := "4532015112830466" // Changed 3 to 4
		if IsValidLuhn(valid) != true {
			t.Error("Original should be valid")
		}
		if IsValidLuhn(invalid) != false {
			t.Error("Single digit change should be detected")
		}
	})

	t.Run("CheckDigitRange", func(t *testing.T) {
		// All check digits should be 0-9
		for i := 0; i < 100; i++ {
			input := "12345678901234567890"[:i%15+1]
			digit := GenerateCheckDigit(input)
			if digit < 0 || digit > 9 {
				t.Errorf("GenerateCheckDigit(%q) = %d, want 0-9", input, digit)
			}
		}
	})
}

func BenchmarkIsValidLuhn(b *testing.B) {
	testNumber := "4532015112830366"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsValidLuhn(testNumber)
	}
}

func BenchmarkGenerateCheckDigit(b *testing.B) {
	testNumber := "453201511283036"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GenerateCheckDigit(testNumber)
	}
}
