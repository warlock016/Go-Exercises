# Delve Debugger Guide: Interactive Go Debugging

## Table of Contents

1. [What is Delve?](#what-is-delve)
2. [Installation](#installation)
3. [Starting Delve](#starting-delve)
4. [Essential Commands](#essential-commands)
5. [Breakpoints](#breakpoints)
6. [Stepping Through Code](#stepping-through-code)
7. [Inspecting Variables](#inspecting-variables)
8. [Working with Pointers](#working-with-pointers)
9. [Debugging Tests](#debugging-tests)
10. [Goroutine Debugging](#goroutine-debugging)
11. [Practical Examples](#practical-examples)
12. [Common Workflows](#common-workflows)
13. [Tips and Tricks](#tips-and-tricks)
14. [Troubleshooting](#troubleshooting)

---

## What is Delve?

Delve (`dlv`) is a debugger for Go programs. Unlike print debugging, Delve lets you:

- **Pause execution** at any line (breakpoints)
- **Step through code** line by line
- **Inspect variables** including complex types (slices, maps, structs, pointers)
- **Modify variables** during execution
- **Navigate the call stack** to see how you got there
- **Debug goroutines** individually

### When to Use Delve vs Printf

| Use Printf When... | Use Delve When... |
|-------------------|-------------------|
| Quick sanity check | Segfault/panic with unclear cause |
| Logging control flow | Need to inspect complex data structures |
| Simple variable values | Pointer/nil issues |
| CI/automated testing | Interactive exploration needed |
| Remote debugging not possible | Understanding unfamiliar code |

---

## Installation

### macOS / Linux

```bash
# Using go install (recommended)
go install github.com/go-delve/delve/cmd/dlv@latest

# Verify installation
dlv version
```

### Verify PATH

Delve installs to `$GOPATH/bin` (usually `~/go/bin`). Ensure it's in your PATH:

```bash
# Add to ~/.zshrc or ~/.bashrc
export PATH=$PATH:$(go env GOPATH)/bin

# Reload shell
source ~/.zshrc
```

### macOS Code Signing (if needed)

On macOS, you may need to allow Delve to control processes:

```bash
# If you get "could not launch process: debugserver" error
sudo /usr/sbin/DevToolsSecurity -enable
```

---

## Starting Delve

### Debug a Package/Main Program

```bash
# From package directory
dlv debug

# Or specify package path
dlv debug ./cmd/myapp

# With arguments
dlv debug ./cmd/myapp -- --config=dev.yaml
```

### Debug Tests

```bash
# Debug all tests in current package
dlv test

# Debug specific test file
dlv test -- -test.run TestSpecificFunction

# Debug tests in specific package
dlv test ./pkg/mypackage
```

### Attach to Running Process

```bash
# Find process ID
ps aux | grep myprogram

# Attach
dlv attach <pid>
```

### Debug Core Dump

```bash
dlv core ./myprogram core.12345
```

---

## Essential Commands

Once inside Delve (`(dlv)` prompt), these are your core commands:

### Navigation

| Command | Shortcut | Description |
|---------|----------|-------------|
| `continue` | `c` | Run until next breakpoint |
| `next` | `n` | Execute current line, step OVER function calls |
| `step` | `s` | Execute current line, step INTO function calls |
| `stepout` | `so` | Run until current function returns |
| `restart` | `r` | Restart the program |

### Breakpoints

| Command | Shortcut | Description |
|---------|----------|-------------|
| `break <location>` | `b` | Set breakpoint |
| `breakpoints` | `bp` | List all breakpoints |
| `clear <id>` | | Remove breakpoint by ID |
| `clearall` | | Remove all breakpoints |
| `condition <id> <expr>` | `cond` | Set conditional breakpoint |

### Inspection

| Command | Shortcut | Description |
|---------|----------|-------------|
| `print <expr>` | `p` | Print variable or expression |
| `locals` | | Show all local variables |
| `args` | | Show function arguments |
| `whatis <expr>` | | Show type of expression |
| `display <expr>` | | Auto-print on each step |

### Context

| Command | Shortcut | Description |
|---------|----------|-------------|
| `list` | `l` | Show source code around current line |
| `stack` | `bt` | Show call stack (backtrace) |
| `frame <n>` | | Switch to stack frame n |
| `goroutines` | `grs` | List all goroutines |
| `goroutine <id>` | `gr` | Switch to goroutine |

### Session

| Command | Description |
|---------|-------------|
| `help` | Show all commands |
| `help <cmd>` | Help for specific command |
| `exit` / `quit` | Exit Delve |

---

## Breakpoints

### Setting Breakpoints

```bash
# By function name
(dlv) break main.main
(dlv) break mini_database.Set

# By file:line
(dlv) break mini_database.go:70
(dlv) break /full/path/to/file.go:42

# By package.function
(dlv) break github.com/user/pkg.Function

# At current location (when paused)
(dlv) break
```

### Listing Breakpoints

```bash
(dlv) breakpoints
Breakpoint 1 at 0x10a5f20 for mini_database.Set() ./mini_database.go:57
Breakpoint 2 at 0x10a6230 for mini_database.Rollback() ./mini_database.go:255
```

### Conditional Breakpoints

Only trigger when condition is true:

```bash
# Set breakpoint
(dlv) break mini_database.go:70
Breakpoint 1 set at 0x10a5f80

# Add condition
(dlv) condition 1 key == "user:123"

# Now breakpoint only triggers when key equals "user:123"
```

### Removing Breakpoints

```bash
# Remove specific breakpoint
(dlv) clear 1

# Remove all breakpoints
(dlv) clearall
```

### Breakpoint on Panic

```bash
# Break when panic occurs
(dlv) break runtime.gopanic
```

---

## Stepping Through Code

### Understanding Step Commands

```go
func outer() {
    x := 1           // Line 1
    inner()          // Line 2
    y := 2           // Line 3
}

func inner() {
    z := 3           // Line A
}
```

At Line 2:
- `next` → Goes to Line 3 (steps OVER inner())
- `step` → Goes to Line A (steps INTO inner())
- `stepout` → After entering inner(), returns to Line 3

### Practical Example

```bash
(dlv) break mini_database.Set
(dlv) continue
> mini_database.Set() ./mini_database.go:57

(dlv) next    # Execute line 57, go to 58
(dlv) next    # Execute line 58, go to 59
(dlv) step    # If line has function call, enter it
(dlv) stepout # Exit current function, return to caller
```

### Step Count

Execute multiple steps at once:

```bash
(dlv) next 5    # Execute 5 lines
(dlv) step 3    # Step into, 3 times
```

---

## Inspecting Variables

### Basic Printing

```bash
# Print simple variable
(dlv) print key
"user:123"

# Print struct
(dlv) print newOp
mini_database.Operation {
    OpType: "update",
    Key: "user:123",
    OldRecord: *mini_database.Record nil,
}

# Print with format
(dlv) print fmt.Sprintf("key=%s, value=%s", key, value)
```

### Printing Pointers

```bash
# Print pointer value (the address)
(dlv) print newOp.OldRecord
*mini_database.Record nil

# Print what pointer points to (dereference)
(dlv) print *newOp.OldRecord
# Error: nil pointer dereference (if nil)

# Check if pointer is nil
(dlv) print newOp.OldRecord == nil
true
```

### Maps and Slices

```bash
# Print entire map
(dlv) print db.Data
map[string]mini_database.Record [
    "user:1": {...},
    "user:2": {...},
]

# Print map element
(dlv) print db.Data["user:1"]

# Print slice
(dlv) print db.transactionLog

# Print slice element
(dlv) print db.transactionLog[0]

# Print slice length
(dlv) print len(db.transactionLog)
```

### Nested Structures

```bash
# Access nested fields
(dlv) print db.stats.RecordCount
5

# Access through pointer
(dlv) print db.stats
*mini_database.Stats {
    RecordCount: 5,
    Reads: 10,
    ...
}
```

### Show All Local Variables

```bash
(dlv) locals
key = "user:123"
value = "Alice"
newRec = mini_database.Record {Key: "user:123", Value: "Alice", ...}
newOp = mini_database.Operation {OpType: "", Key: "user:123", OldRecord: nil}
```

### Show Function Arguments

```bash
(dlv) args
db = (*mini_database.Database)(0xc0000a4000)
key = "user:123"
value = "Alice"
```

### Type Information

```bash
(dlv) whatis newOp.OldRecord
*mini_database.Record

(dlv) whatis db.Data
map[string]mini_database.Record
```

---

## Working with Pointers

Pointer bugs (like your nil dereference) are common. Here's how to debug them:

### Check if Pointer is Nil

```bash
(dlv) print ptr
*SomeType nil

(dlv) print ptr == nil
true
```

### Inspect Valid Pointer

```bash
# Pointer holds an address
(dlv) print ptr
*mini_database.Record 0xc0000a4080

# Dereference to see contents
(dlv) print *ptr
mini_database.Record {
    Key: "user:1",
    Value: "Alice",
    ...
}
```

### Common Nil Pointer Pattern

Your bug was:

```go
newOp := Operation{
    Key: key,
    // OldRecord not initialized → nil
}
*newOp.OldRecord = db.Data[key]  // CRASH: dereferencing nil
```

Debug session:

```bash
(dlv) break mini_database.go:70
(dlv) continue
> mini_database.go:70

(dlv) print newOp
mini_database.Operation {
    OpType: "",
    Key: "user:1",
    OldRecord: *mini_database.Record nil,   # ← nil!
}

(dlv) print newOp.OldRecord == nil
true

# Now you see the problem: OldRecord is nil before dereference
```

---

## Debugging Tests

### Run Specific Test

```bash
# Debug single test
dlv test -- -test.run TestRollback

# Debug test with verbose output
dlv test -- -test.v -test.run TestRollback
```

### Full Test Debug Session

```bash
$ cd /path/to/mini_database
$ dlv test -- -test.run TestTransactionRollback

Type 'help' for list of commands.
(dlv) break Rollback
Breakpoint 1 set at 0x10a7340

(dlv) continue
> mini_database.Rollback() ./mini_database.go:255

(dlv) print db.transactionLog
[]mini_database.Operation len: 2, cap: 2, [
    {OpType: "create", Key: "key1", OldRecord: nil},
    {OpType: "update", Key: "key2", OldRecord: *{Key: "key2", ...}},
]

(dlv) next
(dlv) print i
1

(dlv) print db.transactionLog[i].OpType
"update"
```

### Break on Test Failure

```bash
# Break when test calls t.Error/t.Fail
(dlv) break testing.(*common).Error
(dlv) break testing.(*common).Fail
```

---

## Goroutine Debugging

### List All Goroutines

```bash
(dlv) goroutines
* Goroutine 1 - User: ./main.go:15 main.main (0x10a5f20)
  Goroutine 2 - User: runtime/proc.go:367 runtime.gopark (0x1038d40)
  Goroutine 3 - User: ./worker.go:42 main.worker (0x10a6100)
```

### Switch Goroutine

```bash
(dlv) goroutine 3
Switched to goroutine 3

(dlv) stack
0  main.worker() ./worker.go:42
1  runtime.goexit() runtime/asm_arm64.s:1222
```

### Set Breakpoint in Specific Goroutine

```bash
# Break only in goroutine 3
(dlv) break worker.go:50
(dlv) condition 1 runtime.curg.goid == 3
```

---

## Practical Examples

### Example 1: Debugging Your Nil Pointer Bug

```bash
$ cd "/Users/mode/Documents/Code/Go Exercises/02-data-structures/17_capstone_mini_database"
$ dlv test -- -test.run TestTransaction

(dlv) break mini_database.go:70
Breakpoint 1 set at 0x10a5f80

(dlv) continue
> mini_database.go:70: *newOp.OldRecord = db.Data[key]

# Inspect the operation struct BEFORE the crash
(dlv) print newOp
mini_database.Operation {
    OpType: "",
    Key: "user:1",
    OldRecord: *mini_database.Record nil,  # ← THE BUG
}

# Confirm OldRecord is nil
(dlv) print newOp.OldRecord == nil
true

# See what we're trying to assign
(dlv) print db.Data[key]
mini_database.Record {
    Key: "user:1",
    Value: "Alice",
    CreatedAt: 1699876543,
    UpdatedAt: 0,
}

# The problem: we're trying to dereference nil to assign to it
```

### Example 2: Debugging a Loop

```bash
(dlv) break mini_database.go:262
(dlv) continue

# First iteration
(dlv) print i
2

(dlv) print db.transactionLog[i]
{OpType: "update", Key: "key3", OldRecord: *{...}}

(dlv) next
(dlv) print i
1

# Watch variable across iterations
(dlv) display i
(dlv) display db.transactionLog[i].OpType

(dlv) continue  # Runs to next breakpoint hit
i = 1
db.transactionLog[i].OpType = "create"
```

### Example 3: Examining the Call Stack

```bash
(dlv) break mini_database.go:70
(dlv) continue

# See how we got here
(dlv) stack
0  mini_database.(*Database).Set() ./mini_database.go:70
1  mini_database.TestTransactionRollback() ./mini_database_test.go:245
2  testing.tRunner() testing/testing.go:1576

# Move up the stack to see test code
(dlv) frame 1
(dlv) list
   243: func TestTransactionRollback(t *testing.T) {
   244:     db := NewDatabase()
   245:     db.Set("key1", "value1")  // ← We're in this call
   246:     db.Begin()
```

---

## Common Workflows

### Workflow 1: Finding Where a Panic Occurs

```bash
$ dlv test

(dlv) break runtime.gopanic
(dlv) continue

# Program runs until panic
> runtime.gopanic() runtime/panic.go:1188

# Look at call stack
(dlv) stack
0  runtime.gopanic() runtime/panic.go:1188
1  runtime.panicmem() runtime/panic.go:261
2  runtime.sigpanic() runtime/signal_unix.go:881
3  mini_database.(*Database).Set() ./mini_database.go:70  # ← HERE
4  mini_database.TestTransaction() ./mini_database_test.go:102

# Go to the frame where our code panicked
(dlv) frame 3
(dlv) list
   68:     if _, exists := db.Data[key]; exists {
   69:         // BUG: OldRecord is nil
>  70:         *newOp.OldRecord = db.Data[key]

(dlv) locals
# See all variable values at crash point
```

### Workflow 2: Understanding Data Flow

```bash
(dlv) break mini_database.Set
(dlv) continue

# See what's passed in
(dlv) args
db = (*mini_database.Database)(0xc0000a4000)
key = "user:1"
value = "Alice"

# Step through, watching state changes
(dlv) display db.Data
(dlv) display newOp

(dlv) next
(dlv) next  # Watch display output update
```

### Workflow 3: Comparing Before/After State

```bash
(dlv) break mini_database.Rollback
(dlv) continue

# Capture state before rollback
(dlv) print db.Data
map[string]mini_database.Record [
    "key1": {Key: "key1", Value: "modified"},
    "key2": {Key: "key2", Value: "new"},
]

# Step through rollback
(dlv) next
(dlv) next
...

# Check state after
(dlv) print db.Data
map[string]mini_database.Record [
    "key1": {Key: "key1", Value: "original"},
]
# key2 is gone (was created in transaction), key1 restored
```

---

## Tips and Tricks

### 1. Use Tab Completion

```bash
(dlv) print db.Da<TAB>
(dlv) print db.Data
```

### 2. Command History

Use up/down arrows to navigate previous commands.

### 3. Abbreviations

Most commands have short forms:

```bash
(dlv) b main.go:50    # break
(dlv) c               # continue
(dlv) n               # next
(dlv) s               # step
(dlv) p variable      # print
(dlv) l               # list
(dlv) bt              # stack (backtrace)
```

### 4. Evaluate Expressions

```bash
(dlv) print len(db.Data)
5

(dlv) print db.Data["key1"].Value + " modified"
"Alice modified"

(dlv) print i > 0 && i < len(slice)
true
```

### 5. Set Variables (Modify During Debug)

```bash
(dlv) set i = 0
(dlv) set newOp.OpType = "test"
```

### 6. Source Code Context

```bash
# Show more lines around current position
(dlv) list mini_database.go:60:80   # Lines 60-80

# Show current function
(dlv) list .                         # Current location
```

### 7. Watchpoints (Break on Variable Change)

```bash
# Note: Watchpoints have limited support in dlv
# Use conditional breakpoints instead:
(dlv) break mini_database.go:100
(dlv) condition 1 db.stats.RecordCount > 5
```

### 8. Continue to Specific Location

```bash
# Run until you hit line 100
(dlv) break mini_database.go:100
(dlv) continue
```

---

## Troubleshooting

### "could not launch process: fork/exec: operation not permitted"

**macOS:** Enable developer tools:

```bash
sudo /usr/sbin/DevToolsSecurity -enable
```

### "could not attach to pid: this is a build with PIE"

Build without PIE for debugging:

```bash
go build -gcflags="all=-N -l" -o myapp .
dlv exec ./myapp
```

### "could not find symbol value for..."

Symbols were optimized away. Disable optimizations:

```bash
dlv test -gcflags="all=-N -l"
```

### Breakpoint Not Hitting

1. Check file path is correct: `break ./pkg/file.go:50`
2. Ensure code path is executed (not in dead branch)
3. Check breakpoint is set: `breakpoints`

### Variables Show as "optimized away"

Disable compiler optimizations:

```bash
dlv test -gcflags="all=-N -l" -- -test.run TestMyFunction
```

### Slow Debugging

Delve compiles with debug symbols. For large projects:

```bash
# Pre-build test binary
go test -c -gcflags="all=-N -l" -o test.exe

# Debug the binary
dlv exec ./test.exe -- -test.run TestSpecific
```

---

## Quick Reference Card

```
STARTING
  dlv test                    Debug tests
  dlv test -- -test.run X     Debug specific test
  dlv debug                   Debug main package

BREAKPOINTS
  b <loc>                     Set breakpoint
  bp                          List breakpoints
  clear <id>                  Remove breakpoint
  cond <id> <expr>            Conditional breakpoint

EXECUTION
  c                           Continue
  n                           Next (step over)
  s                           Step (step into)
  so                          Step out
  r                           Restart

INSPECTION
  p <expr>                    Print expression
  locals                      Show local variables
  args                        Show function arguments
  whatis <var>                Show type
  display <expr>              Auto-print on each step

NAVIGATION
  l                           List source
  bt / stack                  Show call stack
  frame <n>                   Go to stack frame
  grs                         List goroutines
  gr <id>                     Switch goroutine

SESSION
  help                        Show help
  exit / quit                 Exit delve
```

---

## Further Reading

- [Delve Documentation](https://github.com/go-delve/delve/tree/master/Documentation)
- [Delve CLI Commands](https://github.com/go-delve/delve/tree/master/Documentation/cli)
- [VS Code Go Debugging](https://github.com/golang/vscode-go/wiki/debugging) (uses Delve)

---

**Version:** 1.0
**Last Updated:** 2025-12-02
**Related:** DEBUGGING_ALGORITHMS.md, VALUES_AND_REFERENCES.md
