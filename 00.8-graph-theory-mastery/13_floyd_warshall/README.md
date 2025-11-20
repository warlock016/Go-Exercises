# Exercise 13: Floyd-Warshall Algorithm

**Tier:** 4 (Mastery)
**Time:** 90 min
**Goal:** All-pairs shortest path using dynamic programming

## Functions

```go
func FloydWarshall(graph WeightedGraph, numVertices int) [][]int
func HasNegativeCycle(distances [][]int) bool
func GetShortestPath(graph WeightedGraph, distances [][]int, start, end int) []int
func TransitiveClosure(graph map[int][]int) [][]bool  // Can i reach j?
```

## Key Concept

Floyd-Warshall computes shortest paths between ALL pairs of vertices in O(V³).

**Algorithm:**
```
for k in vertices:
    for i in vertices:
        for j in vertices:
            dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
```

Try all intermediate vertices k - can we improve path i→j by going through k?

**Example:**
```
    1       2
  A---B---C

Initial: A→C = ∞
Via B: A→B→C = 1+2 = 3  ✓ Better!
```
