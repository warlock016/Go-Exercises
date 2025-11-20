package path_existence

import (
	"reflect"
	"testing"
)

func TestHasPath(t *testing.T) {
	connectedGraph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	disconnectedGraph := map[int][]int{
		0: {1},
		1: {},
		2: {3},
		3: {},
	}

	tests := []struct {
		name  string
		graph map[int][]int
		start int
		end   int
		want  bool
	}{
		{"Path exists (direct)", connectedGraph, 0, 1, true},
		{"Path exists (indirect)", connectedGraph, 0, 3, true},
		{"Start equals end", connectedGraph, 0, 0, true},
		{"No path (disconnected)", disconnectedGraph, 0, 3, false},
		{"Path to self (isolated)", map[int][]int{0: {}}, 0, 0, true},
		{"Empty graph", map[int][]int{}, 0, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasPath(tt.graph, tt.start, tt.end)
			if got != tt.want {
				t.Errorf("HasPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindPathBFS(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {4},
		3: {5},
		4: {5},
		5: {},
	}

	tests := []struct {
		name       string
		graph      map[int][]int
		start      int
		end        int
		wantLength int
		wantNil    bool
	}{
		{"Shortest path 0->5", graph, 0, 5, 3, false}, // Could be [0,1,3,5] or [0,2,4,5]
		{"Direct path", graph, 0, 1, 1, false},
		{"Start equals end", graph, 0, 0, 0, false},
		{"No path", map[int][]int{0: {}, 1: {}}, 0, 1, 0, true},
		{"Path length 2", graph, 1, 5, 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindPathBFS(tt.graph, tt.start, tt.end)

			if tt.wantNil {
				if got != nil {
					t.Errorf("FindPathBFS() = %v, want nil", got)
				}
				return
			}

			if got == nil {
				t.Errorf("FindPathBFS() = nil, want path")
				return
			}

			// Verify path length
			if GetPathLength(got) != tt.wantLength {
				t.Errorf("Path length = %d, want %d. Path: %v",
					GetPathLength(got), tt.wantLength, got)
			}

			// Verify path starts and ends correctly
			if got[0] != tt.start || got[len(got)-1] != tt.end {
				t.Errorf("Path doesn't connect start to end. Got: %v", got)
			}

			// Verify path is valid (each step is connected)
			for i := 0; i < len(got)-1; i++ {
				current := got[i]
				next := got[i+1]
				neighbors := tt.graph[current]

				found := false
				for _, neighbor := range neighbors {
					if neighbor == next {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("Invalid path: no edge from %d to %d", current, next)
				}
			}
		})
	}
}

func TestFindPathDFS(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {4},
		3: {5},
		4: {5},
		5: {},
	}

	tests := []struct {
		name    string
		graph   map[int][]int
		start   int
		end     int
		wantNil bool
	}{
		{"Path exists 0->5", graph, 0, 5, false},
		{"Direct path", graph, 0, 1, false},
		{"Start equals end", graph, 0, 0, false},
		{"No path", map[int][]int{0: {}, 1: {}}, 0, 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindPathDFS(tt.graph, tt.start, tt.end)

			if tt.wantNil {
				if got != nil {
					t.Errorf("FindPathDFS() = %v, want nil", got)
				}
				return
			}

			if got == nil {
				t.Errorf("FindPathDFS() = nil, want path")
				return
			}

			// Verify path starts and ends correctly
			if got[0] != tt.start || got[len(got)-1] != tt.end {
				t.Errorf("Path doesn't connect start to end. Got: %v", got)
			}
		})
	}
}

func TestFindAllPathsDFS(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	tests := []struct {
		name      string
		graph     map[int][]int
		start     int
		end       int
		wantPaths int
	}{
		{"Two paths from 0->3", graph, 0, 3, 2}, // [0,1,3] and [0,2,3]
		{"One path (linear)", map[int][]int{0: {1}, 1: {2}, 2: {}}, 0, 2, 1},
		{"No path", map[int][]int{0: {}, 1: {}}, 0, 1, 0},
		{"Start equals end", graph, 0, 0, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindAllPathsDFS(tt.graph, tt.start, tt.end)

			if len(got) != tt.wantPaths {
				t.Errorf("FindAllPathsDFS() found %d paths, want %d. Paths: %v",
					len(got), tt.wantPaths, got)
			}

			// Verify each path is valid
			for i, path := range got {
				if len(path) == 0 {
					t.Errorf("Path %d is empty", i)
					continue
				}

				if path[0] != tt.start || path[len(path)-1] != tt.end {
					t.Errorf("Path %d doesn't connect start to end: %v", i, path)
				}
			}
		})
	}
}

func TestGetPathLength(t *testing.T) {
	tests := []struct {
		name string
		path []int
		want int
	}{
		{"Empty path", []int{}, 0},
		{"Single vertex", []int{0}, 0},
		{"Two vertices (1 edge)", []int{0, 1}, 1},
		{"Four vertices (3 edges)", []int{0, 1, 2, 3}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetPathLength(tt.path)
			if got != tt.want {
				t.Errorf("GetPathLength(%v) = %d, want %d", tt.path, got, tt.want)
			}
		})
	}
}

// Test that BFS finds shortest path
func TestBFSFindsShortestPath(t *testing.T) {
	// Graph with multiple paths of different lengths
	graph := map[int][]int{
		0: {1, 2},
		1: {2, 3},
		2: {3},
		3: {},
	}

	path := FindPathBFS(graph, 0, 3)
	if path == nil {
		t.Fatal("Expected path, got nil")
	}

	// BFS should find path of length 2 (0->1->3 or 0->2->3)
	// NOT length 3 (0->1->2->3)
	if GetPathLength(path) != 2 {
		t.Errorf("BFS should find shortest path of length 2, got length %d: %v",
			GetPathLength(path), path)
	}
}

// Test path uniqueness in FindAllPaths
func TestFindAllPathsUnique(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	paths := FindAllPathsDFS(graph, 0, 3)

	// Check for duplicates
	seen := make(map[string]bool)
	for _, path := range paths {
		key := pathToString(path)
		if seen[key] {
			t.Errorf("Duplicate path found: %v", path)
		}
		seen[key] = true
	}
}

func pathToString(path []int) string {
	s := ""
	for _, v := range path {
		s += string(rune(v + '0'))
	}
	return s
}

// Benchmark
func BenchmarkFindPathBFS(b *testing.B) {
	graph := make(map[int][]int)
	for i := 0; i < 100; i++ {
		graph[i] = []int{i + 1}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindPathBFS(graph, 0, 50)
	}
}

func BenchmarkFindAllPathsDFS(b *testing.B) {
	graph := map[int][]int{
		0: {1, 2, 3},
		1: {4},
		2: {4},
		3: {4},
		4: {},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindAllPathsDFS(graph, 0, 4)
	}
}
