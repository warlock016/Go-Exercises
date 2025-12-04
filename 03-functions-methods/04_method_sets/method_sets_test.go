package method_sets

import (
	"math"
	"strings"
	"testing"
)

func TestTemperatureCelsius(t *testing.T) {
	tests := []struct {
		name string
		temp Temperature
		want float64
	}{
		{"zero", Temperature(0), 0},
		{"positive", Temperature(25), 25},
		{"negative", Temperature(-10), -10},
		{"decimal", Temperature(36.6), 36.6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.temp.Celsius()
			if math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("Temperature(%f).Celsius() = %f, want %f", tt.temp, got, tt.want)
			}
		})
	}
}

func TestTemperatureFahrenheit(t *testing.T) {
	tests := []struct {
		name string
		temp Temperature
		want float64
	}{
		{"zero celsius", Temperature(0), 32},
		{"room temp", Temperature(25), 77},
		{"body temp", Temperature(37), 98.6},
		{"negative", Temperature(-40), -40}, // same in C and F
		{"boiling", Temperature(100), 212},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.temp.Fahrenheit()
			if math.Abs(got-tt.want) > 0.1 {
				t.Errorf("Temperature(%f).Fahrenheit() = %f, want %f", tt.temp, got, tt.want)
			}
		})
	}
}

func TestTemperatureKelvin(t *testing.T) {
	tests := []struct {
		name string
		temp Temperature
		want float64
	}{
		{"absolute zero", Temperature(-273.15), 0},
		{"zero celsius", Temperature(0), 273.15},
		{"room temp", Temperature(25), 298.15},
		{"boiling", Temperature(100), 373.15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.temp.Kelvin()
			if math.Abs(got-tt.want) > 0.01 {
				t.Errorf("Temperature(%f).Kelvin() = %f, want %f", tt.temp, got, tt.want)
			}
		})
	}
}

func TestStringListJoin(t *testing.T) {
	tests := []struct {
		name      string
		list      StringList
		separator string
		want      string
	}{
		{"empty", StringList{}, ", ", ""},
		{"single", StringList{"apple"}, ", ", "apple"},
		{"multiple", StringList{"apple", "banana", "cherry"}, ", ", "apple, banana, cherry"},
		{"space separator", StringList{"Go", "is", "fun"}, " ", "Go is fun"},
		{"no separator", StringList{"a", "b", "c"}, "", "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.Join(tt.separator)
			if got != tt.want {
				t.Errorf("StringList.Join(%q) = %q, want %q", tt.separator, got, tt.want)
			}
		})
	}
}

func TestStringListContains(t *testing.T) {
	list := StringList{"apple", "banana", "cherry"}

	tests := []struct {
		name string
		item string
		want bool
	}{
		{"present at start", "apple", true},
		{"present in middle", "banana", true},
		{"present at end", "cherry", true},
		{"not present", "grape", false},
		{"empty string", "", false},
		{"case sensitive", "Apple", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := list.Contains(tt.item)
			if got != tt.want {
				t.Errorf("StringList.Contains(%q) = %v, want %v", tt.item, got, tt.want)
			}
		})
	}

	// Test empty list
	emptyList := StringList{}
	if emptyList.Contains("anything") {
		t.Errorf("Empty StringList.Contains() should return false")
	}
}

func TestStringListFilter(t *testing.T) {
	list := StringList{"apple", "banana", "cherry", "date"}

	tests := []struct {
		name      string
		predicate func(string) bool
		want      StringList
	}{
		{
			"length greater than 5",
			func(s string) bool { return len(s) > 5 },
			StringList{"banana", "cherry"},
		},
		{
			"starts with c",
			func(s string) bool { return strings.HasPrefix(s, "c") },
			StringList{"cherry"},
		},
		{
			"contains a",
			func(s string) bool { return strings.Contains(s, "a") },
			StringList{"apple", "banana", "date"},
		},
		{
			"none match",
			func(s string) bool { return len(s) > 10 },
			StringList{},
		},
		{
			"all match",
			func(s string) bool { return true },
			StringList{"apple", "banana", "cherry", "date"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := list.Filter(tt.predicate)
			if len(got) != len(tt.want) {
				t.Errorf("StringList.Filter() length = %d, want %d", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("StringList.Filter()[%d] = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestStringListAppend(t *testing.T) {
	list := StringList{"apple", "banana"}

	list.Append("cherry")
	if len(list) != 3 || list[2] != "cherry" {
		t.Errorf("After Append(cherry), list = %v, want 3 elements with cherry at end", list)
	}

	list.Append("date", "elderberry")
	if len(list) != 5 {
		t.Errorf("After Append(date, elderberry), length = %d, want 5", len(list))
	}
	if list[3] != "date" || list[4] != "elderberry" {
		t.Errorf("After Append, last two elements = %v, %v, want date, elderberry", list[3], list[4])
	}
}

func TestBytesString(t *testing.T) {
	tests := []struct {
		name  string
		bytes Bytes
		want  string
	}{
		{"zero", Bytes(0), "0 B"},
		{"small", Bytes(512), "512 B"},
		{"1 KB", Bytes(1024), "1.00 KB"},
		{"5 KB", Bytes(5120), "5.00 KB"},
		{"1 MB", Bytes(1024 * 1024), "1.00 MB"},
		{"5 MB", Bytes(5 * 1024 * 1024), "5.00 MB"},
		{"1 GB", Bytes(1024 * 1024 * 1024), "1.00 GB"},
		{"2.5 GB", Bytes(2560 * 1024 * 1024), "2.50 GB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bytes.String()
			if got != tt.want {
				t.Errorf("Bytes(%d).String() = %q, want %q", tt.bytes, got, tt.want)
			}
		})
	}
}

func TestBytesKilobytes(t *testing.T) {
	tests := []struct {
		name  string
		bytes Bytes
		want  float64
	}{
		{"zero", Bytes(0), 0},
		{"1 KB", Bytes(1024), 1},
		{"5 KB", Bytes(5120), 5},
		{"half KB", Bytes(512), 0.5},
		{"1 MB", Bytes(1024 * 1024), 1024},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bytes.Kilobytes()
			if math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("Bytes(%d).Kilobytes() = %f, want %f", tt.bytes, got, tt.want)
			}
		})
	}
}

func TestBytesMegabytes(t *testing.T) {
	tests := []struct {
		name  string
		bytes Bytes
		want  float64
	}{
		{"zero", Bytes(0), 0},
		{"1 MB", Bytes(1024 * 1024), 1},
		{"5 MB", Bytes(5 * 1024 * 1024), 5},
		{"half MB", Bytes(512 * 1024), 0.5},
		{"1 GB", Bytes(1024 * 1024 * 1024), 1024},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bytes.Megabytes()
			if math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("Bytes(%d).Megabytes() = %f, want %f", tt.bytes, got, tt.want)
			}
		})
	}
}

func TestBytesGigabytes(t *testing.T) {
	tests := []struct {
		name  string
		bytes Bytes
		want  float64
	}{
		{"zero", Bytes(0), 0},
		{"1 GB", Bytes(1024 * 1024 * 1024), 1},
		{"2 GB", Bytes(2 * 1024 * 1024 * 1024), 2},
		{"half GB", Bytes(512 * 1024 * 1024), 0.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bytes.Gigabytes()
			if math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("Bytes(%d).Gigabytes() = %f, want %f", tt.bytes, got, tt.want)
			}
		})
	}
}
