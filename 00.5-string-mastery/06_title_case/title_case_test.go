package titlecase

import "testing"

// Basic test cases for title case conversion
func TestToTitleCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Simple two words",
			input: "hello world",
			want:  "Hello World",
		},
		{
			name:  "Empty string",
			input: "",
			want:  "",
		},
		{
			name:  "Single character",
			input: "a",
			want:  "A",
		},
		{
			name:  "Already title case",
			input: "Hello World",
			want:  "Hello World",
		},
		{
			name:  "All lowercase",
			input: "the quick brown fox",
			want:  "The Quick Brown Fox",
		},
		{
			name:  "All uppercase (shouting)",
			input: "STOP SHOUTING",
			want:  "Stop Shouting",
		},
		{
			name:  "Mixed case",
			input: "ThIs Is WeIrD",
			want:  "This Is Weird",
		},
		{
			name:  "With hyphens",
			input: "hello-world-test",
			want:  "Hello-World-Test",
		},
		{
			name:  "With apostrophes",
			input: "it's a nice day",
			want:  "It'S A Nice Day",
		},
		{
			name:  "With punctuation",
			input: "hello, world! how are you?",
			want:  "Hello, World! How Are You?",
		},
		{
			name:  "Double spaces",
			input: "hello  world",
			want:  "Hello  World",
		},
		{
			name:  "Leading spaces",
			input: "  leading spaces",
			want:  "  Leading Spaces",
		},
		{
			name:  "Trailing spaces",
			input: "trailing spaces  ",
			want:  "Trailing Spaces  ",
		},
		{
			name:  "Only spaces",
			input: "   ",
			want:  "   ",
		},
		{
			name:  "Numbers and letters",
			input: "test123abc",
			want:  "Test123Abc",
		},
		{
			name:  "Starting with number",
			input: "123abc def",
			want:  "123Abc Def",
		},
		{
			name:  "Unicode accented characters",
			input: "café résumé naïve",
			want:  "Café Résumé Naïve",
		},
		{
			name:  "Unicode Greek",
			input: "γεια σου κόσμε",
			want:  "Γεια Σου Κόσμε",
		},
		{
			name:  "Unicode Cyrillic",
			input: "привет мир",
			want:  "Привет Мир",
		},
		{
			name:  "Only punctuation",
			input: "!@#$%",
			want:  "!@#$%",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToTitleCase(tt.input)
			if got != tt.want {
				t.Errorf("ToTitleCase(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// Test edge cases specifically
func TestToTitleCaseEdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Single space",
			input: " ",
			want:  " ",
		},
		{
			name:  "Single letter",
			input: "a",
			want:  "A",
		},
		{
			name:  "Single uppercase letter",
			input: "A",
			want:  "A",
		},
		{
			name:  "Tab character",
			input: "hello\tworld",
			want:  "Hello\tWorld",
		},
		{
			name:  "Newline character",
			input: "hello\nworld",
			want:  "Hello\nWorld",
		},
		{
			name:  "Multiple punctuation",
			input: "hello...world!!!",
			want:  "Hello...World!!!",
		},
		{
			name:  "Underscore separator",
			input: "hello_world_test",
			want:  "Hello_World_Test",
		},
		{
			name:  "Parentheses",
			input: "hello (world) test",
			want:  "Hello (World) Test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToTitleCase(tt.input)
			if got != tt.want {
				t.Errorf("ToTitleCase(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// Test that demonstrates the state tracking concept
func TestStateTracking(t *testing.T) {
	// This test shows how the "atWordStart" flag changes throughout processing
	input := "a b"
	want := "A B"
	got := ToTitleCase(input)

	if got != want {
		t.Errorf("ToTitleCase(%q) = %q, want %q", input, got, want)
		t.Log("State tracking explanation:")
		t.Log("  'a' - atWordStart=true  → capitalize → 'A', set false")
		t.Log("  ' ' - not a letter     → keep ' ', set atWordStart=true")
		t.Log("  'b' - atWordStart=true  → capitalize → 'B', set false")
	}
}

// Test comparison with strings.Title (deprecated but useful to understand differences)
func TestComparisonNotes(t *testing.T) {
	// Note: strings.Title is deprecated in Go 1.18+ because it doesn't
	// handle Unicode word boundaries correctly. Your implementation is
	// simpler but more predictable.

	testCases := []struct {
		input string
		want  string
		note  string
	}{
		{
			input: "hello world",
			want:  "Hello World",
			note:  "Basic case works",
		},
		{
			input: "HELLO WORLD",
			want:  "Hello World",
			note:  "Your version lowercases non-first letters",
		},
		{
			input: "it's nice",
			want:  "It'S Nice",
			note:  "Your version treats apostrophe as word boundary",
		},
	}

	for _, tc := range testCases {
		got := ToTitleCase(tc.input)
		if got != tc.want {
			t.Errorf("ToTitleCase(%q) = %q, want %q\nNote: %s",
				tc.input, got, tc.want, tc.note)
		} else {
			t.Logf("✓ %s: %q → %q", tc.note, tc.input, got)
		}
	}
}

// Benchmark to verify strings.Builder is used efficiently
func BenchmarkToTitleCase(b *testing.B) {
	testCases := []string{
		"hello world",
		"the quick brown fox jumps over the lazy dog",
		"HELLO WORLD HOW ARE YOU TODAY",
		"café résumé naïve señor",
	}

	for _, tc := range testCases {
		b.Run(tc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = ToTitleCase(tc)
			}
		})
	}
}
