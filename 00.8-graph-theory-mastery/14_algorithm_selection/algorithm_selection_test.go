package algorithm_selection

import "testing"

func TestChooseShortestPathAlgorithm(t *testing.T) {
	tests := []struct {
		name      string
		graphType string
		weighted  bool
		negative  bool
		allPairs  bool
		want      string
	}{
		{"Unweighted", "undirected", false, false, false, "BFS"},
		{"Weighted non-negative single", "directed", true, false, false, "Dijkstra"},
		{"All pairs", "directed", true, false, true, "Floyd-Warshall"},
		{"Negative weights", "directed", true, true, false, "Bellman-Ford"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChooseShortestPathAlgorithm(tt.graphType, tt.weighted, tt.negative, tt.allPairs)
			if got != tt.want {
				t.Errorf("ChooseShortestPathAlgorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}
