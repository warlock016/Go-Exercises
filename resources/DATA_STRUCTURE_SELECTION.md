# Data Structure Selection Guide

> **From Making Code Work to Making Architectural Decisions**

This guide teaches you how to choose the right data structure for the job - a critical skill that separates junior developers from intermediate+ engineers. Good choices make your code fast, maintainable, and scalable. Bad choices create performance bottlenecks, confusing code, and painful refactors.

## Table of Contents

1. [Why This Matters](#why-this-matters)
2. [The Core Question: Access Patterns](#the-core-question-access-patterns)
3. [The Big Three: Array, Slice, Map](#the-big-three-array-slice-map)
4. [Decision Framework](#decision-framework)
5. [Trade-off Analysis](#trade-off-analysis)
6. [Common Scenarios with Solutions](#common-scenarios-with-solutions)
7. [Performance Considerations](#performance-considerations)
8. [The Hybrid Pattern](#the-hybrid-pattern)
9. [Real-World System Design](#real-world-system-design)
10. [Common Mistakes](#common-mistakes)
11. [Benchmarking Guide](#benchmarking-guide)
12. [Quick Reference Tables](#quick-reference-tables)
13. [Practice Exercises](#practice-exercises)

---

## Why This Matters

### Impact on Performance

```go
// BAD: O(n) lookup every time
type Library struct {
    Books []Book  // Need to search through ALL books to find one
}

func GetBook(lib *Library, id int) *Book {
    for i := range lib.Books {
        if lib.Books[i].ID == id {
            return &lib.Books[i]
        }
    }
    return nil
}

// GOOD: O(1) lookup
type Library struct {
    Books map[int]Book  // Direct access by ID
}

func GetBook(lib *Library, id int) Book {
    return lib.Books[id]  // Instant lookup
}
```

**Performance difference with 10,000 books:**
- Slice: Average 5,000 comparisons per lookup
- Map: ~1 hash calculation per lookup
- **5,000x faster with the right structure**

### Impact on Maintainability

```go
// CONFUSING: What does this slice represent?
type Member struct {
    Books []Book  // Full book copies? Which books? Can they change?
}

// CLEAR: Member references books they've checked out
type Member struct {
    CheckedOutBooks []int  // Book IDs - single source of truth
}
```

### How Good Choices Scale

```
10 items: Slice vs Map - barely noticeable difference
100 items: Slice starts to lag for lookups
1,000 items: Slice lookups noticeably slow
10,000 items: Slice unusable for frequent lookups
100,000 items: Only map is viable
```

**The principle:** Choose structures that maintain acceptable performance as data grows. Don't optimize prematurely, but don't box yourself into O(n²) algorithms either.

---

## The Core Question: Access Patterns

**This is the #1 factor in data structure selection.**

Before choosing a structure, ask: **"How will I access this data?"**

### Common Access Patterns

| Pattern | Description | Best Structure |
|---------|-------------|----------------|
| **Lookup by key** | Get item by ID, username, etc. | `map[K]V` |
| **Sequential access** | Process all items in order | `[]T` |
| **Frequent search** | Find items matching criteria | `[]T` (small) or index with `map` (large) |
| **Maintain order** | Items must stay in specific sequence | `[]T` |
| **Unique items** | No duplicates allowed | `map[T]bool` |
| **Count occurrences** | Track how many times each item appears | `map[T]int` |
| **Check membership** | Does X exist in set? | `map[T]bool` |

### Real Example from Exercise 11: Library System

Let's analyze YOUR Library implementation:

```go
type Library struct {
    Books   map[int]Book    // Why map?
    Members map[int]Member  // Why map?
}

type Member struct {
    CheckedOutBooks []int   // Why slice?
}
```

**Why `map[int]Book` for Books?**

Access pattern analysis:
- `AddBook()` - Need to store book by ID
- `RemoveBook(bookID)` - Need to find and delete by ID → **O(1) with map**
- `CheckoutBook(bookID)` - Need to check if book exists and is available → **O(1) with map**
- `FindBooksByAuthor()` - Need to search all books anyway → **Map doesn't hurt**

Primary pattern: **Lookup by ID** → Map is correct choice

**Why `[]int` for CheckedOutBooks?**

Access pattern analysis:
- `CheckoutBook()` - Need to add a book ID → **O(1) append to slice**
- `ReturnBook()` - Need to find and remove a book ID → **O(n) search, but...**
- `GetMemberBooks()` - Need to iterate all checked out books → **O(n) iteration required anyway**

Secondary considerations:
- Members typically have 0-10 books (small size)
- Order might matter (checkout history)
- No need for "does member have book X?" queries (just checking returns)

Primary pattern: **Small sequential access** → Slice is correct choice

### Pattern Recognition Exercise

For each scenario, identify the access pattern:

1. User authentication: Check if username/password exists
   - Pattern: **Lookup by key (username)**

2. Shopping cart: Display items in order added
   - Pattern: **Maintain order + sequential access**

3. View count tracker: Count page views per URL
   - Pattern: **Count occurrences**

4. Friend list on social media: Show all friends, check if X is friend
   - Pattern: **Check membership + sequential access** (might need both structures!)

---

## The Big Three: Array, Slice, Map

### Array `[N]T`

**What it is:** Fixed-size sequence of elements, size is part of the type.

```go
var arr [5]int          // 5 integers, fixed size
arr[0] = 10             // O(1) access
len(arr)                // Always 5
```

**When to use:**
- Size is known at compile time and never changes
- Need stack allocation (arrays can live on stack)
- Working with fixed-size data (IPv4 address, chess board, etc.)
- Want type safety around size

**Performance:**
- Access by index: **O(1)**
- Search: **O(n)**
- Insert/Delete: **Not applicable** (size is fixed)

**Memory:**
- Contiguous memory block
- Size: `N * sizeof(T)` bytes
- No overhead (just the elements)

**Go-specific considerations:**
- Arrays are values - copying an array copies all elements
- Rarely used directly - slices are more flexible
- Can convert to slice: `arr[:]`

**Example use cases:**
```go
// IPv4 address - always 4 bytes
type IPv4 [4]byte

// Chess board - always 8x8
type Board [8][8]Piece

// Fixed buffer for reading
var buffer [4096]byte
```

### Slice `[]T`

**What it is:** Dynamic-size view over an underlying array.

```go
slice := []int{1, 2, 3}          // Literal syntax
slice = append(slice, 4)          // Grows dynamically
slice[0]                          // O(1) access
slice[1:3]                        // O(1) slicing operation
```

**When to use:**
- Size unknown or changes at runtime
- Need ordered collection
- Iterating through all elements
- Stack/queue/list operations
- Most general-purpose collections

**Performance:**
- Access by index: **O(1)**
- Append (amortized): **O(1)**
- Insert at position: **O(n)** (requires shifting)
- Delete at position: **O(n)** (requires shifting)
- Search: **O(n)**
- Sort: **O(n log n)**

**Memory:**
- Slice header: 24 bytes (pointer + len + cap)
- Backing array: Depends on capacity
- Growth: Doubles capacity when full (roughly)

**Go-specific considerations:**
- Slices are references to underlying arrays
- Copying a slice copies the header, not the data
- Multiple slices can share the same backing array
- Be careful with slice tricks like `append(s[:i], s[i+1:]...)` - can cause bugs

**Example use cases:**
```go
// Query results - size unknown
results := []User{}
for rows.Next() {
    results = append(results, user)
}

// Checked out books - small, ordered list
type Member struct {
    CheckedOutBooks []int
}

// Function arguments - most flexible
func ProcessItems(items []Item) {}
```

### Map `map[K]V`

**What it is:** Hash table mapping keys to values.

```go
m := make(map[string]int)
m["alice"] = 25                  // O(1) insert
age := m["alice"]                // O(1) lookup
delete(m, "alice")               // O(1) delete
age, exists := m["bob"]          // Comma-ok idiom
```

**When to use:**
- Need fast lookup by key
- Keys are unique identifiers
- Order doesn't matter (maps are unordered)
- Checking membership (set pattern with `map[K]bool`)
- Counting occurrences (`map[K]int`)

**Performance:**
- Lookup by key: **O(1)** average, O(n) worst case (rare)
- Insert: **O(1)** average
- Delete: **O(1)** average
- Iteration: **O(n)** but unordered

**Memory:**
- Overhead: ~10.79 bytes per key-value pair (plus key/value sizes)
- Grows in buckets
- Never shrinks (even after deletions)

**Go-specific considerations:**
- Iteration order is randomized (intentionally)
- Zero value is nil - must use `make()` before adding
- Retrieving missing key returns zero value (use comma-ok)
- Maps are references - copying map copies reference, not data
- Not thread-safe (need sync.Map or mutex for concurrent access)

**Example use cases:**
```go
// Entity lookup by ID
type Library struct {
    Books map[int]Book
}

// Counting occurrences
wordCount := make(map[string]int)
for _, word := range words {
    wordCount[word]++
}

// Set (unique items)
seen := make(map[string]bool)
if !seen[item] {
    seen[item] = true
    // Process unique item
}

// Cache
cache := make(map[string]Result)
if result, found := cache[key]; found {
    return result
}
```

---

## Decision Framework

### The Decision Tree

```
START: What data structure should I use?
│
├─ Do I need to look up items by a key (ID, name, etc.)?
│  │
│  ├─ YES → Use map[K]V
│  │         ├─ Need ordering too? → Use map + []K (hybrid pattern)
│  │         └─ Just lookup? → map[K]V is perfect
│  │
│  └─ NO → Continue to next question
│
├─ Do I need to maintain a specific order?
│  │
│  ├─ YES → Use []T
│  │         └─ Frequent lookups too? → Consider map + []T (hybrid)
│  │
│  └─ NO → Continue to next question
│
├─ Do I need unique items only (set behavior)?
│  │
│  ├─ YES → Use map[T]bool or map[T]struct{}
│  │
│  └─ NO → Continue to next question
│
├─ Am I counting occurrences?
│  │
│  ├─ YES → Use map[T]int
│  │
│  └─ NO → Continue to next question
│
├─ Is the size fixed and known at compile time?
│  │
│  ├─ YES → Use [N]T (array)
│  │
│  └─ NO → Use []T (slice) as default
```

### Questions to Ask Yourself

Before choosing a structure, answer these:

1. **What's the primary operation?**
   - Lookup by key? → Map
   - Sequential processing? → Slice
   - Membership checking? → Map (set pattern)

2. **How large will this collection get?**
   - 0-10 items: Slice is fine even for lookups
   - 10-100 items: Depends on lookup frequency
   - 100+ items: Map for lookups, slice for iteration
   - 1000+ items: Definitely use map for lookups

3. **Does order matter?**
   - Must maintain order → Slice (or hybrid)
   - Order doesn't matter → Map is fine

4. **How often will I modify it?**
   - Frequent adds/deletes by key → Map
   - Frequent adds to end → Slice
   - Frequent inserts/deletes in middle → Consider other structures (or slice if rare)

5. **What's the access pattern?**
   - By key → Map
   - By index → Slice/Array
   - Sequential iteration → Either works
   - Random access → Depends on key vs index

6. **Will I need to search for items?**
   - Search by key → Map
   - Search by property → Must iterate (slice or map values)
   - Frequent searches → Consider indexing with additional maps

### Red Flags for Wrong Choices

**Red flag: Using slice for lookups with large data**

```go
// BAD: O(n) lookup on every operation
type UserService struct {
    Users []User  // 10,000 users
}

func (s *UserService) GetUserByID(id int) *User {
    for i := range s.Users {  // Searches all 10,000 users!
        if s.Users[i].ID == id {
            return &s.Users[i]
        }
    }
    return nil
}
```

**Red flag: Using map when order matters**

```go
// BAD: Iteration order is random
type Playlist struct {
    Songs map[int]Song
}

func (p *Playlist) Display() {
    for _, song := range p.Songs {  // Random order every time!
        fmt.Println(song.Title)
    }
}
```

**Red flag: Storing full objects when IDs suffice**

```go
// BAD: Duplicate data, can get out of sync
type Member struct {
    CheckedOutBooks []Book  // What if book details change?
}

// GOOD: Single source of truth
type Member struct {
    CheckedOutBooks []int  // Book IDs - lookup in Books map
}
```

---

## Trade-off Analysis

### Slice vs Map: The Library Books Decision

Let's analyze YOUR decision from Exercise 11:

**Scenario:** Store all books in a library

**Option 1: Slice**
```go
type Library struct {
    Books []Book
}
```

Pros:
- Maintains insertion order
- Simple iteration (no need for keys)
- Slightly less memory overhead for small collections

Cons:
- `RemoveBook(id)` requires O(n) search
- `CheckoutBook(id)` requires O(n) search to check availability
- Every ID-based operation is slow

**Option 2: Map (YOUR CHOICE)**
```go
type Library struct {
    Books map[int]Book
}
```

Pros:
- O(1) lookup by ID for all operations
- O(1) insert and delete
- Scales to thousands of books

Cons:
- No guaranteed order
- Slightly more memory overhead
- Can't easily get "nth book"

**Analysis:**

Your operations:
- `AddBook()` - Same cost for both (O(1))
- `RemoveBook(id)` - **Map wins: O(1) vs O(n)**
- `CheckoutBook(id)` - **Map wins: O(1) vs O(n)**
- `FindBooksByAuthor()` - Same cost (O(n) iteration required)

**Verdict:** Map is correct. The primary access pattern is lookup by ID.

### Slice vs Map: The Checked Out Books Decision

**Scenario:** Track which books a member has checked out

**Option 1: Slice (YOUR CHOICE)**
```go
type Member struct {
    CheckedOutBooks []int
}
```

Pros:
- Maintains checkout order (might be useful)
- Simple iteration
- Less memory overhead
- Members typically have 0-10 books (small size)

Cons:
- `ReturnBook()` requires O(n) search to remove
- Can't quickly check "does member have book X?"

**Option 2: Map**
```go
type Member struct {
    CheckedOutBooks map[int]bool
}
```

Pros:
- O(1) check if member has specific book
- O(1) removal

Cons:
- No order information
- More memory overhead
- Overkill for small collections

**Analysis:**

Your operations:
- `CheckoutBook()` - Slice append: O(1) vs Map insert: O(1) - **Tie**
- `ReturnBook()` - Slice: O(n) search + delete vs Map: O(1) delete - **Map wins**
- `GetMemberBooks()` - Slice: O(n) iteration vs Map: O(n) iteration - **Tie**

**But consider:**
- Average checked out books per member: 0-5
- O(n) with n=5 is only 5 operations (negligible)
- Order might be useful (checkout history)
- Slice is simpler and more memory efficient

**Verdict:** Slice is correct for typical library usage. If members could check out 100+ books, reconsider.

**Rule of thumb:** For small collections (< 20 items), slice is fine even if you need searches. The overhead of a map isn't worth it.

### Map with Value vs Map with Pointer

**Option 1: Map with Value**
```go
type Library struct {
    Books map[int]Book
}

func UpdateBook(lib *Library, id int, newTitle string) {
    book := lib.Books[id]        // Gets COPY
    book.Title = newTitle
    lib.Books[id] = book         // Must reassign to map
}
```

Pros:
- Simple value semantics
- No pointer indirection
- No risk of nil pointer panics

Cons:
- Must reassign to map after modification
- Copies data on retrieval (though small structs are cheap)

**Option 2: Map with Pointer**
```go
type Library struct {
    Books map[int]*Book
}

func UpdateBook(lib *Library, id int, newTitle string) {
    lib.Books[id].Title = newTitle  // Modifies directly
}
```

Pros:
- Can modify in place
- No reassignment needed
- Efficient for large structs

Cons:
- Must check for nil: `if book := lib.Books[id]; book != nil { ... }`
- Can accidentally create nil pointers
- More complex semantics

**When to use each:**
- Small structs (< 100 bytes): Use value - simpler and fast enough
- Large structs (> 100 bytes): Consider pointer - avoid copies
- Need nil to represent "not found": Use pointer
- Want simple code: Use value

**Your Library System:** Book struct is small (~50 bytes), value is correct choice.

### String vs []byte

**String:**
- Immutable
- Can use as map key
- String literals are efficient
- Concatenation creates new strings (can be slow)

**[]byte:**
- Mutable (can modify in place)
- Better for building strings
- Can't use as map key directly
- Convert to string when needed

```go
// Building a string - BAD
s := ""
for _, word := range words {
    s += word + " "  // Creates new string every iteration!
}

// Building a string - GOOD
var buf []byte
for _, word := range words {
    buf = append(buf, word...)
    buf = append(buf, ' ')
}
s := string(buf)

// Building a string - BEST
var sb strings.Builder
for _, word := range words {
    sb.WriteString(word)
    sb.WriteString(" ")
}
s := sb.String()
```

---

## Common Scenarios with Solutions

### Scenario 1: Lookup by ID

**Use case:** Get user by ID, get book by ISBN, get product by SKU

**Solution:** `map[K]V`

```go
type UserStore struct {
    Users map[int]User
}

// O(1) lookup
func (s *UserStore) GetUser(id int) (User, bool) {
    user, exists := s.Users[id]
    return user, exists
}
```

**Why:** Primary operation is lookup by unique identifier.

### Scenario 2: Maintain Order

**Use case:** Task list, playlist, feed, history

**Solution:** `[]T`

```go
type TaskList struct {
    Tasks []Task
}

// Add to end - O(1)
func (t *TaskList) AddTask(task Task) {
    t.Tasks = append(t.Tasks, task)
}

// Display in order
func (t *TaskList) Display() {
    for i, task := range t.Tasks {
        fmt.Printf("%d. %s\n", i+1, task.Title)
    }
}
```

**Why:** Order is essential, no need for keyed lookup.

### Scenario 3: Unique Items (Set)

**Use case:** Track seen URLs, unique visitors, deduplicate list

**Solution:** `map[T]bool` or `map[T]struct{}`

```go
// Option 1: map[T]bool - more readable
seen := make(map[string]bool)
for _, url := range urls {
    if !seen[url] {
        seen[url] = true
        ProcessURL(url)
    }
}

// Option 2: map[T]struct{} - saves memory (no bool value)
seen := make(map[string]struct{})
for _, url := range urls {
    if _, exists := seen[url]; !exists {
        seen[url] = struct{}{}
        ProcessURL(url)
    }
}
```

**Why:** Need O(1) membership check, order doesn't matter.

### Scenario 4: Count Occurrences

**Use case:** Word frequency, vote counting, event tracking

**Solution:** `map[T]int`

```go
wordCount := make(map[string]int)
for _, word := range words {
    wordCount[word]++  // Missing key returns 0, so this works!
}

// Display sorted by count
type wordFreq struct {
    word  string
    count int
}
words := []wordFreq{}
for word, count := range wordCount {
    words = append(words, wordFreq{word, count})
}
sort.Slice(words, func(i, j int) bool {
    return words[i].count > words[j].count
})
```

**Why:** Need to track counts per key, map handles missing keys gracefully.

### Scenario 5: Graph Adjacency List

**Use case:** Social network (friends), URL links, dependencies

**Solution:** `map[Node][]Node`

```go
type Graph struct {
    Edges map[string][]string  // Node -> list of connected nodes
}

// Add edge from A to B
func (g *Graph) AddEdge(from, to string) {
    g.Edges[from] = append(g.Edges[from], to)
}

// Get neighbors of node
func (g *Graph) Neighbors(node string) []string {
    return g.Edges[node]
}
```

**Why:** Need to look up neighbors by node (map), each node can have multiple neighbors (slice).

### Scenario 6: Cache

**Use case:** Memoization, expensive computation results, API responses

**Solution:** `map[Key]Value`

```go
type Cache struct {
    Results map[string]Result
}

func (c *Cache) GetOrCompute(key string, compute func() Result) Result {
    if result, found := c.Results[key]; found {
        return result  // Cache hit - O(1)
    }
    result := compute()
    c.Results[key] = result
    return result
}
```

**Why:** Need fast lookup to check if cached, map provides O(1) access.

### Scenario 7: Queue

**Use case:** Job processing, breadth-first search, event handling

**Solution:** `[]T` with append and slicing

```go
type Queue struct {
    Items []Task
}

// Enqueue - O(1) amortized
func (q *Queue) Enqueue(task Task) {
    q.Items = append(q.Items, task)
}

// Dequeue - O(1) but leaves unused memory
func (q *Queue) Dequeue() (Task, bool) {
    if len(q.Items) == 0 {
        return Task{}, false
    }
    task := q.Items[0]
    q.Items = q.Items[1:]  // Advance slice
    return task, true
}
```

**Why:** Need FIFO behavior, slice with append/slice operations works well.

**Note:** For high-throughput queues, consider ring buffer to avoid memory buildup.

### Scenario 8: Stack

**Use case:** Undo/redo, parser, depth-first search, function call tracking

**Solution:** `[]T` with append and slicing

```go
type Stack struct {
    Items []string
}

// Push - O(1) amortized
func (s *Stack) Push(item string) {
    s.Items = append(s.Items, item)
}

// Pop - O(1)
func (s *Stack) Pop() (string, bool) {
    if len(s.Items) == 0 {
        return "", false
    }
    n := len(s.Items) - 1
    item := s.Items[n]
    s.Items = s.Items[:n]  // Shrink slice
    return item, true
}
```

**Why:** Need LIFO behavior, slice with append/slice at end is efficient.

---

## Performance Considerations

### Big-O Complexity Table

| Operation | Array/Slice (by index) | Slice (search) | Map (by key) | Map (iteration) |
|-----------|------------------------|----------------|--------------|-----------------|
| **Access** | O(1) | O(n) | O(1) avg | N/A |
| **Insert at end** | O(1) amortized | O(1) amortized | O(1) avg | N/A |
| **Insert at start/middle** | O(n) | O(n) | N/A | N/A |
| **Delete by key** | O(1) (sets zero) | O(n) search + O(n) shift | O(1) avg | N/A |
| **Search by value** | O(n) | O(n) | O(n) | O(n) |
| **Iterate all** | O(n) | O(n) | N/A | O(n) |
| **Sort** | O(n log n) | O(n log n) | N/A (can't sort map) | N/A |

**Key takeaways:**
- **Map wins for lookup by key** - O(1) vs O(n)
- **Slice wins for ordered access** - Maps are unordered
- **Both are O(n) for iteration** - No clear winner
- **Slice is better for indexed access** - Maps don't have indexes

### When Size Matters

Different sizes require different strategies:

**0-10 items:**
- Slice is fine for everything
- O(n) search on 10 items = 10 operations (negligible)
- Map overhead not worth it
- Keep it simple

**10-100 items:**
- Depends on operation frequency
- Frequent lookups? → Map
- Occasional lookups? → Slice is fine
- Consider both if you need lookup + order

**100-1,000 items:**
- Definitely use map for lookups
- O(n) on 1,000 items = 1,000 operations (noticeable)
- Slice still fine for one-time filtering

**1,000-10,000 items:**
- Map is essential for lookups
- Consider indexing multiple fields (multiple maps)
- Be mindful of memory usage

**10,000+ items:**
- Map for all keyed access
- Consider database if data doesn't fit in memory
- Profile and benchmark your actual usage

### Memory Overhead

**Approximate memory usage:**

```go
// Slice
type Slice []int64  // 8 bytes per element
s := make([]int64, 100)
// Memory: 24 bytes (header) + 800 bytes (data) = 824 bytes

// Map
type Map map[int64]int64  // ~10.79 bytes overhead per entry
m := make(map[int64]int64)
for i := 0; i < 100; i++ {
    m[int64(i)] = int64(i)
}
// Memory: ~2,879 bytes (includes overhead + data)
```

**Memory comparison:**

| Collection | Size | Slice Memory | Map Memory | Map Overhead |
|------------|------|--------------|------------|--------------|
| 10 items | 80B data | 104 bytes | 188 bytes | 1.8x |
| 100 items | 800B data | 824 bytes | 2,879 bytes | 3.5x |
| 1,000 items | 8KB data | 8,024 bytes | 26,790 bytes | 3.3x |
| 10,000 items | 80KB data | 80,024 bytes | 267,900 bytes | 3.3x |

**Rule of thumb:** Maps use ~3-4x more memory than slices for the same data.

**When it matters:**
- Millions of small objects → Consider slice
- Embedded systems / memory-constrained → Prefer slice
- Typical applications → Map overhead is acceptable for performance gain

### Preallocating for Performance

**Slice preallocation:**

```go
// BAD: Grows repeatedly (multiple allocations)
var results []User
for _, id := range userIDs {  // 10,000 IDs
    results = append(results, GetUser(id))
}

// GOOD: Preallocate if size is known
results := make([]User, 0, len(userIDs))
for _, id := range userIDs {
    results = append(results, GetUser(id))
}
```

**Map preallocation:**

```go
// BAD: Grows from 0
m := make(map[int]User)
for _, user := range users {  // 10,000 users
    m[user.ID] = user
}

// GOOD: Preallocate if size is known
m := make(map[int]User, len(users))
for _, user := range users {
    m[user.ID] = user
}
```

**Performance gain:** 20-50% faster for large collections by avoiding reallocation.

---

## The Hybrid Pattern

Sometimes you need BOTH fast lookup AND ordering. That's when you use the hybrid pattern.

### When You Need Both

**Scenario:** Library books where you need:
1. Fast lookup by ID (map behavior)
2. Display in insertion order (slice behavior)

**Problem with map alone:**
```go
type Library struct {
    Books map[int]Book  // Fast lookup but random iteration order
}

// This prints in RANDOM order every time
func (l *Library) DisplayBooks() {
    for _, book := range l.Books {
        fmt.Println(book.Title)
    }
}
```

**Problem with slice alone:**
```go
type Library struct {
    Books []Book  // Ordered but O(n) lookup
}

// This is O(n) for every checkout
func (l *Library) CheckoutBook(bookID int) bool {
    for i := range l.Books {  // Must search all books
        if l.Books[i].ID == bookID {
            // ...
        }
    }
}
```

### The Solution: Map + Slice

**Pattern:** Store data in map, store keys in slice for ordering.

```go
type Library struct {
    Books    map[int]Book  // Fast lookup: O(1)
    BookIDs  []int         // Maintains insertion order
}

// Add book - O(1) for both operations
func (l *Library) AddBook(book Book) {
    l.Books[book.ID] = book
    l.BookIDs = append(l.BookIDs, book.ID)
}

// Get by ID - O(1)
func (l *Library) GetBook(id int) (Book, bool) {
    book, exists := l.Books[id]
    return book, exists
}

// Display in order - O(n) but in insertion order
func (l *Library) DisplayBooks() {
    for _, id := range l.BookIDs {
        book := l.Books[id]
        fmt.Println(book.Title)
    }
}

// Remove book - O(n) to find in slice, O(1) to delete from map
func (l *Library) RemoveBook(id int) bool {
    if _, exists := l.Books[id]; !exists {
        return false
    }

    delete(l.Books, id)  // O(1)

    // Remove from slice - O(n)
    for i, bookID := range l.BookIDs {
        if bookID == id {
            l.BookIDs = append(l.BookIDs[:i], l.BookIDs[i+1:]...)
            break
        }
    }
    return true
}
```

### Real-World Example: Music Playlist

```go
type Playlist struct {
    Songs      map[string]Song  // SongID -> Song (fast lookup)
    SongOrder  []string         // Maintains play order
}

// Add song to end
func (p *Playlist) AddSong(song Song) {
    p.Songs[song.ID] = song
    p.SongOrder = append(p.SongOrder, song.ID)
}

// Insert song at position
func (p *Playlist) InsertSongAt(song Song, position int) {
    p.Songs[song.ID] = song

    // Insert into slice at position
    p.SongOrder = append(p.SongOrder[:position],
                         append([]string{song.ID}, p.SongOrder[position:]...)...)
}

// Play in order
func (p *Playlist) Play() {
    for _, songID := range p.SongOrder {
        song := p.Songs[songID]
        PlaySong(song)
    }
}

// Shuffle
func (p *Playlist) Shuffle() {
    rand.Shuffle(len(p.SongOrder), func(i, j int) {
        p.SongOrder[i], p.SongOrder[j] = p.SongOrder[j], p.SongOrder[i]
    })
}

// Check if song exists - O(1)
func (p *Playlist) HasSong(songID string) bool {
    _, exists := p.Songs[songID]
    return exists
}
```

### Trade-offs of Hybrid Pattern

**Pros:**
- O(1) lookup by key
- Maintains specific ordering
- Can reorder without changing data
- Can have multiple orderings (multiple slices, same map)

**Cons:**
- More memory (storing keys twice: once in map, once in slice)
- More complex to maintain consistency
- Deletes are still O(n) due to slice removal
- Must keep both structures in sync

**When to use:**
- Need both fast lookup AND ordering
- Reordering is common (sorting, shuffling)
- Collection size is medium (100-10,000 items)
- Deletions are rare

**When not to use:**
- Order doesn't matter (just use map)
- Lookups are rare (just use slice)
- Frequent deletes (slice removal is O(n))

---

## Real-World System Design

Let's analyze real systems and understand WHY each structure was chosen.

### System 1: Library System (Your Exercise 11)

```go
type Book struct {
    ID        int
    Title     string
    Author    string
    ISBN      string
    Available bool
}

type Member struct {
    ID              int
    Name            string
    CheckedOutBooks []int    // Why slice?
}

type Library struct {
    Books        map[int]Book    // Why map?
    Members      map[int]Member  // Why map?
    nextBookId   int
    nextMemberId int
}
```

**Analysis:**

**Books - `map[int]Book`:**
- Primary operation: `CheckoutBook(bookID)` - needs to find book by ID
- Also: `RemoveBook(bookID)`, `ReturnBook(bookID)` - all by ID
- Size: Can grow to thousands of books
- Order: Not important for operations
- **Decision: Map is correct** - O(1) lookup scales well

**Members - `map[int]Member`:**
- Primary operation: `CheckoutBook(memberID, bookID)` - needs to find member by ID
- Also: `GetMemberBooks(memberID)` - by ID
- Size: Can grow to thousands of members
- Order: Not important
- **Decision: Map is correct** - O(1) lookup scales well

**CheckedOutBooks - `[]int`:**
- Primary operation: `CheckoutBook()` - append to list
- Secondary: `ReturnBook()` - find and remove (O(n) but small list)
- Also: `GetMemberBooks()` - iterate all (O(n) required anyway)
- Size: Typically 0-10 books per member
- Order: Might be useful (checkout history)
- Alternative: `map[int]bool` for O(1) removal
- **Decision: Slice is correct** - Small size makes O(n) acceptable, simpler code

**Key insight:** Different parts of the same system need different structures based on THEIR access patterns and sizes.

### System 2: Music Playlist System

```go
type Song struct {
    ID       string
    Title    string
    Artist   string
    Duration int
}

type Playlist struct {
    Songs      map[string]Song  // Why map?
    SongOrder  []string         // Why slice?
}
```

**Access patterns:**
- Play songs in order → Need ordering (slice)
- Skip to specific song → Need fast lookup (map)
- Shuffle → Need to reorder (slice)
- Check if song exists → Need fast lookup (map)

**Decision: Hybrid pattern** - Both structures needed.

**Alternative designs:**

```go
// Alternative 1: Just slice - WORSE
type Playlist struct {
    Songs []Song  // O(n) to find specific song for skipping
}

// Alternative 2: Just map - WORSE
type Playlist struct {
    Songs map[string]Song  // Can't maintain play order
}

// Current design - BEST
type Playlist struct {
    Songs     map[string]Song  // O(1) lookup
    SongOrder []string         // Maintains order
}
```

### System 3: E-commerce Shopping Cart

```go
type CartItem struct {
    ProductID string
    Quantity  int
    Price     float64
}

type ShoppingCart struct {
    Items map[string]CartItem  // ProductID -> CartItem
}
```

**Access patterns:**
- Add product → Insert by ID
- Update quantity → Lookup by ID and modify
- Remove product → Delete by ID
- Calculate total → Iterate all items
- Display cart → Iterate all items (order less important)

**Decision: Map is correct** - Primary operations are by ProductID.

**Why not slice?**
```go
// If we used slice:
type ShoppingCart struct {
    Items []CartItem
}

// Adding/updating quantity requires O(n) search
func (c *ShoppingCart) UpdateQuantity(productID string, qty int) {
    for i := range c.Items {
        if c.Items[i].ProductID == productID {  // Must search!
            c.Items[i].Quantity = qty
            return
        }
    }
    // Not found, add new
    c.Items = append(c.Items, CartItem{ProductID: productID, Quantity: qty})
}

// With map - O(1)
func (c *ShoppingCart) UpdateQuantity(productID string, qty int) {
    item := c.Items[productID]
    item.Quantity = qty
    c.Items[productID] = item
}
```

### System 4: Social Media Feed

```go
type Post struct {
    ID        string
    Author    string
    Content   string
    Timestamp time.Time
    Likes     int
}

type Feed struct {
    Posts     []Post           // Why slice?
    PostIndex map[string]int   // Why map? (Optional optimization)
}
```

**Access patterns:**
- Display feed → Iterate in time order (slice)
- Load more → Append older posts (slice)
- Like post → Find by ID and increment (map helps)
- Delete post → Find by ID and remove (map helps)

**Decision: Hybrid might be useful** - Slice for ordering, map for fast ID lookup.

**Simpler version:**
```go
type Feed struct {
    Posts []Post  // Just slice if ID lookups are rare
}
```

**Trade-off:** If likes/deletes are frequent, hybrid is worth it. If rare, slice alone is simpler.

---

## Common Mistakes

### Mistake 1: Using Slice for Lookups with Large Data

**The mistake:**
```go
type UserService struct {
    Users []User  // 50,000 users
}

// Called on every HTTP request - SLOW!
func (s *UserService) GetUserByID(id int) *User {
    for i := range s.Users {  // Average 25,000 comparisons!
        if s.Users[i].ID == id {
            return &s.Users[i]
        }
    }
    return nil
}
```

**Why it's bad:**
- O(n) lookup on every request
- With 50,000 users: average 25,000 comparisons per lookup
- Scales horribly as users grow
- 100 requests/sec = 2.5 million comparisons/sec

**The fix:**
```go
type UserService struct {
    Users map[int]User  // O(1) lookup
}

func (s *UserService) GetUserByID(id int) (User, bool) {
    user, exists := s.Users[id]
    return user, exists
}
```

**Performance impact:** 25,000x faster for large datasets.

### Mistake 2: Using Map When Order Matters

**The mistake:**
```go
type TodoList struct {
    Tasks map[int]Task
}

// Display tasks - RANDOM ORDER EVERY TIME!
func (t *TodoList) Display() {
    for _, task := range t.Tasks {
        fmt.Println(task.Title)
    }
}

// Output on first run:
// 3. Buy groceries
// 1. Finish report
// 2. Call dentist

// Output on second run:
// 1. Finish report
// 3. Buy groceries
// 2. Call dentist
```

**Why it's bad:**
- Maps have randomized iteration order (intentionally in Go)
- Users expect consistent ordering
- Can't show "task #2" because there's no order

**The fix:**
```go
type TodoList struct {
    Tasks []Task  // Maintains insertion order
}

// OR hybrid if you need ID lookup too:
type TodoList struct {
    Tasks     map[int]Task
    TaskOrder []int
}
```

### Mistake 3: Not Considering Size

**The mistake:**
```go
// Overkill: Using map for small data
type Config struct {
    Settings map[string]string  // Only 5 settings
}

// map overhead: ~10.79 bytes per entry
// Total: 5 * 10.79 = 54 bytes overhead + data
// Slice would be: 24 bytes header + data
```

**When to care:**
- Thousands of small collections (each struct has a small map)
- Embedded systems with limited memory
- Micro-optimizing

**Rule of thumb:**
- < 10 items: Slice is fine even for lookups
- 10-100 items: Depends on frequency
- 100+ items: Map for lookups

### Mistake 4: Storing Duplicate Data

**The mistake:**
```go
type Library struct {
    Books   map[int]Book
    Members map[int]Member
}

type Member struct {
    CheckedOutBooks []Book  // Full book copies!
}

// Problems:
// 1. If book.Title changes, member's copy is stale
// 2. Wastes memory (duplicate book data)
// 3. Two sources of truth - can get out of sync
```

**The fix:**
```go
type Member struct {
    CheckedOutBooks []int  // Just IDs - single source of truth
}

// Look up book when needed
func GetMemberBooks(lib *Library, memberID int) []Book {
    member := lib.Members[memberID]
    books := []Book{}
    for _, bookID := range member.CheckedOutBooks {
        books = append(books, lib.Books[bookID])
    }
    return books
}
```

**Principle:** Store IDs to reference entities, not full copies. Single source of truth.

### Mistake 5: Modifying Map Value Without Reassignment

**The mistake:**
```go
type Library struct {
    Books map[int]Book  // Note: map[int]Book, not map[int]*Book
}

// BUG: This doesn't work!
func CheckoutBook(lib *Library, bookID int) {
    lib.Books[bookID].Available = false  // Compile error!
}

// Error: cannot assign to struct field in map
```

**Why it fails:**
- Maps return VALUES, not references
- You're trying to modify a copy
- Go prevents this to avoid confusion

**The fix:**
```go
func CheckoutBook(lib *Library, bookID int) {
    book := lib.Books[bookID]        // Get copy
    book.Available = false           // Modify copy
    lib.Books[bookID] = book         // Reassign to map
}

// OR use map with pointer:
type Library struct {
    Books map[int]*Book  // Pointer values
}

func CheckoutBook(lib *Library, bookID int) {
    lib.Books[bookID].Available = false  // Works! Modifying through pointer
}
```

### Mistake 6: Forgetting to Initialize Maps

**The mistake:**
```go
type Library struct {
    Books map[int]Book  // nil map!
}

lib := &Library{}
lib.Books[1] = Book{}  // PANIC: assignment to entry in nil map
```

**The fix:**
```go
type Library struct {
    Books map[int]Book
}

// Option 1: Initialize in constructor
func NewLibrary() *Library {
    return &Library{
        Books: make(map[int]Book),
    }
}

// Option 2: Initialize inline (Go 1.11+)
lib := &Library{
    Books: make(map[int]Book),
}
```

**Slices don't have this problem:**
```go
type Library struct {
    Books []Book  // nil slice is usable with append
}

lib := &Library{}
lib.Books = append(lib.Books, book)  // Works fine
```

---

## Benchmarking Guide

### When to Benchmark

**DO benchmark when:**
- Processing large datasets (10,000+ items)
- Operation runs in hot path (thousands of times per second)
- Deciding between structures for critical system
- Premature optimization claims "this is faster" - prove it!

**DON'T benchmark when:**
- Small data (< 100 items)
- Operation runs rarely (once per request)
- Code clarity matters more than nanoseconds
- You're still prototyping

**Rule:** Make it work, make it right, make it fast (in that order).

### Example: Slice vs Map for Member Lookup

Let's benchmark whether `Member.CheckedOutBooks` should be slice or map.

**Code to benchmark:**

```go
package benchmark_test

import (
    "testing"
)

// Slice implementation
type MemberSlice struct {
    CheckedOutBooks []int
}

func (m *MemberSlice) HasBook(bookID int) bool {
    for _, id := range m.CheckedOutBooks {
        if id == bookID {
            return true
        }
    }
    return false
}

// Map implementation
type MemberMap struct {
    CheckedOutBooks map[int]bool
}

func (m *MemberMap) HasBook(bookID int) bool {
    return m.CheckedOutBooks[bookID]
}

// Benchmarks
func BenchmarkSliceHasBook5(b *testing.B) {
    m := &MemberSlice{CheckedOutBooks: []int{1, 2, 3, 4, 5}}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        m.HasBook(3)
    }
}

func BenchmarkMapHasBook5(b *testing.B) {
    m := &MemberMap{CheckedOutBooks: map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        m.HasBook(3)
    }
}

func BenchmarkSliceHasBook50(b *testing.B) {
    books := []int{}
    for i := 1; i <= 50; i++ {
        books = append(books, i)
    }
    m := &MemberSlice{CheckedOutBooks: books}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        m.HasBook(25)
    }
}

func BenchmarkMapHasBook50(b *testing.B) {
    books := make(map[int]bool)
    for i := 1; i <= 50; i++ {
        books[i] = true
    }
    m := &MemberMap{CheckedOutBooks: books}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        m.HasBook(25)
    }
}
```

**Run benchmarks:**
```bash
go test -bench=. -benchmem
```

**Sample results:**
```
BenchmarkSliceHasBook5-8     100000000    10.2 ns/op    0 B/op    0 allocs/op
BenchmarkMapHasBook5-8       50000000     24.5 ns/op    0 B/op    0 allocs/op
BenchmarkSliceHasBook50-8    20000000     82.3 ns/op    0 B/op    0 allocs/op
BenchmarkMapHasBook50-8      50000000     24.8 ns/op    0 B/op    0 allocs/op
```

**Analysis:**
- 5 items: Slice is 2.4x faster (10.2ns vs 24.5ns)
- 50 items: Map is 3.3x faster (24.8ns vs 82.3ns)
- **Crossover point: ~10-15 items**

**Conclusion:** For typical library usage (0-10 books per member), slice is actually faster!

### How to Interpret Results

**Look at:**
1. **ns/op** - Nanoseconds per operation (lower is better)
2. **B/op** - Bytes allocated per operation (lower is better)
3. **allocs/op** - Allocations per operation (lower is better)

**Consider:**
- Is the difference significant? (10ns vs 15ns doesn't matter, 10ns vs 1000ns does)
- How often does this run? (1000x/sec vs 1x/hour)
- Are there allocations? (GC pressure matters at scale)

**Real-world example:**
```
Operation: 50ns
Frequency: 10,000 requests/sec
Impact: 50ns * 10,000 = 0.5ms/sec = 0.05% CPU

vs

Operation: 500ns (10x slower)
Frequency: 10,000 requests/sec
Impact: 500ns * 10,000 = 5ms/sec = 0.5% CPU
```

**When to care:** If CPU usage is > 10% or P99 latency is degraded.

---

## Quick Reference Tables

### Operation Complexity Cheat Sheet

| Operation | Array `[N]T` | Slice `[]T` | Map `map[K]V` |
|-----------|--------------|-------------|---------------|
| Access by key/index | O(1) | O(1) | O(1) avg |
| Search by value | O(n) | O(n) | O(n) |
| Insert at end | N/A | O(1) amortized | O(1) avg |
| Insert at start | N/A | O(n) | N/A |
| Insert at position | N/A | O(n) | N/A |
| Delete by key/index | N/A | O(n) | O(1) avg |
| Maintain order | Yes | Yes | No |
| Allows duplicates | Yes | Yes | No (keys unique) |
| Memory overhead | None | 24B header | ~10.79B per entry |
| Zero value usable | Yes | Yes (nil ok for append) | No (must make()) |

### Decision Matrix

| If you need... | Use this |
|----------------|----------|
| Fast lookup by unique key | `map[K]V` |
| Maintain insertion/custom order | `[]T` |
| Both lookup and order | `map[K]V` + `[]K` (hybrid) |
| No duplicates (set) | `map[T]bool` or `map[T]struct{}` |
| Count occurrences | `map[T]int` |
| Small collection (< 10 items) with occasional lookups | `[]T` (simpler) |
| Large collection (100+) with frequent lookups | `map[K]V` |
| Fixed size known at compile time | `[N]T` |
| Stack (LIFO) | `[]T` (append/slice from end) |
| Queue (FIFO) | `[]T` (append/slice from start) |
| Graph adjacency | `map[Node][]Node` |
| Cache | `map[Key]Value` |

### Size Thresholds

| Collection Size | Recommended Structure for Lookups |
|-----------------|-----------------------------------|
| 0-10 items | Slice (O(n) is acceptable) |
| 10-50 items | Depends: frequent lookups → map, rare → slice |
| 50-100 items | Map if lookups are common |
| 100-1,000 items | Map for any keyed lookups |
| 1,000+ items | Definitely map (or database) |

### When to Use Each Structure

| Structure | Primary Use Case | Example |
|-----------|------------------|---------|
| `[N]T` | Fixed-size data | `type IPv4 [4]byte` |
| `[]T` | Ordered, dynamic list | `Tasks []Task` |
| `map[K]V` | Keyed lookup | `Users map[int]User` |
| `map[K]bool` | Set (unique items) | `Visited map[string]bool` |
| `map[K]int` | Counting | `WordCount map[string]int` |
| `map[K][]V` | One-to-many | `Graph map[Node][]Node` |
| `map[K]V` + `[]K` | Lookup + order | `Playlist` (see hybrid pattern) |

---

## Practice Exercises

Test your understanding with these scenarios. Try to answer before revealing the solution.

### Exercise 1: User Session Store

**Scenario:** Build a session store that:
- Stores user sessions by session ID
- Expires sessions after 30 minutes
- Looks up sessions on every request

**Question:** What structure(s) would you use?

<details>
<summary>Answer</summary>

**Structure:**
```go
type SessionStore struct {
    Sessions map[string]Session  // sessionID -> Session
}

type Session struct {
    UserID    int
    ExpiresAt time.Time
}
```

**Reasoning:**
- Primary operation: Lookup session by ID on every request → **Map**
- Size: Could be thousands of sessions → Map scales well
- Order doesn't matter
- Need fast O(1) lookup for performance

**Why not slice?**
- O(n) lookup on every request would be too slow
- No need for ordering

</details>

### Exercise 2: Leaderboard System

**Scenario:** Build a game leaderboard that:
- Stores player scores
- Displays top 10 players (sorted by score)
- Updates scores frequently
- Checks if player is on leaderboard

**Question:** What structure(s) would you use?

<details>
<summary>Answer</summary>

**Structure:**
```go
type Leaderboard struct {
    Scores map[string]int  // playerName -> score
}

// Display top 10
func (l *Leaderboard) TopTen() []Player {
    players := []Player{}
    for name, score := range l.Scores {
        players = append(players, Player{name, score})
    }
    sort.Slice(players, func(i, j int) bool {
        return players[i].Score > players[j].Score
    })
    if len(players) > 10 {
        players = players[:10]
    }
    return players
}
```

**Reasoning:**
- Update score: Need fast lookup by player name → **Map**
- Display top 10: Must sort anyway, so map's lack of order doesn't hurt
- Check if player exists: O(1) with map

**Alternative approach** (if updates are rare but display is frequent):
```go
type Leaderboard struct {
    Scores       map[string]int
    SortedScores []Player  // Pre-sorted, updated on score change
}
```

</details>

### Exercise 3: Browser History

**Scenario:** Build a browser history that:
- Tracks visited URLs in order
- Allows "back" button (go to previous URL)
- Prevents duplicate consecutive URLs
- Displays history in reverse chronological order

**Question:** What structure(s) would you use?

<details>
<summary>Answer</summary>

**Structure:**
```go
type BrowserHistory struct {
    URLs []string  // Ordered list of URLs
}

// Visit URL (prevents consecutive duplicates)
func (b *BrowserHistory) Visit(url string) {
    if len(b.URLs) > 0 && b.URLs[len(b.URLs)-1] == url {
        return  // Skip duplicate
    }
    b.URLs = append(b.URLs, url)
}

// Back button
func (b *BrowserHistory) Back() string {
    if len(b.URLs) < 2 {
        return ""
    }
    b.URLs = b.URLs[:len(b.URLs)-1]  // Remove current
    return b.URLs[len(b.URLs)-1]     // Return previous
}

// Display history (reverse order)
func (b *BrowserHistory) Display() []string {
    reversed := make([]string, len(b.URLs))
    for i, url := range b.URLs {
        reversed[len(b.URLs)-1-i] = url
    }
    return reversed
}
```

**Reasoning:**
- Primary need: Maintain chronological order → **Slice**
- Back button: Need to remove last item → Slice supports this
- Display: Need ordered access → Slice
- No need for fast lookup by URL

**Why not map?**
- Order is essential
- No keyed lookups needed

</details>

### Exercise 4: Tag System

**Scenario:** Build a blog post tagging system where:
- Each post can have multiple tags
- Need to find all posts with a specific tag
- Need to list all tags used across all posts

**Question:** What structure(s) would you use?

<details>
<summary>Answer</summary>

**Structure:**
```go
type Post struct {
    ID    int
    Title string
    Tags  []string  // Small list per post
}

type Blog struct {
    Posts       map[int]Post      // postID -> Post
    TagIndex    map[string][]int  // tag -> []postIDs
}

// Add post
func (b *Blog) AddPost(post Post) {
    b.Posts[post.ID] = post

    // Update tag index
    for _, tag := range post.Tags {
        b.TagIndex[tag] = append(b.TagIndex[tag], post.ID)
    }
}

// Find posts by tag - O(1) to get IDs, O(n) to lookup posts
func (b *Blog) FindByTag(tag string) []Post {
    postIDs := b.TagIndex[tag]
    posts := []Post{}
    for _, id := range postIDs {
        posts = append(posts, b.Posts[id])
    }
    return posts
}

// List all tags - O(n) where n = number of unique tags
func (b *Blog) AllTags() []string {
    tags := []string{}
    for tag := range b.TagIndex {
        tags = append(tags, tag)
    }
    return tags
}
```

**Reasoning:**
- Posts: Need lookup by ID → `map[int]Post`
- Tags per post: Small list (3-10 tags) → `[]string`
- Tag search: Need "find all posts with tag X" → Create index: `map[string][]int`
- This is a **multi-index pattern** (one primary map + search indexes)

</details>

### Exercise 5: Undo/Redo System

**Scenario:** Build an undo/redo system for a text editor:
- Track changes (insertions, deletions)
- Undo reverts last change
- Redo reapplies undone change
- Can undo up to 100 changes

**Question:** What structure(s) would you use?

<details>
<summary>Answer</summary>

**Structure:**
```go
type Change struct {
    Type      string  // "insert" or "delete"
    Position  int
    Text      string
}

type UndoRedoSystem struct {
    UndoStack []Change  // Stack of changes
    RedoStack []Change  // Stack of undone changes
}

// Record change
func (u *UndoRedoSystem) RecordChange(change Change) {
    u.UndoStack = append(u.UndoStack, change)
    u.RedoStack = nil  // Clear redo on new change

    // Limit to 100 changes
    if len(u.UndoStack) > 100 {
        u.UndoStack = u.UndoStack[1:]
    }
}

// Undo - O(1)
func (u *UndoRedoSystem) Undo() (Change, bool) {
    if len(u.UndoStack) == 0 {
        return Change{}, false
    }
    n := len(u.UndoStack) - 1
    change := u.UndoStack[n]
    u.UndoStack = u.UndoStack[:n]
    u.RedoStack = append(u.RedoStack, change)
    return change, true
}

// Redo - O(1)
func (u *UndoRedoSystem) Redo() (Change, bool) {
    if len(u.RedoStack) == 0 {
        return Change{}, false
    }
    n := len(u.RedoStack) - 1
    change := u.RedoStack[n]
    u.RedoStack = u.RedoStack[:n]
    u.UndoStack = append(u.UndoStack, change)
    return change, true
}
```

**Reasoning:**
- Need LIFO (Last In First Out) behavior → **Stack pattern with slice**
- Operations: push/pop from end → Slice is perfect (O(1))
- Order critical (must undo in reverse order)
- No need for keyed lookup

**Why not map?**
- Stacks require ordering
- No need for random access

</details>

---

## Conclusion

**Key Principles:**

1. **Access patterns drive structure choice** - How you use data matters more than what data is
2. **Size matters** - Small collections can use simpler structures
3. **Order vs lookup** - Different structures excel at different operations
4. **Don't over-engineer** - Use the simplest structure that works
5. **Measure when it matters** - Benchmark critical paths, not everything
6. **Hybrid when necessary** - Combine structures when you need multiple access patterns

**The Decision Process:**

```
1. Identify primary operation (lookup? iterate? search?)
2. Estimate size (10 items? 10,000 items?)
3. Consider order requirements (does sequence matter?)
4. Check secondary operations (what else do you need?)
5. Choose simplest structure that meets requirements
6. Benchmark if performance is critical
```

**Remember:** Good architectural decisions come from understanding trade-offs, not memorizing rules. Every structure has strengths and weaknesses. Your job is to match the structure to your specific access patterns and constraints.

Now when you design systems, you won't just "make it work" - you'll make intentional, informed decisions that scale.

---

**Further Reading:**

- Go Blog: [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- Go Blog: [Go maps in action](https://go.dev/blog/maps)
- Effective Go: [Data](https://go.dev/doc/effective_go#data)
- Your exercise: `02-data-structures/11_data_modeling/` - Review your design decisions

**Next Steps:**

1. Review your Exercise 11 code with this guide in mind
2. Try the practice exercises above
3. In your next project, document WHY you chose each structure
4. Benchmark one of your hot paths to see actual performance
