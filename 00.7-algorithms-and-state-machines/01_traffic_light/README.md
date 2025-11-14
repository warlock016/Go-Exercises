# Exercise 01: Traffic Light State Machine

**Difficulty:** Easy
**Time:** 20-30 minutes
**Concepts:** Simple state transitions, enums, state tracking

## Learning Goal

Learn to implement a basic state machine where each state transitions to exactly one next state. This builds intuition for state-based programming.

## The Problem

Implement a traffic light controller that cycles through three states:

```
RED → GREEN → YELLOW → RED → ...
```

Your controller should:
- Start in RED state
- Transition to next state on each `Next()` call
- Return current state with `Current()`
- Track how many times it's cycled (completed RED → GREEN → YELLOW → RED)

## Function Signatures

```go
type TrafficLight struct {
    // Your fields here
}

func NewTrafficLight() *TrafficLight

func (t *TrafficLight) Next()

func (t *TrafficLight) Current() string

func (t *TrafficLight) Cycles() int  // How many complete cycles (RED→GREEN→YELLOW→RED)
```

## Examples

```go
light := NewTrafficLight()
light.Current()  // → "RED"
light.Cycles()   // → 0

light.Next()
light.Current()  // → "GREEN"
light.Cycles()   // → 0

light.Next()
light.Current()  // → "YELLOW"
light.Cycles()   // → 0

light.Next()
light.Current()  // → "RED" (cycled back)
light.Cycles()   // → 1 (completed one cycle)

light.Next()
light.Current()  // → "GREEN"
light.Cycles()   // → 1
```

## What This Teaches

- **State representation:** How to store current state
- **State transitions:** Moving from one state to another
- **Cycle detection:** Recognizing when you've returned to start
- **Encapsulation:** Hiding state inside a struct

Ready to implement? Open `traffic_light.go` and look for `TODO(human)` markers!
