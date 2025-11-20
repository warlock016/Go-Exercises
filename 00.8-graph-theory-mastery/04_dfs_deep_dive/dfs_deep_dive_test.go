package dfs_deep_dive

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	cyclic := map[int][]int{
		0: {1, 2},
		1: {0, 2},
		2: {0, 1},
	}

	acyclic := map[int][]int{
		0: {1, 2},
		1: {0},
		2: {0},
	}

	tests := []struct {
		name  string
		graph map[int][]int
		want  bool
	}{
		{"Triangle (cyclic)", cyclic, true},
		{"Tree (acyclic)", acyclic, false},
		{"Empty graph", map[int][]int{}, false},
		{"Single vertex", map[int][]int{0: {}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasCycle(tt.graph)
			if got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasCycleDirected(t *testing.T) {
	cyclic := map[int][]int{
		0: {1},
		1: {2},
		2: {0},
	}

	acyclic := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	tests := []struct {
		name  string
		graph map[int][]int
		want  bool
	}{
		{"Directed cycle", cyclic, true},
		{"DAG (acyclic)", acyclic, false},
		{"Empty graph", map[int][]int{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasCycleDirected(tt.graph)
			if got != tt.want {
				t.Errorf("HasCycleDirected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountPathsDFS(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	tests := []struct {
		name  string
		start int
		end   int
		want  int
	}{
		{"Two paths", 0, 3, 2}, // [0,1,3] and [0,2,3]
		{"One path", 1, 3, 1},
		{"No path", 0, 99, 0},
		{"Same vertex", 0, 0, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountPathsDFS(graph, tt.start, tt.end)
			if got != tt.want {
				t.Errorf("CountPathsDFS() = %d, want %d", got, tt.want)
			}
		})
	}
}
