package two_dimensional_slices

// CreateMatrix creates a rows x cols matrix initialized with zeros
func CreateMatrix(rows, cols int) [][]int {
	// TODO(human): Create and initialize 2D slice
	matrix := make([][]int, rows)

	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	return matrix
}

// SetMatrixValue sets the value at the specified row and column
// Does nothing if indices are out of bounds
func SetMatrixValue(matrix [][]int, row, col, value int) {
	// TODO(human): Validate bounds and set value

	if row < 0 || col < 0 { // early return if invalid row, col indexes
		return
	}

	for r := range matrix {
		if row == r {
			for c := range matrix[r] {
				if col == c {
					matrix[r][c] = value
				}
			}
		}
	}
}

// GetMatrixValue returns the value at the specified row and column
// Returns (value, true) if found, (0, false) if out of bounds
func GetMatrixValue(matrix [][]int, row, col int) (int, bool) {
	// TODO(human): Validate bounds and return value

	if row < 0 || col < 0 { // early return if invalid row, col indexes
		return 0, false
	}

	for r := range matrix {
		if row == r {
			for c := range matrix[r] {
				if col == c {
					return matrix[r][c], true
				}
			}
		}
	}

	return 0, false
}

// TransposeMatrix returns a new matrix that is the transpose of the input
// (rows become columns, columns become rows)

// 3x4 matrix
// [1 2 3 4]
// [5 6 7 8]
// [9 10 11 12]

// ...becomes 4x3 matrix
// [9 5 1]
// [10 6 2]
// [11 7 3]
// [12 8 4]

// (0/0) -> (0/2),
// (0/1) -> (1/2),
// (0/2) -> (2/2),
// (0/3) -> (3/2),

// (1/0) -> (0/1),
// (1/1) -> (1/1),
// (1/2) -> (2/1),
// (1/3) -> (3/1),

// (2/0) -> (0/0),
// (2/1) -> (1/0),
// (2/2) -> (2/0),
// (2/3) -> (3/0),

// 1, 2, 3, 4
// 5, 6, 7, 8
// 9,10,11,12

func TransposeMatrix(matrix [][]int) [][]int {
	// TODO(human): Swap rows and columns

	if len(matrix) == 0 {
		return nil
	}

	rows := len(matrix)
	var cols int

	for r := range matrix {
		cols = max(len(matrix[r]), cols)
	}

	transposed := make([][]int, cols)

	for r := range transposed {
		transposed[r] = make([]int, rows)
	}

	for r := len(matrix) - 1; r >= 0; r-- {
		for c := range matrix[r] {
			transposed[c][r] = matrix[r][c]
		}
	}

	return transposed
}

// SumMatrix returns the sum of all elements in the matrix
func SumMatrix(matrix [][]int) int {
	// TODO(human): Sum all elements

	var sum int

	for r := range matrix { // r is the row index and NOT the value!
		for c := range matrix[r] { // c is the column index at row r, NOT the value!
			sum += matrix[r][c]
		}
	}

	return sum
}

// CreateChessboard creates a size x size chessboard with alternating "W" and "B"
// Top-left corner should be "W"
func CreateChessboard(size int) [][]string {
	// TODO(human): Create matrix with alternating pattern
	// [0][0] should be "W"

	if size < 1 {
		return nil
	}

	// this section creates a nxn matrix
	result := make([][]string, size)
	for i := range size {
		result[i] = make([]string, size)
	}

	for r := range result {
		for c := range result[r] {
			if (r+c)%2 == 0 {
				result[r][c] = "W"
			} else {
				result[r][c] = "B"
			}
		}
	}

	return result
}

// FlattenMatrix converts a 2D matrix into a 1D slice (row-major order)
func FlattenMatrix(matrix [][]int) []int {
	// TODO(human): Convert 2D to 1D

	result := make([]int, 0)

	for r := range matrix {
		result = append(result, matrix[r]...)
	}

	return result
}
