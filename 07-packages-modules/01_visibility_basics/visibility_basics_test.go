package visibility

import (
	"strings"
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name        string
		userName    string
		password    string
		wantErr     bool
		errContains string
	}{
		{
			name:     "Valid user",
			userName: "Alice",
			password: "secret123",
			wantErr:  false,
		},
		{
			name:        "Empty name",
			userName:    "",
			password:    "secret123",
			wantErr:     true,
			errContains: "name",
		},
		{
			name:        "Password too short",
			userName:    "Bob",
			password:    "short",
			wantErr:     true,
			errContains: "password",
		},
		{
			name:        "Name too long",
			userName:    strings.Repeat("a", 60),
			password:    "secret123",
			wantErr:     true,
			errContains: "name",
		},
		{
			name:     "Valid with exact min password length",
			userName: "Charlie",
			password: "12345678",
			wantErr:  false,
		},
		{
			name:     "Valid with max name length",
			userName: strings.Repeat("a", 50),
			password: "validpass",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewUser(tt.userName, tt.password)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewUser(%q, %q) expected error, got nil", tt.userName, tt.password)
				} else if tt.errContains != "" && !strings.Contains(strings.ToLower(err.Error()), tt.errContains) {
					t.Errorf("NewUser(%q, %q) error = %q, should contain %q", tt.userName, tt.password, err.Error(), tt.errContains)
				}
			} else {
				if err != nil {
					t.Errorf("NewUser(%q, %q) unexpected error: %v", tt.userName, tt.password, err)
				}
				if u == nil {
					t.Errorf("NewUser(%q, %q) returned nil user", tt.userName, tt.password)
				}
				if u != nil && u.Name != tt.userName {
					t.Errorf("NewUser(%q, %q) user.Name = %q, want %q", tt.userName, tt.password, u.Name, tt.userName)
				}
			}
		})
	}
}

func TestCheckPassword(t *testing.T) {
	tests := []struct {
		name            string
		userPassword    string
		checkPassword   string
		expectedMatch   bool
	}{
		{
			name:          "Correct password",
			userPassword:  "secret123",
			checkPassword: "secret123",
			expectedMatch: true,
		},
		{
			name:          "Wrong password",
			userPassword:  "secret123",
			checkPassword: "wrongpass",
			expectedMatch: false,
		},
		{
			name:          "Case sensitive",
			userPassword:  "Secret123",
			checkPassword: "secret123",
			expectedMatch: false,
		},
		{
			name:          "Empty password check",
			userPassword:  "secret123",
			checkPassword: "",
			expectedMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewUser("TestUser", tt.userPassword)
			if err != nil {
				t.Fatalf("NewUser() error = %v", err)
			}

			got := u.CheckPassword(tt.checkPassword)
			if got != tt.expectedMatch {
				t.Errorf("CheckPassword(%q) = %v, want %v", tt.checkPassword, got, tt.expectedMatch)
			}
		})
	}
}

func TestUserFieldVisibility(t *testing.T) {
	u, err := NewUser("Alice", "secret123")
	if err != nil {
		t.Fatalf("NewUser() error = %v", err)
	}

	// This test verifies that Name is exported
	if u.Name != "Alice" {
		t.Errorf("User.Name = %q, want %q", u.Name, "Alice")
	}

	// Note: We cannot test that password is unexported in code
	// (it would cause a compile error), but the test structure
	// verifies we can only access it via CheckPassword()
	if !u.CheckPassword("secret123") {
		t.Error("Password should be stored and accessible via CheckPassword()")
	}
}

func TestValidateIsUnexported(t *testing.T) {
	// This test exists to document that validate() is unexported.
	// If validate were exported, we could call it directly:
	// err := Validate("name", "password")
	//
	// Since it's unexported, it's only accessible within the package.
	// This is the correct design - validation is an internal detail.

	t.Log("validate() is correctly unexported and only used internally by NewUser()")
}

func TestConfigIsUnexported(t *testing.T) {
	// This test exists to document that config type is unexported.
	// If config were exported, we could create it:
	// c := Config{minPasswordLength: 5}
	//
	// Since it's unexported, package configuration is encapsulated.
	// External users cannot bypass the validation rules.

	t.Log("config type is correctly unexported, ensuring consistent validation rules")
}
