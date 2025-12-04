package user

import "testing"

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
