# Exercise 03: Parentheses Validator

**Difficulty:** Medium
**Time:** 30-40 minutes
**Concepts:** Stack data structure, matching pairs, validation patterns

## Learning Goal

Learn to use a stack to track state in validation problems. Stacks are perfect for matching pairs (opening/closing brackets, HTML tags, etc.) because they naturally handle nesting.

## The Problem

Implement a validator that checks if parentheses in a string are balanced and properly nested.

**Valid examples:**
- `"()"` - Single pair
- `"(())"` - Nested pairs
- `"()()"` - Multiple pairs
- `"((()))"` - Deeply nested

**Invalid examples:**
- `"(()"` - Missing closing
- `"())"` - Extra closing
- `")("` - Wrong order
- `"(()))"` - Extra closing at end

## Function Signatures

```go
func IsValid(s string) bool
```

## Examples

```go
IsValid("")
// → true (empty string is valid)

IsValid("()")
// → true

IsValid("(())")
// → true

IsValid("()()")
// → true

IsValid("((()))")
// → true

IsValid("(()")
// → false (missing closing)

IsValid("())")
// → false (extra closing)

IsValid(")(")
// → false (wrong order)
```

## Instructions

1. Create a stack (slice) to track opening parentheses
2. Iterate through each character in the string
3. For opening parenthesis '(':
   - Push onto the stack
4. For closing parenthesis ')':
   - If stack is empty, return false (no matching opening)
   - Otherwise, pop from the stack (matched a pair)
5. After the loop, check if stack is empty
   - Empty stack = all parentheses matched
   - Non-empty stack = some openings never closed

## Hints

### Basic Hint

A **stack** is a Last-In-First-Out (LIFO) data structure. In Go, you can use a slice:

```go
stack := []rune{}

// Push
stack = append(stack, '(')

// Pop
if len(stack) > 0 {
    stack = stack[:len(stack)-1]
}

// Check if empty
isEmpty := len(stack) == 0
```

The key insight: When you see ')' you need to match it with the most recent '(' (which is at the top of the stack).

### Intermediate Hint

Structure your code like this:

```go
func IsValid(s string) bool {
    stack := []rune{}

    for _, char := range s {
        if char == '(' {
            // Push opening parenthesis onto stack
            stack = append(stack, char)
        } else if char == ')' {
            // Check if there's a matching opening
            if len(stack) == 0 {
                return false  // No matching opening
            }
            // Pop from stack (we found a match)
            stack = stack[:len(stack)-1]
        }
        // Ignore other characters
    }

    // Stack should be empty if all pairs matched
    return len(stack) == 0
}
```

### Complete Solution Hint

Here's the full algorithm:

```go
func IsValid(s string) bool {
    stack := []rune{}

    for _, char := range s {
        if char == '(' {
            // Opening: push onto stack
            stack = append(stack, char)
        } else if char == ')' {
            // Closing: check for matching opening
            if len(stack) == 0 {
                return false  // Error: closing without opening
            }
            // Pop the matching opening
            stack = stack[:len(stack)-1]
        }
        // Ignore non-parenthesis characters
    }

    // Valid only if stack is empty (all openings were closed)
    return len(stack) == 0
}
```

**Why this works:**
- Each '(' is pushed, waiting for its matching ')'
- Each ')' pops the most recent unmatched '('
- If we see ')' with empty stack, we have ')' without '('
- If stack is non-empty at end, we have '(' without ')'

## Think About

1. Why does a stack work better than a counter for this problem?
2. What happens if the string contains other characters like "a(b)c"? Should they be valid?
3. How would you extend this to handle multiple bracket types: `()`, `[]`, `{}`?
4. What's the time complexity? Space complexity?

## What This Teaches

- **Stack data structure:** LIFO pattern for matching pairs
- **Validation patterns:** How to check structural correctness
- **Edge case handling:** Empty input, unmatched pairs, wrong order
- **State accumulation:** Stack tracks "unmatched openings"
- **Foundation for parsing:** Brackets/tags work the same way

Ready to implement? Open `parentheses_validator.go` and look for `TODO(human)` markers!
