# Exercise 10: Advanced Maps

## 🎯 Learning Goal
Master complex map structures including maps of slices, maps of maps, maps of structs, and using maps to represent real-world data relationships like graphs and grouped data.

## 📝 Problem Description

While basic maps (`map[string]int`) are straightforward, real-world programs often require more sophisticated map structures:

- **Maps of slices** - Group multiple values under a single key (e.g., students per grade)
- **Maps of maps** - Nested relationships (e.g., sales by region by product)
- **Maps of structs** - Store complex objects with multiple fields
- **Graph representations** - Adjacency lists using maps

Understanding these patterns is essential because:
- They enable efficient data grouping and lookups
- They model real-world relationships naturally
- They're more memory-efficient than alternatives (e.g., nested loops)
- Many algorithms rely on these structures (graph traversal, data aggregation)

In this exercise, you'll implement common patterns using advanced map structures.

## 🔧 Function Signatures

Implement these functions in `advanced_maps.go`:

```go
// Student represents a student with a name and grade
type Student struct {
	Name  string
	Grade string
}

// CreateMapOfSlices creates a map that stores slices of integers
// Initialize with keys "even" and "odd", each with empty slices
func CreateMapOfSlices() map[string][]int

// AddToMapSlice appends a value to the slice at the given key
// Creates the slice if the key doesn't exist
func AddToMapSlice(m map[string][]int, key string, value int)

// CreateMapOfMaps creates a map of maps (nested maps)
// Initialize with keys "user1" and "user2", each containing empty string->int maps
func CreateMapOfMaps() map[string]map[string]int

// SetNestedValue sets a value in the nested map at [outerKey][innerKey]
// Creates nested map if outerKey doesn't exist
func SetNestedValue(m map[string]map[string]int, outerKey, innerKey string, value int)

// CreateMapOfStructs creates a map[int]Student with some initial data
// Add students: {1, "Alice", "A"}, {2, "Bob", "B"}, {3, "Charlie", "A"}
func CreateMapOfStructs() map[int]Student

// GroupStudentsByGrade takes a slice of students and returns a map
// where keys are grades and values are slices of students with that grade
func GroupStudentsByGrade(students []Student) map[string][]Student

// BuildAdjacencyList converts a slice of edges into a graph adjacency list
// Each edge is [from, to], result maps each node to its neighbors
// Example: [[1,2], [1,3], [2,3]] -> {1:[2,3], 2:[3], 3:[]}
func BuildAdjacencyList(edges [][2]int) map[int][]int
```

## 💡 Examples

```go
// Maps of slices
m := CreateMapOfSlices()  // {"even": [], "odd": []}
AddToMapSlice(m, "even", 2)
AddToMapSlice(m, "even", 4)
AddToMapSlice(m, "odd", 1)
// m = {"even": [2, 4], "odd": [1]}

// Maps of maps
mm := CreateMapOfMaps()  // {"user1": {}, "user2": {}}
SetNestedValue(mm, "user1", "score", 100)
SetNestedValue(mm, "user1", "level", 5)
// mm = {"user1": {"score": 100, "level": 5}, "user2": {}}

// Maps of structs
students := CreateMapOfStructs()
alice := students[1]  // Student{Name: "Alice", Grade: "A"}

// Grouping
allStudents := []Student{
    {"Alice", "A"},
    {"Bob", "B"},
    {"Charlie", "A"},
}
grouped := GroupStudentsByGrade(allStudents)
// {"A": [{"Alice", "A"}, {"Charlie", "A"}], "B": [{"Bob", "B"}]}

// Graph adjacency list
edges := [][2]int{{1, 2}, {1, 3}, {2, 3}}
graph := BuildAdjacencyList(edges)
// {1: [2, 3], 2: [3], 3: []}
```

## 📋 Instructions

1. **CreateMapOfSlices:** Use `make(map[string][]int)` and initialize with empty slices
2. **AddToMapSlice:** Check if key exists, create slice if needed, then append
3. **CreateMapOfMaps:** Use `make(map[string]map[string]int)` with nested initialization
4. **SetNestedValue:** Check if outer key exists, create inner map if needed, then set value
5. **CreateMapOfStructs:** Use map literal or make, then add Student structs with given data
6. **GroupStudentsByGrade:** Iterate through students, group by grade using AddToMapSlice pattern
7. **BuildAdjacencyList:** Create map, add each edge's destination to source's neighbor list

## 🧪 Testing

Run tests with:
```bash
go test -v
```

Expected test count: ~30-35 tests across all functions

## 🤔 Think About

1. **Why use `map[string][]int` instead of `map[string]int`?**
   - Single key can map to multiple values
   - Example: Tag -> list of item IDs

2. **What's the difference between nil map and empty map?**
   - Nil map: `var m map[string]int` (cannot add to it, will panic)
   - Empty map: `make(map[string]int)` (ready to use)

3. **Why check if key exists before appending to map of slices?**
   - If key doesn't exist, the slice is nil
   - Appending to nil slice works, but explicit initialization is clearer

4. **When would you use a map of structs vs slice of structs?**
   - Map: O(1) lookup by ID, unique keys
   - Slice: Ordered, can have duplicates, O(n) lookup

## 💡 Hints

<details>
<summary>Hint 1: Maps of slices pattern</summary>

The standard pattern for maps of slices:

```go
// Initialize map
m := make(map[string][]int)

// Adding to a key that might not exist
if _, exists := m[key]; !exists {
    m[key] = []int{}  // Initialize slice
}
m[key] = append(m[key], value)

// Or simpler (append to nil slice works):
m[key] = append(m[key], value)
```

**Note:** Appending to a nil slice is valid and creates a new slice!
</details>

<details>
<summary>Hint 2: Maps of maps pattern</summary>

Nested maps require two-level initialization:

```go
// Create outer map
m := make(map[string]map[string]int)

// Before accessing inner map, check and create if needed
if m[outerKey] == nil {
    m[outerKey] = make(map[string]int)
}

// Now safe to set
m[outerKey][innerKey] = value
```
</details>

<details>
<summary>Hint 3: Grouping with maps</summary>

Grouping pattern:

```go
result := make(map[string][]Student)

for _, student := range students {
    grade := student.Grade
    result[grade] = append(result[grade], student)
}

return result
```

This works because `append` handles nil slices automatically.
</details>

<details>
<summary>Hint 4: Building adjacency lists</summary>

For graphs, ensure all nodes have entries (even if no outgoing edges):

```go
graph := make(map[int][]int)

for _, edge := range edges {
    from, to := edge[0], edge[1]
    graph[from] = append(graph[from], to)

    // Ensure destination node exists in map
    if _, exists := graph[to]; !exists {
        graph[to] = []int{}
    }
}
```
</details>

<details>
<summary>Full Solution</summary>

```go
package advanced_maps

type Student struct {
	Name  string
	Grade string
}

func CreateMapOfSlices() map[string][]int {
	m := make(map[string][]int)
	m["even"] = []int{}
	m["odd"] = []int{}
	return m
}

func AddToMapSlice(m map[string][]int, key string, value int) {
	m[key] = append(m[key], value)
}

func CreateMapOfMaps() map[string]map[string]int {
	m := make(map[string]map[string]int)
	m["user1"] = make(map[string]int)
	m["user2"] = make(map[string]int)
	return m
}

func SetNestedValue(m map[string]map[string]int, outerKey, innerKey string, value int) {
	if m[outerKey] == nil {
		m[outerKey] = make(map[string]int)
	}
	m[outerKey][innerKey] = value
}

func CreateMapOfStructs() map[int]Student {
	return map[int]Student{
		1: {Name: "Alice", Grade: "A"},
		2: {Name: "Bob", Grade: "B"},
		3: {Name: "Charlie", Grade: "A"},
	}
}

func GroupStudentsByGrade(students []Student) map[string][]Student {
	result := make(map[string][]Student)
	for _, student := range students {
		result[student.Grade] = append(result[student.Grade], student)
	}
	return result
}

func BuildAdjacencyList(edges [][2]int) map[int][]int {
	graph := make(map[int][]int)

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)

		// Ensure destination node exists
		if _, exists := graph[to]; !exists {
			graph[to] = []int{}
		}
	}

	return graph
}
```
</details>

## 🎓 What This Teaches

- **Maps of slices** - Storing multiple values per key (one-to-many relationships)
- **Maps of maps** - Nested data structures for hierarchical data
- **Maps of structs** - Storing complex objects with O(1) lookup
- **Nil handling** - Understanding nil slices vs nil maps vs empty collections
- **Grouping operations** - Using maps to aggregate data by key
- **Graph representations** - Adjacency lists for efficient graph algorithms
- **Existence checks** - Using the comma-ok idiom for nested initialization
- **Common patterns** - Real-world data structure techniques used in production code

---

**Next Exercise:** `11_data_modeling` - Design a complete library system using structs and maps
