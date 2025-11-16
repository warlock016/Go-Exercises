package slice_algorithms

import (
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		fn    func(int) int
		want  []int
	}{
		{"double values", []int{1, 2, 3}, func(x int) int { return x * 2 }, []int{2, 4, 6}},
		{"square values", []int{2, 3, 4}, func(x int) int { return x * x }, []int{4, 9, 16}},
		{"add 10", []int{1, 5, 9}, func(x int) int { return x + 10 }, []int{11, 15, 19}},
		{"identity", []int{1, 2, 3}, func(x int) int { return x }, []int{1, 2, 3}},
		{"empty slice", []int{}, func(x int) int { return x * 2 }, []int{}},
		{"negate", []int{1, -2, 3}, func(x int) int { return -x }, []int{-1, 2, -3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.slice, tt.fn)
			if len(got) != len(tt.want) {
				t.Fatalf("Map() length = %d, want %d", len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Map()[%d] = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
	}{
		{"even numbers", []int{1, 2, 3, 4, 5, 6}, func(x int) bool { return x%2 == 0 }, []int{2, 4, 6}},
		{"odd numbers", []int{1, 2, 3, 4, 5}, func(x int) bool { return x%2 != 0 }, []int{1, 3, 5}},
		{"greater than 5", []int{1, 5, 7, 3, 9}, func(x int) bool { return x > 5 }, []int{7, 9}},
		{"all match", []int{2, 4, 6}, func(x int) bool { return x%2 == 0 }, []int{2, 4, 6}},
		{"none match", []int{1, 3, 5}, func(x int) bool { return x%2 == 0 }, []int{}},
		{"empty slice", []int{}, func(x int) bool { return true }, []int{}},
		{"positive only", []int{-2, 0, 1, -3, 5}, func(x int) bool { return x > 0 }, []int{1, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.slice, tt.predicate)
			if len(got) != len(tt.want) {
				t.Fatalf("Filter() length = %d, want %d", len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Filter()[%d] = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		initial int
		fn      func(int, int) int
		want    int
	}{
		{"sum", []int{1, 2, 3, 4}, 0, func(acc, x int) int { return acc + x }, 10},
		{"product", []int{2, 3, 4}, 1, func(acc, x int) int { return acc * x }, 24},
		{"max", []int{3, 7, 2, 9, 1}, 0, func(acc, x int) int {
			if x > acc {
				return x
			}
			return acc
		}, 9},
		{"count", []int{1, 2, 3}, 0, func(acc, x int) int { return acc + 1 }, 3},
		{"empty slice", []int{}, 42, func(acc, x int) int { return acc + x }, 42},
		{"concatenate digits", []int{1, 2, 3}, 0, func(acc, x int) int { return acc*10 + x }, 123},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.slice, tt.initial, tt.fn)
			if got != tt.want {
				t.Errorf("Reduce() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		wantVal   int
		wantFound bool
	}{
		{"find first even", []int{1, 3, 4, 6, 5}, func(x int) bool { return x%2 == 0 }, 4, true},
		{"find first > 5", []int{1, 2, 7, 3, 9}, func(x int) bool { return x > 5 }, 7, true},
		{"not found", []int{1, 3, 5}, func(x int) bool { return x%2 == 0 }, 0, false},
		{"empty slice", []int{}, func(x int) bool { return true }, 0, false},
		{"find first element", []int{1, 2, 3}, func(x int) bool { return x == 1 }, 1, true},
		{"find negative", []int{1, -2, 3}, func(x int) bool { return x < 0 }, -2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, gotFound := Find(tt.slice, tt.predicate)
			if gotVal != tt.wantVal || gotFound != tt.wantFound {
				t.Errorf("Find() = (%d, %v), want (%d, %v)", gotVal, gotFound, tt.wantVal, tt.wantFound)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		value int
		want  bool
	}{
		{"contains first", []int{1, 2, 3}, 1, true},
		{"contains middle", []int{1, 2, 3}, 2, true},
		{"contains last", []int{1, 2, 3}, 3, true},
		{"not contains", []int{1, 2, 3}, 5, false},
		{"empty slice", []int{}, 1, false},
		{"contains zero", []int{0, 1, 2}, 0, true},
		{"contains negative", []int{-1, 0, 1}, -1, true},
		{"single element match", []int{42}, 42, true},
		{"single element no match", []int{42}, 99, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.slice, tt.value)
			if got != tt.want {
				t.Errorf("Contains(%v, %d) = %v, want %v", tt.slice, tt.value, got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"no duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"simple duplicates", []int{1, 2, 2, 3}, []int{1, 2, 3}},
		{"multiple duplicates", []int{1, 2, 2, 3, 1, 4, 3}, []int{1, 2, 3, 4}},
		{"all same", []int{5, 5, 5, 5}, []int{5}},
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"preserves order", []int{3, 1, 2, 1, 3}, []int{3, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Unique(tt.slice)
			if len(got) != len(tt.want) {
				t.Fatalf("Unique(%v) length = %d, want %d", tt.slice, len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Unique(%v)[%d] = %d, want %d", tt.slice, i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		wantTrue  []int
		wantFalse []int
	}{
		{
			"even/odd",
			[]int{1, 2, 3, 4, 5, 6},
			func(x int) bool { return x%2 == 0 },
			[]int{2, 4, 6},
			[]int{1, 3, 5},
		},
		{
			"positive/non-positive",
			[]int{-1, 2, 0, 3, -5},
			func(x int) bool { return x > 0 },
			[]int{2, 3},
			[]int{-1, 0, -5},
		},
		{
			"all match",
			[]int{2, 4, 6},
			func(x int) bool { return x%2 == 0 },
			[]int{2, 4, 6},
			[]int{},
		},
		{
			"none match",
			[]int{1, 3, 5},
			func(x int) bool { return x%2 == 0 },
			[]int{},
			[]int{1, 3, 5},
		},
		{
			"empty slice",
			[]int{},
			func(x int) bool { return true },
			[]int{},
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTrue, gotFalse := Partition(tt.slice, tt.predicate)

			if len(gotTrue) != len(tt.wantTrue) {
				t.Fatalf("Partition() matching length = %d, want %d", len(gotTrue), len(tt.wantTrue))
			}
			for i := range tt.wantTrue {
				if gotTrue[i] != tt.wantTrue[i] {
					t.Errorf("Partition() matching[%d] = %d, want %d", i, gotTrue[i], tt.wantTrue[i])
				}
			}

			if len(gotFalse) != len(tt.wantFalse) {
				t.Fatalf("Partition() non-matching length = %d, want %d", len(gotFalse), len(tt.wantFalse))
			}
			for i := range tt.wantFalse {
				if gotFalse[i] != tt.wantFalse[i] {
					t.Errorf("Partition() non-matching[%d] = %d, want %d", i, gotFalse[i], tt.wantFalse[i])
				}
			}
		})
	}
}
