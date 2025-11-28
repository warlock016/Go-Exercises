# Exercise 06: List Processing

**Tier:** 2 - Patterns
**Estimated Time:** 60 minutes
**Concepts:** Nested recursion, flattening structures, recursive data processing

---

## Learning Goal

Handle nested and complex data structures recursively. Learn to recurse on irregular structures where depth isn't known in advance.

---

## Problem Description

Process nested structures using recursion - structures within structures.

### 1. Flatten Nested Slices

Given a slice that may contain integers or other slices (represented as `interface{}`), flatten it into a single slice.

- Input: `[1, [2, 3], [4, [5, 6]]]`
- Output: `[1, 2, 3, 4, 5, 6]`

### 2. Deep Sum

Calculate the sum of all integers in a nested structure (slice of slices of... integers).

- Input: `[1, [2, [3, 4]], 5]`
- Output: `15`

### 3. Max Depth

Find the maximum nesting depth of a nested slice structure.

- Input: `[1, [2, [3]]]`
- Output: `3`

### 4. Filter Nested

Keep only elements that satisfy a condition, preserving structure.

- Input: `[1, [2, 3], [4, [5, 6]]]`, keep even numbers
- Output: `[2, 4, 6]` (flattened)

---

## Function Signatures

```go
func Flatten(nested interface{}) []int

func DeepSum(nested interface{}) int

func MaxDepth(nested interface{}) int

func FilterNested(nested interface{}, predicate func(int) bool) []int
```

---

## Examples

### Flatten
```go
Flatten(1)                    // [1]
Flatten([]interface{}{1, 2})  // [1, 2]
Flatten([]interface{}{1, []interface{}{2, 3}})  // [1, 2, 3]
Flatten([]interface{}{1, []interface{}{2, []interface{}{3}}})  // [1, 2, 3]
```

### DeepSum
```go
DeepSum(5)                                      // 5
DeepSum([]interface{}{1, 2, 3})                 // 6
DeepSum([]interface{}{1, []interface{}{2, 3}})  // 6
DeepSum([]interface{}{1, []interface{}{2, []interface{}{3, 4}}})  // 10
```

### MaxDepth
```go
MaxDepth(1)                                       // 1
MaxDepth([]interface{}{1, 2})                     // 1
MaxDepth([]interface{}{1, []interface{}{2}})      // 2
MaxDepth([]interface{}{1, []interface{}{2, []interface{}{3}}})  // 3
```

### FilterNested
```go
FilterNested([]interface{}{1, 2, 3, 4}, isEven)  // [2, 4]
FilterNested([]interface{}{1, []interface{}{2, 3}, 4}, isEven)  // [2, 4]
```

---

## Instructions

1. **Understand `interface{}`**:
   ```go
   // interface{} can hold any type
   var x interface{} = 42
   var y interface{} = []interface{}{1, 2}

   // Type assertion to check actual type
   if num, ok := x.(int); ok {
       // x is an int
   }

   if slice, ok := y.([]interface{}); ok {
       // y is a slice
   }
   ```

2. **Pattern for nested recursion**:
   ```go
   func ProcessNested(nested interface{}) result {
       // Base case: nested is an integer
       if num, ok := nested.(int); ok {
           return processInt(num)
       }

       // Recursive case: nested is a slice
       if slice, ok := nested.([]interface{}); ok {
           result := initialValue
           for _, element := range slice {
               result = combine(result, ProcessNested(element))  // Recurse!
           }
           return result
       }

       return defaultValue
   }
   ```

3. **Implement Flatten**:
   - If element is int, return `[]int{element}`
   - If element is slice, recursively flatten each sub-element
   - Combine all flattened results

4. **Implement DeepSum**:
   - If element is int, return it
   - If element is slice, sum all recursive results

5. **Implement MaxDepth**:
   - If element is int, depth is 1
   - If element is slice, depth is 1 + max depth of children

6. **Implement FilterNested**:
   - Similar to Flatten, but only include elements that pass predicate

7. **Run tests**: `go test -v`

---

## Hints

<details>
<summary><strong>Hint 1 - Basic (Type Checking Pattern)</strong></summary>

**The fundamental pattern:**
```go
func ProcessNested(nested interface{}) {
    // Check if it's a base case (integer)
    if num, ok := nested.(int); ok {
        // Handle integer
        return handleInt(num)
    }

    // Check if it's a recursive case (slice)
    if slice, ok := nested.([]interface{}); ok {
        // Handle slice - recurse on each element
        for _, element := range slice {
            ProcessNested(element)  // Recursive call
        }
    }
}
```

**Key insight:** Type assertion `value.(Type)` returns (value, true) if successful, (zero, false) otherwise.

</details>

<details>
<summary><strong>Hint 2 - Intermediate (Combining Results)</strong></summary>

**Flatten - concatenate slices:**
```go
result := []int{}
for _, element := range slice {
    flattened := Flatten(element)  // Returns []int
    result = append(result, flattened...)  // Spread operator
}
```

**DeepSum - add values:**
```go
sum := 0
for _, element := range slice {
    sum += DeepSum(element)  // Add recursive result
}
```

**MaxDepth - take maximum:**
```go
maxChildDepth := 0
for _, element := range slice {
    childDepth := MaxDepth(element)
    if childDepth > maxChildDepth {
        maxChildDepth = childDepth
    }
}
return 1 + maxChildDepth  // Current level + max child depth
```

</details>

<details>
<summary><strong>Hint 3 - Advanced (Complete Flatten Example)</strong></summary>

**Step-by-step Flatten logic:**

Input: `[1, [2, 3], 4]`

1. Check if int → no
2. Check if slice → yes
3. Loop through elements:
   - Element 1:
     - Check if int → yes → return `[]int{1}`
   - Element [2, 3]:
     - Check if int → no
     - Check if slice → yes
     - Loop through 2, 3:
       - 2 → `[]int{2}`
       - 3 → `[]int{3}`
     - Concatenate → `[]int{2, 3}`
   - Element 4:
     - Check if int → yes → return `[]int{4}`
4. Concatenate all: `[]int{1} + []int{2, 3} + []int{4}` → `[]int{1, 2, 3, 4}`

</details>

<details>
<summary><strong>Hint 4 - Complete Solutions</strong></summary>

**Flatten:**
```go
func Flatten(nested interface{}) []int {
    if num, ok := nested.(int); ok {
        return []int{num}
    }

    if slice, ok := nested.([]interface{}); ok {
        result := []int{}
        for _, element := range slice {
            result = append(result, Flatten(element)...)
        }
        return result
    }

    return []int{}
}
```

**DeepSum:**
```go
func DeepSum(nested interface{}) int {
    if num, ok := nested.(int); ok {
        return num
    }

    if slice, ok := nested.([]interface{}); ok {
        sum := 0
        for _, element := range slice {
            sum += DeepSum(element)
        }
        return sum
    }

    return 0
}
```

**MaxDepth:**
```go
func MaxDepth(nested interface{}) int {
    if _, ok := nested.(int); ok {
        return 1
    }

    if slice, ok := nested.([]interface{}); ok {
        if len(slice) == 0 {
            return 1
        }

        maxChildDepth := 0
        for _, element := range slice {
            childDepth := MaxDepth(element)
            if childDepth > maxChildDepth {
                maxChildDepth = childDepth
            }
        }
        return 1 + maxChildDepth
    }

    return 0
}
```

**FilterNested:**
```go
func FilterNested(nested interface{}, predicate func(int) bool) []int {
    if num, ok := nested.(int); ok {
        if predicate(num) {
            return []int{num}
        }
        return []int{}
    }

    if slice, ok := nested.([]interface{}); ok {
        result := []int{}
        for _, element := range slice {
            result = append(result, FilterNested(element, predicate)...)
        }
        return result
    }

    return []int{}
}
```

</details>

---

## Think About

1. **Type safety:** `interface{}` loses type safety. How do other languages handle this? (Generics, algebraic data types)

2. **Irregular structures:** Unlike arrays where depth is uniform, these structures can have different depths. How does recursion handle this naturally?

3. **Real-world applications:** Where do you see nested structures?
   - JSON parsing
   - File systems (folders within folders)
   - DOM trees in web browsers
   - Expression trees in compilers

4. **Performance:** Each type assertion has a cost. Would a different data structure be more efficient?

5. **Alternative representation:** Instead of `interface{}`, could you define a custom type?
   ```go
   type NestedInt struct {
       IsInt  bool
       Value  int
       Children []NestedInt
   }
   ```

---

## What This Teaches

✅ **Nested recursion** - Recursing on irregular structures
✅ **Type assertions** - Working with `interface{}`
✅ **Structure preservation vs flattening** - Different processing strategies
✅ **Higher-order functions** - Functions as parameters (predicate)
✅ **Real-world pattern** - Handling JSON, trees, file systems
✅ **Depth-first traversal** - Visiting all nodes in nested structures

**Connection to trees:** This is essentially tree traversal! Each nested slice is a node with children.

---

**Next Exercise:** 07 - Integer Algorithms (GCD, digit manipulation)
