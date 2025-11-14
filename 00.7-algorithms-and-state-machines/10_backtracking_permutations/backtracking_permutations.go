package backtrackingpermutations

// Permute generates all permutations of the input string.
// For a string of length n, returns n! permutations.
//
// Time complexity: O(n × n!) - must generate all permutations
// Space complexity: O(n) recursion depth + O(n × n!) for results
func Permute(s string) []string {
	// TODO(human): Initialize results slice and convert string to runes
	// Pattern:
	//   results := []string{}
	//   runes := []rune(s)
	//
	// Why []rune?
	// - Strings are immutable in Go
	// - Need mutable slice for swapping
	// - Runes handle Unicode correctly

	// TODO(human): Call backtrack helper function
	// Pattern:
	//   backtrack(runes, 0, &results)
	//
	// Why pass pointer to results?
	// - Avoid copying entire slice on each recursive call
	// - All recursive calls share same results slice

	// TODO(human): Return the results
	return []string{}
}

// backtrack is the recursive helper that generates permutations using swapping.
//
// Algorithm:
// - Base case: when start == len(runes), we've fixed all positions
// - Recursive case: try each remaining character at position 'start'
//   1. Swap runes[start] with runes[i] (choose)
//   2. Recurse with start+1 (explore)
//   3. Swap back (un-choose/backtrack)
func backtrack(runes []rune, start int, results *[]string) {
	// TODO(human): Implement base case
	// Pattern:
	//   if start == len(runes) {
	//       *results = append(*results, string(runes))
	//       return
	//   }
	//
	// When start == len(runes):
	// - We've made choices for all positions [0...n-1]
	// - Current runes slice is a complete permutation
	// - Add it to results and return

	// TODO(human): Implement recursive case with backtracking
	// Pattern:
	//   for i := start; i < len(runes); i++ {
	//       // CHOOSE: Try runes[i] at position 'start'
	//       runes[start], runes[i] = runes[i], runes[start]
	//
	//       // EXPLORE: Fix remaining positions recursively
	//       backtrack(runes, start+1, results)
	//
	//       // UN-CHOOSE: Restore original state (backtrack)
	//       runes[start], runes[i] = runes[i], runes[start]
	//   }
	//
	// Visualization for "abc":
	// start=0: try 'a', 'b', 'c' at position 0
	//   start=1 with "abc": try 'b', 'c' at position 1
	//     start=2 with "abc": base case → add "abc"
	//     start=2 with "acb": base case → add "acb"
	//   start=1 with "bac": try 'a', 'c' at position 1
	//     start=2 with "bac": base case → add "bac"
	//     start=2 with "bca": base case → add "bca"
	//   ... and so on
	//
	// Why swap back?
	// - The for loop tries different values at position 'start'
	// - Each iteration needs the original state
	// - Swapping back restores the array for the next iteration
	//
	// Example trace for "ab":
	// backtrack("ab", 0):
	//   i=0: swap(0,0)="ab" → backtrack("ab",1)
	//          backtrack("ab",1): i=1: swap(1,1)="ab" → backtrack("ab",2)
	//            backtrack("ab",2): base case, add "ab"
	//        swap back (0,0)="ab"
	//   i=1: swap(0,1)="ba" → backtrack("ba",1)
	//          backtrack("ba",1): i=1: swap(1,1)="ba" → backtrack("ba",2)
	//            backtrack("ba",2): base case, add "ba"
	//        swap back (0,1)="ab"
}
