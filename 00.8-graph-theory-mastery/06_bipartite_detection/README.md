# Exercise 06: Bipartite Detection

**Tier:** 2 (Application)
**Time:** 60 min
**Goal:** Determine if a graph can be 2-colored (bipartite)

## Functions

```go
func IsBipartite(graph map[int][]int) bool
func TwoColor(graph map[int][]int) (map[int]int, bool) // Returns coloring or nil
```

## Key Concept

Bipartite = vertices can be divided into two sets where edges only connect different sets.

**Example:**
```
Bipartite:     0 --- 1
               |     |
               2 --- 3

Set A: {0, 3}  Set B: {1, 2}  (No edges within sets)

Not Bipartite: 0 --- 1 --- 2
                \---------/  (Triangle - odd cycle)
```

Use BFS/DFS with 2-coloring. If neighbor has same color = not bipartite.
