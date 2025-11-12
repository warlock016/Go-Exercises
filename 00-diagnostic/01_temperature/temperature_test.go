package temperature

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name    string
		celsius float64
		want    float64
	}{
		{"Freezing point", 0, 32},
		{"Boiling point", 100, 212},
		{"Body temperature", 37, 98.6},
		{"Negative temp", -40, -40},
		{"Room temperature", 20, 68},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CelsiusToFahrenheit(tt.celsius)
			if got != tt.want {
				t.Errorf("CelsiusToFahrenheit(%v) = %v, want %v", tt.celsius, got, tt.want)
			}
		})
	}
}
