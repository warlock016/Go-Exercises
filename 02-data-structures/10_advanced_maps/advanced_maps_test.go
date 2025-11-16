package advanced_maps

import (
	"testing"
)

func TestCreateMapOfSlices(t *testing.T) {
	m := CreateMapOfSlices()

	if m == nil {
		t.Fatal("CreateMapOfSlices() returned nil")
	}

	if _, exists := m["even"]; !exists {
		t.Error("CreateMapOfSlices() missing 'even' key")
	}

	if _, exists := m["odd"]; !exists {
		t.Error("CreateMapOfSlices() missing 'odd' key")
	}

	if m["even"] == nil {
		t.Error("CreateMapOfSlices() 'even' key has nil slice")
	}

	if m["odd"] == nil {
		t.Error("CreateMapOfSlices() 'odd' key has nil slice")
	}

	if len(m["even"]) != 0 {
		t.Errorf("CreateMapOfSlices() 'even' length = %d, want 0", len(m["even"]))
	}

	if len(m["odd"]) != 0 {
		t.Errorf("CreateMapOfSlices() 'odd' length = %d, want 0", len(m["odd"]))
	}
}

func TestAddToMapSlice(t *testing.T) {
	tests := []struct {
		name      string
		initial   map[string][]int
		key       string
		values    []int
		wantSlice []int
	}{
		{
			"add to existing key",
			map[string][]int{"test": {1, 2}},
			"test",
			[]int{3, 4},
			[]int{1, 2, 3, 4},
		},
		{
			"add to new key",
			make(map[string][]int),
			"new",
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			"add single value",
			map[string][]int{"key": {}},
			"key",
			[]int{42},
			[]int{42},
		},
		{
			"add multiple values",
			map[string][]int{},
			"nums",
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.initial
			for _, val := range tt.values {
				AddToMapSlice(m, tt.key, val)
			}

			got := m[tt.key]
			if len(got) != len(tt.wantSlice) {
				t.Fatalf("AddToMapSlice() length = %d, want %d", len(got), len(tt.wantSlice))
			}

			for i := range tt.wantSlice {
				if got[i] != tt.wantSlice[i] {
					t.Errorf("AddToMapSlice()[%d] = %d, want %d", i, got[i], tt.wantSlice[i])
				}
			}
		})
	}
}

func TestCreateMapOfMaps(t *testing.T) {
	m := CreateMapOfMaps()

	if m == nil {
		t.Fatal("CreateMapOfMaps() returned nil")
	}

	if _, exists := m["user1"]; !exists {
		t.Error("CreateMapOfMaps() missing 'user1' key")
	}

	if _, exists := m["user2"]; !exists {
		t.Error("CreateMapOfMaps() missing 'user2' key")
	}

	if m["user1"] == nil {
		t.Error("CreateMapOfMaps() 'user1' has nil inner map")
	}

	if m["user2"] == nil {
		t.Error("CreateMapOfMaps() 'user2' has nil inner map")
	}

	if len(m["user1"]) != 0 {
		t.Errorf("CreateMapOfMaps() 'user1' length = %d, want 0", len(m["user1"]))
	}

	if len(m["user2"]) != 0 {
		t.Errorf("CreateMapOfMaps() 'user2' length = %d, want 0", len(m["user2"]))
	}
}

func TestSetNestedValue(t *testing.T) {
	tests := []struct {
		name      string
		outerKey  string
		innerKey  string
		value     int
		wantValue int
	}{
		{"set in existing map", "user1", "score", 100, 100},
		{"set in new outer key", "user3", "level", 5, 5},
		{"set multiple in same outer key", "user1", "health", 95, 95},
		{"overwrite existing value", "user1", "score", 200, 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := CreateMapOfMaps()
			SetNestedValue(m, tt.outerKey, tt.innerKey, tt.value)

			if m[tt.outerKey] == nil {
				t.Fatalf("SetNestedValue() did not create inner map for %q", tt.outerKey)
			}

			got, exists := m[tt.outerKey][tt.innerKey]
			if !exists {
				t.Errorf("SetNestedValue() did not set %q[%q]", tt.outerKey, tt.innerKey)
			}

			if got != tt.wantValue {
				t.Errorf("SetNestedValue() %q[%q] = %d, want %d", tt.outerKey, tt.innerKey, got, tt.wantValue)
			}
		})
	}
}

func TestCreateMapOfStructs(t *testing.T) {
	m := CreateMapOfStructs()

	if m == nil {
		t.Fatal("CreateMapOfStructs() returned nil")
	}

	expectedStudents := map[int]Student{
		1: {Name: "Alice", Grade: "A"},
		2: {Name: "Bob", Grade: "B"},
		3: {Name: "Charlie", Grade: "A"},
	}

	if len(m) != len(expectedStudents) {
		t.Fatalf("CreateMapOfStructs() length = %d, want %d", len(m), len(expectedStudents))
	}

	for id, expected := range expectedStudents {
		got, exists := m[id]
		if !exists {
			t.Errorf("CreateMapOfStructs() missing student with ID %d", id)
			continue
		}

		if got.Name != expected.Name {
			t.Errorf("CreateMapOfStructs()[%d].Name = %q, want %q", id, got.Name, expected.Name)
		}

		if got.Grade != expected.Grade {
			t.Errorf("CreateMapOfStructs()[%d].Grade = %q, want %q", id, got.Grade, expected.Grade)
		}
	}
}

func TestGroupStudentsByGrade(t *testing.T) {
	tests := []struct {
		name     string
		students []Student
		want     map[string][]Student
	}{
		{
			"empty list",
			[]Student{},
			map[string][]Student{},
		},
		{
			"single student",
			[]Student{{Name: "Alice", Grade: "A"}},
			map[string][]Student{
				"A": {{Name: "Alice", Grade: "A"}},
			},
		},
		{
			"multiple grades",
			[]Student{
				{Name: "Alice", Grade: "A"},
				{Name: "Bob", Grade: "B"},
				{Name: "Charlie", Grade: "A"},
				{Name: "Diana", Grade: "C"},
			},
			map[string][]Student{
				"A": {{Name: "Alice", Grade: "A"}, {Name: "Charlie", Grade: "A"}},
				"B": {{Name: "Bob", Grade: "B"}},
				"C": {{Name: "Diana", Grade: "C"}},
			},
		},
		{
			"all same grade",
			[]Student{
				{Name: "Alice", Grade: "A"},
				{Name: "Bob", Grade: "A"},
			},
			map[string][]Student{
				"A": {{Name: "Alice", Grade: "A"}, {Name: "Bob", Grade: "A"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupStudentsByGrade(tt.students)

			if len(got) != len(tt.want) {
				t.Fatalf("GroupStudentsByGrade() grades = %d, want %d", len(got), len(tt.want))
			}

			for grade, wantStudents := range tt.want {
				gotStudents, exists := got[grade]
				if !exists {
					t.Errorf("GroupStudentsByGrade() missing grade %q", grade)
					continue
				}

				if len(gotStudents) != len(wantStudents) {
					t.Errorf("GroupStudentsByGrade()[%q] length = %d, want %d", grade, len(gotStudents), len(wantStudents))
					continue
				}

				for i := range wantStudents {
					if gotStudents[i].Name != wantStudents[i].Name {
						t.Errorf("GroupStudentsByGrade()[%q][%d].Name = %q, want %q", grade, i, gotStudents[i].Name, wantStudents[i].Name)
					}
					if gotStudents[i].Grade != wantStudents[i].Grade {
						t.Errorf("GroupStudentsByGrade()[%q][%d].Grade = %q, want %q", grade, i, gotStudents[i].Grade, wantStudents[i].Grade)
					}
				}
			}
		})
	}
}

func TestBuildAdjacencyList(t *testing.T) {
	tests := []struct {
		name  string
		edges [][2]int
		want  map[int][]int
	}{
		{
			"empty graph",
			[][2]int{},
			map[int][]int{},
		},
		{
			"single edge",
			[][2]int{{1, 2}},
			map[int][]int{1: {2}, 2: {}},
		},
		{
			"linear graph",
			[][2]int{{1, 2}, {2, 3}, {3, 4}},
			map[int][]int{1: {2}, 2: {3}, 3: {4}, 4: {}},
		},
		{
			"star graph",
			[][2]int{{1, 2}, {1, 3}, {1, 4}},
			map[int][]int{1: {2, 3, 4}, 2: {}, 3: {}, 4: {}},
		},
		{
			"triangle",
			[][2]int{{1, 2}, {2, 3}, {3, 1}},
			map[int][]int{1: {2}, 2: {3}, 3: {1}},
		},
		{
			"multiple edges from same node",
			[][2]int{{1, 2}, {1, 3}, {2, 3}},
			map[int][]int{1: {2, 3}, 2: {3}, 3: {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildAdjacencyList(tt.edges)

			if len(got) != len(tt.want) {
				t.Fatalf("BuildAdjacencyList() nodes = %d, want %d", len(got), len(tt.want))
			}

			for node, wantNeighbors := range tt.want {
				gotNeighbors, exists := got[node]
				if !exists {
					t.Errorf("BuildAdjacencyList() missing node %d", node)
					continue
				}

				if len(gotNeighbors) != len(wantNeighbors) {
					t.Errorf("BuildAdjacencyList()[%d] neighbors = %d, want %d", node, len(gotNeighbors), len(wantNeighbors))
					continue
				}

				for i := range wantNeighbors {
					if gotNeighbors[i] != wantNeighbors[i] {
						t.Errorf("BuildAdjacencyList()[%d][%d] = %d, want %d", node, i, gotNeighbors[i], wantNeighbors[i])
					}
				}
			}
		})
	}
}
