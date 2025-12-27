# Exercise 02: Stringer Interface

**Concept:** Using fmt.Stringer from the standard library
**Difficulty:** Easy
**Estimated Time:** 35 minutes

## Learning Goal

Learn how to implement the `fmt.Stringer` interface to control how your types are printed. This is one of the most commonly used interfaces in Go and demonstrates how interfaces enable powerful standard library features.

## The Problem

By default, printing a struct shows its field values in a generic format:

```go
type Temperature struct {
    Value float64
    Unit  string
}

t := Temperature{Value: 72.5, Unit: "F"}
fmt.Println(t)  // Output: {72.5 F}  <- Not very readable!
```

By implementing `String() string`, you can customize the output:

```go
func (t Temperature) String() string {
    return fmt.Sprintf("%.1f°%s", t.Value, t.Unit)
}

fmt.Println(t)  // Output: 72.5°F  <- Much better!
```

## The fmt.Stringer Interface

The fmt package defines this interface:

```go
type Stringer interface {
    String() string
}
```

When you print something with `fmt.Println`, `fmt.Printf("%v")`, or similar functions, Go automatically checks if the type implements Stringer. If it does, it calls `String()` instead of using the default format.

## Your Task

Implement the `String()` method for three types to make them print nicely.

### Types to Implement

1. **Temperature** - represents a temperature with value and unit
   - Fields: `Value float64`, `Unit string`
   - Format: "72.5°F" or "22.0°C"

2. **Point** - represents a 2D coordinate
   - Fields: `X int`, `Y int`
   - Format: "(3, 5)"

3. **Duration** - represents a time duration
   - Fields: `Hours int`, `Minutes int`
   - Format: "2h 30m" or "0h 45m"

## Function Signatures

```go
type Temperature struct {
    Value float64
    Unit  string
}

type Point struct {
    X int
    Y int
}

type Duration struct {
    Hours   int
    Minutes int
}

func (t Temperature) String() string
func (p Point) String() string
func (d Duration) String() string
```

## Examples

```go
temp := Temperature{Value: 72.5, Unit: "F"}
fmt.Println(temp)  // → "72.5°F"

point := Point{X: 3, Y: 5}
fmt.Println(point)  // → "(3, 5)"

dur := Duration{Hours: 2, Minutes: 30}
fmt.Println(dur)  // → "2h 30m"

// Also works in formatted strings
fmt.Printf("The temperature is %v\n", temp)  // → "The temperature is 72.5°F"
fmt.Printf("Point: %s\n", point)             // → "Point: (3, 5)"
```

## Instructions

1. Open `stringer_interface.go`
2. Define the three structs with their fields
3. Implement String() method for each type
4. Use `fmt.Sprintf` to format the output
5. Run `go test -v`

## Hints

<details>
<summary>Hint 1: fmt.Sprintf</summary>

Use `fmt.Sprintf` to build formatted strings:

```go
func (t Temperature) String() string {
    return fmt.Sprintf("%.1f°%s", t.Value, t.Unit)
}
```

Format verbs:
- `%f` - float (%.1f means 1 decimal place)
- `%d` - integer
- `%s` - string
</details>

<details>
<summary>Hint 2: Temperature String</summary>

```go
func (t Temperature) String() string {
    return fmt.Sprintf("%.1f°%s", t.Value, t.Unit)
}
```

The degree symbol is just a Unicode character: °
</details>

<details>
<summary>Hint 3: Point String</summary>

```go
func (p Point) String() string {
    return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
```
</details>

<details>
<summary>Hint 4: Duration String</summary>

```go
func (d Duration) String() string {
    return fmt.Sprintf("%dh %dm", d.Hours, d.Minutes)
}
```
</details>

## Think About

1. **When is String() called?**
   - Automatically by fmt.Println, fmt.Printf with %v or %s
   - NOT called when printing individual fields
   - Try: `fmt.Println(temp.Value)` vs `fmt.Println(temp)`

2. **What happens if you don't implement String()?**
   - Go uses default representation: `{72.5 F}`
   - Still works, just not as readable

3. **Can you call String() manually?**
   - Yes! `s := temp.String()`
   - But usually you let fmt package call it

4. **Value receiver vs pointer receiver?**
   - For String(), value receiver is usually fine
   - Reading data doesn't need to modify the struct

## What This Teaches

- **fmt.Stringer interface** - standard library interfaces are powerful
- **Custom formatting** - control how your types are displayed
- **Implicit satisfaction** - no need to declare "implements Stringer"
- **String formatting** - using fmt.Sprintf effectively
- **User-friendly output** - making data readable

## Common Mistakes to Avoid

1. **Wrong method signature:**
   ```go
   // WRONG - takes parameters
   func (t Temperature) String(format string) string { ... }

   // RIGHT - no parameters
   func (t Temperature) String() string { ... }
   ```

2. **Infinite recursion:**
   ```go
   // WRONG - calls Println, which calls String(), which calls Println...
   func (t Temperature) String() string {
       return fmt.Sprintf("Temperature: %v", t)  // Infinite loop!
   }

   // RIGHT - format the fields directly
   func (t Temperature) String() string {
       return fmt.Sprintf("%.1f°%s", t.Value, t.Unit)
   }
   ```

3. **Not using fmt.Sprintf:**
   ```go
   // WORKS but less elegant
   func (t Temperature) String() string {
       return strconv.FormatFloat(t.Value, 'f', 1, 64) + "°" + t.Unit
   }

   // BETTER
   func (t Temperature) String() string {
       return fmt.Sprintf("%.1f°%s", t.Value, t.Unit)
   }
   ```

## Challenge Extensions (Optional)

After completing the basic version, try these:

1. **Conditional formatting:**
   ```go
   // Duration: Don't show hours if 0
   // "45m" instead of "0h 45m"
   func (d Duration) String() string {
       if d.Hours == 0 {
           return fmt.Sprintf("%dm", d.Minutes)
       }
       return fmt.Sprintf("%dh %dm", d.Hours, d.Minutes)
   }
   ```

2. **Pluralization:**
   ```go
   // "1 hour 30 minutes" vs "2 hours 15 minutes"
   ```

3. **Add a Color type:**
   ```go
   type Color struct {
       R, G, B uint8
   }
   // Format as "#FF5733" or "rgb(255, 87, 51)"
   ```

## After Completing

You now understand:
- How fmt.Stringer works
- How to implement String() method
- How fmt package uses interfaces
- How to format strings with fmt.Sprintf
- The power of standard library interfaces

This pattern is used everywhere in Go. Many types implement Stringer to provide readable output!

---

**Next up:** Exercise 03 - Error Interface (creating custom error types)
