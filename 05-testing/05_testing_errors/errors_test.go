package errors

import (
	"strings"
	"testing"
)

// TODO(human): Create helper function for error checking
// func assertErrorContains(t *testing.T, err error, substr string) {
//     t.Helper()
//     if err == nil {
//         t.Fatal("expected error, got nil")
//     }
//     if !strings.Contains(err.Error(), substr) {
//         t.Errorf("error %q does not contain %q", err.Error(), substr)
//     }
// }

// TODO(human): Test ParseInt function
// Use table-driven test with subtests
// Test cases:
//   - Valid integers: "123", "-456", "0"
//   - Invalid: "abc", "12.34", "", "999999999999999999999"
// Check both the result AND error
// For error cases, check error message contains meaningful text

// TODO(human): Test Divide function
// Test cases:
//   - Normal division: 10 / 2 = 5
//   - Division by zero: expect error with "division by zero"
//   - Negative numbers: -10 / 2 = -5
//   - Division resulting in float: 5 / 2 = 2.5

// TODO(human): Test ReadFile function
// Test cases:
//   - Empty path: should error with "empty path"
//   - Non-existent file: should error (error message will contain os error)
//   - For this exercise, just test error cases (no need to create actual files)
// Hint: You can check that error contains "failed to read file" for wrapped errors
