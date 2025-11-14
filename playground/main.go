package main

import "fmt"

func main() {
	// Original: only positive values [0, 254]
	positiveSlice := IntSliceGenerator()
	fmt.Printf("Positive only (length %d): %v\n\n", len(positiveSlice), positiveSlice)

	// With negatives: range [-127, 127]
	mixedSlice := IntSliceGeneratorWithNegatives()
	fmt.Printf("Mixed range (length %d): %v\n\n", len(mixedSlice), mixedSlice)

	// Full range: very large positive/negative values
	fullRangeSlice := IntSliceGeneratorFullRange()
	fmt.Printf("Full range (length %d): %v\n", len(fullRangeSlice), fullRangeSlice)
}
