package advanced_maps

import "strconv"

// TODO(human): Define Student struct with Name and Grade fields

type Student struct {
	Name  string
	Grade string
}

// CreateMapOfSlices creates a map that stores slices of integers
// Initialize with keys "even" and "odd", each with empty slices
func CreateMapOfSlices() map[string][]int {
	// TODO(human): Create and initialize map with "even" and "odd" keys

	return map[string][]int{
		"odd":  {},
		"even": {},
	}
}

// AddToMapSlice appends a value to the slice at the given key
// Creates the slice if the key doesn't exist
func AddToMapSlice(m map[string][]int, key string, value int) {
	// TODO(human): Append value to slice at key
	m[key] = append(m[key], value)

}

// CreateMapOfMaps creates a map of maps (nested maps)
// Initialize with keys "user1" and "user2", each containing empty string->int maps
func CreateMapOfMaps() map[string]map[string]int {
	// TODO(human): Create nested map structure with "user1" and "user2" keys
	return map[string]map[string]int{
		"user1": {},
		"user2": {},
	}
}

// SetNestedValue sets a value in the nested map at [outerKey][innerKey]
// Creates nested map if outerKey doesn't exist
func SetNestedValue(m map[string]map[string]int, outerKey, innerKey string, value int) {
	// TODO(human): Initialize inner map if needed, then set value
	if m[outerKey] == nil {
		m[outerKey] = make(map[string]int)
	}

	m[outerKey][innerKey] = value
}

// CreateMapOfStructs creates a map[int]Student with some initial data
// Add students: {1, "Alice", "A"}, {2, "Bob", "B"}, {3, "Charlie", "A"}
func CreateMapOfStructs() map[int]Student {
	// TODO(human): Create map with student data

	// result := make(map[int]Student)

	// result[1] = Student{Name: "Alice", Grade: "A"}
	// result[2] = Student{Name: "Bob", Grade: "B"}
	// result[3] = Student{Name: "Charlie", Grade: "A"}

	// return result
	return map[int]Student{
		1: {Name: "Alice", Grade: "A"},
		2: {Name: "Bob", Grade: "B"},
		3: {Name: "Charlie", Grade: "A"},
	}
}

// GroupStudentsByGrade takes a slice of students and returns a map
// where keys are grades and values are slices of students with that grade
func GroupStudentsByGrade(students []Student) map[string][]Student {
	// TODO(human): Group students by grade

	result := make(map[string][]Student, 0)

	for _, v := range students {
		result[v.Grade] = append(result[v.Grade], v)
	}

	return result
}

// BuildAdjacencyList converts a slice of edges into a graph adjacency list
// Each edge is [from, to], result maps each node to its neighbors
// Example: [[1,2], [1,3], [2,3]] -> {1:[2,3], 2:[3], 3:[]}
func BuildAdjacencyList(edges [][2]int) map[int][]int {
	// TODO(human): Build graph from edges
	result := make(map[int][]int)

	for _, v := range edges {
		if result[v[0]] == nil {
			result[v[0]] = make([]int, 0)
		}
		if result[v[1]] == nil {
			result[v[1]] = make([]int, 0)
		}
		result[v[0]] = append(result[v[0]], v[1])
	}

	return result
}

// --- NEW EXERCISES: Practice `make` Initialization Patterns ---

// PreallocateMap creates a map with expected capacity and populates it
// Returns map[int]string with keys 1-1000, values "item_1" to "item_1000"
// Use capacity hint for performance
func PreallocateMap(size int) map[int]string {
	// TODO(human): Create map with capacity hint, populate with size entries

	result := make(map[int]string, 1000)

	for i := 1; i <= size; i++ {
		result[i] = "item_" + strconv.Itoa(i)
	}

	return result
}

// SafeNestedIncrement increments the value at m[category][item] by 1
// Creates necessary maps/values if they don't exist (defensive initialization)
// Returns the new value after incrementing
func SafeNestedIncrement(m map[string]map[string]int, category, item string) int {
	// TODO(human): Check if category exists, create if nil, then increment item count

	if m[category] == nil {
		m[category] = make(map[string]int)
	}

	m[category][item]++

	return m[category][item]
}

// BuildInventoryIndex converts flat item data into nested structure
// Input: items with Category and Name fields
// Output: map[category]map[itemName]count
// All counts start at 1 (represents 1 item of each type)
func BuildInventoryIndex(items []struct{ Category, Name string }) map[string]map[string]int {
	// TODO(human): Build two-level map from flat data

	result := make(map[string]map[string]int)

	for _, v := range items {
		if result[v.Category] == nil {
			result[v.Category] = make(map[string]int)
		}
		result[v.Category][v.Name]++
	}
	return result
}
