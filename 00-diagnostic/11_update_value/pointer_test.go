package pointer

import "testing"

func TestUpdateValue(t *testing.T) {
	tests := []struct {
		name     string
		initial  int
		newValue int
	}{
		{"Update to positive", 10, 42},
		{"Update to zero", 100, 0},
		{"Update to negative", 5, -10},
		{"Same value", 7, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := tt.initial
			UpdateValue(&num, tt.newValue)
			if num != tt.newValue {
				t.Errorf("After UpdateValue(&num, %v), num = %v, want %v", tt.newValue, num, tt.newValue)
			}
		})
	}
}
