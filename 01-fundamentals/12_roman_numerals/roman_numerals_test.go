package romannumerals

import "testing"

func TestToRoman(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"One", 1, "I"},
		{"Four (subtraction case)", 4, "IV"},
		{"Nine (subtraction case)", 9, "IX"},
		{"Fifty-Eight", 58, "LVIII"},
		{"Nineteen Ninety-Four", 1994, "MCMXCIV"},
		{"Maximum value", 3999, "MMMCMXCIX"},
		{"Five", 5, "V"},
		{"Ten", 10, "X"},
		{"Forty (subtraction case)", 40, "XL"},
		{"Ninety (subtraction case)", 90, "XC"},
		{"Four Hundred (subtraction case)", 400, "CD"},
		{"Nine Hundred (subtraction case)", 900, "CM"},
		{"Twenty-Seven", 27, "XXVII"},
		{"Forty-Eight", 48, "XLVIII"},
		{"Fifty-Nine", 59, "LIX"},
		{"Ninety-Three", 93, "XCIII"},
		{"One Hundred Forty-One", 141, "CXLI"},
		{"One Hundred Sixty-Three", 163, "CLXIII"},
		{"Four Hundred Two", 402, "CDII"},
		{"Five Hundred Seventy-Five", 575, "DLXXV"},
		{"Nine Hundred Eleven", 911, "CMXI"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToRoman(tt.input)
			if got != tt.want {
				t.Errorf("ToRoman(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestFromRoman(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"One", "I", 1},
		{"Four (subtraction case)", "IV", 4},
		{"Nine (subtraction case)", "IX", 9},
		{"Fifty-Eight", "LVIII", 58},
		{"Nineteen Ninety-Four", "MCMXCIV", 1994},
		{"Maximum value", "MMMCMXCIX", 3999},
		{"Five", "V", 5},
		{"Ten", "X", 10},
		{"Forty (subtraction case)", "XL", 40},
		{"Ninety (subtraction case)", "XC", 90},
		{"Four Hundred (subtraction case)", "CD", 400},
		{"Nine Hundred (subtraction case)", "CM", 900},
		{"Twenty-Seven", "XXVII", 27},
		{"Forty-Eight", "XLVIII", 48},
		{"Fifty-Nine", "LIX", 59},
		{"Ninety-Three", "XCIII", 93},
		{"One Hundred Forty-One", "CXLI", 141},
		{"One Hundred Sixty-Three", "CLXIII", 163},
		{"Four Hundred Two", "CDII", 402},
		{"Five Hundred Seventy-Five", "DLXXV", 575},
		{"Nine Hundred Eleven", "CMXI", 911},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromRoman(tt.input)
			if got != tt.want {
				t.Errorf("FromRoman(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

// TestRoundTrip verifies that converting to Roman and back yields the original number
func TestRoundTrip(t *testing.T) {
	tests := []int{1, 4, 9, 27, 48, 59, 93, 141, 163, 402, 575, 911, 1994, 3999}

	for _, n := range tests {
		roman := ToRoman(n)
		got := FromRoman(roman)
		if got != n {
			t.Errorf("Round trip failed for %d: ToRoman=%q, FromRoman=%d", n, roman, got)
		}
	}
}

// BenchmarkToRoman measures performance of decimal to Roman conversion
func BenchmarkToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToRoman(1994)
	}
}

// BenchmarkFromRoman measures performance of Roman to decimal conversion
func BenchmarkFromRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FromRoman("MCMXCIV")
	}
}
