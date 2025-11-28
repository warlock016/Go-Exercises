package subsets_combinations

// GenerateSubsets generates all subsets of the given slice
func GenerateSubsets(nums []int) [][]int {
	// TODO(human): Implement

	result := make([][]int, 0)
	subsetBacktrack(&result, nums, []int{}, 0)
	// fmt.Printf("Final result: %v\n", result)
	return result
}

func subsetBacktrack(result *[][]int, nums, subset []int, i int) {

	// fmt.Printf("Result: %v, input: %v, subset: %v, idx: %d\n", *result, nums, subset, i)
	if i == len(nums) {
		saved := append([]int{}, subset...)
		*result = append(*result, saved)
		return
	}

	subset = append(subset, nums[i])
	subsetBacktrack(result, nums, subset, i+1)

	subset = subset[:len(subset)-1]
	subsetBacktrack(result, nums, subset, i+1)

}

// Combinations generates all k-length combinations from nums
func Combinations(nums []int, k int) [][]int {
	// TODO(human): Implement
	result := make([][]int, 0)
	combinationHelper(&result, nums, []int{}, k, 0)
	// fmt.Printf("Combinations: %v\n", result)
	return result
}

// [1,2,3] ; k = 2; len = 3
// [1,2], [2,3], [1,3]

// [1,2,3,4] ; k = 2; len = 4
// [1,2], [1,3], [1,4], [2,3], [2,4], [3,4]
func combinationHelper(result *[][]int, nums []int, subset []int, k, start int) {

	if len(subset) == k {
		saved := append([]int{}, subset...)
		*result = append(*result, saved)
		return
	}

	for i := start; i < len(nums); i++ {
		subset = append(subset, nums[i])                // sets the reference element for the current iteration
		combinationHelper(result, nums, subset, k, i+1) // recursive search: sets the start for the nested iteration at start = i+1
		subset = subset[:len(subset)-1]                 // removes the reference element after the previous recursion. This resets the reference element before appending the next reference element during the next iteration.
	}
}

/*
★ Insight ─────────────────────────────────────
  The recursive call happens BETWEEN append and remove!

  The order is:
  1. Append (CHOOSE)
  2. Recurse (EXPLORE) ← This goes deep and returns
  3. Remove (UNCHOOSE) ← This happens AFTER recursion returns
  4. Next iteration starts

  Line 3 executes BEFORE the next iteration of line 1, not after!
  ─────────────────────────────────────────────────

  Your Confusion: Order of Execution

  You expected:
  Iteration 1: append → recurse → [next iteration starts] → append overwrites remove?

  What actually happens:
  Iteration 1: append → recurse (goes deep, returns) → remove → THEN next iteration
  Iteration 2: append → recurse (goes deep, returns) → remove → THEN next iteration
  ...

  Detailed Trace: One Loop Iteration

  Let's trace Combinations([1, 2, 3], k=2) with the loop at start=0:

  for i := start; i < len(nums); i++ {
      subset = append(subset, nums[i])           // Line A: CHOOSE
      combinationHelper(result, nums, subset, k, i+1)  // Line B: EXPLORE
      subset = subset[:len(subset)-1]            // Line C: UNCHOOSE
  }

  First iteration (i=0):

  BEFORE loop: subset = []

  Step 1 (Line A): subset = append([], 1)
                   subset = [1]

  Step 2 (Line B): combinationHelper([1], k=2, start=1)
                   ↓
                   This call goes DEEP:
                   ├─ Appends 2 → [1,2] → saves → removes 2 → [1]
                   ├─ Appends 3 → [1,3] → saves → removes 3 → [1]
                   └─ Returns back here with subset = [1]
                   ↑

  Step 3 (Line C): subset = [1][:len([1])-1]
                   subset = [1][:0]
                   subset = []

  AFTER first iteration: subset = []

  Second iteration (i=1):

  BEFORE: subset = []  (cleaned up by previous iteration!)

  Step 1 (Line A): subset = append([], 2)
                   subset = [2]

  Step 2 (Line B): combinationHelper([2], k=2, start=2)
                   ↓
                   Goes deep, adds 3, saves [2,3], removes 3
                   Returns with subset = [2]
                   ↑

  Step 3 (Line C): subset = [2][:0]
                   subset = []

  AFTER second iteration: subset = []

  Visual Timeline

  Time →

  i=0: [append 1] ──→ [RECURSE...deep...returns] ──→ [remove 1] ──→ END i=0
                            ↓                              ↓
                      subset=[1]                     subset=[]
                      goes deep                      cleaned up!

  i=1: [append 2] ──→ [RECURSE...deep...returns] ──→ [remove 2] ──→ END i=1
                            ↓                              ↓
                      subset=[2]                     subset=[]

  i=2: [append 3] ──→ [RECURSE...returns] ──→ [remove 3] ──→ END LOOP
                            ↓                      ↓
                      subset=[3]               subset=[]

  Understanding subset[:len(subset)-1]

  Yes, you're correct! This removes the last element by creating a shorter slice.

  subset := []int{1, 2, 3}

  // len(subset) = 3
  // len(subset) - 1 = 2

  subset = subset[:2]  // Same as subset[:len(subset)-1]

  // Result: subset = [1, 2]

  Visual:
  Before: [1, 2, 3]
           ↑  ↑  ↑
           0  1  2  (indices)

  subset[:2] means "elements from index 0 up to (but not including) index 2"

  After:  [1, 2]
           ↑  ↑
           0  1

  More examples:
  [1, 2, 3, 4][:3]  →  [1, 2, 3]     // Remove last 1
  [1, 2, 3, 4][:2]  →  [1, 2]       // Remove last 2
  [1][:0]           →  []           // Remove the only element
  [][:0]            →  []           // Empty stays empty

  The Key Insight: Recursion Blocks

  The recursive call blocks until it completes:

  subset = append(subset, nums[i])      // 1. Executes immediately
  combinationHelper(...)                 // 2. BLOCKS here until ALL nested calls return
  subset = subset[:len(subset)-1]       // 3. Only executes AFTER recursion returns

  Think of it like a function call:
  fmt.Println("Before")
  someFunction()  // Program waits here until someFunction() finishes
  fmt.Println("After")  // Only runs after someFunction() returns

  Complete Trace: Combinations([1,2,3], k=2)

  Let me show the EXACT order of every operation:

  combinationHelper([], start=0)
  │
  ├── i=0:
  │   ├── append 1      → subset = [1]
  │   ├── RECURSE with [1], start=1
  │   │   │
  │   │   ├── i=1:
  │   │   │   ├── append 2      → subset = [1,2]
  │   │   │   ├── RECURSE with [1,2], start=2
  │   │   │   │   └── len([1,2]) == 2, SAVE [1,2], return
  │   │   │   └── remove last   → subset = [1]
  │   │   │
  │   │   ├── i=2:
  │   │   │   ├── append 3      → subset = [1,3]
  │   │   │   ├── RECURSE with [1,3], start=3
  │   │   │   │   └── len([1,3]) == 2, SAVE [1,3], return
  │   │   │   └── remove last   → subset = [1]
  │   │   │
  │   │   └── loop ends, return
  │   │
  │   └── remove last   → subset = []     ← Back to first level!
  │
  ├── i=1:
  │   ├── append 2      → subset = [2]
  │   ├── RECURSE with [2], start=2
  │   │   │
  │   │   └── i=2:
  │   │       ├── append 3      → subset = [2,3]
  │   │       ├── RECURSE with [2,3], start=3
  │   │       │   └── len([2,3]) == 2, SAVE [2,3], return
  │   │       └── remove last   → subset = [2]
  │   │
  │   └── remove last   → subset = []
  │
  └── i=2:
      ├── append 3      → subset = [3]
      ├── RECURSE with [3], start=3
      │   └── loop doesn't run (start=3 >= len=3), return
      └── remove last   → subset = []

  Result: [[1,2], [1,3], [2,3]]

  Analogy: Stacking and Unstacking

  Think of it like stacking plates:

  CHOOSE (append):  Put plate on top of stack
  EXPLORE (recurse): Do something with current stack, then...
  UNCHOOSE (remove): Take top plate off before trying next option

  Stack: []
    ↓ append 1
  Stack: [1]
    ↓ recurse (goes deep, adds 2, saves [1,2], removes 2, adds 3, saves [1,3], removes 3)
  Stack: [1]  ← Returns here!
    ↓ remove 1
  Stack: []
    ↓ append 2  ← Next iteration starts with clean stack
  Stack: [2]
    ...

  Summary

  1. Order of execution: append → recurse (blocks until done) → remove → next iteration
  2. subset[:len(subset)-1]: Removes the last element (shortens slice by 1)
  3. Why it works: The recursive call BLOCKS, so the remove happens AFTER all nested work completes, but BEFORE the next loop iteration starts
  4. The pattern:
  for each choice:
      make choice       // modify state
      explore further   // recurse
      undo choice       // restore state ← CRITICAL for backtracking!
*/

// Permutations generates all permutations of nums
func Permutations(nums []int) [][]int {
	// TODO(human): Implement

	result := make([][]int, 0)
	used := make([]bool, len(nums))
	permutationHelper(&result, nums, []int{}, used)
	return result
}

func permutationHelper(result *[][]int, nums, subset []int, used []bool) {

	if len(subset) == len(nums) {
		saved := append([]int{}, subset...)
		*result = append(*result, saved)
	}

	for i := range nums {

		if used[i] {
			continue
		}

		// select reference
		used[i] = true
		subset = append(subset, nums[i])

		// recursive composition
		permutationHelper(result, nums, subset, used)

		// deselect reference
		subset = subset[:len(subset)-1]
		used[i] = false

	}

}
