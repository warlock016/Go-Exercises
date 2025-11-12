package findmax

import "testing"

func TestFindMax(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
		wantErr bool
	}{
		{"Basic max", []int{1, 5, 3, 9, 2}, 9, false},
		{"All negative", []int{-10, -5, -20}, -5, false},
		{"Single element", []int{42}, 42, false},
		{"Max at start", []int{100, 1, 2, 3}, 100, false},
		{"Max at end", []int{1, 2, 3, 100}, 100, false},
		{"Empty slice", []int{}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindMax(tt.numbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMax(%v) error = %v, wantErr %v", tt.numbers, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("FindMax(%v) = %v, want %v", tt.numbers, got, tt.want)
			}
		})
	}
}
