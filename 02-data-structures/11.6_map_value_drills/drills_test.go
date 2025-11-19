package drills

import (
	"testing"
)

// ============================================================================
// DRILL 1 TESTS: Single-Level Warmup
// ============================================================================

func TestDrill1_ModifyBookPrice(t *testing.T) {
	tests := []struct {
		name      string
		initial   map[int]Book
		bookID    int
		newPrice  float64
		wantPrice float64
	}{
		{
			name:      "Modify existing book",
			initial:   map[int]Book{1: {Title: "1984", Price: 10.0}},
			bookID:    1,
			newPrice:  15.0,
			wantPrice: 15.0,
		},
		{
			name:      "Modify to zero price",
			initial:   map[int]Book{2: {Title: "Go Programming", Price: 50.0}},
			bookID:    2,
			newPrice:  0.0,
			wantPrice: 0.0,
		},
		{
			name:      "Modify to higher price",
			initial:   map[int]Book{3: {Title: "Clean Code", Price: 25.0}},
			bookID:    3,
			newPrice:  99.99,
			wantPrice: 99.99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ModifyBookPrice(tt.initial, tt.bookID, tt.newPrice)
			got := tt.initial[tt.bookID].Price
			if got != tt.wantPrice {
				t.Errorf("ModifyBookPrice() = %v, want %v", got, tt.wantPrice)
			}
			// Verify title wasn't changed
			if tt.initial[tt.bookID].Title == "" {
				t.Error("Title should not be empty after price modification")
			}
		})
	}
}

// ============================================================================
// DRILL 2 TESTS: Map of Maps
// ============================================================================

func TestDrill2_IncrementNestedValue(t *testing.T) {
	t.Run("Increment existing value", func(t *testing.T) {
		data := map[string]map[string]int{
			"user1": {"score": 10},
		}
		IncrementNestedValue(data, "user1", "score")
		if got := data["user1"]["score"]; got != 11 {
			t.Errorf("IncrementNestedValue() = %v, want 11", got)
		}
	})

	t.Run("Create new inner key", func(t *testing.T) {
		data := map[string]map[string]int{
			"user1": {},
		}
		IncrementNestedValue(data, "user1", "newScore")
		if got := data["user1"]["newScore"]; got != 1 {
			t.Errorf("IncrementNestedValue() = %v, want 1 (started at 0)", got)
		}
	})

	t.Run("Create new outer key with inner map", func(t *testing.T) {
		data := map[string]map[string]int{}
		IncrementNestedValue(data, "newUser", "score")
		if data["newUser"] == nil {
			t.Error("Outer map should be created")
		}
		if got := data["newUser"]["score"]; got != 1 {
			t.Errorf("IncrementNestedValue() = %v, want 1", got)
		}
	})

	t.Run("Multiple increments", func(t *testing.T) {
		data := map[string]map[string]int{
			"user1": {"score": 5},
		}
		IncrementNestedValue(data, "user1", "score")
		IncrementNestedValue(data, "user1", "score")
		IncrementNestedValue(data, "user1", "score")
		if got := data["user1"]["score"]; got != 8 {
			t.Errorf("After 3 increments, got %v, want 8", got)
		}
	})
}

// ============================================================================
// DRILL 3 TESTS: Struct with Map Field
// ============================================================================

func TestDrill3_UpdateUserSetting(t *testing.T) {
	t.Run("Update existing setting", func(t *testing.T) {
		users := map[string]Profile{
			"alice": {Name: "Alice", Settings: map[string]int{"theme": 1}},
		}
		UpdateUserSetting(users, "alice", "theme", 2)
		if got := users["alice"].Settings["theme"]; got != 2 {
			t.Errorf("UpdateUserSetting() = %v, want 2", got)
		}
		// Verify name wasn't changed
		if users["alice"].Name != "Alice" {
			t.Error("User name should not change")
		}
	})

	t.Run("Add new setting to existing user", func(t *testing.T) {
		users := map[string]Profile{
			"bob": {Name: "Bob", Settings: map[string]int{"theme": 1}},
		}
		UpdateUserSetting(users, "bob", "fontSize", 14)
		if got := users["bob"].Settings["fontSize"]; got != 14 {
			t.Errorf("UpdateUserSetting() = %v, want 14", got)
		}
		// Verify existing setting still exists
		if users["bob"].Settings["theme"] != 1 {
			t.Error("Existing settings should not be affected")
		}
	})

	t.Run("Multiple updates to same user", func(t *testing.T) {
		users := map[string]Profile{
			"charlie": {Name: "Charlie", Settings: map[string]int{}},
		}
		UpdateUserSetting(users, "charlie", "theme", 1)
		UpdateUserSetting(users, "charlie", "fontSize", 12)
		UpdateUserSetting(users, "charlie", "volume", 50)

		if len(users["charlie"].Settings) != 3 {
			t.Errorf("User should have 3 settings, got %v", len(users["charlie"].Settings))
		}
	})
}

// ============================================================================
// DRILL 4 TESTS: Double Nesting (Music System Pattern)
// ============================================================================

func TestDrill4_AddSongToPlaylist(t *testing.T) {
	t.Run("Add song to playlist", func(t *testing.T) {
		lib := &Library{
			Users: map[string]User{
				"alice": {
					Playlists: map[string]Playlist{
						"favorites": {Songs: []Song{{Title: "Song1"}}},
					},
				},
			},
		}

		AddSongToPlaylist(lib, "alice", "favorites", Song{Title: "Song2"})

		count := GetPlaylistSongCount(lib, "alice", "favorites")
		if count != 2 {
			t.Errorf("GetPlaylistSongCount() = %v, want 2", count)
		}

		// Verify the songs are actually in the playlist
		songs := lib.Users["alice"].Playlists["favorites"].Songs
		if len(songs) != 2 {
			t.Errorf("Playlist should have 2 songs, got %v", len(songs))
		}
		if songs[1].Title != "Song2" {
			t.Errorf("Second song title = %v, want Song2", songs[1].Title)
		}
	})

	t.Run("Add multiple songs", func(t *testing.T) {
		lib := &Library{
			Users: map[string]User{
				"bob": {
					Playlists: map[string]Playlist{
						"workout": {Songs: []Song{}},
					},
				},
			},
		}

		AddSongToPlaylist(lib, "bob", "workout", Song{Title: "Song1"})
		AddSongToPlaylist(lib, "bob", "workout", Song{Title: "Song2"})
		AddSongToPlaylist(lib, "bob", "workout", Song{Title: "Song3"})

		count := GetPlaylistSongCount(lib, "bob", "workout")
		if count != 3 {
			t.Errorf("After adding 3 songs, count = %v, want 3", count)
		}
	})

	t.Run("Add duplicate songs allowed", func(t *testing.T) {
		lib := &Library{
			Users: map[string]User{
				"charlie": {
					Playlists: map[string]Playlist{
						"repeat": {Songs: []Song{}},
					},
				},
			},
		}

		song := Song{Title: "FavoriteSong"}
		AddSongToPlaylist(lib, "charlie", "repeat", song)
		AddSongToPlaylist(lib, "charlie", "repeat", song)

		count := GetPlaylistSongCount(lib, "charlie", "repeat")
		if count != 2 {
			t.Errorf("Duplicates should be allowed, count = %v, want 2", count)
		}
	})
}

func TestDrill4_GetPlaylistSongCount(t *testing.T) {
	t.Run("Get count from existing playlist", func(t *testing.T) {
		lib := &Library{
			Users: map[string]User{
				"alice": {
					Playlists: map[string]Playlist{
						"favorites": {Songs: []Song{{Title: "Song1"}, {Title: "Song2"}}},
					},
				},
			},
		}

		count := GetPlaylistSongCount(lib, "alice", "favorites")
		if count != 2 {
			t.Errorf("GetPlaylistSongCount() = %v, want 2", count)
		}
	})

	t.Run("Get count from empty playlist", func(t *testing.T) {
		lib := &Library{
			Users: map[string]User{
				"bob": {
					Playlists: map[string]Playlist{
						"empty": {Songs: []Song{}},
					},
				},
			},
		}

		count := GetPlaylistSongCount(lib, "bob", "empty")
		if count != 0 {
			t.Errorf("GetPlaylistSongCount() = %v, want 0 for empty playlist", count)
		}
	})

	t.Run("User doesn't exist", func(t *testing.T) {
		lib := &Library{Users: map[string]User{}}

		count := GetPlaylistSongCount(lib, "nonexistent", "playlist")
		if count != 0 {
			t.Errorf("GetPlaylistSongCount() = %v, want 0 for nonexistent user", count)
		}
	})

	t.Run("Playlist doesn't exist", func(t *testing.T) {
		lib := &Library{
			Users: map[string]User{
				"alice": {Playlists: map[string]Playlist{}},
			},
		}

		count := GetPlaylistSongCount(lib, "alice", "nonexistent")
		if count != 0 {
			t.Errorf("GetPlaylistSongCount() = %v, want 0 for nonexistent playlist", count)
		}
	})
}

// ============================================================================
// DRILL 5 TESTS: Mixed Operations (Mastery)
// ============================================================================

func TestDrill5_CreateUserWithPlaylist(t *testing.T) {
	lib := &Library{Users: map[string]User{}}

	CreateUserWithPlaylist(lib, "newUser", "myPlaylist")

	user, exists := lib.Users["newUser"]
	if !exists {
		t.Fatal("User should be created")
	}

	if user.Playlists == nil {
		t.Fatal("User.Playlists should be initialized")
	}

	playlist, exists := user.Playlists["myPlaylist"]
	if !exists {
		t.Fatal("Playlist should be created")
	}

	if playlist.Songs == nil {
		t.Error("Playlist.Songs should be initialized (can be empty slice)")
	}

	CreateUserWithPlaylist(lib, "newUser", "myPlaylist2")

	if len(user.Playlists) != 2 {
		t.Error("Second playlist was not added correctly")
	}
}

func TestDrill5_ModifySongInPlaylist(t *testing.T) {
	t.Run("Modify song title", func(t *testing.T) {
		lib := &Library{
			Users: map[string]User{
				"alice": {
					Playlists: map[string]Playlist{
						"favorites": {
							Songs: []Song{
								{Title: "OldTitle1"},
								{Title: "OldTitle2"},
								{Title: "OldTitle3"},
							},
						},
					},
				},
			},
		}

		ModifySongInPlaylist(lib, "alice", "favorites", 1, "NewTitle2")

		songs := lib.Users["alice"].Playlists["favorites"].Songs
		if songs[1].Title != "NewTitle2" {
			t.Errorf("Song title = %v, want NewTitle2", songs[1].Title)
		}
		// Verify other songs unchanged
		if songs[0].Title != "OldTitle1" {
			t.Error("Other songs should not be modified")
		}
		if songs[2].Title != "OldTitle3" {
			t.Error("Other songs should not be modified")
		}
	})
}

func TestDrill5_RemovePlaylist(t *testing.T) {
	t.Run("Remove existing playlist", func(t *testing.T) {
		lib := &Library{
			Users: map[string]User{
				"bob": {
					Playlists: map[string]Playlist{
						"workout": {Songs: []Song{{Title: "Song1"}}},
						"chill":   {Songs: []Song{{Title: "Song2"}}},
					},
				},
			},
		}

		RemovePlaylist(lib, "bob", "workout")

		user := lib.Users["bob"]
		if _, exists := user.Playlists["workout"]; exists {
			t.Error("Playlist should be removed")
		}
		if _, exists := user.Playlists["chill"]; !exists {
			t.Error("Other playlists should remain")
		}
	})
}

func TestDrill5_IntegrationTest(t *testing.T) {
	// This test combines all operations
	lib := &Library{Users: map[string]User{}}

	// 1. Create user with playlist
	CreateUserWithPlaylist(lib, "testUser", "playlist1")

	// 2. Add songs
	AddSongToPlaylist(lib, "testUser", "playlist1", Song{Title: "Song1"})
	AddSongToPlaylist(lib, "testUser", "playlist1", Song{Title: "Song2"})
	AddSongToPlaylist(lib, "testUser", "playlist1", Song{Title: "Song3"})

	// 3. Verify count
	count := GetPlaylistSongCount(lib, "testUser", "playlist1")
	if count != 3 {
		t.Errorf("Count = %v, want 3", count)
	}

	// 4. Modify a song
	ModifySongInPlaylist(lib, "testUser", "playlist1", 1, "ModifiedSong")

	// 5. Verify modification
	songs := lib.Users["testUser"].Playlists["playlist1"].Songs
	if songs[1].Title != "ModifiedSong" {
		t.Errorf("Modified song title = %v, want ModifiedSong", songs[1].Title)
	}

	// 6. Create another playlist
	CreateUserWithPlaylist(lib, "testUser", "playlist2")

	// 7. Verify both playlists exist
	user := lib.Users["testUser"]
	if len(user.Playlists) != 2 {
		t.Errorf("User should have 2 playlists, got %v", len(user.Playlists))
	}

	// 8. Remove one playlist
	RemovePlaylist(lib, "testUser", "playlist1")

	// 9. Verify removal
	user = lib.Users["testUser"]
	if len(user.Playlists) != 1 {
		t.Errorf("After removal, user should have 1 playlist, got %v", len(user.Playlists))
	}
	if _, exists := user.Playlists["playlist2"]; !exists {
		t.Error("playlist2 should still exist")
	}
}
