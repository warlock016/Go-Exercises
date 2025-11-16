package two_dimensional_slices

import (
	"testing"
)

func TestCreateMatrix(t *testing.T) {
	tests := []struct {
		name     string
		rows     int
		cols     int
		wantRows int
		wantCols int
	}{
		{"0x0 matrix", 0, 0, 0, 0},
		{"3x4 matrix", 3, 4, 3, 4},
		{"1x1 matrix", 1, 1, 1, 1},
		{"5x2 matrix", 5, 2, 5, 2},
		{"10x10 matrix", 10, 10, 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateMatrix(tt.rows, tt.cols)
			if len(got) != tt.wantRows {
				t.Errorf("CreateMatrix(%d, %d) rows = %d, want %d", tt.rows, tt.cols, len(got), tt.wantRows)
			}
			for i, row := range got {
				if len(row) != tt.wantCols {
					t.Errorf("CreateMatrix(%d, %d)[%d] cols = %d, want %d", tt.rows, tt.cols, i, len(row), tt.wantCols)
				}
				// Check all values are zero
				for j, val := range row {
					if val != 0 {
						t.Errorf("CreateMatrix(%d, %d)[%d][%d] = %d, want 0", tt.rows, tt.cols, i, j, val)
					}
				}
			}
		})
	}
}

func TestSetMatrixValue(t *testing.T) {
	tests := []struct {
		name     string
		rows     int
		cols     int
		setRow   int
		setCol   int
		setValue int
		wantSet  bool
	}{
		{"set valid position", 3, 3, 1, 1, 42, true},
		{"set corner", 3, 3, 0, 0, 99, true},
		{"set last position", 3, 3, 2, 2, 7, true},
		{"row out of bounds", 3, 3, 5, 1, 10, false},
		{"col out of bounds", 3, 3, 1, 5, 10, false},
		{"negative row", 3, 3, -1, 1, 10, false},
		{"negative col", 3, 3, 1, -1, 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := CreateMatrix(tt.rows, tt.cols)
			SetMatrixValue(matrix, tt.setRow, tt.setCol, tt.setValue)

			if tt.wantSet {
				got, ok := GetMatrixValue(matrix, tt.setRow, tt.setCol)
				if !ok {
					t.Errorf("SetMatrixValue(%d, %d, %d) position not accessible", tt.setRow, tt.setCol, tt.setValue)
				}
				if got != tt.setValue {
					t.Errorf("SetMatrixValue(%d, %d, %d) = %d, want %d", tt.setRow, tt.setCol, tt.setValue, got, tt.setValue)
				}
			}
		})
	}
}

func TestGetMatrixValue(t *testing.T) {
	matrix := CreateMatrix(3, 3)
	SetMatrixValue(matrix, 1, 2, 42)
	SetMatrixValue(matrix, 0, 0, 10)
	SetMatrixValue(matrix, 2, 2, 99)

	tests := []struct {
		name    string
		row     int
		col     int
		wantVal int
		wantOk  bool
	}{
		{"get set value", 1, 2, 42, true},
		{"get corner", 0, 0, 10, true},
		{"get last", 2, 2, 99, true},
		{"get zero value", 0, 1, 0, true},
		{"row out of bounds", 5, 1, 0, false},
		{"col out of bounds", 1, 5, 0, false},
		{"negative row", -1, 1, 0, false},
		{"negative col", 1, -1, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := GetMatrixValue(matrix, tt.row, tt.col)
			if ok != tt.wantOk {
				t.Errorf("GetMatrixValue(%d, %d) ok = %v, want %v", tt.row, tt.col, ok, tt.wantOk)
			}
			if ok && got != tt.wantVal {
				t.Errorf("GetMatrixValue(%d, %d) = %d, want %d", tt.row, tt.col, got, tt.wantVal)
			}
		})
	}
}

func TestTransposeMatrix(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			"2x3 matrix",
			[][]int{{1, 2, 3}, {4, 5, 6}},
			[][]int{{1, 4}, {2, 5}, {3, 6}},
		},
		{
			"3x2 matrix",
			[][]int{{1, 2}, {3, 4}, {5, 6}},
			[][]int{{1, 3, 5}, {2, 4, 6}},
		},
		{
			"1x4 matrix",
			[][]int{{1, 2, 3, 4}},
			[][]int{{1}, {2}, {3}, {4}},
		},
		{
			"4x1 matrix",
			[][]int{{1}, {2}, {3}, {4}},
			[][]int{{1, 2, 3, 4}},
		},
		{
			"2x2 matrix",
			[][]int{{1, 2}, {3, 4}},
			[][]int{{1, 3}, {2, 4}},
		},
		{
			"empty matrix",
			[][]int{},
			[][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TransposeMatrix(tt.input)
			if len(got) != len(tt.want) {
				t.Fatalf("TransposeMatrix() rows = %d, want %d", len(got), len(tt.want))
			}
			for i := range tt.want {
				if len(got[i]) != len(tt.want[i]) {
					t.Fatalf("TransposeMatrix()[%d] cols = %d, want %d", i, len(got[i]), len(tt.want[i]))
				}
				for j := range tt.want[i] {
					if got[i][j] != tt.want[i][j] {
						t.Errorf("TransposeMatrix()[%d][%d] = %d, want %d", i, j, got[i][j], tt.want[i][j])
					}
				}
			}
		})
	}
}

func TestSumMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{"empty matrix", [][]int{}, 0},
		{"single element", [][]int{{5}}, 5},
		{"single row", [][]int{{1, 2, 3, 4}}, 10},
		{"single column", [][]int{{1}, {2}, {3}}, 6},
		{"2x2 matrix", [][]int{{1, 2}, {3, 4}}, 10},
		{"3x3 matrix", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 45},
		{"with zeros", [][]int{{0, 1}, {2, 0}, {0, 3}}, 6},
		{"with negatives", [][]int{{-1, 2}, {3, -4}}, 0},
		{"all zeros", [][]int{{0, 0}, {0, 0}}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumMatrix(tt.matrix)
			if got != tt.want {
				t.Errorf("SumMatrix(%v) = %d, want %d", tt.matrix, got, tt.want)
			}
		})
	}
}

func TestCreateChessboard(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"1x1 board", 1},
		{"2x2 board", 2},
		{"3x3 board", 3},
		{"8x8 board", 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateChessboard(tt.size)
			if len(got) != tt.size {
				t.Fatalf("CreateChessboard(%d) rows = %d, want %d", tt.size, len(got), tt.size)
			}

			for i := range got {
				if len(got[i]) != tt.size {
					t.Fatalf("CreateChessboard(%d)[%d] cols = %d, want %d", tt.size, i, len(got[i]), tt.size)
				}

				for j := range got[i] {
					expected := "W"
					if (i+j)%2 != 0 {
						expected = "B"
					}
					if got[i][j] != expected {
						t.Errorf("CreateChessboard(%d)[%d][%d] = %q, want %q", tt.size, i, j, got[i][j], expected)
					}
				}
			}

			// Check top-left is "W"
			if tt.size > 0 && got[0][0] != "W" {
				t.Errorf("CreateChessboard(%d)[0][0] = %q, want %q (top-left must be W)", tt.size, got[0][0], "W")
			}
		})
	}
}

func TestFlattenMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{"empty matrix", [][]int{}, []int{}},
		{"single element", [][]int{{5}}, []int{5}},
		{"single row", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"single column", [][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
		{"2x3 matrix", [][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{"3x2 matrix", [][]int{{1, 2}, {3, 4}, {5, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{"jagged matrix", [][]int{{1}, {2, 3}, {4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{"with zeros", [][]int{{0, 1}, {2, 0}}, []int{0, 1, 2, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenMatrix(tt.matrix)
			if len(got) != len(tt.want) {
				t.Fatalf("FlattenMatrix(%v) length = %d, want %d", tt.matrix, len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("FlattenMatrix(%v)[%d] = %d, want %d", tt.matrix, i, got[i], tt.want[i])
				}
			}
		})
	}
}
