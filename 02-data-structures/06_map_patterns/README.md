# Exercise 06: Map Patterns

## 🎯 Learning Goal
Master common map patterns including frequency counting, grouping, inversion, merging, and filtering. Understand map iteration order and key extraction.

## 📝 Problem Description

Maps (hash tables) are fundamental data structures in Go. This exercise covers real-world patterns you'll use constantly: counting occurrences, grouping data by keys, inverting key-value relationships, and more.

**Key Concepts:**
- Map initialization with `make()` or literals
- Iteration with `range` (order is random!)
- The "comma ok" idiom: `value, exists := map[key]`
- Maps store references (modification affects original)
- Extracting and sorting keys

## 🔧 Function Signatures

```go
// CountWords returns word frequency map from text
func CountWords(text string) map[string]int

// GroupByLength groups strings by their length
func GroupByLength(words []string) map[int][]string

// InvertMap inverts keys/values (values become keys, original keys in slice)
func InvertMap(m map[string]int) map[int][]string

// MergeMaps combines two maps (m2 values overwrite m1 on conflicts)
func MergeMaps(m1, m2 map[string]int) map[string]int

// MapKeys returns sorted slice of all keys
func MapKeys(m map[string]int) []string

// MapValues returns slice of all values (order unspecified)
func MapValues(m map[string]int) []int

// FilterMap returns new map with entries matching predicate
func FilterMap(m map[string]int, predicate func(string, int) bool) map[string]int
```

## 💡 Examples

```go
// Word frequency
text := "hello world hello"
counts := CountWords(text)
// map[string]int{"hello": 2, "world": 1}

// Group by length
words := []string{"cat", "elephant", "dog", "ant"}
groups := GroupByLength(words)
// map[int][]string{3: ["cat", "dog", "ant"], 8: ["elephant"]}

// Invert map
m := map[string]int{"a": 1, "b": 2, "c": 1}
inverted := InvertMap(m)
// map[int][]string{1: ["a", "c"], 2: ["b"]}

// Merge maps
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 3, "c": 4}
merged := MergeMaps(m1, m2)
// map[string]int{"a": 1, "b": 3, "c": 4}  // m2's "b" overwrites m1's

// Extract keys
m := map[string]int{"z": 1, "a": 2, "m": 3}
keys := MapKeys(m)
// []string{"a", "m", "z"}  // sorted alphabetically

// Filter map
m := map[string]int{"a": 5, "b": 10, "c": 3}
filtered := FilterMap(m, func(k string, v int) bool { return v > 4 })
// map[string]int{"a": 5, "b": 10}
```

## 📋 Instructions

1. **CountWords:** Split text by spaces (`strings.Fields()`), iterate and increment counts
2. **GroupByLength:** Iterate words, use `len(word)` as key, append to slice value
3. **InvertMap:** Iterate original map, append key to slice at inverted[value]
4. **MergeMaps:** Create result map, copy m1, then iterate m2 and overwrite
5. **MapKeys:** Extract keys to slice, sort with `sort.Strings()`
6. **MapValues:** Extract values to slice (order doesn't matter)
7. **FilterMap:** Create result map, iterate and add entries matching predicate

## 🧪 Testing

```bash
go test -v
```

Expected test count: ~35+ tests

## 🤔 Think About

1. **Why is map iteration order random?**
   - Go intentionally randomizes to prevent code relying on order (hash table iteration is undefined)

2. **How do you check if a key exists?**
   - Use two-value form: `value, exists := map[key]`
   - If key doesn't exist, `exists` is false and `value` is zero value

3. **Can you modify a map during iteration?**
   - Yes! But newly added keys may or may not appear in current iteration

4. **Why return []string instead of map[string]bool for keys?**
   - Slices can be sorted, maintain order, and are easier to iterate

5. **What happens when you invert a map with duplicate values?**
   - Multiple keys map to same value, so inverted map needs []string values

## 💡 Hints

<details>
<summary>Hint 1: CountWords implementation</summary>

```go
func CountWords(text string) map[string]int {
    counts := make(map[string]int)
    words := strings.Fields(text)  // Split on whitespace
    for _, word := range words {
        counts[word]++  // Increment count (zero value is 0)
    }
    return counts
}
```
</details>

<details>
<summary>Hint 2: GroupByLength implementation</summary>

```go
func GroupByLength(words []string) map[int][]string {
    groups := make(map[int][]string)
    for _, word := range words {
        length := len(word)
        groups[length] = append(groups[length], word)
    }
    return groups
}
```
</details>

<details>
<summary>Hint 3: InvertMap implementation</summary>

```go
func InvertMap(m map[string]int) map[int][]string {
    result := make(map[int][]string)
    for key, val := range m {
        result[val] = append(result[val], key)
    }
    return result
}
```
</details>

<details>
<summary>Hint 4: MapKeys with sorting</summary>

```go
func MapKeys(m map[string]int) []string {
    keys := make([]string, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    sort.Strings(keys)  // Need: import "sort"
    return keys
}
```
</details>

<details>
<summary>Complete Solution</summary>

```go
import (
    "sort"
    "strings"
)

func CountWords(text string) map[string]int {
    counts := make(map[string]int)
    words := strings.Fields(text)
    for _, word := range words {
        counts[word]++
    }
    return counts
}

func GroupByLength(words []string) map[int][]string {
    groups := make(map[int][]string)
    for _, word := range words {
        length := len(word)
        groups[length] = append(groups[length], word)
    }
    return groups
}

func InvertMap(m map[string]int) map[int][]string {
    result := make(map[int][]string)
    for key, val := range m {
        result[val] = append(result[val], key)
    }
    return result
}

func MergeMaps(m1, m2 map[string]int) map[string]int {
    result := make(map[string]int)
    for key, val := range m1 {
        result[key] = val
    }
    for key, val := range m2 {
        result[key] = val  // Overwrites m1's value
    }
    return result
}

func MapKeys(m map[string]int) []string {
    keys := make([]string, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    sort.Strings(keys)
    return keys
}

func MapValues(m map[string]int) []int {
    values := make([]int, 0, len(m))
    for _, val := range m {
        values = append(values, val)
    }
    return values
}

func FilterMap(m map[string]int, predicate func(string, int) bool) map[string]int {
    result := make(map[string]int)
    for key, val := range m {
        if predicate(key, val) {
            result[key] = val
        }
    }
    return result
}
```
</details>

## 🎓 What This Teaches

- **Frequency counting** - Classic map pattern for aggregation
- **Grouping by key** - Creating map[K][]V structures
- **Map inversion** - Swapping keys and values
- **Map merging** - Combining multiple maps
- **Key extraction** - Getting sorted keys from maps
- **Map filtering** - Creating subsets based on predicates
- **Zero values** - Maps return zero value for missing keys
- **Map iteration** - Understanding randomized order

---

**Next Exercise:** `07_struct_composition` - Nested structs and composition patterns
