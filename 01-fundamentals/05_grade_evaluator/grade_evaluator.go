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

	var grade string

	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}
	return grade
}

// IsPassing returns true if the score is a passing grade (>= 60).
//
// Example:
//
//	IsPassing(85)  // true
//	IsPassing(60)  // true
//	IsPassing(59)  // false
func IsPassing(score int) bool {

	return score >= 60
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

func LetterGradeMod(score int) (string, int) {
	// TODO(human): Implement letter grade conversion

	var grade string
	var remainder int

	if score >= 90 {
		grade = "A"
		remainder = score % 90
	} else if score >= 80 {
		grade = "B"
		remainder = score % 80
	} else if score >= 70 {
		grade = "C"
		remainder = score % 70
	} else if score >= 60 {
		grade = "D"
		remainder = score % 60
	} else {
		grade = "F"
		remainder = 0
	}
	return grade, remainder
}

func GradeWithPlus(score int) string {
	grade, mod := LetterGradeMod(score)

	if mod >= 7 && grade != "F" {
		grade += "+"
	} else if mod < 3 && grade != "F" {
		grade += "-"
	}

	return grade
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

	var cumScore float64

	if len(scores) == 0 {
		return 0, "F"
	}

	for _, v := range scores {
		cumScore += float64(v)
	}

	avgScore := cumScore / float64(len(scores))

	return avgScore, LetterGrade(int(avgScore))
}
