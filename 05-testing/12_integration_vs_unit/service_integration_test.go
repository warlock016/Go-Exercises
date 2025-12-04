//go:build integration

package service

import "testing"

// TODO(human): Write INTEGRATION test
// This would test with a real database
// For this exercise, just demonstrate the pattern with MockDB
// but in real code, this would use a real database connection
//
// func TestServiceIntegration(t *testing.T) {
//     // In real code: connect to test database
//     // For demo, we'll use MockDB
//     db := &MockDB{Data: make(map[string]string)}
//     service := &Service{DB: db}
//
//     // Test full workflow
//     err := db.Set("user:123", "Alice")
//     if err != nil {
//         t.Fatal(err)
//     }
//
//     user, err := service.GetUser("123")
//     // ... assertions
// }

// Run integration tests: go test -tags=integration
// Run all tests: go test -tags=integration ./...
// Run only unit tests: go test (default)
