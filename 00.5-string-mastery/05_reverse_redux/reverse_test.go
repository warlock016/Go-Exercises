package reverse

import "testing"

// Test cases shared across all three implementations
var reverseTests = []struct {
	name  string
	input string
	want  string
}{
	{"Simple ASCII", "hello", "olleh"},
	{"Empty string", "", ""},
	{"Single char", "a", "a"},
	{"Palindrome", "racecar", "racecar"},
	{"With spaces", "hello world", "dlrow olleh"},
	{"With accent", "café", "éfac"},
	{"Chinese chars", "Hello, 世界", "界世 ,olleH"},
	{"Emoji", "👍🎉", "🎉👍"},
	{"Mixed complex", "Go🚀2024", "4202🚀oG"},
	{"Arabic", "مرحبا", "ابحرم"},
	{"Multiple emoji", "🔥💯🎯", "🎯💯🔥"},
}

func TestReverseSimple(t *testing.T) {
	for _, tt := range reverseTests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseSimple(tt.input)
			if got != tt.want {
				t.Errorf("ReverseSimple(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestReverseBuilder(t *testing.T) {
	for _, tt := range reverseTests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseBuilder(tt.input)
			if got != tt.want {
				t.Errorf("ReverseBuilder(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestReverseUTF8(t *testing.T) {
	for _, tt := range reverseTests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseUTF8(tt.input)
			if got != tt.want {
				t.Errorf("ReverseUTF8(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// Verify all three methods produce identical results
func TestConsistency(t *testing.T) {
	testStrings := []string{
		"hello",
		"café",
		"Hello, 世界",
		"👍🎉🔥",
		"",
		"a",
	}

	for _, s := range testStrings {
		simple := ReverseSimple(s)
		builder := ReverseBuilder(s)
		utf8 := ReverseUTF8(s)

		if simple != builder {
			t.Errorf("For %q: ReverseSimple = %q, ReverseBuilder = %q (should match)", s, simple, builder)
		}
		if simple != utf8 {
			t.Errorf("For %q: ReverseSimple = %q, ReverseUTF8 = %q (should match)", s, simple, utf8)
		}
	}
}

// This test demonstrates WHY reversing bytes breaks UTF-8
func TestWhyByteReversalBreaks(t *testing.T) {
	s := "café"

	// WRONG way - reverse bytes
	bytes := []byte(s)
	// Don't actually reverse here, just show what would happen
	t.Logf("Original string: %q", s)
	t.Logf("Bytes: %v", bytes)
	t.Logf("If we reversed bytes, the 'é' (bytes [195 169]) would be split!")

	// RIGHT way - reverse runes
	reversed := ReverseSimple(s)
	t.Logf("Correctly reversed: %q", reversed)

	if reversed != "éfac" {
		t.Errorf("Expected %q, got %q", "éfac", reversed)
	}
}

// Performance comparison of the three approaches
func BenchmarkReverse(b *testing.B) {
	testString := "Hello, 世界! How are you today? 😊🚀💻"

	b.Run("Simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSimple(testString)
		}
	})

	b.Run("Builder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseBuilder(testString)
		}
	})

	b.Run("UTF8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseUTF8(testString)
		}
	})
}

// Milestone test - comparing to your diagnostic performance
func TestMilestone(t *testing.T) {
	// This is the same test case from your diagnostic
	testCases := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"Go!", "!oG"},
		{"café", "éfac"},
		{"", ""},
	}

	t.Log("🎉 Congratulations! You're solving the problem that took 45 minutes in the diagnostic.")
	t.Log("Now you understand:")
	t.Log("  ✓ Why strings are byte sequences")
	t.Log("  ✓ Why runes represent characters")
	t.Log("  ✓ How to iterate correctly")
	t.Log("  ✓ Three different approaches to reversal")

	for _, tc := range testCases {
		got := ReverseSimple(tc.input)
		if got != tc.want {
			t.Errorf("ReverseSimple(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}

	t.Log("\n🚀 From 45-minute struggle to confident understanding. Well done!")
}
