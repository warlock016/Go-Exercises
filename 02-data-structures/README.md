# Module 02: Data Structures

**Status:** In Progress
**Estimated Time:** 8-12 hours
**Difficulty:** Foundation → Intermediate
**Prerequisites:** Module 01 Fundamentals

---

## 🎯 Learning Objectives

By completing this module, you will:

1. **Master Go's built-in data structures** - Understand slices, maps, arrays, and structs at a deep level
2. **Learn internal implementations** - How slices manage capacity, how maps handle collisions, memory layout of structs
3. **Recognize appropriate use cases** - When to use each data structure based on performance and design requirements
4. **Build custom collections** - Implement stacks, queues, and other structures using Go's primitives
5. **Optimize memory usage** - Understand preallocation, capacity management, and performance implications
6. **Design composite types** - Create complex data models using struct embedding and composition

---

## 📋 Module Overview

This module shifts your focus from **control flow patterns** (loops, conditionals) to **data structure design decisions**. You'll move beyond "how do I iterate?" to "which structure should I choose and why?"

**Core Topics:**
- **Slices**: Creation, operations, capacity vs length, slice internals, algorithms
- **Maps**: Key-value operations, existence checking, complex keys, maps of slices/structs
- **Structs**: Definition, composition, embedding, custom types with methods
- **Arrays**: Fixed-size arrays, multi-dimensional arrays, array vs slice trade-offs
- **Performance**: Memory efficiency, preallocation strategies, benchmarking

---

## 🏗️ Exercise Structure (17 Exercises, 4 Tiers)

### Tier 1: Foundation (Exercises 01-04)
**Goal:** Master the basics of each data structure

- `01_slice_basics` - Slice creation, append, len/cap, iteration
- `02_slice_operations` - Slicing syntax, copy, append patterns, capacity growth
- `03_map_fundamentals` - CRUD operations, existence checking, iteration, delete
- `04_struct_basics` - Definition, instantiation, zero values, field access, comparison

**Expected Time:** 2-3 hours
**Success Criteria:** All tests passing, understand when to use each structure

---

### Tier 2: Application (Exercises 05-08)
**Goal:** Apply structures to solve common programming patterns

- `05_slice_algorithms` - Filter, map, reduce, find patterns (functional-style operations)
- `06_map_patterns` - Counting, grouping, frequency analysis, lookup tables
- `07_struct_composition` - Nested structs, anonymous fields, embedding patterns
- `08_collections` - Implement stack and queue using slices

**Expected Time:** 3-4 hours
**Success Criteria:** Can implement common algorithms using appropriate structures

---

### Tier 3: Integration (Exercises 09-13)
**Goal:** Combine structures to solve complex problems

- `09_two_dimensional_slices` - Matrices, grids, game boards, image representation
- `10_advanced_maps` - Maps of slices, maps of structs, complex keys, nested maps
- `11_data_modeling` - Real-world struct design (e.g., library system, inventory)
- `12_custom_types` - Type definitions with methods, method sets (preview Module 03)
- `13_slice_internals` - Understanding backing arrays, capacity management, sharing

**Expected Time:** 4-5 hours
**Success Criteria:** Can design appropriate data models for real-world problems

---

### Tier 4: Mastery (Exercises 14-17)
**Goal:** Optimize performance and build production-quality structures

- `14_performance_optimization` - Preallocation, memory efficiency, benchmarks, profiling
- `15_generic_collections` - Building reusable data structures (preview generics)
- `16_graph_basics` - Adjacency lists/matrices with maps and slices
- `17_capstone_mini_database` - In-memory key-value store combining all concepts

**Expected Time:** 5-6 hours
**Success Criteria:** Can make informed performance trade-offs and design complex systems

---

## 📚 Key Concepts Covered

### Slices
- Slice literal vs `make()` vs array slicing
- Length vs capacity: `len(s)` and `cap(s)`
- Append mechanics and reallocation
- Slice expressions: `s[i:j]`, `s[:n]`, `s[n:]`, `s[i:j:k]` (full slice expression)
- Copying slices correctly
- Slice sharing and the backing array
- Common pitfalls: append modifying shared backing arrays

### Maps
- Map creation: literal, `make()`, nil maps
- Key requirements (comparable types)
- Value lookup with `ok` idiom: `val, ok := m[key]`
- Checking existence vs zero value
- Deleting keys: `delete(m, key)`
- Iteration order is random (non-deterministic)
- Maps of slices, maps of structs
- Map as a set (map[T]bool or map[T]struct{})

### Structs
- Struct definition and field types
- Struct literals: named vs positional
- Anonymous structs
- Zero values for struct fields
- Struct comparison (when valid)
- Struct embedding vs composition
- Pointer vs value semantics
- Method sets (preview for Module 03)

### Performance Considerations
- Preallocating slices: `make([]T, 0, capacity)`
- Map preallocation: `make(map[K]V, capacity)`
- Memory layout and cache locality
- Benchmarking with `testing.B`
- Profiling memory allocations

---

## 🎓 Success Criteria

**To complete this module:**
- ✅ All 17 exercises completed with tests passing (>85% test pass rate)
- ✅ Can explain when to use slice vs array
- ✅ Can explain when to use map vs slice for lookups
- ✅ Understand slice capacity and backing array behavior
- ✅ Can design structs for real-world data models
- ✅ Can implement common collection types (stack, queue, set)
- ✅ Understand performance implications of data structure choices

**Self-Assessment Questions:**
1. What happens when you append to a slice at capacity?
2. Why can't you use a slice as a map key?
3. What's the difference between `nil` and `make([]T, 0)`?
4. When should you use a pointer to a struct vs a value?
5. How do you implement a set in Go?
6. What's the time complexity of map lookups?

---

## 💡 Learning Tips

1. **Visualize memory** - Draw diagrams of slice backing arrays and capacity growth
2. **Run benchmarks** - Use `go test -bench=.` to see performance differences
3. **Read source code** - Look at how the standard library uses these structures
4. **Experiment in playground** - Test your assumptions about slice sharing and map behavior
5. **Focus on trade-offs** - Every structure has pros/cons - learn when to use each

---

## 🔗 Resources

**Official Documentation:**
- [Slices](https://go.dev/blog/slices-intro) - Official Go blog on slice internals
- [Maps](https://go.dev/blog/maps) - How Go maps work
- [Structs](https://go.dev/ref/spec#Struct_types) - Language specification

**Recommended Reading:**
- "Go Slices: usage and internals" - Go blog
- "Effective Go" sections on slices, maps, and embedding
- [Go by Example](https://gobyexample.com/) - Slices, Maps, Structs sections

**In This Repo:**
- `resources/` - Check for data structure guides (may be created as needed)

---

## 🚀 Getting Started

1. **Start with Tier 1** - Even if you think you know slices/maps, work through exercises 01-04 systematically
2. **Read the README for each exercise** - Understand the learning goal before coding
3. **Run tests frequently** - `go test -v` in each exercise directory
4. **Take your time with internals** - Exercise 13 (slice internals) is crucial for deep understanding
5. **Benchmark your solutions** - Exercise 14 introduces benchmarking - use it throughout

**Target Pace:** 2-3 exercises per day with all tests passing

---

## 📊 Progress Tracking

Track your progress in `PROGRESS.md`:
- Mark exercises complete when all tests pass
- Note time spent and concepts that "clicked"
- Update skill matrix levels as you progress
- Record "aha!" moments in the insights section

---

**Module Created:** 2025-11-16
**Ready to Start:** Yes - All 17 exercises created and ready
**Next Module:** 03 - Functions & Methods (unlocks after completing 80%+ of Module 02)
