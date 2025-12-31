# Exercise 10: Pipeline Pattern

**Learning Goal:** Build composable pipeline stages that process data streams

**Difficulty:** Tier 3 - Integration
**Estimated Time:** 50-55 minutes

---

## Problem Description

A **pipeline** is a series of stages connected by channels, where each stage:
1. Receives values from upstream via an inbound channel
2. Processes those values (transform, filter, aggregate)
3. Sends results downstream via an outbound channel

Pipelines enable:
- **Separation of concerns** - Each stage does one thing
- **Composability** - Stages can be rearranged or reused
- **Bounded parallelism** - Easy to add workers per stage
- **Streaming** - Process data as it arrives, not in batches

Think: assembly line in a factory.

---

## Function Signatures

```go
// Stage is a function that processes input channel and returns output channel
type Stage[I, O any] func(input <-chan I) <-chan O

// Generator creates a channel that emits values from a slice
func Generator[T any](values ...T) <-chan T

// Filter keeps only values that pass the predicate
func Filter[T any](predicate func(T) bool) Stage[T, T]

// Map transforms each value
func Map[I, O any](transform func(I) O) Stage[I, O]

// Take emits only the first n values
func Take[T any](n int) Stage[T, T]

// Skip discards the first n values
func Skip[T any](n int) Stage[T, T]

// Reduce collects all values into a single result
func Reduce[T, R any](initial R, reducer func(R, T) R) func(<-chan T) R

// Pipeline chains multiple stages together
func Pipeline[T any](input <-chan T, stages ...Stage[T, T]) <-chan T

// PipelineAsync runs each stage with n workers
func PipelineAsync[T any](input <-chan T, n int, stages ...Stage[T, T]) <-chan T
```

---

## Examples

### Generator
```go
numbers := Generator(1, 2, 3, 4, 5)
for n := range numbers {
    fmt.Println(n) // 1, 2, 3, 4, 5
}
```

### Filter + Map
```go
numbers := Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

evens := Filter(func(n int) bool { return n%2 == 0 })(numbers)
doubled := Map(func(n int) int { return n * 2 })(evens)

for n := range doubled {
    fmt.Println(n) // 4, 8, 12, 16, 20
}
```

### Pipeline
```go
input := Generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

output := Pipeline(input,
    Filter(func(n int) bool { return n%2 == 0 }),
    Map(func(n int) int { return n * 2 }),
    Take(3),
)

for n := range output {
    fmt.Println(n) // 4, 8, 12
}
```

### Reduce
```go
numbers := Generator(1, 2, 3, 4, 5)
sum := Reduce(0, func(acc, n int) int { return acc + n })(numbers)
// sum: 15
```

---

## Instructions

1. Implement `Generator` - create channel, spawn goroutine to send values, close when done
2. Implement `Filter` - return stage that only passes matching values
3. Implement `Map` - return stage that transforms each value
4. Implement `Take` - return stage that stops after n values
5. Implement `Skip` - return stage that discards first n values
6. Implement `Reduce` - consume channel and return single value
7. Implement `Pipeline` - chain stages together
8. (Bonus) Implement `PipelineAsync` - parallel workers per stage
9. Run tests with `go test -v`

---

## Hints

### Basic
- Each stage is a function that takes input channel and returns output channel
- Stages spawn their own goroutine to process values
- Remember to close output channels when input is exhausted
- Use closure to capture parameters (n, predicate, transform)

### Intermediate
- Take needs to close output early without consuming entire input
- Skip counts items before forwarding; after skip count, forwards all
- Pipeline composes stages by passing output of one to input of next
- For PipelineAsync, you can use fan-out/fan-in within each stage

### Solution Pattern
```go
func Filter[T any](predicate func(T) bool) Stage[T, T] {
    return func(input <-chan T) <-chan T {
        output := make(chan T)
        go func() {
            defer close(output)
            for v := range input {
                if predicate(v) {
                    output <- v
                }
            }
        }()
        return output
    }
}
```

---

## Think About

1. What happens if a downstream stage stops consuming (closes early)?
2. How would you add error handling to a pipeline?
3. Why return a Stage function instead of taking the channel directly?
4. How could you add metrics/logging to observe pipeline throughput?

---

## What This Teaches

- **Pipeline pattern** - Composable data processing stages
- **Higher-order functions** - Functions that return functions
- **Channel composition** - Building complex flows from simple parts
- **Generics** - Type-safe reusable pipeline components
- **Lazy evaluation** - Processing on-demand as data flows
- **Separation of concerns** - Each stage has single responsibility
