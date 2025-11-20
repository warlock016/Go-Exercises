# Exercise 05: Connected Components

**Tier:** 2 (Application)
**Time:** 60 min
**Goal:** Identify and analyze disconnected subgraphs

## Functions

```go
func CountComponents(graph map[int][]int) int
func GetAllComponents(graph map[int][]int) [][]int
func IsConnected(graph map[int][]int) bool
func FindBridges(graph map[int][]int) [][2]int  // Edge removal disconnects graph
```

## Key Concept

Connected component = maximal set of vertices where each can reach all others.

**Example:**
```
0--1    2--3    4

Components: [[0,1], [2,3], [4]]
Count: 3
IsConnected: false
Bridge in {0,1}: edge (0,1) - removing it disconnects 0 and 1
```

## Hints

Use BFS/DFS from each unvisited vertex. Each traversal finds one component.

For bridges: An edge is a bridge if removing it increases component count.
