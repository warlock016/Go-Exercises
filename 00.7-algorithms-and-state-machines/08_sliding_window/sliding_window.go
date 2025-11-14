package slidingwindow

// MaxSumSubarray finds the maximum sum of any contiguous subarray of length k
// using the sliding window technique.
//
// Time complexity: O(n) where n is len(nums)
// Space complexity: O(1)
func MaxSumSubarray(nums []int, k int) int {
	// TODO(human): Handle edge cases
	// - If nums is empty or k is invalid (k <= 0 or k > len(nums)), return 0
	// - If k equals len(nums), there's only one window
	//
	// Think: What makes k "invalid"?

	// TODO(human): Calculate the sum of the first window
	// Pattern:
	//   windowSum := 0
	//   for i := 0; i < k; i++ {
	//       windowSum += nums[i]
	//   }
	//   maxSum := windowSum  // Initialize with first window's sum
	//
	// Why initialize maxSum with first window?
	// - It's a valid window that exists
	// - Handles case where all numbers are negative

	// TODO(human): Slide the window across the array
	// Pattern:
	//   for i := k; i < len(nums); i++ {
	//       // Remove element leaving window: nums[i-k]
	//       // Add element entering window: nums[i]
	//       windowSum = windowSum + nums[i] - nums[i-k]
	//
	//       // Update maximum if current window is larger
	//       if windowSum > maxSum {
	//           maxSum = windowSum
	//       }
	//   }
	//
	// Visualization of sliding:
	// [2, 1, 5, 1, 3, 2], k=3
	//  ^-----^              windowSum = 2+1+5 = 8, maxSum = 8
	//     ^-----^           windowSum = 8+1-2 = 7, maxSum = 8
	//        ^-----^        windowSum = 7+3-1 = 9, maxSum = 9
	//           ^-----^     windowSum = 9+2-5 = 6, maxSum = 9

	// TODO(human): Return the maximum sum found
	return 0
}
