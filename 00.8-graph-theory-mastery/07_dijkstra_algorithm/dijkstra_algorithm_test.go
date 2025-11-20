package dijkstra_algorithm

import "testing"

func TestDijkstra(t *testing.T) {
	graph := WeightedGraph{
		0: {1: 1, 2: 5},
		1: {3: 2},
		2: {3: 1},
		3: {},
	}

	distances := Dijkstra(graph, 0)

	if distances[0] != 0 || distances[1] != 1 || distances[3] != 3 {
		t.Errorf("Dijkstra distances incorrect: %v", distances)
	}
}
