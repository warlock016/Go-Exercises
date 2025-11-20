# Exercise 07: Dijkstra's Algorithm

**Tier:** 2 (Application)
**Time:** 90 min
**Goal:** Find shortest path in weighted graphs

## Functions

```go
type WeightedGraph map[int]map[int]int  // vertex -> (neighbor -> weight)

func Dijkstra(graph WeightedGraph, start int) map[int]int  // vertex -> distance
func DijkstraPath(graph WeightedGraph, start, end int) ([]int, int)  // path, cost
```

## Key Concept

Dijkstra = BFS + priority queue. Always process lowest-cost node next.

**Example:**
```
    A --1-- B
    |       |
    5       2
    |       |
    C --1-- D

Dijkstra(A): {A:0, B:1, C:5, D:3}
Shortest A→D: A→B→D (cost 3), not A→C→D (cost 6)
```

Use priority queue (min-heap) to always explore cheapest unvisited node.
