package gradeevaluator

import (
	"math"
	"testing"
)

func TestLetterGrade(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		// A grades
		{"perfect score", 100, "A"},
		{"high A", 95, "A"},
		{"lowest A", 90, "A"},

		// B grades
		{"boundary B to A", 89, "B"},
		{"mid B", 85, "B"},
		{"lowest B", 80, "B"},

		// C grades
		{"boundary C to B", 79, "C"},
		{"mid C", 75, "C"},
		{"lowest C", 70, "C"},

		// D grades
		{"boundary D to C", 69, "D"},
		{"mid D", 65, "D"},
		{"lowest D (passing)", 60, "D"},

		// F grades
		{"boundary F to D", 59, "F"},
		{"mid F", 30, "F"},
		{"zero score", 0, "F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LetterGrade(tt.score)
			if got != tt.want {
				t.Errorf("LetterGrade(%d) = %q, want %q", tt.score, got, tt.want)
			}
		})
	}
}

func TestIsPassing(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  bool
	}{
		{"perfect score passes", 100, true},
		{"high score passes", 85, true},
		{"minimum passing", 60, true},
		{"one below passing", 59, false},
		{"low failing", 30, false},
		{"zero fails", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPassing(tt.score)
			if got != tt.want {
				t.Errorf("IsPassing(%d) = %v, want %v", tt.score, got, tt.want)
			}
		})
	}
}

func TestGradeWithPlus(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		// A+ grades
		{"perfect A+", 100, "A+"},
		{"high A+", 98, "A+"},
		{"lowest A+", 97, "A+"},

		// A grades
		{"boundary A to A+", 96, "A"},
		{"mid A", 94, "A"},
		{"lowest A", 93, "A"},

		// A- grades
		{"boundary A- to A", 92, "A-"},
		{"mid A-", 91, "A-"},
		{"lowest A-", 90, "A-"},

		// B+ grades
		{"boundary B+ to A-", 89, "B+"},
		{"mid B+", 88, "B+"},
		{"lowest B+", 87, "B+"},

		// B grades
		{"boundary B to B+", 86, "B"},
		{"mid B", 84, "B"},
		{"lowest B", 83, "B"},

		// B- grades
		{"boundary B- to B", 82, "B-"},
		{"mid B-", 81, "B-"},
		{"lowest B-", 80, "B-"},

		// C+ grades
		{"boundary C+ to B-", 79, "C+"},
		{"mid C+", 78, "C+"},
		{"lowest C+", 77, "C+"},

		// C grades
		{"boundary C to C+", 76, "C"},
		{"mid C", 74, "C"},
		{"lowest C", 73, "C"},

		// C- grades
		{"boundary C- to C", 72, "C-"},
		{"mid C-", 71, "C-"},
		{"lowest C-", 70, "C-"},

		// D+ grades
		{"boundary D+ to C-", 69, "D+"},
		{"mid D+", 68, "D+"},
		{"lowest D+", 67, "D+"},

		// D grades
		{"boundary D to D+", 66, "D"},
		{"mid D", 64, "D"},
		{"lowest D", 63, "D"},

		// D- grades
		{"boundary D- to D", 62, "D-"},
		{"mid D-", 61, "D-"},
		{"lowest D- (passing)", 60, "D-"},

		// F grades
		{"boundary F to D-", 59, "F"},
		{"mid F", 30, "F"},
		{"zero F", 0, "F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GradeWithPlus(tt.score)
			if got != tt.want {
				t.Errorf("GradeWithPlus(%d) = %q, want %q", tt.score, got, tt.want)
			}
		})
	}
}

func TestClassAverage(t *testing.T) {
	tests := []struct {
		name        string
		scores      []int
		wantAverage float64
		wantGrade   string
	}{
		{
			name:        "empty class",
			scores:      []int{},
			wantAverage: 0.0,
			wantGrade:   "F",
		},
		{
			name:        "single perfect score",
			scores:      []int{100},
			wantAverage: 100.0,
			wantGrade:   "A",
		},
		{
			name:        "all A students",
			scores:      []int{100, 95, 90},
			wantAverage: 95.0,
			wantGrade:   "A",
		},
		{
			name:        "mixed grades averaging B",
			scores:      []int{90, 85, 78, 92, 88},
			wantAverage: 86.6,
			wantGrade:   "B",
		},
		{
			name:        "mixed grades averaging C",
			scores:      []int{70, 75, 80, 65},
			wantAverage: 72.5,
			wantGrade:   "C",
		},
		{
			name:        "barely passing class",
			scores:      []int{60, 65, 70},
			wantAverage: 65.0,
			wantGrade:   "D",
		},
		{
			name:        "failing class",
			scores:      []int{45, 50, 55},
			wantAverage: 50.0,
			wantGrade:   "F",
		},
		{
			name:        "boundary case - rounds to A",
			scores:      []int{89, 90, 91},
			wantAverage: 90.0,
			wantGrade:   "A",
		},
		{
			name:        "boundary case - rounds to B",
			scores:      []int{88, 89, 90},
			wantAverage: 89.0,
			wantGrade:   "B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAverage, gotGrade := ClassAverage(tt.scores)

			// Compare floats with small epsilon for floating point precision
			if math.Abs(gotAverage-tt.wantAverage) > 0.01 {
				t.Errorf("ClassAverage(%v) average = %.2f, want %.2f", tt.scores, gotAverage, tt.wantAverage)
			}

			if gotGrade != tt.wantGrade {
				t.Errorf("ClassAverage(%v) grade = %q, want %q", tt.scores, gotGrade, tt.wantGrade)
			}
		})
	}
}

// Test edge cases and validation
func TestGradeEvaluatorEdgeCases(t *testing.T) {
	t.Run("negative scores", func(t *testing.T) {
		// While negative scores shouldn't happen in practice,
		// they should be treated as failing grades
		grade := LetterGrade(-10)
		if grade != "F" {
			t.Errorf("LetterGrade(-10) = %q, want %q", grade, "F")
		}

		passing := IsPassing(-10)
		if passing {
			t.Errorf("IsPassing(-10) = %v, want false", passing)
		}
	})

	t.Run("scores above 100", func(t *testing.T) {
		// Extra credit might push scores above 100
		grade := LetterGrade(105)
		if grade != "A" {
			t.Errorf("LetterGrade(105) = %q, want %q", grade, "A")
		}

		gradeWithPlus := GradeWithPlus(105)
		if gradeWithPlus != "A+" {
			t.Errorf("GradeWithPlus(105) = %q, want %q", gradeWithPlus, "A+")
		}
	})

	t.Run("class with single student", func(t *testing.T) {
		avg, grade := ClassAverage([]int{85})
		if avg != 85.0 || grade != "B" {
			t.Errorf("ClassAverage([85]) = (%.1f, %q), want (85.0, %q)", avg, grade, "B")
		}
	})

	t.Run("large class", func(t *testing.T) {
		// Test with many students
		scores := make([]int, 100)
		for i := range scores {
			scores[i] = 80 // All B students
		}
		avg, grade := ClassAverage(scores)
		if avg != 80.0 || grade != "B" {
			t.Errorf("ClassAverage(100 students with 80) = (%.1f, %q), want (80.0, %q)", avg, grade, "B")
		}
	})
}
