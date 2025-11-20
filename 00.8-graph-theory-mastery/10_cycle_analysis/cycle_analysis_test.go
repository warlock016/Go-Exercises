package cycle_analysis

import "testing"

func TestFindAllCycles(t *testing.T) {
	graph := map[int][]int{0: {1, 2}, 1: {0, 2}, 2: {0, 1}}
	cycles := FindAllCycles(graph)

	if len(cycles) == 0 {
		t.Error("Should find at least one cycle")
	}
}
