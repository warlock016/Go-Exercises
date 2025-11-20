# Algorithmic Thinking Guide

A guide to problem-solving patterns and strategies that transfer across all algorithms.

---

## Table of Contents

1. [What is Algorithmic Thinking?](#what-is-algorithmic-thinking)
2. [Problem-Solving Framework](#problem-solving-framework)
3. [Pattern Recognition](#pattern-recognition)
4. [Data Structure Selection](#data-structure-selection)
5. [Algorithm Design Strategies](#algorithm-design-strategies)
6. [Debugging Algorithms](#debugging-algorithms)
7. [Complexity Reasoning](#complexity-reasoning)

---

## What is Algorithmic Thinking?

**Algorithmic thinking** is the ability to:
1. **Break down** complex problems into smaller steps
2. **Recognize patterns** from problems you've solved before
3. **Choose** the right data structures and approaches
4. **Reason** about correctness and efficiency
5. **Adapt** solutions to new contexts

**It's NOT about:**
- ❌ Memorizing templates
- ❌ Copying solutions without understanding
- ❌ Knowing every algorithm by heart

**It IS about:**
- ✅ Understanding underlying principles
- ✅ Making informed design choices
- ✅ Deriving solutions from first principles

---

## Problem-Solving Framework

### The 5-Step Process

When faced with any algorithmic problem, follow this framework:

#### **Step 1: Understand**

**Questions to ask:**
- What are the inputs? (types, constraints, edge cases)
- What are the outputs? (format, requirements)
- What are the constraints? (time limits, memory limits)
- What does "success" look like?

**Example:**
```
Problem: "Find if there's a path between two nodes"

Inputs:
- Graph (what representation?)
- Start node (integer)
- End node (integer)

Output:
- Boolean (true/false)

Constraints:
- Graph may be disconnected
- Nodes may not exist in graph

Success:
- Returns true if path exists, false otherwise
- Handles edge cases gracefully
```

#### **Step 2: Decompose**

Break the problem into smaller subproblems.

**Technique: Work backwards**
```
Goal: "Find path from A to B"

Work backwards:
- To know if path exists, I need to explore neighbors
- To explore neighbors, I need to track what I've visited
- To track visited, I need a set/map
- To explore systematically, I need BFS or DFS
```

**Technique: Identify subtasks**
```
Problem: "Find shortest path"

Subtasks:
1. Represent the graph
2. Track visited nodes
3. Track distances
4. Explore in order of increasing distance
5. Return path reconstruction
```

#### **Step 3: Pattern Match**

Ask: "Have I seen something like this before?"

**Common patterns:**
- Need to process in order? → **Queue** (BFS)
- Need to backtrack? → **Stack** (DFS)
- Need minimum/maximum? → **Priority Queue** (Heap)
- Need to group items? → **Map** or **Union-Find**
- Need ordering with dependencies? → **Topological Sort**

#### **Step 4: Design**

Choose your approach and data structures.

**Design checklist:**
- [ ] What data structures do I need?
- [ ] What's the high-level algorithm?
- [ ] What are the edge cases?
- [ ] What's the expected complexity?

**Example:**
```
Problem: Find shortest path (unweighted graph)

Design:
- Data structure: Queue (BFS) + visited map + distance map
- Algorithm: BFS, tracking distance at each level
- Edge cases: Start == end, node doesn't exist, disconnected
- Complexity: O(V + E)
```

#### **Step 5: Implement & Test**

Code your solution and verify with examples.

**Implementation tips:**
- Write the simplest version first
- Add edge case handling
- Test with manual traces
- Optimize if needed

---

## Pattern Recognition

### Exploration Patterns

**Problem Signal:** "Visit all nodes," "explore graph," "find path"

| Signal | Pattern | Data Structure |
|--------|---------|----------------|
| "Level by level" | BFS | Queue (FIFO) |
| "Go deep first" | DFS | Stack (LIFO) |
| "Shortest path" | BFS/Dijkstra | Queue/PriorityQueue |
| "All paths" | DFS | Stack + backtracking |

**Example:**
- "Find if nodes are connected" → BFS or DFS (doesn't matter which)
- "Find shortest distance" → BFS (if unweighted)
- "Find all possible routes" → DFS (explores all branches)

---

### Ordering Patterns

**Problem Signal:** "Sort with constraints," "dependencies," "schedule tasks"

| Signal | Pattern | Algorithm |
|--------|---------|-----------|
| "X must come before Y" | Topological order | Topological Sort (DFS) |
| "Process in order of cost" | Greedy | Priority Queue |
| "Respect precedence" | Dependency graph | DFS on DAG |

**Example:**
- "Build tasks in order" → Topological sort
- "Process cheapest first" → Priority queue (greedy)

---

### Grouping Patterns

**Problem Signal:** "Find groups," "connected components," "clusters"

| Signal | Pattern | Algorithm |
|--------|---------|-----------|
| "Separate into groups" | Connected components | BFS/DFS from each unvisited |
| "Merge sets" | Disjoint sets | Union-Find |
| "Two categories" | Bipartite check | BFS with 2-coloring |

---

### Optimization Patterns

**Problem Signal:** "Minimum cost," "shortest," "cheapest," "optimal"

| Signal | Pattern | Algorithm |
|--------|---------|-----------|
| "Min cost to connect all" | MST | Kruskal/Prim |
| "Shortest path" | Shortest path | Dijkstra/BFS |
| "Optimal sequence" | Dynamic Programming | DP on graph |

---

## Data Structure Selection

### The Decision Tree

**Need to process in order?**
```
First-in-first-out (FIFO) → Queue (BFS)
Last-in-first-out (LIFO)  → Stack (DFS)
Smallest-first            → Priority Queue (Dijkstra)
```

**Need to track membership?**
```
Fast lookup (visited?)    → Map/Set
Grouping elements         → Union-Find
Unique values             → Set
Count occurrences         → Map
```

**Need to track relationships?**
```
Connections between items → Graph (adjacency list)
Hierarchy                 → Tree
Linear sequence           → Slice/Array
```

---

### Data Structure Trade-offs

| Structure | Access | Insert | Delete | Search | Use Case |
|-----------|--------|--------|--------|--------|----------|
| Slice | O(1) | O(1)* | O(n) | O(n) | Sequential, ordered |
| Map | O(1) | O(1) | O(1) | O(1) | Key-value, membership |
| Set | - | O(1) | O(1) | O(1) | Unique items, visited tracking |
| Priority Queue | O(1) peek | O(log n) | O(log n) | - | Always need min/max |

*Amortized for append

---

## Algorithm Design Strategies

### 1. **Brute Force First**

**Strategy:** Solve the problem the simplest way, even if inefficient.

**Why?**
- Gets you thinking about correctness
- Establishes a baseline
- Often reveals patterns for optimization

**Example:**
```
Problem: Find if path exists

Brute force:
- Try all possible paths
- Very inefficient (exponential!)
- But: Leads you to realize you need to avoid revisiting (→ visited set)
```

---

### 2. **Greedy**

**Strategy:** Make the locally optimal choice at each step.

**When to use:**
- "Always pick the best available option"
- Examples: Dijkstra (pick shortest distance), Kruskal (pick cheapest edge)

**Caution:** Greedy doesn't always work! Only for specific problems.

---

### 3. **Divide and Conquer**

**Strategy:** Break problem into smaller subproblems, solve recursively.

**Pattern:**
```
1. Divide problem into smaller pieces
2. Solve each piece recursively
3. Combine results
```

**Examples:** Merge sort, quicksort, binary search

---

### 4. **Dynamic Programming**

**Strategy:** Store results of subproblems to avoid recomputation.

**When to use:**
- Overlapping subproblems (same calculation repeated)
- Optimal substructure (optimal solution uses optimal subsolutions)

**Examples:** Fibonacci, shortest paths (Floyd-Warshall), longest path

---

### 5. **Backtracking**

**Strategy:** Try options, undo if they don't work.

**Pattern:**
```
1. Make a choice
2. Recurse with that choice
3. If it doesn't work, undo and try next choice
```

**Examples:** Finding all paths, N-queens, maze solving

---

## Debugging Algorithms

### Debugging Checklist

When your algorithm doesn't work:

1. **Trace by hand**
   - Run through your algorithm manually with a small example
   - Compare expected vs actual at each step
   - Find where they diverge

2. **Check edge cases**
   - Empty input
   - Single element
   - All same values
   - Duplicates
   - Negative values
   - Maximum/minimum values

3. **Verify data structure usage**
   - Are you popping from the right end? (queue vs stack)
   - Are you checking visited before or after adding to collection?
   - Are you updating the right variables?

4. **Add logging**
   - Print state at each iteration
   - Print data structure contents
   - Verify assumptions

5. **Simplify**
   - Remove optimizations
   - Test with trivial input
   - Isolate the failing component

---

### Common Algorithm Bugs

#### Bug: Infinite loop
**Cause:** Not marking nodes as visited, or checking visited after processing
**Fix:** Mark visited **before** adding to queue/stack

#### Bug: Wrong traversal order
**Cause:** Using stack instead of queue (or vice versa)
**Fix:** BFS = queue (front), DFS = stack (back)

#### Bug: Missing nodes
**Cause:** Only starting from one node in disconnected graph
**Fix:** Loop over all nodes, start BFS/DFS from unvisited ones

#### Bug: Incorrect path reconstruction
**Cause:** Not tracking parent pointers during traversal
**Fix:** Use `parent[child] = current` during exploration

---

## Complexity Reasoning

### How to Analyze Complexity

#### **Time Complexity**

Ask: "How many times does each operation run?"

**Example: BFS**
```go
for len(queue) > 0 {           // Each node visited once: V iterations
    current := queue[0]
    queue = queue[1:]

    visited[current] = true

    for _, neighbor := range graph[current] {  // Each edge checked once: E iterations
        if !visited[neighbor] {
            queue = append(queue, neighbor)
        }
    }
}
```
**Analysis:** Outer loop runs V times, inner loop runs E times total → **O(V + E)**

---

#### **Space Complexity**

Ask: "How much memory grows with input size?"

**Example: BFS**
```go
visited := make(map[int]bool)    // O(V) - one entry per vertex
queue := []int{start}            // O(V) - worst case all vertices in queue
result := []int{}                // O(V) - one result per vertex
```
**Analysis:** All scale with V → **O(V)**

---

### Big-O Quick Reference

| Complexity | Name | Example |
|------------|------|---------|
| O(1) | Constant | Map lookup, array access |
| O(log n) | Logarithmic | Binary search, balanced tree |
| O(n) | Linear | Iterate through array |
| O(n log n) | Linearithmic | Merge sort, heap sort |
| O(n²) | Quadratic | Nested loops over n |
| O(2ⁿ) | Exponential | Recursive Fibonacci (naive) |

**Goal:** Aim for O(n) or O(n log n) when possible.

---

## Mental Models

### The Algorithm Toolbox

Think of algorithms as tools in a toolbox:

```
Problem → Choose Tool → Apply

"Need shortest path?"
  ↓
  Unweighted? → BFS
  Weighted? → Dijkstra
  All pairs? → Floyd-Warshall

"Need to visit all nodes?"
  ↓
  Level-order? → BFS (queue)
  Depth-first? → DFS (stack)

"Need to group items?"
  ↓
  Connected components? → BFS/DFS
  Merge sets? → Union-Find
```

**The more problems you solve, the bigger your toolbox!**

---

### The "Why" Before the "How"

**Always ask:**
1. **Why does this work?** (Correctness)
2. **Why is this efficient?** (Complexity)
3. **Why this approach over others?** (Trade-offs)

**Example: Why does BFS find shortest path?**
- BFS explores level by level
- Level N contains all nodes at distance N
- First time we reach target, we've used minimum edges
- Therefore, BFS finds shortest path in unweighted graphs

**Understanding "why" lets you:**
- Derive algorithms instead of memorizing
- Adapt to new problems
- Explain your reasoning

---

## Final Principles

### 1. **Simplicity First**

Start with the simplest solution. Optimize only if needed.

**Example:**
- Brute force first → Measure → Optimize bottlenecks
- Don't prematurely optimize!

### 2. **Correctness Over Cleverness**

A simple, correct solution beats a clever, buggy one.

**Example:**
- Readable BFS with comments > Obfuscated "optimized" version

### 3. **Patterns Over Memorization**

Learn the patterns, not individual solutions.

**Example:**
- "Queue = process in order" applies to BFS, level-order, task scheduling, etc.

### 4. **Practice Deliberately**

Focus on one concept at a time until automatic.

**Example:**
- Master BFS/DFS before moving to Dijkstra
- Depth > Breadth

### 5. **Visualize Before Coding**

Draw the problem. Trace by hand. Then code.

**Example:**
```
Problem → Draw graph → Trace algorithm → Pseudocode → Code → Test
```

---

## Summary

**Algorithmic thinking is:**
1. Breaking problems into patterns
2. Choosing the right tools (data structures + algorithms)
3. Reasoning about correctness and efficiency
4. Building intuition through practice

**You're building:**
- 🧠 Mental models (queue = FIFO, stack = LIFO)
- 🔧 Toolbox (BFS, DFS, Dijkstra, etc.)
- 🎯 Pattern recognition (shortest path → BFS, dependencies → topo sort)
- 🤔 Reasoning skills (why does this work? what's the complexity?)

**Remember:** Every expert was once a beginner. The difference is deliberate practice and understanding "why," not just "how."

Keep building your algorithmic intuition, one problem at a time! 🚀
