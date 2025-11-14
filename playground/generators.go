package main

import "math/rand"

/*
Key Concepts

1. rand.Int() vs rand.Intn(n)

- rand.Int() - Returns [0, math.MaxInt] which is 9,223,372,036,854,775,807 on 64-bit systems
- Use for: Random values when you want huge numbers
- ⚠️ Never use for lengths/counts - will create massive slices
- rand.Intn(n) - Returns [0, n) (0 up to but not including n)
- Use for: Bounded ranges, slice lengths, array indices
- Example: rand.Intn(255) gives [0, 254]

2. Generating Positive and Negative Values

Method 1: Shift the range (generators.go:29-31)
value := rand.Intn(255) - 127  // Gives [-127, 128)
- Subtract to shift: rand.Intn(200) - 100 gives [-100, 100)
- Predictable, bounded range

Method 2: Random negation (generators.go:43-49)
value := rand.Int()
if rand.Intn(2) == 0 {  // 50% chance
	value = -value
}
- Creates very large positive/negative values
- Unbounded range (±9 quintillion)
- rand.Intn(2) returns either 0 or 1 (coin flip)

3. Performance Tip (generators.go:6-7)

If you know the final size upfront, use:
slice := make([]int, 0, randomLength)  // Pre-allocate capacity
This avoids multiple memory reallocations during append.
*/

func IntSliceGenerator() []int {
	// Using make() with pre-allocated capacity avoids reallocations during append
	// if we know the final size in advance (avoids multiple allocations during append)
	randomLength := rand.Intn(254) + 1
	slice := make([]int, 0, randomLength)

	// rand.Intn(254) returns [0, 254), adding 1 gives us [1, 254]
	// This ensures we always get at least 1 element and cap the max at 254
	// Using rand.Int() here would risk creating slices with billions of elements

	for range randomLength {
		// rand.Intn(255) returns values in [0, 255) - only non-negative integers
		// This generates values from 0 to 254 inclusive
		slice = append(slice, rand.Intn(255))
	}
	return slice
}

// IntSliceGeneratorWithNegatives generates a slice with both positive and negative values
func IntSliceGeneratorWithNegatives() []int {
	randomLength := rand.Intn(254) + 1
	slice := make([]int, 0, randomLength)

	for range randomLength {
		// Method 1: Generate in range [-127, 127]
		// rand.Intn(255) gives [0, 255), subtracting 127 shifts to [-127, 128)
		value := rand.Intn(255) - 127
		slice = append(slice, value)
	}
	return slice
}

// IntSliceGeneratorFullRange generates a slice with potentially very large positive/negative values
func IntSliceGeneratorFullRange() []int {
	randomLength := rand.Intn(254) + 1
	slice := make([]int, 0, randomLength)

	for range randomLength {
		// Method 2: Use rand.Int() and randomly negate it
		// rand.Int() returns only positive values (0 to math.MaxInt)
		value := rand.Int()

		// Randomly make it negative (50% chance)
		if rand.Intn(2) == 0 {
			value = -value
		}

		slice = append(slice, value)
	}
	return slice
}
