package leapyear

// IsLeapYear determines if a given year is a leap year.
//
// Leap year rules:
// - Divisible by 4 → leap year
// - EXCEPT divisible by 100 → not leap year
// - EXCEPT divisible by 400 → leap year
//
// Examples:
//
//	IsLeapYear(2000) → true  (divisible by 400)
//	IsLeapYear(1900) → false (divisible by 100 but not 400)
//	IsLeapYear(2004) → true  (divisible by 4 but not 100)
//	IsLeapYear(2001) → false (not divisible by 4)
func IsLeapYear(year int) bool {
	// TODO(human): Implement the leap year logic

	if year%4 == 0 {
		if !(year%100 == 0) || year%400 == 0 {
			return true
		}
	}
	return false
}
