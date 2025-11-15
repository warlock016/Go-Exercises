package leapyear

import "testing"

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name string
		year int
		want bool
	}{
		// Century years divisible by 400 (leap years)
		{"year 1600 is a leap year", 1600, true},
		{"year 2000 is a leap year", 2000, true},

		// Century years not divisible by 400 (not leap years)
		{"year 1700 is not a leap year", 1700, false},
		{"year 1800 is not a leap year", 1800, false},
		{"year 1900 is not a leap year", 1900, false},
		{"year 2100 is not a leap year", 2100, false},

		// Non-century years divisible by 4 (leap years)
		{"year 2004 is a leap year", 2004, true},
		{"year 2020 is a leap year", 2020, true},
		{"year 1996 is a leap year", 1996, true},

		// Years not divisible by 4 (not leap years)
		{"year 2001 is not a leap year", 2001, false},
		{"year 2019 is not a leap year", 2019, false},
		{"year 1999 is not a leap year", 1999, false},

		// Edge cases
		{"year 4 is a leap year", 4, true},
		{"year 1 is not a leap year", 1, false},
		{"year 100 is not a leap year", 100, false},
		{"year 400 is a leap year", 400, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsLeapYear(tt.year)
			if got != tt.want {
				t.Errorf("IsLeapYear(%d) = %v, want %v", tt.year, got, tt.want)
			}
		})
	}
}

func TestIsLeapYearHistoricalDates(t *testing.T) {
	tests := []struct {
		name string
		year int
		want bool
	}{
		{"Gregorian calendar adoption year 1582", 1582, false},
		{"year 1984 (divisible by 4)", 1984, true},
		{"year 2024 (current decade)", 2024, true},
		{"year 2400 (future century divisible by 400)", 2400, true},
		{"year 2500 (future century not divisible by 400)", 2500, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsLeapYear(tt.year)
			if got != tt.want {
				t.Errorf("IsLeapYear(%d) = %v, want %v", tt.year, got, tt.want)
			}
		})
	}
}

// BenchmarkIsLeapYear measures performance of leap year calculation
func BenchmarkIsLeapYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsLeapYear(2000)
	}
}
