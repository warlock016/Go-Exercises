# Module 04: Error Handling

**Status:** Ready to Start
**Estimated Time:** 10-14 hours
**Difficulty:** Foundation → Mastery
**Prerequisites:** Module 01 Fundamentals, Module 02 Data Structures

---

## 🎯 Learning Objectives

By completing this module, you will:

1. **Master Go's error handling philosophy** - Understand explicit error checking, error as values, and why Go doesn't use exceptions
2. **Create meaningful errors** - Use errors.New, fmt.Errorf, and custom error types to provide context
3. **Build error hierarchies** - Implement sentinel errors, error wrapping, and error chains
4. **Inspect errors effectively** - Use errors.Is and errors.As to check error types and extract values
5. **Design production-ready APIs** - Create consistent error responses for HTTP services
6. **Implement resilience patterns** - Handle transient failures with retry logic and exponential backoff
7. **Apply best practices** - Error logging, context preservation, panic recovery

---

## 📋 Module Overview

Go's approach to error handling is **explicit and intentional**. Instead of try-catch blocks that hide control flow, Go treats errors as values that must be checked and handled explicitly. This module teaches you to think in Go's error paradigm.

**Core Philosophy:**
- Errors are values, not exceptions
- Handle errors explicitly at each step
- Provide context as errors propagate up the call stack
- Make errors inspectable and actionable
- Fail fast, but fail gracefully

**Core Topics:**
- **Basic patterns**: `if err != nil`, early returns, error propagation
- **Error creation**: errors.New, fmt.Errorf, custom error types
- **Error wrapping**: %w verb, error chains, context preservation
- **Error inspection**: errors.Is, errors.As, type assertions
- **Production patterns**: API error responses, retry logic, structured logging
- **Panic recovery**: defer/recover for middleware and library code

---

## 🏗️ Exercise Structure (14 Exercises, 4 Tiers)

### Tier 1: Foundation (Exercises 01-04)
**Goal:** Master the fundamentals of Go error handling

- `01_error_checking_patterns` - if err != nil patterns, early returns, error propagation
- `02_error_creation` - errors.New, fmt.Errorf with formatting, adding context
- `03_sentinel_errors` - Package-level error variables (ErrNotFound, ErrInvalidInput), errors.Is
- `04_custom_error_types` - Implementing the error interface on custom structs

**Expected Time:** 2-3 hours
**Success Criteria:** Understand when to check errors, how to create meaningful errors, recognize sentinel error patterns

---

### Tier 2: Application (Exercises 05-08)
**Goal:** Apply error handling patterns to real-world scenarios

- `05_error_wrapping` - Using %w verb with fmt.Errorf, building error chains, preserving context
- `06_error_inspection` - errors.Is for sentinel checking, errors.As for type extraction
- `07_validation_errors` - Aggregating multiple validation errors, field-level errors
- `08_http_error_responses` - Mapping errors to HTTP status codes, JSON error bodies

**Expected Time:** 3-4 hours
**Success Criteria:** Can wrap errors with context, inspect error chains, design validation APIs, handle HTTP errors

---

### Tier 3: Integration (Exercises 09-12)
**Goal:** Build production-ready error handling systems

- `09_api_error_format` - Consistent API error structure, RFC 7807 Problem Details format
- `10_retry_with_backoff` - Transient error detection, exponential backoff, max retry logic
- `11_error_logging` - Structured logging with error context, log levels, traceability
- `12_repository_errors` - Data access layer error patterns, wrapping database errors

**Expected Time:** 4-5 hours
**Success Criteria:** Design consistent error APIs, implement retry mechanisms, integrate error logging

---

### Tier 4: Mastery (Exercises 13-14)
**Goal:** Advanced error handling for production systems

- `13_error_middleware` - HTTP panic recovery middleware, error logging, stack traces
- `14_weather_api_errors` - Refactor Weather CLI with proper error handling patterns

**Expected Time:** 3-4 hours
**Success Criteria:** Build production-grade error handling, apply patterns to real projects

---

## 📚 Key Concepts Covered

### Error Checking Patterns
- The `if err != nil` idiom
- Early returns to avoid nested error handling
- Error propagation up the call stack
- When to handle vs when to return errors
- Never ignore errors (use `_ = err` explicitly if intentional)

### Error Creation
- `errors.New("simple message")` - Static error messages
- `fmt.Errorf("format %s", val)` - Dynamic error messages with formatting
- `fmt.Errorf("context: %w", err)` - Error wrapping with %w verb (Go 1.13+)
- Custom error types implementing `error` interface

### Sentinel Errors
- Package-level error variables (e.g., `var ErrNotFound = errors.New("not found")`)
- Checking sentinel errors with `errors.Is(err, ErrNotFound)`
- When to use sentinel errors vs custom types
- Exported vs unexported sentinel errors

### Custom Error Types
```go
type ValidationError struct {
    Field string
    Value interface{}
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s (got %v)", e.Field, e.Message, e.Value)
}
```

### Error Wrapping
- `fmt.Errorf("operation failed: %w", originalErr)` - Wraps error with context
- Error chains preserve the full context path
- `errors.Unwrap(err)` - Gets the wrapped error
- Best practice: Add context at each layer without losing original error

### Error Inspection
- `errors.Is(err, target)` - Checks if err or any wrapped error equals target
- `errors.As(err, &target)` - Extracts a specific error type from the chain
- Type assertions for custom error types
- Avoiding brittle string matching on error messages

### HTTP Error Patterns
- Mapping domain errors to HTTP status codes
- JSON error response structures
- RFC 7807 Problem Details format
- Client vs server error distinction (4xx vs 5xx)

### Resilience Patterns
- Detecting transient errors (network timeouts, temporary unavailability)
- Exponential backoff: delay *= 2 with jitter
- Maximum retry attempts
- Circuit breaker pattern (preview)

### Logging and Observability
- Structured logging (key-value pairs)
- Log levels: DEBUG, INFO, WARN, ERROR
- Error context preservation
- Correlation IDs for tracing

### Panic and Recovery
- `panic()` for unrecoverable errors only
- `defer` with `recover()` for graceful recovery
- Middleware panic recovery for HTTP servers
- Converting panics to errors at boundaries

---

## 🎓 Success Criteria

**To complete this module:**
- ✅ All 14 exercises completed with tests passing (>90% test pass rate)
- ✅ Can explain Go's error handling philosophy vs exceptions
- ✅ Understand when to use sentinel errors vs custom types
- ✅ Can wrap errors with context using %w
- ✅ Can inspect error chains with errors.Is and errors.As
- ✅ Can design consistent error APIs for HTTP services
- ✅ Can implement retry logic with exponential backoff
- ✅ Understand when to use panic vs error returns

**Self-Assessment Questions:**
1. Why does Go use explicit error checking instead of exceptions?
2. What's the difference between `%v` and `%w` in fmt.Errorf?
3. When should you create a sentinel error vs a custom error type?
4. How do you check if an error is a specific type in a wrapped chain?
5. What HTTP status code should you return for validation errors?
6. When is it appropriate to use panic()?

---

## 💡 Learning Tips

1. **Check errors immediately** - Never defer error checking; handle it right after the call
2. **Add context at each layer** - Wrap errors with information about what operation failed
3. **Make errors actionable** - Include enough information for users/developers to fix the issue
4. **Test error paths** - Write tests that trigger error conditions, not just happy paths
5. **Read standard library code** - See how `io`, `os`, and `net/http` handle errors
6. **Don't panic** - Panic is for programmer errors (bugs), not expected errors

---

## 🔗 Resources

**Official Documentation:**
- [Error handling and Go](https://go.dev/blog/error-handling-and-go) - Official Go blog
- [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors) - Error wrapping and inspection
- [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover) - Understanding panic recovery

**Recommended Reading:**
- [Effective Go: Errors](https://go.dev/doc/effective_go#errors) - Best practices
- [Go by Example: Errors](https://gobyexample.com/errors) - Code examples
- [RFC 7807 Problem Details](https://tools.ietf.org/html/rfc7807) - Standard API error format

**Standard Library Examples:**
- `errors` package - errors.New, errors.Is, errors.As, errors.Unwrap
- `fmt` package - fmt.Errorf with %w verb
- `net/http` package - HTTP status codes and error responses
- `io` package - io.EOF and other sentinel errors

---

## 🚀 Getting Started

1. **Start with Tier 1** - Build muscle memory for `if err != nil` patterns
2. **Read error messages carefully** - They teach you Go's error conventions
3. **Experiment with wrapping** - See how %w preserves error chains
4. **Write error tests first** - Think about failure modes before implementing
5. **Apply to real code** - Exercise 14 refactors your Weather CLI with proper error handling

**Target Pace:** 2-3 exercises per day with all tests passing

---

## 📊 Progress Tracking

Track your progress in `PROGRESS.md`:
- Mark exercises complete when all tests pass
- Note patterns that "clicked" vs patterns that need more practice
- Update skill matrix levels as you progress
- Record insights about error handling design decisions

---

**Module Created:** 2025-12-04
**Ready to Start:** Yes - All 14 exercises created and ready
**Next Module:** 05 - Concurrency Fundamentals (unlocks after completing 80%+ of Module 04)
