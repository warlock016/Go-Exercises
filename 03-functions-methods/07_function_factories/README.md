# Exercise 07: Function Factories

**Learning Goal:** Master functions that create and return configured functions

---

## 📝 Problem Description

Function factories are functions that return other functions with pre-configured behavior. This pattern is useful for:
- Creating specialized versions of generic functions
- Dependency injection
- Configuration without globals
- Creating validators, formatters, and processors

---

## 🎯 Function Signatures

```go
func MakeAdder(x int) func(int) int

func MakeGreeter(greeting string) func(string) string

func MakeValidator(min, max int) func(int) bool

func MakeFormatter(prefix, suffix string) func(string) string

func MakePowerFunction(exponent int) func(float64) float64
```

---

## 📖 Examples

```go
// MakeAdder
add5 := MakeAdder(5)
add5(3)   // 8
add5(10)  // 15

// MakeGreeter
hello := MakeGreeter("Hello")
hola := MakeGreeter("Hola")
hello("Alice")  // "Hello, Alice!"
hola("Bob")     // "Hola, Bob!"

// MakeValidator
inRange := MakeValidator(1, 10)
inRange(5)   // true
inRange(15)  // false

// MakeFormatter
brackets := MakeFormatter("[", "]")
parens := MakeFormatter("(", ")")
brackets("test")  // "[test]"
parens("test")    // "(test)"

// MakePowerFunction
square := MakePowerFunction(2)
cube := MakePowerFunction(3)
square(4.0)  // 16.0
cube(3.0)    // 27.0
```

---

## 📋 Instructions

1. Implement `MakeAdder` that creates addition functions
2. Implement `MakeGreeter` that creates greeting functions
3. Implement `MakeValidator` that creates range validators
4. Implement `MakeFormatter` that creates string formatters
5. Implement `MakePowerFunction` that creates exponent functions
6. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Basic Concept</summary>

Function factories capture configuration and return specialized functions:
```go
func MakeMultiplier(factor int) func(int) int {
    return func(n int) int {
        return n * factor
    }
}
```

</details>

<details>
<summary>Complete Solution</summary>

```go
package function_factories

import (
	"fmt"
	"math"
)

func MakeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func MakeGreeter(greeting string) func(string) string {
	return func(name string) string {
		return fmt.Sprintf("%s, %s!", greeting, name)
	}
}

func MakeValidator(min, max int) func(int) bool {
	return func(n int) bool {
		return n >= min && n <= max
	}
}

func MakeFormatter(prefix, suffix string) func(string) string {
	return func(s string) string {
		return prefix + s + suffix
	}
}

func MakePowerFunction(exponent int) func(float64) float64 {
	return func(base float64) float64 {
		return math.Pow(base, float64(exponent))
	}
}
```

</details>

---

## 🤔 Think About

1. How do factories differ from closures?
2. When would you use a factory instead of a struct with methods?
3. What are the memory implications of creating many factory functions?
4. How could you use factories for dependency injection?

---

## 🎓 What This Teaches

- **Function factories**: Creating specialized functions from generic ones
- **Configuration**: Pre-configuring behavior without global state
- **Partial application**: Fixing some arguments, leaving others open
- **Code reuse**: Creating variations without duplication
- **Functional patterns**: Higher-order programming techniques

---

**Tier:** 2 - Application
**Estimated Time:** 30-40 minutes
