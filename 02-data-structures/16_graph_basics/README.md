# Exercise 16: Graph Basics

## 🎯 Learning Goal
Master fundamental graph representations and traversals. Learn to represent graphs using adjacency lists and adjacency matrices, implement basic graph operations, and perform BFS (Breadth-First Search) and DFS (Depth-First Search) traversals.

## 📝 Problem Description

Graphs are one of the most important data structures in computer science, used for modeling networks, relationships, dependencies, paths, and more. A graph consists of vertices (nodes) and edges (connections between nodes).

There are two primary ways to represent graphs in memory:
1. **Adjacency List** - `map[int][]int` - Each vertex maps to a list of its neighbors
2. **Adjacency Matrix** - `[][]bool` - 2D array where `matrix[i][j]` is true if edge exists from i to j

In this exercise, you'll implement both representations and basic graph algorithms.

## 🔧 Function Signatures

Implement these functions in `graph_basics.go`:

```go
// Adjacency List Functions

// NewAdjacencyList creates a new empty adjacency list
func NewAdjacencyList() map[int][]int

// AddEdge adds a directed edge from source to destination in adjacency list
func AddEdge(graph map[int][]int, source, dest int)

// AddUndirectedEdge adds an undirected edge (bidirectional) between two vertices
func AddUndirectedEdge(graph map[int][]int, v1, v2 int)

// RemoveEdge removes a directed edge from source to destination
func RemoveEdge(graph map[int][]int, source, dest int)

// HasEdge checks if a directed edge exists from source to destination
func HasEdge(graph map[int][]int, source, dest int) bool

// GetNeighbors returns all neighbors of a vertex (empty slice if vertex doesn't exist)
func GetNeighbors(graph map[int][]int, vertex int) []int

// Adjacency Matrix Functions

// NewAdjacencyMatrix creates a new adjacency matrix for n vertices (initialized to false)
func NewAdjacencyMatrix(n int) [][]bool

// AddEdgeMatrix adds a directed edge from source to destination in adjacency matrix
func AddEdgeMatrix(matrix [][]bool, source, dest int)

// AddUndirectedEdgeMatrix adds an undirected edge between two vertices in matrix
func AddUndirectedEdgeMatrix(matrix [][]bool, v1, v2 int)

// HasEdgeMatrix checks if a directed edge exists from source to destination in matrix
func HasEdgeMatrix(matrix [][]bool, source, dest int) bool

// GetNeighborsMatrix returns all neighbors of a vertex from adjacency matrix
func GetNeighborsMatrix(matrix [][]bool, vertex int) []int

// Conversion Functions

// ListToMatrix converts an adjacency list to an adjacency matrix
// n is the number of vertices (assumes vertices are numbered 0 to n-1)
func ListToMatrix(graph map[int][]int, n int) [][]bool

// MatrixToList converts an adjacency matrix to an adjacency list
func MatrixToList(matrix [][]bool) map[int][]int

// Traversal Functions

// BFS performs breadth-first search starting from start vertex
// Returns vertices in BFS order
func BFS(graph map[int][]int, start int) []int

// DFS performs depth-first search starting from start vertex
// Returns vertices in DFS order
func DFS(graph map[int][]int, start int) []int
```

## 💡 Examples

```go
// Create adjacency list
graph := NewAdjacencyList()
AddEdge(graph, 0, 1)  // 0 -> 1
AddEdge(graph, 0, 2)  // 0 -> 2
AddEdge(graph, 1, 2)  // 1 -> 2
AddEdge(graph, 2, 3)  // 2 -> 3

HasEdge(graph, 0, 1)  // true
HasEdge(graph, 1, 0)  // false (directed graph)
GetNeighbors(graph, 0) // [1, 2]

// Undirected edge
AddUndirectedEdge(graph, 3, 4)  // Adds both 3->4 and 4->3

// Adjacency matrix (4 vertices)
matrix := NewAdjacencyMatrix(4)
AddEdgeMatrix(matrix, 0, 1)
AddEdgeMatrix(matrix, 1, 2)
HasEdgeMatrix(matrix, 0, 1)  // true

// Conversion
graph := map[int][]int{0: {1, 2}, 1: {2}}
matrix := ListToMatrix(graph, 3)  // [[false, true, true], [false, false, true], [false, false, false]]

// BFS traversal
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
BFS(graph, 0)  // [0, 1, 2, 3, 4] (level-by-level)

// DFS traversal
DFS(graph, 0)  // [0, 1, 3, 2, 4] (depth-first)
```

## 📋 Instructions

1. **Adjacency List Operations:**
   - NewAdjacencyList: Return `make(map[int][]int)`
   - AddEdge: Append dest to `graph[source]` slice
   - HasEdge: Check if dest is in `graph[source]` slice
   - GetNeighbors: Return `graph[vertex]` (or empty slice if doesn't exist)

2. **Adjacency Matrix Operations:**
   - NewAdjacencyMatrix: Create n×n 2D slice initialized to false
   - AddEdgeMatrix: Set `matrix[source][dest] = true`
   - HasEdgeMatrix: Return `matrix[source][dest]`
   - GetNeighborsMatrix: Iterate through row, collect indices where value is true

3. **Conversion:**
   - ListToMatrix: Create matrix, iterate over map, set matrix[source][dest] = true for each edge
   - MatrixToList: Create map, iterate through matrix, add edges where value is true

4. **BFS (Breadth-First Search):**
   - Use a queue (slice) to process vertices level by level
   - Use a visited map to avoid processing vertices twice
   - Start with the start vertex in the queue
   - While queue not empty: dequeue, add neighbors to queue if not visited

5. **DFS (Depth-First Search):**
   - Use recursion or a stack to explore as deep as possible before backtracking
   - Use a visited map to avoid infinite loops
   - Process vertex, then recursively process unvisited neighbors

## 🧪 Testing

Run tests with:
```bash
go test -v
```

Expected test count: ~40-50 tests covering all functions

## 🤔 Think About

1. **When to use adjacency list vs adjacency matrix?**
   - **List**: Sparse graphs (few edges), memory efficient, faster to iterate neighbors
   - **Matrix**: Dense graphs (many edges), O(1) edge lookup, easier for some algorithms

2. **What's the difference between BFS and DFS?**
   - **BFS**: Explores level-by-level (nearest first), uses queue, finds shortest path
   - **DFS**: Explores depth-first (as far as possible), uses stack/recursion, good for topological sort

3. **Why do we need a visited map?**
   - Prevents infinite loops in cyclic graphs
   - Ensures each vertex is processed only once
   - Critical for correctness of both BFS and DFS

4. **Directed vs Undirected graphs?**
   - **Directed**: Edge has direction (A→B doesn't mean B→A)
   - **Undirected**: Edge is bidirectional (A-B means A→B and B→A)
   - For undirected: add edges in both directions

## 💡 Hints

<details>
<summary>Hint 1: Adjacency list structure</summary>

```go
// Adjacency list is a map from vertex to its neighbors
graph := make(map[int][]int)

// Add edge 0 -> 1
graph[0] = append(graph[0], 1)

// Add edge 0 -> 2
graph[0] = append(graph[0], 2)

// Now graph[0] = [1, 2]

// Check if edge exists
func HasEdge(graph map[int][]int, source, dest int) bool {
    neighbors, exists := graph[source]
    if !exists {
        return false
    }
    for _, neighbor := range neighbors {
        if neighbor == dest {
            return true
        }
    }
    return false
}
```
</details>

<details>
<summary>Hint 2: Adjacency matrix structure</summary>

```go
// Create n×n matrix
func NewAdjacencyMatrix(n int) [][]bool {
    matrix := make([][]bool, n)
    for i := range matrix {
        matrix[i] = make([]bool, n)
    }
    return matrix
}

// Add edge
matrix[0][1] = true  // Edge from 0 to 1

// Check edge
hasEdge := matrix[0][1]  // O(1) lookup!

// Get neighbors
func GetNeighborsMatrix(matrix [][]bool, vertex int) []int {
    var neighbors []int
    for dest := range matrix[vertex] {
        if matrix[vertex][dest] {
            neighbors = append(neighbors, dest)
        }
    }
    return neighbors
}
```
</details>

<details>
<summary>Hint 3: BFS implementation</summary>

```go
func BFS(graph map[int][]int, start int) []int {
    visited := make(map[int]bool)
    queue := []int{start}
    result := []int{}

    visited[start] = true

    for len(queue) > 0 {
        // Dequeue
        vertex := queue[0]
        queue = queue[1:]

        result = append(result, vertex)

        // Enqueue unvisited neighbors
        for _, neighbor := range graph[vertex] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }

    return result
}
```

Key: Use queue for level-by-level exploration
</details>

<details>
<summary>Hint 4: DFS implementation</summary>

```go
func DFS(graph map[int][]int, start int) []int {
    visited := make(map[int]bool)
    result := []int{}

    var dfsHelper func(int)
    dfsHelper = func(vertex int) {
        visited[vertex] = true
        result = append(result, vertex)

        for _, neighbor := range graph[vertex] {
            if !visited[neighbor] {
                dfsHelper(neighbor)  // Recursive call
            }
        }
    }

    dfsHelper(start)
    return result
}
```

Key: Use recursion (or explicit stack) to explore depth-first
</details>

<details>
<summary>Full Solution (partial)</summary>

```go
package graph_basics

func NewAdjacencyList() map[int][]int {
    return make(map[int][]int)
}

func AddEdge(graph map[int][]int, source, dest int) {
    graph[source] = append(graph[source], dest)
}

func AddUndirectedEdge(graph map[int][]int, v1, v2 int) {
    AddEdge(graph, v1, v2)
    AddEdge(graph, v2, v1)
}

func HasEdge(graph map[int][]int, source, dest int) bool {
    neighbors, exists := graph[source]
    if !exists {
        return false
    }
    for _, neighbor := range neighbors {
        if neighbor == dest {
            return true
        }
    }
    return false
}

func NewAdjacencyMatrix(n int) [][]bool {
    matrix := make([][]bool, n)
    for i := range matrix {
        matrix[i] = make([]bool, n)
    }
    return matrix
}

func BFS(graph map[int][]int, start int) []int {
    visited := make(map[int]bool)
    queue := []int{start}
    result := []int{}
    visited[start] = true

    for len(queue) > 0 {
        vertex := queue[0]
        queue = queue[1:]
        result = append(result, vertex)

        for _, neighbor := range graph[vertex] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
    return result
}

// ... (rest of implementation)
```
</details>

## 🎓 What This Teaches

- **Graph representations** - Adjacency lists vs adjacency matrices, trade-offs
- **Graph terminology** - Vertices, edges, directed vs undirected, neighbors
- **Adjacency list** - Using `map[int][]int` for sparse graphs, dynamic structure
- **Adjacency matrix** - Using `[][]bool` for dense graphs, O(1) lookups
- **BFS algorithm** - Queue-based level-order traversal, shortest path foundation
- **DFS algorithm** - Stack/recursion-based depth-first exploration
- **Visited tracking** - Using maps to prevent infinite loops and duplicate processing
- **Data structure conversion** - Transforming between representations
- **Algorithm complexity** - Understanding space/time trade-offs in graph representations
- **Real-world applications** - Social networks, routing, dependency resolution

---

**Next Exercise:** `17_capstone_mini_database` - In-memory key-value store combining all concepts
