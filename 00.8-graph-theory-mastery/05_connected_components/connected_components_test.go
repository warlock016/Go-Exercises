package connected_components

import "testing"

func TestCountComponents(t *testing.T) {
	tests := []struct {
		name  string
		graph map[int][]int
		want  int
	}{
		{"Three components", map[int][]int{0: {1}, 1: {0}, 2: {3}, 3: {2}, 4: {}}, 3},
		{"One component", map[int][]int{0: {1}, 1: {0, 2}, 2: {1}}, 1},
		{"Empty", map[int][]int{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountComponents(tt.graph); got != tt.want {
				t.Errorf("CountComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsConnected(t *testing.T) {
	connected := map[int][]int{0: {1, 2}, 1: {0, 2}, 2: {0, 1}}
	disconnected := map[int][]int{0: {1}, 1: {0}, 2: {}}

	if !IsConnected(connected) {
		t.Error("Expected connected graph to return true")
	}
	if IsConnected(disconnected) {
		t.Error("Expected disconnected graph to return false")
	}
}
