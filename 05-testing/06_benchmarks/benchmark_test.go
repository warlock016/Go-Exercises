package benchmark

import "testing"

// TODO(human): Write benchmark for ConcatStrings
// func BenchmarkConcatStrings(b *testing.B) {
//     strs := []string{"hello", "world", "from", "Go", "benchmarks"}
//
//     b.ResetTimer()
//     for i := 0; i < b.N; i++ {
//         _ = ConcatStrings(strs)
//     }
// }
//
// Assign result to _ to prevent compiler from optimizing away the call

// TODO(human): Write benchmark for ConcatBuilder
// Same structure as above, but call ConcatBuilder

// TODO(human): Write benchmark for FibonacciRecursive
// Use n = 20 as input
// func BenchmarkFibonacciRecursive(b *testing.B) {
//     const n = 20
//     for i := 0; i < b.N; i++ {
//         _ = FibonacciRecursive(n)
//     }
// }

// TODO(human): Write benchmark for FibonacciIterative
// Use n = 20 as input

// After writing benchmarks, run:
//   go test -bench=.
//   go test -bench=. -benchmem
//   go test -bench=Concat
//   go test -bench=Fibonacci
//
// Observe:
// - How much faster is Builder than + operator?
// - How much faster is iterative than recursive?
// - Look at allocations (B/op and allocs/op)
//
// Try increasing Fibonacci input to 30 - what happens to recursive?
