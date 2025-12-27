# Exercise 09: Mock for Testing

**Concept:** Dependency injection and test doubles using interfaces
**Difficulty:** Medium-Hard
**Estimated Time:** 60 minutes

## Learning Goal

Learn how interfaces enable testable code through dependency injection and mocking.

## The Problem

Code that depends on concrete types is hard to test:

```go
// Hard to test - depends on real time
type Scheduler struct{}

func (s *Scheduler) ShouldRun() bool {
    now := time.Now()
    return now.Hour() >= 9 && now.Hour() < 17
}
```

With interfaces, you can inject test doubles:

```go
type TimeProvider interface {
    Now() time.Time
}

type Scheduler struct {
    Time TimeProvider
}

func (s *Scheduler) ShouldRun() bool {
    now := s.Time.Now()
    return now.Hour() >= 9 && now.Hour() < 17
}
```

## Your Task

Create a testable scheduler using dependency injection.

### Interface

**TimeProvider** - provides current time
- `Now() time.Time`

### Types

1. **RealTime** - returns actual current time
2. **MockTime** - returns configurable time for testing
3. **Scheduler** - uses TimeProvider to check business hours (9 AM - 5 PM)

## Function Signatures

```go
type TimeProvider interface {
    Now() time.Time
}

type RealTime struct{}
type MockTime struct {
    CurrentTime time.Time
}

type Scheduler struct {
    TimeProvider TimeProvider
}

func (r *RealTime) Now() time.Time
func (m *MockTime) Now() time.Time
func (s *Scheduler) ShouldRun() bool
func (s *Scheduler) TimeUntilNextRun() time.Duration
```

## Examples

```go
// Production code uses RealTime
scheduler := &Scheduler{TimeProvider: &RealTime{}}
if scheduler.ShouldRun() {
    // Run task
}

// Test code uses MockTime
mockTime := &MockTime{CurrentTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)}
scheduler := &Scheduler{TimeProvider: mockTime}
scheduler.ShouldRun()  // → true (10 AM is during business hours)
```

## What This Teaches

- Dependency injection
- Test doubles/mocking
- Interface-based design
- Testable code patterns

---

**Next up:** Exercise 10 - Small Interfaces
