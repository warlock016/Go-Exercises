package reverse

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Simple string", "hello", "olleh"},
		{"With punctuation", "Go!", "!oG"},
		{"Unicode characters", "café", "éfac"},
		{"Empty string", "", ""},
		{"Single character", "a", "a"},
		{"Palindrome", "racecar", "racecar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.input)
			if got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
