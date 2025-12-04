# Module 05: Testing

**Status:** Ready to Start
**Estimated Time:** 10-14 hours
**Difficulty:** Foundation → Mastery
**Prerequisites:** Modules 01-02 (Fundamentals, Data Structures)

---

## 🎯 Learning Objectives

By completing this module, you will:

1. **Master Go's testing package** - Understand testing.T, testing.B, and the standard library testing tools
2. **Write table-driven tests** - The idiomatic Go testing pattern used throughout the standard library
3. **Organize tests with subtests** - Use t.Run() for hierarchical test organization
4. **Test HTTP handlers and APIs** - Use httptest package for web service testing
5. **Write benchmarks** - Measure performance and memory allocations
6. **Understand test coverage** - Use coverage tools to identify untested code paths
7. **Apply TDD workflow** - Red-green-refactor cycle for test-driven development

---

## 📋 Module Overview

**CRITICAL CONTEXT:** Testing is a priority learning gap (40% skill level from diagnostic). This module is designed with extra explanations and progressive difficulty.

**Key Difference:** Unlike other modules where you implement code:
- The `.go` files contain **WORKING implementations** (the code to be tested)
- You implement the **TESTS** in `_test.go` files
- The `_test.go` files have TODO(human) markers with test structure guidance

This approach teaches you to:
- Read and understand existing code
- Identify test cases and edge cases
- Write comprehensive test coverage
- Use Go's testing idioms correctly

**Core Topics:**
- **Basic Testing**: testing.T methods, error reporting, test functions
- **Table-Driven Tests**: The standard Go pattern for comprehensive testing
- **Subtests**: Hierarchical test organization with t.Run()
- **Test Helpers**: Reducing duplication with t.Helper()
- **Error Testing**: Validating error conditions and error messages
- **Benchmarking**: Performance measurement with testing.B
- **HTTP Testing**: httptest.ResponseRecorder and httptest.Server
- **Coverage Analysis**: go test -cover and identifying gaps
- **TDD Workflow**: Red-green-refactor development cycle

---

## 🏗️ Exercise Structure (15 Exercises, 4 Tiers)

### Tier 1: Foundation (Exercises 01-04)
**Goal:** Learn the basics of Go testing

- `01_first_test` - testing.T methods (t.Error, t.Errorf, t.Fatal), running tests
- `02_table_driven_basics` - Test tables with struct slices, range iteration
- `03_subtests` - t.Run() for organized output, -run flag for selective running
- `04_test_helpers` - t.Helper(), setup/teardown patterns, test utilities

**Expected Time:** 2-3 hours
**Success Criteria:** All tests passing, understand basic testing.T API

---

### Tier 2: Application (Exercises 05-09)
**Goal:** Apply testing patterns to common scenarios

- `05_testing_errors` - Testing functions that return errors, error message validation
- `06_benchmarks` - testing.B, b.N loop, b.ReportAllocs(), performance measurement
- `07_test_fixtures` - testdata/ directory, golden files, file-based tests
- `08_http_handler_tests` - httptest.ResponseRecorder, testing HTTP handlers (🌐 HTTP)
- `09_testing_json_apis` - JSON response testing, content-type validation (🌐 HTTP)

**Expected Time:** 3-4 hours
**Success Criteria:** Can write comprehensive tests for functions and HTTP handlers

---

### Tier 3: Integration (Exercises 10-13)
**Goal:** Advanced testing techniques and integration tests

- `10_mocking_http_clients` - httptest.Server, interface mocking, client testing (🌐 HTTP)
- `11_test_coverage` - go test -cover, coverage profiles, identifying untested paths
- `12_integration_vs_unit` - Build tags, test isolation strategies, integration test patterns
- `13_testing_middleware` - Testing HTTP middleware functions, request/response manipulation (🌐 HTTP)

**Expected Time:** 4-5 hours
**Success Criteria:** Can write both unit and integration tests, understand mocking

---

### Tier 4: Mastery (Exercises 14-15)
**Goal:** Apply testing to real projects and TDD workflow

- `14_test_weather_client` - Comprehensive Weather CLI client tests (🌐 HTTP)
- `15_test_driven_feature` - TDD workflow: red-green-refactor, building a feature test-first

**Expected Time:** 4-5 hours
**Success Criteria:** Can apply TDD to new features, write production-quality tests

---

## 📚 Key Concepts Covered

### Testing Basics
- Test function signature: `func TestXxx(t *testing.T)`
- Test file naming: `package_test.go` vs `package.go`
- Running tests: `go test`, `go test -v`, `go test -run TestName`
- Error reporting: `t.Error()`, `t.Errorf()`, `t.Fatal()`, `t.Fatalf()`
- Test failure vs test error
- Skipping tests: `t.Skip()`, `t.SkipNow()`

### Table-Driven Testing
```go
tests := []struct {
    name string
    input Type
    want Type
}{
    {"description", inputValue, expectedValue},
}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        got := Function(tt.input)
        if got != tt.want {
            t.Errorf("Function(%v) = %v, want %v", tt.input, got, tt.want)
        }
    })
}
```

**Why Table-Driven Tests?**
- **Comprehensive**: Easy to add many test cases
- **Readable**: Each case is clear and self-documenting
- **Maintainable**: Test logic written once, data varies
- **Standard**: Used throughout Go standard library and ecosystem

### Subtests
- Hierarchical test organization with `t.Run(name, func)`
- Selective test running: `go test -run TestMain/SubtestName`
- Parallel subtests: `t.Parallel()` for concurrent test execution
- Cleanup with `t.Cleanup(func)` instead of defer

### Benchmarking
- Benchmark function signature: `func BenchmarkXxx(b *testing.B)`
- The `b.N` loop: framework controls iteration count
- Running benchmarks: `go test -bench=.`
- Memory profiling: `go test -bench=. -benchmem`
- Reporting allocations: `b.ReportAllocs()`
- Resetting timer: `b.ResetTimer()` after setup

### HTTP Testing
- `httptest.ResponseRecorder`: Captures handler responses without network
- `httptest.NewServer()`: Creates test HTTP server for client testing
- Testing handlers: status codes, headers, body content
- Testing clients: mocking external APIs
- Testing middleware: request transformation, response interception

### Test Coverage
- Basic coverage: `go test -cover`
- Coverage profile: `go test -coverprofile=coverage.out`
- HTML report: `go tool cover -html=coverage.out`
- Interpreting coverage: 80%+ is good, 100% is often unnecessary
- Coverage-driven development: use coverage to find untested paths

---

## 🎓 Success Criteria

**To complete this module:**
- ✅ All 15 exercises completed with comprehensive test suites
- ✅ Can write table-driven tests for any function
- ✅ Understand when to use t.Error vs t.Fatal
- ✅ Can organize complex tests with subtests
- ✅ Can write benchmarks and interpret results
- ✅ Can test HTTP handlers and APIs using httptest
- ✅ Understand test coverage metrics and how to improve them
- ✅ Can apply TDD workflow to new features

**Self-Assessment Questions:**
1. What's the difference between t.Error() and t.Fatal()?
2. Why is table-driven testing idiomatic in Go?
3. When should you use t.Helper()?
4. What does b.N represent in a benchmark?
5. How do you test an HTTP handler without starting a server?
6. What's a good test coverage percentage?
7. What are the three steps in TDD (red-green-refactor)?

---

## 💡 Learning Tips

1. **Read the working code first** - Understand what you're testing before writing tests
2. **Start with happy path** - Test normal operation first, then edge cases
3. **Think about edge cases** - Empty inputs, nil values, boundary conditions
4. **Use descriptive test names** - "empty string returns error" is better than "test2"
5. **Run tests frequently** - `go test -v` after every test case you add
6. **Check coverage** - `go test -cover` shows what you've missed
7. **Write tests you'd want to read** - Future you will thank present you

**Testing Philosophy:**
- Tests are documentation - they show how to use your code
- Tests give you confidence to refactor
- Good tests fail clearly when something breaks
- Test behavior, not implementation details

---

## 🔗 Resources

**Official Documentation:**
- [testing package](https://pkg.go.dev/testing) - Standard library testing documentation
- [Table-driven tests](https://go.dev/wiki/TableDrivenTests) - Go Wiki guide
- [net/http/httptest](https://pkg.go.dev/net/http/httptest) - HTTP testing utilities
- [go test command](https://pkg.go.dev/cmd/go#hdr-Test_packages) - All test flags and options

**Recommended Reading:**
- "Testing" chapter in "The Go Programming Language" book
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) - TDD-focused tutorial
- [Advanced Testing in Go](https://www.youtube.com/watch?v=8hQG7QlcLBk) - Mitchell Hashimoto talk

**In This Repo:**
- `resources/` - Check for testing guides and best practices

---

## 🚀 Getting Started

1. **Start with Exercise 01** - Even if you've written tests before, start here to learn Go idioms
2. **Read the working code** - Each exercise has working Go code for you to test
3. **Follow the TODO(human) markers** - The test files have structural guidance
4. **Run tests as you write them** - Use `go test -v` to see your tests in action
5. **Check coverage** - Use `go test -cover` to verify you've tested all paths
6. **Experiment** - Try breaking the working code to verify your tests catch the bugs

**Target Pace:** 1-2 exercises per day with comprehensive test coverage

**Commands You'll Use:**
```bash
# Run tests in current directory
go test -v

# Run tests with coverage
go test -cover

# Run specific test
go test -v -run TestFunctionName

# Run specific subtest
go test -v -run TestMain/SubtestName

# Run benchmarks
go test -bench=.

# Run benchmarks with memory stats
go test -bench=. -benchmem

# Generate coverage profile
go test -coverprofile=coverage.out

# View coverage in browser
go tool cover -html=coverage.out
```

---

## 📊 Progress Tracking

Track your progress in `PROGRESS.md`:
- Mark exercises complete when all tests pass with good coverage
- Note time spent and testing patterns learned
- Update skill matrix: Testing should improve from 40% → 80%+
- Record "aha!" moments about test design

**Special Focus Areas (Based on Diagnostic):**
- Table-driven test syntax (struggled with in diagnostic)
- When to use t.Run() vs flat tests
- Error testing patterns
- HTTP testing with httptest

---

## 🌟 Why This Module Matters

Testing is NOT just about catching bugs. Good tests:
- **Document** how your code should be used
- **Enable refactoring** with confidence
- **Catch regressions** when you change code
- **Serve as examples** for other developers
- **Force better design** - testable code is often better code

By mastering testing, you'll:
- Write more reliable code
- Refactor without fear
- Ship with confidence
- Understand codebases faster (read the tests!)
- Become a better Go developer

---

**Module Created:** 2025-12-04
**Ready to Start:** Yes - All 15 exercises created with working code
**Next Module:** 06 - Interfaces (unlocks after completing 80%+ of Module 05)
**Testing Target:** Improve from 40% → 80%+ proficiency
