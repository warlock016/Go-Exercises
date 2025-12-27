# Module 06: Interfaces

**Status:** 🆕 New
**Type:** Core Curriculum
**Estimated Time:** 10-12 hours
**Prerequisites:** Module 01 Fundamentals, Module 02 Structs & Methods

---

## Why This Module?

Interfaces are Go's most powerful feature for abstraction and polymorphism. Unlike other languages, Go uses **implicit satisfaction** - types don't declare they implement an interface, they just do it by having the right methods. This makes interfaces incredibly flexible and is central to Go's philosophy of composition over inheritance.

**Key Insight:** "The bigger the interface, the weaker the abstraction." - Rob Pike

This module teaches you to design small, focused interfaces that enable powerful, testable, composable code.

---

## Learning Objectives

By the end of this module, you will be able to:

✅ **Understand implicit interface satisfaction** - no "implements" keyword needed
✅ **Use standard library interfaces** - io.Reader, io.Writer, fmt.Stringer, error
✅ **Design small, focused interfaces** - following Go proverbs
✅ **Compose interfaces** - combining smaller interfaces into larger ones
✅ **Use type assertions and type switches** - working with interface{}/any
✅ **Write testable code with interfaces** - dependency injection for mocking
✅ **Apply interface segregation principle** - accept interfaces, return structs

---

## Module Structure

### **Tier 1: Introduction (1-3)** - Understanding the Basics (2-3 hours)

Learn what makes Go interfaces unique and how to use them.

| # | Exercise | Concepts | Time |
|---|----------|----------|------|
| 01 | Implicit Satisfaction | Duck typing, no "implements" keyword | 30min |
| 02 | Stringer Interface | fmt.Stringer, custom String() methods | 35min |
| 03 | Error Interface | Custom error types, Error() method | 40min |

**Learning Focus:** Interfaces are satisfied implicitly by having the right methods

---

### **Tier 2: Application (4-7)** - Standard Library Interfaces (5-6 hours)

Master the interfaces you'll use every day in real Go code.

| # | Exercise | Concepts | Time |
|---|----------|----------|------|
| 04 | Reader Basics | io.Reader, Read([]byte) method | 45min |
| 05 | Writer Basics | io.Writer, Write([]byte) method | 45min |
| 05.5 | Wrapper Chains | Chaining Readers, data pipelines | 55min |
| 05.6 | io.Copy Bridge | Connecting Reader to Writer, io.Copy | 50min |
| 06 | Interface Composition | Embedding interfaces | 50min |
| 07 | Type Assertions | Comma-ok idiom, type checking | 50min |

**Learning Focus:** Standard library patterns, wrapper composition, and interface composition

---

### **Tier 3: Integration (8-12)** - Advanced Patterns (3-4 hours)

Apply interfaces to solve real-world design problems.

| # | Exercise | Concepts | Time |
|---|----------|----------|------|
| 08 | Type Switch | switch v.(type) pattern | 50min |
| 09 | Mock for Testing | Dependency injection, test doubles | 60min |
| 10 | Small Interfaces | Single-method interfaces | 55min |
| 11 | Interface Segregation | Accept interfaces, return structs | 60min |
| 12 | Polymorphic Collection | Interface slices, runtime polymorphism | 60min |

**Learning Focus:** Design principles and testing with interfaces

---

## Key Concepts

### Implicit Satisfaction

```go
// NO "implements" keyword in Go!
type Speaker interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

// Dog implicitly satisfies Speaker because it has Speak() method
var s Speaker = Dog{}  // This just works!
```

### Accept Interfaces, Return Structs

```go
// GOOD: Function accepts interface (flexible)
func ProcessData(r io.Reader) error {
    // Works with files, network, strings.Reader, etc.
}

// GOOD: Function returns concrete type (clear)
func NewLogger() *FileLogger {
    return &FileLogger{}
}

// BAD: Unnecessary interface return
func NewLogger() Logger {  // Don't do this unless needed
    return &FileLogger{}
}
```

### Small Interfaces

```go
// GOOD: Small, focused interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// GOOD: Compose when needed
type ReadWriter interface {
    Reader
    Writer
}

// BAD: Large, kitchen-sink interfaces
type DataManager interface {
    Read() error
    Write() error
    Validate() error
    Transform() error
    // ... 10 more methods
}
```

---

## Go Proverbs About Interfaces

1. **"The bigger the interface, the weaker the abstraction."**
   - Keep interfaces small and focused

2. **"Don't design with interfaces, discover them."**
   - Write code first, extract interfaces when you need them

3. **"Accept interfaces, return structs."**
   - Make functions flexible with interface parameters
   - Make APIs clear with concrete returns

4. **"The empty interface says nothing."**
   - `interface{}` (now `any`) provides no guarantees
   - Use only when truly necessary

---

## Exercise List (Detailed)

### Tier 1: Introduction

**01 - Implicit Satisfaction**
- Implement Speaker interface with Dog, Cat, Robot types
- Learn: No "implements" keyword needed
- Function: MakeSpeak(s Speaker) string

**02 - Stringer Interface**
- Implement fmt.Stringer for Temperature, Point, Duration
- Learn: How fmt.Println uses interfaces
- Method: String() string

**03 - Error Interface**
- Create ValidationError and NetworkError types
- Learn: error is just an interface
- Method: Error() string

### Tier 2: Application

**04 - Reader Basics**
- Implement RepeatReader and CountingReader
- Learn: io.Reader pattern
- Method: Read([]byte) (int, error)

**05 - Writer Basics**
- Implement UpperWriter and LimitWriter
- Learn: io.Writer pattern
- Method: Write([]byte) (int, error)

**05.5 - Wrapper Chains**
- Chain multiple wrapper Readers into pipelines
- Learn: How wrappers compose via delegation
- Types: UppercaseReader, LimitReader, TeeReader
- Function: BuildPipeline(source, limit) → chained Reader

**05.6 - io.Copy Bridge**
- Connect Readers and Writers with io.Copy
- Learn: io.Copy as the universal connector
- Functions: CopyWithStats, ProcessData, CopyN
- Pattern: Read → Transform → Write pipelines

**06 - Interface Composition**
- Combine Logger and Closer interfaces
- Learn: Interface embedding
- Type: FileLogger implementing LogCloser

**07 - Type Assertions**
- Work with Shape interface (Circle, Rectangle)
- Learn: Checking concrete types at runtime
- Pattern: value, ok := i.(Type)

### Tier 3: Integration

**08 - Type Switch**
- Handle multiple types with type switch
- Learn: switch v.(type) pattern
- Function: Stringify(v any) string

**09 - Mock for Testing**
- Create TimeProvider interface for testing
- Learn: Dependency injection
- Types: RealTime, MockTime, Scheduler

**10 - Small Interfaces**
- Design Sizer and Namer interfaces
- Learn: Single-method interfaces
- Functions: TotalSize, ListNames

**11 - Interface Segregation**
- Separate CRUD operations into small interfaces
- Learn: Interface Segregation Principle (ISP)
- Functions: Accept minimal interfaces

**12 - Polymorphic Collection**
- Work with slices of interfaces
- Learn: Runtime polymorphism
- Type: PriorityQueue with interface slice

---

## How to Use This Module

### Before You Start

1. **Read about interfaces:**
   ```bash
   # Essential reading
   https://go.dev/tour/methods/9
   https://go.dev/doc/effective_go#interfaces
   ```

2. **Understand the philosophy:**
   - Interfaces enable abstraction WITHOUT inheritance
   - Small interfaces are more powerful than large ones
   - Code to interfaces, not implementations

### Working Through Exercises

1. **Read each exercise README carefully** - understand the interface contract
2. **Implement the methods** - satisfy the interface implicitly
3. **Run tests** - `go test -v`
4. **Write EXPLANATION.md** - explain why the interface is useful

### Success Criteria

To complete this module:
- [ ] Complete all 14 exercises (tests passing)
- [ ] Understand implicit interface satisfaction
- [ ] Can implement io.Reader and io.Writer
- [ ] Can chain wrapper Readers/Writers into pipelines
- [ ] Understand io.Copy as the Reader↔Writer bridge
- [ ] Know when to use type assertions vs type switches
- [ ] Can design small, focused interfaces
- [ ] Understand "accept interfaces, return structs"

---

## Quick Reference

### Common Standard Library Interfaces

```go
// fmt.Stringer - custom string representation
type Stringer interface {
    String() string
}

// error - custom error types
type error interface {
    Error() string
}

// io.Reader - reading data
type Reader interface {
    Read(p []byte) (n int, err error)
}

// io.Writer - writing data
type Writer interface {
    Write(p []byte) (n int, err error)
}

// io.Closer - cleanup resources
type Closer interface {
    Close() error
}
```

### Type Assertion Patterns

```go
// Comma-ok idiom (safe)
if val, ok := i.(ConcreteType); ok {
    // Use val
}

// Direct assertion (panics if wrong type)
val := i.(ConcreteType)  // Dangerous!

// Type switch
switch v := i.(type) {
case string:
    // v is string
case int:
    // v is int
default:
    // Unknown type
}
```

---

## Study Tips

1. **Start simple** - Begin with single-method interfaces
2. **Use the compiler** - Let it tell you what methods are needed
3. **Think about behavior** - Interfaces describe what something does, not what it is
4. **Read stdlib source** - See how Go's standard library uses interfaces
5. **Practice small interfaces** - Resist the urge to make large interfaces

---

## After This Module

Once completed, you will:
- Write flexible, testable Go code
- Understand Go's approach to abstraction
- Use standard library interfaces confidently
- Design interfaces that follow Go idioms
- Mock dependencies for testing
- Apply SOLID principles in Go

**Interfaces unlock Go's true power. Master them!**

---

**Ready to learn Go's superpower?** Start with **Exercise 01: Implicit Satisfaction** 🚀
