package floyd_warshall

import "testing"

func TestFloydWarshall(t *testing.T) {
	graph := WeightedGraph{
		0: {1: 1, 2: 5},
		1: {2: 2},
		2: {},
	}

	distances := FloydWarshall(graph, 3)

	if distances[0][2] != 3 {
		t.Errorf("Expected shortest 0→2 = 3, got %d", distances[0][2])
	}
}
