package graph_properties

import (
	"slices"
)

// GetDegree returns the degree of a vertex in an undirected graph
func GetDegree(graph map[int][]int, vertex int) int {
	// TODO(human): Implement
	return len(graph[vertex])
}

// GetInDegree returns the in-degree of a vertex in a directed graph
func GetInDegree(graph map[int][]int, vertex int) int {
	// TODO(human): Implement

	/*
	   0: {1, 2}
	   1: {3}
	   2: {3}
	   3: {}
	*/

	result := 0

	for i, v := range graph {
		if i != vertex {
			for _, w := range v {
				if w == vertex {
					result++
				}
			}
		}
	}

	return result
}

// GetOutDegree returns the out-degree of a vertex in a directed graph
func GetOutDegree(graph map[int][]int, vertex int) int {
	// TODO(human): Implement

	result := 0

	for _, v := range graph[vertex] {
		if v != vertex {
			result++
		}
	}

	return result
}

// GetGraphDensity calculates the density of a graph
func GetGraphDensity(graph map[int][]int, numVertices int, directed bool) float64 {
	// TODO(human): Implement
	// directed graph max edges: n*(n-1)
	// undirected graph max edges: n*(n-1)/2

	numVtx := len(graph)

	var numEdges int

	for _, v := range graph {
		numEdges += len(v)
	}

	maxEdges := numVtx * (numVtx - 1)

	return float64(numEdges) / float64(maxEdges)
}

// IsDirected determines if a graph is directed
func IsDirected(graph map[int][]int) bool {
	// TODO(human): Implement

	for i, v := range graph {
		for _, w := range v {
			if !slices.Contains(graph[w], i) {
				return true
			}
		}
	}

	return false
}

// CountVertices returns the total number of vertices
func CountVertices(graph map[int][]int) int {
	// TODO(human): Implement
	return len(graph)
}

// CountEdges returns the total number of edges
func CountEdges(graph map[int][]int, directed bool) int {
	// TODO(human): Implement

	result := 0
	for _, v := range graph {
		result += len(v)
	}

	if !directed {
		result /= 2
	}

	return result
}
