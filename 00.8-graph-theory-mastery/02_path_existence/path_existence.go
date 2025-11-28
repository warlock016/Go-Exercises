package path_existence

import (
	"slices"
)

// HasPath returns true if there's a path from start to end
func HasPath(graph map[int][]int, start, end int) bool {
	// TODO(human): Implement
	stack := []int{start}
	visited := map[int]bool{}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current == end {
			return true
		}

		if visited[current] {
			continue
		}

		visited[current] = true

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}

	return false
}

// FindPathBFS finds a path using BFS, returns nil if no path exists
func FindPathBFS(graph map[int][]int, start, end int) []int {
	// TODO(human): Implement

	result := []int{}
	parent := map[int]int{}
	visited := map[int]bool{}
	queue := []int{start}

	if start == end {
		return []int{start}
	}

	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			break
		}

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
				parent[neighbor] = current
			}
		}
	}

	current := end
	if _, exists := parent[end]; !exists { // end node has no parent
		return nil
	}

	for current != start {
		result = append(result, current)
		current = parent[current]
	}

	result = append(result, start)
	slices.Reverse(result)

	return result
}

// FindPathDFS finds a path using DFS, returns nil if no path exists
func FindPathDFS(graph map[int][]int, start, end int) []int {
	// TODO(human): Implement
	result := []int{}
	visited := map[int]bool{}
	stack := []int{}
	parent := map[int]int{}

	if start == end {
		return []int{start}
	}

	stack = append(stack, start)
	visited[start] = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current == end {
			break
		}

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				stack = append(stack, neighbor)
				parent[neighbor] = current
			}
		}
	}

	if _, exists := parent[end]; !exists {
		return nil
	}

	current := end
	for current != start {
		result = append(result, current)
		current = parent[current]
	}

	result = append(result, start)
	slices.Reverse(result)

	return result
}

/*
    0
   / \
  1   2
   \ /
    3

Adjacency List: {0:[1,2], 1:[3], 2:[3], 3:[]} Find all paths from 0 to 3:
*/

func dfsAllPaths(graph map[int][]int, current, end int, path []int, visited map[int]bool, result *[][]int) {

	if current == end {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
		// fmt.Printf("Reached end node. Path result: %v\n", result)
		return
	}

	visited[current] = true
	// fmt.Printf("Current node: %v -> neighbors: %v, visited: %v\n", current, graph[current], visited)

	for _, neighbor := range graph[current] {
		if !visited[neighbor] {
			path = append(path, neighbor)
			dfsAllPaths(graph, neighbor, end, path, visited, result)
			path = path[:len(path)-1]
			// fmt.Printf("Updated path in neighbor loop: %v\n", path)
		}
	}

	visited[current] = false
}

// FindAllPathsDFS finds ALL paths from start to end
func FindAllPathsDFS(graph map[int][]int, start, end int) [][]int {
	// TODO(human): Implement
	// fmt.Printf("Graph: %v\n", graph)

	var result [][]int
	currentPath := []int{start}
	visited := make(map[int]bool)

	dfsAllPaths(graph, start, end, currentPath, visited, &result)

	return result
}

// GetPathLength returns the number of edges in a path
func GetPathLength(path []int) int {
	// TODO(human): Implement

	if len(path) == 0 {
		return 0
	}

	return len(path) - 1
}
