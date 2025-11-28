package bfs_deep_dive

import "fmt"

// ShortestDistance returns the shortest distance from start to end, or -1 if no path exists
func ShortestDistance(graph map[int][]int, start, end int) int {
	// TODO(human): Implement

	path := []int{}
	parent := make(map[int]int)
	visited := map[int]bool{}
	queue := []int{start}

	if start == end {
		return 0
	}

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

	// fmt.Printf("graph: %v\nparent map: %v\nstart: %d\nend: %d\n", graph, parent, start, end)

	if _, exists := parent[end]; !exists {
		return -1
	}

	current := end

	for start != current {
		path = append(path, current)
		current = parent[current]
	}

	path = append(path, start) // we can append initial node, then return len(path) -1 OR omit appending it and return len(path)
	// slices.Reverse(path) 		// no need to reverse order since we only care about distance

	// fmt.Printf("Path: %v\n", path)

	return len(path) - 1
}

// LevelOrder returns vertices grouped by their distance from start
func LevelOrder(graph map[int][]int, start int) [][]int {
	// TODO(human): Implement

	fmt.Printf("Graph: %v\nstart: %d\n", graph, start)
	visited := make(map[int]bool)
	parent := map[int]int{}
	queue := []int{start}
	// result := make([][]int, 0)

	// level := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// result = append(result, []int{})
		// result[level] = make([]int, 0)
		// result[level] = append(result[level], neighbor)

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
				parent[neighbor] = current
				// level++
			}
		}

		// level++ // tracks queue iteration
		// if !visited[current] {
		// 	level++
		// }
		// fmt.Printf("Current node: %d || Level: %d\n", current, level)
	}

	// fmt.Printf("Parent map: %v\n", parent)
	// fmt.Printf("Level 2D array: %v\n", result)
	// fmt.Printf("Levels: %v\n", result)

	return nil
}

// FindClosestVertex finds the closest vertex satisfying a condition
func FindClosestVertex(graph map[int][]int, start int, condition func(int) bool) (vertex int, distance int) {
	// TODO(human): Implement
	return -1, -1
}

// CountComponentsBFS counts connected components using BFS
func CountComponentsBFS(graph map[int][]int) int {
	// TODO(human): Implement
	return 0
}

// GetComponentBFS returns all vertices in the same component as start
func GetComponentBFS(graph map[int][]int, start int) []int {
	// TODO(human): Implement
	return nil
}
