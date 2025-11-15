package numberclassifier

import "testing"

func TestClassifyParity(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"positive even", 4, "even"},
		{"positive odd", 7, "odd"},
		{"zero is even", 0, "even"},
		{"negative even", -8, "even"},
		{"negative odd", -3, "odd"},
		{"large even", 1000, "even"},
		{"large odd", 999, "odd"},
		{"one", 1, "odd"},
		{"two", 2, "even"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClassifyParity(tt.input)
			if got != tt.want {
				t.Errorf("ClassifyParity(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestClassifySign(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"positive small", 5, "positive"},
		{"positive large", 1000, "positive"},
		{"negative small", -5, "negative"},
		{"negative large", -1000, "negative"},
		{"zero", 0, "zero"},
		{"one", 1, "positive"},
		{"negative one", -1, "negative"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClassifySign(tt.input)
			if got != tt.want {
				t.Errorf("ClassifySign(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestClassifyMagnitude(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"small zero", 0, "small"},
		{"small positive", 5, "small"},
		{"small boundary", 10, "small"},
		{"medium low", 11, "medium"},
		{"medium mid", 50, "medium"},
		{"medium boundary", 100, "medium"},
		{"large low", 101, "large"},
		{"large high", 1000, "large"},
		{"small negative", -5, "small"},
		{"medium negative", -50, "medium"},
		{"large negative", -500, "large"},
		{"negative boundary 10", -10, "small"},
		{"negative boundary 100", -100, "medium"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClassifyMagnitude(tt.input)
			if got != tt.want {
				t.Errorf("ClassifyMagnitude(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestDayName(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"Monday", 1, "Monday"},
		{"Tuesday", 2, "Tuesday"},
		{"Wednesday", 3, "Wednesday"},
		{"Thursday", 4, "Thursday"},
		{"Friday", 5, "Friday"},
		{"Saturday", 6, "Saturday"},
		{"Sunday", 7, "Sunday"},
		{"invalid zero", 0, "Invalid day"},
		{"invalid eight", 8, "Invalid day"},
		{"invalid negative", -1, "Invalid day"},
		{"invalid large", 100, "Invalid day"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DayName(tt.input)
			if got != tt.want {
				t.Errorf("DayName(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// Benchmark tests to compare performance characteristics
func BenchmarkClassifyParity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClassifyParity(i)
	}
}

func BenchmarkClassifySign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClassifySign(i - b.N/2) // Mix of positive and negative
	}
}

func BenchmarkClassifyMagnitude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClassifyMagnitude(i)
	}
}

func BenchmarkDayName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DayName((i % 7) + 1) // Cycle through valid days
	}
}
