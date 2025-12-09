package stringutils

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple", "hello", "olleh"},
		{"complex", "hola, adios", "soida ,aloh"},
		{"palindrome", "race car", "rac ecar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.input)

			if got != tt.want {
				t.Errorf("Reverse(%s) got: %s want: %s", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"not", "hello", false},
		{"simple", "racecar", true},
		{"complex", "race car", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.input)

			if got != tt.want {
				t.Errorf("Reverse(%s) got: %v want: %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{"simple", "a", 1},
		{"complex", "mississipi", 4},
		{"consonants", "qwytznvplkjhmxc", 0},
		{"vowels", "aeiou", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountVowels(tt.str)
			if got != tt.want {
				t.Errorf("Unexpected count: got %d, want %d", got, tt.want)
			}
		})
	}
}
