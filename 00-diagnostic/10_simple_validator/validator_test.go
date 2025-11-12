package validator

import "testing"

func TestEmailValidator(t *testing.T) {
	var v Validator = EmailValidator{}

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"Valid email", "user@example.com", true},
		{"No @ symbol", "notanemail", false},
		{"Empty string", "", false},
		{"Just @ symbol", "@", true},
		{"Multiple @ symbols", "user@test@example.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := v.Validate(tt.input)
			if got != tt.want {
				t.Errorf("Validate(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
