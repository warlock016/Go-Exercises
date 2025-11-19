package drills

// ============================================================================
// DRILL 1: Single-Level Warmup
// ============================================================================

type Book struct {
	Title string
	Price float64
}

// ModifyBookPrice changes the price of a book in the map
func ModifyBookPrice(books map[int]Book, bookID int, newPrice float64) {
	// TODO(human): Implement retrieve-modify-reassign pattern

	result := books[bookID]
	result.Price = newPrice
	books[bookID] = result
}

// ============================================================================
// DRILL 2: Map of Maps (Nested Maps without Structs)
// ============================================================================

// IncrementNestedValue increments a value in a nested map structure
// If the outer key doesn't exist, create it with an empty inner map
// If the inner key doesn't exist, start at 0
func IncrementNestedValue(data map[string]map[string]int, outerKey, innerKey string) {

	if data == nil {
		data = make(map[string]map[string]int)
		data[outerKey][innerKey] = 0
	}
	if data[outerKey] == nil {
		data[outerKey] = map[string]int{}
		data[outerKey][innerKey] = 0
	}
	data[outerKey][innerKey]++
}

// ============================================================================
// DRILL 3: Struct with Map Field
// ============================================================================

type Profile struct {
	Name     string
	Settings map[string]int
}

// UpdateUserSetting changes a user's setting value
// User must exist in the map
func UpdateUserSetting(users map[string]Profile, userID string, setting string, value int) {
	// TODO(human): Implement retrieve-modify-reassign for struct with map field

	// newProfile := users[userID]
	// newProfile.Settings[setting] = value
	// users[userID] = newProfile

	users[userID].Settings[setting] = value
}

// ============================================================================
// DRILL 4: Double Nesting - Your Music System Pattern
// ============================================================================

type Song struct {
	Title string
}

type Playlist struct {
	Songs []Song
}

type User struct {
	Playlists map[string]Playlist
}

type Library struct {
	Users map[string]User
}

// AddSongToPlaylist adds a song to a user's playlist
// Assumes user and playlist exist
func AddSongToPlaylist(lib *Library, userID, playlistID string, song Song) {
	// TODO(human): Implement TWO-LEVEL retrieve-modify-reassign
	// CRITICAL: Get playlist from USER COPY, not from lib.Users[userID]!

	currentUser := lib.Users[userID]
	currentPlaylist := currentUser.Playlists[playlistID]
	currentPlaylist.Songs = append(currentPlaylist.Songs, song)
	currentUser.Playlists[playlistID] = currentPlaylist
	lib.Users[userID] = currentUser
}

// GetPlaylistSongCount returns the number of songs in a playlist
// Returns 0 if user or playlist doesn't exist
func GetPlaylistSongCount(lib *Library, userID, playlistID string) int {

	currentUser := lib.Users[userID]
	songs := currentUser.Playlists[playlistID].Songs
	// TODO(human): Implement
	return len(songs)
}

// ============================================================================
// DRILL 5: Mixed Operations (Prove Mastery)
// ============================================================================

// CreateUserWithPlaylist creates a new user with an empty playlist
func CreateUserWithPlaylist(lib *Library, userID, playlistID string) {
	// TODO(human): Implement
	newUser := lib.Users[userID]
	// userPlaylist := newUser.Playlists
	newPlaylist := Playlist{
		Songs: make([]Song, 0),
	}

	if newUser.Playlists == nil {
		newUser.Playlists = make(map[string]Playlist)
	}

	if _, exists := newUser.Playlists[playlistID]; !exists {
		newUser.Playlists[playlistID] = newPlaylist
	}
	lib.Users[userID] = newUser
}

// ModifySongInPlaylist changes the title of a song at a specific index in a playlist
// Assumes user, playlist, and song at index exist
func ModifySongInPlaylist(lib *Library, userID, playlistID string, songIndex int, newTitle string) {
	// TODO(human): Implement nested retrieve-modify-reassign

	if _, ok := lib.Users[userID]; !ok {
		return
	}

	currentUser := lib.Users[userID]
	currentPlaylist := currentUser.Playlists[playlistID]
	currentSong := currentPlaylist.Songs[songIndex]
	currentSong.Title = newTitle

	lib.Users[userID].Playlists[playlistID].Songs[songIndex] = currentSong
}

// RemovePlaylist removes a playlist from a user
// Assumes user and playlist exist
func RemovePlaylist(lib *Library, userID, playlistID string) {
	// TODO(human): Implement (requires modifying user's Playlists map)
	currentUser := lib.Users[userID]
	delete(currentUser.Playlists, playlistID)
	lib.Users[userID] = currentUser
}
