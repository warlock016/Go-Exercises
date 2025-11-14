package frequencycounter

// TopKFrequent returns the k most frequent elements from nums.
// The result can be returned in any order.
//
// Time complexity: O(n + u log u) where n is len(nums), u is unique elements
// Space complexity: O(u)
func TopKFrequent(nums []int, k int) []int {
	// TODO(human): Handle edge cases
	// - If nums is empty or k is 0, return empty slice
	// - Consider: what if k > number of unique elements?

	// TODO(human): Build frequency map
	// Pattern:
	//   freq := make(map[int]int)
	//   for _, num := range nums {
	//       freq[num]++  // Increment count for this number
	//   }
	//
	// Example: [1, 1, 1, 2, 2, 3]
	// Result: freq = {1: 3, 2: 2, 3: 1}

	// TODO(human): Convert map to sortable slice of pairs
	// Pattern:
	//   type pair struct {
	//       num   int  // The number itself
	//       count int  // How many times it appears
	//   }
	//
	//   pairs := make([]pair, 0, len(freq))
	//   for num, count := range freq {
	//       pairs = append(pairs, pair{num, count})
	//   }
	//
	// Why convert to slice?
	// - Maps can't be sorted directly
	// - Need ordered access to find "top k"

	// TODO(human): Sort pairs by frequency (highest first)
	// Pattern:
	//   import "sort"
	//
	//   sort.Slice(pairs, func(i, j int) bool {
	//       return pairs[i].count > pairs[j].count  // > for descending
	//   })
	//
	// After sorting: pairs[0] has highest frequency
	// Think: What does the comparator function return?
	// - true if pairs[i] should come BEFORE pairs[j]
	// - We want highest counts first, so return count[i] > count[j]

	// TODO(human): Extract the top k elements
	// Pattern:
	//   // Ensure k doesn't exceed available elements
	//   if k > len(pairs) {
	//       k = len(pairs)
	//   }
	//
	//   result := make([]int, k)
	//   for i := 0; i < k; i++ {
	//       result[i] = pairs[i].num  // Just the number, not the count
	//   }
	//
	// Why not include counts in result?
	// - Problem only asks for the numbers themselves
	// - Counts were just used for sorting

	// TODO(human): Return the result
	return []int{}
}
