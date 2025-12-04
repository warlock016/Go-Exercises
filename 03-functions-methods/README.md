# Module 03: Functions & Methods

**Status:** Ready to Start
**Estimated Time:** 10-14 hours
**Difficulty:** Foundation → Advanced
**Prerequisites:** Module 01 Fundamentals, Module 02 Data Structures

---

## 🎯 Learning Objectives

By completing this module, you will:

1. **Master function design patterns** - Variadic functions, multiple returns, closures, higher-order functions
2. **Understand methods and receivers** - Value vs pointer receivers, method sets, method chaining
3. **Learn functional programming techniques** - Closures, function factories, composition patterns
4. **Build middleware patterns** - HTTP middleware chains, reusable wrapper functions
5. **Implement advanced patterns** - Functional options, dependency injection, builder pattern
6. **Design flexible APIs** - Fluent interfaces, configuration patterns, testable code

---

## 📋 Module Overview

This module shifts your focus from **data manipulation** to **behavior design**. You'll move beyond "how do I store data?" to "how do I design flexible, testable, and composable functions?"

**Core Topics:**
- **Functions**: Variadic parameters, multiple returns, named returns, defer
- **Methods**: Receiver types (value vs pointer), method sets, method expressions
- **Closures**: Capturing state, function factories, stateful functions
- **Higher-Order Functions**: Functions as parameters, functions as return values
- **Design Patterns**: Middleware, options pattern, builder pattern, dependency injection
- **HTTP Integration**: Handlers, middleware chains, request processing

---

## 🏗️ Exercise Structure (15 Exercises, 4 Tiers)

### Tier 1: Foundation (Exercises 01-04)
**Goal:** Master core function and method concepts

- `01_variadic_functions` - Sum, Max, Concat with variable arguments
- `02_multiple_returns` - Functions returning (result, error) patterns
- `03_value_vs_pointer_receivers` - Method receiver behavior differences
- `04_method_sets` - Attaching methods to custom types

**Expected Time:** 2-3 hours
**Success Criteria:** Understand when to use each function pattern, know value vs pointer receiver rules

---

### Tier 2: Application (Exercises 05-08)
**Goal:** Apply functional programming patterns and HTTP handlers

- `05_closures` - Counter, accumulator with captured state
- `06_higher_order_functions` - Map, Filter, Reduce accepting function parameters
- `07_function_factories` - Functions that return configured functions
- `08_http_handler_functions` - HTTP handlers, http.HandlerFunc type (🌐 HTTP)

**Expected Time:** 3-4 hours
**Success Criteria:** Can write closures and higher-order functions, understand HTTP handler patterns

---

### Tier 3: Integration (Exercises 09-12)
**Goal:** Combine patterns to build real-world functionality

- `09_middleware_pattern` - Chainable middleware for logging, timing (🌐 HTTP)
- `10_method_chaining` - Builder pattern with fluent interface
- `11_dependency_injection` - Passing functions/interfaces for testability
- `12_pipeline_builder` - Functional transformation pipelines

**Expected Time:** 4-5 hours
**Success Criteria:** Can build middleware chains, implement builder pattern, design testable code

---

### Tier 4: Mastery (Exercises 13-15)
**Goal:** Master advanced patterns used in production Go code

- `13_options_pattern` - Functional options for flexible configuration
- `14_http_middleware_chain` - Complete middleware stack with recovery (🌐 HTTP)
- `15_api_client_builder` - Fluent HTTP client builder (🌐 HTTP)

**Expected Time:** 4-5 hours
**Success Criteria:** Can implement options pattern, build production-ready middleware, design fluent APIs

---

## 📚 Key Concepts Covered

### Functions
- Function signatures and parameter lists
- Variadic parameters: `func Sum(nums ...int) int`
- Multiple return values: `func Divide(a, b int) (int, error)`
- Named return values and naked returns
- Anonymous functions: `func(x int) int { return x * 2 }`
- Functions as values: `var fn func(int) int = add`
- Function types and type aliases

### Methods
- Method declaration: `func (t Type) Method() {}`
- Value receivers: `func (t Type)` - operates on copy
- Pointer receivers: `func (t *Type)` - operates on original
- Method sets: which methods are available on T vs *T
- Method expressions: `Type.Method` and `(*Type).Method`
- Method values: `instance.Method` as a function value
- Methods on non-struct types: `type Counter int`

### Closures
- Lexical scoping and variable capture
- Closure state: private variables
- Function factories: functions returning functions
- Stateful closures: counters, accumulators
- Closure gotchas: loop variable capture
- Use cases: callbacks, event handlers, iterators

### Higher-Order Functions
- Functions accepting functions: `func Apply(fn func(int) int, x int) int`
- Functions returning functions: `func Multiplier(factor int) func(int) int`
- Callback patterns
- Function composition: combining multiple functions
- Partial application and currying
- Map/Filter/Reduce implementations

### Design Patterns
- **Middleware Pattern**: Wrapping handlers with additional behavior
- **Options Pattern**: Functional options for flexible initialization
- **Builder Pattern**: Fluent method chaining for object construction
- **Dependency Injection**: Passing dependencies as parameters
- **Pipeline Pattern**: Chaining transformations
- **Adapter Pattern**: Converting function signatures

### HTTP Integration
- `http.Handler` interface: `ServeHTTP(w ResponseWriter, r *Request)`
- `http.HandlerFunc` type: function signature matching Handler
- Handler wrapping and middleware
- Request/response processing
- Middleware chaining
- Recovery and error handling

---

## 🎓 Success Criteria

**To complete this module:**
- ✅ All 15 exercises completed with tests passing (>85% test pass rate)
- ✅ Can explain value vs pointer receivers and when to use each
- ✅ Can write closures and explain variable capture
- ✅ Can implement higher-order functions (map, filter, reduce)
- ✅ Can build HTTP middleware chains
- ✅ Can implement the options pattern for configuration
- ✅ Can design fluent APIs with method chaining
- ✅ Understand function types and how to use them effectively

**Self-Assessment Questions:**
1. When should you use a pointer receiver vs a value receiver?
2. What is a closure and what are common use cases?
3. How do you chain multiple HTTP middleware functions?
4. What problems does the options pattern solve?
5. How do you make a function accept another function as a parameter?
6. What's the difference between `http.Handler` and `http.HandlerFunc`?

---

## 💡 Learning Tips

1. **Visualize function flow** - Draw diagrams showing how closures capture variables
2. **Experiment with receivers** - Try both value and pointer receivers to see behavior differences
3. **Study standard library** - Look at `http.HandlerFunc`, `sort.Sort`, `io.Reader` for patterns
4. **Build incrementally** - Start with simple closures before complex middleware chains
5. **Focus on testability** - Understand how dependency injection makes code easier to test
6. **Read production code** - Popular libraries use these patterns extensively

---

## 🔗 Resources

**Official Documentation:**
- [Functions](https://go.dev/ref/spec#Function_declarations) - Language specification
- [Methods](https://go.dev/ref/spec#Method_declarations) - Method declaration rules
- [net/http package](https://pkg.go.dev/net/http) - HTTP handlers and middleware

**Recommended Reading:**
- "Effective Go" sections on functions, methods, and interfaces
- [Go by Example](https://gobyexample.com/) - Functions, Closures, Methods sections
- "Functional Options in Go" by Dave Cheney
- "The Middleware Pattern in Go" - various blog posts

**In This Repo:**
- `resources/` - Check for function design guides (may be created as needed)
- `02-data-structures/12_custom_types/` - Preview of methods on custom types

---

## 🚀 Getting Started

1. **Start with Tier 1** - Master basic function and method syntax before advanced patterns
2. **Read each README carefully** - Understand the pattern before implementing
3. **Run tests frequently** - `go test -v` in each exercise directory
4. **Experiment with variations** - Try different receiver types, closure patterns
5. **Build on previous exercises** - Later exercises combine multiple patterns

**Target Pace:** 2-3 exercises per day with all tests passing

---

## 📊 Progress Tracking

Track your progress in `PROGRESS.md`:
- Mark exercises complete when all tests pass
- Note time spent and which patterns felt natural vs challenging
- Update skill matrix levels as you progress
- Record insights about when to use each pattern

---

## 🌟 Key Patterns to Master

By the end of this module, you should be comfortable with:

1. **Variadic functions** - For flexible parameter lists
2. **Error handling pattern** - Returning `(result, error)`
3. **Method receivers** - Choosing value vs pointer appropriately
4. **Closures** - For stateful behavior and callbacks
5. **Higher-order functions** - For generic algorithms
6. **HTTP middleware** - For request/response processing
7. **Options pattern** - For flexible configuration
8. **Builder pattern** - For fluent APIs
9. **Dependency injection** - For testable code

---

**Module Created:** 2025-12-04
**Ready to Start:** Yes - All 15 exercises created and ready
**Next Module:** 04 - Error Handling (unlocks after completing 80%+ of Module 03)
