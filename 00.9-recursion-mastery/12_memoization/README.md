# Exercise 12: Memoization

**Tier:** 4 - Advanced
**Estimated Time:** 60 minutes
**Concepts:** Dynamic programming, caching, optimization, memoization pattern

## Learning Goal

Transform slow recursive algorithms into fast ones using memoization - caching results to avoid recomputation.

## Problems

1. **FibonacciMemo** - Fibonacci with O(n) instead of O(2^n)
2. **GridPaths** - Count paths in grid with memoization
3. **ClimbStairsMemo** - Optimized stair climbing (1, 2, or 3 steps)

## The Memoization Pattern

```go
func solveWithMemo(n int) int {
    memo := make(map[int]int)
    return helper(n, memo)
}

func helper(n int, memo map[int]int) int {
    if result, exists := memo[n]; exists {
        return result  // Return cached result
    }
    
    // Compute result
    result := ...recursive calls...
    
    memo[n] = result  // Cache it
    return result
}
```

## Performance Impact

- **Without memo:** Fibonacci(40) = seconds
- **With memo:** Fibonacci(40) = microseconds

**Next Exercise:** 13 - N-Queens
