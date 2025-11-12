package iterate

import (
	"reflect"
	"testing"
)

func TestCollectBytes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []byte
	}{
		{"ASCII", "hi", []byte{104, 105}},
		{"Empty", "", []byte{}},
		{"With accent", "café", []byte{99, 97, 102, 195, 169}}, // é is bytes 195, 169
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CollectBytes(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CollectBytes(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCollectRunes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []rune
	}{
		{"ASCII", "hi", []rune{'h', 'i'}},
		{"Empty", "", []rune{}},
		{"With accent", "café", []rune{'c', 'a', 'f', 'é'}},
		{"Emoji", "👍🎉", []rune{'👍', '🎉'}},
		{"Chinese", "世界", []rune{'世', '界'}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CollectRunes(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CollectRunes(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestGetRuneAt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		index   int
		want    rune
		wantErr bool
	}{
		{"First char", "hello", 0, 'h', false},
		{"Last char", "hello", 4, 'o', false},
		{"Middle char", "hello", 2, 'l', false},
		{"Out of bounds positive", "hello", 5, 0, true},
		{"Out of bounds negative", "hello", -1, 0, true},
		{"Accented char", "café", 3, 'é', false},
		{"Emoji first", "👍🎉", 0, '👍', false},
		{"Emoji second", "👍🎉", 1, '🎉', false},
		{"Emoji out of bounds", "👍🎉", 2, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRuneAt(tt.input, tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRuneAt(%q, %d) error = %v, wantErr %v", tt.input, tt.index, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("GetRuneAt(%q, %d) = %q, want %q", tt.input, tt.index, got, tt.want)
			}
		})
	}
}

// Demonstrates the danger of indexing strings directly
func TestDirectIndexingIsBroken(t *testing.T) {
	s := "café"

	// Wrong way - byte indexing
	// s[3] is NOT 'é', it's the first byte of é!
	byteAt3 := s[3] // This is 195 (first byte of é in UTF-8)

	// Right way - rune indexing
	runeAt3, _ := GetRuneAt(s, 3) // This is 'é'

	t.Logf("String: %q", s)
	t.Logf("s[3] (byte) = %d (not the character!)", byteAt3)
	t.Logf("GetRuneAt(s, 3) = %q (correct character)", runeAt3)

	if byteAt3 == 'é' {
		t.Error("Surprisingly, s[3] == 'é', but this should not happen with UTF-8!")
	}

	if runeAt3 != 'é' {
		t.Errorf("GetRuneAt should return 'é', got %q", runeAt3)
	}
}
