# Exercise 11: Minimum Spanning Tree

**Tier:** 3 (Integration)
**Time:** 75 min
**Goal:** Find minimum cost to connect all vertices

## Functions

```go
func KruskalMST(edges [][3]int, numVertices int) ([][2]int, int)  // edges, total cost
func PrimMST(graph WeightedGraph, start int) ([][2]int, int)
func IsMST(graph WeightedGraph, tree [][2]int) bool  // Verify if given tree is MST
```

## Key Concept

MST = minimum cost to connect all vertices. Greedy algorithms:
- Kruskal: Sort edges by weight, add if doesn't create cycle (union-find)
- Prim: Grow tree from start, always add cheapest edge to tree

**Example:**
```
    1       2
  A---B---C
   \ /
   3

MST: {(A,B,1), (B,C,2)} cost=3  (skip A-B edge weight 3 - creates cycle)
```
