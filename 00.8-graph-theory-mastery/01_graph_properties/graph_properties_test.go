package graph_properties

import (
	"math"
	"testing"
)

func TestGetDegree(t *testing.T) {
	undirectedGraph := map[int][]int{
		0: {1, 2},
		1: {0, 3},
		2: {0, 3},
		3: {1, 2},
	}

	tests := []struct {
		name   string
		graph  map[int][]int
		vertex int
		want   int
	}{
		{"Vertex with 2 neighbors", undirectedGraph, 0, 2},
		{"Vertex with 2 neighbors", undirectedGraph, 1, 2},
		{"Non-existent vertex", undirectedGraph, 999, 0},
		{"Isolated vertex", map[int][]int{0: {}}, 0, 0},
		{"Empty graph", map[int][]int{}, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetDegree(tt.graph, tt.vertex)
			if got != tt.want {
				t.Errorf("GetDegree(%v, %d) = %d, want %d", tt.graph, tt.vertex, got, tt.want)
			}
		})
	}
}

func TestGetInDegree(t *testing.T) {
	directedGraph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	tests := []struct {
		name   string
		graph  map[int][]int
		vertex int
		want   int
	}{
		{"Root vertex (no incoming)", directedGraph, 0, 0},
		{"Vertex with one incoming", directedGraph, 1, 1},
		{"Sink vertex (2 incoming)", directedGraph, 3, 2},  // Fixed: vertex 3 has edges from 1 and 2
		{"Non-existent vertex", directedGraph, 999, 0},
		{"Empty graph", map[int][]int{}, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetInDegree(tt.graph, tt.vertex)
			if got != tt.want {
				t.Errorf("GetInDegree(%v, %d) = %d, want %d", tt.graph, tt.vertex, got, tt.want)
			}
		})
	}
}

func TestGetOutDegree(t *testing.T) {
	directedGraph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	tests := []struct {
		name   string
		graph  map[int][]int
		vertex int
		want   int
	}{
		{"Root vertex (2 outgoing)", directedGraph, 0, 2},
		{"Vertex with one outgoing", directedGraph, 1, 1},
		{"Sink vertex (no outgoing)", directedGraph, 3, 0},
		{"Non-existent vertex", directedGraph, 999, 0},
		{"Empty graph", map[int][]int{}, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetOutDegree(tt.graph, tt.vertex)
			if got != tt.want {
				t.Errorf("GetOutDegree(%v, %d) = %d, want %d", tt.graph, tt.vertex, got, tt.want)
			}
		})
	}
}

func TestGetGraphDensity(t *testing.T) {
	sparseUndirected := map[int][]int{
		0: {1},
		1: {0},
		2: {3},
		3: {2},
	}

	denseUndirected := map[int][]int{
		0: {1, 2, 3},
		1: {0, 2, 3},
		2: {0, 1, 3},
		3: {0, 1, 2},
	}

	directedGraph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	tests := []struct {
		name        string
		graph       map[int][]int
		numVertices int
		directed    bool
		want        float64
		tolerance   float64
	}{
		{"Sparse undirected (2 edges, 4 vertices)", sparseUndirected, 4, false, 0.333, 0.01},
		{"Dense undirected (6 edges, 4 vertices)", denseUndirected, 4, false, 1.0, 0.01},
		{"Directed graph (4 edges, 4 vertices)", directedGraph, 4, true, 0.333, 0.01},
		{"Empty graph", map[int][]int{}, 0, false, 0.0, 0.01},
		{"Single vertex", map[int][]int{0: {}}, 1, false, 0.0, 0.01},
		{"Two vertices, one edge undirected", map[int][]int{0: {1}, 1: {0}}, 2, false, 1.0, 0.01},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetGraphDensity(tt.graph, tt.numVertices, tt.directed)
			if math.Abs(got-tt.want) > tt.tolerance {
				t.Errorf("GetGraphDensity(%v, %d, %v) = %.3f, want %.3f",
					tt.graph, tt.numVertices, tt.directed, got, tt.want)
			}
		})
	}
}

func TestIsDirected(t *testing.T) {
	undirectedGraph := map[int][]int{
		0: {1, 2},
		1: {0, 3},
		2: {0, 3},
		3: {1, 2},
	}

	directedGraph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	mixedGraph := map[int][]int{
		0: {1},
		1: {0, 2},
		2: {}, // 1→2 exists but 2→1 doesn't
	}

	tests := []struct {
		name  string
		graph map[int][]int
		want  bool
	}{
		{"Undirected graph (all edges bidirectional)", undirectedGraph, false},
		{"Directed graph (not all edges bidirectional)", directedGraph, true},
		{"Mixed graph (some edges unidirectional)", mixedGraph, true},
		{"Empty graph", map[int][]int{}, false},
		{"Single vertex", map[int][]int{0: {}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsDirected(tt.graph)
			if got != tt.want {
				t.Errorf("IsDirected(%v) = %v, want %v", tt.graph, got, tt.want)
			}
		})
	}
}

func TestCountVertices(t *testing.T) {
	tests := []struct {
		name  string
		graph map[int][]int
		want  int
	}{
		{"Four vertices", map[int][]int{0: {1}, 1: {0}, 2: {3}, 3: {2}}, 4},
		{"Empty graph", map[int][]int{}, 0},
		{"Single vertex", map[int][]int{0: {}}, 1},
		{"Ten vertices", map[int][]int{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountVertices(tt.graph)
			if got != tt.want {
				t.Errorf("CountVertices(%v) = %d, want %d", tt.graph, got, tt.want)
			}
		})
	}
}

func TestCountEdges(t *testing.T) {
	undirectedGraph := map[int][]int{
		0: {1, 2},
		1: {0, 3},
		2: {0, 3},
		3: {1, 2},
	}

	directedGraph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	tests := []struct {
		name     string
		graph    map[int][]int
		directed bool
		want     int
	}{
		{"Undirected graph (4 edges)", undirectedGraph, false, 4},
		{"Directed graph (4 edges)", directedGraph, true, 4},
		{"Empty graph", map[int][]int{}, false, 0},
		{"Single edge undirected", map[int][]int{0: {1}, 1: {0}}, false, 1},
		{"Single edge directed", map[int][]int{0: {1}, 1: {}}, true, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountEdges(tt.graph, tt.directed)
			if got != tt.want {
				t.Errorf("CountEdges(%v, %v) = %d, want %d", tt.graph, tt.directed, got, tt.want)
			}
		})
	}
}

// Benchmark tests
func BenchmarkGetInDegree(b *testing.B) {
	graph := map[int][]int{
		0: {1, 2, 3, 4},
		1: {2, 3, 4, 5},
		2: {3, 4, 5, 6},
		3: {4, 5, 6, 7},
		4: {5, 6, 7, 8},
		5: {6, 7, 8, 9},
		6: {7, 8, 9},
		7: {8, 9},
		8: {9},
		9: {},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetInDegree(graph, 5)
	}
}

func BenchmarkIsDirected(b *testing.B) {
	graph := map[int][]int{
		0: {1, 2, 3},
		1: {0, 2, 3},
		2: {0, 1, 3},
		3: {0, 1, 2},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsDirected(graph)
	}
}
