package vowelcounter

import (
	"reflect"
	"testing"
)

func TestCountVowels(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"simple word", "hello", 2},
		{"all vowels uppercase", "AEIOU", 5},
		{"all vowels lowercase", "aeiou", 5},
		{"no vowels", "rhythm", 0},
		{"mixed case", "Programming", 3},
		{"empty string", "", 0},
		{"unicode with accented vowel", "café", 2},
		{"single vowel", "a", 1},
		{"single consonant", "b", 0},
		{"with numbers and symbols", "h3ll0 w0rld!", 1},
		{"multiple of same vowel", "aaa", 3},
		{"sentence", "The Quick Brown Fox", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountVowels(tt.input)
			if got != tt.want {
				t.Errorf("CountVowels(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCountConsonants(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"simple word", "hello", 3},
		{"all vowels", "AEIOU", 0},
		{"all consonants", "bcdfg", 5},
		{"mixed case", "Programming", 8},
		{"empty string", "", 0},
		{"with numbers and symbols", "123 abc!", 3},
		{"single consonant", "b", 1},
		{"single vowel", "a", 0},
		{"with spaces", "go lang", 4},
		{"unicode word", "café", 2},
		{"sentence", "The Quick Brown Fox", 11},
		{"only non-letters", "123!@#", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountConsonants(tt.input)
			if got != tt.want {
				t.Errorf("CountConsonants(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestVowelFrequency(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[rune]int
	}{
		{"simple word", "hello", map[rune]int{'e': 1, 'o': 1}},
		{"all vowels", "AEIOU", map[rune]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}},
		{"mixed case vowels", "AaEeIiOoUu", map[rune]int{'a': 2, 'e': 2, 'i': 2, 'o': 2, 'u': 2}},
		{"no vowels", "rhythm", map[rune]int{}},
		{"programming", "Programming", map[rune]int{'o': 1, 'a': 1, 'i': 1}},
		{"empty string", "", map[rune]int{}},
		{"multiple same vowel", "aaa", map[rune]int{'a': 3}},
		{"sentence", "The Quick Brown Fox", map[rune]int{'e': 1, 'u': 1, 'i': 1, 'o': 2}},
		{"unicode with accented", "café", map[rune]int{'a': 1, 'é': 1}},
		{"with numbers", "h3ll0 w0rld", map[rune]int{'o': 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VowelFrequency(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VowelFrequency(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkCountVowels(b *testing.B) {
	s := "The Quick Brown Fox Jumps Over The Lazy Dog"
	for i := 0; i < b.N; i++ {
		CountVowels(s)
	}
}

func BenchmarkCountConsonants(b *testing.B) {
	s := "The Quick Brown Fox Jumps Over The Lazy Dog"
	for i := 0; i < b.N; i++ {
		CountConsonants(s)
	}
}

func BenchmarkVowelFrequency(b *testing.B) {
	s := "The Quick Brown Fox Jumps Over The Lazy Dog"
	for i := 0; i < b.N; i++ {
		VowelFrequency(s)
	}
}
