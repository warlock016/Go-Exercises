package n_queens

import "testing"

func TestCountNQueens(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 0},   // impossible
		{3, 0},   // impossible
		{4, 2},
		{5, 10},
		{6, 4},
		{7, 40},
		{8, 92},
	}

	for _, tt := range tests {
		got := CountNQueens(tt.n)
		if got != tt.want {
			t.Errorf("CountNQueens(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		n         int
		wantCount int
	}{
		{4, 2},
		{5, 10},
		{6, 4},
	}

	for _, tt := range tests {
		solutions := SolveNQueens(tt.n)
		if len(solutions) != tt.wantCount {
			t.Errorf("SolveNQueens(%d) returned %d solutions, want %d", tt.n, len(solutions), tt.wantCount)
		}
	}
}
