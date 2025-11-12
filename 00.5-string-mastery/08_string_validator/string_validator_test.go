package stringvalidator

import "testing"

// Test IsValidEmail function
func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		// Valid emails
		{
			name:  "Simple valid email",
			input: "user@example.com",
			want:  true,
		},
		{
			name:  "Email with dots in local",
			input: "first.last@example.com",
			want:  true,
		},
		{
			name:  "Email with subdomain",
			input: "user@mail.example.com",
			want:  true,
		},
		{
			name:  "Short email",
			input: "a@b.c",
			want:  true,
		},
		{
			name:  "Email with numbers",
			input: "user123@example123.com",
			want:  true,
		},

		// Invalid emails
		{
			name:  "No @ symbol",
			input: "invalid.email.com",
			want:  false,
		},
		{
			name:  "Multiple @ symbols",
			input: "user@@example.com",
			want:  false,
		},
		{
			name:  "No local part",
			input: "@example.com",
			want:  false,
		},
		{
			name:  "No domain part",
			input: "user@",
			want:  false,
		},
		{
			name:  "Domain without dot",
			input: "user@example",
			want:  false,
		},
		{
			name:  "Domain ends with dot",
			input: "user@example.",
			want:  false,
		},
		{
			name:  "Domain ends with .com.",
			input: "user@example.com.",
			want:  false,
		},
		{
			name:  "Empty string",
			input: "",
			want:  false,
		},
		{
			name:  "Just @",
			input: "@",
			want:  false,
		},
		{
			name:  "Multiple @ in middle",
			input: "us@er@example.com",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidEmail(tt.input)
			if got != tt.want {
				t.Errorf("IsValidEmail(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

// Test IsValidPassword function
func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		// Valid passwords
		{
			name:  "Valid with all requirements",
			input: "Strong123",
			want:  true,
		},
		{
			name:  "Valid longer password",
			input: "MySecurePassword1",
			want:  true,
		},
		{
			name:  "Valid exactly 8 chars",
			input: "Valid123",
			want:  true,
		},
		{
			name:  "Valid with special chars",
			input: "P@ssw0rd!",
			want:  true,
		},

		// Invalid passwords
		{
			name:  "Too short",
			input: "Short1",
			want:  false,
		},
		{
			name:  "No uppercase",
			input: "lowercase123",
			want:  false,
		},
		{
			name:  "No lowercase",
			input: "UPPERCASE123",
			want:  false,
		},
		{
			name:  "No digit",
			input: "NoDigitsHere",
			want:  false,
		},
		{
			name:  "Only 7 chars (boundary)",
			input: "Valid12",
			want:  false,
		},
		{
			name:  "Empty string",
			input: "",
			want:  false,
		},
		{
			name:  "Only lowercase",
			input: "alllowercase",
			want:  false,
		},
		{
			name:  "Only uppercase",
			input: "ALLUPPERCASE",
			want:  false,
		},
		{
			name:  "Only digits",
			input: "12345678",
			want:  false,
		},
		{
			name:  "Long but no digit",
			input: "VeryLongPasswordButNoDigit",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidPassword(tt.input)
			if got != tt.want {
				t.Errorf("IsValidPassword(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

// Test ContainsOnly function
func TestContainsOnly(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		allowed string
		want    bool
	}{
		// Valid cases
		{
			name:    "All chars allowed",
			input:   "abc",
			allowed: "abcdef",
			want:    true,
		},
		{
			name:    "Numbers allowed",
			input:   "123",
			allowed: "0123456789",
			want:    true,
		},
		{
			name:    "Repeated chars in input",
			input:   "aaa",
			allowed: "abc",
			want:    true,
		},
		{
			name:    "Empty input",
			input:   "",
			allowed: "abc",
			want:    true,
		},
		{
			name:    "Single char",
			input:   "a",
			allowed: "abc",
			want:    true,
		},
		{
			name:    "Input equals allowed",
			input:   "abc",
			allowed: "abc",
			want:    true,
		},

		// Invalid cases
		{
			name:    "Char not allowed",
			input:   "abc!",
			allowed: "abc",
			want:    false,
		},
		{
			name:    "Number not allowed",
			input:   "abc1",
			allowed: "abc",
			want:    false,
		},
		{
			name:    "Empty allowed set",
			input:   "a",
			allowed: "",
			want:    false,
		},
		{
			name:    "Space not allowed",
			input:   "a b",
			allowed: "ab",
			want:    false,
		},
		{
			name:    "Case sensitive",
			input:   "ABC",
			allowed: "abc",
			want:    false,
		},
		{
			name:    "Unicode char not allowed",
			input:   "café",
			allowed: "cafe",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsOnly(tt.input, tt.allowed)
			if got != tt.want {
				t.Errorf("ContainsOnly(%q, %q) = %v, want %v", tt.input, tt.allowed, got, tt.want)
			}
		})
	}
}

// Test edge cases for email validation
func TestIsValidEmailEdgeCases(t *testing.T) {
	edgeCases := []struct {
		input string
		want  bool
		note  string
	}{
		{"a@b.c", true, "minimal valid email"},
		{"user@sub.domain.example.com", true, "multiple dots in domain"},
		{"user.name@example.com", true, "dot in local part"},
		{"@.", false, "just @ and ."},
		{"user@.com", false, "domain starts with dot"},
		{"user@domain..com", true, "consecutive dots in domain (we don't reject this)"},
	}

	for _, tc := range edgeCases {
		got := IsValidEmail(tc.input)
		if got != tc.want {
			t.Errorf("IsValidEmail(%q) = %v, want %v (%s)", tc.input, got, tc.want, tc.note)
		}
	}
}

// Test boundary conditions for password validation
func TestIsValidPasswordBoundaries(t *testing.T) {
	boundaryTests := []struct {
		input string
		want  bool
		note  string
	}{
		{"Valid12", false, "7 chars - too short"},
		{"Valid123", true, "8 chars - exactly minimum"},
		{"ValidOne", false, "8 chars but no digit"},
		{"VALID123", false, "8 chars but no lowercase"},
		{"valid123", false, "8 chars but no uppercase"},
		{"VeryLongPasswordThatMeetsAllRequirements1", true, "very long password"},
	}

	for _, tc := range boundaryTests {
		got := IsValidPassword(tc.input)
		if got != tc.want {
			t.Errorf("IsValidPassword(%q) = %v, want %v (%s)", tc.input, got, tc.want, tc.note)
		}
	}
}

// Test Unicode handling in ContainsOnly
func TestContainsOnlyUnicode(t *testing.T) {
	tests := []struct {
		input   string
		allowed string
		want    bool
		note    string
	}{
		{"café", "caféè", true, "accented chars allowed"},
		{"😀😁", "😀😁😂", true, "emoji allowed"},
		{"hello", "héllo", false, "non-accented not in accented set"},
		{"привет", "привет", true, "Cyrillic exact match"},
	}

	for _, tc := range tests {
		got := ContainsOnly(tc.input, tc.allowed)
		if got != tc.want {
			t.Errorf("ContainsOnly(%q, %q) = %v, want %v (%s)",
				tc.input, tc.allowed, got, tc.want, tc.note)
		}
	}
}

// Benchmark validators
func BenchmarkIsValidEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsValidEmail("user@example.com")
	}
}

func BenchmarkIsValidPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsValidPassword("Strong123")
	}
}

func BenchmarkContainsOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ContainsOnly("abcdefghijklmnop", "abcdefghijklmnopqrstuvwxyz")
	}
}

// Demonstrate validation patterns
func TestValidationPatterns(t *testing.T) {
	// This test demonstrates common validation patterns

	// Pattern 1: Multiple conditions with AND
	hasLength := len("Test123") >= 8
	hasDigit := true // (assuming we checked)
	isValid := hasLength && hasDigit
	t.Logf("Pattern: Multiple AND conditions result in %v", isValid)

	// Pattern 2: Early return on failure
	validate := func(s string) bool {
		if len(s) == 0 {
			return false // Exit early
		}
		if len(s) < 3 {
			return false // Exit early
		}
		return true // Only if all checks pass
	}
	t.Logf("Pattern: Early return = %v", validate("test"))

	// Pattern 3: Building up boolean flags
	flags := make([]bool, 3)
	flags[0] = true // found uppercase
	flags[1] = true // found lowercase
	flags[2] = true // found digit
	allTrue := flags[0] && flags[1] && flags[2]
	t.Logf("Pattern: Combining flags = %v", allTrue)
}
