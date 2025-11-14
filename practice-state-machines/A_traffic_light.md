# Exercise A: Traffic Light State Machine

**Time Estimate:** 15-20 minutes
**Difficulty:** Easy
**Goal:** Learn basic state transitions without complex logic

## The Problem

Implement a traffic light controller that cycles through states:
```
RED → GREEN → YELLOW → RED → ...
```

Each call to `Next()` advances to the next state.

## Function Signature

```go
type TrafficLight struct {
    state string
}

func NewTrafficLight() *TrafficLight {
    // Start in RED state
}

func (t *TrafficLight) Next() {
    // Transition to next state
}

func (t *TrafficLight) Current() string {
    // Return current state
}
```

## Examples

```go
light := NewTrafficLight()
light.Current() // → "RED"

light.Next()
light.Current() // → "GREEN"

light.Next()
light.Current() // → "YELLOW"

light.Next()
light.Current() // → "RED" (cycles back)
```

## Your Task

1. Create `traffic_light.go`
2. Implement the three functions
3. Use a `switch` statement on `state` to determine next state
4. Test manually by calling `Next()` multiple times

## Solution Approach

```go
func (t *TrafficLight) Next() {
    switch t.state {
    case "RED":
        t.state = "GREEN"
    case "GREEN":
        t.state = "YELLOW"
    case "YELLOW":
        t.state = "RED"
    }
}
```

## Key Learning

- **State:** A variable that determines behavior
- **State Transition:** Moving from one state to another based on current state
- **No complex logic:** Just simple mappings

## After Completing

Try extending:
- Add a `CanGo() bool` method (true for GREEN, false otherwise)
- Add duration tracking (how long in current state)
- Add `Prev()` to go backwards

**Time to complete:** Write the code and test it!
