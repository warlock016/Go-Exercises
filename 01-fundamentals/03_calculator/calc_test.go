package calc

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"Positive numbers", 5, 3, 8},
		{"With zero", 10, 0, 10},
		{"Negative numbers", -5, -3, -8},
		{"Mixed signs", -5, 3, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"Positive result", 10, 4, 6},
		{"Negative result", 3, 10, -7},
		{"With zero", 5, 0, 5},
		{"Negative numbers", -5, -3, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Subtract(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"Positive numbers", 6, 7, 42},
		{"With zero", 5, 0, 0},
		{"With one", 8, 1, 8},
		{"Negative numbers", -3, -4, 12},
		{"Mixed signs", -5, 3, -15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Multiply(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name          string
		a, b          int
		wantQuotient  int
		wantRemainder int
		wantErr       bool
	}{
		{"Simple division", 17, 5, 3, 2, false},
		{"Exact division", 20, 4, 5, 0, false},
		{"Divide by zero", 10, 0, 0, 0, true},
		{"Negative dividend", -17, 5, -3, -2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQ, gotR, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide(%v, %v) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if gotQ != tt.wantQuotient {
					t.Errorf("Divide(%v, %v) quotient = %v, want %v", tt.a, tt.b, gotQ, tt.wantQuotient)
				}
				if gotR != tt.wantRemainder {
					t.Errorf("Divide(%v, %v) remainder = %v, want %v", tt.a, tt.b, gotR, tt.wantRemainder)
				}
			}
		})
	}
}

func TestModulo(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		{"Basic modulo", 17, 5, 2, false},
		{"No remainder", 20, 4, 0, false},
		{"Divide by zero", 10, 0, 0, true},
		{"Negative dividend", -17, 5, -2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Modulo(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Modulo(%v, %v) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Modulo(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
