package romannumerals

import (
	"strings"
)

// ToRoman converts a decimal number (1-3999) to Roman numerals.
//
// Roman numerals use seven symbols:
// I=1, V=5, X=10, L=50, C=100, D=500, M=1000
//
// Special subtraction cases:
// IV=4, IX=9, XL=40, XC=90, CD=400, CM=900
//
// Example:
//
//	ToRoman(1994) returns "MCMXCIV"
//	(M=1000, CM=900, XC=90, IV=4)
func ToRoman(n int) string {
	var result strings.Builder

	if n <= 0 {
		return ""
	}

	remainder := n

	thousand := remainder / 1000
	remainder -= thousand * 1000

	fivehundred := remainder / 500
	remainder -= fivehundred * 500

	hundred := remainder / 100
	remainder -= hundred * 100

	fifty := remainder / 50
	remainder -= fifty * 50

	ten := remainder / 10
	remainder -= ten * 10

	five := remainder / 5
	remainder -= five * 5

	one := remainder / 1
	remainder -= one * 1

	// Roman numerals use seven symbols:
	// I=1, V=5, X=10, L=50, C=100, D=500, M=1000
	//
	// Special subtraction cases:
	// IV=4, IX=9, XL=40, XC=90, CD=400, CM=900

	if thousand != 0 {
		result.WriteString(strings.Repeat("M", thousand))
	}

	if fivehundred == 1 {
		if hundred == 0 {
			result.WriteString("D")
		} else if hundred > 0 && hundred < 4 {
			result.WriteString("D")
			result.WriteString(strings.Repeat("C", hundred))
		} else if hundred == 4 {
			result.WriteString("CM")
		}
	} else {
		if hundred == 4 {
			result.WriteString("CD")
		} else if hundred > 0 && hundred < 4 {
			result.WriteString(strings.Repeat("C", hundred))
		}
	}

	if fifty == 1 {
		if ten == 4 {
			result.WriteString("XC")
		} else if ten > 0 && ten < 4 {
			result.WriteString("L")
			result.WriteString(strings.Repeat("X", ten))
		} else if ten == 0 {
			result.WriteString("L")
		}
	} else {
		if ten == 4 {
			result.WriteString("XL")
		} else if ten > 0 && ten < 4 {
			result.WriteString(strings.Repeat("X", ten))
		}
	}

	if five == 1 {
		if one == 4 {
			result.WriteString("IX")
		} else if one > 0 && one < 4 {
			result.WriteString("V")
			result.WriteString(strings.Repeat("I", one))
		} else if one == 0 {
			result.WriteString("V")
		}
	} else {
		if one == 4 {
			result.WriteString("IV")
		} else if one > 0 && one < 4 {
			result.WriteString(strings.Repeat("I", one))
		}
	}

	return result.String()
}

// FromRoman converts a Roman numeral string to a decimal number.
//
// Handles subtraction cases where a smaller value appears before
// a larger value (e.g., IV=4, IX=9, XL=40, XC=90, CD=400, CM=900).
//
// Example:
//
//	FromRoman("MCMXCIV") returns 1994
func FromRoman(s string) int {
	// TODO(human): Implement Roman numeral to decimal conversion

	runes := []rune(s)
	roman := make(map[rune]int)
	var decimal int

	roman['I'] = 1
	roman['V'] = 5
	roman['X'] = 10
	roman['L'] = 50
	roman['C'] = 100
	roman['D'] = 500
	roman['M'] = 1000

	// L V I I I (len=5)
	for i := 0; i < len(runes); i++ {
		if i+1 < len(runes) {
			if roman[runes[i]] >= roman[runes[i+1]] {
				decimal += roman[runes[i]]
			} else {
				decimal += roman[runes[i+1]] - roman[runes[i]]
				i++
			}
		} else {
			decimal += roman[runes[i]]
		}
	}

	return decimal
}
