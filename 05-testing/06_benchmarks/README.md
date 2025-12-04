# Exercise 06: Benchmarks

**Learning Goal:** Learn to write benchmarks with testing.B to measure performance and memory allocations.

---

## Problem Description

Benchmarks measure how fast code runs and how much memory it allocates. Go's testing package makes benchmarking easy with `testing.B`. This is crucial for:
- Comparing different implementations
- Finding performance bottlenecks
- Ensuring optimizations actually help
- Tracking performance regressions

---

## Functions to Benchmark

```go
func ConcatStrings(strs []string) string  // Using + operator
func ConcatBuilder(strs []string) string  // Using strings.Builder
func FibonacciRecursive(n int) int
func FibonacciIterative(n int) int
```

---

## Benchmark Function Signature

```go
func BenchmarkXxx(b *testing.B) {
    // Setup (not timed)

    b.ResetTimer()  // Reset timer after setup

    for i := 0; i < b.N; i++ {
        // Code to benchmark
        // b.N is controlled by the testing framework
    }
}
```

---

## The b.N Loop

The testing framework automatically:
1. Starts with b.N = 1
2. Runs your benchmark
3. If too fast, increases b.N and tries again
4. Continues until it gets stable timing
5. Reports operations per second

**You never set b.N - just loop `for i := 0; i < b.N; i++`**

---

## Running Benchmarks

```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkConcatStrings

# Run with memory allocation stats
go test -bench=. -benchmem

# Run benchmarks multiple times for accuracy
go test -bench=. -benchtime=3s

# Compare two implementations
go test -bench=Concat
```

---

## Example Output

```
BenchmarkConcatStrings-8     1000000   1234 ns/op   512 B/op   10 allocs/op
BenchmarkConcatBuilder-8     5000000    245 ns/op   128 B/op    2 allocs/op
```

Reading this:
- `-8`: GOMAXPROCS (number of CPUs)
- `1000000`: b.N (iterations run)
- `1234 ns/op`: Time per operation
- `512 B/op`: Bytes allocated per operation
- `10 allocs/op`: Number of allocations per operation

---

## Your Task

Write benchmarks for the string concatenation and Fibonacci functions. Compare:
1. String concatenation: + operator vs strings.Builder
2. Fibonacci: recursive vs iterative

Observe the dramatic performance differences!

---

## Example: Benchmarking String Concat

```go
func BenchmarkConcatStrings(b *testing.B) {
    strs := []string{"hello", "world", "from", "Go"}

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = ConcatStrings(strs)
    }
}

func BenchmarkConcatBuilder(b *testing.B) {
    strs := []string{"hello", "world", "from", "Go"}

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = ConcatBuilder(strs)
    }
}
```

---

## Key Benchmark Methods

### b.ResetTimer()
Reset timer after setup code:
```go
func BenchmarkWithSetup(b *testing.B) {
    // Setup (not timed)
    data := generateTestData()

    b.ResetTimer()  // Start timing here

    for i := 0; i < b.N; i++ {
        ProcessData(data)
    }
}
```

### b.ReportAllocs()
Report allocation stats:
```go
func BenchmarkAllocs(b *testing.B) {
    b.ReportAllocs()  // Include allocation stats

    for i := 0; i < b.N; i++ {
        _ = make([]int, 100)
    }
}
```

### b.StopTimer() / b.StartTimer()
Pause/resume timing:
```go
func BenchmarkWithPauses(b *testing.B) {
    for i := 0; i < b.N; i++ {
        b.StopTimer()
        data := setupForThisIteration()
        b.StartTimer()

        Process(data)
    }
}
```

---

## Instructions

1. Open `benchmark_test.go`
2. Write benchmarks for string concatenation (both versions)
3. Write benchmarks for Fibonacci (both versions)
4. Run with `go test -bench=. -benchmem`
5. Compare the results - which implementations are faster?

**Questions to Answer:**
- How much faster is strings.Builder than + operator?
- How much faster is iterative Fibonacci than recursive?
- How many allocations does each method make?

---

## Hints

<details>
<summary>Hint 1: Basic Benchmark Structure</summary>

```go
func BenchmarkFunctionName(b *testing.B) {
    // Optional setup
    input := prepareInput()

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _ = FunctionToTest(input)
        // Assign to _ to prevent compiler optimizing away the call
    }
}
```
</details>

<details>
<summary>Hint 2: Comparing Implementations</summary>

Create similar benchmarks with descriptive names:
- `BenchmarkConcatStrings`
- `BenchmarkConcatBuilder`

Run together: `go test -bench=Concat -benchmem`
</details>

<details>
<summary>Hint 3: Fibonacci Benchmark Input</summary>

For Fibonacci, use a reasonable n value like 20:
- Not too small (results unreliable)
- Not too large (recursive will be very slow)

```go
const fibInput = 20

func BenchmarkFibonacciRecursive(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = FibonacciRecursive(fibInput)
    }
}
```
</details>

<details>
<summary>Full Solution</summary>

```go
package benchmark

import "testing"

func BenchmarkConcatStrings(b *testing.B) {
    strs := []string{"hello", "world", "from", "Go", "benchmarks"}

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = ConcatStrings(strs)
    }
}

func BenchmarkConcatBuilder(b *testing.B) {
    strs := []string{"hello", "world", "from", "Go", "benchmarks"}

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = ConcatBuilder(strs)
    }
}

func BenchmarkFibonacciRecursive(b *testing.B) {
    const n = 20

    for i := 0; i < b.N; i++ {
        _ = FibonacciRecursive(n)
    }
}

func BenchmarkFibonacciIterative(b *testing.B) {
    const n = 20

    for i := 0; i < b.N; i++ {
        _ = FibonacciIterative(n)
    }
}

// You can also benchmark with different input sizes
func BenchmarkFibonacciRecursive10(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = FibonacciRecursive(10)
    }
}

func BenchmarkFibonacciRecursive30(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = FibonacciRecursive(30)
    }
}
```
</details>

---

## What This Teaches

- **Benchmarking basics** - testing.B and benchmark functions
- **The b.N loop** - How the framework controls iterations
- **Performance comparison** - Measuring different implementations
- **Memory profiling** - Understanding allocation costs
- **Benchmark tools** - Running and interpreting benchmark results

---

## Think About

1. Why does the testing framework control b.N instead of letting you set it?
2. What happens if you forget `b.ResetTimer()` after expensive setup?
3. Why assign results to `_` instead of just calling the function?
4. How would you benchmark code that has random variation?
5. What's a "reasonable" benchmark time? 100ns? 1ms? 1s?

---

## Expected Results

You should observe:
- **String concatenation**: Builder is 5-10x faster with far fewer allocations
- **Fibonacci**: Iterative is exponentially faster than recursive

This demonstrates why benchmark-driven optimization matters!

---

**Next Exercise:** `07_test_fixtures` - Using testdata/ and golden files
