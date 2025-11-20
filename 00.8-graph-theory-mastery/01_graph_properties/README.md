# Exercise 01: Graph Properties & Metrics

**Tier:** 1 (Foundation)
**Estimated Time:** 60 minutes
**Learning Goal:** Understand graph characteristics through metric calculation

---

## Learning Objective

Before implementing complex graph algorithms, you need to **understand what makes graphs different from each other**. This exercise teaches you to analyze graphs by calculating fundamental properties like degree, density, and type classification.

**Why this matters:** These metrics help you choose the right algorithm. For example:
- Dense graphs (many edges) → Use adjacency matrix
- Sparse graphs (few edges) → Use adjacency list
- High in-degree nodes → Important in social networks (influencers)

---

## Problem Description

You'll implement functions to calculate various graph properties and metrics. These aren't algorithms per se - they're **analytical tools** that help you understand graph structure.

---

## Function Signatures

```go
// GetDegree returns the degree of a vertex in an undirected graph
func GetDegree(graph map[int][]int, vertex int) int

// GetInDegree returns the in-degree of a vertex in a directed graph
func GetInDegree(graph map[int][]int, vertex int) int

// GetOutDegree returns the out-degree of a vertex in a directed graph
func GetOutDegree(graph map[int][]int, vertex int) int

// GetGraphDensity calculates the density of a graph (edges / possible edges)
// For directed graph: density = E / (V * (V-1))
// For undirected graph: density = 2E / (V * (V-1))
func GetGraphDensity(graph map[int][]int, numVertices int, directed bool) float64

// IsDirected determines if a graph is directed by checking if all edges are bidirectional
func IsDirected(graph map[int][]int) bool

// CountVertices returns the total number of vertices in the graph
func CountVertices(graph map[int][]int) int

// CountEdges returns the total number of edges in the graph
func CountEdges(graph map[int][]int, directed bool) int
```

---

## Examples

### Example 1: Undirected Graph
```
Graph:
    0 --- 1
    |     |
    2 --- 3

Adjacency List:
{
    0: [1, 2],
    1: [0, 3],
    2: [0, 3],
    3: [1, 2],
}

GetDegree(graph, 0) → 2  (connected to 1 and 2)
GetDegree(graph, 1) → 2  (connected to 0 and 3)

CountVertices(graph) → 4
CountEdges(graph, false) → 4  (undirected: count each edge once)

GetGraphDensity(graph, 4, false) → 0.667
  (4 edges, max possible = 6 for 4 vertices)
  (Density = 2*4 / (4*3) = 8/12 = 0.667)

IsDirected(graph) → false  (all edges are bidirectional)
```

### Example 2: Directed Graph
```
Graph:
    0 → 1
    ↓   ↓
    2 → 3

Adjacency List:
{
    0: [1, 2],
    1: [3],
    2: [3],
    3: [],
}

GetOutDegree(graph, 0) → 2  (edges going OUT from 0)
GetInDegree(graph, 0) → 0   (edges coming IN to 0)

GetOutDegree(graph, 3) → 0  (no outgoing edges)
GetInDegree(graph, 3) → 3   (edges from 0,1,2)

CountEdges(graph, true) → 4  (directed: count each edge as-is)

GetGraphDensity(graph, 4, true) → 0.333
  (4 edges, max possible = 12 for 4 vertices)
  (Density = 4 / (4*3) = 4/12 = 0.333)

IsDirected(graph) → true  (not all edges are bidirectional)
```

### Example 3: Sparse vs Dense
```
Sparse Graph (few edges):
    0 --- 1     2 --- 3
    (2 edges, 4 vertices)
    Density = 2*2 / (4*3) = 0.333

Dense Graph (many edges):
    0 - 1
    |\ /|
    | X |
    |/ \|
    2 - 3
    (6 edges, 4 vertices - fully connected!)
    Density = 2*6 / (4*3) = 1.0  (maximum density)
```

---

## Instructions

1. Implement all 7 functions in `graph_properties.go`
2. For `GetInDegree`, you need to count how many other vertices point TO the given vertex
3. For `IsDirected`, check if every edge (u, v) has a corresponding edge (v, u)
4. For `GetGraphDensity`, use the formulas provided in the function signature
5. Handle edge cases: non-existent vertices, empty graphs

---

## Visualization Exercise

**Before coding**, draw these graphs on paper and manually calculate their properties:

**Graph A (Undirected):**
```
    1
   / \
  0   2
   \ /
    3
```

Calculate by hand:
- Degree of each vertex?
- Total edges?
- Density?

**Graph B (Directed):**
```
  0 → 1 → 2
  ↓       ↓
  3 ← 4 ← 5
```

Calculate by hand:
- In-degree and out-degree of vertex 1?
- Which vertex has highest in-degree?
- Is this graph sparse or dense?

---

## Hints

<details>
<summary>Hint 1: Basic Concepts</summary>

**Degree (undirected):** Count the neighbors
- Vertex 0 with neighbors [1,2] → degree 2

**In-degree (directed):** Count how many vertices have an edge TO this vertex
- For vertex 3, count vertices that have 3 in their neighbor list

**Out-degree (directed):** Count the neighbors (same as degree for adjacency list)
- For vertex 0 with neighbors [1,2] → out-degree 2

</details>

<details>
<summary>Hint 2: GetInDegree Algorithm</summary>

```
GetInDegree(graph, target):
    count = 0
    for each vertex in graph:
        for each neighbor of vertex:
            if neighbor == target:
                count++
    return count
```

You need to check ALL vertices to see who points to the target.

</details>

<details>
<summary>Hint 3: IsDirected Logic</summary>

A graph is undirected if EVERY edge is bidirectional:
- If vertex u has neighbor v, then vertex v must have neighbor u
- If you find ANY edge that's not bidirectional, it's directed

```
for each vertex u:
    for each neighbor v of u:
        if v doesn't have u as neighbor:
            return true  // directed!
return false  // all edges bidirectional
```

</details>

<details>
<summary>Hint 4: Counting Edges</summary>

**Directed:** Just count all neighbors across all vertices
```
edges = 0
for each vertex:
    edges += len(neighbors)
```

**Undirected:** Each edge appears twice (u→v and v→u), so divide by 2
```
edges = 0
for each vertex:
    edges += len(neighbors)
return edges / 2
```

</details>

<details>
<summary>Hint 5: Density Formula</summary>

**Maximum possible edges:**
- Directed: V × (V-1) (every vertex can point to every other vertex)
- Undirected: V × (V-1) / 2 (each pair connected once)

**Density:**
- Directed: E / (V × (V-1))
- Undirected: 2E / (V × (V-1)) which simplifies to E / (V×(V-1)/2)

**Edge case:** If V < 2, density is 0 (can't have edges with 0 or 1 vertex)

</details>

<details>
<summary>Complete Solution</summary>

```go
func GetDegree(graph map[int][]int, vertex int) int {
    if neighbors, exists := graph[vertex]; exists {
        return len(neighbors)
    }
    return 0
}

func GetInDegree(graph map[int][]int, vertex int) int {
    count := 0
    for _, neighbors := range graph {
        for _, neighbor := range neighbors {
            if neighbor == vertex {
                count++
            }
        }
    }
    return count
}

func GetOutDegree(graph map[int][]int, vertex int) int {
    return GetDegree(graph, vertex)
}

func GetGraphDensity(graph map[int][]int, numVertices int, directed bool) float64 {
    if numVertices < 2 {
        return 0.0
    }

    edges := float64(CountEdges(graph, directed))
    maxEdges := float64(numVertices * (numVertices - 1))

    if directed {
        return edges / maxEdges
    }
    return (2 * edges) / maxEdges
}

func IsDirected(graph map[int][]int) bool {
    for u, neighbors := range graph {
        for _, v := range neighbors {
            // Check if v has u as neighbor (bidirectional edge)
            vNeighbors, exists := graph[v]
            if !exists {
                return true // v doesn't exist, so edge is directed
            }

            hasBackEdge := false
            for _, neighbor := range vNeighbors {
                if neighbor == u {
                    hasBackEdge = true
                    break
                }
            }

            if !hasBackEdge {
                return true // No back edge, so it's directed
            }
        }
    }
    return false
}

func CountVertices(graph map[int][]int) int {
    return len(graph)
}

func CountEdges(graph map[int][]int, directed bool) int {
    count := 0
    for _, neighbors := range graph {
        count += len(neighbors)
    }

    if !directed {
        count = count / 2 // Each undirected edge counted twice
    }

    return count
}
```

</details>

---

## Think About

1. **Why is density useful?**
   - Helps choose between adjacency list (sparse) vs matrix (dense)
   - Social networks are typically sparse (you're not friends with everyone)
   - Street grids are dense (most intersections connect to neighbors)

2. **Why does in-degree matter in directed graphs?**
   - High in-degree = popular/influential node (many incoming connections)
   - Example: Celebrity in Twitter (millions follow them)
   - Zero in-degree = potential starting points (topological sort)

3. **What's the relationship between degree and edges?**
   - Sum of all degrees = 2 × number of edges (each edge contributes to 2 vertices)
   - This is called the Handshaking Lemma!

---

## What This Teaches

✅ **Analyzing graph structure** before algorithmic design
✅ **Metric calculation** from data structures (map iteration, counting)
✅ **Directed vs undirected** graph characteristics
✅ **Edge case handling** (empty graphs, non-existent vertices)
✅ **Foundation** for choosing appropriate algorithms based on graph properties

**Next:** Exercise 02 will use these concepts to find paths in graphs.
