package topological_sort

import "testing"

func TestTopologicalSort(t *testing.T) {
	dag := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	result := TopologicalSort(dag)
	if len(result) != 4 || result[0] != 0 || result[len(result)-1] != 3 {
		t.Errorf("Topological sort incorrect: %v", result)
	}
}

func TestHasValidOrdering(t *testing.T) {
	dag := map[int][]int{0: {1}, 1: {2}, 2: {}}
	cycle := map[int][]int{0: {1}, 1: {2}, 2: {0}}

	if !HasValidOrdering(dag) {
		t.Error("DAG should have valid ordering")
	}
	if HasValidOrdering(cycle) {
		t.Error("Cycle should not have valid ordering")
	}
}
