# Exercise 12: Strongly Connected Components

**Tier:** 3 (Integration)
**Time:** 90 min
**Goal:** Find strongly connected components (SCCs) in directed graphs

## Functions

```go
func KosarajuSCC(graph map[int][]int) [][]int  // Returns all SCCs
func CountSCC(graph map[int][]int) int
func IsStronglyConnected(graph map[int][]int) bool
```

## Key Concept

SCC = maximal set of vertices where each can reach all others (directed graphs only).

**Kosaraju's Algorithm:**
1. DFS on original graph, record finish order
2. Create transpose graph (reverse all edges)
3. DFS on transpose in reverse finish order
4. Each DFS tree = one SCC

**Example:**
```
0→1→2
↑   ↓
└───3

SCCs: [[0,1,2,3]] - all can reach each other
```
