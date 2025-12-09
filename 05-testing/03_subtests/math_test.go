package math

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		Name  string
		input int
		want  int
	}{
		{"negative", -1, 1},
		{"positive", 1, 1},
		{"neg zero", -0, 0},
		{"zero", 0, 0},
		{"large number", 10e6, 10e6},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := Abs(tt.input)
			if got != tt.want {
				t.Errorf("Abs(%d) got %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		Name string
		a, b int
		want int
	}{
		{"negative", -1, 1, 1},
		{"positive", 2, 1, 2},
		{"neg zero", -0, -1, 0},
		{"zero", 0, 0, 0},
		{"obvious", 0, 10, 10},
		{"negative vs positive", -5, 5, 5},
		{"positive vs negative", 4, -4, 4},
		{"double negative", -5, -2, -2},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := Max(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Max(%d, %d) got %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		min, max int
		want     int
	}{
		{"clamp min", -10, 0, 10, 0},
		{"clamp max", 50, 0, 20, 20},
		{"neg clamp min", -5, -20, -5, -5},
		{"neg clamp max", -5, -20, -10, -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Clamp(tt.value, tt.min, tt.max)
			if got != tt.want {
				t.Errorf("Clamp(%d [%d:%d]) got %d, want %d", tt.value, tt.min, tt.max, got, tt.want)
			}
		})
	}
}
