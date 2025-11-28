package string_recursion

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Empty string", "", ""},
		{"Single character", "a", "a"},
		{"Two characters", "ab", "ba"},
		{"Simple word", "hello", "olleh"},
		{"With punctuation", "Go!", "!oG"},
		{"Longer string", "recursion", "noisrucer"},
		{"Palindrome", "racecar", "racecar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseString(tt.input)
			if got != tt.want {
				t.Errorf("ReverseString(%q) = %q, want %q", tt.input, got, tt.want)
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
		{"Empty string", "", true},
		{"Single character", "a", true},
		{"Simple palindrome", "racecar", true},
		{"Simple palindrome uppercase", "Racecar", true},
		{"Not palindrome", "hello", false},
		{"With spaces (palindrome)", "race car", true},
		{"With spaces (not palindrome)", "race care", false},
		{"Complex palindrome", "A man a plan a canal Panama", true},
		{"Two characters same", "aa", true},
		{"Two characters different", "ab", false},
		{"Even length palindrome", "noon", true},
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

func TestCountVowels(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Empty string", "", 0},
		{"No vowels", "rhythm", 0},
		{"Single vowel", "a", 1},
		{"All vowels", "aeiou", 5},
		{"All vowels uppercase", "AEIOU", 5},
		{"Simple word", "hello", 2},
		{"Mixed case", "HeLLo", 2},
		{"With spaces", "hello world", 3},
		{"Consonants only", "sky", 0},
		{"Multiple same vowels", "aardvark", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountVowels(tt.input)
			if got != tt.want {
				t.Errorf("CountVowels(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestRemoveChar(t *testing.T) {
	tests := []struct {
		name string
		s    string
		char rune
		want string
	}{
		{"Empty string", "", 'a', ""},
		{"Character not present", "hello", 'z', "hello"},
		{"Remove single occurrence", "hello", 'h', "ello"},
		{"Remove multiple occurrences", "hello", 'l', "heo"},
		{"Remove all characters", "aaaa", 'a', ""},
		{"Remove from middle", "aardvark", 'a', "rdvrk"},
		{"Remove last character", "hello", 'o', "hell"},
		{"Case sensitive", "Hello", 'h', "Hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveChar(tt.s, tt.char)
			if got != tt.want {
				t.Errorf("RemoveChar(%q, %q) = %q, want %q", tt.s, tt.char, got, tt.want)
			}
		})
	}
}

// Benchmarks
func BenchmarkReverseString(b *testing.B) {
	s := "this is a test string for benchmarking"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseString(s)
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	s := "A man a plan a canal Panama"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindrome(s)
	}
}

func BenchmarkCountVowels(b *testing.B) {
	s := "this is a test string with several vowels"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountVowels(s)
	}
}
