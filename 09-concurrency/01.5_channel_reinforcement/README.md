# Channel Reinforcement Exercises

## Purpose

These 6 exercises bridge the gap between basic goroutine/channel concepts (01-02) and more advanced patterns (03+). They specifically target the mental model challenges you encountered:

1. **Sequential thinking trap** - Goroutines run concurrently, not one after another
2. **Channel closing** - The sender closes, from inside the goroutine
3. **Synchronization via channels** - Making main wait by receiving
4. **Generator pattern mastery** - Variations beyond the basic pattern
5. **Multi-goroutine coordination** - Beyond two goroutines

Complete these exercises **in order** (1 → 2 → 3 → 4 → 5 → 6).

---

## Exercise 1: Echo

### Learning Goal
Reinforce the basic send/receive pattern with a processing step.

### Connection to Previous Work
Similar to `Sum(values)` from 02_channel_fundamentals, but simpler - one value in, one value out.

### Function Signature
```go
// Echo sends a value to a goroutine which doubles it and returns the result
func Echo(value int) int
```

### Examples
```
Echo(5)  → 10
Echo(0)  → 0
Echo(-3) → -6
```

### Visual Diagram
```
    main                goroutine
      │                     │
      │──── value ─────────►│
      │                     │ (doubles it)
      │◄──── result ────────│
      │                     │
   return
```

### Hints
<details>
<summary>Hint 1 (Basic)</summary>
You need two channels: one for sending to the goroutine, one for receiving back.
</details>

<details>
<summary>Hint 2 (Pattern)</summary>
The goroutine receives from one channel, processes, sends to another channel.
</details>

<details>
<summary>Hint 3 (Solution structure)</summary>

```
1. Create two channels
2. Start goroutine that: receives → doubles → sends
3. Main sends value
4. Main receives and returns result
```
</details>

### Think About
- Could you use a single channel for both directions? Why or why not?
- Does this function need to close any channels? Why?

---

## Exercise 2: Countdown

### Learning Goal
Prove Generator pattern mastery by reversing the direction.

### Connection to Previous Work
Same pattern as `Generator(n)` from 02_channel_fundamentals, but counts down instead of up.

### Function Signature
```go
// Countdown returns a channel that sends n, n-1, ..., 1, then closes
func Countdown(n int) <-chan int
```

### Examples
```
Countdown(3) sends: 3, 2, 1 (then closes)
Countdown(1) sends: 1 (then closes)
Countdown(0) sends: nothing (closes immediately)
```

### Visual Diagram
```
    main                    goroutine
      │                         │
      │◄─── return ch ──────────│
      │                         │
      │         ┌───────────────┤
      │         │ for i := n; i >= 1; i--
      │◄── n ───┤               │
      │◄── n-1 ─┤               │
      │◄── ... ─┤               │
      │◄── 1 ───┤               │
      │         │ close(ch)     │
      │         └───────────────┘
```

### Hints
<details>
<summary>Hint 1 (Basic)</summary>
Same structure as Generator: create channel, start goroutine, return channel.
</details>

<details>
<summary>Hint 2 (Loop)</summary>
The loop goes from n down to 1 (inclusive). Handle n=0 correctly.
</details>

### Think About
- Why does the close go inside the goroutine?
- What happens if you forget to close the channel?
- What happens with Countdown(0)?

---

## Exercise 3: Relay

### Learning Goal
Chain multiple goroutines together, each passing a value to the next.

### Connection to Previous Work
Extends the PingPong concept, but with a clearer linear flow instead of bouncing.

### Function Signature
```go
// Relay passes a value through n goroutines, each adding 1
// Returns the final value (should equal value + stages)
func Relay(value int, stages int) int
```

### Examples
```
Relay(10, 3) → 13  // 10 → 11 → 12 → 13
Relay(0, 5)  → 5   // 0 → 1 → 2 → 3 → 4 → 5
Relay(5, 0)  → 5   // No stages, value passes through unchanged
Relay(5, 1)  → 6   // 5 → 6
```

### Visual Diagram
```
Relay(10, 3):

main ─────► goroutine1 ─────► goroutine2 ─────► goroutine3 ─────► main
     10           11              12               13
        ch[0]         ch[1]           ch[2]           ch[3]
```

### Hints
<details>
<summary>Hint 1 (Basic)</summary>
You need stages+1 channels to connect stages goroutines.
</details>

<details>
<summary>Hint 2 (Pattern)</summary>
Create all channels first, then start all goroutines, then send, then receive.
</details>

<details>
<summary>Hint 3 (Goroutine logic)</summary>
Each goroutine: receive from left channel, add 1, send to right channel.
</details>

### Think About
- What happens if stages is 0?
- Do you need to close channels here? Why or why not?
- How would you handle errors in a relay?

---

## Exercise 4: FanIn

### Learning Goal
Multiple producers sending to a **single shared channel** concurrently (fan-in pattern).

### Connection to Previous Work
Prepares you for worker pools and concurrent aggregation patterns.

### Function Signature
```go
// FanIn launches one goroutine per value, each sends to a shared channel
// Returns the sum of all values
func FanIn(values []int) int
```

### Expected Pattern

**IMPORTANT:** This exercise demonstrates the "many-to-one" pattern:
- Create **ONE** shared channel
- Launch N goroutines that ALL send to this same channel
- Main receives exactly N values (order is non-deterministic!)
- Sum all received values

This is **different from Relay** (which chains goroutines sequentially).

### Examples
```
FanIn([]int{1, 2, 3})     → 6
FanIn([]int{10})          → 10
FanIn([]int{})            → 0
FanIn([]int{-1, 0, 1})    → 0
```

### Visual Diagram
```
FanIn([]int{1, 2, 3}):

goroutine1 ────┐
    (sends 1)  │
               │
goroutine2 ────┼────► shared ch ────► main (receives N times, sums)
    (sends 2)  │
               │
goroutine3 ────┘
    (sends 3)

Key: All goroutines send to the SAME channel!
     Order of arrival is NON-DETERMINISTIC.
```

### Hints
<details>
<summary>Hint 1 (Structure)</summary>

```go
ch := make(chan int)  // ONE shared channel

for _, v := range values {
    go func(val int) {
        ch <- val  // All send to same channel
    }(v)
}

// Main receives len(values) times
```
</details>

<details>
<summary>Hint 2 (Counting)</summary>
You know how many values to expect: len(values). No need to close the channel!
</details>

<details>
<summary>Hint 3 (Closure trap)</summary>
Remember to pass `v` as a parameter to the goroutine, not capture it directly:
```go
go func(val int) { ch <- val }(v)  // ✅ Correct
go func() { ch <- v }()            // ❌ Bug: all goroutines share same v
```
</details>

### Think About
- In what order will values arrive at main? Is it deterministic?
- Why is fan-in with multiple senders to one channel safe?
- How would you close the channel properly if you didn't know len(values)?

---

## Exercise 5: Ticker

### Learning Goal
Time-based channel sending, tested with Go 1.25's `testing/synctest`.

### Connection to Previous Work
Generator pattern + time, introduces fake time testing.

### Function Signature
```go
// Ticker returns a channel that sends 0, 1, 2, ..., n-1
// Each send is separated by the given interval
func Ticker(n int, interval time.Duration) <-chan int
```

### Examples
```
Ticker(3, time.Second) sends: 0 (immediately), 1 (after 1s), 2 (after 2s), closes
Ticker(0, time.Second) sends: nothing (closes immediately)
```

### Visual Diagram
```
Ticker(3, 1s):

    time     0s      1s      2s      3s
              │       │       │       │
goroutine:   send 0  send 1  send 2  close
              │       │       │       │
              ▼       ▼       ▼       ▼
    ch:      ──0──────1───────2───────X (closed)
```

### About synctest

Go 1.25 introduced `testing/synctest` for testing time-dependent code:

```go
import "testing/synctest"

func TestTicker(t *testing.T) {
    synctest.Test(t, func(t *testing.T) {
        ch := Ticker(3, time.Second)

        // Inside synctest.Test, time is fake!
        // time.Sleep and time.After complete instantly
        // when there's nothing else to do

        for i := range 3 {
            if got := <-ch; got != i {
                t.Errorf("got %d, want %d", got, i)
            }
        }
    })
}
```

### Hints
<details>
<summary>Hint 1 (Basic)</summary>
Generator pattern, but with time.Sleep or time.After between sends.
</details>

<details>
<summary>Hint 2 (First send)</summary>
The first value (0) is sent immediately, not after waiting.
</details>

<details>
<summary>Hint 3 (Edge case)</summary>
When n=0, just close the channel immediately.
</details>

### Think About
- Why is fake time testing valuable?
- How would you test this without synctest?
- What happens if the receiver is slow?

---

## Exercise 6: Pipeline

### Learning Goal
Compose multiple channel stages into a processing pipeline with 5 stages.

### Connection to Previous Work
Combines Generator, filtering, transformation, and aggregation - the capstone of channel reinforcement.

### Function Signature
```go
// Pipeline: Generate → FilterEven → Double → AddOne → FilterOdd → Sum
func Pipeline(n int) int
```

### Processing Stages
1. **Generate**: 1, 2, 3, ..., n
2. **FilterEven**: keep only even numbers (removes odd numbers)
3. **Double**: x → x * 2
4. **AddOne**: x → x + 1
5. **FilterOdd**: keep only odd numbers (after +1, all are odd, so all pass)
6. **Sum**: add all values

### Examples
```
Pipeline(5):
  Generate:    1, 2, 3, 4, 5
  FilterEven:  2, 4           (removes 1, 3, 5)
  Double:      4, 8
  AddOne:      5, 9
  FilterOdd:   5, 9           (all pass - all are odd)
  Sum:         14

Pipeline(10):
  Generate:    1, 2, 3, 4, 5, 6, 7, 8, 9, 10
  FilterEven:  2, 4, 6, 8, 10
  Double:      4, 8, 12, 16, 20
  AddOne:      5, 9, 13, 17, 21
  FilterOdd:   5, 9, 13, 17, 21
  Sum:         65

Pipeline(1):
  Generate:    1
  FilterEven:  (empty - 1 is odd)
  Sum:         0

Pipeline(0) → 0 (no values generated)
```

### Visual Diagram
```
Pipeline(5):

Generate → FilterEven → Double → AddOne → FilterOdd → Sum
   │           │          │         │          │        │
 1,2,3,4,5    2,4        4,8       5,9        5,9      14
   │           │          │         │          │        │
 close      close      close     close      close    return
```

### Hints
<details>
<summary>Hint 1 (Basic)</summary>
Each stage is a function that takes an input channel and returns an output channel.
You need 5 helper functions: generate, filterEven, double, addOne, filterOdd.
</details>

<details>
<summary>Hint 2 (Structure)</summary>

```go
func Pipeline(n int) int {
    ch1 := generate(n)
    ch2 := filterEven(ch1)
    ch3 := double(ch2)
    ch4 := addOne(ch3)
    ch5 := filterOdd(ch4)
    return sum(ch5)
}
```
</details>

<details>
<summary>Hint 3 (Stage pattern)</summary>

```go
// Transform stage (double, addOne)
func double(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for v := range in {
            out <- v * 2
        }
        close(out)
    }()
    return out
}

// Filter stage (filterEven, filterOdd)
func filterEven(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for v := range in {
            if v%2 == 0 {
                out <- v
            }
        }
        close(out)
    }()
    return out
}
```
</details>

### Think About
- Why does each stage need to close its output channel?
- What's the relationship between `range` and channel closing?
- How would you add error handling to a pipeline?

---

## Success Criteria

After completing these exercises, you should be able to:

- [ ] Explain why `close()` goes inside the goroutine
- [ ] Draw channel flow diagrams for any concurrent function
- [ ] Write Generator pattern variants without reference
- [ ] Coordinate N goroutines with channels
- [ ] Use synctest for time-dependent tests
- [ ] Compose channel stages into pipelines

## Running Tests

```bash
cd 09-concurrency/01.5_channel_reinforcement
go test -v                    # Run all tests
go test -v -run TestEcho      # Run specific test
go test -bench=.              # Run benchmarks
```

## Estimated Time

- Exercise 1 (Echo): 10-15 minutes
- Exercise 2 (Countdown): 10 minutes
- Exercise 3 (Relay): 20-30 minutes
- Exercise 4 (FanIn): 15-20 minutes
- Exercise 5 (Ticker): 20-30 minutes
- Exercise 6 (Pipeline): 25-35 minutes

**Total: 2-3 hours**
