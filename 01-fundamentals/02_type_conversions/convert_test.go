package convert

import "testing"

func TestIntToFloat(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want float64
	}{
		{"Positive", 42, 42.0},
		{"Zero", 0, 0.0},
		{"Negative", -10, -10.0},
		{"Large", 1000000, 1000000.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntToFloat(tt.n)
			if got != tt.want {
				t.Errorf("IntToFloat(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestFloatToInt(t *testing.T) {
	tests := []struct {
		name string
		f    float64
		want int
	}{
		{"Truncate decimals", 3.14, 3},
		{"Round down", 9.99, 9},
		{"Whole number", 42.0, 42},
		{"Negative", -5.7, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FloatToInt(tt.f)
			if got != tt.want {
				t.Errorf("FloatToInt(%v) = %v, want %v", tt.f, got, tt.want)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    int
		wantErr bool
	}{
		{"Valid positive", "123", 123, false},
		{"Valid negative", "-42", -42, false},
		{"Zero", "0", 0, false},
		{"Invalid letters", "abc", 0, true},
		{"Invalid mixed", "12abc", 0, true},
		{"Empty", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInt(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToInt(%q) error = %v, wantErr %v", tt.s, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("StringToInt(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{"Positive", 42, "42"},
		{"Zero", 0, "0"},
		{"Negative", -10, "-10"},
		{"Large", 123456, "123456"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntToString(tt.n)
			if got != tt.want {
				t.Errorf("IntToString(%v) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}
