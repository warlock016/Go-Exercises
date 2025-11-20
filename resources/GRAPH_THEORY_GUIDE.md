# Graph Theory Guide

A comprehensive visual guide to graph concepts, representations, and algorithms.

---

## Table of Contents

1. [What is a Graph?](#what-is-a-graph)
2. [Graph Terminology](#graph-terminology)
3. [Graph Representations](#graph-representations)
4. [Traversal Algorithms](#traversal-algorithms)
5. [Common Graph Algorithms](#common-graph-algorithms)
6. [Complexity Analysis](#complexity-analysis)
7. [When to Use Which Algorithm](#when-to-use-which-algorithm)

---

## What is a Graph?

A **graph** is a data structure that represents relationships between objects.

**Components:**
- **Vertices (nodes):** The objects themselves (people, cities, web pages)
- **Edges:** The connections between objects (friendships, roads, hyperlinks)

**Visual Example:**
```
    A -------- B
    |          |
    |          |
    C -------- D
```

**Real-World Examples:**
- **Social Network:** Vertices = people, Edges = friendships
- **Road Map:** Vertices = cities, Edges = roads
- **Web:** Vertices = pages, Edges = hyperlinks
- **Dependencies:** Vertices = tasks, Edges = "must do before"

---

## Graph Terminology

### 1. **Directed vs Undirected**

**Undirected Graph:** Edges have no direction (symmetric relationship)
```
A ←→ B    (A and B are friends)
```

**Directed Graph (Digraph):** Edges have direction (asymmetric relationship)
```
A → B     (A follows B, but B doesn't follow A)
```

### 2. **Weighted vs Unweighted**

**Unweighted:** All edges are equal
```
A ---- B    (distance doesn't matter)
```

**Weighted:** Edges have costs/distances/weights
```
A --5-- B   (5 km between A and B)
A -10-- C   (10 km between A and C)
```

### 3. **Degree**

**Degree** = Number of edges connected to a vertex

**Undirected:**
```
    A (degree 2)
   / \
  B   C (both degree 1)
```

**Directed:**
- **In-degree:** Edges coming IN
- **Out-degree:** Edges going OUT
```
A → B → C
A: in=0, out=1
B: in=1, out=1
C: in=1, out=0
```

### 4. **Path**

A **path** is a sequence of vertices connected by edges.

```
A → B → C → D
Path from A to D: [A, B, C, D]
Path length: 3 (number of edges)
```

**Simple Path:** No repeated vertices
**Cycle:** A path that starts and ends at the same vertex

### 5. **Connected Components**

A **connected component** is a subgraph where all vertices can reach each other.

```
Component 1:    Component 2:
  A --- B          E --- F
  |     |
  C --- D

This graph has 2 connected components.
```

### 6. **DAG (Directed Acyclic Graph)**

A directed graph with **no cycles**.

```
   A
  / \
 B   C   ← DAG (no way to get back to A)
 |
 D
```

**Use cases:** Task scheduling, dependency resolution, build systems

---

## Graph Representations

### 1. **Adjacency List**

**Structure:** Map of vertex → list of neighbors

**Example:**
```
Graph:        Adjacency List:
  0 → 1       0: [1, 2]
  0 → 2       1: [2]
  1 → 2       2: []
  2: (none)
```

**In Go:**
```go
graph := map[int][]int{
    0: {1, 2},
    1: {2},
    2: {},
}
```

**Pros:**
- ✅ Space efficient for sparse graphs: O(V + E)
- ✅ Fast neighbor lookup: O(1) average
- ✅ Easy to add edges

**Cons:**
- ❌ Checking if edge exists: O(degree(v))
- ❌ Not cache-friendly

**When to use:** Most real-world graphs (social networks, web graphs) are sparse.

---

### 2. **Adjacency Matrix**

**Structure:** 2D array, `matrix[i][j] = true` if edge from i to j exists

**Example:**
```
Graph:        Adjacency Matrix:
  0 → 1          0  1  2
  0 → 2       0 [F  T  T]
  1 → 2       1 [F  F  T]
  2: (none)   2 [F  F  F]
```

**In Go:**
```go
graph := [][]bool{
    {false, true, true},   // 0 → 1, 0 → 2
    {false, false, true},  // 1 → 2
    {false, false, false}, // 2 → nothing
}
```

**Pros:**
- ✅ Fast edge checking: O(1)
- ✅ Simple to implement
- ✅ Cache-friendly for dense graphs

**Cons:**
- ❌ Space inefficient: O(V²) even for sparse graphs
- ❌ Adding vertices requires reallocation

**When to use:** Dense graphs, or when you need O(1) edge checks frequently.

---

### 3. **Edge List**

**Structure:** List of (source, dest, weight) tuples

**Example:**
```
Graph:        Edge List:
  0 → 1       [(0,1), (0,2), (1,2)]
  0 → 2
  1 → 2
```

**In Go:**
```go
type Edge struct {
    from, to, weight int
}
edges := []Edge{
    {0, 1, 5},
    {0, 2, 3},
    {1, 2, 2},
}
```

**When to use:** Kruskal's MST algorithm, reading from files

---

## Traversal Algorithms

### BFS (Breadth-First Search)

**Idea:** Explore **level by level** (like ripples in water)

**Data Structure:** **Queue** (FIFO - First In, First Out)

**Visual:**
```
Start at 0:

Level 0:     0
            / \
Level 1:   1   2
          /     \
Level 2: 3       4

BFS order: [0, 1, 2, 3, 4] (level by level)
```

**Algorithm:**
```
1. Add start to queue
2. While queue not empty:
   a. Remove from front (current)
   b. Mark as visited
   c. Add all unvisited neighbors to back of queue
```

**Step-by-Step:**
```
Queue: [0]           Visit: []
Queue: [1, 2]        Visit: [0]
Queue: [2, 3]        Visit: [0, 1]
Queue: [3, 4]        Visit: [0, 1, 2]
Queue: [4]           Visit: [0, 1, 2, 3]
Queue: []            Visit: [0, 1, 2, 3, 4] ← Done!
```

**When to use:**
- ✅ Shortest path in **unweighted** graphs
- ✅ Level-order traversal
- ✅ Finding minimum steps/moves
- ✅ Testing if graph is bipartite

**Time:** O(V + E)
**Space:** O(V) for queue

---

### DFS (Depth-First Search)

**Idea:** Explore **as deep as possible** before backtracking

**Data Structure:** **Stack** (LIFO - Last In, First Out)

**Visual:**
```
Start at 0:

    0
   / \
  1   2
 /     \
3       4

DFS order: [0, 2, 4, 1, 3] (go deep first!)
```

**Algorithm:**
```
1. Add start to stack
2. While stack not empty:
   a. Remove from top (current)
   b. Mark as visited
   c. Add all unvisited neighbors to top of stack
```

**Step-by-Step:**
```
Stack: [0]           Visit: []
Stack: [1, 2]        Visit: [0]       ← Added neighbors of 0
Stack: [1, 4]        Visit: [0, 2]    ← Popped 2, added neighbors
Stack: [1]           Visit: [0, 2, 4] ← Popped 4, no neighbors
Stack: [3]           Visit: [0, 2, 4, 1] ← Popped 1, added neighbors
Stack: []            Visit: [0, 2, 4, 1, 3] ← Done!
```

**When to use:**
- ✅ Cycle detection
- ✅ Finding **all** paths
- ✅ Topological sort
- ✅ Maze solving
- ✅ Backtracking problems

**Time:** O(V + E)
**Space:** O(V) for stack

---

### BFS vs DFS: The Key Difference

**Same graph, different orders:**
```
      0
     / \
    1   2
   /|   |\
  3 4   5 6
```

**BFS (Queue):** [0, 1, 2, 3, 4, 5, 6]
**DFS (Stack):** [0, 2, 6, 5, 1, 4, 3]

**Why?**
- **BFS:** Process nodes in order discovered (queue → FIFO)
- **DFS:** Process most recently discovered nodes first (stack → LIFO)

**Mental Model:**
- **BFS:** Expanding ripple, exploring outward
- **DFS:** Following one path to the end, then backtracking

---

## Common Graph Algorithms

### 1. **Shortest Path (Unweighted) - BFS**

**Problem:** Find shortest path from A to B

**Solution:** BFS automatically finds shortest path in unweighted graphs!

**Why?** BFS explores level by level, so the first time you reach the target, you've found the shortest path.

**Example:**
```
A → B → D
↓   ↓
C → E

Shortest path A to E:
BFS visits: A → [B,C] → [D,E]
Path: A → C → E (2 edges)
```

---

### 2. **Shortest Path (Weighted) - Dijkstra**

**Problem:** Find shortest path in **weighted** graph

**Data Structure:** **Priority Queue** (always process lowest-cost node next)

**Example:**
```
    A --1-- B
    |       |
    5       2
    |       |
    C --1-- D

Dijkstra from A to D:
1. Distance[A] = 0
2. Visit A, update: Distance[B]=1, Distance[C]=5
3. Visit B (lowest=1), update: Distance[D]=3
4. Visit D (lowest=3) ← Found shortest path!

Shortest: A → B → D (cost 3)
```

**Time:** O((V + E) log V) with binary heap

---

### 3. **Cycle Detection**

**Undirected:** Use DFS, detect if you visit an already-visited node (that isn't your parent)

**Directed:** Use DFS with three colors:
- White: Unvisited
- Gray: Visiting (in current path)
- Black: Visited (completely explored)

**Cycle exists if:** You encounter a **gray** node (node in current DFS path)

---

### 4. **Topological Sort**

**Problem:** Order tasks respecting dependencies

**Requirements:** Graph must be a **DAG** (no cycles!)

**Algorithm (DFS-based):**
1. Do DFS from all unvisited nodes
2. When finishing a node (all descendants explored), add to front of result

**Example:**
```
   put_on_socks → put_on_shoes
   put_on_pants → put_on_shoes
   put_on_pants → put_on_belt

Topological order:
[put_on_pants, put_on_belt, put_on_socks, put_on_shoes]
```

---

### 5. **Connected Components**

**Problem:** Find all disconnected subgraphs

**Algorithm:**
```
1. For each unvisited vertex:
   a. Do BFS/DFS to find all reachable vertices
   b. Mark them as one component
   c. Increment component count
```

**Example:**
```
A--B    C--D    E

Components: [[A,B], [C,D], [E]]
Count: 3
```

---

## Complexity Analysis

### Time Complexity

**Common notation:**
- V = number of vertices
- E = number of edges

| Algorithm | Adjacency List | Adjacency Matrix |
|-----------|---------------|------------------|
| BFS/DFS | O(V + E) | O(V²) |
| Dijkstra | O((V+E) log V) | O(V²) |
| Floyd-Warshall | O(V³) | O(V³) |

**Why BFS/DFS is O(V + E)?**
- Visit each vertex once: O(V)
- Check each edge once: O(E)
- Total: O(V + E)

### Space Complexity

| Representation | Space | Notes |
|----------------|-------|-------|
| Adjacency List | O(V + E) | Efficient for sparse graphs |
| Adjacency Matrix | O(V²) | Wastes space if sparse |
| Edge List | O(E) | Minimal space |

**Sparse vs Dense:**
- **Sparse:** E ≈ V (few edges) → Use adjacency list
- **Dense:** E ≈ V² (many edges) → Matrix might be better

---

## When to Use Which Algorithm

### Decision Tree

**Need shortest path?**
- Unweighted → **BFS**
- Weighted (non-negative) → **Dijkstra**
- Weighted (negative edges) → **Bellman-Ford**
- All pairs → **Floyd-Warshall**

**Need to visit all nodes?**
- Level by level → **BFS**
- Depth-first → **DFS**

**Dependency ordering?**
- → **Topological Sort** (requires DAG)

**Find cycles?**
- → **DFS with colors**

**Find components?**
- → **BFS or DFS** from each unvisited node

**Minimum cost to connect all?**
- → **MST** (Kruskal or Prim)

---

## Quick Reference

### BFS Template
```go
func BFS(graph map[int][]int, start int) []int {
    visited := make(map[int]bool)
    queue := []int{start}
    result := []int{}

    for len(queue) > 0 {
        current := queue[0]      // Front
        queue = queue[1:]

        if visited[current] {
            continue
        }

        visited[current] = true
        result = append(result, current)

        for _, neighbor := range graph[current] {
            if !visited[neighbor] {
                queue = append(queue, neighbor) // Back
            }
        }
    }

    return result
}
```

### DFS Template
```go
func DFS(graph map[int][]int, start int) []int {
    visited := make(map[int]bool)
    stack := []int{start}
    result := []int{}

    for len(stack) > 0 {
        current := stack[len(stack)-1] // Top
        stack = stack[:len(stack)-1]

        if visited[current] {
            continue
        }

        visited[current] = true
        result = append(result, current)

        for _, neighbor := range graph[current] {
            if !visited[neighbor] {
                stack = append(stack, neighbor) // Top
            }
        }
    }

    return result
}
```

**Notice:** The ONLY difference is which end you pop from!

---

## Final Tips

1. **Draw the graph** before coding - visualization prevents bugs
2. **Trace by hand** - understand the algorithm before implementing
3. **Know your data structure** - Queue (BFS), Stack (DFS), PriorityQueue (Dijkstra)
4. **Check for edge cases** - Empty graph, single node, disconnected components
5. **Understand the "why"** - Don't just memorize templates

**Remember:** Graph algorithms are just patterns of exploration. Once you understand BFS/DFS, the rest are variations!
