package memoization

import "testing"

func TestFibonacciMemo(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 1},
		{10, 55},
		{20, 6765},
		{30, 832040},
	}

	for _, tt := range tests {
		got := FibonacciMemo(tt.input)
		if got != tt.want {
			t.Errorf("FibonacciMemo(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestGridPaths(t *testing.T) {
	tests := []struct {
		m, n int
		want int
	}{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 6},
		{3, 4, 10},
	}

	for _, tt := range tests {
		got := GridPaths(tt.m, tt.n)
		if got != tt.want {
			t.Errorf("GridPaths(%d, %d) = %d, want %d", tt.m, tt.n, got, tt.want)
		}
	}
}

func TestClimbStairsMemo(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 7},
		{5, 13},
	}

	for _, tt := range tests {
		got := ClimbStairsMemo(tt.n)
		if got != tt.want {
			t.Errorf("ClimbStairsMemo(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
