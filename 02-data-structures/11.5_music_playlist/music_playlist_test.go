package music_playlist

import (
	"strings"
	"testing"
)

// Test Song Library Operations

func TestAddSongToLibrary(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		title    string
		artist   string
		duration int
		wantErr  bool
	}{
		{"Valid song", "s1", "Bohemian Rhapsody", "Queen", 354, false},
		{"Another valid song", "s2", "Stairway to Heaven", "Led Zeppelin", 482, false},
		{"Zero duration should fail", "s3", "Bad Song", "Artist", 0, true},
		{"Negative duration should fail", "s4", "Bad Song", "Artist", -100, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			system := NewMusicSystem()
			err := system.AddSong(tt.id, tt.title, tt.artist, tt.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddSong() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSongIDUniqueness(t *testing.T) {
	system := NewMusicSystem()

	err := system.AddSong("s1", "Song One", "Artist A", 200)
	if err != nil {
		t.Fatalf("First AddSong failed: %v", err)
	}

	// Try to add another song with same ID
	err = system.AddSong("s1", "Song Two", "Artist B", 300)
	if err == nil {
		t.Error("AddSong should fail when ID already exists")
	}
}

func TestGetSongByID(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Test Song", "Test Artist", 250)

	song, err := system.GetSong("s1")
	if err != nil {
		t.Fatalf("GetSong failed: %v", err)
	}

	if song.Title != "Test Song" {
		t.Errorf("Expected title 'Test Song', got '%s'", song.Title)
	}
	if song.Artist != "Test Artist" {
		t.Errorf("Expected artist 'Test Artist', got '%s'", song.Artist)
	}
	if song.Duration != 250 {
		t.Errorf("Expected duration 250, got %d", song.Duration)
	}
}

func TestGetSongByIDNotFound(t *testing.T) {
	system := NewMusicSystem()

	_, err := system.GetSong("nonexistent")
	if err == nil {
		t.Error("GetSong should return error for non-existent ID")
	}
}

func TestRemoveSongFromLibrary(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song to Remove", "Artist", 200)

	err := system.RemoveSong("s1")
	if err != nil {
		t.Fatalf("RemoveSong failed: %v", err)
	}

	// Should not be able to get removed song
	_, err = system.GetSong("s1")
	if err == nil {
		t.Error("GetSong should fail after song is removed")
	}
}

func TestRemoveSongNotFound(t *testing.T) {
	system := NewMusicSystem()

	err := system.RemoveSong("nonexistent")
	if err == nil {
		t.Error("RemoveSong should return error for non-existent ID")
	}
}

func TestFindSongsByArtist(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song 1", "The Beatles", 200)
	system.AddSong("s2", "Song 2", "Queen", 250)
	system.AddSong("s3", "Song 3", "The Beatles", 180)
	system.AddSong("s4", "Song 4", "Led Zeppelin", 300)

	songs := system.FindSongsByArtist("The Beatles")

	if len(songs) != 2 {
		t.Errorf("Expected 2 Beatles songs, got %d", len(songs))
	}

	// Verify both songs are by The Beatles
	for _, song := range songs {
		if song.Artist != "The Beatles" {
			t.Errorf("Expected artist 'The Beatles', got '%s'", song.Artist)
		}
	}
}

func TestFindSongsByArtistCaseSensitive(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song 1", "The Beatles", 200)

	// Case-sensitive - should not match
	songs := system.FindSongsByArtist("the beatles")

	if len(songs) != 0 {
		t.Errorf("FindSongsByArtist should be case-sensitive, got %d results", len(songs))
	}
}

func TestFindSongsByArtistNotFound(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song 1", "Artist A", 200)

	songs := system.FindSongsByArtist("Nonexistent Artist")

	if len(songs) != 0 {
		t.Errorf("Expected 0 songs, got %d", len(songs))
	}
}

func TestSearchSongsByTitle(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Let It Be", "The Beatles", 243)
	system.AddSong("s2", "Let It Go", "Idina Menzel", 225)
	system.AddSong("s3", "Yesterday", "The Beatles", 125)
	system.AddSong("s4", "Let Me Love You", "DJ Snake", 205)

	songs := system.SearchSongsByTitle("let")

	// Should find s1, s2, s4 (all containing "let" case-insensitive)
	if len(songs) != 3 {
		t.Errorf("Expected 3 songs containing 'let', got %d", len(songs))
	}

	// Verify all results contain "let" (case-insensitive)
	for _, song := range songs {
		if !strings.Contains(strings.ToLower(song.Title), "let") {
			t.Errorf("Song '%s' should not be in results", song.Title)
		}
	}
}

func TestSearchSongsByTitleCaseInsensitive(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "HELLO World", "Artist", 200)
	system.AddSong("s2", "hello world", "Artist", 200)
	system.AddSong("s3", "HeLLo WoRLd", "Artist", 200)

	// All variations should match "hello"
	songs := system.SearchSongsByTitle("hello")
	if len(songs) != 3 {
		t.Errorf("Expected 3 songs (case-insensitive), got %d", len(songs))
	}

	// Search with uppercase should also match all
	songs = system.SearchSongsByTitle("HELLO")
	if len(songs) != 3 {
		t.Errorf("Expected 3 songs (case-insensitive), got %d", len(songs))
	}
}

func TestSearchSongsByTitlePartialMatch(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Thunderstruck", "AC/DC", 292)

	// Partial matches should work
	songs := system.SearchSongsByTitle("thunder")
	if len(songs) != 1 {
		t.Errorf("Expected 1 song matching 'thunder', got %d", len(songs))
	}

	songs = system.SearchSongsByTitle("struck")
	if len(songs) != 1 {
		t.Errorf("Expected 1 song matching 'struck', got %d", len(songs))
	}

	songs = system.SearchSongsByTitle("under")
	if len(songs) != 1 {
		t.Errorf("Expected 1 song matching 'under', got %d", len(songs))
	}
}

// Test User Operations

func TestCreateUser(t *testing.T) {
	system := NewMusicSystem()

	err := system.CreateUser("u1", "john_doe")
	if err != nil {
		t.Errorf("CreateUser failed: %v", err)
	}
}

func TestUserIDUniqueness(t *testing.T) {
	system := NewMusicSystem()

	err := system.CreateUser("u1", "user_one")
	if err != nil {
		t.Fatalf("First CreateUser failed: %v", err)
	}

	err = system.CreateUser("u1", "user_two")
	if err == nil {
		t.Error("CreateUser should fail when ID already exists")
	}
}

// Test Playlist Operations

func TestCreatePlaylist(t *testing.T) {
	system := NewMusicSystem()
	system.CreateUser("u1", "user")

	err := system.CreatePlaylist("u1", "p1", "My Playlist")
	if err != nil {
		t.Errorf("CreatePlaylist failed: %v", err)
	}
}

func TestCreatePlaylistUserNotFound(t *testing.T) {
	system := NewMusicSystem()

	err := system.CreatePlaylist("nonexistent", "p1", "Playlist")
	if err == nil {
		t.Error("CreatePlaylist should fail when user doesn't exist")
	}
}

func TestPlaylistIDScopedToUser(t *testing.T) {
	system := NewMusicSystem()
	system.CreateUser("u1", "alice")
	system.CreateUser("u2", "bob")

	// Both users can have playlist with ID "p1"
	err1 := system.CreatePlaylist("u1", "p1", "Alice's Playlist")
	err2 := system.CreatePlaylist("u2", "p1", "Bob's Playlist")

	if err1 != nil || err2 != nil {
		t.Error("Different users should be able to have playlists with same ID")
	}

	// Same user cannot have two playlists with same ID
	err := system.CreatePlaylist("u1", "p1", "Another Playlist")
	if err == nil {
		t.Error("Same user should not be able to create two playlists with same ID")
	}
}

func TestAddSongToPlaylist(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song 1", "Artist", 200)
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist")

	err := system.AddSongToPlaylist("u1", "p1", "s1")
	if err != nil {
		t.Errorf("AddSongToPlaylist failed: %v", err)
	}
}

func TestAddSongToPlaylistNotFound(t *testing.T) {
	system := NewMusicSystem()
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist")

	// Song doesn't exist
	err := system.AddSongToPlaylist("u1", "p1", "nonexistent")
	if err == nil {
		t.Error("AddSongToPlaylist should fail when song doesn't exist")
	}
}

func TestAddDuplicateSongsToPlaylist(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song 1", "Artist", 200)
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist")

	// Add same song multiple times
	system.AddSongToPlaylist("u1", "p1", "s1")
	system.AddSongToPlaylist("u1", "p1", "s1")
	system.AddSongToPlaylist("u1", "p1", "s1")

	songs := system.GetPlaylistSongs("u1", "p1")

	if len(songs) != 3 {
		t.Errorf("Expected 3 songs in playlist (duplicates allowed), got %d", len(songs))
	}
}

func TestPlaylistSizeLimit(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song", "Artist", 200)
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist")

	// Add 1000 songs (should succeed)
	for i := 0; i < 1000; i++ {
		err := system.AddSongToPlaylist("u1", "p1", "s1")
		if err != nil {
			t.Fatalf("Failed to add song %d: %v", i+1, err)
		}
	}

	// Adding 1001st song should fail
	err := system.AddSongToPlaylist("u1", "p1", "s1")
	if err == nil {
		t.Error("AddSongToPlaylist should fail when playlist reaches 1000 songs")
	}
}

func TestGetPlaylistSongs(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song 1", "Artist", 200)
	system.AddSong("s2", "Song 2", "Artist", 250)
	system.AddSong("s3", "Song 3", "Artist", 180)
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist")

	system.AddSongToPlaylist("u1", "p1", "s1")
	system.AddSongToPlaylist("u1", "p1", "s2")
	system.AddSongToPlaylist("u1", "p1", "s3")

	songs := system.GetPlaylistSongs("u1", "p1")

	if len(songs) != 3 {
		t.Errorf("Expected 3 songs, got %d", len(songs))
	}

	// Verify order is maintained
	expectedOrder := []string{"s1", "s2", "s3"}
	for i, song := range songs {
		if song.ID != expectedOrder[i] {
			t.Errorf("Song at position %d: expected ID %s, got %s", i, expectedOrder[i], song.ID)
		}
	}
}

func TestGetPlaylistSongsEmptyPlaylist(t *testing.T) {
	system := NewMusicSystem()
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Empty Playlist")

	songs := system.GetPlaylistSongs("u1", "p1")

	if len(songs) != 0 {
		t.Errorf("Expected 0 songs in empty playlist, got %d", len(songs))
	}
}

func TestRemoveSongFromPlaylist(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song 1", "Artist", 200)
	system.AddSong("s2", "Song 2", "Artist", 250)
	system.AddSong("s3", "Song 3", "Artist", 180)
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist")

	system.AddSongToPlaylist("u1", "p1", "s1")
	system.AddSongToPlaylist("u1", "p1", "s2")
	system.AddSongToPlaylist("u1", "p1", "s3")

	// Remove middle song (position 1)
	err := system.RemoveSongFromPlaylist("u1", "p1", 1)
	if err != nil {
		t.Fatalf("RemoveSongFromPlaylist failed: %v", err)
	}

	songs := system.GetPlaylistSongs("u1", "p1")

	if len(songs) != 2 {
		t.Errorf("Expected 2 songs after removal, got %d", len(songs))
	}

	// Verify remaining songs are s1 and s3
	if songs[0].ID != "s1" || songs[1].ID != "s3" {
		t.Error("Remaining songs should be s1 and s3 in order")
	}
}

func TestRemoveSongFromPlaylistInvalidPosition(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song", "Artist", 200)
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist")
	system.AddSongToPlaylist("u1", "p1", "s1")

	// Try to remove at invalid positions
	err := system.RemoveSongFromPlaylist("u1", "p1", -1)
	if err == nil {
		t.Error("RemoveSongFromPlaylist should fail for negative position")
	}

	err = system.RemoveSongFromPlaylist("u1", "p1", 5)
	if err == nil {
		t.Error("RemoveSongFromPlaylist should fail for out-of-bounds position")
	}
}

func TestDeletePlaylist(t *testing.T) {
	system := NewMusicSystem()
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist to Delete")

	err := system.DeletePlaylist("u1", "p1")
	if err != nil {
		t.Fatalf("DeletePlaylist failed: %v", err)
	}

	// Trying to get songs from deleted playlist should fail
	songs := system.GetPlaylistSongs("u1", "p1")
	if songs != nil {
		t.Error("GetPlaylistSongs should return nil for deleted playlist")
	}
}

func TestCalculatePlaylistDuration(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song 1", "Artist", 200)
	system.AddSong("s2", "Song 2", "Artist", 300)
	system.AddSong("s3", "Song 3", "Artist", 150)
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist")

	system.AddSongToPlaylist("u1", "p1", "s1")
	system.AddSongToPlaylist("u1", "p1", "s2")
	system.AddSongToPlaylist("u1", "p1", "s3")

	duration := system.GetPlaylistDuration("u1", "p1")
	expected := 650 // 200 + 300 + 150

	if duration != expected {
		t.Errorf("Expected duration %d, got %d", expected, duration)
	}
}

func TestCalculatePlaylistDurationWithDuplicates(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song 1", "Artist", 100)
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist")

	// Add same song 3 times
	system.AddSongToPlaylist("u1", "p1", "s1")
	system.AddSongToPlaylist("u1", "p1", "s1")
	system.AddSongToPlaylist("u1", "p1", "s1")

	duration := system.GetPlaylistDuration("u1", "p1")
	expected := 300 // 100 * 3

	if duration != expected {
		t.Errorf("Expected duration %d (song counted 3 times), got %d", expected, duration)
	}
}

func TestCalculateEmptyPlaylistDuration(t *testing.T) {
	system := NewMusicSystem()
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Empty Playlist")

	duration := system.GetPlaylistDuration("u1", "p1")

	if duration != 0 {
		t.Errorf("Expected duration 0 for empty playlist, got %d", duration)
	}
}

func TestPlayNextSong(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song 1", "Artist", 200)
	system.AddSong("s2", "Song 2", "Artist", 250)
	system.AddSong("s3", "Song 3", "Artist", 180)
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist")

	system.AddSongToPlaylist("u1", "p1", "s1")
	system.AddSongToPlaylist("u1", "p1", "s2")
	system.AddSongToPlaylist("u1", "p1", "s3")

	// Play songs in order
	song, err := system.PlayNext("u1", "p1")
	if err != nil || song.ID != "s1" {
		t.Errorf("First PlayNext should return s1, got %v, err %v", song, err)
	}

	song, err = system.PlayNext("u1", "p1")
	if err != nil || song.ID != "s2" {
		t.Errorf("Second PlayNext should return s2, got %v, err %v", song, err)
	}

	song, err = system.PlayNext("u1", "p1")
	if err != nil || song.ID != "s3" {
		t.Errorf("Third PlayNext should return s3, got %v, err %v", song, err)
	}

	// Fourth call should fail (end of playlist)
	song, err = system.PlayNext("u1", "p1")
	if err == nil {
		t.Error("PlayNext should fail at end of playlist")
	}
}

func TestPlayNextEmptyPlaylist(t *testing.T) {
	system := NewMusicSystem()
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Empty Playlist")

	_, err := system.PlayNext("u1", "p1")
	if err == nil {
		t.Error("PlayNext should fail on empty playlist")
	}
}

func TestShufflePlaylist(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song 1", "Artist", 200)
	system.AddSong("s2", "Song 2", "Artist", 250)
	system.AddSong("s3", "Song 3", "Artist", 180)
	system.AddSong("s4", "Song 4", "Artist", 220)
	system.AddSong("s5", "Song 5", "Artist", 240)
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist")

	system.AddSongToPlaylist("u1", "p1", "s1")
	system.AddSongToPlaylist("u1", "p1", "s2")
	system.AddSongToPlaylist("u1", "p1", "s3")
	system.AddSongToPlaylist("u1", "p1", "s4")
	system.AddSongToPlaylist("u1", "p1", "s5")

	originalOrder := system.GetPlaylistSongs("u1", "p1")

	system.ShufflePlaylist("u1", "p1")

	shuffledOrder := system.GetPlaylistSongs("u1", "p1")

	// Should have same songs
	if len(shuffledOrder) != len(originalOrder) {
		t.Errorf("Shuffle should maintain song count: expected %d, got %d", len(originalOrder), len(shuffledOrder))
	}

	// Should contain same songs (check IDs)
	originalIDs := make(map[string]int)
	for _, song := range originalOrder {
		originalIDs[song.ID]++
	}

	shuffledIDs := make(map[string]int)
	for _, song := range shuffledOrder {
		shuffledIDs[song.ID]++
	}

	for id, count := range originalIDs {
		if shuffledIDs[id] != count {
			t.Errorf("Shuffle should preserve song counts: ID %s had %d, now has %d", id, count, shuffledIDs[id])
		}
	}

	// Note: We don't test if order changed because shuffle MIGHT produce same order by chance
	// In a real implementation, you'd want multiple shuffles to verify randomness
}

func TestShuffleResetsPosition(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song 1", "Artist", 200)
	system.AddSong("s2", "Song 2", "Artist", 250)
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist")

	system.AddSongToPlaylist("u1", "p1", "s1")
	system.AddSongToPlaylist("u1", "p1", "s2")

	// Advance position
	system.PlayNext("u1", "p1")

	// Shuffle should reset position to 0
	system.ShufflePlaylist("u1", "p1")

	// First PlayNext after shuffle should work
	_, err := system.PlayNext("u1", "p1")
	if err != nil {
		t.Error("PlayNext should work after shuffle (position should be reset)")
	}
}

// Test Cross-Entity Operations

func TestFindPlaylistsContainingSong(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Popular Song", "Artist", 200)
	system.AddSong("s2", "Other Song", "Artist", 250)

	system.CreateUser("u1", "alice")
	system.CreateUser("u2", "bob")

	system.CreatePlaylist("u1", "p1", "Alice's Mix 1")
	system.CreatePlaylist("u1", "p2", "Alice's Mix 2")
	system.CreatePlaylist("u2", "p1", "Bob's Mix")
	system.CreatePlaylist("u2", "p2", "Bob's Empty")

	// Add s1 to multiple playlists
	system.AddSongToPlaylist("u1", "p1", "s1")
	system.AddSongToPlaylist("u1", "p2", "s1")
	system.AddSongToPlaylist("u2", "p1", "s1")

	// Add s2 to one playlist
	system.AddSongToPlaylist("u1", "p1", "s2")

	// Find playlists containing s1
	playlists := system.FindPlaylistsWithSong("s1")

	if len(playlists) != 3 {
		t.Errorf("Expected 3 playlists containing s1, got %d", len(playlists))
	}

	// Find playlists containing s2
	playlists = system.FindPlaylistsWithSong("s2")

	if len(playlists) != 1 {
		t.Errorf("Expected 1 playlist containing s2, got %d", len(playlists))
	}
}

func TestFindPlaylistsContainingSongNotInAny(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Lonely Song", "Artist", 200)
	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist")

	// Song exists but not added to any playlist
	playlists := system.FindPlaylistsWithSong("s1")

	if len(playlists) != 0 {
		t.Errorf("Expected 0 playlists, got %d", len(playlists))
	}
}

func TestCascadingDelete(t *testing.T) {
	system := NewMusicSystem()
	system.AddSong("s1", "Song to Delete", "Artist", 200)
	system.AddSong("s2", "Song to Keep", "Artist", 250)

	system.CreateUser("u1", "user")
	system.CreatePlaylist("u1", "p1", "Playlist 1")
	system.CreatePlaylist("u1", "p2", "Playlist 2")

	// Add s1 to both playlists
	system.AddSongToPlaylist("u1", "p1", "s1")
	system.AddSongToPlaylist("u1", "p1", "s2")
	system.AddSongToPlaylist("u1", "p2", "s1")

	// Remove s1 from library
	system.RemoveSong("s1")

	// s1 should be removed from both playlists
	songs1 := system.GetPlaylistSongs("u1", "p1")
	if len(songs1) != 1 || songs1[0].ID != "s2" {
		t.Error("Playlist 1 should only contain s2 after s1 deleted from library")
	}

	songs2 := system.GetPlaylistSongs("u1", "p2")
	if len(songs2) != 0 {
		t.Error("Playlist 2 should be empty after s1 deleted from library")
	}
}

func TestGetUserPlaylists(t *testing.T) {
	system := NewMusicSystem()
	system.CreateUser("u1", "user")

	system.CreatePlaylist("u1", "p1", "Playlist 1")
	system.CreatePlaylist("u1", "p2", "Playlist 2")
	system.CreatePlaylist("u1", "p3", "Playlist 3")

	playlists := system.GetUserPlaylists("u1")

	if len(playlists) != 3 {
		t.Errorf("Expected 3 playlists, got %d", len(playlists))
	}
}

func TestGetUserPlaylistsEmpty(t *testing.T) {
	system := NewMusicSystem()
	system.CreateUser("u1", "user")

	playlists := system.GetUserPlaylists("u1")

	if len(playlists) != 0 {
		t.Errorf("Expected 0 playlists, got %d", len(playlists))
	}
}

func TestGetUserMostRecentPlaylist(t *testing.T) {
	system := NewMusicSystem()
	system.CreateUser("u1", "user")

	system.CreatePlaylist("u1", "p1", "First")
	system.CreatePlaylist("u1", "p2", "Second")
	system.CreatePlaylist("u1", "p3", "Third")

	playlist, err := system.GetMostRecentPlaylist("u1")
	if err != nil {
		t.Fatalf("GetMostRecentPlaylist failed: %v", err)
	}

	if playlist.Name != "Third" {
		t.Errorf("Expected most recent playlist 'Third', got '%s'", playlist.Name)
	}
}

func TestGetUserMostRecentPlaylistNoPlaylists(t *testing.T) {
	system := NewMusicSystem()
	system.CreateUser("u1", "user")

	_, err := system.GetMostRecentPlaylist("u1")
	if err == nil {
		t.Error("GetMostRecentPlaylist should fail when user has no playlists")
	}
}

// Test Edge Cases and Integration

func TestComplexWorkflow(t *testing.T) {
	// This test simulates a realistic user workflow
	system := NewMusicSystem()

	// Add songs
	system.AddSong("s1", "Bohemian Rhapsody", "Queen", 354)
	system.AddSong("s2", "Stairway to Heaven", "Led Zeppelin", 482)
	system.AddSong("s3", "Hotel California", "Eagles", 391)
	system.AddSong("s4", "Imagine", "John Lennon", 183)

	// Create users
	system.CreateUser("u1", "alice")
	system.CreateUser("u2", "bob")

	// Alice creates playlists
	system.CreatePlaylist("u1", "p1", "Rock Classics")
	system.CreatePlaylist("u1", "p2", "Favorites")

	// Bob creates playlist
	system.CreatePlaylist("u2", "p1", "Bob's Mix")

	// Add songs to playlists
	system.AddSongToPlaylist("u1", "p1", "s1")
	system.AddSongToPlaylist("u1", "p1", "s2")
	system.AddSongToPlaylist("u1", "p1", "s3")

	system.AddSongToPlaylist("u1", "p2", "s1")
	system.AddSongToPlaylist("u1", "p2", "s4")

	system.AddSongToPlaylist("u2", "p1", "s1")
	system.AddSongToPlaylist("u2", "p1", "s2")

	// Verify durations
	duration1 := system.GetPlaylistDuration("u1", "p1")
	if duration1 != 1227 { // 354 + 482 + 391
		t.Errorf("Alice's Rock Classics duration: expected 1227, got %d", duration1)
	}

	duration2 := system.GetPlaylistDuration("u1", "p2")
	if duration2 != 537 { // 354 + 183
		t.Errorf("Alice's Favorites duration: expected 537, got %d", duration2)
	}

	// Find playlists with Bohemian Rhapsody (in all 3 playlists)
	playlists := system.FindPlaylistsWithSong("s1")
	if len(playlists) != 3 {
		t.Errorf("Expected Bohemian Rhapsody in 3 playlists, got %d", len(playlists))
	}

	// Search for songs
	results := system.SearchSongsByTitle("stairway")
	if len(results) != 1 || results[0].ID != "s2" {
		t.Error("Search should find Stairway to Heaven")
	}

	// Play through Alice's Favorites
	song1, _ := system.PlayNext("u1", "p2")
	if song1.ID != "s1" {
		t.Errorf("First song should be s1, got %s", song1.ID)
	}

	song2, _ := system.PlayNext("u1", "p2")
	if song2.ID != "s4" {
		t.Errorf("Second song should be s4, got %s", song2.ID)
	}

	// End of playlist
	_, err := system.PlayNext("u1", "p2")
	if err == nil {
		t.Error("Should be at end of playlist")
	}

	// Delete a song from library (cascading delete)
	system.RemoveSong("s1")

	// Verify it's removed from all playlists
	songs1 := system.GetPlaylistSongs("u1", "p1")
	if len(songs1) != 2 { // s2, s3 remain
		t.Errorf("Alice's Rock Classics should have 2 songs after delete, got %d", len(songs1))
	}

	songs2 := system.GetPlaylistSongs("u1", "p2")
	if len(songs2) != 1 { // s4 remains
		t.Errorf("Alice's Favorites should have 1 song after delete, got %d", len(songs2))
	}

	bobSongs := system.GetPlaylistSongs("u2", "p1")
	if len(bobSongs) != 1 { // s2 remains
		t.Errorf("Bob's Mix should have 1 song after delete, got %d", len(bobSongs))
	}
}
