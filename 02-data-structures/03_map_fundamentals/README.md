# Exercise 03: Map Fundamentals

## 🎯 Learning Goal
Master the fundamentals of Go maps: creation methods, CRUD operations, the comma-ok idiom for safe lookups, deletion, and iteration patterns.

## 📝 Problem Description

Maps are Go's built-in hash table implementation - they provide fast key-value lookups in O(1) average time. Unlike slices (which use integer indices), maps use keys of any comparable type.

**Key Concepts:**
- Map creation: literals vs `make()` vs nil maps
- The comma-ok idiom: `val, ok := m[key]`
- CRUD operations: Create, Read, Update, Delete
- Iteration with `range`
- Map keys must be comparable (no slices, maps, or functions)

**Important:** Maps are reference types. Passing a map to a function allows that function to modify the original map.

## 🔧 Function Signatures

```go
// CreateMapLiteral creates and returns a map with keys "one", "two", "three"
// and values 1, 2, 3 respectively
func CreateMapLiteral() map[string]int

// CreateMapWithMake creates an empty map with the specified initial capacity
func CreateMapWithMake(size int) map[string]int

// GetValue safely retrieves a value from the map using the comma-ok idiom
// Returns the value and true if the key exists, or zero value and false if not
func GetValue(m map[string]int, key string) (int, bool)

// SetValue sets the key to the specified value in the map
func SetValue(m map[string]int, key string, value int)

// DeleteKey removes the key from the map
// If the key doesn't exist, this is a no-op (safe to call)
func DeleteKey(m map[string]int, key string)

// GetKeys returns a slice containing all keys from the map
// The order is not guaranteed (maps are unordered)
func GetKeys(m map[string]int) []string

// CountOccurrences counts how many times each word appears in the slice
// Returns a map where keys are words and values are their frequencies
func CountOccurrences(words []string) map[string]int
```

## 💡 Examples

```go
// Creating maps
m1 := CreateMapLiteral()  // map[string]int{"one": 1, "two": 2, "three": 3}

m2 := CreateMapWithMake(10)  // Empty map with capacity hint of 10

// Safe lookup with comma-ok
m := map[string]int{"age": 25}
val, ok := GetValue(m, "age")     // val=25, ok=true
val2, ok2 := GetValue(m, "name")  // val2=0, ok2=false

// Setting values
m := make(map[string]int)
SetValue(m, "score", 100)  // m["score"] = 100

// Deleting keys
m := map[string]int{"a": 1, "b": 2}
DeleteKey(m, "a")  // m is now {"b": 2}
DeleteKey(m, "c")  // Safe, nothing happens

// Getting keys
m := map[string]int{"x": 1, "y": 2, "z": 3}
keys := GetKeys(m)  // []string{"x", "y", "z"} (order varies)

// Counting occurrences
words := []string{"cat", "dog", "cat", "bird", "dog", "cat"}
counts := CountOccurrences(words)  // {"cat": 3, "dog": 2, "bird": 1}
```

## 📋 Instructions

1. **CreateMapLiteral:** Use map literal syntax `map[string]int{...}`
2. **CreateMapWithMake:** Use `make(map[string]int, size)` - capacity is a hint, not a limit
3. **GetValue:** Use comma-ok idiom: `val, ok := m[key]`
4. **SetValue:** Simple assignment: `m[key] = value`
5. **DeleteKey:** Use built-in `delete(m, key)` function
6. **GetKeys:** Create a slice and iterate with `for key := range m`
7. **CountOccurrences:** Create a map, iterate through words, increment counts

## 🧪 Testing

```bash
go test -v
```

Expected test count: ~30 tests across all functions

## 🤔 Think About

1. **What's the difference between a nil map and an empty map?**
   - Nil map: `var m map[string]int` - Cannot add keys, will panic!
   - Empty map: `m := make(map[string]int)` - Can add keys normally

2. **Why do we need the comma-ok idiom?**
   - Distinguishes between "key exists with value 0" and "key doesn't exist"
   - Without it, both return 0 (the zero value)

3. **Are maps ordered?**
   - No! Iteration order is random and can vary between runs

4. **What types can be map keys?**
   - Any comparable type: strings, numbers, bools, pointers, structs/arrays of comparable types
   - NOT slices, maps, or functions (they're not comparable)

## 💡 Hints

<details>
<summary>Hint 1: Creating maps</summary>

Three ways to create a map:

```go
// 1. Map literal
m1 := map[string]int{"key": 42}

// 2. Make with capacity hint
m2 := make(map[string]int, 100)

// 3. Nil map (read-only, cannot add keys!)
var m3 map[string]int  // m3 == nil
```

**Important:** Always initialize with literal or `make()` before adding keys!
</details>

<details>
<summary>Hint 2: The comma-ok idiom</summary>

The comma-ok idiom is the standard way to check if a key exists:

```go
m := map[string]int{"age": 25}

// Check if key exists
val, ok := m["age"]
if ok {
    fmt.Println("Age is", val)  // Age is 25
}

// Key doesn't exist
val, ok := m["name"]
// val = 0 (zero value for int)
// ok = false
```

Without comma-ok, you can't distinguish between missing keys and zero values!
</details>

<details>
<summary>Hint 3: CRUD operations</summary>

```go
m := make(map[string]int)

// Create/Update (same syntax!)
m["key"] = 100

// Read with comma-ok
val, ok := m["key"]

// Delete
delete(m, "key")

// Check if deleted
_, exists := m["key"]  // exists = false
```
</details>

<details>
<summary>Hint 4: Iterating over maps</summary>

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}

// Iterate over keys and values
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// Iterate over just keys
for key := range m {
    fmt.Println(key)
}

// Iterate over just values
for _, value := range m {
    fmt.Println(value)
}
```

**Important:** Iteration order is random!
</details>

<details>
<summary>Hint 5: Counting pattern</summary>

A common pattern for counting occurrences:

```go
func CountOccurrences(words []string) map[string]int {
    counts := make(map[string]int)
    for _, word := range words {
        counts[word]++  // Zero value for int is 0, so this works!
    }
    return counts
}
```

The zero value property makes counting elegant - no need to check if key exists first!
</details>

<details>
<summary>Full Solution</summary>

```go
package map_fundamentals

func CreateMapLiteral() map[string]int {
    return map[string]int{
        "one":   1,
        "two":   2,
        "three": 3,
    }
}

func CreateMapWithMake(size int) map[string]int {
    return make(map[string]int, size)
}

func GetValue(m map[string]int, key string) (int, bool) {
    val, ok := m[key]
    return val, ok
}

func SetValue(m map[string]int, key string, value int) {
    m[key] = value
}

func DeleteKey(m map[string]int, key string) {
    delete(m, key)
}

func GetKeys(m map[string]int) []string {
    keys := make([]string, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    return keys
}

func CountOccurrences(words []string) map[string]int {
    counts := make(map[string]int)
    for _, word := range words {
        counts[word]++
    }
    return counts
}
```
</details>

## 🎓 What This Teaches

- **Map creation** - Literals vs make(), understanding capacity hints
- **Comma-ok idiom** - Safe key existence checking
- **CRUD operations** - The four fundamental map operations
- **The delete() function** - Built-in function for removing keys
- **Map iteration** - Using range with maps (unordered)
- **Zero values** - How Go's zero value property simplifies counting
- **Reference semantics** - Maps are reference types (mutations visible across function calls)
- **Counting pattern** - The classic use case for maps in Go

---

**Next Exercise:** `04_struct_basics` - Defining and using custom types with structs
