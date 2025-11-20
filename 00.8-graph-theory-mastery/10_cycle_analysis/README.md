# Exercise 10: Cycle Analysis

**Tier:** 3 (Integration)
**Time:** 75 min
**Goal:** Find and analyze cycles in graphs

## Functions

```go
func FindAllCycles(graph map[int][]int) [][]int
func FindShortestCycle(graph map[int][]int) []int
func HasNegativeCycle(graph WeightedGraph) bool  // For weighted directed graphs
```

## Key Concept

Advanced cycle detection: not just "does cycle exist?" but "what are ALL cycles?"

Use DFS with backtracking to enumerate all cycles.
