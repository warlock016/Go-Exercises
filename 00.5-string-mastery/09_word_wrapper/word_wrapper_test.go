package wordwrapper

import (
	"strings"
	"testing"
)

// Test basic wrapping functionality
func TestWrapText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		maxWidth int
		want     string
	}{
		// Basic cases
		{
			name:     "Fits in one line",
			text:     "hello world",
			maxWidth: 20,
			want:     "hello world",
		},
		{
			name:     "Wrap at word boundary",
			text:     "hello world",
			maxWidth: 5,
			want:     "hello\nworld",
		},
		{
			name:     "Multiple lines",
			text:     "The quick brown fox",
			maxWidth: 10,
			want:     "The quick\nbrown fox",
		},
		{
			name:     "Wrap multiple times",
			text:     "one two three four five",
			maxWidth: 10,
			want:     "one two\nthree four\nfive",
		},

		// Edge cases
		{
			name:     "Empty string",
			text:     "",
			maxWidth: 10,
			want:     "",
		},
		{
			name:     "Single word",
			text:     "hello",
			maxWidth: 10,
			want:     "hello",
		},
		{
			name:     "Word longer than maxWidth",
			text:     "verylongword",
			maxWidth: 5,
			want:     "verylongword",
		},
		{
			name:     "Multiple spaces normalized",
			text:     "hello  world",
			maxWidth: 20,
			want:     "hello world",
		},
		{
			name:     "Leading and trailing spaces",
			text:     "  hello world  ",
			maxWidth: 20,
			want:     "hello world",
		},
		{
			name:     "Single character words",
			text:     "a b c d e f g",
			maxWidth: 5,
			want:     "a b c\nd e f\ng",
		},
		{
			name:     "Exact fit",
			text:     "hello",
			maxWidth: 5,
			want:     "hello",
		},
		{
			name:     "Just over limit",
			text:     "hello world",
			maxWidth: 11,
			want:     "hello world",
		},
		{
			name:     "Just under limit forces wrap",
			text:     "hello world",
			maxWidth: 10,
			want:     "hello\nworld",
		},

		// Special cases
		{
			name:     "All single letters",
			text:     "a b c d",
			maxWidth: 3,
			want:     "a b\nc d",
		},
		{
			name:     "MaxWidth of 1",
			text:     "a b c",
			maxWidth: 1,
			want:     "a\nb\nc",
		},
		{
			name:     "MaxWidth of 0",
			text:     "hello world",
			maxWidth: 0,
			want:     "hello world",
		},
		{
			name:     "MaxWidth negative",
			text:     "hello world",
			maxWidth: -1,
			want:     "hello world",
		},
		{
			name:     "Only spaces",
			text:     "     ",
			maxWidth: 10,
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WrapText(tt.text, tt.maxWidth)
			if got != tt.want {
				t.Errorf("WrapText(%q, %d) =\n%q\nwant:\n%q",
					tt.text, tt.maxWidth, got, tt.want)
				// Show line-by-line comparison
				gotLines := strings.Split(got, "\n")
				wantLines := strings.Split(tt.want, "\n")
				t.Logf("Got %d lines, want %d lines", len(gotLines), len(wantLines))
				for i := 0; i < max(len(gotLines), len(wantLines)); i++ {
					g := ""
					w := ""
					if i < len(gotLines) {
						g = gotLines[i]
					}
					if i < len(wantLines) {
						w = wantLines[i]
					}
					match := "✓"
					if g != w {
						match = "✗"
					}
					t.Logf("  Line %d: %s got=%q (len=%d) want=%q (len=%d)",
						i, match, g, len(g), w, len(w))
				}
			}
		})
	}
}

// Test that lines don't exceed maxWidth (except for words longer than maxWidth)
func TestWrapTextMaxWidth(t *testing.T) {
	tests := []struct {
		text     string
		maxWidth int
	}{
		{"The quick brown fox jumps over the lazy dog", 20},
		{"one two three four five six seven eight nine ten", 15},
		{"Lorem ipsum dolor sit amet consectetur", 25},
	}

	for _, tt := range tests {
		result := WrapText(tt.text, tt.maxWidth)
		lines := strings.Split(result, "\n")

		for i, line := range lines {
			if len(line) > tt.maxWidth {
				// Check if it's a single word longer than maxWidth
				words := strings.Fields(line)
				if len(words) == 1 {
					t.Logf("Line %d exceeds maxWidth but is a single long word (acceptable): %q (len=%d, max=%d)",
						i, line, len(line), tt.maxWidth)
				} else {
					t.Errorf("Line %d exceeds maxWidth: %q (len=%d, max=%d)",
						i, line, len(line), tt.maxWidth)
				}
			}
		}
	}
}

// Test specific algorithm behaviors
func TestWrapTextAlgorithm(t *testing.T) {
	// Test that spaces are handled correctly
	t.Run("Space handling", func(t *testing.T) {
		result := WrapText("a b c", 3)
		if !strings.HasPrefix(result, "a b") {
			t.Errorf("Expected lines to start without extra spaces, got: %q", result)
		}
		lines := strings.Split(result, "\n")
		for i, line := range lines {
			if strings.HasPrefix(line, " ") || strings.HasSuffix(line, " ") {
				t.Errorf("Line %d has leading or trailing space: %q", i, line)
			}
		}
	})

	// Test that last line is included
	t.Run("Last line included", func(t *testing.T) {
		result := WrapText("one two three", 7)
		if !strings.Contains(result, "three") {
			t.Errorf("Last word 'three' missing from result: %q", result)
		}
	})

	// Test boundary condition: exactly maxWidth
	t.Run("Exact maxWidth fit", func(t *testing.T) {
		result := WrapText("12345 67890", 11)
		if result != "12345 67890" {
			t.Errorf("Should fit in one line when exactly maxWidth: %q", result)
		}
	})

	// Test boundary condition: one char over
	t.Run("One char over maxWidth", func(t *testing.T) {
		result := WrapText("12345 678901", 11)
		if !strings.Contains(result, "\n") {
			t.Errorf("Should wrap when over maxWidth: %q", result)
		}
	})
}

// Test with longer realistic text
func TestWrapTextRealWorld(t *testing.T) {
	text := "The quick brown fox jumps over the lazy dog. This is a test of the word wrapping algorithm."
	maxWidth := 30

	result := WrapText(text, maxWidth)
	lines := strings.Split(result, "\n")

	t.Logf("Wrapped %d chars into %d lines (max width: %d):", len(text), len(lines), maxWidth)
	for i, line := range lines {
		t.Logf("  Line %d (len=%d): %q", i, len(line), line)
		if len(line) > maxWidth {
			words := strings.Fields(line)
			if len(words) != 1 {
				t.Errorf("Line %d exceeds maxWidth and has multiple words", i)
			}
		}
	}
}

// Test edge cases with long words
func TestWrapTextLongWords(t *testing.T) {
	tests := []struct {
		name string
		text string
		max  int
		note string
	}{
		{
			name: "Word longer than max",
			text: "supercalifragilisticexpialidocious",
			max:  10,
			note: "Long word should be on its own line",
		},
		{
			name: "Mix of long and short",
			text: "short supercalifragilisticexpialidocious short",
			max:  10,
			note: "Long word should be on its own line, surrounded by short words",
		},
		{
			name: "Multiple long words",
			text: "verylongword anotherlongword",
			max:  5,
			note: "Each long word gets its own line",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WrapText(tt.text, tt.max)
			t.Logf("%s:\n  Input: %q\n  Output: %q", tt.note, tt.text, result)

			// Verify all words are present
			inputWords := strings.Fields(tt.text)
			resultWords := strings.Fields(result)
			if len(inputWords) != len(resultWords) {
				t.Errorf("Word count mismatch: input has %d words, output has %d",
					len(inputWords), len(resultWords))
			}
		})
	}
}

// Benchmark the wrapping algorithm
func BenchmarkWrapText(b *testing.B) {
	text := "The quick brown fox jumps over the lazy dog. " +
		"This is a test of the word wrapping algorithm. " +
		"It should handle multiple sentences and wrap them correctly."

	b.Run("Width20", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = WrapText(text, 20)
		}
	})

	b.Run("Width40", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = WrapText(text, 40)
		}
	})

	b.Run("Width80", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = WrapText(text, 80)
		}
	})
}

// Helper function for max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Demonstrate the algorithm step by step
func TestAlgorithmDemonstration(t *testing.T) {
	text := "The quick brown"
	maxWidth := 10

	t.Logf("Wrapping: %q (maxWidth=%d)", text, maxWidth)
	t.Logf("")
	t.Logf("Words: %v", strings.Fields(text))
	t.Logf("")
	t.Logf("Process:")
	t.Logf("  1. Start with empty line")
	t.Logf("  2. Add 'The' → line='The' (len=3, ok)")
	t.Logf("  3. Try add 'quick' → future='The quick' (len=9, ok)")
	t.Logf("  4. Try add 'brown' → future='The quick brown' (len=15, > 10)")
	t.Logf("  5. WRAP: save 'The quick', start new line")
	t.Logf("  6. Add 'brown' → line='brown' (len=5, ok)")
	t.Logf("  7. End: save 'brown'")
	t.Logf("")

	result := WrapText(text, maxWidth)
	t.Logf("Result: %q", result)
	t.Logf("")
	t.Logf("Lines:")
	for i, line := range strings.Split(result, "\n") {
		t.Logf("  %d: %q (len=%d)", i, line, len(line))
	}
}
