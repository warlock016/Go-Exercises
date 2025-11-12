# Getting Started with Your Go Learning Journey

Welcome! This guide will help you begin your structured Go learning journey.

## First Steps (Do These Now!)

### 1. Verify Your Go Installation
```bash
go version
```
You should see Go 1.21 or later. If not, install from https://go.dev/dl/

### 2. Navigate to Your Learning Directory
```bash
cd "/Users/mode/Documents/Code/Go Exercises"
```

### 3. Read the Main README
```bash
cat README.md
```
This explains the overall structure and philosophy.

### 4. Start Your Diagnostic Assessment
```bash
cd 00-diagnostic
cat README.md
```

## Understanding the Learning System

### Directory Structure
```
Go Exercises/
├── README.md              # Start here - explains everything
├── PROGRESS.md            # Your skill metrics and advancement
├── JOURNAL.md             # Daily reflections (use this!)
├── 00-diagnostic/         # → START HERE (30-60 min assessment)
├── 01-fundamentals/       # Variables, types, control flow
├── 02-data-structures/    # Slices, maps, structs
├── 03-functions-methods/  # Functions, methods, receivers
├── 04-error-handling/     # Go's error handling patterns
├── 05-testing/            # Writing tests, TDD
├── 06-interfaces/         # Interface design
├── 07-packages-modules/   # Code organization
├── 08-http-apis/          # Web servers and APIs
├── 09-concurrency/        # Goroutines, channels
├── 10-databases/          # SQL integration
├── 11-performance/        # Profiling and optimization
├── projects/              # Milestone capstone projects
└── resources/             # Helpful guides and references
```

### How Exercises Work

Each exercise has:
- **README.md** - Problem description, hints, examples
- **[name].go** - File with TODO(human) where you implement
- **[name]_test.go** - Tests to verify your solution

### The Learning Cycle

1. **Read the README** - Understand the problem
2. **Implement the function** - Edit the .go file
3. **Run tests** - `go test -v`
4. **Debug if needed** - Read errors, fix code
5. **Repeat** - Until tests pass
6. **Reflect** - Update your JOURNAL.md

## Daily Learning Routine (4 hours)

### Hour 1: Review & Warm-up (9:00-10:00)
- Review yesterday's concepts
- Do a spaced repetition exercise
- Update your JOURNAL.md with yesterday's reflection

### Hour 2-3: New Learning (10:00-12:00)
- Read exercise README carefully
- Attempt implementation
- Run tests frequently
- Debug systematically
- Don't rush - understanding > speed

### Hour 4: Independent Practice (12:00-1:00)
- Additional exercises
- Write code explanations
- Experiment with variations
- Update PROGRESS.md

### Break! (Take real breaks)

## Anti-AI-Reliance Guidelines (Weeks 1-4)

### ✅ Allowed
- Reading Go documentation (go.dev)
- Searching for "how to X in Go" (reading articles)
- Reading example code to understand patterns
- Using the Go compiler errors to guide you
- Asking "what does this error mean?"

### ❌ Not Allowed (for now)
- Asking AI to "write a function that..."
- Copying code without understanding it
- Having AI debug your code
- AI-generated solutions

### Why?
You're building **foundational skills**. Using AI now is like using a calculator before learning arithmetic - it prevents genuine understanding.

## Tips for Success

### 1. Read Error Messages Carefully
Go's compiler is extremely helpful. The error message usually tells you:
- Exactly what's wrong
- Which line has the problem
- Often suggests the fix

Example:
```
./temperature.go:7:2: undefined: result
```
This tells you: file (temperature.go), line (7), column (2), problem (undefined variable)

### 2. Start Simple, Then Improve
Don't try to write perfect code immediately:
1. Get something working (even if hacky)
2. Make sure it passes tests
3. Then refactor and improve

### 3. Test Frequently
Don't write 50 lines then test. Write 3-5 lines, test, repeat.
```bash
go test -v  # Verbose output
go test     # Quiet (just pass/fail)
```

### 4. When Stuck (After 15-20 minutes)
1. Read the error message again (carefully!)
2. Check the exercise README for hints
3. Look at the test file to understand what's expected
4. Search Go.dev for the specific concept
5. Take a 5-minute break
6. Try explaining the problem out loud

### 5. Use Your Journal
After each session, spend 5 minutes writing:
- What did I learn?
- What was hard?
- What clicked?
- What do I need to review?

This reflection dramatically improves retention.

## Common Beginner Mistakes

### 1. Not Reading the Error
**Bad:** See error → panic → change random things
**Good:** Read error → understand problem → make targeted fix

### 2. Copying Without Understanding
**Bad:** Copy code from Stack Overflow → "it works!" → move on
**Good:** Copy code → understand each line → explain in comments → then move on

### 3. Skipping Tests
**Bad:** Write function → "looks good!" → move on
**Good:** Write function → run tests → fix until green → understand why

### 4. Rushing
**Bad:** Complete as many exercises as possible
**Good:** Deeply understand each concept before moving on

### 5. Not Asking "Why?"
**Bad:** "This is the syntax, memorize it"
**Good:** "Why is it designed this way? What problem does it solve?"

## Keyboard Shortcuts for Efficiency

### Terminal
- `Ctrl+C` - Cancel running program
- `Ctrl+L` - Clear screen
- `Up Arrow` - Previous command
- `Ctrl+R` - Search command history

### Go Commands
```bash
go test -v              # Run tests (verbose)
go test -v -run TestName # Run specific test
go run main.go          # Run a program
go build                # Compile package
go fmt                  # Format code
go doc package.Function # View documentation
```

## First Day Checklist

- [ ] Verified Go installation (`go version`)
- [ ] Read main README.md
- [ ] Read this getting started guide
- [ ] Opened JOURNAL.md and made first entry
- [ ] Started diagnostic assessment (00-diagnostic/)
- [ ] Completed at least 3 diagnostic exercises
- [ ] Ran `go test -v` successfully
- [ ] Updated JOURNAL.md with reflections

## Resources You'll Need

### Official Documentation
- **go.dev/doc** - Official Go documentation
- **go.dev/tour** - Interactive tour (good refresher)
- **pkg.go.dev** - Package documentation

### Recommended Reading Order
1. Effective Go: https://go.dev/doc/effective_go
2. Go by Example: https://gobyexample.com
3. Go standard library: https://pkg.go.dev/std

### When You Need Help
1. Read the error message
2. Check exercise hints
3. Search go.dev documentation
4. Search "golang [your question]"
5. After genuine attempt (20min+), ask for guidance

## What Success Looks Like

### After Week 1
- Completed diagnostic assessment
- Understanding of basic syntax
- Can write simple functions independently
- Comfortable reading Go documentation
- Beginning to think in Go patterns

### After Week 4
- Completed first capstone project (CLI tool)
- Strong fundamentals (variables, loops, functions)
- Good grasp of slices, maps, structs
- Writing tests for your code
- Debugging efficiently

### After Week 8
- Completed REST API project
- Confident with error handling
- Understanding interfaces
- Writing idiomatic Go code
- Learning from your mistakes quickly

## Your First Task

**Right now, do this:**

1. Open your JOURNAL.md:
```bash
nano JOURNAL.md  # or use your preferred editor
```

2. Fill in today's entry (Day 1):
- What are you hoping to learn?
- Why are you learning Go?
- What are your goals?
- How are you feeling about this journey?

3. Begin the diagnostic:
```bash
cd 00-diagnostic
cat README.md
```

---

**Remember:** This is a marathon, not a sprint. You have 20 hours/week for focused learning. Quality of practice matters far more than quantity. Every expert was once a beginner who refused to give up.

Let's begin! 🚀
