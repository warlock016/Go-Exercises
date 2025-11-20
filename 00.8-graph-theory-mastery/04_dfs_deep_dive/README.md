# Exercise 04: DFS Deep Dive

**Tier:** 1 (Foundation)
**Estimated Time:** 75 minutes
**Learning Goal:** Master DFS applications - cycle detection, topological prerequisites, exhaustive search

---

## Learning Objective

DFS excels at problems requiring **depth-first exploration** and **backtracking**. This exercise teaches you when DFS is the right tool: detecting cycles, finding all solutions, and exploring paths completely before backtracking.

---

## Function Signatures

```go
// HasCycle detects if an undirected graph contains a cycle
func HasCycle(graph map[int][]int) bool

// HasCycleDirected detects if a directed graph contains a cycle
func HasCycleDirected(graph map[int][]int) bool

// FindCycle returns vertices forming a cycle, or nil if no cycle exists
func FindCycle(graph map[int][]int) []int

// CountPathsDFS counts all paths from start to end
func CountPathsDFS(graph map[int][]int, start, end int) int

// LongestPath finds the longest simple path from start to end (no repeated vertices)
func LongestPath(graph map[int][]int, start, end int) []int
```

---

## Examples

**Cycle Detection:**
```
Undirected:     0 --- 1
                |     |      ← Cycle: 0-1-2-0
                2 ----+

HasCycle(graph) → true

Acyclic:        0 --- 1 --- 2
HasCycle(graph) → false
```

**Directed Cycle:**
```
Has cycle:      0 → 1 → 2
                ↑_______↓     ← Cycle: 0→1→2→0

HasCycleDirected(graph) → true
```

---

## Hints

<details>
<summary>Hint: Cycle Detection (Undirected)</summary>

Use DFS. A cycle exists if you visit a node that's already visited AND it's not your parent.

```go
func dfs(current, parent int, visited map[int]bool) bool {
    visited[current] = true
    for _, neighbor := range graph[current] {
        if !visited[neighbor] {
            if dfs(neighbor, current, visited) {
                return true  // Cycle found in recursion
            }
        } else if neighbor != parent {
            return true  // Visited non-parent = cycle!
        }
    }
    return false
}
```
</details>

<details>
<summary>Hint: Cycle Detection (Directed)</summary>

Use three states:
- WHITE: Unvisited
- GRAY: Currently exploring (in recursion stack)
- BLACK: Fully explored

Cycle exists if you encounter a GRAY node (back edge).
</details>

---

## What This Teaches

✅ **Cycle detection** with DFS
✅ **Backtracking** for exhaustive search
✅ **DFS = stack** (recursion or explicit)
✅ **When DFS > BFS** (all paths, cycles, deep exploration)
