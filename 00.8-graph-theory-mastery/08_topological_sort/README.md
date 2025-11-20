# Exercise 08: Topological Sort

**Tier:** 2 (Application)  
**Time:** 75 min
**Goal:** Order tasks respecting dependencies (DAG)

## Functions

```go
func TopologicalSort(graph map[int][]int) []int
func HasValidOrdering(graph map[int][]int) bool  // Is it a DAG?
func FindPrerequisites(graph map[int][]int, task int) []int  // All dependencies
```

## Key Concept

Topological sort = linear ordering where all edges go left→right.

**Example:**
```
put_on_socks → put_on_shoes
put_on_pants → put_on_shoes  
put_on_pants → put_on_belt

Valid order: [pants, belt, socks, shoes]
```

Use DFS: add to result AFTER exploring all descendants.
