# Exercise 02: Channel Fundamentals

**Learning Goal:** Master channel creation, send/receive operations, and the blocking nature of unbuffered channels

**Difficulty:** Tier 1 - Introduction
**Estimated Time:** 30-35 minutes

---

## Problem Description

Channels are Go's primary mechanism for communication between goroutines. They're typed conduits: a `chan int` carries integers, a `chan string` carries strings.

Unbuffered channels block the sender until a receiver is ready, and vice versa. This creates a synchronization point - the sender and receiver must "meet" at the channel.

In this exercise, you'll build functions that use channels for goroutine communication, learning the send (`<-`), receive (`<-`), close, and range operations.

---

## Function Signatures

```go
// SendReceive demonstrates basic channel communication
// Creates a channel, sends value in a goroutine, returns the received value
func SendReceive(value int) int

// Generator returns a channel that will receive the numbers 1 to n, then close
func Generator(n int) <-chan int

// Sum receives all values from the channel and returns their sum
func Sum(ch <-chan int) int

// Ping sends a message to the returned channel (send-only return type)
func Ping(msg string) <-chan string

// PingPong bounces a counter between two goroutines n times
// Returns the final counter value
func PingPong(n int) int
```

---

## Examples

### SendReceive
```
Input:  SendReceive(42)
Output: 42

// Value sent in goroutine is received in main
```

### Generator
```
Input:  Generator(5)
Output: <-chan int that produces: 1, 2, 3, 4, 5 (then closes)

// Example usage:
for v := range Generator(5) {
    fmt.Println(v)  // prints 1, 2, 3, 4, 5
}
```

### Sum
```
Input:  Sum(Generator(5))
Output: 15  // 1+2+3+4+5

// Sums all values from the channel
```

### Ping
```
Input:  Ping("hello")
Output: <-chan string containing "hello"

// Returned channel is receive-only (caller can't send)
```

### PingPong
```
Input:  PingPong(3)
Output: 6

// Counter starts at 0
// Ping adds 1, sends to Pong (counter=1)
// Pong adds 1, sends to Ping (counter=2)
// Ping adds 1, sends to Pong (counter=3)
// ... continues for n round trips
// Final value after 3 round trips: 6
```

---

## Instructions

1. Implement `SendReceive` - create channel, spawn sender goroutine, return received value
2. Implement `Generator` - return a channel, spawn goroutine that sends 1 to n then closes
3. Implement `Sum` - use `range` to receive all values until channel closes
4. Implement `Ping` - send a message to a channel, return receive-only channel
5. Implement `PingPong` - two goroutines passing a counter back and forth
6. Run tests with `go test -v`

---

## Hints

### Basic
- Create channels with `make(chan Type)`
- Send: `ch <- value`, Receive: `value := <-ch`
- Close: `close(ch)` - signals no more values
- Range: `for v := range ch` - receives until channel closes

### Intermediate
- `<-chan T` is a receive-only channel (can only read from it)
- `chan<- T` is a send-only channel (can only write to it)
- These direction types enforce correct usage at compile time
- When returning a channel, spawn a goroutine to populate it

### Solution Pattern
```go
func Generator(n int) <-chan int {
    ch := make(chan int)
    go func() {
        // send values
        close(ch)  // important: signals end of data
    }()
    return ch
}
```

---

## Think About

1. What happens if you try to receive from a closed channel?
2. What happens if you try to send to a closed channel?
3. Why do we return `<-chan T` instead of `chan T` from Generator?
4. Why does PingPong need two channels, not one?

---

## What This Teaches

- **Channel creation** - `make(chan T)` creates an unbuffered channel
- **Send and receive** - The `<-` operator for bidirectional communication
- **Blocking behavior** - Unbuffered channels synchronize sender and receiver
- **Closing channels** - Signals no more values; required for `range` to terminate
- **Channel direction** - `<-chan` and `chan<-` for type safety
- **Generator pattern** - Spawning a goroutine that populates a returned channel
- **Range over channel** - Idiomatic way to consume all values until close
