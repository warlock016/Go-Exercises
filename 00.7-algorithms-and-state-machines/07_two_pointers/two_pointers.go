package twopointers

// RemoveDuplicates removes duplicate elements from a sorted array
// Returns a new slice containing only unique elements
func RemoveDuplicates(nums []int) []int {
	// TODO(human): Implement the two-pointer technique
	//
	// Algorithm:
	// 1. Handle edge case: if array is empty, return it
	// 2. Initialize slow pointer to 0 (first element is always unique)
	// 3. Loop with fast pointer from 1 to end:
	//    a. If nums[fast] != nums[slow]:
	//       - We found a new unique element
	//       - Increment slow
	//       - Copy nums[fast] to nums[slow]
	//    b. If they're equal, just move fast forward (skip duplicate)
	// 4. Return nums[:slow+1] (all unique elements)
	//
	// Visualization for [1, 1, 2, 2, 3]:
	//
	// Initial: slow=0, fast=1
	// [1, 1, 2, 2, 3]
	//  s  f
	//
	// nums[fast]=1 == nums[slow]=1 → skip (fast++)
	//
	// fast=2:
	// [1, 1, 2, 2, 3]
	//  s     f
	//
	// nums[fast]=2 != nums[slow]=1 → unique!
	// slow++, nums[slow] = nums[fast]
	// [1, 2, 2, 2, 3]
	//     s  f
	//
	// Continue...
	//
	// Think: Why start fast at 1 instead of 0?
	// Answer: The first element is always unique (no previous element to compare)

	return nil
}

// RemoveDuplicatesKeepTwo keeps at most 2 occurrences of each element
// BONUS CHALLENGE: Only attempt after completing RemoveDuplicates!
func RemoveDuplicatesKeepTwo(nums []int) []int {
	// TODO(human): BONUS - Implement variation that keeps up to 2 duplicates
	//
	// Algorithm modification:
	// - Instead of comparing nums[fast] with nums[slow]
	// - Compare nums[fast] with nums[slow-1]
	// - This allows one duplicate to pass through
	//
	// Example: [1, 1, 1, 2, 2, 3] → [1, 1, 2, 2, 3]
	//
	// Only attempt this after mastering the basic version!

	return nil
}

// MoveZeroes moves all zeros to the end while maintaining relative order
// BONUS CHALLENGE: Another two-pointer application!
func MoveZeroes(nums []int) []int {
	// TODO(human): BONUS - Move all zeros to the end
	//
	// Algorithm:
	// - Use slow pointer to track position for non-zero elements
	// - Fast pointer scans through array
	// - When fast finds non-zero, swap it with position at slow
	//
	// Example: [0, 1, 0, 3, 12] → [1, 3, 12, 0, 0]
	//
	// This is the same pattern as RemoveDuplicates!

	return nil
}
