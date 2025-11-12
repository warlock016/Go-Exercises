package builder

import "testing"

func TestJoinWords(t *testing.T) {
	tests := []struct {
		name      string
		words     []string
		separator string
		want      string
	}{
		{"Two words space", []string{"Hello", "World"}, " ", "Hello World"},
		{"Three words comma", []string{"a", "b", "c"}, ",", "a,b,c"},
		{"Hyphen separator", []string{"foo", "bar", "baz"}, "-", "foo-bar-baz"},
		{"Single word", []string{"alone"}, ",", "alone"},
		{"Empty slice", []string{}, ",", ""},
		{"Empty separator", []string{"a", "b"}, "", "ab"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinWords(tt.words, tt.separator)
			if got != tt.want {
				t.Errorf("JoinWords(%v, %q) = %q, want %q", tt.words, tt.separator, got, tt.want)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		name string
		s    string
		n    int
		want string
	}{
		{"Repeat Go 3 times", "Go", 3, "GoGoGo"},
		{"Repeat Hi! twice", "Hi!", 2, "Hi!Hi!"},
		{"Repeat once", "test", 1, "test"},
		{"Repeat zero times", "test", 0, ""},
		{"Repeat emoji", "👍", 3, "👍👍👍"},
		{"Repeat space", " ", 5, "     "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Repeat(tt.s, tt.n)
			if got != tt.want {
				t.Errorf("Repeat(%q, %d) = %q, want %q", tt.s, tt.n, got, tt.want)
			}
		})
	}
}

func TestBuildList(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		want  string
	}{
		{
			"Three items",
			[]string{"Apple", "Banana", "Cherry"},
			"1. Apple\n2. Banana\n3. Cherry",
		},
		{
			"Single item",
			[]string{"Only"},
			"1. Only",
		},
		{
			"Empty list",
			[]string{},
			"",
		},
		{
			"Two items",
			[]string{"First", "Second"},
			"1. First\n2. Second",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildList(tt.items)
			if got != tt.want {
				t.Errorf("BuildList() =\n%q\nwant:\n%q", got, tt.want)
			}
		})
	}
}

func TestBuildCSV(t *testing.T) {
	tests := []struct {
		name string
		rows [][]string
		want string
	}{
		{
			"Simple table",
			[][]string{
				{"Name", "Age", "City"},
				{"Alice", "25", "NYC"},
				{"Bob", "30", "LA"},
			},
			"Name,Age,City\nAlice,25,NYC\nBob,30,LA",
		},
		{
			"Single row",
			[][]string{
				{"a", "b", "c"},
			},
			"a,b,c",
		},
		{
			"Empty table",
			[][]string{},
			"",
		},
		{
			"Single column",
			[][]string{
				{"A"},
				{"B"},
				{"C"},
			},
			"A\nB\nC",
		},
		{
			"Two columns",
			[][]string{
				{"Name", "Score"},
				{"Alice", "95"},
				{"Bob", "87"},
			},
			"Name,Score\nAlice,95\nBob,87",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildCSV(tt.rows)
			if got != tt.want {
				t.Errorf("BuildCSV() =\n%q\nwant:\n%q", got, tt.want)
			}
		})
	}
}

// Demonstrates why strings.Builder is important for performance
func BenchmarkStringConcatenation(b *testing.B) {
	words := []string{"word1", "word2", "word3", "word4", "word5"}

	b.Run("Using += (slow)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := ""
			for _, word := range words {
				result += word
			}
			_ = result
		}
	})

	b.Run("Using strings.Builder (fast)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := JoinWords(words, "")
			_ = result
		}
	})
}
