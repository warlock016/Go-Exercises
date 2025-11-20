package bfs_deep_dive

import (
	"reflect"
	"sort"
	"testing"
)

func TestShortestDistance(t *testing.T) {
	graph := map[int][]int{
		0: {1, 3},
		1: {0, 2},
		2: {1, 5},
		3: {0, 4},
		4: {3, 5},
		5: {2, 4},
	}

	tests := []struct {
		name  string
		start int
		end   int
		want  int
	}{
		{"Same vertex", 0, 0, 0},
		{"Distance 1", 0, 1, 1},
		{"Distance 2", 0, 2, 2},
		{"Distance 3", 0, 5, 3},
		{"No path", 0, 99, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShortestDistance(graph, tt.start, tt.end)
			if got != tt.want {
				t.Errorf("ShortestDistance() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestLevelOrder(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2, 3},
		1: {0, 4},
		2: {0},
		3: {0, 5},
		4: {1},
		5: {3},
	}

	result := LevelOrder(graph, 0)

	// Verify structure
	if len(result) == 0 {
		t.Fatal("Expected non-empty level order")
	}

	if !reflect.DeepEqual(result[0], []int{0}) {
		t.Errorf("Level 0 should be [0], got %v", result[0])
	}

	// Level 1 should contain {1, 2, 3}
	if len(result) > 1 {
		level1 := result[1]
		sort.Ints(level1)
		expected := []int{1, 2, 3}
		if !reflect.DeepEqual(level1, expected) {
			t.Errorf("Level 1 should contain %v, got %v", expected, level1)
		}
	}
}

func TestFindClosestVertex(t *testing.T) {
	graph := map[int][]int{
		0: {1},
		1: {0, 2},
		2: {1, 3},
		3: {2, 4},
		4: {3},
	}

	isEven := func(v int) bool { return v%2 == 0 }
	isGreaterThan3 := func(v int) bool { return v > 3 }

	tests := []struct {
		name          string
		start         int
		condition     func(int) bool
		wantVertex    int
		wantDistance  int
	}{
		{"Find even (self)", 0, isEven, 0, 0},
		{"Find even (distance 1)", 1, isEven, 0, 1},
		{"Find > 3 from 0", 0, isGreaterThan3, 4, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, gotD := FindClosestVertex(graph, tt.start, tt.condition)
			if gotV != tt.wantVertex || gotD != tt.wantDistance {
				t.Errorf("FindClosestVertex() = (%d, %d), want (%d, %d)",
					gotV, gotD, tt.wantVertex, tt.wantDistance)
			}
		})
	}
}

func TestCountComponentsBFS(t *testing.T) {
	threeComponents := map[int][]int{
		0: {1},
		1: {0},
		2: {3},
		3: {2},
		4: {},
	}

	tests := []struct {
		name  string
		graph map[int][]int
		want  int
	}{
		{"Three components", threeComponents, 3},
		{"One component", map[int][]int{0: {1}, 1: {0, 2}, 2: {1}}, 1},
		{"Empty graph", map[int][]int{}, 0},
		{"Isolated vertices", map[int][]int{0: {}, 1: {}, 2: {}}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountComponentsBFS(tt.graph)
			if got != tt.want {
				t.Errorf("CountComponentsBFS() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestGetComponentBFS(t *testing.T) {
	graph := map[int][]int{
		0: {1},
		1: {0, 2},
		2: {1},
		3: {4},
		4: {3},
		5: {},
	}

	tests := []struct {
		name  string
		start int
		want  []int
	}{
		{"Component with 0", 0, []int{0, 1, 2}},
		{"Component with 3", 3, []int{3, 4}},
		{"Isolated vertex", 5, []int{5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetComponentBFS(graph, tt.start)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetComponentBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
