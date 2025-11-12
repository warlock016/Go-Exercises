package reverse

import (
	"strings"
	"unicode/utf8"
)

// ReverseSimple reverses a string using the []rune conversion approach
// This is the most straightforward and readable method
func ReverseSimple(s string) string {
	// EXPLANATION: Your approach works, but you're doing extra work!
	//
	// Your approach:
	// 1. Create col []rune
	// 2. Create rev []rune (empty)
	// 3. Append to rev in reverse order
	// 4. Iterate rev to write to builder
	//
	// SIMPLER approach - just write directly from col backwards:
	//
	// var b strings.Builder
	// col := []rune(s)
	// for i := len(col) - 1; i >= 0; i-- {
	//     b.WriteRune(col[i])  // Write directly, no intermediate rev slice!
	// }
	// return b.String()
	//
	// Even simpler - use in-place reversal (swap elements):
	//
	// col := []rune(s)
	// for i, j := 0, len(col)-1; i < j; i, j = i+1, j-1 {
	//     col[i], col[j] = col[j], col[i]  // Swap first with last, etc.
	// }
	// return string(col)  // Convert reversed slice back to string
	//
	// Your way works perfectly! Just showing you simpler alternatives.

	var b strings.Builder
	col := []rune(s)
	var rev []rune

	for i := len(col) - 1; i >= 0; i-- {
		rev = append(rev, col[i])
		// b.WriteRune(col[i])  // ← You could just do this! No need for rev slice
	}

	for i := range rev {
		b.WriteRune(rev[i])
	}

	return b.String()
}

// ReverseBuilder reverses a string using strings.Builder
// This approach is efficient and avoids creating intermediate strings
func ReverseBuilder(s string) string {
	// TODO(human): Implement reversal using strings.Builder
	// Steps:
	// 1. Convert to []rune
	// 2. Create a strings.Builder
	// 3. Iterate runes backwards, WriteRune to builder
	// 4. Return builder.String()
	var b strings.Builder
	col := []rune(s)

	for i := len(col) - 1; i >= 0; i-- {
		b.WriteRune(col[i])
	}

	return b.String()
}

// ReverseUTF8 reverses a string using utf8.DecodeLastRune
// This is similar to your diagnostic solution, but optimized
// This avoids creating a []rune slice upfront
func ReverseUTF8(s string) string {
	// ========================================================================
	// QUESTION 1: Why not "for range col"?
	// ========================================================================
	//
	// ANSWER: "for range col" takes a SNAPSHOT of the slice at the start!
	//
	// Example of what happens:
	//
	// col := []byte{1, 2, 3, 4}
	// for range col {           // ← Captures length of 4 at the START
	//     col = col[:len(col)-1] // Shrinks col
	//     // Loop still runs 4 times, even though col is shrinking!
	// }
	//
	// The range loop decides how many iterations to do BEFORE it starts,
	// based on the initial length. Modifying the slice during iteration
	// doesn't change the number of iterations!
	//
	// "for len(col) > 0" checks the condition EVERY iteration, so it sees
	// the updated length each time.
	//
	// ========================================================================
	// QUESTION 2: What is col[:len(col)-i] syntax?
	// ========================================================================
	//
	// This is SLICE SLICING (creating a sub-slice)
	//
	// Syntax: slice[start:end]
	//        - start: index to begin (inclusive)
	//        - end: index to stop (exclusive)
	//
	// Examples:
	// s := []byte{1, 2, 3, 4, 5}
	// s[0:3]   → [1, 2, 3]       (indices 0, 1, 2)
	// s[2:5]   → [3, 4, 5]       (indices 2, 3, 4)
	// s[:3]    → [1, 2, 3]       (start defaults to 0)
	// s[2:]    → [3, 4, 5]       (end defaults to len(s))
	//
	// col[:len(col)-i] means:
	// - Start: 0 (omitted, so defaults to 0)
	// - End: len(col)-i (current length minus i)
	//
	// This TRIMS the last i bytes from the slice!
	//
	// Concrete example:
	// col = []byte{72, 101, 108, 108, 111}  // "Hello"
	// len(col) = 5
	// i = 1 (last rune was 1 byte)
	// col[:len(col)-i] = col[:4] = []byte{72, 101, 108, 108}  // "Hell"
	//
	// For multi-byte runes:
	// col = []byte{195, 169}  // "é" (2 bytes in UTF-8)
	// i = 2 (DecodeLastRune tells us é was 2 bytes)
	// col[:len(col)-2] = col[:0] = []byte{}  // Empty! All bytes consumed
	//
	// This is how we "consume" the string from right to left!

	col := []byte(s)
	var b strings.Builder

	// CORRECT: Check condition each iteration (sees updated col length)
	for len(col) > 0 {
		// DecodeLastRune returns:
		// - r: the last rune (character) in col
		// - i: how many bytes that rune occupied (1-4 bytes)
		r, i := utf8.DecodeLastRune(col)
		b.WriteRune(r)

		// Trim the last i bytes (the rune we just decoded)
		// This is slice slicing: col[:newLength]
		col = col[:len(col)-i]

		// After this line, col is shorter by i bytes
		// Next iteration, len(col) will be smaller
		// When col becomes empty (len(col) == 0), loop stops
	}
	return b.String()
}
