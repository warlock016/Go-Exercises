# Exercise 05: Closures

**Learning Goal:** Master closures and variable capture for stateful functions

---

## 📝 Problem Description

A closure is a function that "closes over" variables from its surrounding scope. This creates functions with private state - a powerful pattern for:
- Counters and accumulators
- Callbacks with context
- Factory functions
- Iterator patterns

You'll implement:
1. `Counter()` - Returns a function that increments and returns a count
2. `Accumulator()` - Returns a function that adds to a running sum
3. `Multiplier()` - Returns a function that multiplies by a fixed factor
4. `RateLimiter()` - Returns a function that tracks calls within a time window

---

## 🎯 Function Signatures

```go
func Counter() func() int

func Accumulator() func(int) int

func Multiplier(factor int) func(int) int

func RateLimiter(maxCalls int, window time.Duration) func() bool
```

---

## 📖 Examples

```go
// Counter example
count := Counter()
count()    // 1
count()    // 2
count()    // 3

count2 := Counter()  // independent counter
count2()   // 1
count()    // 4

// Accumulator example
acc := Accumulator()
acc(5)     // 5
acc(3)     // 8
acc(-2)    // 6

// Multiplier example
double := Multiplier(2)
triple := Multiplier(3)

double(5)  // 10
triple(5)  // 15
double(7)  // 14

// RateLimiter example
limiter := RateLimiter(3, 1*time.Second)
limiter()  // true (call 1)
limiter()  // true (call 2)
limiter()  // true (call 3)
limiter()  // false (limit reached)
// After 1 second...
limiter()  // true (window reset)
```

---

## 📋 Instructions

1. Implement `Counter()` that returns a function incrementing a count
2. Implement `Accumulator()` that returns a function adding to a sum
3. Implement `Multiplier()` that captures a factor and returns a multiplier function
4. Implement `RateLimiter()` that tracks calls within a sliding time window
5. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Basic Concept</summary>

Closures capture variables from outer scope:
```go
func MakeAdder(x int) func(int) int {
    return func(y int) int {
        return x + y  // x is captured
    }
}

add5 := MakeAdder(5)
add5(3)  // 8
```

Each closure has its own copy of captured variables.

</details>

<details>
<summary>Intermediate Hint</summary>

For `Counter()`:
```go
func Counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

For `RateLimiter()`:
- Store a slice of call timestamps
- On each call, filter out timestamps outside the window
- Check if remaining count < maxCalls

</details>

<details>
<summary>Complete Solution</summary>

```go
package closures

import "time"

func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func Accumulator() func(int) int {
	sum := 0
	return func(n int) int {
		sum += n
		return sum
	}
}

func Multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func RateLimiter(maxCalls int, window time.Duration) func() bool {
	var calls []time.Time

	return func() bool {
		now := time.Now()

		// Remove calls outside the window
		cutoff := now.Add(-window)
		validCalls := []time.Time{}
		for _, t := range calls {
			if t.After(cutoff) {
				validCalls = append(validCalls, t)
			}
		}
		calls = validCalls

		// Check limit
		if len(calls) < maxCalls {
			calls = append(calls, now)
			return true
		}
		return false
	}
}
```

</details>

---

## 🤔 Think About

1. Why does each call to `Counter()` create an independent counter?
2. What would happen if you used a global variable instead of a closure?
3. How does the closure "remember" its state between calls?
4. What are the memory implications of closures with large captured data?

---

## 🎓 What This Teaches

- **Lexical scoping**: Functions capture variables from outer scope
- **State encapsulation**: Private state without classes/objects
- **Function factories**: Functions that create configured functions
- **Functional programming**: First-class functions and composition
- **Memory**: Each closure maintains its own captured variables
- **Iterator pattern**: Stateful iteration without explicit objects

---

**Tier:** 2 - Application
**Estimated Time:** 40-50 minutes
