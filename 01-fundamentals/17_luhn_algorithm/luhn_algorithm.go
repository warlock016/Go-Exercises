package luhnalgorithm

import (
	"strconv"
	"strings"
	"unicode"
)

// IsValidLuhn validates whether the given string is a valid number
// according to the Luhn algorithm (modulus 10 algorithm).
//
// The function should:
//   - Remove any non-digit characters (spaces, hyphens, etc.)
//   - Return false for empty strings or invalid input
//   - Apply the Luhn algorithm: starting from the rightmost digit,
//     double every second digit, subtract 9 if result > 9, sum all digits
//   - Return true if the sum is divisible by 10
//
// Examples:
//
//	IsValidLuhn("4532015112830366") → true
//	IsValidLuhn("1234567812345678") → false
func IsValidLuhn(s string) bool {
	// TODO(human): Implement the Luhn validation algorithm
	lowerCase := strings.ToLower(s)
	digits := []rune{}
	doubled := []int{}
	totalSum := 0

	for _, r := range lowerCase {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		}
	}

	if len(digits) == 0 {
		return false
	}

	for i, r := range digits {
		positionFromRight := len(digits) - 1 - i
		switch positionFromRight%2 == 1 {
		case true:
			a, _ := strconv.Atoi(string(r))
			a *= 2
			if a > 9 {
				prodSum := 0
				numeric := strconv.Itoa(a)

				for _, v := range numeric {
					val, _ := strconv.Atoi(string(v))
					prodSum += val
				}
				doubled = append(doubled, prodSum)
			} else {
				doubled = append(doubled, a)
			}
		case false:
			a, _ := strconv.Atoi(string(r))
			doubled = append(doubled, a)
		}
	}

	for _, v := range doubled {
		totalSum += v
	}

	return totalSum%10 == 0
}

// GenerateCheckDigit calculates the check digit needed to make the given
// number valid according to the Luhn algorithm.
//
// The input string represents a partial number WITHOUT the check digit.
// The function returns the single digit (0-9) that should be appended
// to make the complete number valid.
//
// Examples:
//
//	GenerateCheckDigit("453201511283036") → 6
//	GenerateCheckDigit("123456781234567") → 0
func GenerateCheckDigit(s string) int {
	// TODO(human): Implement check digit generation
	return 0
}
