package two_dimensional_slices

// CreateMatrix creates a rows x cols matrix initialized with zeros
func CreateMatrix(rows, cols int) [][]int {
	// TODO(human): Create and initialize 2D slice
	return nil
}

// SetMatrixValue sets the value at the specified row and column
// Does nothing if indices are out of bounds
func SetMatrixValue(matrix [][]int, row, col, value int) {
	// TODO(human): Validate bounds and set value
}

// GetMatrixValue returns the value at the specified row and column
// Returns (value, true) if found, (0, false) if out of bounds
func GetMatrixValue(matrix [][]int, row, col int) (int, bool) {
	// TODO(human): Validate bounds and return value
	return 0, false
}

// TransposeMatrix returns a new matrix that is the transpose of the input
// (rows become columns, columns become rows)
func TransposeMatrix(matrix [][]int) [][]int {
	// TODO(human): Swap rows and columns
	return nil
}

// SumMatrix returns the sum of all elements in the matrix
func SumMatrix(matrix [][]int) int {
	// TODO(human): Sum all elements
	return 0
}

// CreateChessboard creates a size x size chessboard with alternating "W" and "B"
// Top-left corner should be "W"
func CreateChessboard(size int) [][]string {
	// TODO(human): Create matrix with alternating pattern
	return nil
}

// FlattenMatrix converts a 2D matrix into a 1D slice (row-major order)
func FlattenMatrix(matrix [][]int) []int {
	// TODO(human): Convert 2D to 1D
	return nil
}
