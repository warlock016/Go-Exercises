package minimum_spanning_tree

import "testing"

func TestKruskalMST(t *testing.T) {
	edges := [][3]int{
		{0, 1, 1},
		{1, 2, 2},
		{0, 2, 3},
	}

	mst, cost := KruskalMST(edges, 3)
	if cost != 3 || len(mst) != 2 {
		t.Errorf("Expected MST cost 3 with 2 edges, got cost=%d, edges=%d", cost, len(mst))
	}
}
