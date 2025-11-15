package palindromecheck

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"simple palindrome", "racecar", true},
		{"case sensitive - not palindrome", "Racecar", false},
		{"not palindrome", "hello", false},
		{"empty string", "", true},
		{"single character", "a", true},
		{"two same characters", "aa", true},
		{"two different characters", "ab", false},
		{"even length palindrome", "noon", true},
		{"odd length palindrome", "civic", true},
		{"not palindrome - similar ends", "abcda", false},
		{"unicode palindrome", "こんにちんこ", false}, // Not a palindrome
		{"unicode single char", "日", true},
		{"longer palindrome", "madam", true},
		{"numbers palindrome", "12321", true},
		{"numbers not palindrome", "12345", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPalindromeIgnoreCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"lowercase palindrome", "racecar", true},
		{"mixed case palindrome", "Racecar", true},
		{"all caps palindrome", "RACECAR", true},
		{"mixed case complex", "RaceCar", true},
		{"not palindrome", "Hello", false},
		{"single char upper", "A", true},
		{"single char lower", "a", true},
		{"mixed case two chars", "Aa", true},
		{"mixed case not palindrome", "Ab", false},
		{"camel case palindrome", "AbA", true},
		{"empty string", "", true},
		{"long mixed case", "TacoCat", true},
		{"not palindrome mixed", "TacoTac", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindromeIgnoreCase(tt.input)
			if got != tt.want {
				t.Errorf("IsPalindromeIgnoreCase(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPalindromeIgnoreSpaces(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"classic phrase", "A man a plan a canal Panama", true},
		{"with spaces not palindrome", "race car", false},
		{"another classic", "Was it a car or a cat I saw", true},
		{"not palindrome", "hello world", false},
		{"NASA palindrome", "A Santa at NASA", true},
		{"no spaces palindrome", "racecar", true},
		{"no spaces not palindrome", "hello", false},
		{"single word with spaces", " a ", true},
		{"empty string", "", true},
		{"only spaces", "   ", true},
		{"multiple spaces between", "a  b  a", true},
		{"taco cat", "taco cat", true},
		{"never odd or even", "never odd or even", true},
		{"multiple words not palindrome", "hello there world", false},
		{"single letter repeated", "a a a a a", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindromeIgnoreSpaces(tt.input)
			if got != tt.want {
				t.Errorf("IsPalindromeIgnoreSpaces(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestLongestPalindromeSubstring(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		want          string
		alternatives  []string // Some test cases have multiple valid answers
		wantMinLength int      // Minimum acceptable length
	}{
		{
			name:          "two valid answers",
			input:         "babad",
			want:          "bab",
			alternatives:  []string{"bab", "aba"},
			wantMinLength: 3,
		},
		{
			name:  "even length palindrome",
			input: "cbbd",
			want:  "bb",
		},
		{
			name:  "entire string is palindrome",
			input: "racecar",
			want:  "racecar",
		},
		{
			name:          "double letters",
			input:         "hello",
			want:          "ll",
			alternatives:  []string{"ll"},
			wantMinLength: 2,
		},
		{
			name:  "single character",
			input: "a",
			want:  "a",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "no repeating chars",
			input: "abcde",
			want:  "a", // First single character
		},
		{
			name:  "long palindrome in middle",
			input: "forgeeksskeegfor",
			want:  "geeksskeeg",
		},
		{
			name:          "multiple same-length palindromes",
			input:         "aabbaa",
			want:          "aabbaa",
			alternatives:  []string{"aabbaa"},
			wantMinLength: 6,
		},
		{
			name:          "palindrome at start",
			input:         "abaxyz",
			want:          "aba",
			alternatives:  []string{"aba"},
			wantMinLength: 3,
		},
		{
			name:          "palindrome at end",
			input:         "xyzaba",
			want:          "aba",
			alternatives:  []string{"aba"},
			wantMinLength: 3,
		},
		{
			name:  "all same character",
			input: "aaaa",
			want:  "aaaa",
		},
		{
			name:          "two character palindrome",
			input:         "ac",
			want:          "a",
			alternatives:  []string{"a", "c"},
			wantMinLength: 1,
		},
		{
			name:  "complex case",
			input: "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendure",
			want:  "ranynar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestPalindromeSubstring(tt.input)

			// Check if result matches any of the valid answers
			if tt.alternatives != nil {
				found := false
				for _, alt := range tt.alternatives {
					if got == alt {
						found = true
						break
					}
				}
				if !found && len(got) >= tt.wantMinLength {
					// Verify it's actually a palindrome of acceptable length
					if IsPalindrome(got) {
						found = true
					}
				}
				if !found {
					t.Errorf("LongestPalindromeSubstring(%q) = %q, want one of %v (min length %d)",
						tt.input, got, tt.alternatives, tt.wantMinLength)
				}
			} else {
				if got != tt.want {
					// Verify the result is at least a palindrome
					if !IsPalindrome(got) {
						t.Errorf("LongestPalindromeSubstring(%q) = %q (not a palindrome), want %q",
							tt.input, got, tt.want)
					} else if len(got) != len(tt.want) {
						t.Errorf("LongestPalindromeSubstring(%q) = %q (length %d), want %q (length %d)",
							tt.input, got, len(got), tt.want, len(tt.want))
					}
				}
			}

			// Verify the result is actually a palindrome (unless empty)
			if got != "" && !IsPalindrome(got) {
				t.Errorf("LongestPalindromeSubstring(%q) returned %q which is not a palindrome",
					tt.input, got)
			}
		})
	}
}

func TestExpandAroundCenter(t *testing.T) {
	tests := []struct {
		name  string
		runes []rune
		left  int
		right int
		want  int
	}{
		{"odd palindrome center", []rune("aba"), 1, 1, 3},
		{"even palindrome center", []rune("abba"), 1, 2, 4},
		{"no expansion", []rune("abc"), 0, 0, 1},
		{"no match", []rune("ab"), 0, 1, 0},
		{"single char", []rune("a"), 0, 0, 1},
		{"full expansion", []rune("aaa"), 1, 1, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expandAroundCenter(tt.runes, tt.left, tt.right)
			if got != tt.want {
				t.Errorf("expandAroundCenter(%q, %d, %d) = %d, want %d",
					string(tt.runes), tt.left, tt.right, got, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkIsPalindrome(b *testing.B) {
	testString := "racecar"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindrome(testString)
	}
}

func BenchmarkIsPalindromeIgnoreCase(b *testing.B) {
	testString := "RaceCar"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindromeIgnoreCase(testString)
	}
}

func BenchmarkIsPalindromeIgnoreSpaces(b *testing.B) {
	testString := "A man a plan a canal Panama"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindromeIgnoreSpaces(testString)
	}
}

func BenchmarkLongestPalindromeSubstring(b *testing.B) {
	testString := "forgeeksskeegfor"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestPalindromeSubstring(testString)
	}
}

func BenchmarkLongestPalindromeSubstringLong(b *testing.B) {
	testString := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendure"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestPalindromeSubstring(testString)
	}
}
