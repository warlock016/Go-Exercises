package patternprinter

import "testing"

func TestPrintSquare(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "1x1 square",
			n:    1,
			want: "*\n",
		},
		{
			name: "3x3 square",
			n:    3,
			want: "***\n***\n***\n",
		},
		{
			name: "5x5 square",
			n:    5,
			want: "*****\n*****\n*****\n*****\n*****\n",
		},
		{
			name: "zero size",
			n:    0,
			want: "",
		},
		{
			name: "negative size",
			n:    -1,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrintSquare(tt.n)
			if got != tt.want {
				t.Errorf("PrintSquare(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestPrintTriangle(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "1 row triangle",
			n:    1,
			want: "*\n",
		},
		{
			name: "4 row triangle",
			n:    4,
			want: "*\n**\n***\n****\n",
		},
		{
			name: "6 row triangle",
			n:    6,
			want: "*\n**\n***\n****\n*****\n******\n",
		},
		{
			name: "zero size",
			n:    0,
			want: "",
		},
		{
			name: "negative size",
			n:    -3,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrintTriangle(tt.n)
			if got != tt.want {
				t.Errorf("PrintTriangle(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestPrintPyramid(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "1 row pyramid",
			n:    1,
			want: "*\n",
		},
		{
			name: "3 row pyramid",
			n:    3,
			want: "  *\n ***\n*****\n",
		},
		{
			name: "5 row pyramid",
			n:    5,
			want: "    *\n   ***\n  *****\n *******\n*********\n",
		},
		{
			name: "2 row pyramid",
			n:    2,
			want: " *\n***\n",
		},
		{
			name: "zero size",
			n:    0,
			want: "",
		},
		{
			name: "negative size",
			n:    -2,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrintPyramid(tt.n)
			if got != tt.want {
				t.Errorf("PrintPyramid(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestPrintNumberSquare(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "1x1 number square",
			n:    1,
			want: "1\n",
		},
		{
			name: "3x3 number square",
			n:    3,
			want: "111\n222\n333\n",
		},
		{
			name: "5x5 number square",
			n:    5,
			want: "11111\n22222\n33333\n44444\n55555\n",
		},
		{
			name: "10x10 number square (double digits)",
			n:    10,
			want: "1111111111\n2222222222\n3333333333\n4444444444\n5555555555\n6666666666\n7777777777\n8888888888\n9999999999\n10101010101010101010\n",
		},
		{
			name: "zero size",
			n:    0,
			want: "",
		},
		{
			name: "negative size",
			n:    -5,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrintNumberSquare(tt.n)
			if got != tt.want {
				t.Errorf("PrintNumberSquare(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

// Benchmark tests to demonstrate performance differences
func BenchmarkPrintSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintSquare(10)
	}
}

func BenchmarkPrintTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintTriangle(10)
	}
}

func BenchmarkPrintPyramid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintPyramid(10)
	}
}

func BenchmarkPrintNumberSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintNumberSquare(10)
	}
}
