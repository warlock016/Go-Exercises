package graph_basics

import (
	"slices"
	"testing"
)

// Helper function to check if two slices contain the same elements (order doesn't matter)
func containsSameElements(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := make([]int, len(a))
	bCopy := make([]int, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	slices.Sort(aCopy)
	slices.Sort(bCopy)
	return slices.Equal(aCopy, bCopy)
}

// Adjacency List Tests

func TestNewAdjacencyList(t *testing.T) {
	graph := NewAdjacencyList()
	if graph == nil {
		t.Fatal("NewAdjacencyList() returned nil")
	}
	if len(graph) != 0 {
		t.Errorf("NewAdjacencyList() length = %d, want 0", len(graph))
	}
}

func TestAddEdge(t *testing.T) {
	graph := NewAdjacencyList()
	AddEdge(graph, 0, 1)
	AddEdge(graph, 0, 2)
	AddEdge(graph, 1, 2)

	if !containsSameElements(graph[0], []int{1, 2}) {
		t.Errorf("graph[0] = %v, want [1, 2]", graph[0])
	}
	if !containsSameElements(graph[1], []int{2}) {
		t.Errorf("graph[1] = %v, want [2]", graph[1])
	}
}

func TestAddUndirectedEdge(t *testing.T) {
	graph := NewAdjacencyList()
	AddUndirectedEdge(graph, 0, 1)

	if !HasEdge(graph, 0, 1) {
		t.Error("Should have edge 0 -> 1")
	}
	if !HasEdge(graph, 1, 0) {
		t.Error("Should have edge 1 -> 0 (undirected)")
	}
}

func TestHasEdge(t *testing.T) {
	graph := NewAdjacencyList()
	AddEdge(graph, 0, 1)
	AddEdge(graph, 1, 2)

	tests := []struct {
		name   string
		source int
		dest   int
		want   bool
	}{
		{"existing edge 0->1", 0, 1, true},
		{"existing edge 1->2", 1, 2, true},
		{"non-existing edge 0->2", 0, 2, false},
		{"reverse edge 1->0", 1, 0, false},
		{"non-existing vertex", 5, 6, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasEdge(graph, tt.source, tt.dest)
			if got != tt.want {
				t.Errorf("HasEdge(%d, %d) = %v, want %v", tt.source, tt.dest, got, tt.want)
			}
		})
	}
}

func TestGetNeighbors(t *testing.T) {
	graph := NewAdjacencyList()
	AddEdge(graph, 0, 1)
	AddEdge(graph, 0, 2)
	AddEdge(graph, 0, 3)
	AddEdge(graph, 1, 2)

	tests := []struct {
		name   string
		vertex int
		want   []int
	}{
		{"vertex 0", 0, []int{1, 2, 3}},
		{"vertex 1", 1, []int{2}},
		{"vertex 2 (no neighbors)", 2, []int{}},
		{"non-existing vertex", 5, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetNeighbors(graph, tt.vertex)
			if !containsSameElements(got, tt.want) {
				t.Errorf("GetNeighbors(%d) = %v, want %v", tt.vertex, got, tt.want)
			}
		})
	}
}

func TestRemoveEdge(t *testing.T) {
	graph := NewAdjacencyList()
	AddEdge(graph, 0, 1)
	AddEdge(graph, 0, 2)
	AddEdge(graph, 1, 2)

	RemoveEdge(graph, 0, 1)

	if HasEdge(graph, 0, 1) {
		t.Error("Edge 0->1 should be removed")
	}
	if !HasEdge(graph, 0, 2) {
		t.Error("Edge 0->2 should still exist")
	}

	// Remove non-existing edge should not error
	RemoveEdge(graph, 5, 6)
}

// Adjacency Matrix Tests

func TestNewAdjacencyMatrix(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"3x3 matrix", 3},
		{"5x5 matrix", 5},
		{"1x1 matrix", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := NewAdjacencyMatrix(tt.n)
			if len(matrix) != tt.n {
				t.Fatalf("Matrix rows = %d, want %d", len(matrix), tt.n)
			}
			for i := range matrix {
				if len(matrix[i]) != tt.n {
					t.Errorf("Matrix[%d] cols = %d, want %d", i, len(matrix[i]), tt.n)
				}
				// Check all values are false
				for j := range matrix[i] {
					if matrix[i][j] {
						t.Errorf("Matrix[%d][%d] = true, want false (initial state)", i, j)
					}
				}
			}
		})
	}
}

func TestAddEdgeMatrix(t *testing.T) {
	matrix := NewAdjacencyMatrix(4)
	AddEdgeMatrix(matrix, 0, 1)
	AddEdgeMatrix(matrix, 0, 2)
	AddEdgeMatrix(matrix, 1, 3)

	if !matrix[0][1] {
		t.Error("matrix[0][1] should be true")
	}
	if !matrix[0][2] {
		t.Error("matrix[0][2] should be true")
	}
	if !matrix[1][3] {
		t.Error("matrix[1][3] should be true")
	}
	if matrix[1][0] {
		t.Error("matrix[1][0] should be false (directed graph)")
	}
}

func TestAddUndirectedEdgeMatrix(t *testing.T) {
	matrix := NewAdjacencyMatrix(3)
	AddUndirectedEdgeMatrix(matrix, 0, 1)

	if !matrix[0][1] {
		t.Error("matrix[0][1] should be true")
	}
	if !matrix[1][0] {
		t.Error("matrix[1][0] should be true (undirected)")
	}
}

func TestHasEdgeMatrix(t *testing.T) {
	matrix := NewAdjacencyMatrix(4)
	AddEdgeMatrix(matrix, 0, 1)
	AddEdgeMatrix(matrix, 1, 2)

	tests := []struct {
		name   string
		source int
		dest   int
		want   bool
	}{
		{"existing edge 0->1", 0, 1, true},
		{"existing edge 1->2", 1, 2, true},
		{"non-existing edge 0->2", 0, 2, false},
		{"reverse edge 1->0", 1, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasEdgeMatrix(matrix, tt.source, tt.dest)
			if got != tt.want {
				t.Errorf("HasEdgeMatrix(%d, %d) = %v, want %v", tt.source, tt.dest, got, tt.want)
			}
		})
	}
}

func TestGetNeighborsMatrix(t *testing.T) {
	matrix := NewAdjacencyMatrix(4)
	AddEdgeMatrix(matrix, 0, 1)
	AddEdgeMatrix(matrix, 0, 2)
	AddEdgeMatrix(matrix, 0, 3)
	AddEdgeMatrix(matrix, 1, 2)

	tests := []struct {
		name   string
		vertex int
		want   []int
	}{
		{"vertex 0", 0, []int{1, 2, 3}},
		{"vertex 1", 1, []int{2}},
		{"vertex 2 (no neighbors)", 2, []int{}},
		{"vertex 3 (no neighbors)", 3, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetNeighborsMatrix(matrix, tt.vertex)
			if !containsSameElements(got, tt.want) {
				t.Errorf("GetNeighborsMatrix(%d) = %v, want %v", tt.vertex, got, tt.want)
			}
		})
	}
}

// Conversion Tests

func TestListToMatrix(t *testing.T) {
	graph := NewAdjacencyList()
	AddEdge(graph, 0, 1)
	AddEdge(graph, 0, 2)
	AddEdge(graph, 1, 2)
	AddEdge(graph, 2, 3)

	matrix := ListToMatrix(graph, 4)

	if !matrix[0][1] || !matrix[0][2] {
		t.Error("Edges from vertex 0 not converted correctly")
	}
	if !matrix[1][2] {
		t.Error("Edge 1->2 not converted correctly")
	}
	if !matrix[2][3] {
		t.Error("Edge 2->3 not converted correctly")
	}
	if matrix[1][0] {
		t.Error("Non-existing edge should be false")
	}
}

func TestMatrixToList(t *testing.T) {
	matrix := NewAdjacencyMatrix(4)
	AddEdgeMatrix(matrix, 0, 1)
	AddEdgeMatrix(matrix, 0, 2)
	AddEdgeMatrix(matrix, 1, 2)
	AddEdgeMatrix(matrix, 2, 3)

	graph := MatrixToList(matrix)

	if !containsSameElements(graph[0], []int{1, 2}) {
		t.Errorf("graph[0] = %v, want [1, 2]", graph[0])
	}
	if !containsSameElements(graph[1], []int{2}) {
		t.Errorf("graph[1] = %v, want [2]", graph[1])
	}
	if !containsSameElements(graph[2], []int{3}) {
		t.Errorf("graph[2] = %v, want [3]", graph[2])
	}
}

// BFS Tests

func TestBFS(t *testing.T) {
	// Create graph:
	//     0
	//    / \
	//   1   2
	//   |   |
	//   3   4
	graph := NewAdjacencyList()
	AddEdge(graph, 0, 1)
	AddEdge(graph, 0, 2)
	AddEdge(graph, 1, 3)
	AddEdge(graph, 2, 4)

	result := BFS(graph, 0)

	// BFS should visit level by level: 0, then 1 and 2, then 3 and 4
	if len(result) != 5 {
		t.Fatalf("BFS result length = %d, want 5", len(result))
	}
	if result[0] != 0 {
		t.Errorf("BFS[0] = %d, want 0 (start)", result[0])
	}

	// Level 1: 1 and 2 (order may vary)
	level1 := []int{result[1], result[2]}
	if !containsSameElements(level1, []int{1, 2}) {
		t.Errorf("BFS level 1 = %v, want [1, 2]", level1)
	}

	// Level 2: 3 and 4 (order may vary)
	level2 := []int{result[3], result[4]}
	if !containsSameElements(level2, []int{3, 4}) {
		t.Errorf("BFS level 2 = %v, want [3, 4]", level2)
	}
}

func TestBFSLinear(t *testing.T) {
	// Linear graph: 0 -> 1 -> 2 -> 3
	graph := NewAdjacencyList()
	AddEdge(graph, 0, 1)
	AddEdge(graph, 1, 2)
	AddEdge(graph, 2, 3)

	result := BFS(graph, 0)

	want := []int{0, 1, 2, 3}
	if !slices.Equal(result, want) {
		t.Errorf("BFS(linear graph) = %v, want %v", result, want)
	}
}

func TestBFSSingleNode(t *testing.T) {
	graph := NewAdjacencyList()
	AddEdge(graph, 0, 0) // Self-loop should be handled

	result := BFS(graph, 0)

	if len(result) != 1 || result[0] != 0 {
		t.Errorf("BFS(single node) = %v, want [0]", result)
	}
}

// DFS Tests

func TestDFS(t *testing.T) {
	// Create graph:
	//     0
	//    / \
	//   1   2
	//   |   |
	//   3   4
	graph := NewAdjacencyList()
	AddEdge(graph, 0, 1)
	AddEdge(graph, 0, 2)
	AddEdge(graph, 1, 3)
	AddEdge(graph, 2, 4)

	result := DFS(graph, 0)

	// DFS should visit depth-first
	if len(result) != 5 {
		t.Fatalf("DFS result length = %d, want 5", len(result))
	}
	if result[0] != 0 {
		t.Errorf("DFS[0] = %d, want 0 (start)", result[0])
	}

	// Should contain all vertices
	if !containsSameElements(result, []int{0, 1, 2, 3, 4}) {
		t.Errorf("DFS should visit all vertices, got %v", result)
	}
}

func TestDFSLinear(t *testing.T) {
	// Linear graph: 0 -> 1 -> 2 -> 3
	graph := NewAdjacencyList()
	AddEdge(graph, 0, 1)
	AddEdge(graph, 1, 2)
	AddEdge(graph, 2, 3)

	result := DFS(graph, 0)

	want := []int{0, 1, 2, 3}
	if !slices.Equal(result, want) {
		t.Errorf("DFS(linear graph) = %v, want %v", result, want)
	}
}

func TestDFSSingleNode(t *testing.T) {
	graph := NewAdjacencyList()
	result := DFS(graph, 0)

	if len(result) != 1 || result[0] != 0 {
		t.Errorf("DFS(single node) = %v, want [0]", result)
	}
}

func TestDFSCycle(t *testing.T) {
	// Graph with cycle: 0 -> 1 -> 2 -> 0
	graph := NewAdjacencyList()
	AddEdge(graph, 0, 1)
	AddEdge(graph, 1, 2)
	AddEdge(graph, 2, 0) // Cycle back

	result := DFS(graph, 0)

	// Should visit each vertex exactly once despite cycle
	if len(result) != 3 {
		t.Errorf("DFS(cycle) length = %d, want 3 (each vertex once)", len(result))
	}
	if !containsSameElements(result, []int{0, 1, 2}) {
		t.Errorf("DFS(cycle) = %v, want [0, 1, 2]", result)
	}
}
