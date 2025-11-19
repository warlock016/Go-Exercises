package music_playlist

import (
	"fmt"
	"reflect"
	"strings"
)

type Library struct {
	Users map[string]User
	Songs map[string]Song
}

type User struct {
	ID            string
	Username      string
	Playlists     map[string]Playlist // [playlistID]Playlist{}
	PlaylistOrder []string
}

type Playlist struct {
	ID        string
	Name      string
	Songs     []Song
	SongOrder int
}

type Song struct {
	ID       string
	Title    string
	Artist   string
	Duration int
}

func NewMusicSystem() *Library {
	return &Library{
		Users: make(map[string]User),
		Songs: make(map[string]Song),
	}
}

func (lib *Library) AddSong(songID string, title string, artist string, duration int) error {

	if _, exists := lib.Songs[songID]; exists {
		return fmt.Errorf("song already exists")
	}

	if duration <= 0 {
		return fmt.Errorf("invalid duration")
	}

	newSong := Song{
		ID:       songID,
		Title:    title,
		Artist:   artist,
		Duration: duration,
	}

	lib.Songs[songID] = newSong

	return nil
}

func (lib *Library) RemoveSong(songID string) error {

	if _, exists := lib.Songs[songID]; !exists {
		return fmt.Errorf("song does not exist")
	}

	delete(lib.Songs, songID)

	allUsers := lib.Users

	for i, u := range allUsers {
		singleUser := u
		for j, p := range u.Playlists {

			singlePlaylist := p
			songList := []Song{}

			for _, s := range p.Songs {
				if s.ID != songID {
					songList = append(songList, s)
				}
			}

			singlePlaylist.Songs = songList
			singlePlaylist.SongOrder = len(singlePlaylist.Songs)
			singleUser.Playlists[j] = singlePlaylist
		}
		allUsers[i] = singleUser
	}

	fmt.Println(reflect.DeepEqual(lib.Users, allUsers))
	// lib.Users = allUsers

	return nil
}

func (lib *Library) GetSong(songID string) (Song, error) {

	song, exists := lib.Songs[songID]

	if exists {
		return song, nil
	}

	return Song{}, fmt.Errorf("song not available")
}

func (lib *Library) FindSongsByArtist(artist string) []Song {

	retrievedSongs := []Song{}

	for _, v := range lib.Songs {
		if strings.Contains(v.Artist, artist) && len([]rune(artist)) != 0 {
			retrievedSongs = append(retrievedSongs, v)
		}
	}

	return retrievedSongs
}

func (lib *Library) SearchSongsByTitle(title string) []Song {

	foundSongs := []Song{}

	for _, v := range lib.Songs {
		if strings.Contains(strings.ToLower(v.Title), strings.ToLower(title)) {
			foundSongs = append(foundSongs, v)
		}
	}

	return foundSongs
}

func (lib *Library) CreateUser(userID string, userName string) error {

	if _, exists := lib.Users[userID]; exists {
		return fmt.Errorf("user already exists")
	}

	newUser := User{
		ID:            userID,
		Username:      userName,
		Playlists:     map[string]Playlist{},
		PlaylistOrder: []string{},
	}

	lib.Users[userID] = newUser

	return nil
}

func (lib *Library) DeleteUser(userID string) error {

	if _, exists := lib.Users[userID]; !exists {
		return fmt.Errorf("user does not exist")
	}

	delete(lib.Users, userID)

	return nil
}

func (lib *Library) CreatePlaylist(userID string, playlistID string, playlistName string) error {

	if _, exists := lib.Users[userID]; !exists {
		return fmt.Errorf("userID %s does not exist", userID)
	}

	if _, exists := lib.Users[userID].Playlists[playlistID]; exists {
		return fmt.Errorf("playlistID %s already exists for userID %s", playlistID, userID)
	}

	currentUser := lib.Users[userID]

	newPlaylist := Playlist{
		ID:    playlistID,
		Name:  playlistName,
		Songs: make([]Song, 0),
		// SongOrder: 0,
	}

	currentUser.Playlists[playlistID] = newPlaylist
	currentUser.PlaylistOrder = append(currentUser.PlaylistOrder, playlistID)
	lib.Users[userID] = currentUser

	return nil
}

func (lib *Library) DeletePlaylist(userID string, playlistID string) error {

	if _, exists := lib.Users[userID]; !exists {
		return fmt.Errorf("userID %s does not exist", userID)
	}

	if _, exists := lib.Users[userID].Playlists[playlistID]; !exists {
		return fmt.Errorf("playlistID %s does not exist for userID %s", playlistID, userID)
	}

	currentUser := lib.Users[userID]
	delete(currentUser.Playlists, playlistID)

	newOrder := []string{}

	for _, v := range currentUser.PlaylistOrder {
		if v != playlistID {
			newOrder = append(newOrder, v)
		}
	}

	currentUser.PlaylistOrder = newOrder
	lib.Users[userID] = currentUser

	return nil
}

func (lib *Library) AddSongToPlaylist(userID string, playlistID string, songID string) error {

	if _, exists := lib.Songs[songID]; !exists {
		return fmt.Errorf("songID %s does not exist", songID)
	}

	if _, exists := lib.Users[userID]; !exists {
		return fmt.Errorf("userID %s does not exist", userID)
	}

	if _, exists := lib.Users[userID].Playlists[playlistID]; !exists {
		return fmt.Errorf("playlistID %s does not exist for userId %s", playlistID, userID)
	}

	if len(lib.Users[userID].Playlists[playlistID].Songs) >= 1000 {
		return fmt.Errorf("max. playlist size reached")
	}

	currentUser := lib.Users[userID]
	currentUserPlaylist := currentUser.Playlists[playlistID]
	foundSong := lib.Songs[songID]

	currentUserPlaylist.Songs = append(currentUserPlaylist.Songs, foundSong)

	currentUser.Playlists[playlistID] = currentUserPlaylist
	lib.Users[userID] = currentUser

	return nil
}

func (lib *Library) RemoveSongFromPlaylist(userID string, playlistID string, songPos int) error {

	if _, exists := lib.Users[userID]; !exists {
		return fmt.Errorf("userID %s does not exist", userID)
	}

	if _, exists := lib.Users[userID].Playlists[playlistID]; !exists {
		return fmt.Errorf("playlistID %s does not exist for userID %s", playlistID, userID)
	}

	if len(lib.Users[userID].Playlists[playlistID].Songs) == 0 {
		return fmt.Errorf("song list is empty")
	}

	if songPos < 0 || songPos >= len(lib.Users[userID].Playlists[playlistID].Songs) {
		return fmt.Errorf("invalid song position %d in playlistId %s for userId %s", songPos, playlistID, userID)
	}

	currentUser := lib.Users[userID]
	userPlaylist := currentUser.Playlists[playlistID]
	userPlaylist.SongOrder -= 1

	newPlaylistSongs := []Song{}

	for i, v := range userPlaylist.Songs {
		if i != songPos {
			newPlaylistSongs = append(newPlaylistSongs, v)
		}
	}

	userPlaylist.Songs = newPlaylistSongs
	currentUser.Playlists[playlistID] = userPlaylist
	lib.Users[userID] = currentUser

	return nil
}

func (lib *Library) GetUserPlaylists(userID string) map[string]Playlist {

	result := map[string]Playlist{}

	if _, exists := lib.Users[userID]; !exists {
		return result
	}

	result = lib.Users[userID].Playlists

	return result
}

func (lib *Library) GetMostRecentPlaylist(userID string) (Playlist, error) {

	result := Playlist{}

	if _, exists := lib.Users[userID]; !exists {
		return result, fmt.Errorf("userId %s does not exist", userID)
	}

	if len(lib.Users[userID].PlaylistOrder) == 0 {
		return result, fmt.Errorf("no playlist found for userId %s", userID)
	}

	pos := len(lib.Users[userID].PlaylistOrder) - 1
	lastPlaylistId := lib.Users[userID].PlaylistOrder[pos]
	result = lib.Users[userID].Playlists[lastPlaylistId]

	return result, nil
}

func (lib *Library) GetPlaylistSongs(userID string, playlistID string) []Song {

	result := []Song{}

	if _, exists := lib.Users[userID]; !exists {
		return nil
	}

	if _, exists := lib.Users[userID].Playlists[playlistID]; !exists {
		return nil
	}

	userPlaylist := lib.Users[userID].Playlists[playlistID].Songs

	result = append(result, userPlaylist...)

	return result
}

func (lib *Library) GetPlaylistDuration(userID string, playlistID string) int {

	result := 0

	if _, exists := lib.Users[userID]; !exists {
		return result
	}

	if _, exists := lib.Users[userID].Playlists[playlistID]; !exists {
		return result
	}

	currentPlaylist := lib.Users[userID].Playlists[playlistID]

	if len(currentPlaylist.Songs) == 0 {
		return result
	}

	for _, v := range currentPlaylist.Songs {
		result += v.Duration
	}

	return result
}

func (lib *Library) PlayNext(userID string, playlistID string) (Song, error) {

	next := Song{}

	if _, exists := lib.Users[userID]; !exists {
		return next, fmt.Errorf("userId %s does not exist", userID)
	}

	if _, exists := lib.Users[userID].Playlists[playlistID]; !exists {
		return next, fmt.Errorf("playlistId %s for userId %s does not exist", playlistID, userID)
	}

	currentUser := lib.Users[userID]
	userPlaylist := currentUser.Playlists[playlistID]
	playlistLength := len(userPlaylist.Songs)

	if playlistLength == 0 {
		return next, fmt.Errorf("playlistId %s for userId %s is empty", playlistID, userID)
	}

	if userPlaylist.SongOrder >= playlistLength {
		// playlistPos = 0
		return next, fmt.Errorf("reached end of playlist")
	}

	next = userPlaylist.Songs[userPlaylist.SongOrder]
	userPlaylist.SongOrder++

	currentUser.Playlists[playlistID] = userPlaylist

	lib.Users[userID] = currentUser

	return next, nil
}

func (lib *Library) ShufflePlaylist(userID string, playlistID string) {

	if _, exists := lib.Users[userID]; !exists {
		return
	}

	if _, exists := lib.Users[userID].Playlists[playlistID]; !exists {
		return
	}

	currentUser := lib.Users[userID]
	userPlaylist := lib.Users[userID].Playlists[playlistID]
	shuffle := make(map[Song]struct{})
	result := []Song{}

	for _, v := range userPlaylist.Songs {
		shuffle[v] = struct{}{}
	}

	for i := range shuffle {
		result = append(result, i)
	}

	userPlaylist.Songs = result
	currentUser.Playlists[playlistID] = userPlaylist
	lib.Users[userID] = currentUser
}

func (lib *Library) FindPlaylistsWithSong(songID string) []Song {
	result := []Song{}

	for _, u := range lib.Users {
		for _, p := range u.Playlists {
			for _, s := range p.Songs {
				if s.ID == songID {
					result = append(result, s)
				}
			}
		}
	}

	return result
}
