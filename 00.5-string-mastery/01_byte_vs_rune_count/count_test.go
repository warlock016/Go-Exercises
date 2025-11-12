package count

import "testing"

func TestByteCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"ASCII only", "hello", 5},
		{"Empty string", "", 0},
		{"With accents", "café", 5}, // é is 2 bytes in UTF-8
		{"Emoji", "👍", 4},           // Most emoji are 4 bytes
		{"Mixed", "Hello, 世界", 13},  // 7 ASCII + 6 for two Chinese chars
		{"Multiple emoji", "👍👎", 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ByteCount(tt.input)
			if got != tt.want {
				t.Errorf("ByteCount(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestRuneCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"ASCII only", "hello", 5},
		{"Empty string", "", 0},
		{"With accents", "café", 4}, // 4 characters
		{"Emoji", "👍", 1},           // 1 character
		{"Mixed", "Hello, 世界", 9},  // 7 + comma + space + 2 Chinese chars
		{"Multiple emoji", "👍👎", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RuneCount(tt.input)
			if got != tt.want {
				t.Errorf("RuneCount(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

// This test demonstrates the KEY difference
func TestByteVsRuneDifference(t *testing.T) {
	testCases := []struct {
		str   string
		bytes int
		runes int
	}{
		{"hello", 5, 5},     // ASCII: bytes == runes
		{"café", 5, 4},      // Non-ASCII: bytes > runes
		{"👍", 4, 1},         // Emoji: bytes >> runes
		{"Hello, 世界", 13, 9}, // Mixed
	}

	for _, tc := range testCases {
		t.Run(tc.str, func(t *testing.T) {
			gotBytes := ByteCount(tc.str)
			gotRunes := RuneCount(tc.str)

			if gotBytes != tc.bytes {
				t.Errorf("%q: ByteCount = %v, want %v", tc.str, gotBytes, tc.bytes)
			}
			if gotRunes != tc.runes {
				t.Errorf("%q: RuneCount = %v, want %v", tc.str, gotRunes, tc.runes)
			}

			// The key insight
			if tc.bytes != tc.runes {
				t.Logf("✓ Confirmed: %q has %d bytes but only %d runes (characters)", tc.str, tc.bytes, tc.runes)
			}
		})
	}
}
