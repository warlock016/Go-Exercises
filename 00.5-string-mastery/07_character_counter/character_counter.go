package charactercounter

import (
	"strings"
	"unicode"
)

// CountRunes counts the frequency of each rune in the string.
// It returns a map where keys are runes and values are their counts.
//
// Examples:
//
//	CountRunes("hello") → map[h:1 e:1 l:2 o:1]
//	CountRunes("aabbcc") → map[a:2 b:2 c:2]
//	CountRunes("") → map[] (empty map)
func CountRunes(s string) map[rune]int {
	// TODO(human): Implement rune counting
	//
	// APPROACH:
	// 1. Create a map[rune]int to store counts
	// 2. Iterate through each rune in the string
	// 3. Increment the count for that rune
	//
	// KEY INSIGHT:
	// In Go, accessing a non-existent map key returns the zero value.
	// For int, the zero value is 0.
	// So counts[r]++ works even on the first occurrence!
	//
	// PSEUDOCODE:
	//   counts := make(map[rune]int)
	//   for _, r := range s {
	//       counts[r]++  // Auto-initializes to 0 if not present
	//   }
	//   return counts

	// counts := make(map[rune]int)
	counts := map[rune]int{}

	for _, v := range s {
		counts[v]++
	}
	// Your code here

	return counts
}

// MostCommon returns the most frequently occurring rune in the string.
// If there is a tie, it returns the rune that appears first in the string.
// If the string is empty, it returns the zero value for rune (0 or '\x00').
//
// Examples:
//
//	MostCommon("hello") → 'l' (appears 2 times)
//	MostCommon("aabbcc") → 'a' (tie: all appear twice, 'a' is first)
//	MostCommon("mississippi") → 'i' (appears 4 times).   idx [0,1,2,3,4,5,6,7,8,9,0], [0:1, 1:4, 2:4, 3:2]
//	MostCommon("") → '\x00' (zero value)
func MostCommon(s string) rune {
	// TODO(human): Implement finding the most common rune
	//
	// CHALLENGE:
	// Maps are unordered! To respect "first occurrence" in case of ties,
	// you need to iterate through the original string while tracking what
	// you've already seen.
	//
	// APPROACH:
	// 1. Get the counts using CountRunes
	// 2. Iterate through the original string (to preserve order)
	// 3. Track the rune with max count
	// 4. Use a "seen" map to avoid checking the same rune twice
	//
	// PSEUDOCODE:
	//   if s == "" {
	//       return 0  // Zero value for rune
	//   }
	//
	//   counts := CountRunes(s)
	//   seen := make(map[rune]bool)
	//   var mostCommon rune
	//   maxCount := 0
	//
	//   for _, r := range s {
	//       if !seen[r] {  // Haven't checked this rune yet
	//           if counts[r] > maxCount {
	//               maxCount = counts[r]
	//               mostCommon = r
	//           }
	//           seen[r] = true
	//       }
	//   }
	//
	//   return mostCommon

	// Your code here
	// we need to check which element map[r] has the highest count && the lowest r
	// counts := make(map[rune]int) // example: mississippi -> [m:1, i:4, s:4, p:2] or [s:4, p:2, i:4, m:1] or [s:4, p:2, m:1, i:4] ... -> 4! == 24 permutations
	counts := map[rune]int{}

	for _, r := range s {
		counts[r]++
	} // this loop can be omitted by reusing the previous function CountRunes(s)

	// counts := CountRunes(s)

	seen := make(map[rune]bool)
	var mostCommon rune
	var highestCount int

	for _, r := range s { // we iterate over the string runes to preserve order or runes, which map does not guarantee!
		if !seen[r] {

			if counts[r] > highestCount /* && r < mostCommon */ {
				highestCount = counts[r]
				mostCommon = r
			}
			seen[r] = true // can be before or after the inner if conditional
		}
	}

	return mostCommon
}

// IsAnagram checks if two strings are anagrams of each other.
// An anagram is a word formed by rearranging the letters of another.
// This function is case-insensitive and ignores spaces.
//
// Examples:
//
//	IsAnagram("listen", "silent") → true
//	IsAnagram("Hello", "hello") → true (case insensitive)
//	IsAnagram("Astronomer", "Moon starer") → true (ignore spaces)
//	IsAnagram("hello", "world") → false
func IsAnagram(s1, s2 string) bool {
	// TODO(human): Implement anagram detection
	//
	// APPROACH:
	// Two strings are anagrams if they contain the same characters
	// with the same frequencies (ignoring case and spaces).
	//
	// STEPS:
	// 1. Normalize both strings (lowercase + remove spaces)
	// 2. Count runes in both normalized strings
	// 3. Compare the two count maps
	//
	// NORMALIZING:
	// Create a helper function or inline logic:
	//   - Convert to lowercase: unicode.ToLower()
	//   - Skip spaces: unicode.IsSpace()
	//   - Build result with strings.Builder
	//
	// COMPARING MAPS:
	// Two maps are equal if:
	//   - They have the same number of keys
	//   - Each key in map1 exists in map2 with the same value
	//
	// PSEUDOCODE:
	//   normalize := func(s string) string {
	//       var builder strings.Builder
	//       for _, r := range s {
	//           if !unicode.IsSpace(r) {
	//               builder.WriteRune(unicode.ToLower(r))
	//           }
	//       }
	//       return builder.String()
	//   }
	//
	//   n1 := normalize(s1)
	//   n2 := normalize(s2)
	//
	//   counts1 := CountRunes(n1)
	//   counts2 := CountRunes(n2)
	//
	//   if len(counts1) != len(counts2) {
	//       return false
	//   }
	//
	//   for r, count := range counts1 {
	//       if counts2[r] != count {
	//           return false
	//       }
	//   }
	//
	//   return true

	_ = strings.Builder{} // Hint: use this for normalization
	_ = unicode.ToLower   // Hint: use this for case conversion
	_ = unicode.IsSpace   // Hint: use this to detect spaces

	// Your code here

	// var a strings.Builder
	// var a []rune
	a := make(map[rune]int)
	for _, r := range s1 {
		if unicode.IsLetter(r) {
			// a.WriteRune(unicode.ToLower(r))
			a[unicode.ToLower(r)]++
		}
	}

	// var b strings.Builder
	for _, r := range s2 {
		if unicode.IsLetter(r) {
			a[unicode.ToLower(r)]--
		}
	}

	for _, v := range a {
		if v != 0 {
			return false
		}
	}

	return true
}
