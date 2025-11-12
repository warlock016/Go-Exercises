package charactercounter

import (
	"reflect"
	"testing"
)

// Test CountRunes function
func TestCountRunes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[rune]int
	}{
		{
			name:  "Simple word",
			input: "hello",
			want:  map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1},
		},
		{
			name:  "Empty string",
			input: "",
			want:  map[rune]int{},
		},
		{
			name:  "Single character",
			input: "a",
			want:  map[rune]int{'a': 1},
		},
		{
			name:  "Repeated characters",
			input: "aabbcc",
			want:  map[rune]int{'a': 2, 'b': 2, 'c': 2},
		},
		{
			name:  "All same character",
			input: "aaaa",
			want:  map[rune]int{'a': 4},
		},
		{
			name:  "With spaces and punctuation",
			input: "Hello, World!",
			want:  map[rune]int{'H': 1, 'e': 1, 'l': 3, 'o': 2, ',': 1, ' ': 1, 'W': 1, 'r': 1, 'd': 1, '!': 1},
		},
		{
			name:  "Unicode characters",
			input: "café",
			want:  map[rune]int{'c': 1, 'a': 1, 'f': 1, 'é': 1},
		},
		{
			name:  "Numbers",
			input: "123321",
			want:  map[rune]int{'1': 2, '2': 2, '3': 2},
		},
		{
			name:  "Mixed case",
			input: "AaBbCc",
			want:  map[rune]int{'A': 1, 'a': 1, 'B': 1, 'b': 1, 'C': 1, 'c': 1},
		},
		{
			name:  "Emoji",
			input: "🔥🔥💯",
			want:  map[rune]int{'🔥': 2, '💯': 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountRunes(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountRunes(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

// Test MostCommon function
func TestMostCommon(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  rune
	}{
		{
			name:  "Clear winner",
			input: "hello",
			want:  'l',
		},
		{
			name:  "Tie - first wins",
			input: "aabbcc",
			want:  'a',
		},
		{
			name:  "Tie - second wins",
			input: "abbc",
			want:  'b',
		},
		{
			name:  "Single character",
			input: "a",
			want:  'a',
		},
		{
			name:  "All same",
			input: "aaaa",
			want:  'a',
		},
		{
			name:  "Mississippi",
			input: "mississippi",
			want:  'i',
		},
		{
			name:  "Complex tie",
			input: "aabbccdd",
			want:  'a',
		},
		{
			name:  "With spaces",
			input: "a a a b b c",
			want:  ' ',
		},
		{
			name:  "Numbers",
			input: "112233",
			want:  '1',
		},
		{
			name:  "Empty string",
			input: "",
			want:  '\x00',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MostCommon(tt.input)
			if got != tt.want {
				t.Errorf("MostCommon(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// Test IsAnagram function
func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "Classic anagram",
			s1:   "listen",
			s2:   "silent",
			want: true,
		},
		{
			name: "Case insensitive",
			s1:   "Hello",
			s2:   "hello",
			want: true,
		},
		{
			name: "Mixed case anagram",
			s1:   "Listen",
			s2:   "Silent",
			want: true,
		},
		{
			name: "With spaces",
			s1:   "Astronomer",
			s2:   "Moon starer",
			want: true,
		},
		{
			name: "Multiple spaces",
			s1:   "a b c",
			s2:   "c b a",
			want: true,
		},
		{
			name: "Not anagrams",
			s1:   "hello",
			s2:   "world",
			want: false,
		},
		{
			name: "Different lengths",
			s1:   "abc",
			s2:   "abcd",
			want: false,
		},
		{
			name: "Different frequencies",
			s1:   "aab",
			s2:   "abb",
			want: false,
		},
		{
			name: "Empty strings",
			s1:   "",
			s2:   "",
			want: true,
		},
		{
			name: "One empty",
			s1:   "a",
			s2:   "",
			want: false,
		},
		{
			name: "Same string",
			s1:   "test",
			s2:   "test",
			want: true,
		},
		{
			name: "Spaces only matter for content",
			s1:   "a b c",
			s2:   "abc",
			want: true,
		},
		{
			name: "Unicode anagram",
			s1:   "café",
			s2:   "éfac",
			want: true,
		},
		{
			name: "Numbers",
			s1:   "123",
			s2:   "321",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAnagram(tt.s1, tt.s2)
			if got != tt.want {
				t.Errorf("IsAnagram(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

// Test edge cases for CountRunes
func TestCountRunesEdgeCases(t *testing.T) {
	// Test that map is properly initialized
	result := CountRunes("")
	if result == nil {
		t.Error("CountRunes should return empty map, not nil")
	}

	// Test multi-byte runes are counted as single characters
	result = CountRunes("😀😀😀")
	if result['😀'] != 3 {
		t.Errorf("Multi-byte emoji should be counted as single runes, got count: %d", result['😀'])
	}
}

// Test that MostCommon respects order for ties
func TestMostCommonOrderPreservation(t *testing.T) {
	// When there's a tie, the first occurrence should win
	testCases := []struct {
		input string
		want  rune
		note  string
	}{
		{"aabbcc", 'a', "all tied, 'a' is first"},
		{"bbaac", 'b', "all tied, 'b' is first"},
		{"xxyyzz", 'x', "all tied, 'x' is first"},
	}

	for _, tc := range testCases {
		got := MostCommon(tc.input)
		if got != tc.want {
			t.Errorf("MostCommon(%q) = %q, want %q (%s)", tc.input, got, tc.want, tc.note)
		}
	}
}

// Test IsAnagram normalization
func TestIsAnagramNormalization(t *testing.T) {
	// These should all be considered anagrams after normalization
	testCases := []struct {
		s1   string
		s2   string
		note string
	}{
		{"ABC", "abc", "case difference"},
		{"A B C", "abc", "spaces in first"},
		{"abc", "A B C", "spaces in second"},
		{"A B C", "C B A", "spaces in both"},
		{"  abc  ", "cba", "leading/trailing spaces"},
	}

	for _, tc := range testCases {
		if !IsAnagram(tc.s1, tc.s2) {
			t.Errorf("IsAnagram(%q, %q) should be true (%s)", tc.s1, tc.s2, tc.note)
		}
	}
}

// Benchmark tests
func BenchmarkCountRunes(b *testing.B) {
	testString := "The quick brown fox jumps over the lazy dog. The quick brown fox jumps over the lazy dog."
	for i := 0; i < b.N; i++ {
		_ = CountRunes(testString)
	}
}

func BenchmarkMostCommon(b *testing.B) {
	testString := "The quick brown fox jumps over the lazy dog. The quick brown fox jumps over the lazy dog."
	for i := 0; i < b.N; i++ {
		_ = MostCommon(testString)
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsAnagram("Astronomer", "Moon starer")
	}
}

// Demonstration of how maps work with zero values
func TestMapZeroValueBehavior(t *testing.T) {
	counts := make(map[rune]int)

	// Accessing a non-existent key returns zero value (0 for int)
	if counts['a'] != 0 {
		t.Error("Non-existent map key should return 0")
	}

	// This means counts[r]++ works even on first access
	counts['a']++ // 0 + 1 = 1
	counts['a']++ // 1 + 1 = 2

	if counts['a'] != 2 {
		t.Errorf("Expected 2, got %d", counts['a'])
	}

	t.Log("✓ Map zero value behavior allows simple increment pattern")
}
