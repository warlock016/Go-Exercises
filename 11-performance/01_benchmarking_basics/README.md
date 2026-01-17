# Exercise 01: Benchmarking Basics

## Learning Goal
Learn to write, run, and interpret Go benchmarks using the `testing` package.

## Problem Description

Benchmarking is essential for performance optimization. Go's `testing` package includes built-in support for benchmarks that measure execution time and memory allocations.

Key concepts:
- Benchmark functions start with `Benchmark` and take `*testing.B`
- `b.N` is the number of iterations (adjusted automatically)
- `b.ResetTimer()` excludes setup time from measurement
- `b.ReportAllocs()` includes allocation statistics
- Sub-benchmarks allow comparing variations

## Function Signatures

```go
// SumLoop adds numbers from 0 to n using a loop
func SumLoop(n int) int

// SumFormula adds numbers from 0 to n using the formula n*(n+1)/2
func SumFormula(n int) int

// FibonacciRecursive calculates the nth Fibonacci number recursively
func FibonacciRecursive(n int) int

// FibonacciIterative calculates the nth Fibonacci number iteratively
func FibonacciIterative(n int) int

// ReverseStringNaive reverses a string using concatenation
func ReverseStringNaive(s string) string

// ReverseStringBuilder reverses a string using strings.Builder
func ReverseStringBuilder(s string) string
```

## Benchmark Examples

```go
func BenchmarkSumLoop(b *testing.B) {
    for i := 0; i < b.N; i++ {
        SumLoop(1000)
    }
}

func BenchmarkSumFormula(b *testing.B) {
    for i := 0; i < b.N; i++ {
        SumFormula(1000)
    }
}

// Sub-benchmarks for different input sizes
func BenchmarkSum(b *testing.B) {
    sizes := []int{10, 100, 1000, 10000}
    for _, size := range sizes {
        b.Run(fmt.Sprintf("Loop/%d", size), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                SumLoop(size)
            }
        })
        b.Run(fmt.Sprintf("Formula/%d", size), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                SumFormula(size)
            }
        })
    }
}
```

## Running Benchmarks

```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkSum

# Include memory stats
go test -bench=. -benchmem

# Run multiple times for stable results
go test -bench=. -count=5

# Set minimum benchmark time
go test -bench=. -benchtime=3s
```

## Instructions

1. Implement the functions in `benchmarking_basics.go`
2. Write benchmarks in `benchmarking_basics_test.go`
3. Run benchmarks and observe the differences
4. Experiment with different input sizes using sub-benchmarks

## Hints

### Basic
- The sum formula is: n * (n + 1) / 2
- Fibonacci: F(n) = F(n-1) + F(n-2), with F(0)=0, F(1)=1
- strings.Builder is more efficient for building strings

### Intermediate
- Use `b.ResetTimer()` after expensive setup
- Recursive Fibonacci has exponential complexity
- String concatenation creates new strings each time

### Benchmark Interpretation
```
BenchmarkSumLoop-8     100000    12345 ns/op    0 B/op    0 allocs/op
                 │          │         │          │           │
                 │          │         │          │           └── allocations per op
                 │          │         │          └── bytes allocated per op
                 │          │         └── nanoseconds per operation
                 │          └── number of iterations
                 └── GOMAXPROCS
```

## Think About

1. Why does SumFormula benchmark faster regardless of input size?
2. Why is FibonacciRecursive so much slower for large n?
3. Why does ReverseStringNaive allocate more memory?
4. How does b.N adapt to different operation speeds?

## What This Teaches

- Writing benchmark functions
- Using b.N for iteration counts
- Sub-benchmarks for comparing variations
- Reading benchmark output
- Understanding algorithmic complexity through measurement
