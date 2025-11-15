package gradeevaluator

// LetterGrade converts a numeric score (0-100) to a letter grade.
//
// Grading scale:
//   - A: 90-100
//   - B: 80-89
//   - C: 70-79
//   - D: 60-69
//   - F: 0-59
//
// Example:
//
//	LetterGrade(95)  // "A"
//	LetterGrade(73)  // "C"
//	LetterGrade(42)  // "F"
func LetterGrade(score int) string {
	// TODO(human): Implement letter grade conversion
	// Hint: Use if/else-if chain, starting with highest grade
	// if score >= 90 {
	//     return "A"
	// } else if ...
	return ""
}

// IsPassing returns true if the score is a passing grade (>= 60).
//
// Example:
//
//	IsPassing(85)  // true
//	IsPassing(60)  // true
//	IsPassing(59)  // false
func IsPassing(score int) bool {
	// TODO(human): Implement pass/fail check
	// Hint: Single comparison is enough
	return false
}

// GradeWithPlus converts a numeric score to a detailed letter grade with
// plus/minus modifiers (e.g., A+, A, A-, B+, etc.).
//
// Grading scale:
//   - A+: 97-100
//   - A:  93-96
//   - A-: 90-92
//   - B+: 87-89
//   - B:  83-86
//   - B-: 80-82
//   - C+: 77-79
//   - C:  73-76
//   - C-: 70-72
//   - D+: 67-69
//   - D:  63-66
//   - D-: 60-62
//   - F:  0-59 (no F+ or F-)
//
// Example:
//
//	GradeWithPlus(98)  // "A+"
//	GradeWithPlus(92)  // "A"
//	GradeWithPlus(88)  // "B+"
//	GradeWithPlus(55)  // "F"
func GradeWithPlus(score int) string {
	// TODO(human): Implement detailed grading with plus/minus
	// Hint: Long if/else-if chain, check highest grades first
	// if score >= 97 {
	//     return "A+"
	// } else if score >= 93 {
	//     return "A"
	// } else if ...
	return ""
}

// ClassAverage calculates the average score of a class and returns both
// the numeric average and the corresponding letter grade.
//
// Returns (0.0, "F") if the scores slice is empty.
//
// Example:
//
//	ClassAverage([]int{90, 85, 78, 92, 88})  // (86.6, "B")
//	ClassAverage([]int{100, 95, 90})         // (95.0, "A")
//	ClassAverage([]int{})                    // (0.0, "F")
func ClassAverage(scores []int) (average float64, grade string) {
	// TODO(human): Implement class average calculation
	// Hint:
	// 1. Check if slice is empty first
	// 2. Sum all scores using a for loop
	// 3. Convert to float64 before dividing: float64(sum) / float64(len(scores))
	// 4. Use LetterGrade to convert average to letter grade
	// 5. Return both values
	return 0.0, "F"
}
