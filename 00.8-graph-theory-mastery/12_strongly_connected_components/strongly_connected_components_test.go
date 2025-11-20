package strongly_connected_components

import "testing"

func TestKosarajuSCC(t *testing.T) {
	graph := map[int][]int{
		0: {1},
		1: {2},
		2: {0},
		3: {},
	}

	sccs := KosarajuSCC(graph)
	if len(sccs) != 2 {
		t.Errorf("Expected 2 SCCs, got %d", len(sccs))
	}
}
