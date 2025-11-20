package graph_basics

import (
	"fmt"
	"slices"
)

// Adjacency List Functions

// NewAdjacencyList creates a new empty adjacency list
func NewAdjacencyList() map[int][]int {
	// TODO(human): Implement
	return make(map[int][]int)
}

// AddEdge adds a directed edge from source to destination in adjacency list
func AddEdge(graph map[int][]int, source, dest int) {
	// TODO(human): Implement
	if _, exists := graph[source]; !exists { // directed edge does not exist
		graph[source] = []int{}
		graph[source] = append(graph[source], dest)
	} else if !slices.Contains(graph[source], dest) {
		currDest := graph[source]
		currDest = append(currDest, dest) // add edge to graph
		graph[source] = currDest          // reassign updated slice to source graph
	}

}

// AddUndirectedEdge adds an undirected edge (bidirectional) between two vertices
func AddUndirectedEdge(graph map[int][]int, v1, v2 int) {
	// TODO(human): Implement
	// same as AddEdge but repeat twice to cover both directions: v1->v2 && v2->v1

	if _, exists := graph[v1]; !exists {
		graph[v1] = []int{}
		graph[v1] = append(graph[v1], v2)
	} else if !slices.Contains(graph[v1], v2) {
		currDest := graph[v1]
		currDest = append(currDest, v2)
		graph[v1] = currDest
	}

	if _, exists := graph[v2]; !exists {
		graph[v2] = []int{}
		graph[v2] = append(graph[v2], v1)
	} else if !slices.Contains(graph[v2], v1) {
		currDest := graph[v2]
		currDest = append(currDest, v1)
		graph[v2] = currDest
	}
}

// RemoveEdge removes a directed edge from source to destination
func RemoveEdge(graph map[int][]int, source, dest int) {
	// TODO(human): Implement
	if _, exists := graph[source]; exists {
		currDest := graph[source]

		newDest := make([]int, 0, len(currDest))

		for _, v := range currDest {
			if v != dest {
				newDest = append(newDest, v)
			}
		}

		graph[source] = newDest
	}
}

// HasEdge checks if a directed edge exists from source to destination
func HasEdge(graph map[int][]int, source, dest int) bool {
	// TODO(human): Implement

	if _, exists := graph[source]; exists {
		return slices.Contains(graph[source], dest)
	}

	return false
}

// GetNeighbors returns all neighbors of a vertex (empty slice if vertex doesn't exist)
func GetNeighbors(graph map[int][]int, vertex int) []int {
	// TODO(human): Implement
	result := []int{}
	if _, exists := graph[vertex]; exists {
		result = graph[vertex]
	}
	return result
}

// Adjacency Matrix Functions

// NewAdjacencyMatrix creates a new adjacency matrix for n vertices (initialized to false)
func NewAdjacencyMatrix(n int) [][]bool {
	// TODO(human): Implement
	result := make([][]bool, n)

	for i := range result {
		result[i] = make([]bool, n)
	}

	return result
}

// AddEdgeMatrix adds a directed edge from source to destination in adjacency matrix
func AddEdgeMatrix(matrix [][]bool, source, dest int) {
	// TODO(human): Implement
	if source < 0 || dest < 0 {
		return
	}

	if source >= len(matrix) {
		matrix = append(matrix, []bool{})
		matrix[source] = make([]bool, len(matrix))
	}

	matrix[source][dest] = true
}

// AddUndirectedEdgeMatrix adds an undirected edge between two vertices in matrix
func AddUndirectedEdgeMatrix(matrix [][]bool, v1, v2 int) {
	// TODO(human): Implement
	if v1 < 0 || v2 < 0 {
		return
	}

	// 3x3 -> v1 = 4, v2 = 5 => ext = 5 => ext - len(3x3) = 2
	ext := max(v1, v2)

	if ext >= len(matrix) {
		for i := range ext - len(matrix) {
			matrix = append(matrix, []bool{})
			matrix[len(matrix)-1+i] = append(matrix[len(matrix)-1+i], false)
		}
	}

	matrix[v1][v2] = true
	matrix[v2][v1] = true
}

// HasEdgeMatrix checks if a directed edge exists from source to destination in matrix
func HasEdgeMatrix(matrix [][]bool, source, dest int) bool {
	// TODO(human): Implement
	if source < len(matrix) && dest < len(matrix[source]) {
		return matrix[source][dest]
	}
	return false
}

// GetNeighborsMatrix returns all neighbors of a vertex from adjacency matrix
func GetNeighborsMatrix(matrix [][]bool, vertex int) []int {
	// TODO(human): Implement
	result := []int{}

	if vertex < len(matrix) {
		for i, v := range matrix[vertex] {
			if v {
				result = append(result, i)
			}
		}
	}

	return result
}

// Conversion Functions

// ListToMatrix converts an adjacency list to an adjacency matrix
// n is the number of vertices (assumes vertices are numbered 0 to n-1)
func ListToMatrix(graph map[int][]int, n int) [][]bool {
	// TODO(human): Implement

	result := make([][]bool, n)

	for i := range result {
		result[i] = make([]bool, n)
	}

	for i := range graph {
		for _, v := range graph[i] {
			result[i][v] = true
		}
	}

	return result
}

// MatrixToList converts an adjacency matrix to an adjacency list
func MatrixToList(matrix [][]bool) map[int][]int {
	// TODO(human): Implement

	result := map[int][]int{}

	for i, v := range matrix {
		for j, w := range v {
			if w {
				result[i] = append(result[i], j)
			}
		}
	}

	return result
}

// Traversal Functions

// BFS performs breadth-first search starting from start vertex
// Returns vertices in BFS order
func BFS(graph map[int][]int, start int) []int {
	// TODO(human): Implement

	/*
			Graph:
			0 → 1, 2
			1 → 3
			2 → 4

			 0
			/ \
		   1  2
		   |  |
		   3  4
	*/

	queue := []int{start}
	visited := make(map[int]bool)
	result := []int{}

	for len(queue) > 0 {

		fmt.Printf("Queue before pop: %v ", queue)

		current := queue[0]
		queue = queue[1:]

		fmt.Printf("frontier: %v, current: %d\n", queue, current)

		if visited[current] {
			continue
		}

		visited[current] = true
		result = append(result, current)

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

// DFS performs depth-first search starting from start vertex
// Returns vertices in DFS order
func DFS(graph map[int][]int, start int) []int {
	// TODO(human): Implement

	/*
		Graph:
		0 → 1, 2
		1 → 3
		2 → 4

		0 → 1 → 3
		0 → 2 → 4
	*/

	stack := []int{start}
	visited := make(map[int]bool)
	result := []int{}

	for len(stack) > 0 {

		fmt.Printf("Stack before pop: %v ", stack)

		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Printf("frontier: %v, current: %d\n", stack, current)

		if visited[current] {
			continue
		}

		visited[current] = true
		result = append(result, current)

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}

	return result
}
