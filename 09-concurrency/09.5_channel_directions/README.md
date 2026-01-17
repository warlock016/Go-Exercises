# Exercise 09.5: Channel Directions

## Learning Goal

Master channel direction types (`<-chan T`, `chan<- T`) by implementing common concurrent patterns from scratch—including designing function signatures.

## Why This Exercise?

Channel directions are compile-time safety features that:
- Prevent accidental closes by receivers
- Document data flow in function signatures
- Enable the compiler to catch ownership bugs

This exercise forces you to **think about** the correct directions rather than just copying signatures.

## The Challenge

Unlike previous exercises, this one provides **no function signatures**. You must:

1. **Design the signatures** - Decide the correct channel direction for each parameter and return type
2. **Implement the functions** - Write the concurrent logic
3. **Write the tests** - Practice testing concurrent code

## Functions to Implement

| Function | Pattern | Description |
|----------|---------|-------------|
| `Numbers` | Generator | Emit integers 1 to n |
| `Square` | Pipeline Stage | Square each value |
| `Filter` | Pipeline Stage | Keep values matching predicate |
| `Sum` | Consumer/Sink | Sum all values (blocking) |
| `Merge` | Fan-In | Combine multiple channels |
| `Tee` | Broadcast | Split one channel into two |

## Key Questions to Answer

Before implementing each function, ask yourself:

1. **Who creates this channel?** → Creator closes it
2. **Who sends to this channel?** → Should be `chan<-` or bidirectional
3. **Who receives from this channel?** → Should be `<-chan` or bidirectional
4. **Should this function block or return immediately?**

## Testing Guidelines

The test file contains detailed comments about what to test. Key points:

- Run with `-race` flag: `go test -race ./...`
- Run multiple times: `go test -race -count=10 ./...`
- `Tee` tests MUST consume both outputs concurrently (or deadlock!)
- Use helper functions to reduce boilerplate

## Success Criteria

- [ ] All functions compile with correct channel direction types
- [ ] All tests pass
- [ ] No race conditions (`go test -race -count=100`)
- [ ] You can explain WHY each direction type was chosen

## Estimated Time

- 30-45 minutes for implementation
- 20-30 minutes for tests
- Total: ~1 hour

## Hints (Progressive)

<details>
<summary>Hint 1: Numbers signature</summary>

The function creates and owns the channel, so it should return...?
Think: can the caller close a channel returned from a generator?

</details>

<details>
<summary>Hint 2: Square/Filter signature</summary>

These accept a channel they don't own (input) and create one they do own (output).
Input: can only receive. Output: return as receive-only.

</details>

<details>
<summary>Hint 3: Tee deadlock</summary>

Tee sends each input value to BOTH outputs. If output channels are unbuffered
and one consumer is slow, Tee blocks. Tests must read from both concurrently.

</details>

<details>
<summary>Hint 4: Merge knowing when all inputs are done</summary>

Use sync.WaitGroup with one goroutine per input channel.
Only close output after wg.Wait() completes.

</details>
