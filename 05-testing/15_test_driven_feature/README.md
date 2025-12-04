# Exercise 15: Test-Driven Development (TDD)

**Learning Goal:** Experience the TDD workflow - red, green, refactor. Build a feature test-first.

---

## Problem Description

Test-Driven Development (TDD) is a development process where you:
1. **Red**: Write a failing test
2. **Green**: Write minimal code to make it pass
3. **Refactor**: Improve the code while keeping tests passing

This exercise guides you through building a URL shortener using TDD.

---

## The TDD Cycle

```
1. RED: Write a failing test
   ├─ Think about the API/interface you want
   ├─ Write test for one small behavior
   └─ Run test - it MUST fail

2. GREEN: Make the test pass
   ├─ Write minimal code to pass the test
   ├─ Don't worry about perfect code yet
   └─ Run test - it MUST pass

3. REFACTOR: Improve the code
   ├─ Clean up duplication
   ├─ Improve names and structure
   └─ Run tests - they MUST still pass

Repeat for each feature!
```

---

## Feature to Build: URL Shortener

Build a simple URL shortener with these features:

```go
type URLShortener struct { ... }

// Shorten creates a short code for a URL
func (s *URLShortener) Shorten(url string) (string, error)

// Expand retrieves the original URL
func (s *URLShortener) Expand(code string) (string, error)

// Stats returns usage statistics
func (s *URLShortener) Stats(code string) (*Stats, error)
```

---

## TDD Steps to Follow

### Step 1: Shorten a URL
1. **RED**: Write test for `Shorten("https://example.com")`
2. **GREEN**: Return a hardcoded short code
3. **REFACTOR**: (nothing to refactor yet)

### Step 2: Expand a URL
1. **RED**: Write test for `Expand(code)` returning original URL
2. **GREEN**: Store URL in map, return it
3. **REFACTOR**: (maybe extract map initialization)

### Step 3: Handle Errors
1. **RED**: Test `Expand("nonexistent")` returns error
2. **GREEN**: Check map, return error if not found
3. **REFACTOR**: Add error constants

### Step 4: Validate URLs
1. **RED**: Test `Shorten("")` returns error
2. **GREEN**: Add validation
3. **REFACTOR**: Extract validation function

### Step 5: Track Statistics
1. **RED**: Test `Stats(code)` returns access count
2. **GREEN**: Track access count on Expand
3. **REFACTOR**: Create Stats struct

### Step 6: Handle Duplicates
1. **RED**: Test shortening same URL twice returns same code
2. **GREEN**: Check if URL already exists before generating code
3. **REFACTOR**: Extract lookup logic

---

## Your Task

Follow the TDD steps above. For each step:
1. Write the test FIRST
2. Run `go test` - it should FAIL
3. Write minimal code to pass
4. Run `go test` - it should PASS
5. Refactor if needed
6. Move to next step

---

## Example: Step 1

### RED - Write Failing Test

```go
func TestShorten(t *testing.T) {
    s := NewURLShortener()
    code, err := s.Shorten("https://example.com")

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if code == "" {
        t.Error("expected non-empty code")
    }
}
```

Run: `go test` → FAILS (function doesn't exist)

### GREEN - Make It Pass

```go
type URLShortener struct{}

func NewURLShortener() *URLShortener {
    return &URLShortener{}
}

func (s *URLShortener) Shorten(url string) (string, error) {
    return "abc123", nil  // Hardcoded for now
}
```

Run: `go test` → PASSES

### REFACTOR - (nothing to refactor yet)

Move to Step 2!

---

## Instructions

1. Create `shortener.go` and `shortener_test.go`
2. Follow Steps 1-6 in order
3. For each step: RED → GREEN → REFACTOR
4. Run tests after every change
5. Commit after each step completes (optional but recommended)

---

## Key TDD Principles

**RED Phase:**
- Test should fail for the right reason
- Test one thing at a time
- Think about the interface you want

**GREEN Phase:**
- Write simplest code that passes
- It's OK to hardcode initially
- Don't write extra features

**REFACTOR Phase:**
- Tests must keep passing
- Improve names, extract functions
- Remove duplication

---

## Success Criteria

- [ ] All 6 steps completed
- [ ] Each step followed RED-GREEN-REFACTOR
- [ ] All tests passing
- [ ] Code is clean and readable
- [ ] No premature optimization
- [ ] Tests document the features

---

## What This Teaches

- **TDD workflow** - Red, green, refactor cycle
- **Design through tests** - Tests drive the API design
- **Incremental development** - Build features one at a time
- **Confidence** - Tests let you refactor safely
- **Living documentation** - Tests show how to use the code

---

## Think About

1. How does writing tests first change your design?
2. What happens if you skip the RED phase?
3. When should you refactor vs keep going?
4. How does TDD help with requirements understanding?
5. What are the limits of TDD? When might you not use it?

---

## After Completing

Reflect on the TDD experience:
- Was it slower or faster than coding first?
- Did writing tests first change your design?
- Did you feel more confident in your code?
- What was hardest about the TDD process?
- Will you use TDD in future projects?

---

**Congratulations!** You've completed Module 05: Testing!

You now know:
- Basic testing with testing.T
- Table-driven tests and subtests
- Test helpers and fixtures
- Error testing patterns
- Benchmarking
- HTTP testing with httptest
- Test coverage analysis
- Integration vs unit testing
- Middleware testing
- Test-Driven Development

**Next Steps:**
- Apply these patterns to your projects
- Write tests for existing code
- Try TDD on your next feature
- Explore testing in other modules
