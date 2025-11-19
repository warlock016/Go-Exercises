# Music Playlist System

## Learning Goal

Practice architectural decision-making by designing a complete system from scratch. You'll choose data structures based on requirements and performance targets, learning how design choices create trade-offs between time complexity, memory usage, and code complexity.

## Problem Description

Build a music playlist management system similar to Spotify or Apple Music. You'll implement a library of songs, user-owned playlists, and operations like adding/removing songs, searching, and playlist management.

**Your Task:** Design the entire system architecture from scratch. You decide:
- What structs to create and what fields they contain
- What data structures to use (slices, maps, combinations)
- What function signatures look like
- How to organize relationships between entities
- How to achieve the performance targets

Multiple valid solutions exist. The tests verify behavior, not implementation.

## System Requirements

### Entities

Your system must model three entities with these properties:

**Song**
- Unique identifier (never reused, even after deletion)
- Title
- Artist name
- Duration in seconds

**Playlist**
- Unique identifier (scoped to user - different users can have playlists with same ID)
- Name
- Ordered collection of songs
- Current playback position (0-indexed)

**User**
- Unique identifier
- Username
- Collection of playlists they own

### Core Operations

Implement functionality for the following operations:

#### Library Management
1. **Add song to library** - Store a new song with unique ID
2. **Remove song from library** - Delete song and remove from all playlists containing it
3. **Get song by ID** - Retrieve song details by its identifier
4. **Find songs by artist** - Get all songs by a specific artist (case-sensitive exact match)
5. **Search songs by title** - Find songs with titles containing search term (case-insensitive partial match)

#### Playlist Management
6. **Create playlist** - New empty playlist for a user
7. **Delete playlist** - Remove playlist from user's collection
8. **Add song to playlist** - Append song to end of playlist (maintains order)
9. **Remove song from playlist** - Remove song at specific position (shifts remaining songs)
10. **Get playlist songs** - Return ordered list of songs in playlist
11. **Play next song** - Advance position and return current song (returns nil/error at end)
12. **Calculate playlist duration** - Sum of all song durations in seconds
13. **Shuffle playlist** - Randomize song order, reset position to 0

#### Cross-Entity Operations
14. **Find playlists containing song** - Get all playlists across all users that contain a specific song ID
15. **Get user's playlists** - Return all playlists owned by a user
16. **Get user's most recent playlist** - Return the last created playlist for a user

### Business Rules

Your implementation must enforce these constraints:

- **Song ID Uniqueness**: Song IDs must be unique across the entire library. Once used, an ID cannot be reused even after the song is deleted.
- **Playlist ID Scoping**: Playlist IDs are unique per user. User A and User B can both have a playlist with ID "1", but User A cannot have two playlists with ID "1".
- **Duplicate Songs in Playlists**: A playlist can contain the same song multiple times (like Spotify allows). Example: [SongA, SongB, SongA] is valid.
- **Playlist Size Limit**: Maximum 1000 songs per playlist.
- **Positive Duration**: Song duration must be greater than 0.
- **Empty Playlists**: Playlists can be empty (0 songs).
- **Cascading Deletes**: Removing a song from the library must remove it from all playlists.
- **Playback Position**: Default position is 0. Position must be valid (0 to len(songs)-1). Empty playlist has no valid position.

### Performance Targets

Your design should achieve these time complexities:

- **Find song by ID**: O(1) - constant time lookup
- **Add song to playlist**: O(1) - constant time append
- **Find playlists containing song**: O(p) where p = number of playlists total
- **Get all songs by artist**: O(n) where n = total songs in library
- **Search songs by title**: O(n) where n = total songs in library

## Examples

### Example 1: Basic Song and Playlist Operations

```go
// Add songs to library
song1 := AddSong("1", "Bohemian Rhapsody", "Queen", 354)
song2 := AddSong("2", "Stairway to Heaven", "Led Zeppelin", 482)
song3 := AddSong("3", "Hotel California", "Eagles", 391)

// Create user and playlist
user := CreateUser("u1", "john_doe")
playlist := CreatePlaylist(user, "p1", "Rock Classics")

// Add songs to playlist
AddSongToPlaylist(playlist, song1)
AddSongToPlaylist(playlist, song2)
AddSongToPlaylist(playlist, song3)

// Get playlist songs
songs := GetPlaylistSongs(playlist)
// Returns: [song1, song2, song3]

// Calculate duration
duration := CalculatePlaylistDuration(playlist)
// Returns: 1227 (354 + 482 + 391)
```

### Example 2: Playback and Position

```go
// Create playlist with 3 songs
playlist := CreatePlaylist(user, "p1", "My Mix")
AddSongToPlaylist(playlist, song1)
AddSongToPlaylist(playlist, song2)
AddSongToPlaylist(playlist, song3)

// Initial position is 0
current := PlayNext(playlist)
// Returns: song1, position advances to 1

current = PlayNext(playlist)
// Returns: song2, position advances to 2

current = PlayNext(playlist)
// Returns: song3, position advances to 3

current = PlayNext(playlist)
// Returns: nil/error (end of playlist)
```

### Example 3: Search and Discovery

```go
// Add songs
AddSong("1", "Let It Be", "The Beatles", 243)
AddSong("2", "Let It Go", "Idina Menzel", 225)
AddSong("3", "Yesterday", "The Beatles", 125)
AddSong("4", "Let Me Love You", "DJ Snake", 205)

// Search by title (case-insensitive partial match)
results := SearchSongsByTitle("let")
// Returns: [song1, song2, song4] (all containing "let")

// Find by artist (case-sensitive exact match)
results = FindSongsByArtist("The Beatles")
// Returns: [song1, song3]
```

### Example 4: Duplicate Songs in Playlist

```go
// Same song can appear multiple times
playlist := CreatePlaylist(user, "p1", "Repeat Favorites")
AddSongToPlaylist(playlist, song1)
AddSongToPlaylist(playlist, song2)
AddSongToPlaylist(playlist, song1)  // Same as first
AddSongToPlaylist(playlist, song1)  // Same again

songs := GetPlaylistSongs(playlist)
// Returns: [song1, song2, song1, song1]

duration := CalculatePlaylistDuration(playlist)
// Returns: duration of song1*3 + song2*1
```

### Example 5: Finding Playlists Containing a Song

```go
// Multiple users, multiple playlists
user1 := CreateUser("u1", "alice")
user2 := CreateUser("u2", "bob")

playlist1 := CreatePlaylist(user1, "p1", "Workout")
playlist2 := CreatePlaylist(user1, "p2", "Chill")
playlist3 := CreatePlaylist(user2, "p1", "Road Trip")

song := AddSong("1", "Thunderstruck", "AC/DC", 292)

AddSongToPlaylist(playlist1, song)
AddSongToPlaylist(playlist3, song)

// Find all playlists containing this song
playlists := FindPlaylistsContainingSong(song)
// Returns: [playlist1, playlist3]
```

### Example 6: Cascading Deletes

```go
// Song appears in multiple playlists
song := AddSong("1", "Song X", "Artist Y", 200)
playlist1 := CreatePlaylist(user, "p1", "List 1")
playlist2 := CreatePlaylist(user, "p2", "List 2")

AddSongToPlaylist(playlist1, song)
AddSongToPlaylist(playlist2, song)

// Remove song from library
RemoveSong(song)

// Song is removed from all playlists
songs1 := GetPlaylistSongs(playlist1)
// Returns: [] (empty)

songs2 := GetPlaylistSongs(playlist2)
// Returns: [] (empty)
```

### Example 7: Shuffle Playlist

```go
playlist := CreatePlaylist(user, "p1", "Mix")
AddSongToPlaylist(playlist, song1)
AddSongToPlaylist(playlist, song2)
AddSongToPlaylist(playlist, song3)
AddSongToPlaylist(playlist, song4)

// Original order: [song1, song2, song3, song4]
ShufflePlaylist(playlist)
// New order: [song3, song1, song4, song2] (randomized)
// Position reset to 0
```

### Example 8: User Playlist Management

```go
user := CreateUser("u1", "sarah")

// Create playlists (tracked in order)
p1 := CreatePlaylist(user, "p1", "Morning")
p2 := CreatePlaylist(user, "p2", "Evening")
p3 := CreatePlaylist(user, "p3", "Workout")

// Get all user's playlists
playlists := GetUserPlaylists(user)
// Returns: [p1, p2, p3]

// Get most recently created
recent := GetUserMostRecentPlaylist(user)
// Returns: p3

// Delete a playlist
DeletePlaylist(user, p2)

playlists = GetUserPlaylists(user)
// Returns: [p1, p3]
```

## Testing

Run the tests from the exercise directory:

```bash
cd /Users/mode/Documents/Code/Go\ Exercises/02-data-structures/11.5_music_playlist/
go test -v
```

The test file is provided with ~50 test cases covering:
- All core operations
- Edge cases (empty playlists, boundaries, invalid inputs)
- Business rule enforcement (uniqueness, limits, cascading)
- Performance characteristics (where testable)

Your implementation must make all tests pass.

## Think About

Before and during implementation, consider these questions:

**Data Structure Selection:**
- How will you achieve O(1) song lookup by ID?
- How will you maintain playlist song order efficiently?
- What data structure ensures song ID uniqueness?
- How will you track which playlists contain a specific song?
- Should you use `map[K]V` or `map[K]*V`? What are the trade-offs?

**Design Trade-offs:**
- Fast lookup vs memory usage: Which is more important?
- Should you store song references or song copies in playlists?
- How do you handle cascading deletes efficiently?
- Should playback position be stored in the playlist struct or managed separately?

**Architecture Decisions:**
- Do you need a central "Library" or "MusicSystem" struct?
- Should operations be functions or methods?
- How do you model relationships: User owns Playlists, Playlists contain Songs?
- Should Song/Playlist/User be exported or internal?

**Performance Considerations:**
- Does your `FindPlaylistsContainingSong` scan all playlists (O(p)) or use an index?
- Does `SearchSongsByTitle` require full library scan (acceptable) or can it be optimized?
- What happens to performance as the library grows to 10,000 songs?

**Edge Cases:**
- What happens if you try to add song 1001 to a playlist?
- Can you create a playlist with the same ID for the same user twice?
- What if you remove a song from a playlist position that doesn't exist?
- How do you handle shuffling an empty playlist?

## Success Criteria

Your implementation is complete when:

1. All tests pass (`go test -v`)
2. Business rules are enforced (IDs, limits, cascading deletes)
3. Performance targets are met (O(1) song lookup, O(1) playlist append)
4. Code handles edge cases gracefully (empty collections, invalid indices)
5. Design decisions are intentional (you can explain why you chose specific data structures)

## What This Teaches

This exercise develops several critical skills:

**System Design**: Building a multi-entity system with relationships and constraints.

**Data Structure Selection**: Choosing the right tool for the job:
- Maps for fast lookups
- Slices for ordered collections
- Combinations for complex requirements

**Trade-off Analysis**: Understanding that design choices have costs:
- Time complexity vs space complexity
- Code simplicity vs performance
- Flexibility vs enforcement

**Architectural Thinking**: Designing interfaces and abstractions:
- What should be a function vs a method?
- How to model entity relationships?
- Where to encapsulate complexity?

**Real-world Patterns**: This mirrors actual systems:
- E-commerce (users, carts, products)
- Social media (users, posts, likes)
- Content management (articles, tags, authors)

**Performance Awareness**: Understanding Big-O in practice:
- When O(n) is acceptable (rare operations)
- When O(1) is critical (frequent operations)
- How data structure choice affects complexity

This is not just about making tests pass - it's about learning to design systems that are correct, efficient, and maintainable.
