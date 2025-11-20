package graph_visualization

import (
	"strings"
	"testing"
)

func TestToAdjacencyMatrix(t *testing.T) {
	graph := map[int][]int{0: {1}, 1: {0}}
	result := ToAdjacencyMatrix(graph)

	if !strings.Contains(result, "0") || !strings.Contains(result, "1") {
		t.Error("Matrix should contain vertex labels")
	}
}
