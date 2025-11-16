package graph_basics

// Adjacency List Functions

// NewAdjacencyList creates a new empty adjacency list
func NewAdjacencyList() map[int][]int {
	// TODO(human): Implement
	return nil
}

// AddEdge adds a directed edge from source to destination in adjacency list
func AddEdge(graph map[int][]int, source, dest int) {
	// TODO(human): Implement
}

// AddUndirectedEdge adds an undirected edge (bidirectional) between two vertices
func AddUndirectedEdge(graph map[int][]int, v1, v2 int) {
	// TODO(human): Implement
}

// RemoveEdge removes a directed edge from source to destination
func RemoveEdge(graph map[int][]int, source, dest int) {
	// TODO(human): Implement
}

// HasEdge checks if a directed edge exists from source to destination
func HasEdge(graph map[int][]int, source, dest int) bool {
	// TODO(human): Implement
	return false
}

// GetNeighbors returns all neighbors of a vertex (empty slice if vertex doesn't exist)
func GetNeighbors(graph map[int][]int, vertex int) []int {
	// TODO(human): Implement
	return nil
}

// Adjacency Matrix Functions

// NewAdjacencyMatrix creates a new adjacency matrix for n vertices (initialized to false)
func NewAdjacencyMatrix(n int) [][]bool {
	// TODO(human): Implement
	return nil
}

// AddEdgeMatrix adds a directed edge from source to destination in adjacency matrix
func AddEdgeMatrix(matrix [][]bool, source, dest int) {
	// TODO(human): Implement
}

// AddUndirectedEdgeMatrix adds an undirected edge between two vertices in matrix
func AddUndirectedEdgeMatrix(matrix [][]bool, v1, v2 int) {
	// TODO(human): Implement
}

// HasEdgeMatrix checks if a directed edge exists from source to destination in matrix
func HasEdgeMatrix(matrix [][]bool, source, dest int) bool {
	// TODO(human): Implement
	return false
}

// GetNeighborsMatrix returns all neighbors of a vertex from adjacency matrix
func GetNeighborsMatrix(matrix [][]bool, vertex int) []int {
	// TODO(human): Implement
	return nil
}

// Conversion Functions

// ListToMatrix converts an adjacency list to an adjacency matrix
// n is the number of vertices (assumes vertices are numbered 0 to n-1)
func ListToMatrix(graph map[int][]int, n int) [][]bool {
	// TODO(human): Implement
	return nil
}

// MatrixToList converts an adjacency matrix to an adjacency list
func MatrixToList(matrix [][]bool) map[int][]int {
	// TODO(human): Implement
	return nil
}

// Traversal Functions

// BFS performs breadth-first search starting from start vertex
// Returns vertices in BFS order
func BFS(graph map[int][]int, start int) []int {
	// TODO(human): Implement
	return nil
}

// DFS performs depth-first search starting from start vertex
// Returns vertices in DFS order
func DFS(graph map[int][]int, start int) []int {
	// TODO(human): Implement
	return nil
}
