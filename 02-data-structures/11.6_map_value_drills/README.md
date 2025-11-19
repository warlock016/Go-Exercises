# Exercise 11.6: Map Value Semantics Drills

## 🎯 Learning Goal

Master the retrieve-modify-reassign pattern for map values, from single-level to nested structures. This exercise isolates the exact pattern you need for Exercise 11.5 (Music Playlist).

## 📝 Problem Description

Maps in Go store **values**, not references. When you retrieve a struct from a map, you get a **copy**. Modifying that copy doesn't affect the original map unless you reassign it.

**The Challenge:** This becomes tricky with nested structures where you need to reassign at MULTIPLE levels.

## Why This Exercise Exists

After Exercise 11 (Library System), you learned the basic pattern:
```go
book := lib.Books[bookID]     // Get copy
book.Available = false        // Modify copy
lib.Books[bookID] = book      // Reassign ✅
```

But Exercise 11.5 (Music Playlist) has **nested** maps:
```go
type Library struct {
    Users map[string]User  // User is a VALUE
}
type User struct {
    Playlists map[string]Playlist  // Playlist is a VALUE (nested!)
}
```

**This requires TWO reassignments!** These drills build that skill progressively.

---

## 📋 The 5 Drills

### Drill 1: Single-Level Warmup (5 minutes)
**Goal:** Reinforce the basic pattern you already know

Modify a book's price in a simple map. You've done this before!

**Pattern:**
```
Get → Modify → Reassign (1 level)
```

---

### Drill 2: Map of Maps (10 minutes)
**Goal:** Apply pattern to simple nested map (no structs yet)

Work with `map[string]map[string]int` - nested maps without structs.

**Pattern:**
```
Get outer map value → Modify it → Reassign (still 1 level, but nested!)
```

**Insight:** The value IS a map, so modifying it works... but you still need to understand the copy behavior.

---

### Drill 3: Struct with Map Field (15 minutes)
**Goal:** THIS is your exact pattern from 11.5!

```go
type Profile struct {
    Settings map[string]int  // ← Nested map inside struct
}
users := map[string]Profile  // ← Profile is a VALUE
```

**Pattern:**
```
Get Profile copy → Modify its Settings map → Reassign Profile (1 level)
```

**Critical Insight:** When you get `Profile` from the map, its `Settings` field is a map (reference type), BUT the `Profile` itself is a copy. You must reassign the Profile!

---

### Drill 4: Double Nesting - Your Music System (20 minutes)
**Goal:** Master the two-level retrieve-modify-reassign

```go
type Library struct {
    Users map[string]User  // User is VALUE
}
type User struct {
    Playlists map[string]Playlist  // Playlist is VALUE (nested!)
}
```

**Pattern:**
```
Get User copy → Get Playlist copy from User's map → Modify Playlist
→ Reassign Playlist to User's map → Reassign User to Library's map (2 levels!)
```

**This is Exercise 11.5 in isolation!**

---

### Drill 5: Mixed Operations (20 minutes)
**Goal:** Prove mastery with create, modify, delete operations

All the patterns combined with realistic operations.

---

## 🔧 Function Signatures

All functions are in `drills.go`. Tests are in `drills_test.go`.

### Drill 1 Functions
```go
func ModifyBookPrice(books map[int]Book, bookID int, newPrice float64)
```

### Drill 2 Functions
```go
func IncrementNestedValue(data map[string]map[string]int, outerKey, innerKey string)
```

### Drill 3 Functions
```go
func UpdateUserSetting(users map[string]Profile, userID string, setting string, value int)
```

### Drill 4 Functions
```go
func AddSongToPlaylist(lib *Library, userID, playlistID string, song Song)
func GetPlaylistSongCount(lib *Library, userID, playlistID string) int
```

### Drill 5 Functions
```go
func CreateUserWithPlaylist(lib *Library, userID, playlistID string)
func ModifySongInPlaylist(lib *Library, userID, playlistID string, songIndex int, newTitle string)
func RemovePlaylist(lib *Library, userID, playlistID string)
```

---

## 💡 Examples

### Drill 1: Single-Level (You know this!)
```go
books := map[int]Book{1: {Title: "1984", Price: 10.0}}
ModifyBookPrice(books, 1, 15.0)
// books[1].Price is now 15.0 ✅
```

### Drill 3: Struct with Map Field (Tricky!)
```go
users := map[string]Profile{
    "alice": {Name: "Alice", Settings: map[string]int{"theme": 1}},
}
UpdateUserSetting(users, "alice", "theme", 2)
// users["alice"].Settings["theme"] is now 2 ✅

// Common mistake:
users["alice"].Settings["theme"] = 2  // ❌ Won't compile!
// "cannot assign to struct field users["alice"].Settings in map"
```

### Drill 4: Double Nesting (Your exact problem!)
```go
lib := &Library{Users: map[string]User{
    "alice": {Playlists: map[string]Playlist{
        "favorites": {Songs: []Song{{Title: "Song1"}}},
    }},
}}

AddSongToPlaylist(lib, "alice", "favorites", Song{Title: "Song2"})
// lib.Users["alice"].Playlists["favorites"].Songs now has 2 songs ✅

// Common mistake:
user := lib.Users["alice"]
playlist := lib.Users["alice"].Playlists["favorites"]  // ❌ Gets from ORIGINAL, not user copy!
playlist.Songs = append(playlist.Songs, newSong)
user.Playlists["favorites"] = playlist
lib.Users["alice"] = user
// The playlist you modified is NOT the one inside user!

// Correct:
user := lib.Users["alice"]                    // ← Get outer copy
playlist := user.Playlists["favorites"]       // ← Get from COPY (not lib.Users!)
playlist.Songs = append(playlist.Songs, newSong)
user.Playlists["favorites"] = playlist        // ← Put back into copy
lib.Users["alice"] = user                     // ← Put copy back into map ✅
```

---

## 🧪 Testing

Run all drills:
```bash
cd 02-data-structures/11.6_map_value_drills
go test -v
```

Run specific drill:
```bash
go test -v -run TestDrill1
go test -v -run TestDrill4
```

**Expected:** All tests should pass when you implement the functions correctly.

---

## 🤔 Think About

1. **Why does `map[string]map[string]int` (Drill 2) work differently than `map[string]Profile` (Drill 3)?**
   - In Drill 2, the value IS a map (reference type), so modifying it affects all references
   - In Drill 3, the value is a struct (value type), so you get a copy

2. **In Drill 4, why must you get the Playlist from `user.Playlists` and NOT `lib.Users[userID].Playlists`?**
   - Because `lib.Users[userID]` returns a NEW copy each time
   - The `user` variable holds a copy
   - Getting playlist from original map creates a DIFFERENT copy
   - Modifying that different copy and putting it in `user` means they're mismatched!

3. **When would you use `map[K]*V` (pointers) instead of `map[K]V` (values)?**
   - If you need to modify frequently (avoid retrieve-modify-reassign pattern)
   - If structs are large (> 100 bytes)
   - If you need shared state between references
   - Trade-off: Must handle nil, more GC pressure, concurrency risks

4. **What's the "nesting depth" rule?**
   - Map nesting depth = Number of reassignments needed
   - Depth 1: `map[K]V` → 1 reassignment
   - Depth 2: `map[K]V` where `V` has `map[K2]V2` → 2 reassignments
   - Depth 3: Three levels → 3 reassignments (gets messy, consider pointers!)

---

## 📊 Success Criteria

You've mastered map value semantics when:

- ✅ All 5 drills pass without looking at hints
- ✅ You can draw memory diagrams showing why double nesting needs two reassignments
- ✅ You can explain to someone else why `user.Playlists[...] ≠ lib.Users[userID].Playlists[...]`
- ✅ You return to Exercise 11.5 and fix GetPlaylistSongs, RemoveSongFromPlaylist without help
- ✅ You recognize the pattern instantly in future code

---

## 🎓 What This Teaches

- **Map value semantics** - Maps store values, return copies
- **Nested structure modification** - Multi-level retrieve-modify-reassign pattern
- **Mental model** - Visualizing copies vs references
- **Common pitfalls** - Getting from original vs copy, forgetting reassignments
- **Decision making** - When to use values vs pointers
- **Real-world patterns** - How databases, caches, and state management work

---

## 🔗 Related Resources

- `/resources/VALUES_AND_REFERENCES.md` - Section 3: "The Map Value Problem"
- `/resources/NESTED_MAP_VALUES_QUICKSTART.md` - 1-page cheat sheet (coming soon)
- Exercise 11 (`data_modeling.go`) - Single-level pattern
- Exercise 11.5 (`music_playlist.go`) - Your real challenge

---

**Ready?** Open `drills.go`, implement the functions, run the tests, and watch the pattern click! 🧠
