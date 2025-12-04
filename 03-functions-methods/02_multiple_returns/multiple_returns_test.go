package multiple_returns

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		want      float64
		wantError bool
	}{
		{"simple division", 10, 2, 5.0, false},
		{"division with remainder", 7, 3, 7.0 / 3.0, false},
		{"divide by zero", 5, 0, 0, true},
		{"zero divided by number", 0, 5, 0.0, false},
		{"negative numbers", -10, 2, -5.0, false},
		{"both negative", -10, -2, 5.0, false},
		{"fractional result", 1, 3, 1.0 / 3.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if tt.wantError {
				if err == nil {
					t.Errorf("Divide(%f, %f) expected error, got nil", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%f, %f) unexpected error: %v", tt.a, tt.b, err)
				}
				if math.Abs(got-tt.want) > 0.0001 {
					t.Errorf("Divide(%f, %f) = %f, want %f", tt.a, tt.b, got, tt.want)
				}
			}
		})
	}
}

func TestParseName(t *testing.T) {
	tests := []struct {
		name      string
		fullName  string
		wantFirst string
		wantLast  string
		wantError bool
	}{
		{"simple name", "John Doe", "John", "Doe", false},
		{"three part name", "Mary Jane Watson", "Mary Jane", "Watson", false},
		{"single name", "Alice", "", "", true},
		{"empty name", "", "", "", true},
		{"whitespace only", "   ", "", "", true},
		{"leading/trailing spaces", "  John Doe  ", "John", "Doe", false},
		{"multiple spaces between", "John  Doe", "John", "Doe", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirst, gotLast, err := ParseName(tt.fullName)
			if tt.wantError {
				if err == nil {
					t.Errorf("ParseName(%q) expected error, got nil", tt.fullName)
				}
			} else {
				if err != nil {
					t.Errorf("ParseName(%q) unexpected error: %v", tt.fullName, err)
				}
				if gotFirst != tt.wantFirst || gotLast != tt.wantLast {
					t.Errorf("ParseName(%q) = (%q, %q), want (%q, %q)",
						tt.fullName, gotFirst, gotLast, tt.wantFirst, tt.wantLast)
				}
			}
		})
	}
}

func TestStats(t *testing.T) {
	tests := []struct {
		name      string
		numbers   []int
		wantMin   int
		wantMax   int
		wantAvg   float64
		wantError bool
	}{
		{"normal case", []int{5, 2, 9, 1, 7}, 1, 9, 4.8, false},
		{"all negative", []int{-5, -10, -2}, -10, -2, -5.666666666666667, false},
		{"single element", []int{42}, 42, 42, 42.0, false},
		{"empty slice", []int{}, 0, 0, 0.0, true},
		{"all same", []int{5, 5, 5}, 5, 5, 5.0, false},
		{"two elements", []int{10, 20}, 10, 20, 15.0, false},
		{"mixed positive negative", []int{-5, 0, 5}, -5, 5, 0.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax, gotAvg, err := Stats(tt.numbers)
			if tt.wantError {
				if err == nil {
					t.Errorf("Stats(%v) expected error, got nil", tt.numbers)
				}
			} else {
				if err != nil {
					t.Errorf("Stats(%v) unexpected error: %v", tt.numbers, err)
				}
				if gotMin != tt.wantMin || gotMax != tt.wantMax {
					t.Errorf("Stats(%v) min/max = (%d, %d), want (%d, %d)",
						tt.numbers, gotMin, gotMax, tt.wantMin, tt.wantMax)
				}
				if math.Abs(gotAvg-tt.wantAvg) > 0.0001 {
					t.Errorf("Stats(%v) avg = %f, want %f",
						tt.numbers, gotAvg, tt.wantAvg)
				}
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		target    int
		wantIndex int
		wantFound bool
	}{
		{"found in middle", []int{1, 2, 3}, 2, 1, true},
		{"not found", []int{10, 20, 30}, 40, 0, false},
		{"found at start", []int{5, 10, 15}, 5, 0, true},
		{"found at end", []int{5, 10, 15}, 15, 2, true},
		{"single element found", []int{42}, 42, 0, true},
		{"single element not found", []int{42}, 99, 0, false},
		{"empty slice", []int{}, 1, 0, false},
		{"duplicate elements", []int{1, 2, 2, 3}, 2, 1, true}, // returns first occurrence
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotFound := FindIndex(tt.slice, tt.target)
			if gotFound != tt.wantFound {
				t.Errorf("FindIndex(%v, %d) found = %v, want %v",
					tt.slice, tt.target, gotFound, tt.wantFound)
			}
			if gotFound && gotIndex != tt.wantIndex {
				t.Errorf("FindIndex(%v, %d) index = %d, want %d",
					tt.slice, tt.target, gotIndex, tt.wantIndex)
			}
		})
	}
}
