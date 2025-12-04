# Exercise 06: Higher-Order Functions

**Learning Goal:** Master functions that accept or return other functions

---

## 📝 Problem Description

Higher-order functions are functions that:
1. Accept functions as parameters
2. Return functions as results
3. Or both

This enables powerful patterns like map/filter/reduce, callbacks, and function composition. You'll implement generic algorithms that work with any function.

---

## 🎯 Function Signatures

```go
func Map(slice []int, fn func(int) int) []int

func Filter(slice []int, predicate func(int) bool) []int

func Reduce(slice []int, initial int, fn func(int, int) int) int

func Compose(f func(int) int, g func(int) int) func(int) int

func ForEach(slice []int, fn func(int))

func Any(slice []int, predicate func(int) bool) bool

func All(slice []int, predicate func(int) bool) bool
```

---

## 📖 Examples

```go
// Map example
nums := []int{1, 2, 3, 4}
doubled := Map(nums, func(n int) int { return n * 2 })
// []int{2, 4, 6, 8}

// Filter example
evens := Filter(nums, func(n int) bool { return n%2 == 0 })
// []int{2, 4}

// Reduce example
sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
// 10

product := Reduce(nums, 1, func(acc, n int) int { return acc * n })
// 24

// Compose example
double := func(n int) int { return n * 2 }
addTen := func(n int) int { return n + 10 }
composed := Compose(double, addTen)
composed(5)  // double(addTen(5)) = double(15) = 30

// Any/All examples
hasEven := Any(nums, func(n int) bool { return n%2 == 0 })  // true
allPositive := All(nums, func(n int) bool { return n > 0 })  // true
allEven := All(nums, func(n int) bool { return n%2 == 0 })   // false
```

---

## 📋 Instructions

1. Implement `Map` to transform each element
2. Implement `Filter` to keep only matching elements
3. Implement `Reduce` to accumulate a single value
4. Implement `Compose` to combine two functions (f ∘ g)
5. Implement `ForEach` to execute a function on each element
6. Implement `Any` to check if at least one element matches
7. Implement `All` to check if all elements match
8. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Basic Concept</summary>

Higher-order functions accept functions as parameters:
```go
func Apply(n int, fn func(int) int) int {
    return fn(n)
}

result := Apply(5, func(x int) int { return x * 2 })  // 10
```

</details>

<details>
<summary>Intermediate Hint</summary>

For `Map`:
- Create a new slice with same length
- Apply `fn` to each element
- Return the new slice

For `Reduce`:
- Start with `initial` as accumulator
- Call `fn(accumulator, element)` for each element
- Return final accumulator

For `Compose`:
- Return a function that calls `g` then `f`
- `Compose(f, g)(x)` = `f(g(x))`

</details>

<details>
<summary>Complete Solution</summary>

```go
package higher_order_functions

func Map(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func Filter(slice []int, predicate func(int) bool) []int {
	result := []int{}
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce(slice []int, initial int, fn func(int, int) int) int {
	acc := initial
	for _, v := range slice {
		acc = fn(acc, v)
	}
	return acc
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func ForEach(slice []int, fn func(int)) {
	for _, v := range slice {
		fn(v)
	}
}

func Any(slice []int, predicate func(int) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

func All(slice []int, predicate func(int) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}
```

</details>

---

## 🤔 Think About

1. How are these functions different from loops?
2. What are the benefits of separating the algorithm from the operation?
3. Why does Go not have generics for these in the standard library? (Note: Go 1.18+ has generics)
4. How would you implement `MapStrings` for string slices?

---

## 🎓 What This Teaches

- **Higher-order functions**: Functions as first-class values
- **Functional patterns**: Map, filter, reduce operations
- **Function composition**: Combining simple functions into complex ones
- **Separation of concerns**: Algorithm separate from operation
- **Declarative style**: What to do, not how to do it
- **Code reuse**: Generic algorithms work with any function

---

**Tier:** 2 - Application
**Estimated Time:** 40-50 minutes
