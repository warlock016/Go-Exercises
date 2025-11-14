package regexmatcher

import "testing"

func TestKleeneStar(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		text    string
		want    bool
	}{
		{"zero matches", "a*", "", true},
		{"one match", "a*", "a", true},
		{"three matches", "a*", "aaa", true},
		{"wrong character", "a*", "b", false},
		{"partial match", "a*", "aab", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Match(tt.pattern, tt.text)
			if got != tt.want {
				t.Errorf("Match(%q, %q) = %v, want %v", tt.pattern, tt.text, got, tt.want)
			}
		})
	}
}

func TestPlus(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		text    string
		want    bool
	}{
		{"zero matches (should fail)", "a+", "", false},
		{"one match", "a+", "a", true},
		{"three matches", "a+", "aaa", true},
		{"wrong character", "a+", "b", false},
		{"partial match", "a+", "aab", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Match(tt.pattern, tt.text)
			if got != tt.want {
				t.Errorf("Match(%q, %q) = %v, want %v", tt.pattern, tt.text, got, tt.want)
			}
		})
	}
}

func TestLiterals(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		text    string
		want    bool
	}{
		{"exact match", "abc", "abc", true},
		{"too short", "abc", "ab", false},
		{"too long", "abc", "abcd", false},
		{"wrong chars", "abc", "xyz", false},
		{"empty both", "", "", true},
		{"empty text", "a", "", false},
		{"empty pattern", "", "a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Match(tt.pattern, tt.text)
			if got != tt.want {
				t.Errorf("Match(%q, %q) = %v, want %v", tt.pattern, tt.text, got, tt.want)
			}
		})
	}
}

func TestStarFollowedByChar(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		text    string
		want    bool
	}{
		{"zero a's then b", "a*b", "b", true},
		{"one a then b", "a*b", "ab", true},
		{"three a's then b", "a*b", "aaab", true},
		{"just a's (no b)", "a*b", "aaa", false},
		{"just b (wrong char)", "a*b", "c", false},
		{"extra char after", "a*b", "aabc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Match(tt.pattern, tt.text)
			if got != tt.want {
				t.Errorf("Match(%q, %q) = %v, want %v", tt.pattern, tt.text, got, tt.want)
			}
		})
	}
}

func TestPlusFollowedByChar(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		text    string
		want    bool
	}{
		{"one a then b", "a+b", "ab", true},
		{"three a's then b", "a+b", "aaab", true},
		{"zero a's (should fail)", "a+b", "b", false},
		{"just a's (no b)", "a+b", "aaa", false},
		{"extra char after", "a+b", "aabc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Match(tt.pattern, tt.text)
			if got != tt.want {
				t.Errorf("Match(%q, %q) = %v, want %v", tt.pattern, tt.text, got, tt.want)
			}
		})
	}
}

func TestCharStarChar(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		text    string
		want    bool
	}{
		{"zero b's", "ab*c", "ac", true},
		{"one b", "ab*c", "abc", true},
		{"four b's", "ab*c", "abbbbc", true},
		{"missing a", "ab*c", "bbbbc", false},
		{"missing c", "ab*c", "abbbb", false},
		{"wrong order", "ab*c", "acb", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Match(tt.pattern, tt.text)
			if got != tt.want {
				t.Errorf("Match(%q, %q) = %v, want %v", tt.pattern, tt.text, got, tt.want)
			}
		})
	}
}

func TestMultipleStars(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		text    string
		want    bool
	}{
		{"all zero", "a*b*c*", "", true},
		{"all one", "a*b*c*", "abc", true},
		{"multiple each", "a*b*c*", "aabbcc", true},
		{"zero a's", "a*b*c*", "bbc", true},
		{"zero b's", "a*b*c*", "aac", true},
		{"zero c's", "a*b*c*", "aabb", true},
		{"wrong order", "a*b*c*", "cba", false},
		{"interspersed", "a*b*c*", "ababc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Match(tt.pattern, tt.text)
			if got != tt.want {
				t.Errorf("Match(%q, %q) = %v, want %v", tt.pattern, tt.text, got, tt.want)
			}
		})
	}
}

func TestMixedOperators(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		text    string
		want    bool
	}{
		{"a+ b* c", "a+b*c", "abc", true},
		{"a+ b* c (no b)", "a+b*c", "aac", true},
		{"a+ b* c (no a - fail)", "a+b*c", "bbc", false},
		{"a* b+ c*", "a*b+c*", "bbc", true},
		{"a* b+ c* (zero b's - fail)", "a*b+c*", "aac", false},
		{"complex", "a+b*c+", "abbccc", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Match(tt.pattern, tt.text)
			if got != tt.want {
				t.Errorf("Match(%q, %q) = %v, want %v", tt.pattern, tt.text, got, tt.want)
			}
		})
	}
}

func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		text    string
		want    bool
	}{
		{"empty pattern and text", "", "", true},
		{"star on empty text", "a*", "", true},
		{"plus on empty text", "a+", "", false},
		{"literal on empty", "a", "", false},
		{"pattern but empty text", "abc", "", false},
		{"single char star", "a*", "aaaaaaaaaa", true},
		{"single char plus", "a+", "aaaaaaaaaa", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Match(tt.pattern, tt.text)
			if got != tt.want {
				t.Errorf("Match(%q, %q) = %v, want %v", tt.pattern, tt.text, got, tt.want)
			}
		})
	}
}

func TestUnicode(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		text    string
		want    bool
	}{
		{"emoji star", "😀*", "", true},
		{"emoji star one", "😀*", "😀", true},
		{"emoji star many", "😀*", "😀😀😀", true},
		{"emoji plus", "😀+", "😀😀", true},
		{"emoji literal", "😀", "😀", true},
		{"unicode star char", "ñ*a", "ñña", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Match(tt.pattern, tt.text)
			if got != tt.want {
				t.Errorf("Match(%q, %q) = %v, want %v", tt.pattern, tt.text, got, tt.want)
			}
		})
	}
}

// TestParsePattern verifies the pattern parser works correctly
func TestParsePattern(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		want    []Token
	}{
		{
			"star only",
			"a*",
			[]Token{{char: 'a', op: "*"}},
		},
		{
			"plus only",
			"a+",
			[]Token{{char: 'a', op: "+"}},
		},
		{
			"literal only",
			"abc",
			[]Token{
				{char: 'a', op: ""},
				{char: 'b', op: ""},
				{char: 'c', op: ""},
			},
		},
		{
			"star then literal",
			"a*b",
			[]Token{
				{char: 'a', op: "*"},
				{char: 'b', op: ""},
			},
		},
		{
			"mixed operators",
			"a+b*c",
			[]Token{
				{char: 'a', op: "+"},
				{char: 'b', op: "*"},
				{char: 'c', op: ""},
			},
		},
		{
			"multiple stars",
			"a*b*c*",
			[]Token{
				{char: 'a', op: "*"},
				{char: 'b', op: "*"},
				{char: 'c', op: "*"},
			},
		},
		{
			"empty pattern",
			"",
			[]Token{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parsePattern(tt.pattern)
			if len(got) != len(tt.want) {
				t.Fatalf("parsePattern(%q) returned %d tokens, want %d\nGot:  %+v\nWant: %+v",
					tt.pattern, len(got), len(tt.want), got, tt.want)
			}
			for i := range got {
				if got[i].char != tt.want[i].char || got[i].op != tt.want[i].op {
					t.Errorf("parsePattern(%q)[%d] = %+v, want %+v",
						tt.pattern, i, got[i], tt.want[i])
				}
			}
		})
	}
}
