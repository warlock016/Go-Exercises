# Exercise 14: Algorithm Selection & Analysis

**Tier:** 4 (Mastery)
**Time:** 75 min
**Goal:** Choose the right algorithm for the problem

## Functions

```go
func ChooseShortestPathAlgorithm(graphType string, weighted bool, negative bool, allPairs bool) string
func ComparePerformance(graph interface{}, algorithms []string) map[string]time.Duration
func AnalyzeGraphProperties(graph map[int][]int) map[string]interface{}  // density, diameter, etc.
```

## Key Concept

Meta-algorithmic reasoning: **when** to use **which** algorithm and **why**.

**Decision Tree:**
```
Need shortest path?
├─ Unweighted? → BFS (O(V+E))
├─ Weighted, non-negative?
│  ├─ Single source? → Dijkstra (O((V+E)logV))
│  └─ All pairs? → Floyd-Warshall (O(V³))
└─ Negative weights? → Bellman-Ford (O(VE))
```

Benchmark different approaches on the same graph.
