# Exercise 12: Integration vs Unit Tests

**Learning Goal:** Understand the difference between unit and integration tests, use build tags to separate them.

---

## Problem Description

**Unit tests**: Test single functions in isolation, fast, no external dependencies
**Integration tests**: Test multiple components together, slower, may use databases/APIs

Go uses build tags to separate test types:
```go
//go:build integration

package mypackage
```

---

## Your Task

1. Write unit tests (fast, isolated)
2. Write integration tests (build tag: `//go:build integration`)
3. Run separately:
   - Unit: `go test`
   - Integration: `go test -tags=integration`
   - All: `go test -tags=integration ./...`

---

**Next Exercise:** `13_testing_middleware` - Test HTTP middleware
