package bipartite_detection

import "testing"

func TestIsBipartite(t *testing.T) {
	bipartite := map[int][]int{0: {1, 3}, 1: {0, 2}, 2: {1, 3}, 3: {0, 2}}
	triangle := map[int][]int{0: {1, 2}, 1: {0, 2}, 2: {0, 1}}

	if !IsBipartite(bipartite) {
		t.Error("Expected bipartite graph to return true")
	}
	if IsBipartite(triangle) {
		t.Error("Triangle (odd cycle) should not be bipartite")
	}
}
