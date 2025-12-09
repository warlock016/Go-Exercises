package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"1+2", 1, 2, 3},
		{"-2+2", -2, 2, 0},
		{"1-2", 1, -2, -1},
		{"bla", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if tt.want != got {
				t.Errorf("Add(%d, %d) got: %d, want: %d", tt.a, tt.b, got, tt.want)
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
		{"simple subtraction", 1, 1, 0},
		{"2 - 0 subtraction", 2, 0, 2},
		{"0 - 2 subtraction", 0, 2, -2},
		{"negative b", 5, -5, 10},
		{"negative a", -5, 5, -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Subtract(%d, %d) got %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		{"simple division", 3, 1, 3, false},
		{"zero-div error", 2, 0, 0, true},
		{"negative division", 6, -2, -3, false},
		{"valid zero", 0, 5, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Divide(%d, %d) expected error, got nil", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%d, %d) unexpected error: %v", tt.a, tt.b, err)
				}
				if got != tt.want {
					t.Errorf("Divide(%d, %d) got %d, want %d", tt.a, tt.b, got, tt.want)
				}
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
		{"simple multiplication", 1, 2, 2},
		{"negative multiplication", -1, 1, -1},
		{"double negative mult", -1, -2, 2},
		{"zeroA multiplication", 0, 1, 0},
		{"zeroB multiplication", 1, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Multiplty(%d, %d) got %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		name string
		a    int
		want bool
	}{
		{"simple even", 2, true},
		{"simple odd", 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEven(tt.a)
			if tt.want != got {
				t.Errorf("IsEven(%d) unexpected result %v", tt.a, got)
			}
		})
	}
}
