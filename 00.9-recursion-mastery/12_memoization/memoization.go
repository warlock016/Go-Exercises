package memoization

// FibonacciMemo calculates nth Fibonacci number with memoization
func FibonacciMemo(n int) int {
	// TODO(human): Implement
	memo := make(map[int]int)
	memo[0], memo[1], memo[2] = 0, 1, 1
	return fibHelper(n, memo)
}

func fibHelper(n int, memo map[int]int) int {
	if result, exists := memo[n]; exists {
		return result
	}

	result := fibHelper(n-1, memo) + fibHelper(n-2, memo)

	memo[n] = result
	return result
}

/*
  The Core Question

  GridPaths(3, 3) asks: "How many unique paths exist from top-left to bottom-right in a 3×3 grid, moving only right (→) or down (↓)?"

  Start
    ↓
  ┌───┬───┬───┐
  │ S │   │   │   Row 1
  ├───┼───┼───┤
  │   │   │   │   Row 2
  ├───┼───┼───┤
  │   │   │ E │   Row 3
  └───┴───┴───┘
                ↑
               End

  ---
  Why Your Initial Approach Felt Natural

  You wanted to:
  1. Start at S
  2. Step right or down
  3. Track each path until reaching E
  4. Count all paths

  This simulation approach works but is inefficient. You'd literally walk every path.

  ---
  The Recursive Insight: Work Backwards

  Instead of asking "how do I get FROM start TO end?", ask:

  "How many ways can I ARRIVE at any cell?"

  For any cell (m, n), you can only arrive from:
  - The cell above it (m-1, n) - you came DOWN
  - The cell to the left (m, n-1) - you came RIGHT

            ┌─────────┐
            │ (m-1,n) │  ← came from here (moved DOWN)
            │         │
            └────┬────┘
                 │ ↓
  ┌─────────┬────▼────┐
  │ (m,n-1) │→ (m,n)  │  ← current cell
  │         │         │
  └─────────┴─────────┘
       ↑
  came from here (moved RIGHT)

  So: paths to (m,n) = paths to (m-1,n) + paths to (m,n-1)

  ---
  Concrete Example: 3×3 Grid

  Let me fill in the number of paths to reach each cell:

  ┌───┬───┬───┐
  │ 1 │ 1 │ 1 │   ← Top row: only 1 way (all rights →→)
  ├───┼───┼───┤
  │ 1 │ 2 │ 3 │   ← Middle: 2 = 1+1, 3 = 1+2
  ├───┼───┼───┤
  │ 1 │ 3 │ 6 │   ← Bottom: 3 = 1+2, 6 = 3+3
  └───┴───┴───┘

  Reading the grid:
  - Cell (1,1): 1 way (you start here)
  - Cell (1,3): 1 way (→→)
  - Cell (3,1): 1 way (↓↓)
  - Cell (2,2): 2 ways = paths from above (1) + paths from left (1)
  - Cell (3,3): 6 ways = paths from above (3) + paths from left (3)

  ---
  Why Edge Rows/Columns = 1

  If you're in the first row (m == 1), you can ONLY move right. There's exactly 1 path.
  If you're in the first column (n == 1), you can ONLY move down. There's exactly 1 path.

  if m == 1 || n == 1 {
      return 1  // Only one way to traverse a single row or column
  }

  ---
  Trace Through Your Code

  gridHelper(3, 3, memo)

  1. m=3, n=3: Not base case, not in memo
  2. Calculate: gridHelper(2, 3) + gridHelper(3, 2)

  gridHelper(2, 3):
    → gridHelper(1, 3) + gridHelper(2, 2)
    → 1 + gridHelper(2, 2)
    → 1 + [gridHelper(1, 2) + gridHelper(2, 1)]
    → 1 + [1 + 1]
    → 1 + 2 = 3

  gridHelper(3, 2):
    → gridHelper(2, 2) + gridHelper(3, 1)
    → 2 (from memo!) + 1
    → 3

  gridHelper(3, 3) = 3 + 3 = 6

  ---
  The Mental Shift

  | Your Initial Thinking      | Recursive Counting                       |
  |----------------------------|------------------------------------------|
  | Simulate walking each path | Count paths mathematically               |
  | Track visited cells        | No tracking needed                       |
  | Store path strings "RRDD"  | Just store numbers                       |
  | Forward: start → end       | Backward: "how many ways to reach here?" |

  ---
*/

// GridPaths counts paths from top-left to bottom-right in m×n grid (can only move right or down)
func GridPaths(m, n int) int {
	// TODO(human): Implement
	memo := make(map[[2]int]int)

	return gridHelper(m, n, memo)
}

func gridHelper(m, n int, memo map[[2]int]int) int {
	if m == 1 || n == 1 {
		return 1
	}

	key := [2]int{m, n}
	if result, exists := memo[key]; exists {
		return result
	}

	result := gridHelper(m-1, n, memo) + gridHelper(m, n-1, memo)
	memo[key] = result
	return result
}

// ClimbStairsMemo counts ways to climb n stairs (1, 2, or 3 steps at a time) with memoization
func ClimbStairsMemo(n int) int {
	// TODO(human): Implement
	memo := make(map[int]int)
	return stairsHelper(n, memo)
}

func stairsHelper(n int, memo map[int]int) int {

	if n < 0 {
		return 0
	}

	if n == 0 || n == 1 {
		return 1
	}

	if result, exists := memo[n]; exists {
		return result
	}

	result := stairsHelper(n-1, memo) + stairsHelper(n-2, memo) + stairsHelper(n-3, memo)
	memo[n] = result

	return result
}
