package runlengthencoding

import (
	"strings"
	"testing"
)

// Test Encode function
func TestEncode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Basic cases
		{
			name:  "Multiple runs",
			input: "aaabbc",
			want:  "3a2b1c",
		},
		{
			name:  "No repetition",
			input: "abc",
			want:  "1a1b1c",
		},
		{
			name:  "All same character",
			input: "aaa",
			want:  "3a",
		},
		{
			name:  "Empty string",
			input: "",
			want:  "",
		},
		{
			name:  "Single character",
			input: "a",
			want:  "1a",
		},

		// More complex cases
		{
			name:  "Simple word",
			input: "hello",
			want:  "1h1e2l1o",
		},
		{
			name:  "Multiple runs",
			input: "aabbccddee",
			want:  "2a2b2c2d2e",
		},
		{
			name:  "Long run",
			input: "aaaaaaaaaa",
			want:  "10a",
		},
		{
			name:  "Mixed runs",
			input: "mississippi",
			want:  "1m1i2s1i2s1i2p1i",
		},

		// Edge cases
		{
			name:  "Two characters",
			input: "aa",
			want:  "2a",
		},
		{
			name:  "Alternating",
			input: "ababab",
			want:  "1a1b1a1b1a1b",
		},
		{
			name:  "Spaces",
			input: "a a a",
			want:  "1a1 1a1 1a",
		},
		{
			name:  "Multiple spaces",
			input: "a   b",
			want:  "1a3 1b",
		},

		// Unicode
		{
			name:  "Unicode characters",
			input: "😀😀😀",
			want:  "3😀",
		},
		{
			name:  "Accented characters",
			input: "ééé",
			want:  "3é",
		},
		{
			name:  "Mixed Unicode",
			input: "aéééb",
			want:  "1a3é1b",
		},

		// Numbers in input
		{
			name:  "Digit characters",
			input: "111222",
			want:  "31122",
		},
		{
			name:  "Mixed letters and digits",
			input: "aa11bb",
			want:  "2a2112b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Encode(tt.input)
			if got != tt.want {
				t.Errorf("Encode(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// Test Decode function
func TestDecode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Basic cases
		{
			name:  "Multiple runs",
			input: "3a2b1c",
			want:  "aaabbc",
		},
		{
			name:  "No repetition",
			input: "1a1b1c",
			want:  "abc",
		},
		{
			name:  "All same character",
			input: "3a",
			want:  "aaa",
		},
		{
			name:  "Empty string",
			input: "",
			want:  "",
		},
		{
			name:  "Single character",
			input: "1a",
			want:  "a",
		},

		// More complex cases
		{
			name:  "Simple word",
			input: "1h1e2l1o",
			want:  "hello",
		},
		{
			name:  "Multiple runs",
			input: "2a2b2c2d2e",
			want:  "aabbccddee",
		},
		{
			name:  "Long run",
			input: "10a",
			want:  "aaaaaaaaaa",
		},
		{
			name:  "Very long run",
			input: "100a",
			want:  strings.Repeat("a", 100),
		},

		// Edge cases
		{
			name:  "Two characters",
			input: "2a",
			want:  "aa",
		},
		{
			name:  "Alternating",
			input: "1a1b1a1b1a1b",
			want:  "ababab",
		},
		{
			name:  "Spaces",
			input: "1a1 1a1 1a",
			want:  "a a a",
		},
		{
			name:  "Multiple spaces",
			input: "1a3 1b",
			want:  "a   b",
		},

		// Unicode
		{
			name:  "Unicode characters",
			input: "3😀",
			want:  "😀😀😀",
		},
		{
			name:  "Accented characters",
			input: "3é",
			want:  "ééé",
		},

		// Multi-digit counts
		{
			name:  "Two-digit count",
			input: "13a",
			want:  strings.Repeat("a", 13),
		},
		{
			name:  "Three-digit count",
			input: "123a",
			want:  strings.Repeat("a", 123),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Decode(tt.input)
			if got != tt.want {
				t.Errorf("Decode(%q) = %q, want %q", tt.input, got, tt.want)
				if len(got) != len(tt.want) {
					t.Errorf("  Length: got %d, want %d", len(got), len(tt.want))
				}
			}
		})
	}
}

// Test round-trip property: Decode(Encode(s)) == s
func TestRoundTrip(t *testing.T) {
	testStrings := []string{
		"hello",
		"aaabbc",
		"abc",
		"",
		"a",
		"aaaaaa",
		"The quick brown fox",
		"mississippi",
		"a b c d e",
		"😀😀😀hello",
		"123456",
		"aaa111bbb",
		strings.Repeat("a", 100),
	}

	for _, original := range testStrings {
		encoded := Encode(original)
		decoded := Decode(encoded)

		if decoded != original {
			t.Errorf("Round-trip failed for %q", original)
			t.Errorf("  Original: %q", original)
			t.Errorf("  Encoded:  %q", encoded)
			t.Errorf("  Decoded:  %q", decoded)
		} else {
			t.Logf("✓ Round-trip success: %q → %q → %q", original, encoded, decoded)
		}
	}
}

// Test compression effectiveness
func TestCompressionEffectiveness(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectation string
	}{
		{
			name:        "Good compression",
			input:       "aaaaaabbbbbb",
			expectation: "compressed",
		},
		{
			name:        "No compression",
			input:       "abcdefgh",
			expectation: "expanded",
		},
		{
			name:        "Moderate compression",
			input:       "aabbccdd",
			expectation: "same or slightly compressed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := Encode(tt.input)
			ratio := float64(len(encoded)) / float64(len(tt.input))

			t.Logf("Input:  %q (len=%d)", tt.input, len(tt.input))
			t.Logf("Output: %q (len=%d)", encoded, len(encoded))
			t.Logf("Ratio:  %.2f (%s)", ratio, tt.expectation)

			if ratio < 1.0 {
				t.Logf("  Compression achieved: %.1f%% of original size", ratio*100)
			} else {
				t.Logf("  Expansion occurred: %.1f%% of original size", ratio*100)
			}
		})
	}
}

// Test edge cases specifically
func TestEncodeEdgeCases(t *testing.T) {
	// Single character repeated many times
	input := strings.Repeat("a", 50)
	encoded := Encode(input)
	if encoded != "50a" {
		t.Errorf("Encode(%d a's) = %q, want %q", 50, encoded, "50a")
	}

	// Alternating pattern
	input = "ababababab"
	encoded = Encode(input)
	expected := "1a1b1a1b1a1b1a1b1a1b"
	if encoded != expected {
		t.Errorf("Encode(%q) = %q, want %q", input, encoded, expected)
	}
}

func TestDecodeEdgeCases(t *testing.T) {
	// Multi-digit count
	input := "99a"
	decoded := Decode(input)
	if len(decoded) != 99 {
		t.Errorf("Decode(%q) should produce 99 characters, got %d", input, len(decoded))
	}

	// Multiple multi-digit counts
	input = "10a20b30c"
	decoded = Decode(input)
	expected := strings.Repeat("a", 10) + strings.Repeat("b", 20) + strings.Repeat("c", 30)
	if decoded != expected {
		t.Errorf("Decode(%q) failed", input)
		t.Errorf("  Got:  %q (len=%d)", decoded, len(decoded))
		t.Errorf("  Want: %q (len=%d)", expected, len(expected))
	}
}

// Test that Encode handles Unicode correctly
func TestEncodeUnicode(t *testing.T) {
	tests := []struct {
		input string
		want  string
		note  string
	}{
		{"😀😀😀", "3😀", "emoji"},
		{"ééé", "3é", "accented chars"},
		{"🔥🔥💯", "2🔥1💯", "mixed emoji"},
		{"привет", "1п1р1и1в1е1т", "Cyrillic"},
	}

	for _, tt := range tests {
		got := Encode(tt.input)
		if got != tt.want {
			t.Errorf("Encode(%q) = %q, want %q (%s)",
				tt.input, got, tt.want, tt.note)
		}
	}
}

// Benchmark encoding
func BenchmarkEncode(b *testing.B) {
	testCases := []string{
		"hello",
		"aaabbbccc",
		strings.Repeat("a", 100),
		"The quick brown fox jumps over the lazy dog",
	}

	for _, tc := range testCases {
		b.Run(tc[:min(len(tc), 20)], func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Encode(tc)
			}
		})
	}
}

// Benchmark decoding
func BenchmarkDecode(b *testing.B) {
	testCases := []string{
		"1h1e2l1o",
		"3a3b3c",
		"100a",
		"1T1h1e1 1q1u1i1c1k1 1b1r1o1w1n",
	}

	for _, tc := range testCases {
		b.Run(tc[:min(len(tc), 20)], func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Decode(tc)
			}
		})
	}
}

// Helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Demonstrate the algorithm step by step
func TestAlgorithmDemonstration(t *testing.T) {
	input := "aaabbc"

	t.Log("=== ENCODING DEMONSTRATION ===")
	t.Logf("Input: %q", input)
	t.Logf("")
	t.Logf("Process:")
	t.Logf("  Start: currentRune='a', count=1")
	t.Logf("  i=1: 'a' == 'a' → count=2")
	t.Logf("  i=2: 'a' == 'a' → count=3")
	t.Logf("  i=3: 'b' != 'a' → write '3a', reset to 'b', count=1")
	t.Logf("  i=4: 'b' == 'b' → count=2")
	t.Logf("  i=5: 'c' != 'b' → write '2b', reset to 'c', count=1")
	t.Logf("  End: write '1c'")
	t.Logf("")

	encoded := Encode(input)
	t.Logf("Result: %q", encoded)
	t.Logf("")

	t.Log("=== DECODING DEMONSTRATION ===")
	t.Logf("Input: %q", encoded)
	t.Logf("")
	t.Logf("Process:")
	t.Logf("  Read '3' → count=3")
	t.Logf("  Read 'a' → write 'a' 3 times → 'aaa'")
	t.Logf("  Read '2' → count=2")
	t.Logf("  Read 'b' → write 'b' 2 times → 'aaabb'")
	t.Logf("  Read '1' → count=1")
	t.Logf("  Read 'c' → write 'c' 1 time → 'aaabbc'")
	t.Logf("")

	decoded := Decode(encoded)
	t.Logf("Result: %q", decoded)
	t.Logf("")

	if decoded == input {
		t.Log("✓ Round-trip successful!")
	} else {
		t.Errorf("✗ Round-trip failed: got %q, want %q", decoded, input)
	}
}

// Test compression vs expansion
func TestCompressionAnalysis(t *testing.T) {
	t.Log("=== COMPRESSION ANALYSIS ===")
	t.Log("")

	tests := []string{
		"aaaaaabbbbbb",      // Good compression
		"abcdefgh",          // Expansion
		"aabbccdd",          // Neutral
		strings.Repeat("a", 100), // Excellent compression
	}

	for _, input := range tests {
		encoded := Encode(input)
		ratio := float64(len(encoded)) / float64(len(input))

		displayInput := input
		if len(input) > 30 {
			displayInput = input[:30] + "..."
		}

		t.Logf("Input:  %q (len=%d)", displayInput, len(input))
		t.Logf("Output: %q (len=%d)", encoded, len(encoded))
		t.Logf("Ratio:  %.2f", ratio)

		if ratio < 1.0 {
			improvement := (1.0 - ratio) * 100
			t.Logf("Result: Compression - %.1f%% smaller", improvement)
		} else if ratio > 1.0 {
			increase := (ratio - 1.0) * 100
			t.Logf("Result: Expansion - %.1f%% larger", increase)
		} else {
			t.Log("Result: No change")
		}
		t.Log("")
	}
}
