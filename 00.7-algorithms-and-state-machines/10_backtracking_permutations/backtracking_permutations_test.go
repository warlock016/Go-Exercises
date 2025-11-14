package backtrackingpermutations

import (
	"sort"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "Three characters",
			input: "abc",
			want:  []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name:  "Two characters",
			input: "ab",
			want:  []string{"ab", "ba"},
		},
		{
			name:  "Single character",
			input: "a",
			want:  []string{"a"},
		},
		{
			name:  "Empty string",
			input: "",
			want:  []string{""},
		},
		{
			name:  "Four characters",
			input: "abcd",
			want: []string{
				"abcd", "abdc", "acbd", "acdb", "adbc", "adcb",
				"bacd", "badc", "bcad", "bcda", "bdac", "bdca",
				"cabd", "cadb", "cbad", "cbda", "cdab", "cdba",
				"dabc", "dacb", "dbac", "dbca", "dcab", "dcba",
			},
		},
		{
			name:  "Numbers as string",
			input: "123",
			want:  []string{"123", "132", "213", "231", "312", "321"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Permute(tt.input)

			// Check count
			expectedCount := factorial(len([]rune(tt.input)))
			if len(got) != expectedCount {
				t.Errorf("Permute(%q) returned %d permutations, want %d", tt.input, len(got), expectedCount)
			}

			// Sort both slices for comparison (order doesn't matter)
			sortedGot := make([]string, len(got))
			copy(sortedGot, got)
			sort.Strings(sortedGot)

			sortedWant := make([]string, len(tt.want))
			copy(sortedWant, tt.want)
			sort.Strings(sortedWant)

			// Compare
			if !stringSlicesEqual(sortedGot, sortedWant) {
				t.Errorf("Permute(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestPermuteCount(t *testing.T) {
	// Test that we generate correct number of permutations
	tests := []struct {
		input         string
		expectedCount int
	}{
		{"", 1},      // 0! = 1
		{"a", 1},     // 1! = 1
		{"ab", 2},    // 2! = 2
		{"abc", 6},   // 3! = 6
		{"abcd", 24}, // 4! = 24
		{"12345", 120}, // 5! = 120
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := Permute(tt.input)
			if len(got) != tt.expectedCount {
				t.Errorf("Permute(%q) generated %d permutations, want %d", tt.input, len(got), tt.expectedCount)
			}
		})
	}
}

func TestPermuteAllUnique(t *testing.T) {
	// Verify all permutations are unique (for distinct characters)
	input := "abc"
	got := Permute(input)

	seen := make(map[string]bool)
	for _, perm := range got {
		if seen[perm] {
			t.Errorf("Permute(%q) generated duplicate: %q", input, perm)
		}
		seen[perm] = true
	}
}

func TestPermuteAllValid(t *testing.T) {
	// Verify each permutation contains exactly the input characters
	input := "abc"
	got := Permute(input)

	for _, perm := range got {
		if !isPermutation(perm, input) {
			t.Errorf("Permute(%q) generated invalid permutation: %q", input, perm)
		}
	}
}

func TestPermuteWithDuplicates(t *testing.T) {
	// When input has duplicate characters, we'll get duplicate permutations
	// This is expected for the basic implementation
	input := "aab"
	got := Permute(input)

	expectedCount := factorial(len(input)) // 3! = 6
	if len(got) != expectedCount {
		t.Errorf("Permute(%q) returned %d permutations, want %d", input, len(got), expectedCount)
	}

	// Each permutation should still be valid
	for _, perm := range got {
		if !isPermutation(perm, input) {
			t.Errorf("Permute(%q) generated invalid permutation: %q", input, perm)
		}
	}
}

func TestPermuteUnicode(t *testing.T) {
	// Test with Unicode characters
	input := "café"
	got := Permute(input)

	expectedCount := factorial(4) // 4 characters → 4! = 24
	if len(got) != expectedCount {
		t.Errorf("Permute(%q) returned %d permutations, want %d", input, len(got), expectedCount)
	}

	// Verify all contain the same runes
	for _, perm := range got {
		if !isPermutation(perm, input) {
			t.Errorf("Permute(%q) generated invalid permutation: %q", input, perm)
		}
	}
}

// Helper functions

// factorial calculates n!
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// stringSlicesEqual checks if two string slices are equal
func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// isPermutation checks if s1 is a valid permutation of s2
func isPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// Count character frequencies
	freq1 := make(map[rune]int)
	freq2 := make(map[rune]int)

	for _, r := range s1 {
		freq1[r]++
	}
	for _, r := range s2 {
		freq2[r]++
	}

	// Compare frequencies
	if len(freq1) != len(freq2) {
		return false
	}

	for r, count := range freq1 {
		if freq2[r] != count {
			return false
		}
	}

	return true
}

// Benchmark to measure performance
func BenchmarkPermute3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Permute("abc")
	}
}

func BenchmarkPermute4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Permute("abcd")
	}
}

func BenchmarkPermute5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Permute("abcde")
	}
}
