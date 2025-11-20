# Exercise 03: BFS Deep Dive

**Tier:** 1 (Foundation)
**Estimated Time:** 75 minutes
**Learning Goal:** Master BFS applications - shortest distance, level-order, connected components

---

## Learning Objective

BFS isn't just about visiting nodes - it's a powerful tool for solving specific types of problems. This exercise teaches you to **think in BFS patterns**: levels, shortest paths, and breadth-first exploration.

**Key Insight:** BFS naturally solves problems involving "minimum steps," "closest node," or "level-by-level processing" because it explores in expanding circles from the source.

---

## Function Signatures

```go
// ShortestDistance returns the shortest distance (number of edges) from start to end
// Returns -1 if no path exists
func ShortestDistance(graph map[int][]int, start, end int) int

// LevelOrder returns vertices grouped by their distance from start
// Example: [[0], [1,2], [3,4,5]] means 0 is level 0, 1&2 are level 1, etc.
func LevelOrder(graph map[int][]int, start int) [][]int

// FindClosestVertex finds the closest vertex to start that satisfies a condition
// condition is a function that returns true for target vertices
// Returns the vertex and distance, or (-1, -1) if none found
func FindClosestVertex(graph map[int][]int, start int, condition func(int) bool) (vertex int, distance int)

// CountComponentsBFS counts the number of connected components in the graph
func CountComponentsBFS(graph map[int][]int) int

// GetComponentBFS returns all vertices in the same connected component as start
func GetComponentBFS(graph map[int][]int, start int) []int
```

---

## Examples

### Example 1: Shortest Distance
```
Graph:
    0 --- 1 --- 2
    |           |
    3 --- 4 --- 5

ShortestDistance(graph, 0, 5) → 3 (path: 0->1->2->5 or 0->3->4->5)
ShortestDistance(graph, 0, 0) → 0
ShortestDistance(graph, 0, 99) → -1 (doesn't exist)
```

### Example 2: Level Order
```
Graph (tree-like):
        0
       /|\
      1 2 3
      |   |
      4   5

LevelOrder(graph, 0) → [[0], [1,2,3], [4,5]]
- Level 0: {0}
- Level 1: {1,2,3} (distance 1 from 0)
- Level 2: {4,5} (distance 2 from 0)
```

### Example 3: Find Closest Vertex
```
Graph:
    0 --- 1 --- 2 --- 3
                |
                4

// Find closest even number to 0
condition := func(v int) bool { return v % 2 == 0 }
FindClosestVertex(graph, 0, condition) → (2, 2)
// Vertex 2 is at distance 2, and it's even
```

### Example 4: Connected Components
```
Graph:
    0 --- 1     2 --- 3     4

CountComponentsBFS(graph) → 3
GetComponentBFS(graph, 0) → [0, 1]
GetComponentBFS(graph, 2) → [2, 3]
GetComponentBFS(graph, 4) → [4]
```

---

## Visualization Exercise

**Trace BFS by hand:**

Graph:
```
    0
   / \
  1   2
  |\ /|
  3 4 5
```

Starting from 0, fill in the table:

| Iteration | Queue | Current | Distance from 0 | Level | Action |
|-----------|-------|---------|-----------------|-------|--------|
| 0 | [0] | - | - | - | Initialize |
| 1 | [1,2] | 0 | {0:0} | 0 | Visit 0 |
| 2 | [2,3,4] | 1 | {0:0, 1:1} | 1 | Visit 1 |
| 3 | [3,4,4,5] | 2 | {0:0, 1:1, 2:1} | 1 | Visit 2 |
| ... | ... | ... | ... | ... | Continue |

**Questions:**
1. What's the shortest distance from 0 to 5?
2. What vertices are at level 2?
3. How do you avoid adding 4 twice to the queue?

---

## Hints

<details>
<summary>Hint 1: Tracking Distance</summary>

BFS naturally tracks distance! Each "level" of BFS represents vertices at distance d from start.

```go
distance := make(map[int]int)
distance[start] = 0

for len(queue) > 0 {
    current := queue[0]
    queue = queue[1:]

    for _, neighbor := range graph[current] {
        if _, visited := distance[neighbor]; !visited {
            distance[neighbor] = distance[current] + 1  // One step further
            queue = append(queue, neighbor)
        }
    }
}
```

</details>

<details>
<summary>Hint 2: Level Order Grouping</summary>

Group vertices by their distance from start:

```go
levels := [][]int{}
levels[0] = []int{start}

currentLevel := 0
for there are nodes at currentLevel {
    nextLevel := []int{}
    for each node in levels[currentLevel] {
        for each unvisited neighbor {
            add neighbor to nextLevel
        }
    }
    levels = append(levels, nextLevel)
    currentLevel++
}
```

</details>

<details>
<summary>Hint 3: Connected Components</summary>

Run BFS from each unvisited vertex:

```go
count := 0
visited := make(map[int]bool)

for vertex := range graph {
    if !visited[vertex] {
        BFS(graph, vertex, visited)  // Marks entire component as visited
        count++
    }
}
return count
```

</details>

---

## What This Teaches

✅ **BFS finds shortest path** in unweighted graphs automatically
✅ **Level-order traversal** groups nodes by distance
✅ **Queue = FIFO** ensures breadth-first exploration
✅ **Connected components** via repeated BFS
✅ **Distance tracking** is natural with BFS

**Next:** Exercise 04 will explore DFS applications and contrast with BFS.
