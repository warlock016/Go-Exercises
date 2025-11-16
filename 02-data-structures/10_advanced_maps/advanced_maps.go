package advanced_maps

// TODO(human): Define Student struct with Name and Grade fields

// CreateMapOfSlices creates a map that stores slices of integers
// Initialize with keys "even" and "odd", each with empty slices
func CreateMapOfSlices() map[string][]int {
	// TODO(human): Create and initialize map with "even" and "odd" keys
	return nil
}

// AddToMapSlice appends a value to the slice at the given key
// Creates the slice if the key doesn't exist
func AddToMapSlice(m map[string][]int, key string, value int) {
	// TODO(human): Append value to slice at key
}

// CreateMapOfMaps creates a map of maps (nested maps)
// Initialize with keys "user1" and "user2", each containing empty string->int maps
func CreateMapOfMaps() map[string]map[string]int {
	// TODO(human): Create nested map structure with "user1" and "user2" keys
	return nil
}

// SetNestedValue sets a value in the nested map at [outerKey][innerKey]
// Creates nested map if outerKey doesn't exist
func SetNestedValue(m map[string]map[string]int, outerKey, innerKey string, value int) {
	// TODO(human): Initialize inner map if needed, then set value
}

// CreateMapOfStructs creates a map[int]Student with some initial data
// Add students: {1, "Alice", "A"}, {2, "Bob", "B"}, {3, "Charlie", "A"}
func CreateMapOfStructs() map[int]Student {
	// TODO(human): Create map with student data
	return nil
}

// GroupStudentsByGrade takes a slice of students and returns a map
// where keys are grades and values are slices of students with that grade
func GroupStudentsByGrade(students []Student) map[string][]Student {
	// TODO(human): Group students by grade
	return nil
}

// BuildAdjacencyList converts a slice of edges into a graph adjacency list
// Each edge is [from, to], result maps each node to its neighbors
// Example: [[1,2], [1,3], [2,3]] -> {1:[2,3], 2:[3], 3:[]}
func BuildAdjacencyList(edges [][2]int) map[int][]int {
	// TODO(human): Build graph from edges
	return nil
}
