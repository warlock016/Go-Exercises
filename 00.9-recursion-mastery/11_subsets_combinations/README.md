# Exercise 11: Subsets & Combinations

**Tier:** 4 - Advanced
**Estimated Time:** 65 minutes
**Concepts:** Backtracking, choose-explore-unchoose, power set, combinatorics

## Learning Goal

Master backtracking with the choose-explore-unchoose pattern. Generate all possible combinations of elements.

## Problems

1. **GenerateSubsets** - All subsets (power set) of [1,2,3]
2. **Combinations** - All k-length combinations
3. **Permutations** - All orderings of elements

## Examples

```go
GenerateSubsets([]int{1, 2}) 
// [[], [1], [2], [1,2]]

Combinations([]int{1, 2, 3}, 2)
// [[1,2], [1,3], [2,3]]

Permutations([]int{1, 2, 3})
// [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
```

## The Backtracking Pattern

```go
func backtrack(choices, current path) {
    if isComplete(current) {
        results.add(copy(current))
        return
    }
    
    for each choice {
        current.add(choice)      // CHOOSE
        backtrack(...)           // EXPLORE
        current.remove(choice)   // UNCHOOSE
    }
}
```

**Next Exercise:** 12 - Memoization
