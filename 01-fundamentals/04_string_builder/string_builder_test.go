package stringbuilder

import (
	"strings"
	"testing"
)

func TestConcat(t *testing.T) {
	tests := []struct {
		name  string
		parts []string
		want  string
	}{
		{
			name:  "two strings",
			parts: []string{"Hello", "World"},
			want:  "HelloWorld",
		},
		{
			name:  "three strings with space",
			parts: []string{"Hello", " ", "World"},
			want:  "Hello World",
		},
		{
			name:  "multiple strings",
			parts: []string{"Go", "lang", "is", "awesome"},
			want:  "Golangisawesome",
		},
		{
			name:  "empty strings included",
			parts: []string{"Hello", "", "World"},
			want:  "HelloWorld",
		},
		{
			name:  "single string",
			parts: []string{"Solo"},
			want:  "Solo",
		},
		{
			name:  "no arguments",
			parts: []string{},
			want:  "",
		},
		{
			name:  "all empty strings",
			parts: []string{"", "", ""},
			want:  "",
		},
		{
			name:  "unicode strings",
			parts: []string{"Hello", " ", "世界", "!"},
			want:  "Hello 世界!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Concat(tt.parts...)
			if got != tt.want {
				t.Errorf("Concat(%v) = %q, want %q", tt.parts, got, tt.want)
			}
		})
	}
}

func TestBuildGreeting(t *testing.T) {
	tests := []struct {
		name  string
		pName string
		title string
		want  string
	}{
		{
			name:  "with title Dr",
			pName: "Smith",
			title: "Dr.",
			want:  "Hello, Dr. Smith!",
		},
		{
			name:  "with title Prof",
			pName: "Alice",
			title: "Prof.",
			want:  "Hello, Prof. Alice!",
		},
		{
			name:  "no title",
			pName: "Johnson",
			title: "",
			want:  "Hello, Johnson!",
		},
		{
			name:  "with title Mr",
			pName: "Bob",
			title: "Mr.",
			want:  "Hello, Mr. Bob!",
		},
		{
			name:  "unicode name with title",
			pName: "李明",
			title: "Dr.",
			want:  "Hello, Dr. 李明!",
		},
		{
			name:  "unicode name no title",
			pName: "José",
			title: "",
			want:  "Hello, José!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildGreeting(tt.pName, tt.title)
			if got != tt.want {
				t.Errorf("BuildGreeting(%q, %q) = %q, want %q", tt.pName, tt.title, got, tt.want)
			}
		})
	}
}

func TestRepeatWithSeparator(t *testing.T) {
	tests := []struct {
		name string
		s    string
		n    int
		sep  string
		want string
	}{
		{
			name: "repeat 3 times with dash",
			s:    "Go",
			n:    3,
			sep:  "-",
			want: "Go-Go-Go",
		},
		{
			name: "repeat 5 times no separator",
			s:    "*",
			n:    5,
			sep:  "",
			want: "*****",
		},
		{
			name: "repeat once with separator",
			s:    "ha",
			n:    1,
			sep:  ",",
			want: "ha",
		},
		{
			name: "repeat zero times",
			s:    "x",
			n:    0,
			sep:  "-",
			want: "",
		},
		{
			name: "repeat negative times",
			s:    "x",
			n:    -5,
			sep:  "-",
			want: "",
		},
		{
			name: "repeat with comma separator",
			s:    "item",
			n:    4,
			sep:  ", ",
			want: "item, item, item, item",
		},
		{
			name: "unicode string with separator",
			s:    "🚀",
			n:    3,
			sep:  " ",
			want: "🚀 🚀 🚀",
		},
		{
			name: "empty string repeated",
			s:    "",
			n:    3,
			sep:  "-",
			want: "--",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RepeatWithSeparator(tt.s, tt.n, tt.sep)
			if got != tt.want {
				t.Errorf("RepeatWithSeparator(%q, %d, %q) = %q, want %q", tt.s, tt.n, tt.sep, got, tt.want)
			}
		})
	}
}

func TestFormatTable(t *testing.T) {
	tests := []struct {
		name    string
		headers []string
		values  []string
		want    string
	}{
		{
			name:    "three columns",
			headers: []string{"Name", "Age", "City"},
			values:  []string{"Alice", "30", "NYC"},
			want:    "Name | Age | City\nAlice | 30 | NYC",
		},
		{
			name:    "two columns",
			headers: []string{"ID", "Status"},
			values:  []string{"42", "Active"},
			want:    "ID | Status\n42 | Active",
		},
		{
			name:    "single column",
			headers: []string{"Name"},
			values:  []string{"Bob"},
			want:    "Name\nBob",
		},
		{
			name:    "empty slices",
			headers: []string{},
			values:  []string{},
			want:    "",
		},
		{
			name:    "unicode values",
			headers: []string{"Name", "Country"},
			values:  []string{"José", "España"},
			want:    "Name | Country\nJosé | España",
		},
		{
			name:    "long values",
			headers: []string{"Username", "Email", "Role"},
			values:  []string{"john_doe", "john@example.com", "Admin"},
			want:    "Username | Email | Role\njohn_doe | john@example.com | Admin",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatTable(tt.headers, tt.values)
			if got != tt.want {
				t.Errorf("FormatTable() = %q, want %q", got, tt.want)
				// Show line-by-line comparison for better debugging
				gotLines := strings.Split(got, "\n")
				wantLines := strings.Split(tt.want, "\n")
				t.Logf("Got %d lines, want %d lines", len(gotLines), len(wantLines))
				for i := 0; i < len(gotLines) || i < len(wantLines); i++ {
					g := ""
					w := ""
					if i < len(gotLines) {
						g = gotLines[i]
					}
					if i < len(wantLines) {
						w = wantLines[i]
					}
					if g != w {
						t.Logf("  Line %d: got %q, want %q", i, g, w)
					}
				}
			}
		})
	}
}

// Benchmarks to demonstrate performance differences

func BenchmarkConcat(b *testing.B) {
	parts := []string{"Hello", " ", "World", "!", " ", "How", " ", "are", " ", "you", "?"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Concat(parts...)
	}
}

func BenchmarkConcatMany(b *testing.B) {
	// Test with many parts to show Builder efficiency
	parts := make([]string, 100)
	for i := range parts {
		parts[i] = "x"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Concat(parts...)
	}
}

func BenchmarkBuildGreeting(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = BuildGreeting("Smith", "Dr.")
	}
}

func BenchmarkRepeatWithSeparator(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RepeatWithSeparator("Go", 10, "-")
	}
}

func BenchmarkFormatTable(b *testing.B) {
	headers := []string{"Name", "Age", "City", "Country", "Email"}
	values := []string{"Alice", "30", "NYC", "USA", "alice@example.com"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FormatTable(headers, values)
	}
}
