# Exercise 09: Two-Dimensional Slices

## 🎯 Learning Goal
Master two-dimensional slices (slices of slices) for representing matrices, grids, and tabular data. Understand how to create, manipulate, and traverse 2D data structures.

## 📝 Problem Description

Two-dimensional slices are essential for representing grids, matrices, game boards, images, and tabular data. Unlike some languages that have built-in 2D arrays, Go uses slices of slices: `[][]T`.

Understanding 2D slices is crucial because:
- Each row can have different lengths (jagged arrays)
- You must initialize both the outer slice and each inner slice
- Access patterns affect performance (row-major vs column-major)
- Many algorithms (image processing, dynamic programming, graphs) rely on 2D structures

In this exercise, you'll implement matrix operations and grid manipulations that demonstrate practical 2D slice usage.

## 🔧 Function Signatures

Implement these functions in `two_dimensional_slices.go`:

```go
// CreateMatrix creates a rows x cols matrix initialized with zeros
func CreateMatrix(rows, cols int) [][]int

// SetMatrixValue sets the value at the specified row and column
// Does nothing if indices are out of bounds
func SetMatrixValue(matrix [][]int, row, col, value int)

// GetMatrixValue returns the value at the specified row and column
// Returns (value, true) if found, (0, false) if out of bounds
func GetMatrixValue(matrix [][]int, row, col int) (int, bool)

// TransposeMatrix returns a new matrix that is the transpose of the input
// (rows become columns, columns become rows)
func TransposeMatrix(matrix [][]int) [][]int

// SumMatrix returns the sum of all elements in the matrix
func SumMatrix(matrix [][]int) int

// CreateChessboard creates a size x size chessboard with alternating "W" and "B"
// Top-left corner should be "W"
func CreateChessboard(size int) [][]string

// FlattenMatrix converts a 2D matrix into a 1D slice (row-major order)
func FlattenMatrix(matrix [][]int) []int
```

## 💡 Examples

```go
// Creating a 3x4 matrix
m := CreateMatrix(3, 4)
// [[0, 0, 0, 0],
//  [0, 0, 0, 0],
//  [0, 0, 0, 0]]

// Setting values
SetMatrixValue(m, 1, 2, 42)
// [[0, 0, 0, 0],
//  [0, 0, 42, 0],
//  [0, 0, 0, 0]]

// Getting values
val, ok := GetMatrixValue(m, 1, 2)  // val=42, ok=true
val, ok = GetMatrixValue(m, 5, 5)   // val=0, ok=false (out of bounds)

// Transposing
original := [][]int{
    {1, 2, 3},
    {4, 5, 6},
}
transposed := TransposeMatrix(original)
// [[1, 4],
//  [2, 5],
//  [3, 6]]

// Summing
matrix := [][]int{{1, 2}, {3, 4}, {5, 6}}
sum := SumMatrix(matrix)  // 21

// Chessboard
board := CreateChessboard(3)
// [["W", "B", "W"],
//  ["B", "W", "B"],
//  ["W", "B", "W"]]

// Flattening
matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
flat := FlattenMatrix(matrix)  // [1, 2, 3, 4, 5, 6]
```

## 📋 Instructions

1. **CreateMatrix:** Create the outer slice with `make([][]int, rows)`, then create each row with `make([]int, cols)`
2. **SetMatrixValue:** Check bounds before setting: `row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[row])`
3. **GetMatrixValue:** Similar bounds checking, return `(0, false)` for out-of-bounds access
4. **TransposeMatrix:** Create a new matrix with swapped dimensions, then copy `result[j][i] = matrix[i][j]`
5. **SumMatrix:** Nested loops - outer loop for rows, inner loop for columns
6. **CreateChessboard:** Pattern is "W" when `(row + col) % 2 == 0`, "B" otherwise
7. **FlattenMatrix:** Calculate total size, create 1D slice, copy all elements in row-major order

## 🧪 Testing

Run tests with:
```bash
go test -v
```

Expected test count: ~35-40 tests across all functions

## 🤔 Think About

1. **Why do we need to initialize each row separately?**
   - `[][]int` is a slice of slices, not a contiguous 2D array
   - Each inner slice is independent and needs its own allocation

2. **What happens if rows have different lengths?**
   - This is valid! Go supports "jagged arrays"
   - You must check `len(matrix[row])` before accessing columns

3. **What's the difference between `matrix[i][j]` and `matrix[j][i]`?**
   - `[i][j]` is row i, column j
   - `[j][i]` is row j, column i (transpose relationship)

4. **How much memory does a 1000x1000 matrix use?**
   - Each int is 8 bytes (on 64-bit systems)
   - 1000 * 1000 * 8 = 8,000,000 bytes = ~7.6 MB
   - Plus overhead for slice headers

## 💡 Hints

<details>
<summary>Hint 1: Creating 2D slices</summary>

Creating a 2D slice requires two steps:

```go
// Step 1: Create outer slice (rows)
matrix := make([][]int, rows)

// Step 2: Create each inner slice (columns)
for i := range matrix {
    matrix[i] = make([]int, cols)
}
```

**Common mistake:** Forgetting step 2 results in a slice of nil slices!
</details>

<details>
<summary>Hint 2: Bounds checking</summary>

Always validate both row and column indices:

```go
// Check if row is valid
if row < 0 || row >= len(matrix) {
    return 0, false
}

// Check if column is valid for this row
if col < 0 || col >= len(matrix[row]) {
    return 0, false
}

// Now safe to access matrix[row][col]
```
</details>

<details>
<summary>Hint 3: Transposing matrices</summary>

Transpose swaps rows and columns:

```go
// Original: rows x cols
// Result: cols x rows

rows := len(matrix)
cols := len(matrix[0])  // Assumes non-empty and rectangular

result := CreateMatrix(cols, rows)

for i := 0; i < rows; i++ {
    for j := 0; j < cols; j++ {
        result[j][i] = matrix[i][j]  // Swap indices
    }
}
```
</details>

<details>
<summary>Hint 4: Nested iteration</summary>

Use nested loops for 2D traversal:

```go
// Iterate through all elements
for i := 0; i < len(matrix); i++ {      // Each row
    for j := 0; j < len(matrix[i]); j++ { // Each column in row
        // Process matrix[i][j]
    }
}

// Or with range
for i, row := range matrix {
    for j, val := range row {
        // Process val at position [i][j]
    }
}
```
</details>

<details>
<summary>Hint 5: Chessboard pattern</summary>

The alternating pattern is based on sum of indices:

```go
for i := 0; i < size; i++ {
    for j := 0; j < size; j++ {
        if (i + j) % 2 == 0 {
            board[i][j] = "W"
        } else {
            board[i][j] = "B"
        }
    }
}
```
</details>

<details>
<summary>Full Solution</summary>

```go
package two_dimensional_slices

func CreateMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

func SetMatrixValue(matrix [][]int, row, col, value int) {
	if row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[row]) {
		matrix[row][col] = value
	}
}

func GetMatrixValue(matrix [][]int, row, col int) (int, bool) {
	if row < 0 || row >= len(matrix) {
		return 0, false
	}
	if col < 0 || col >= len(matrix[row]) {
		return 0, false
	}
	return matrix[row][col], true
}

func TransposeMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}

	rows := len(matrix)
	cols := len(matrix[0])

	result := CreateMatrix(cols, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}

func SumMatrix(matrix [][]int) int {
	sum := 0
	for _, row := range matrix {
		for _, val := range row {
			sum += val
		}
	}
	return sum
}

func CreateChessboard(size int) [][]string {
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				board[i][j] = "W"
			} else {
				board[i][j] = "B"
			}
		}
	}

	return board
}

func FlattenMatrix(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	// Calculate total size
	totalSize := 0
	for _, row := range matrix {
		totalSize += len(row)
	}

	result := make([]int, 0, totalSize)
	for _, row := range matrix {
		result = append(result, row...)
	}

	return result
}
```
</details>

## 🎓 What This Teaches

- **2D slice initialization** - Creating slices of slices requires careful allocation
- **Bounds checking** - Must validate both dimensions before accessing elements
- **Matrix operations** - Transpose, sum, and other fundamental matrix algorithms
- **Nested iteration** - Working with two levels of loops for 2D traversal
- **Memory layout** - Understanding row-major ordering in Go
- **Jagged arrays** - Each row can have different lengths (unlike true 2D arrays)
- **Pattern generation** - Creating structured data like chessboards

---

**Next Exercise:** `10_advanced_maps` - Maps of slices, maps of structs, and nested maps
