package romannumerals

import "strings"

// ToRoman converts a decimal number (1-3999) to Roman numerals.
//
// Roman numerals use seven symbols:
// I=1, V=5, X=10, L=50, C=100, D=500, M=1000
//
// Special subtraction cases:
// IV=4, IX=9, XL=40, XC=90, CD=400, CM=900
//
// Example:
//   ToRoman(1994) returns "MCMXCIV"
//   (M=1000, CM=900, XC=90, IV=4)
func ToRoman(n int) string {
	var result strings.Builder

	// TODO(human): Create parallel slices for values and symbols
	// values should contain: 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1
	// symbols should contain: "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"

	// TODO(human): Loop through each value-symbol pair
	// For each pair:
	//   - While n >= current value:
	//     - Append the symbol to result
	//     - Subtract the value from n

	return result.String()
}

// FromRoman converts a Roman numeral string to a decimal number.
//
// Handles subtraction cases where a smaller value appears before
// a larger value (e.g., IV=4, IX=9, XL=40, XC=90, CD=400, CM=900).
//
// Example:
//   FromRoman("MCMXCIV") returns 1994
func FromRoman(s string) int {
	// TODO(human): Create a map of rune to int for symbol values
	// Map should contain: I=1, V=5, X=10, L=50, C=100, D=500, M=1000

	// TODO(human): Iterate through the string from RIGHT to LEFT (len(s)-1 down to 0)
	// For each character:
	//   - Get the current value from the map
	//   - If current < prev, subtract it (subtraction case like IV)
	//   - Otherwise, add it
	//   - Update prev to current

	return 0
}
