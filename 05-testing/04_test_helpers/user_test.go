package user

import "testing"

func setupUser(t *testing.T, name, email string) *User {
	t.Helper()
	user, err := NewUser(name, email)
	assertNoError(t, err)
	return user
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func assertEqual(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestNewUser(t *testing.T) {
	tests := []struct {
		Name    string
		User    string
		Email   string
		wantErr bool
	}{
		{"valid", "Alice", "alice@gmail.com", false},
		{"empty name", "", "alice@gmail.com", true},
		{"invalid email", "Charlie", "charlieyahoo.co.uk", true},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			user, err := NewUser(tt.User, tt.Email)
			if tt.wantErr {
				assertError(t, err)
			} else {
				assertNoError(t, err)
				assertEqual(t, user.Name, tt.User)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		Name    string
		User    string
		Email   string
		wantErr bool
	}{
		{
			Name:    "valid user",
			User:    "Alice",
			Email:   "alice@gmail.com",
			wantErr: false,
		},
		{
			Name:    "empty name",
			User:    "",
			Email:   "bob@gmail.com",
			wantErr: true,
		},
		{
			Name:    "invalid email",
			User:    "Charlie",
			Email:   "charlieoutlook.com",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var user *User
			if tt.wantErr {
				user = &User{
					Name:  tt.User,
					Email: tt.Email,
				}
				err := user.Validate()
				assertError(t, err)
			} else {
				user = setupUser(t, tt.User, tt.Email)
				err := user.Validate()
				assertNoError(t, err)
				assertEqual(t, user.Name, tt.User)
				assertEqual(t, user.Email, tt.Email)
			}
		})
	}
}

func TestUpdateEmail(t *testing.T) {
	tests := []struct {
		Name     string
		User     string
		OldEmail string
		Email    string
		wantErr  bool
	}{
		{
			Name:     "valid email",
			User:     "Alice",
			OldEmail: "alice@gmail.com",
			Email:    "alice@gmail.de",
			wantErr:  false,
		},
		{
			Name:     "invalid email",
			User:     "Bob",
			OldEmail: "bob@gmx.de",
			Email:    "bobbygreengmx.de",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			user := setupUser(t, tt.User, tt.OldEmail)
			err := user.UpdateEmail(tt.Email)
			if tt.wantErr {
				assertError(t, err)
				assertEqual(t, user.Email, tt.OldEmail)
			} else {
				assertNoError(t, err)
				assertEqual(t, user.Email, tt.Email)
			}
		})
	}
}

// TODO(human): Create helper functions at the top of this file
//
// Helper 1: assertNoError(t *testing.T, err error)
//   - Call t.Helper() first
//   - If err != nil, call t.Fatalf("unexpected error: %v", err)
//
// Helper 2: assertError(t *testing.T, err error)
//   - Call t.Helper() first
//   - If err == nil, call t.Fatal("expected error, got nil")
//
// Helper 3: assertEqual(t *testing.T, got, want interface{})
//   - Call t.Helper() first
//   - If got != want, call t.Errorf("got %v, want %v", got, want)
//
// Helper 4: setupUser(t *testing.T, name, email string) *User
//   - Call t.Helper() first
//   - Call NewUser(name, email)
//   - Use assertNoError to check the error
//   - Return the user

// TODO(human): Test NewUser with subtests
// Test cases:
// - "valid user" - create user with valid name and email
// - "empty name" - expect error
// - "invalid email" - expect error (no @ symbol)
//
// Use your helper functions!

// TODO(human): Test User.Validate() with subtests
// Test cases:
// - "valid user" - setupUser then validate, should pass
// - You can manually create invalid users to test error cases

// TODO(human): Test User.UpdateEmail() with subtests
// Test cases:
// - "valid email" - setupUser, update email, check it changed
// - "invalid email" - setupUser, try invalid email, should error and email unchanged
//
// After writing tests:
// 1. Run go test -v to see all tests pass
// 2. Intentionally remove t.Helper() from a helper
// 3. Break a test to see where the error points
// 4. Add t.Helper() back and see the difference
