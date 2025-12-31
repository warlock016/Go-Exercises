# Exercise 07: Error Handling in Concurrent Code

**Learning Goal:** Propagate and collect errors from multiple goroutines correctly

**Difficulty:** Tier 2 - Application
**Estimated Time:** 35-40 minutes

---

## Problem Description

Error handling in concurrent code is tricky. When multiple goroutines can fail, you need strategies to:
1. Collect errors from multiple sources
2. Return early on first error (fail-fast)
3. Aggregate multiple errors
4. Avoid losing errors due to unbuffered channels

This exercise explores different patterns for concurrent error handling, from basic error channels to the powerful `errgroup` package.

---

## Function Signatures

```go
// Result wraps a value and potential error
type Result[T any] struct {
    Value T
    Err   error
}

// ProcessWithErrors processes items concurrently, collecting all errors
func ProcessWithErrors(items []int, process func(int) error) []error

// FirstError runs tasks concurrently, returns first error (or nil if all succeed)
func FirstError(tasks []func() error) error

// ProcessResults processes items and returns all results
func ProcessResults[T any](items []int, process func(int) (T, error)) []Result[T]

// RunWithTimeout runs a task with timeout, returns error if timeout exceeded
func RunWithTimeout(task func() error, timeout time.Duration) error

// ParallelFetch fetches from multiple URLs concurrently
// Returns results map and any errors encountered
func ParallelFetch(urls []string, fetch func(string) (string, error)) (map[string]string, []error)
```

---

## Examples

### ProcessWithErrors
```go
items := []int{1, 2, 3, 4, 5}
process := func(i int) error {
    if i%2 == 0 {
        return fmt.Errorf("even number: %d", i)
    }
    return nil
}

errs := ProcessWithErrors(items, process)
// errs: [even number: 2, even number: 4]
```

### FirstError (Fail-Fast)
```go
tasks := []func() error{
    func() error { time.Sleep(100*time.Millisecond); return nil },
    func() error { return errors.New("quick failure") },
    func() error { time.Sleep(200*time.Millisecond); return nil },
}

err := FirstError(tasks)
// err: "quick failure" (returned quickly, didn't wait for slow tasks)
```

### ProcessResults
```go
items := []int{1, 2, 3}
process := func(i int) (string, error) {
    if i == 2 {
        return "", errors.New("failed")
    }
    return fmt.Sprintf("result-%d", i), nil
}

results := ProcessResults(items, process)
// results[0]: {Value: "result-1", Err: nil}
// results[1]: {Value: "", Err: error}
// results[2]: {Value: "result-3", Err: nil}
```

---

## Instructions

1. Implement `ProcessWithErrors` - spawn goroutines for each item, collect all errors
2. Implement `FirstError` - return as soon as any task fails (cancel others conceptually)
3. Implement `ProcessResults` using the Result type to capture both values and errors
4. Implement `RunWithTimeout` combining goroutines with timeout pattern
5. Implement `ParallelFetch` - real-world pattern for concurrent I/O
6. Run tests with `go test -v`

---

## Hints

### Basic
- Use buffered error channel: `errs := make(chan error, len(items))`
- A buffered channel prevents goroutine leaks when not all errors are read
- For FirstError, select between error channel and completion
- sync.WaitGroup helps know when all goroutines are done

### Intermediate
- For FirstError with cancellation, consider using a done channel
- For ProcessResults, you need to maintain order—use index in result
- Error channel pattern: `errCh := make(chan error, 1)` for single error capture
- Consider what happens if process function panics

### Solution Pattern
```go
func ProcessWithErrors(items []int, process func(int) error) []error {
    errs := make(chan error, len(items))
    var wg sync.WaitGroup

    for _, item := range items {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            if err := process(i); err != nil {
                errs <- err
            }
        }(item)
    }

    wg.Wait()
    close(errs)

    var result []error
    for err := range errs {
        result = append(result, err)
    }
    return result
}
```

---

## Think About

1. Why use a buffered channel for errors?
2. What happens to goroutines still running when FirstError returns?
3. How does errgroup.Group solve these problems more elegantly?
4. When would you want all errors vs fail-fast behavior?

---

## What This Teaches

- **Error channels** - Collecting errors from concurrent operations
- **Result type pattern** - Bundling values with errors
- **Fail-fast vs collect-all** - Different error handling strategies
- **Goroutine leak prevention** - Buffered channels prevent leaks
- **Timeout integration** - Combining error handling with timeouts
- **Real-world patterns** - How production code handles concurrent errors
