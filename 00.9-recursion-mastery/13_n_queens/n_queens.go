package n_queens

import (
	"fmt"
)

/*
1. Only one queen per row
2. Only one queen per column
3. Only one queen per diagional

Description:
1. add a queen at the first possible position in first row
2. determine the available column positions for all additional rows
3. add a new queen at the first possible column position in next row
4. recurse through steps 2 + 3 and check if number of available positions == pending rows
4a. if 4 == true, then place next queen via 3
4b. else backtrack to step??? (unknown how to backtrack)

[0, Q, 0, 0]
[0, 0, 0, Q]
[Q, 0, 0, 0]
[0, 0, Q, 0]
*/

// unresolved questions:
// appropriate data structure to track progress: map, slice, ...?? --> in theory, we only need to track the position of each queen for each row. Therefore a map[int]int could suffice (map[row]queen_pos)
// number of required backtracking steps to determine initial valid position

// ideas:
// we place the first queen at each possible position and iterate from there until we reach the end, then we check if generated configuration is valid; if true then add it to the result, else discard.
// for checking valid positions at next row, we take previous position "i" and mark i+1 and i-1 as invalid, as long as i > 0 || i < n, where 0 and n are row boundaries.

// corrections by Claude:
// 1. []int is simpler and more appropriate than map[int]int: we can track queens[row] = column (position)
// 2. use backtracking pattern from previous exercises:
// -> for each possible column in current row:
// 	  	CHOOSE: place queen
//    	EXPLORE: recurse to next row
//    	UNCHOOSE: remove queen
// 3. improve valid position algorithm since current approach only validates next row but does not consider entire diagonal path

// diagonal paths:
// notes: i: row index, j: column index
// [0,0; 1,1; 2,2; 3,3; 4,4; 5,5; 6,6; 7,7] : (i: 0->n, j: 0->n)
// [7,0; 6,1; 5,2; 4,3; 3,4; 2,5; 1,6; 0,7] : (i: n->0, j: 0->n)

// SolveNQueens finds all solutions to the N-Queens problem
// 3D slice allows to store valid board configs in inner 2D slices,
// where each valid config takes one position in the outer slice
func SolveNQueens(n int) [][][]string {
	// TODO(human): Implement

	result := make([][][]string, 0)
	tracker := []int{} // queen[row] = column
	usedCols := make(map[int]bool)
	usedDiag1 := make(map[int]bool)
	usedDiag2 := make(map[int]bool)
	queenHelper(n, tracker, usedCols, usedDiag1, usedDiag2, &result)
	return result
}

func queenHelper(n int, tracker []int, usedCols, usedDiag1, usedDiag2 map[int]bool, result *[][][]string) {
	if len(tracker) == n {
		// fmt.Printf("end reached: %v\n", tracker)
		newBoard := boardCreator(n, tracker)
		*result = append(*result, newBoard)
		return
	}

	// fmt.Printf("board size: %d, tracker: %v, result: %v\n", n, tracker, result)
	for i := range n {
		if validPos(i, usedCols, usedDiag1, usedDiag2, tracker) {
			usedCols[i] = true
			usedDiag1[i-len(tracker)] = true
			usedDiag2[i+len(tracker)] = true
			tracker = append(tracker, i)

			queenHelper(n, tracker, usedCols, usedDiag1, usedDiag2, result)
			tracker = tracker[:len(tracker)-1]
			usedCols[i] = false
			usedDiag1[i-len(tracker)] = false
			usedDiag2[i+len(tracker)] = false

		}
	}
}

/*
First Iteration:
0. queens = [0]
[Q...] pos = 0

1. queens = [0, 2]
[Q...]
[..Q.] invalid pos: 0,1; pos = 2; valid pos: 2,3

2. queens = [0,2]
[Q...]
[..Q.]
[....] invalid pos: 0,1,2,3
!!! incomplete set -> prune

backtrack...
1. queens = [0,3]
[Q...]
[...Q]
[.Q..] invalid pos = 0,3,2; valid pos: 1; pos = 1

2. queens = [0,3,1]
[Q...]
[...Q]
[.Q..]
[....] invalid pos = 0,3,1,2

Second iteration:
0. queens = [1]
invalid position calculation:
1-1 = 0
1+1 = 2
1 = 1
--> next = 3; append 3 to queens

1. queens = [1,3]
invalid position calculation:
3-1 = 2
3+1 = out-of-bounds
3 = 3
1 already taken
--> next = 0; append 0

2. queens = [1,3,0]
invalid position calculation:
0-1 = out-of-bounds
0 = 0
0 + 1 = 1 (taken)
3 = 3 taken
--> next = 2; append 2

queens = [1,3,0,2] -> len queens == n --> finish!


queens = [1,3,0,2]
[.Q..] pos = 1
[...Q] invalid pos: 0,1,2; valid pos = 3
[Q...] invalid pos: 1,2,3; valid pos = 0
[..Q.] invalid pos: 0,1,3; valid pos = 2

queen[0] = 0
queen[1] = !0 && !1
queen[2] = !0 && !2
queen[3] = !0 && !3
queen[4] = !0 && !4
queen[5] = !0 && !5
queen[6] = !0 && !6
queen[7] = !0 && !7
*/

func validPos(pos int, usedCols, usedDiag1, usedDiag2 map[int]bool, tracker []int) bool {

	if len(tracker) == 0 { // no queens exist yet
		return true
	}

	if usedCols[pos] {
		return false
	}

	diag1 := pos - len(tracker)
	diag2 := pos + len(tracker)
	if usedDiag1[diag1] || usedDiag2[diag2] {
		return false
	}

	return true
}

func rowCreator(n, pos int) ([]string, error) {

	if pos >= n {
		return nil, fmt.Errorf("queen position out of bounds")
	}

	row := make([]string, 0, n)

	for i := range n {
		if i == pos {
			row = append(row, "Q")
		} else {
			row = append(row, ".")
		}
	}
	return row, nil
}

func boardCreator(n int, board []int) [][]string {

	if n != len(board) {
		panic("error: incomplete board")
	}

	result := make([][]string, 0, n)

	for i, pos := range board {
		row, err := rowCreator(n, pos)
		if err != nil {
			panic(fmt.Sprintf("row :%d, %v", i, err))
		}
		result = append(result, row)
	}

	return result
}

// CountNQueens counts the number of solutions (more efficient)
func CountNQueens(n int) int {
	// TODO(human): Implement
	usedCols := make(map[int]bool)
	usedDiag1 := make(map[int]bool)
	usedDiag2 := make(map[int]bool)
	return countHelper(n, usedCols, usedDiag1, usedDiag2, []int{})
}

func countHelper(n int, usedCols, usedDiag1, usedDiag2 map[int]bool, tracker []int) int {
	if len(tracker) == n {
		return 1
	}

	result := 0

	for i := range n {
		if validPos(i, usedCols, usedDiag1, usedDiag2, tracker) {
			usedDiag1[i-len(tracker)] = true
			usedDiag2[i+len(tracker)] = true
			usedCols[i] = true
			tracker = append(tracker, i)

			result += countHelper(n, usedCols, usedDiag1, usedDiag2, tracker)
			tracker = tracker[:len(tracker)-1]
			usedCols[i] = false
			usedDiag1[i-len(tracker)] = false
			usedDiag2[i+len(tracker)] = false
		}
	}

	return result
}
