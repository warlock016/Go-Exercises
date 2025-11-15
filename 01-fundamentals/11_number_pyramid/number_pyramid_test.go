package numberpyramid

import (
	"testing"
)

func TestNumberTriangle(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "Single row",
			n:    1,
			want: "1\n",
		},
		{
			name: "Two rows",
			n:    2,
			want: "1\n12\n",
		},
		{
			name: "Three rows",
			n:    3,
			want: "1\n12\n123\n",
		},
		{
			name: "Five rows",
			n:    5,
			want: "1\n12\n123\n1234\n12345\n",
		},
		{
			name: "Seven rows",
			n:    7,
			want: "1\n12\n123\n1234\n12345\n123456\n1234567\n",
		},
		{
			name: "Zero rows (edge case)",
			n:    0,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumberTriangle(tt.n)
			if got != tt.want {
				t.Errorf("NumberTriangle(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestReversePyramid(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "Single row",
			n:    1,
			want: "1\n",
		},
		{
			name: "Two rows",
			n:    2,
			want: "12\n1\n",
		},
		{
			name: "Three rows",
			n:    3,
			want: "123\n12\n1\n",
		},
		{
			name: "Four rows",
			n:    4,
			want: "1234\n123\n12\n1\n",
		},
		{
			name: "Five rows",
			n:    5,
			want: "12345\n1234\n123\n12\n1\n",
		},
		{
			name: "Six rows",
			n:    6,
			want: "123456\n12345\n1234\n123\n12\n1\n",
		},
		{
			name: "Zero rows (edge case)",
			n:    0,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReversePyramid(tt.n)
			if got != tt.want {
				t.Errorf("ReversePyramid(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestMultiplicationTable(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "1x1 table",
			n:    1,
			want: "1\n",
		},
		{
			name: "2x2 table",
			n:    2,
			want: "1 2\n2 4\n",
		},
		{
			name: "3x3 table",
			n:    3,
			want: "1 2 3\n2 4 6\n3 6 9\n",
		},
		{
			name: "4x4 table",
			n:    4,
			want: "1 2 3 4\n2 4 6 8\n3 6 9 12\n4 8 12 16\n",
		},
		{
			name: "5x5 table",
			n:    5,
			want: "1 2 3 4 5\n2 4 6 8 10\n3 6 9 12 15\n4 8 12 16 20\n5 10 15 20 25\n",
		},
		{
			name: "Zero table (edge case)",
			n:    0,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MultiplicationTable(tt.n)
			if got != tt.want {
				t.Errorf("MultiplicationTable(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

// TestMultiplicationTableNoTrailingSpaces verifies no trailing spaces after last number
func TestMultiplicationTableNoTrailingSpaces(t *testing.T) {
	result := MultiplicationTable(3)
	lines := []string{"1 2 3\n", "2 4 6\n", "3 6 9\n"}

	expected := lines[0] + lines[1] + lines[2]
	if result != expected {
		t.Errorf("MultiplicationTable(3) has incorrect spacing or format.\nGot:  %q\nWant: %q", result, expected)
	}

	// Check that no line ends with space before newline
	for i, line := range lines {
		if len(line) > 1 && line[len(line)-2] == ' ' {
			t.Errorf("Row %d has trailing space before newline: %q", i+1, line)
		}
	}
}

// Benchmark functions to compare approaches
func BenchmarkNumberTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberTriangle(10)
	}
}

func BenchmarkReversePyramid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReversePyramid(10)
	}
}

func BenchmarkMultiplicationTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiplicationTable(10)
	}
}
