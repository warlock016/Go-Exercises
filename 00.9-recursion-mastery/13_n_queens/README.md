# Exercise 13: N-Queens

**Tier:** 4 - Advanced
**Estimated Time:** 90 minutes
**Concepts:** Complex backtracking, constraint satisfaction, 2D board recursion

## Learning Goal

Solve the classic N-Queens problem - place N queens on an N×N chessboard so no two queens attack each other.

## The N-Queens Problem

Place N queens on an N×N board such that:
- No two queens share the same row
- No two queens share the same column  
- No two queens share the same diagonal

## Examples

```go
SolveNQueens(4)
// [
//   [".Q..","...Q","Q...","..Q."],
//   ["..Q.","Q...","...Q",".Q.."]
// ]

CountNQueens(4) // 2
CountNQueens(8) // 92
```

## Approach

1. Place one queen per row (avoids row conflicts)
2. For each column in current row:
   - Check if safe (no column/diagonal conflicts)
   - Place queen (CHOOSE)
   - Recurse to next row (EXPLORE)
   - Remove queen (UNCHOOSE)

## This Is It!

This exercise combines everything:
- Backtracking (choose-explore-unchoose)
- Constraint checking
- 2D recursion
- Multiple solutions

**Congratulations!** Complete this and you've mastered recursion.
