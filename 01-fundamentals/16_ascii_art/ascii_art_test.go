package asciiart

import "testing"

func TestDrawBox(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
		char   rune
		want   string
	}{
		{
			name:   "5x3 box with asterisks",
			width:  5,
			height: 3,
			want:   "*****\n*   *\n*****\n",
			char:   '*',
		},
		{
			name:   "7x4 box with hashes",
			width:  7,
			height: 4,
			want:   "#######\n#     #\n#     #\n#######\n",
			char:   '#',
		},
		{
			name:   "3x3 box with plus signs",
			width:  3,
			height: 3,
			want:   "+++\n+ +\n+++\n",
			char:   '+',
		},
		{
			name:   "2x2 minimum box",
			width:  2,
			height: 2,
			want:   "**\n**\n",
			char:   '*',
		},
		{
			name:   "10x5 box with equals",
			width:  10,
			height: 5,
			want:   "==========\n=        =\n=        =\n=        =\n==========\n",
			char:   '=',
		},
		{
			name:   "width too small",
			width:  1,
			height: 3,
			want:   "",
			char:   '*',
		},
		{
			name:   "height too small",
			width:  5,
			height: 1,
			want:   "",
			char:   '*',
		},
		{
			name:   "both dimensions too small",
			width:  0,
			height: 0,
			want:   "",
			char:   '*',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DrawBox(tt.width, tt.height, tt.char)
			if got != tt.want {
				t.Errorf("DrawBox(%d, %d, %q) = \n%q\nwant:\n%q", tt.width, tt.height, tt.char, got, tt.want)
			}
		})
	}
}

func TestDrawDiamond(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "diamond with n=1",
			n:    1,
			want: "*\n",
		},
		{
			name: "diamond with n=3",
			n:    3,
			want: "  *\n ***\n*****\n ***\n  *\n",
		},
		{
			name: "diamond with n=4",
			n:    4,
			want: "   *\n  ***\n *****\n*******\n *****\n  ***\n   *\n",
		},
		{
			name: "diamond with n=2",
			n:    2,
			want: " *\n***\n *\n",
		},
		{
			name: "diamond with n=5",
			n:    5,
			want: "    *\n   ***\n  *****\n *******\n*********\n *******\n  *****\n   ***\n    *\n",
		},
		{
			name: "n too small",
			n:    0,
			want: "",
		},
		{
			name: "negative n",
			n:    -1,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DrawDiamond(tt.n)
			if got != tt.want {
				t.Errorf("DrawDiamond(%d) = \n%q\nwant:\n%q", tt.n, got, tt.want)
			}
		})
	}
}

func TestDrawChessboard(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "1x1 chessboard",
			n:    1,
			want: "█\n",
		},
		{
			name: "2x2 chessboard",
			n:    2,
			want: "█░\n░█\n",
		},
		{
			name: "3x3 chessboard",
			n:    3,
			want: "█░█\n░█░\n█░█\n",
		},
		{
			name: "4x4 chessboard",
			n:    4,
			want: "█░█░\n░█░█\n█░█░\n░█░█\n",
		},
		{
			name: "5x5 chessboard",
			n:    5,
			want: "█░█░█\n░█░█░\n█░█░█\n░█░█░\n█░█░█\n",
		},
		{
			name: "8x8 chessboard",
			n:    8,
			want: "█░█░█░█░\n░█░█░█░█\n█░█░█░█░\n░█░█░█░█\n█░█░█░█░\n░█░█░█░█\n█░█░█░█░\n░█░█░█░█\n",
		},
		{
			name: "n too small",
			n:    0,
			want: "",
		},
		{
			name: "negative n",
			n:    -5,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DrawChessboard(tt.n)
			if got != tt.want {
				t.Errorf("DrawChessboard(%d) = \n%q\nwant:\n%q", tt.n, got, tt.want)
			}
		})
	}
}

// TestDrawBoxVisual prints the box for manual verification during development
func TestDrawBoxVisual(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping visual test in short mode")
	}

	box := DrawBox(8, 5, '█')
	t.Logf("\n8x5 Box:\n%s", box)
}

// TestDrawDiamondVisual prints the diamond for manual verification during development
func TestDrawDiamondVisual(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping visual test in short mode")
	}

	diamond := DrawDiamond(5)
	t.Logf("\nDiamond (n=5):\n%s", diamond)
}

// TestDrawChessboardVisual prints the chessboard for manual verification during development
func TestDrawChessboardVisual(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping visual test in short mode")
	}

	board := DrawChessboard(8)
	t.Logf("\n8x8 Chessboard:\n%s", board)
}

// Benchmark tests to measure performance
func BenchmarkDrawBox(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DrawBox(20, 10, '*')
	}
}

func BenchmarkDrawDiamond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DrawDiamond(10)
	}
}

func BenchmarkDrawChessboard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DrawChessboard(16)
	}
}
