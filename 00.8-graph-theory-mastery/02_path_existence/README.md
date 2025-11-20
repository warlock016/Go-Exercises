# Exercise 02: Path Existence

**Tier:** 1 (Foundation)
**Estimated Time:** 60 minutes
**Learning Goal:** Apply BFS/DFS to answer "is there a path?" and reconstruct paths

---

## Learning Objective

You already know HOW to implement BFS and DFS from Module 02. Now you'll learn **WHY** and **WHEN** to use them by solving path-finding problems.

**Key Insight:** BFS and DFS aren't just traversal algorithms - they're **search** algorithms. The difference between "visiting nodes" and "answering questions" is what separates memorization from understanding.

---

## Problem Description

Implement functions to determine path existence and reconstruct paths between vertices. You'll use both BFS and DFS to see how they differ in the paths they find.

---

## Function Signatures

```go
// HasPath returns true if there's a path from start to end
func HasPath(graph map[int][]int, start, end int) bool

// FindPathBFS finds a path from start to end using BFS
// Returns the path as a slice of vertices, or nil if no path exists
func FindPathBFS(graph map[int][]int, start, end int) []int

// FindPathDFS finds a path from start to end using DFS
// Returns the path as a slice of vertices, or nil if no path exists
func FindPathDFS(graph map[int][]int, start, end int) []int

// FindAllPathsDFS finds ALL paths from start to end using DFS with backtracking
// Returns a slice of paths, where each path is a slice of vertices
func FindAllPathsDFS(graph map[int][]int, start, end int) [][]int

// GetPathLength returns the length (number of edges) of a path
func GetPathLength(path []int) int
```

---

## Examples

### Example 1: Simple Path
```
Graph:
    0 → 1 → 2
    ↓       ↓
    3 → 4 → 5

HasPath(graph, 0, 5) → true
FindPathBFS(graph, 0, 5) → [0, 1, 2, 5] (shortest path)
FindPathDFS(graph, 0, 5) → [0, 3, 4, 5] (depends on neighbor order)

FindAllPathsDFS(graph, 0, 5) → [
    [0, 1, 2, 5],
    [0, 3, 4, 5],
]

GetPathLength([0, 1, 2, 5]) → 3 (3 edges)
```

### Example 2: No Path (Disconnected Graph)
```
Graph:
    0 → 1       3 → 4

HasPath(graph, 0, 4) → false (disconnected)
FindPathBFS(graph, 0, 4) → nil
FindPathDFS(graph, 0, 4) → nil
```

### Example 3: BFS vs DFS Paths
```
Graph (undirected):
        0
       /|\
      1 2 3
      | | |
      4 5 6

BFS(0, 4) → [0, 1, 4] (shortest: 2 edges)
DFS(0, 4) → [0, 3, 6, ...] or [0, 1, 4] (depends on order)

Why? BFS explores level-by-level, DFS goes deep first!
```

---

## Instructions

1. **HasPath:** Use either BFS or DFS to check if end is reachable from start
2. **FindPathBFS:** Track parent pointers during BFS to reconstruct the path
3. **FindPathDFS:** Track parent pointers during DFS to reconstruct the path
4. **FindAllPathsDFS:** Use DFS with backtracking to explore all possible paths
5. Handle edge cases: start == end, vertices don't exist, disconnected graph

---

## Visualization Exercise

**Before coding**, trace by hand:

**Graph:**
```
    0
   / \
  1   2
  |   |
  3   4
   \ /
    5
```

**Question 1:** Find path from 0 to 5 using BFS
- Level 0: [0]
- Level 1: [1, 2] (neighbors of 0)
- Level 2: [3, 4] (neighbors of 1,2)
- Level 3: [5] (neighbors of 3,4) ← FOUND!
- Reconstruct path: How do you know 5 came from 3 or 4?
- **Answer:** Track parents! When you add 5 to queue, remember parent[5] = 3

**Question 2:** Find ALL paths from 0 to 5 using DFS
- Path 1: 0 → 1 → 3 → 5
- Path 2: 0 → 2 → 4 → 5
- How do you find both? **Backtracking!**

---

## Hints

<details>
<summary>Hint 1: Path Reconstruction</summary>

To reconstruct a path, you need to **track where you came from** during traversal.

**Idea:** Use a `parent` map
```go
parent := map[int]int{}
parent[neighbor] = current  // neighbor came from current
```

**After finding the end vertex, walk backwards:**
```
path = []
current = end
while current != start:
    path = [current] + path  // Prepend
    current = parent[current]
path = [start] + path
```

</details>

<details>
<summary>Hint 2: FindPathBFS Structure</summary>

```go
func FindPathBFS(graph map[int][]int, start, end int) []int {
    if start == end {
        return []int{start}
    }

    visited := make(map[int]bool)
    parent := make(map[int]int)
    queue := []int{start}

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if current == end {
            // Found! Reconstruct path using parent map
            return reconstructPath(parent, start, end)
        }

        visited[current] = true

        for _, neighbor := range graph[current] {
            if !visited[neighbor] {
                parent[neighbor] = current  // Track parent!
                queue = append(queue, neighbor)
            }
        }
    }

    return nil  // No path found
}
```

</details>

<details>
<summary>Hint 3: FindAllPathsDFS (Backtracking)</summary>

Finding ALL paths requires exploring every possibility, then **undoing** choices (backtracking).

**Key idea:** Pass current path as parameter, add to results when reaching end

```go
func FindAllPathsDFS(graph map[int][]int, start, end int) [][]int {
    var paths [][]int
    currentPath := []int{start}
    visited := make(map[int]bool)
    dfsAllPaths(graph, start, end, currentPath, visited, &paths)
    return paths
}

func dfsAllPaths(graph map[int][]int, current, end int, path []int, visited map[int]bool, result *[][]int) {
    if current == end {
        // Found a path! Add a COPY to results
        pathCopy := make([]int, len(path))
        copy(pathCopy, path)
        *result = append(*result, pathCopy)
        return
    }

    visited[current] = true

    for _, neighbor := range graph[current] {
        if !visited[neighbor] {
            path = append(path, neighbor)  // Add to path
            dfsAllPaths(graph, neighbor, end, path, visited, result)
            path = path[:len(path)-1]  // BACKTRACK: Remove from path
        }
    }

    visited[current] = false  // BACKTRACK: Unmark visited
}
```

**Why backtrack?** To allow revisiting nodes in DIFFERENT paths!

</details>

<details>
<summary>Hint 4: BFS vs DFS for Shortest Path</summary>

**BFS guarantees shortest path** in unweighted graphs because it explores level-by-level.

**DFS does NOT guarantee shortest path** - it finds A path, but not necessarily the shortest.

**Example:**
```
    0
   / \
  1   2
   \ /
    3

BFS(0,3): [0,1,3] or [0,2,3] (length 2) ✓ shortest
DFS(0,3): Could be [0,1,3] or [0,2,3] (length 2) ✓
          Or [0,1,2,3] if unlucky order (length 3) ✗
```

</details>

<details>
<summary>Complete Solution: FindPathBFS</summary>

```go
func HasPath(graph map[int][]int, start, end int) bool {
    if start == end {
        return true
    }

    visited := make(map[int]bool)
    queue := []int{start}

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if current == end {
            return true
        }

        if visited[current] {
            continue
        }

        visited[current] = true

        for _, neighbor := range graph[current] {
            if !visited[neighbor] {
                queue = append(queue, neighbor)
            }
        }
    }

    return false
}

func FindPathBFS(graph map[int][]int, start, end int) []int {
    if start == end {
        return []int{start}
    }

    visited := make(map[int]bool)
    parent := make(map[int]int)
    queue := []int{start}
    visited[start] = true

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if current == end {
            return reconstructPath(parent, start, end)
        }

        for _, neighbor := range graph[current] {
            if !visited[neighbor] {
                visited[neighbor] = true
                parent[neighbor] = current
                queue = append(queue, neighbor)
            }
        }
    }

    return nil
}

func reconstructPath(parent map[int]int, start, end int) []int {
    path := []int{}
    current := end

    for current != start {
        path = append([]int{current}, path...)  // Prepend
        current = parent[current]
    }

    path = append([]int{start}, path...)
    return path
}

func GetPathLength(path []int) int {
    if len(path) == 0 {
        return 0
    }
    return len(path) - 1  // Number of edges = vertices - 1
}
```

</details>

---

## Think About

1. **When would you use BFS vs DFS for path finding?**
   - BFS: When you need the **shortest** path
   - DFS: When you need **any** path, or **all** paths

2. **Why does backtracking work for finding all paths?**
   - Backtracking "undoes" choices, allowing you to explore alternative routes
   - Without unmarking visited, you'd only find one path

3. **What's the time complexity?**
   - HasPath: O(V + E) - visit each vertex/edge once
   - FindPathBFS: O(V + E) - same as BFS
   - FindAllPathsDFS: O(V!) in worst case - exponential (all possible orderings)

---

## What This Teaches

✅ **BFS/DFS as search tools**, not just traversal
✅ **Path reconstruction** using parent pointers
✅ **Backtracking** for exhaustive search
✅ **BFS finds shortest path** in unweighted graphs
✅ **When to use which algorithm** (shortest vs any vs all)

**Next:** Exercise 03 will dive deeper into BFS applications.
