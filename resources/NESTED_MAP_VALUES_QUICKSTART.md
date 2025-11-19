# Nested Map Values: Quick Reference

**For:** Exercise 11.5 (Music Playlist) and any nested map-of-structs patterns

**Problem:** Maps of structs where structs contain maps of structs = nested value semantics

---

## The Golden Rule

**Always get nested values from the COPY, never from the original map!**

```go
// ❌ WRONG
user := lib.Users[userID]
playlist := lib.Users[userID].Playlists[playlistID]  // Gets from ORIGINAL!

// ✅ CORRECT
user := lib.Users[userID]
playlist := user.Playlists[playlistID]  // Gets from COPY!
```

---

## Visual Decision Tree

```
Do you need to modify a struct field?
├─ No → Just read it, no worries
└─ Yes ↓

Is the struct in a map?
├─ No → Modify directly
└─ Yes ↓

Is it a NESTED map (map inside map)?
├─ No (single level) → Use 1-level pattern (see below)
└─ Yes ↓

                  ┌─────────────────────────────────────┐
                  │  NESTED MAP MODIFICATION PATTERN    │
                  └─────────────────────────────────────┘

                  1. Get outer struct from outer map
                           ↓
                  2. Get inner struct from OUTER's map
                           ↓
                  3. Modify inner struct
                           ↓
                  4. Reassign inner to outer's map
                           ↓
                  5. Reassign outer to outer map
```

---

## Side-by-Side Comparison

| Single-Level Map | Nested Map (Two Levels) |
|------------------|-------------------------|
| `map[K]V` | `map[K]V` where `V` has `map[K2]V2` field |
| **1 retrieve** | **2 retrieves** |
| **1 modify** | **1 modify** |
| **1 reassign** | **2 reassigns** |

### Single-Level (Exercise 11)

```go
type Library struct {
    Books map[int]Book  // Book is VALUE
}

// Modify book
book := lib.Books[bookID]      // 1. Get
book.Available = false         // 2. Modify
lib.Books[bookID] = book       // 3. Reassign
```

### Nested (Exercise 11.5)

```go
type Library struct {
    Users map[string]User  // User is VALUE
}
type User struct {
    Playlists map[string]Playlist  // Playlist is VALUE (nested!)
}

// Modify playlist
user := lib.Users[userID]                    // 1. Get outer
playlist := user.Playlists[playlistID]       // 2. Get inner FROM OUTER
playlist.Songs = append(playlist.Songs, song) // 3. Modify inner
user.Playlists[playlistID] = playlist        // 4. Reassign inner to outer
lib.Users[userID] = user                     // 5. Reassign outer to map
```

---

## Common Mistakes & Fixes

### Mistake 1: Getting Inner from Original

```go
// ❌ WRONG - Gets playlist from a DIFFERENT copy than user
user := lib.Users[userID]
playlist := lib.Users[userID].Playlists[playlistID]
//          ↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑ This creates a NEW copy of User!

// ✅ CORRECT - Gets playlist from the SAME copy
user := lib.Users[userID]
playlist := user.Playlists[playlistID]
//          ↑↑↑↑ Uses the copy we already have
```

**Why this matters:** Each map access returns a new copy. Mixing copies from different accesses = mismatch!

### Mistake 2: Forgetting Final Reassignment

```go
// ❌ WRONG - Modifies user copy but never puts it back
user := lib.Users[userID]
playlist := user.Playlists[playlistID]
playlist.Songs = append(playlist.Songs, song)
user.Playlists[playlistID] = playlist
// Missing: lib.Users[userID] = user  ← MUST have this!

// ✅ CORRECT
user := lib.Users[userID]
playlist := user.Playlists[playlistID]
playlist.Songs = append(playlist.Songs, song)
user.Playlists[playlistID] = playlist
lib.Users[userID] = user  // ✅ Put outer back
```

### Mistake 3: Wrong Order of Operations

```go
// ❌ WRONG - Tries to put back before modifying
user := lib.Users[userID]
playlist := user.Playlists[playlistID]
lib.Users[userID] = user  // Too early!
playlist.Songs = append(playlist.Songs, song)
user.Playlists[playlistID] = playlist  // Modifying AFTER putting back!

// ✅ CORRECT - Modify fully, THEN put back
user := lib.Users[userID]
playlist := user.Playlists[playlistID]
playlist.Songs = append(playlist.Songs, song)
user.Playlists[playlistID] = playlist
lib.Users[userID] = user  // Last step!
```

---

## Pattern Templates

### Template 1: Add to Nested Collection

```go
// Add item to nested slice/map
func AddItemToNested(lib *Library, outerKey, innerKey string, item Item) {
    outer := lib.Outers[outerKey]               // Get outer
    inner := outer.Inners[innerKey]             // Get inner from outer
    inner.Items = append(inner.Items, item)     // Modify inner
    outer.Inners[innerKey] = inner              // Put inner back
    lib.Outers[outerKey] = outer                // Put outer back
}
```

### Template 2: Remove from Nested Collection

```go
// Remove item from nested slice
func RemoveItemFromNested(lib *Library, outerKey, innerKey string, itemIndex int) {
    outer := lib.Outers[outerKey]
    inner := outer.Inners[innerKey]
    inner.Items = append(inner.Items[:itemIndex], inner.Items[itemIndex+1:]...)
    outer.Inners[innerKey] = inner
    lib.Outers[outerKey] = outer
}
```

### Template 3: Modify Nested Field

```go
// Change a field in nested struct
func ModifyNestedField(lib *Library, outerKey, innerKey string, newValue int) {
    outer := lib.Outers[outerKey]
    inner := outer.Inners[innerKey]
    inner.Value = newValue
    outer.Inners[innerKey] = inner
    lib.Outers[outerKey] = outer
}
```

---

## Quick Diagnosis: "Why Isn't My Modification Working?"

Ask yourself:

1. **Did I get the inner value from the outer COPY?**
   ```go
   // ✅ YES: inner := outer.Inners[key]
   // ❌ NO:  inner := lib.Outers[outerKey].Inners[key]
   ```

2. **Did I reassign the inner value to the outer copy's map?**
   ```go
   // ✅ YES: outer.Inners[key] = modifiedInner
   // ❌ NO:  (forgot this line!)
   ```

3. **Did I reassign the outer copy back to the library's map?**
   ```go
   // ✅ YES: lib.Outers[outerKey] = outer
   // ❌ NO:  (forgot this line!)
   ```

**If all 3 are YES, your code should work!**

---

## When to Use Pointers Instead

If you're doing this pattern CONSTANTLY, consider:

```go
// Instead of:
type Library struct {
    Users map[string]User  // Value semantics
}

// Use:
type Library struct {
    Users map[string]*User  // Pointer semantics
}

// Now you can modify directly:
lib.Users[userID].Playlists[playlistID].Songs = append(...)  // ✅ Works!
```

**Trade-offs:**

| `map[K]V` | `map[K]*V` |
|-----------|------------|
| Safer (no nil) | Must check for nil |
| Explicit copies | Aliasing risks |
| Less GC pressure | More allocations |
| Retrieve-modify-reassign | Direct modification |

**Rule of thumb:**
- Learning phase: Use `map[K]V` to understand value semantics
- Production code: Use `map[K]*V` for frequently-modified nested structures

---

## Nesting Depth Guide

| Depth | Pattern | Example |
|-------|---------|---------|
| 1 | Single map | `map[int]Book` → 1 reassignment |
| 2 | Map with map field | `map[string]User` where `User` has `map[string]Playlist` → 2 reassignments |
| 3+ | Triple nesting | **Consider pointers!** Gets messy. |

**Formula:** `Reassignments = Depth`

---

## Memory Aid

**"From Copy, To Copy, To Map"**

1. **From Copy:** Get inner values from the outer COPY (not from original map)
2. **To Copy:** Put modified inner back into the outer COPY
3. **To Map:** Put modified outer COPY back into the map

---

## Exercises for Practice

See `/02-data-structures/11.6_map_value_drills/` for:
- Drill 4: Exact Exercise 11.5 pattern in isolation
- Drill 5: Mixed operations with nested structures

**Test your understanding:**
- Can you explain why getting from the original map breaks things?
- Can you draw a memory diagram showing two different copies?
- Can you spot this bug in code review?

---

## Related Resources

- **VALUES_AND_REFERENCES.md** - Section 4: "Nested Map Values" (detailed explanation)
- **MAKE_AND_NESTED_STRUCTURES.md** - Section 4.1: "Maps of Maps" (initialization patterns)
- **Exercise 11** (`data_modeling.go`) - Single-level pattern practice
- **Exercise 11.5** (`music_playlist.go`) - Your real challenge
- **Exercise 11.6** (`map_value_drills/`) - Focused practice

---

**Print this page** and keep it next to you while coding Exercise 11.5! 📄
