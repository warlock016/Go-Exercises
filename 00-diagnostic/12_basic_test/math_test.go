package math

import (
	"testing"
)

// TODO(human): Implement TestAdd using table-driven testing
// Test cases:
// - Add(2, 3) should return 5
// - Add(0, 0) should return 0
// - Add(-1, 1) should return 0
// - Add(10, -5) should return 5

func TestAdd(t *testing.T) {
	tests := []struct {
		run      string
		valA     int
		valB     int
		expected int
	}{
		{"1", 2, 3, 5},
		{"2", 0, 0, 0},
		{"3", -1, 1, 0},
		{"4", 10, -5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.run, func(t *testing.T) {
			num := Add(tt.valA, tt.valB)

			if num != tt.expected {
				t.Errorf("got %d, want %d", num, tt.expected)
			}
		})
	}
}
