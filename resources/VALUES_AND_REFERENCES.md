# Go Values and References: A Practical Guide

**Why this guide exists:** You just hit the classic "cannot modify map values directly" issue in Exercise 11. This happens because Go's value/reference semantics work differently than most languages. This guide will fix your mental model.

**What you tried:**
```go
// Exercise 11, line 96 - This doesn't work!
lib.Members[memberID].CheckedOutBooks = append(...)  // COMPILE ERROR
```

**What you're thinking:** "Why can't I just modify the field? The member is right there in the map!"

**Reality:** Map values are not directly modifiable. You get a **copy** when you access them.

Let's build the correct mental model from the ground up.

---

## Table of Contents

1. [Core Principle: Everything is Pass-by-Value](#core-principle-everything-is-pass-by-value)
2. [Type Categories](#type-categories)
3. [The Map Value Problem (Your Current Issue)](#the-map-value-problem)
4. [Nested Map Values: The Two-Level Challenge](#nested-map-values-the-two-level-challenge)
5. [Understanding Memory: Box and Address Model](#understanding-memory-box-and-address-model)
6. [Common Scenarios](#common-scenarios)
7. [Decision Tree: When Do I Need a Pointer?](#decision-tree-when-do-i-need-a-pointer)
8. [Practical Patterns](#practical-patterns)
9. [Common Mistakes](#common-mistakes)
10. [Quick Reference](#quick-reference)

---

## Core Principle: Everything is Pass-by-Value

### The Golden Rule

**In Go, EVERYTHING is passed by value. Always. No exceptions.**

```go
func foo(x int) {
    x = 42  // Modifies the COPY, not the original
}

func main() {
    n := 10
    foo(n)
    fmt.Println(n)  // Still 10, not 42
}
```

When you pass `n` to `foo`, Go **copies** the value of `n` into parameter `x`. Changes to `x` don't affect `n`.

### But Wait... Some Types *Do* Modify the Original?

Yes! But here's the key insight:

**Some types contain ADDRESSES (pointers) internally, so when you copy them, you're copying the address, not the data.**

```go
func appendToSlice(s []int) {
    s = append(s, 99)  // Might or might not affect original - DANGEROUS
}

func modifySliceElement(s []int) {
    s[0] = 99  // DOES affect original - safe pattern
}

func main() {
    nums := []int{1, 2, 3}
    modifySliceElement(nums)
    fmt.Println(nums)  // [99, 2, 3] - element modified

    appendToSlice(nums)
    fmt.Println(nums)  // [99, 2, 3] - append didn't affect original!
}
```

Why the difference?
- A slice is a **struct containing a pointer** to the underlying array
- When you copy the slice, you copy the pointer (so you point to the same array)
- Modifying elements works (same array)
- But `append` might create a NEW array if capacity is exceeded, and that new array pointer doesn't get back to the caller

---

## Type Categories

Go types fall into two mental buckets:

### Value Types (Copy the Data)

When you pass these, you copy **all the data**:

| Type       | What Gets Copied                    | Size               |
|------------|-------------------------------------|--------------------|
| `int`      | The number                          | 8 bytes (on 64-bit)|
| `bool`     | true/false                          | 1 byte             |
| `string`   | Pointer + length (immutable)        | 16 bytes           |
| `struct`   | All fields                          | Sum of field sizes |
| `array`    | All elements                        | element_size × len |

**Key insight:** Structs and arrays copy ALL their contents. A `[1000]int` array is 8000 bytes of copying!

### Reference Types (Copy the Address)

When you pass these, you copy a **descriptor that contains pointers**:

| Type        | What Gets Copied                           | Points To             |
|-------------|--------------------------------------------|-----------------------|
| `slice`     | pointer + len + cap (24 bytes)             | Underlying array      |
| `map`       | pointer to map structure (8 bytes)         | Hash table in heap    |
| `chan`      | pointer to channel structure (8 bytes)     | Channel buffer        |
| `*T`        | The address (8 bytes)                      | The actual T value    |
| `interface` | type descriptor + pointer (16 bytes)       | Underlying value      |

**Key insight:** These are sometimes called "reference types" but that's misleading. They're still passed by value - it's just that the value being copied contains a pointer.

### The String Exception

Strings are weird:

```go
s := "hello"
```

Internally, a string is:
```
struct {
    ptr *byte  // Pointer to character data
    len int    // Length
}
```

BUT strings are **immutable**. You can never modify the underlying bytes. So even though they contain a pointer, they behave like value types for practical purposes.

---

## The Map Value Problem

### Your Exercise 11 Issue (Line 96)

Here's what you tried:

```go
// data_modeling.go, CheckoutBook function
member, exists := lib.Members[memberID]  // Gets a COPY of the Member
book, expected := lib.Books[bookID]      // Gets a COPY of the Book

if exists && expected && book.Available {
    book.Available = false                                        // Modifies the copy
    member.CheckedOutBooks = append(member.CheckedOutBooks, bookID)  // Modifies the copy

    // YOU MUST PUT THE COPIES BACK!
    lib.Members[memberID] = member  // ✅ Reassign the modified copy
    lib.Books[bookID] = book        // ✅ Reassign the modified copy
    return true
}
```

**Why you can't do `lib.Members[memberID].CheckedOutBooks = ...` directly:**

When you write `lib.Members[memberID]`, Go:
1. Looks up the key in the map
2. **Copies** the struct value out of the map
3. Returns that copy to you

The struct sitting in the map is **not addressable**. You can't take its address, and you can't modify it in place.

### Visual Model: What's Happening

```
Map memory layout:
┌─────────────────────────────────────┐
│  Members map                        │
│  ┌───────────────────────────────┐  │
│  │ Key: 1 → Value: Member{...}   │  │ ← This is stored directly in map
│  └───────────────────────────────┘  │
└─────────────────────────────────────┘

When you do: member := lib.Members[1]
┌─────────────────────────────────────┐
│  Members map                        │
│  ┌───────────────────────────────┐  │
│  │ Key: 1 → Value: Member{...}   │  │ ← Original stays here
│  └───────────────────────────────┘  │
└─────────────────────────────────────┘
                    │
                    │ COPY operation
                    ↓
              ┌─────────────┐
              │   member    │ ← Copy on stack
              │  Member{..} │
              └─────────────┘

Modifications happen to the COPY, not the original in the map!
```

### The Retrieve-Modify-Reassign Pattern

This is the correct pattern for modifying structs in maps:

```go
// 1. RETRIEVE - Get a copy
member := lib.Members[memberID]

// 2. MODIFY - Change the copy
member.CheckedOutBooks = append(member.CheckedOutBooks, bookID)

// 3. REASSIGN - Put the copy back in the map
lib.Members[memberID] = member
```

This is what you correctly did in Exercise 11!

### Alternative: Use Map of Pointers

If you're doing LOTS of modifications, consider `map[K]*V` instead of `map[K]V`:

```go
// Instead of:
type Library struct {
    Members map[int]Member  // map[int]Member
}

// Use:
type Library struct {
    Members map[int]*Member  // map[int]*Member
}

// Now you can modify directly:
lib.Members[memberID].CheckedOutBooks = append(...)  // ✅ Works!
```

**Why this works:**

```
Map of pointers:
┌─────────────────────────────────────┐
│  Members map                        │
│  ┌───────────────────────────────┐  │
│  │ Key: 1 → Value: *Member       │  │ ← Map stores POINTER (8 bytes)
│  │              │                │  │
│  │              └────────────────┼──┼─→ ┌──────────────┐
│  └───────────────────────────────┘  │   │ Member{...}  │ ← Actual struct on heap
└─────────────────────────────────────┘   └──────────────┘

When you do: lib.Members[1].CheckedOutBooks = append(...)
1. Get the pointer from map (copies 8 bytes)
2. Follow the pointer to the actual struct (on heap)
3. Modify the struct directly (no copy needed!)
```

### Trade-offs: `map[K]V` vs `map[K]*V`

| Aspect            | `map[K]V`                              | `map[K]*V`                                  |
|-------------------|----------------------------------------|---------------------------------------------|
| **Syntax**        | Retrieve-modify-reassign               | Direct modification                         |
| **Memory**        | Values inline in map                   | Pointers in map, values on heap             |
| **Cache**         | Better (values together)               | Worse (pointer chasing)                     |
| **Allocation**    | Less GC pressure                       | More GC pressure (heap allocs)              |
| **Nil values**    | Can't have nil, always zero value      | Can have nil entries (extra state)          |
| **When to use**   | Small structs, infrequent modification | Large structs, frequent modification        |

**Rule of thumb:**
- Small structs (< 5 fields, < 40 bytes): Use `map[K]V` with retrieve-modify-reassign
- Large structs or frequent modifications: Use `map[K]*V`

### Common Gotcha: Iterating and Modifying

```go
// ❌ WRONG - Tries to modify the copy from range
for id, member := range lib.Members {
    member.Name = "Modified"  // Modifies the copy, doesn't affect map!
}

// ✅ CORRECT - Retrieve-modify-reassign
for id := range lib.Members {
    member := lib.Members[id]      // Get copy
    member.Name = "Modified"       // Modify copy
    lib.Members[id] = member       // Put back
}

// ✅ ALSO CORRECT - Use keys only
for id := range lib.Members {
    m := lib.Members[id]
    m.Name = "Modified"
    lib.Members[id] = m
}
```

---

## Nested Map Values: The Two-Level Challenge

### The Problem Amplified

If modifying a single-level map value is tricky, **nested** map values are where most developers get stuck. This is your Exercise 11.5 scenario:

```go
type Library struct {
    Users map[string]User  // Level 1: User is a VALUE type
}

type User struct {
    Playlists map[string]Playlist  // Level 2: Playlist is a VALUE type
}
```

**The Challenge:** You have **maps inside maps**, both storing values (not pointers). This requires TWO retrieve-modify-reassign operations!

### Your Exact Bug from Exercise 11.5

```go
func (lib *Library) AddSongToPlaylist(userID, playlistID, songID string) error {
    foundSong := lib.Songs[songID]
    foundUser := lib.Users[userID]                    // ✅ Get User copy
    foundPlaylist := lib.Users[userID].Playlists[playlistID]  // ❌ BUG!
    //               ↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑
    // This gets Playlist from ORIGINAL map, not from foundUser copy!

    foundPlaylist.Songs = append(foundPlaylist.Songs, foundSong)
    foundUser.Playlists[playlistID] = foundPlaylist  // Put modified playlist in foundUser
    lib.Users[userID] = foundUser                    // Put foundUser back in map

    // Result: foundUser contains a DIFFERENT playlist than the one you modified!
    // The playlist you appended to is not the same one you put in foundUser.
    return nil
}
```

**Why This Fails:**

1. `foundUser := lib.Users[userID]` creates a **COPY** of User (call it Copy A)
2. `lib.Users[userID].Playlists[playlistID]` accesses the **ORIGINAL** User in the map (not Copy A!), and gets its Playlist
3. You modify that Playlist and put it in Copy A's Playlists map
4. But Copy A's Playlists already had a different copy of that Playlist!
5. You're mixing copies from different sources = data corruption

### The Correct Pattern

```go
func (lib *Library) AddSongToPlaylist(userID, playlistID, songID string) error {
    foundSong := lib.Songs[songID]

    // Step 1: Get outer struct (User) from outer map
    foundUser := lib.Users[userID]                // Get User COPY

    // Step 2: Get inner struct (Playlist) from COPY's map (not original!)
    foundPlaylist := foundUser.Playlists[playlistID]  // ✅ Get from foundUser, not lib.Users[userID]!

    // Step 3: Modify inner struct
    foundPlaylist.Songs = append(foundPlaylist.Songs, foundSong)

    // Step 4: Put inner struct back into outer struct's map
    foundUser.Playlists[playlistID] = foundPlaylist

    // Step 5: Put outer struct back into outer map
    lib.Users[userID] = foundUser

    return nil
}
```

**The Key Rule:** Always get nested values from the COPY, never from the original map!

### Visual: Why Getting from Original Breaks

```
Step 1: foundUser := lib.Users[userID]

lib.Users (original map)           foundUser (local copy)
┌──────────────────┐               ┌──────────────────┐
│ Key: "alice"     │               │ Copy of User     │
│ Value: User{     │     COPY      │   Playlists:     │
│   Playlists: map───────────────→ │     (reference   │
│     "fav": Play1 │               │      to SAME map)│
│ }                │               │ }                │
└──────────────────┘               └──────────────────┘

Both foundUser and lib.Users["alice"] point to the SAME Playlists map (it's a reference type).


Step 2A: ❌ WRONG - Get from original
foundPlaylist := lib.Users[userID].Playlists[playlistID]

This accesses lib.Users again, getting a FRESH copy of User, then its Playlist.
This is a DIFFERENT copy than foundUser!

┌──────────────────┐               ┌──────────────────┐               ┌──────────────────┐
│ lib.Users map    │               │ foundUser (copy) │               │ foundPlaylist    │
│   "alice": User  │───fresh───→   │   (Copy A)       │               │   (from NEW copy)│
│     Playlists─────copy of User   │                  │               │   NOT from       │
│       "fav": P1  │               │                  │               │   foundUser!     │
└──────────────────┘               └──────────────────┘               └──────────────────┘


Step 2B: ✅ CORRECT - Get from foundUser copy
foundPlaylist := foundUser.Playlists[playlistID]

This gets Playlist from the copy we already have.

┌──────────────────┐               ┌──────────────────┐
│ lib.Users map    │               │ foundUser (copy) │
│   "alice": User  │               │   Playlists:     │
│     Playlists    │               │     "fav": P1────┼──→ foundPlaylist
└──────────────────┘               └──────────────────┘
                                         ↑
                                         └──── We get Playlist from THIS copy
```

### Memory Trace: Wrong vs Right

**WRONG WAY (your original bug):**

```go
foundUser := lib.Users[userID]                        // Copy A of User
foundPlaylist := lib.Users[userID].Playlists[playlistID]  // Gets from Copy B (different copy!)

// Memory state:
// Copy A (foundUser):      User{Playlists: map{"favorites": Playlist{Songs: [Song1]}}}
// Copy B (from line 2):    User{Playlists: map{"favorites": Playlist{Songs: [Song1]}}}
// foundPlaylist:           Playlist{Songs: [Song1]} (from Copy B)

foundPlaylist.Songs = append(foundPlaylist.Songs, Song2)
// foundPlaylist:           Playlist{Songs: [Song1, Song2]}

foundUser.Playlists[playlistID] = foundPlaylist
// Copy A (foundUser):      User{Playlists: map{"favorites": Playlist{Songs: [Song1, Song2]}}}
// But this Playlist came from Copy B! Copy A originally had its own Playlist instance.

lib.Users[userID] = foundUser
// lib.Users["alice"]:      User{Playlists: map{"favorites": Playlist{Songs: [Song1, Song2]}}}

// Tests read from lib.Users["alice"].Playlists["favorites"].Songs
// Result: ✅ Actually works! But...wait, why did your tests fail then?
```

**Wait - Let me re-examine your actual bug...**

Actually, looking more carefully at the code, the issue is more subtle. When you do:
```go
foundUser := lib.Users[userID]
```

You get a copy of User. But User.Playlists is a `map[string]Playlist`. Maps are reference types, so both `foundUser.Playlists` and `lib.Users[userID].Playlists` point to the **SAME underlying map**!

So actually, doing:
```go
foundUser.Playlists[playlistID] = modifiedPlaylist
```

**DOES** modify the original map! Because maps are reference types.

But you still need to reassign `lib.Users[userID] = foundUser` if you want to modify other fields of User (like if User had a `TotalPlaylists int` counter).

**The REAL reason your tests were failing:**

Let me check... If both copies share the same Playlists map, then modifications SHOULD work... unless you're reading the Playlist wrong in your getter functions!

### The Nested Pattern Rule

Regardless of the subtlety above, **the safe pattern for nested maps is always:**

```
Get outer struct copy
Get inner struct from OUTER COPY (not from original map)
Modify inner struct
Reassign inner to outer copy's map
Reassign outer copy to original map
```

**Number of reassignments = Number of value-type nesting levels**

- `map[K]V`: 1 reassignment
- `map[K]V` where `V` has `map[K2]V2` field: 2 reassignments
- `map[K]V` where `V` has `map[K2]V2` where `V2` has `map[K3]V3`: 3 reassignments (getting messy - use pointers!)

### When Nested Maps Work Differently

**Important Distinction:**

```go
// Case 1: Struct with map field
type User struct {
    Playlists map[string]Playlist  // ← Map is reference type
}
users := map[string]User           // ← User is value type

user := users["alice"]             // Copy of User
user.Playlists["new"] = Playlist{} // ← This DOES modify the original map!
// Because Playlists is a map (reference), both copies point to same underlying map
```

vs

```go
// Case 2: Map of maps
data := map[string]map[string]int{
    "alice": {"score": 10},
}

inner := data["alice"]             // Reference to same map
inner["score"] = 20                // ← This DOES modify original!
// Maps are reference types, so inner and data["alice"] point to same map
```

vs

```go
// Case 3: Map of structs where struct has value fields
type Profile struct {
    Score int  // ← Value type field
}
profiles := map[string]Profile

profile := profiles["alice"]       // Copy of Profile
profile.Score = 20                 // ← Modifies the COPY
profiles["alice"] = profile        // ← MUST reassign!
```

**The Pattern:** If the field you're modifying is a value type (int, struct, array), you need reassignment. If it's a reference type (map, slice, pointer), modifications work directly, BUT you still might need reassignment if other fields changed.

### Practice Exercise

Fix this bug:

```go
type Company struct {
    Departments map[string]Department
}
type Department struct {
    Employees []string
}
companies := map[string]Company{
    "Acme": {Departments: map[string]Department{
        "Engineering": {Employees: []string{"Alice"}},
    }},
}

// BUG: Add "Bob" to Engineering department
company := companies["Acme"]
dept := companies["Acme"].Departments["Engineering"]  // ❌ Gets from original, not company copy!
dept.Employees = append(dept.Employees, "Bob")
company.Departments["Engineering"] = dept
companies["Acme"] = company
```

**Fix:**
```go
company := companies["Acme"]
dept := company.Departments["Engineering"]  // ✅ Get from company copy!
dept.Employees = append(dept.Employees, "Bob")
company.Departments["Engineering"] = dept
companies["Acme"] = company
```

---

## Understanding Memory: Box and Address Model

### The Mental Model

Think of memory as a warehouse with numbered storage boxes:

```
Memory (Warehouse):
┌──────────┬──────────┬──────────┬──────────┬──────────┐
│  Box 100 │  Box 101 │  Box 102 │  Box 103 │  Box 104 │
│  ┌────┐  │  ┌────┐  │  ┌────┐  │  ┌────┐  │  ┌────┐  │
│  │ 42 │  │  │ 17 │  │  │ 99 │  │  │ 8  │  │  │ ...│  │
│  └────┘  │  └────┘  │  └────┘  │  └────┘  │  └────┘  │
└──────────┴──────────┴──────────┴──────────┴──────────┘
  Address    Address    Address    Address    Address
  0x100      0x101      0x102      0x103      0x104
```

### Value Types: Copy the Contents

```go
x := 42          // Box 100 contains: 42
y := x           // Box 101 contains: 42 (COPIED from box 100)
y = 99           // Box 101 now contains: 99
fmt.Println(x)   // 42 (box 100 unchanged)
```

Visual:
```
Before y = 99:              After y = 99:
┌──────────┬──────────┐    ┌──────────┬──────────┐
│  Box 100 │  Box 101 │    │  Box 100 │  Box 101 │
│  ┌────┐  │  ┌────┐  │    │  ┌────┐  │  ┌────┐  │
│  │ 42 │  │  │ 42 │  │    │  │ 42 │  │  │ 99 │  │
│  └────┘  │  └────┘  │    │  └────┘  │  └────┘  │
└──────────┴──────────┘    └──────────┴──────────┘
    x          y                x          y
```

### Pointers: Store the Address

```go
x := 42          // Box 100 contains: 42
p := &x          // Box 101 contains: "address 100" (not the value 42!)
*p = 99          // Follow address 100, change its contents to 99
fmt.Println(x)   // 99 (box 100 was modified through pointer)
```

Visual:
```
Before *p = 99:              After *p = 99:
┌──────────┬──────────┐    ┌──────────┬──────────┐
│  Box 100 │  Box 101 │    │  Box 100 │  Box 101 │
│  ┌────┐  │  ┌────┐  │    │  ┌────┐  │  ┌────┐  │
│  │ 42 │  │  │0x100│  │    │  │ 99 │  │  │0x100│  │
│  └────┘  │  └────┘  │    │  └────┘  │  └────┘  │
└──────────┴──────────┘    └──────────┴──────────┘
    x          p                x          p
               │                           │
               └──→ "points to x"          └──→ "points to x"
```

### Structs: Multiple Boxes Together

```go
type Point struct {
    X int
    Y int
}

p1 := Point{10, 20}  // Boxes 100-101: [10][20]
p2 := p1             // Boxes 102-103: [10][20] (ENTIRE struct copied)
p2.X = 30            // Only changes box 102
fmt.Println(p1.X)    // 10 (box 100 unchanged)
```

Visual:
```
After p2 := p1:
┌──────────┬──────────┐  ┌──────────┬──────────┐
│  Box 100 │  Box 101 │  │  Box 102 │  Box 103 │
│  ┌────┐  │  ┌────┐  │  │  ┌────┐  │  ┌────┐  │
│  │ 10 │  │  │ 20 │  │  │  │ 10 │  │  │ 20 │  │
│  └────┘  │  └────┘  │  │  └────┘  │  └────┘  │
└──────────┴──────────┘  └──────────┴──────────┘
       p1                       p2
    (2 boxes)                (2 boxes)

After p2.X = 30:
┌──────────┬──────────┐  ┌──────────┬──────────┐
│  Box 100 │  Box 101 │  │  Box 102 │  Box 103 │
│  ┌────┐  │  ┌────┐  │  │  ┌────┐  │  ┌────┐  │
│  │ 10 │  │  │ 20 │  │  │  │ 30 │  │  │ 20 │  │
│  └────┘  │  └────┘  │  │  └────┘  │  └────┘  │
└──────────┴──────────┘  └──────────┴──────────┘
       p1                       p2
```

### Pointer to Struct: Share the Boxes

```go
p1 := &Point{10, 20}  // Box 100: pointer to boxes 200-201
p2 := p1              // Box 101: copy of pointer (still points to 200-201)
p2.X = 30             // Follow pointer, change box 200
fmt.Println(p1.X)     // 30 (both pointers point to same struct)
```

Visual:
```
p1 := &Point{10, 20}:
┌──────────┐                    ┌──────────┬──────────┐
│  Box 100 │                    │  Box 200 │  Box 201 │
│  ┌────┐  │                    │  ┌────┐  │  ┌────┐  │
│  │0x200│  │ ──────────────→   │  │ 10 │  │  │ 20 │  │
│  └────┘  │                    │  └────┘  │  └────┘  │
└──────────┘                    └──────────┴──────────┘
    p1                              Point struct

p2 := p1:
┌──────────┬──────────┐         ┌──────────┬──────────┐
│  Box 100 │  Box 101 │         │  Box 200 │  Box 201 │
│  ┌────┐  │  ┌────┐  │         │  ┌────┐  │  ┌────┐  │
│  │0x200│  │  │0x200│  │  ─────→  │  10 │  │  │ 20 │  │
│  └────┘  │  └────┘  │  ─────→  │  └────┘  │  └────┘  │
└──────────┴──────────┘         └──────────┴──────────┘
    p1         p2                   Point struct
                                 (SHARED by both pointers)

p2.X = 30:
┌──────────┬──────────┐         ┌──────────┬──────────┐
│  Box 100 │  Box 101 │         │  Box 200 │  Box 201 │
│  ┌────┐  │  ┌────┐  │         │  ┌────┐  │  ┌────┐  │
│  │0x200│  │  │0x200│  │  ─────→  │  30 │  │  │ 20 │  │ ← Changed!
│  └────┘  │  └────┘  │  ─────→  │  └────┘  │  └────┘  │
└──────────┴──────────┘         └──────────┴──────────┘
    p1         p2                   Point struct
```

### Slices: The Shared Array Model

```go
s1 := []int{1, 2, 3}  // Slice header points to array [1, 2, 3]
s2 := s1              // Copy slice header (still points to SAME array)
s2[0] = 99            // Modifies the shared array
fmt.Println(s1[0])    // 99 (both slices see the change)
```

Visual:
```
s1 := []int{1, 2, 3}:
┌─────────────────┐
│   Slice s1      │
│  ┌───────────┐  │
│  │ ptr: 0x300│  │ ────┐
│  │ len: 3    │  │     │
│  │ cap: 3    │  │     │
│  └───────────┘  │     │
└─────────────────┘     │
                        │
                        ↓
             ┌──────────┬──────────┬──────────┐
             │  Box 300 │  Box 301 │  Box 302 │
  Underlying │  ┌────┐  │  ┌────┐  │  ┌────┐  │
  array      │  │ 1  │  │  │ 2  │  │  │ 3  │  │
             │  └────┘  │  └────┘  │  └────┘  │
             └──────────┴──────────┴──────────┘

s2 := s1 (copies the header, NOT the array):
┌─────────────────┐  ┌─────────────────┐
│   Slice s1      │  │   Slice s2      │
│  ┌───────────┐  │  │  ┌───────────┐  │
│  │ ptr: 0x300│  │  │  │ ptr: 0x300│  │ ← SAME pointer!
│  │ len: 3    │  │  │  │ len: 3    │  │
│  │ cap: 3    │  │  │  │ cap: 3    │  │
│  └───────────┘  │  │  └───────────┘  │
└─────────────────┘  └─────────────────┘
         │                    │
         └──────────┬─────────┘
                    ↓
             ┌──────────┬──────────┬──────────┐
  Underlying │  Box 300 │  Box 301 │  Box 302 │
  array      │  ┌────┐  │  ┌────┐  │  ┌────┐  │
  (SHARED!)  │  │ 1  │  │  │ 2  │  │  │ 3  │  │
             │  └────┘  │  └────┘  │  └────┘  │
             └──────────┴──────────┴──────────┘

s2[0] = 99 (modifies shared array):
┌─────────────────┐  ┌─────────────────┐
│   Slice s1      │  │   Slice s2      │
│  ┌───────────┐  │  │  ┌───────────┐  │
│  │ ptr: 0x300│  │  │  │ ptr: 0x300│  │
│  │ len: 3    │  │  │  │ len: 3    │  │
│  │ cap: 3    │  │  │  │ cap: 3    │  │
│  └───────────┘  │  │  └───────────┘  │
└─────────────────┘  └─────────────────┘
         │                    │
         └──────────┬─────────┘
                    ↓
             ┌──────────┬──────────┬──────────┐
  Underlying │  Box 300 │  Box 301 │  Box 302 │
  array      │  ┌────┐  │  ┌────┐  │  ┌────┐  │
             │  │ 99 │  │  │ 2  │  │  │ 3  │  │ ← Changed!
             │  └────┘  │  └────┘  │  └────┘  │
             └──────────┴──────────┴──────────┘
```

Key insight: `s1` and `s2` are **different slice headers** but point to the **same underlying array**.

---

## Common Scenarios

### Scenario 1: Modifying Function Parameters

#### Problem: Changes Don't Persist

```go
type Config struct {
    Host string
    Port int
}

func setDefaults(cfg Config) {
    cfg.Host = "localhost"  // Modifies the COPY
    cfg.Port = 8080
}

func main() {
    config := Config{}
    setDefaults(config)
    fmt.Println(config.Host)  // "" (empty, not "localhost"!)
}
```

Why? The function receives a **copy** of the struct.

```
main():                       setDefaults():
┌──────────────┐              ┌──────────────┐
│   config     │   COPY →     │   cfg        │
│ ┌──────────┐ │              │ ┌──────────┐ │
│ │ Host: "" │ │              │ │ Host: "" │ │ → Modified to "localhost"
│ │ Port: 0  │ │              │ │ Port: 0  │ │ → Modified to 8080
│ └──────────┘ │              │ └──────────┘ │
└──────────────┘              └──────────────┘
   (Original)                    (Copy, then discarded)
```

#### Solution A: Return the Modified Copy

```go
func setDefaults(cfg Config) Config {
    cfg.Host = "localhost"
    cfg.Port = 8080
    return cfg
}

func main() {
    config := Config{}
    config = setDefaults(config)  // Replace with returned copy
    fmt.Println(config.Host)  // "localhost" ✅
}
```

#### Solution B: Use a Pointer

```go
func setDefaults(cfg *Config) {
    cfg.Host = "localhost"  // Modifies through pointer
    cfg.Port = 8080
}

func main() {
    config := Config{}
    setDefaults(&config)  // Pass address
    fmt.Println(config.Host)  // "localhost" ✅
}
```

```
main():                       setDefaults():
┌──────────────┐              ┌──────────────┐
│   config     │              │   cfg        │
│ ┌──────────┐ │              │ ┌──────────┐ │
│ │ Host: "" │ │ ←────────────│ │ 0x100    │ │ (pointer to config)
│ │ Port: 0  │ │   points to  │ └──────────┘ │
│ └──────────┘ │              └──────────────┘
└──────────────┘
   (Modified directly through pointer!)
```

**When to use which:**
- **Return value**: Small structs (< 5 fields), functional style, immutability preferred
- **Pointer**: Large structs, method receivers, need to express "can be nil"

### Scenario 2: Slice Modifications in Functions

#### The Confusing Part

```go
func appendValue(s []int, val int) {
    s = append(s, val)  // Might work, might not!
}

func main() {
    nums := []int{1, 2, 3}
    appendValue(nums, 4)
    fmt.Println(nums)  // [1, 2, 3] - append didn't persist!
}
```

Why? `append` might allocate a NEW array if capacity is exceeded. The new array pointer stays in the copy of the slice header, never reaches the caller.

```
Before append (capacity available):
main():                       appendValue():
┌─────────────────┐           ┌─────────────────┐
│   nums          │  COPY →   │   s             │
│ ┌─────────────┐ │           │ ┌─────────────┐ │
│ │ ptr: 0x300  │─┼───┐       │ │ ptr: 0x300  │─┼───┐
│ │ len: 3      │ │   │       │ │ len: 3      │ │   │
│ │ cap: 5      │ │   │       │ │ cap: 5      │ │   │
│ └─────────────┘ │   │       │ └─────────────┘ │   │
└─────────────────┘   │       └─────────────────┘   │
                      │                             │
                      └───────────┬─────────────────┘
                                  ↓
                      Underlying array (cap=5):
                      [1][2][3][_][_]

After append (new element fits):
main():                       appendValue():
┌─────────────────┐           ┌─────────────────┐
│   nums          │           │   s             │
│ ┌─────────────┐ │           │ ┌─────────────┐ │
│ │ ptr: 0x300  │─┼───┐       │ │ ptr: 0x300  │─┼───┐
│ │ len: 3      │ │   │       │ │ len: 4      │ │   │ ← Changed!
│ │ cap: 5      │ │   │       │ │ cap: 5      │ │   │
│ └─────────────┘ │   │       │ └─────────────┘ │   │
└─────────────────┘   │       └─────────────────┘   │
   (len unchanged!)   │                             │
                      └───────────┬─────────────────┘
                                  ↓
                      Underlying array:
                      [1][2][3][4][_]
                                 ↑
                            Added here, but len not updated in main!

If capacity is exceeded, new array allocated:
main():                       appendValue():
┌─────────────────┐           ┌─────────────────┐
│   nums          │           │   s             │
│ ┌─────────────┐ │           │ ┌─────────────┐ │
│ │ ptr: 0x300  │─┼───┐       │ │ ptr: 0x400  │─┼───┐ NEW array!
│ │ len: 3      │ │   │       │ │ len: 4      │ │   │
│ │ cap: 5      │ │   │       │ │ cap: 10     │ │   │ NEW capacity!
│ └─────────────┘ │   │       │ └─────────────┘ │   │
└─────────────────┘   │       └─────────────────┘   │
                      │                             │
                      ↓                             ↓
              Old array:                    New array:
              [1][2][3][_][_]              [1][2][3][4][_][_][_][_][_][_]
```

#### Solution A: Return the Slice

```go
func appendValue(s []int, val int) []int {
    return append(s, val)
}

func main() {
    nums := []int{1, 2, 3}
    nums = appendValue(nums, 4)  // Update with returned slice
    fmt.Println(nums)  // [1, 2, 3, 4] ✅
}
```

This is the **idiomatic Go pattern**. Most stdlib functions work this way.

#### Solution B: Pass Pointer to Slice

```go
func appendValue(s *[]int, val int) {
    *s = append(*s, val)  // Dereference, append, assign back
}

func main() {
    nums := []int{1, 2, 3}
    appendValue(&nums, 4)
    fmt.Println(nums)  // [1, 2, 3, 4] ✅
}
```

This works but is **less idiomatic**. Use when you have multiple return values already.

#### When Modification DOES Work Without Return

Modifying **elements** always works:

```go
func doubleElements(s []int) {
    for i := range s {
        s[i] *= 2  // Modifies underlying array
    }
}

func main() {
    nums := []int{1, 2, 3}
    doubleElements(nums)
    fmt.Println(nums)  // [2, 4, 6] ✅ Works!
}
```

Why? You're modifying the shared underlying array, not the slice header.

**Summary:**
- Modifying existing elements: Works without return value
- Changing length (append, reslice): Need to return or use pointer

### Scenario 3: Range Loop Variable Address

#### The Dangerous Bug

```go
type Student struct {
    Name string
}

students := []Student{
    {Name: "Alice"},
    {Name: "Bob"},
    {Name: "Charlie"},
}

var pointers []*Student
for _, student := range students {
    pointers = append(pointers, &student)  // ❌ BUG!
}

for _, p := range pointers {
    fmt.Println(p.Name)  // "Charlie" "Charlie" "Charlie" - All point to same!
}
```

Why? The range loop reuses the **same variable** for each iteration!

```
Range loop memory:
┌────────────────┐
│   student      │ ← Single variable, reused each iteration!
│  ┌──────────┐  │
│  │ Name: "" │  │
│  └──────────┘  │
└────────────────┘
   Address: 0x100 (never changes!)

Iteration 1:
student = students[0]  // student.Name = "Alice"
&student → 0x100       // All pointers end up as 0x100!

Iteration 2:
student = students[1]  // student.Name = "Bob" (OVERWRITES student at 0x100)
&student → 0x100

Iteration 3:
student = students[2]  // student.Name = "Charlie" (OVERWRITES again)
&student → 0x100

Result: All pointers point to 0x100, which now contains "Charlie"
```

#### Solution A: Copy to New Variable

```go
for _, student := range students {
    student := student  // Create NEW variable (shadows loop var)
    pointers = append(pointers, &student)  // ✅ Each iteration gets unique address
}
```

#### Solution B: Use Index

```go
for i := range students {
    pointers = append(pointers, &students[i])  // ✅ Address of slice element
}
```

#### Solution C: Use Slice of Pointers from the Start

```go
students := []*Student{
    {Name: "Alice"},
    {Name: "Bob"},
    {Name: "Charlie"},
}

var pointers []*Student
for _, student := range students {
    pointers = append(pointers, student)  // ✅ Already a pointer
}
```

### Scenario 4: Method Receivers

#### When to Use Pointer Receivers

```go
type Counter struct {
    Count int
}

// ❌ Value receiver - doesn't modify original
func (c Counter) IncrementWrong() {
    c.Count++  // Modifies the copy
}

// ✅ Pointer receiver - modifies original
func (c *Counter) Increment() {
    c.Count++  // Modifies through pointer
}

func main() {
    counter := Counter{Count: 0}

    counter.IncrementWrong()
    fmt.Println(counter.Count)  // 0 (unchanged)

    counter.Increment()  // Go automatically does &counter
    fmt.Println(counter.Count)  // 1 ✅
}
```

**Rules for method receivers:**
1. Need to modify the receiver? → Use pointer receiver `*T`
2. Receiver is large (> 4-5 fields)? → Use pointer receiver `*T`
3. Receiver has pointer fields and you're modifying them? → Use pointer receiver `*T`
4. All other methods use pointer receiver? → Be consistent, use pointer receiver `*T`
5. Otherwise → Value receiver `T` is fine

**Consistency rule:** If ANY method needs a pointer receiver, make ALL methods use pointer receivers. Don't mix.

---

## Decision Tree: When Do I Need a Pointer?

```
┌─────────────────────────────────────────┐
│ Do I need to modify the original value? │
└──────────────┬──────────────────────────┘
               │
       ┌───────┴───────┐
       │               │
      YES             NO
       │               │
       ↓               ↓
   Use pointer    ┌────────────────────┐
   or return      │ Is it a large      │
   modified       │ struct (>40 bytes)?│
   value          └────┬────────────────┘
                       │
                ┌──────┴──────┐
                │             │
               YES           NO
                │             │
                ↓             ↓
           Use pointer   Value is fine
          (avoid copy    (cheap to copy)
           overhead)

Special cases:

┌──────────────────────────────┐
│ Is it a method receiver?     │
└──────┬───────────────────────┘
       │
   See "Method Receiver Rules" above
       │
┌──────┴───────────────────────┐
│ Is it a slice/map/channel?   │
└──────┬───────────────────────┘
       │
       ↓
  Usually don't need pointer
  (they already contain pointers internally)

  Exception: Slice operations that change
  len/cap (append) → return the slice

┌──────────────────────────────┐
│ Map value that you need to   │
│ modify?                      │
└──────┬───────────────────────┘
       │
       ↓
  Option 1: Retrieve-modify-reassign
  Option 2: Use map[K]*V instead
```

---

## Practical Patterns

### Pattern 1: Retrieve-Modify-Reassign (Your Exercise 11)

**Use case:** Modifying struct values in a map

```go
// From your Exercise 11: data_modeling.go
func CheckoutBook(lib *Library, memberID, bookID int) bool {
    // RETRIEVE
    member, memberExists := lib.Members[memberID]
    book, bookExists := lib.Books[bookID]

    if !memberExists || !bookExists || !book.Available {
        return false
    }

    // MODIFY
    book.Available = false
    member.CheckedOutBooks = append(member.CheckedOutBooks, bookID)

    // REASSIGN
    lib.Members[memberID] = member
    lib.Books[bookID] = book

    return true
}
```

**Key points:**
- Get copies of values from map
- Modify the copies
- Write copies back to map
- This is **thread-safe** for the individual operations (but you'd need a mutex for concurrent access)

### Pattern 2: Constructor Returns Pointer

```go
// ✅ Idiomatic - return pointer to struct
func NewLibrary() *Library {
    return &Library{
        Books:   make(map[int]Book),
        Members: make(map[int]Member),
    }
}

// Usage:
lib := NewLibrary()  // lib is *Library
lib.Books[1] = Book{...}  // Can modify through pointer
```

**Why return pointer?**
- Allows methods to modify the library
- Efficient (no copying large structs)
- Convention: "New" prefix returns pointer

### Pattern 3: Slice Append Always Returns

```go
// ✅ Correct pattern
func addBooks(books []Book, newBooks ...Book) []Book {
    return append(books, newBooks...)
}

// Usage:
books := []Book{}
books = addBooks(books, book1, book2)  // Reassign!
```

**Why?**
- `append` might reallocate array (new pointer)
- Returned slice has updated len/cap/ptr
- Caller must capture return value

### Pattern 4: Defensive Nil Checks for Nested Maps

```go
// From Exercise 10: SafeNestedIncrement
func SafeNestedIncrement(m map[string]map[string]int, category, item string) int {
    // CHECK if inner map exists
    if m[category] == nil {
        m[category] = make(map[string]int)  // CREATE if missing
    }

    // Now safe to use
    m[category][item]++
    return m[category][item]
}
```

**Pattern:**
1. Check if inner map is nil
2. Initialize if needed
3. Proceed with operation

**When to use:**
- Maps of maps
- Maps of slices
- Any nested reference type

### Pattern 5: Zero Values Are Your Friend

```go
// This works without initialization:
m := make(map[string]int)
m["count"]++  // Zero value for int is 0, increments to 1

// This does NOT work:
m := make(map[string]map[string]int)
m["outer"]["inner"]++  // PANIC! m["outer"] is nil

// Must initialize inner map:
if m["outer"] == nil {
    m["outer"] = make(map[string]int)
}
m["outer"]["inner"]++  // Now works
```

**Zero values:**
- `int`: 0
- `bool`: false
- `string`: ""
- `pointer`: nil
- `slice`: nil (but usable! `append` works on nil slices)
- `map`: nil (NOT usable - must `make` first)

### Pattern 6: Slice of Pointers for Shared Data

```go
// Use when you want multiple slices to "see" updates to same objects
type Library struct {
    AllBooks      []*Book  // Master list
    AvailableBooks []*Book  // Subset of AllBooks
}

func (lib *Library) CheckoutBook(bookID int) {
    // Find book in AvailableBooks
    for i, book := range lib.AvailableBooks {
        if book.ID == bookID {
            book.Available = false  // ✅ Updates the shared *Book

            // Remove from available list
            lib.AvailableBooks = append(
                lib.AvailableBooks[:i],
                lib.AvailableBooks[i+1:]...,
            )
            return
        }
    }
}
```

**Key insight:** Both slices contain pointers to the same `Book` structs. Modifying through one pointer affects all.

---

## Common Mistakes

### Mistake 1: Forgetting to Reassign Map Values

```go
// ❌ WRONG
func updateStudent(students map[int]Student, id int, newGrade string) {
    students[id].Grade = newGrade  // COMPILE ERROR: cannot assign to struct field
}

// ✅ CORRECT
func updateStudent(students map[int]Student, id int, newGrade string) {
    student := students[id]
    student.Grade = newGrade
    students[id] = student
}
```

### Mistake 2: Not Capturing `append` Return Value

```go
// ❌ WRONG
func addItem(items []int, item int) {
    append(items, item)  // Return value ignored!
}

// ✅ CORRECT
func addItem(items []int, item int) []int {
    return append(items, item)
}

// Usage:
items = addItem(items, 42)  // Reassign!
```

### Mistake 3: Range Loop Variable Pointer

```go
// ❌ WRONG
var pointers []*Student
for _, student := range students {
    pointers = append(pointers, &student)  // All point to same loop var!
}

// ✅ CORRECT - Option 1: Shadow variable
for _, student := range students {
    student := student  // New variable each iteration
    pointers = append(pointers, &student)
}

// ✅ CORRECT - Option 2: Use index
for i := range students {
    pointers = append(pointers, &students[i])
}
```

### Mistake 4: Assuming Slice Assignment Creates Deep Copy

```go
// ❌ WRONG ASSUMPTION
original := []int{1, 2, 3}
copy := original          // Shallow copy - shares array!
copy[0] = 99
fmt.Println(original[0])  // 99 (not 1!)

// ✅ CORRECT - Deep copy
original := []int{1, 2, 3}
copy := make([]int, len(original))
copy(copy, original)  // Built-in copy function
copy[0] = 99
fmt.Println(original[0])  // 1 ✅
```

### Mistake 5: Nil Map Usage

```go
// ❌ WRONG
var m map[string]int  // nil map
m["key"] = 42         // PANIC: assignment to entry in nil map

// ✅ CORRECT
m := make(map[string]int)
m["key"] = 42

// OR
var m map[string]int
if m == nil {
    m = make(map[string]int)
}
m["key"] = 42
```

**Note:** Reading from nil map returns zero value (no panic), but writing panics!

### Mistake 6: Modifying Value Receiver in Method

```go
type Counter struct {
    Count int
}

// ❌ WRONG - Won't modify original
func (c Counter) Increment() {
    c.Count++  // Modifies copy
}

// ✅ CORRECT
func (c *Counter) Increment() {
    c.Count++  // Modifies original through pointer
}
```

### Mistake 7: Nested Map Without Initialization

```go
// ❌ WRONG
m := make(map[string]map[string]int)
m["outer"]["inner"] = 42  // PANIC: nil inner map

// ✅ CORRECT
m := make(map[string]map[string]int)
m["outer"] = make(map[string]int)  // Initialize inner map
m["outer"]["inner"] = 42
```

---

## Quick Reference

### Memory Model Summary

| Category       | Type          | Pass Behavior           | Modify Behavior                    |
|----------------|---------------|-------------------------|------------------------------------|
| **Value**      | `int`         | Copy value              | Copy modified, original unchanged  |
| **Value**      | `bool`        | Copy value              | Copy modified, original unchanged  |
| **Value**      | `string`      | Copy header (ptr+len)   | Immutable (can't modify)           |
| **Value**      | `struct`      | Copy all fields         | Copy modified, original unchanged  |
| **Value**      | `array`       | Copy all elements       | Copy modified, original unchanged  |
| **Reference**  | `slice`       | Copy header (ptr+len+cap)| Modify elements: affects original<br>Append: affects original only if cap sufficient |
| **Reference**  | `map`         | Copy pointer            | Modifications affect original      |
| **Reference**  | `chan`        | Copy pointer            | Operations affect original         |
| **Reference**  | `*T`          | Copy pointer            | Dereference to modify original     |
| **Reference**  | `interface`   | Copy descriptor+pointer | Depends on underlying type         |

### When to Use Pointers

| Scenario                                    | Use Pointer? | Reason                                      |
|---------------------------------------------|--------------|---------------------------------------------|
| Need to modify parameter                    | ✅ Yes        | Value types pass by copy                    |
| Struct > 40 bytes                           | ✅ Yes        | Avoid expensive copy                        |
| Method needs to modify receiver             | ✅ Yes        | Mutation requires pointer                   |
| Function returns new value (functional)     | ❌ No         | Return modified copy instead                |
| Slice/map parameter (not modifying len/cap) | ❌ No         | Already contains pointer                    |
| Need to express "optional" (can be nil)     | ✅ Yes        | Zero value vs nil distinction               |
| Small struct (< 3-4 fields), immutable      | ❌ No         | Copying is cheap                            |

### Map Operations Cheat Sheet

| Operation                        | Code                                                    | Works? | Notes                                |
|----------------------------------|---------------------------------------------------------|--------|--------------------------------------|
| Read from map                    | `v := m[key]`                                           | ✅      | Returns zero value if missing        |
| Check existence                  | `v, ok := m[key]`                                       | ✅      | ok=false if missing                  |
| Write to map                     | `m[key] = value`                                        | ✅      |                                      |
| Delete from map                  | `delete(m, key)`                                        | ✅      | Safe even if key doesn't exist       |
| Modify map value field           | `m[key].field = x`                                      | ❌      | Cannot assign to struct field        |
| Retrieve-modify-reassign         | `v := m[key]; v.field = x; m[key] = v`                  | ✅      | Correct pattern for `map[K]V`        |
| Modify pointer map value         | `m[key].field = x` (if `m` is `map[K]*V`)               | ✅      | Works with map of pointers           |
| Nested map without init          | `m[outer][inner] = x`                                   | ❌      | Panics if `m[outer]` is nil          |
| Nested map with check            | `if m[outer] == nil { m[outer] = make(...) }; m[outer][inner] = x` | ✅ | Defensive initialization             |
| Iterate and modify values        | `for k := range m { v := m[k]; v.field = x; m[k] = v }` | ✅      | Use retrieve-modify-reassign         |

### Slice Operations Cheat Sheet

| Operation                        | Code                                     | Modifies Original? | Notes                                    |
|----------------------------------|------------------------------------------|--------------------|------------------------------------------|
| Modify element                   | `s[i] = x`                               | ✅ Yes              | Shared underlying array                  |
| Append (capacity available)      | `s = append(s, x)`                       | Maybe              | Affects original if cap sufficient       |
| Append (capacity exceeded)       | `s = append(s, x)`                       | ❌ No               | New array allocated                      |
| Append in function               | `func f(s []int) { s = append(s, x) }`   | ❌ No               | Must return slice                        |
| Append with return               | `func f(s []int) []int { return append(s, x) }` | ✅ Yes      | Correct pattern                          |
| Pass to function, modify element | `func f(s []int) { s[0] = x }`           | ✅ Yes              | Shared array                             |
| Copy slice header                | `s2 := s1`                               | Shared array       | Both point to same array                 |
| Deep copy slice                  | `s2 := make([]T, len(s1)); copy(s2, s1)` | Independent        | Separate arrays                          |
| Reslice                          | `s2 := s1[1:3]`                          | Shared array       | Subslice shares array                    |

### Function Parameter Patterns

| Need                              | Pattern                                    | Example                                                     |
|-----------------------------------|--------------------------------------------|-------------------------------------------------------------|
| Read only, small struct           | Pass by value                              | `func Display(s Student) { ... }`                           |
| Modify struct                     | Pass pointer                               | `func Update(s *Student) { s.Grade = "A" }`                 |
| Append to slice (common)          | Return slice                               | `func Add(s []int, x int) []int { return append(s, x) }`    |
| Modify slice elements             | Pass by value                              | `func Double(s []int) { for i := range s { s[i] *= 2 } }`   |
| Modify map                        | Pass by value                              | `func Add(m map[K]V, k K, v V) { m[k] = v }`                |
| Return new value (functional)     | Return value                               | `func Increment(c Counter) Counter { c.Count++; return c }` |

---

## Practice: Spot the Bug

Test your understanding! Each example has a bug or gotcha. Can you spot it?

### Exercise 1

```go
type Product struct {
    Name  string
    Price float64
}

func applyDiscount(products []Product, discount float64) {
    for _, product := range products {
        product.Price *= (1 - discount)
    }
}

func main() {
    products := []Product{
        {"Laptop", 1000.0},
        {"Mouse", 25.0},
    }
    applyDiscount(products, 0.10)
    fmt.Println(products[0].Price)  // What prints?
}
```

<details>
<summary>Answer</summary>

**Bug:** Loop variable `product` is a **copy** of each struct. Modifications don't affect the original slice.

**Fix:**
```go
func applyDiscount(products []Product, discount float64) {
    for i := range products {
        products[i].Price *= (1 - discount)  // Modify slice element directly
    }
}
```

**Output:**
- Buggy version: 1000.0 (unchanged)
- Fixed version: 900.0
</details>

### Exercise 2

```go
func collectPointers() []*int {
    var result []*int
    for i := 0; i < 3; i++ {
        result = append(result, &i)
    }
    return result
}

func main() {
    pointers := collectPointers()
    for _, p := range pointers {
        fmt.Println(*p)  // What prints?
    }
}
```

<details>
<summary>Answer</summary>

**Bug:** Taking address of loop variable `i`. All pointers point to the **same variable**, which ends with value 3.

**Output:** 3, 3, 3 (not 0, 1, 2)

**Fix:**
```go
func collectPointers() []*int {
    var result []*int
    for i := 0; i < 3; i++ {
        i := i  // Create new variable each iteration
        result = append(result, &i)
    }
    return result
}
```
</details>

### Exercise 3

```go
type Config struct {
    Settings map[string]string
}

func NewConfig() Config {
    return Config{
        Settings: nil,  // Intentionally nil
    }
}

func main() {
    config := NewConfig()
    config.Settings["key"] = "value"  // What happens?
}
```

<details>
<summary>Answer</summary>

**Bug:** Writing to nil map causes **panic**.

**Fix:**
```go
func NewConfig() Config {
    return Config{
        Settings: make(map[string]string),  // Initialize map
    }
}
```

**Note:** Reading from nil map is safe (returns zero value), but writing panics.
</details>

### Exercise 4

```go
func modifySlice(s []int) {
    s = append(s, 99)
}

func main() {
    nums := []int{1, 2, 3}
    modifySlice(nums)
    fmt.Println(len(nums))  // What prints?
}
```

<details>
<summary>Answer</summary>

**Bug:** `append` updates the slice header (len/cap/ptr), but the modified header stays in the function. Caller's slice unchanged.

**Output:** 3 (not 4)

**Fix:**
```go
func modifySlice(s []int) []int {
    return append(s, 99)
}

func main() {
    nums := []int{1, 2, 3}
    nums = modifySlice(nums)  // Capture return value
    fmt.Println(len(nums))    // 4
}
```
</details>

### Exercise 5

```go
type Cache struct {
    Data map[string]int
}

func (c Cache) Set(key string, value int) {
    c.Data[key] = value
}

func main() {
    cache := Cache{Data: make(map[string]int)}
    cache.Set("answer", 42)
    fmt.Println(cache.Data["answer"])  // What prints?
}
```

<details>
<summary>Answer</summary>

**Not a bug!** This actually works.

**Output:** 42

**Why?** Even though `Set` has a value receiver (receives a copy of `Cache`), the `Data` field is a map, which **contains a pointer**. Both the original and the copy point to the same map in memory.

**However:** This is still bad practice! If `Cache` had other fields, they wouldn't be shared. Be consistent:
- Methods that modify: Use pointer receiver `*Cache`
- Methods that read only: Can use value receiver `Cache` (if small)

**Better:**
```go
func (c *Cache) Set(key string, value int) {
    c.Data[key] = value
}
```
</details>

---

## Summary: Key Takeaways

### The Three Rules of Go Memory

1. **Everything is pass-by-value** - Always. No exceptions. Some types just contain pointers.
2. **Map values are not addressable** - Use retrieve-modify-reassign pattern.
3. **Slices share arrays but have independent headers** - `append` changes headers, so return the slice.

### When You Hit This Issue

**If you can't modify a struct in a map:**
- Pattern A: Retrieve → Modify → Reassign (what you did in Exercise 11) ✅
- Pattern B: Use `map[K]*V` instead of `map[K]V`

**If changes to a parameter don't persist:**
- Is it a value type? → Pass pointer or return modified value
- Is it a slice and you used `append`? → Return the slice

**If all pointers from a loop point to the same thing:**
- Don't take address of loop variable → Create new variable each iteration

### Your Mental Model Checklist

When you write Go code, ask:

- [ ] Am I passing a value type? (struct, array) → **Expensive copy, consider pointer**
- [ ] Am I modifying a parameter? → **Use pointer or return value**
- [ ] Am I using `append` on a slice? → **Must return and reassign**
- [ ] Am I modifying a struct in a map? → **Retrieve-modify-reassign pattern**
- [ ] Am I taking address in a loop? → **Shadow loop variable first**
- [ ] Is this a method that modifies? → **Use pointer receiver**

### Resources for Deeper Learning

- [Go FAQ: Why do T and *T have different method sets?](https://go.dev/doc/faq#methods_on_values_or_pointers)
- [Effective Go: Pointers vs Values](https://go.dev/doc/effective_go#pointers_vs_values)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Go maps in action](https://go.dev/blog/maps)

---

**Now go forth and modify with confidence!** The map value bug will never trip you up again.
