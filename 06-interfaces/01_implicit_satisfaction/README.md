# Exercise 01: Implicit Satisfaction

**Concept:** Duck typing and implicit interface satisfaction
**Difficulty:** Easy
**Estimated Time:** 30 minutes

## Learning Goal

Understand Go's unique approach to interfaces: **implicit satisfaction**. Unlike Java or C#, Go doesn't require you to declare that a type implements an interface. If a type has all the methods the interface requires, it automatically satisfies that interface.

This is sometimes called "duck typing" - if it walks like a duck and quacks like a duck, it's a duck!

## The Problem

In traditional OOP languages, you must explicitly declare interface implementation:

```java
// Java - explicit declaration required
class Dog implements Speaker {
    public String speak() {
        return "Woof!";
    }
}
```

In Go, this is unnecessary. You just define the methods:

```go
// Go - no declaration needed!
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

// Dog now satisfies Speaker interface automatically
var s Speaker = Dog{}  // This works!
```

## Your Task

Create three types (Dog, Cat, Robot) that all satisfy a Speaker interface, then write a function that accepts any Speaker and makes it speak.

### Part 1: Define the Interface

Create a `Speaker` interface with one method:
- `Speak() string` - returns what the speaker says

### Part 2: Implement Types

Create three types that satisfy Speaker:

1. **Dog** - says "Woof!"
2. **Cat** - says "Meow!"
3. **Robot** - says "Beep boop!"

### Part 3: Use Polymorphism

Implement `MakeSpeak(s Speaker) string` that:
- Takes any type that satisfies Speaker
- Calls its Speak method
- Returns the result

## Function Signatures

```go
type Speaker interface {
    Speak() string
}

type Dog struct{}
type Cat struct{}
type Robot struct{}

func (d Dog) Speak() string
func (c Cat) Speak() string
func (r Robot) Speak() string

func MakeSpeak(s Speaker) string
```

## Examples

```go
dog := Dog{}
cat := Cat{}
robot := Robot{}

MakeSpeak(dog)    // → "Woof!"
MakeSpeak(cat)    // → "Meow!"
MakeSpeak(robot)  // → "Beep boop!"

// Also works with interface variable
var s Speaker
s = Dog{}
s.Speak()         // → "Woof!"
s = Cat{}
s.Speak()         // → "Meow!"
```

## Instructions

1. Open `implicit_satisfaction.go`
2. Define the Speaker interface
3. Create Dog, Cat, and Robot types (empty structs are fine)
4. Implement the Speak() method for each type
5. Implement MakeSpeak function
6. Run `go test -v`

## Hints

<details>
<summary>Hint 1: Interface Definition</summary>

An interface is just a collection of method signatures:

```go
type Speaker interface {
    Speak() string  // Any type with this method satisfies Speaker
}
```
</details>

<details>
<summary>Hint 2: Implementing Methods</summary>

Add methods to your types using receiver syntax:

```go
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

// Dog now satisfies Speaker!
```

You can use value receivers (d Dog) or pointer receivers (*d Dog). For empty structs, value receivers are fine.
</details>

<details>
<summary>Hint 3: MakeSpeak Function</summary>

The function just needs to call the Speak method:

```go
func MakeSpeak(s Speaker) string {
    return s.Speak()
}
```

This works with ANY type that has a Speak() string method!
</details>

<details>
<summary>Hint 4: Full Solution Pattern</summary>

```go
package implicit_satisfaction

// Define interface
type Speaker interface {
    Speak() string
}

// Define types
type Dog struct{}
type Cat struct{}
type Robot struct{}

// Implement methods (satisfies interface implicitly)
func (d Dog) Speak() string {
    return "Woof!"
}

func (c Cat) Speak() string {
    return "Meow!"
}

func (r Robot) Speak() string {
    return "Beep boop!"
}

// Use interface
func MakeSpeak(s Speaker) string {
    return s.Speak()
}
```
</details>

## Think About

1. **What makes this "implicit"?**
   - You never declared "Dog implements Speaker"
   - Go checks this automatically at compile time
   - If Dog doesn't have Speak(), it won't compile

2. **What are the advantages?**
   - Types can satisfy interfaces they don't know about
   - No coupling between packages
   - Easier to refactor and test

3. **What if you forget a method?**
   - Try commenting out one Speak() method
   - Compiler will tell you the type doesn't satisfy Speaker
   - Error: "Dog does not implement Speaker (missing Speak method)"

4. **Can a type satisfy multiple interfaces?**
   - Yes! If it has all the required methods
   - Example: A type with Read() and Write() satisfies both Reader and Writer

## What This Teaches

- **Implicit interface satisfaction** - no "implements" keyword needed
- **Duck typing** - if it has the right methods, it satisfies the interface
- **Polymorphism** - one function works with many types
- **Compile-time checking** - Go verifies interfaces at compile time
- **Decoupling** - types don't need to know about interfaces

## Common Mistakes to Avoid

1. **Trying to declare implementation:**
   ```go
   // WRONG - Go has no "implements" keyword
   type Dog struct{} implements Speaker
   ```

2. **Wrong method signature:**
   ```go
   // WRONG - returns wrong type
   func (d Dog) Speak() int {
       return 42
   }
   // Error: Dog does not implement Speaker
   ```

3. **Missing method:**
   ```go
   type Dog struct{}
   // Missing Speak() method entirely
   // Error: cannot use Dog{} (type Dog) as type Speaker
   ```

4. **Method on wrong receiver type:**
   ```go
   // If you need *Dog (pointer), but only have Dog (value)
   func (d *Dog) Speak() string { ... }
   var s Speaker = Dog{}  // ERROR
   var s Speaker = &Dog{} // OK
   ```

## After Completing

You now understand:
- How Go's implicit interface satisfaction works
- The power of duck typing
- How to create and satisfy interfaces
- How polymorphism works in Go

This is the foundation for everything else in this module!

---

**Next up:** Exercise 02 - Stringer Interface (using fmt.Stringer from the standard library)
