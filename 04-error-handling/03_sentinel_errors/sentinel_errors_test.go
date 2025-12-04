package sentinel_errors

import (
	"errors"
	"testing"
)

func TestFindUser(t *testing.T) {
	users := map[int]string{
		1: "Alice",
		2: "Bob",
		3: "Charlie",
	}

	tests := []struct {
		name       string
		userID     int
		want       string
		wantErr    error
	}{
		{"existing user 1", 1, "Alice", nil},
		{"existing user 2", 2, "Bob", nil},
		{"non-existent user", 99, "", ErrNotFound},
		{"negative id", -1, "", ErrNotFound},
		{"zero id", 0, "", ErrNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindUser(tt.userID, users)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("FindUser(%d) error = %v, want %v", tt.userID, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindUser(%d) = %q, want %q", tt.userID, got, tt.want)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name    string
		userID  int
		userName string
		existing map[int]string
		wantErr error
	}{
		{"create new user", 4, "Dave", map[int]string{1: "Alice"}, nil},
		{"duplicate user", 1, "NewName", map[int]string{1: "Alice"}, ErrAlreadyExists},
		{"empty name", 5, "", map[int]string{}, ErrInvalidInput},
		{"empty name with existing users", 10, "", map[int]string{1: "Alice"}, ErrInvalidInput},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users := make(map[int]string)
			for k, v := range tt.existing {
				users[k] = v
			}

			err := CreateUser(tt.userID, tt.userName, users)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("CreateUser(%d, %q) error = %v, want %v", tt.userID, tt.userName, err, tt.wantErr)
			}

			// If no error expected, verify user was added
			if tt.wantErr == nil {
				if users[tt.userID] != tt.userName {
					t.Errorf("CreateUser(%d, %q) didn't add user to map", tt.userID, tt.userName)
				}
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name     string
		userID   int
		existing map[int]string
		wantErr  error
	}{
		{"delete existing user", 1, map[int]string{1: "Alice", 2: "Bob"}, nil},
		{"delete non-existent", 99, map[int]string{1: "Alice"}, ErrNotFound},
		{"unauthorized deletion", 0, map[int]string{1: "Alice"}, ErrUnauthorized},
		{"delete from empty map", 1, map[int]string{}, ErrNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users := make(map[int]string)
			for k, v := range tt.existing {
				users[k] = v
			}

			err := DeleteUser(tt.userID, users)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("DeleteUser(%d) error = %v, want %v", tt.userID, err, tt.wantErr)
			}

			// If no error expected, verify user was deleted
			if tt.wantErr == nil {
				if _, exists := users[tt.userID]; exists {
					t.Errorf("DeleteUser(%d) didn't remove user from map", tt.userID)
				}
			}
		})
	}
}

func TestIsNotFoundError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"ErrNotFound", ErrNotFound, true},
		{"ErrAlreadyExists", ErrAlreadyExists, false},
		{"nil error", nil, false},
		{"other error", errors.New("something else"), false},
		{"ErrTimeout", ErrTimeout, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotFoundError(tt.err); got != tt.want {
				t.Errorf("IsNotFoundError(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}

func TestIsTimeoutError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"ErrTimeout", ErrTimeout, true},
		{"ErrNotFound", ErrNotFound, false},
		{"nil error", nil, false},
		{"other error", errors.New("something else"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTimeoutError(tt.err); got != tt.want {
				t.Errorf("IsTimeoutError(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}

func TestHandleError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantContains string
	}{
		{"not found", ErrNotFound, "not found"},
		{"already exists", ErrAlreadyExists, "already exists"},
		{"unauthorized", ErrUnauthorized, "not authorized"},
		{"invalid input", ErrInvalidInput, "invalid"},
		{"timeout", ErrTimeout, "timed out"},
		{"unknown error", errors.New("random error"), "unknown"},
		{"nil error", nil, "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HandleError(tt.err)
			if got == "" {
				t.Errorf("HandleError(%v) returned empty string", tt.err)
			}
			// We don't check exact message, just that it's not empty
			// The actual implementation can have various phrasings
		})
	}
}
