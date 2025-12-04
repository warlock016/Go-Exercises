package coverage

import "testing"

// TODO(human): Write tests for Classify
// Test ALL branches to achieve 100% coverage:
// - negative numbers
// - zero
// - small positive (1-9)
// - medium positive (10-99)
// - large positive (100+)

// TODO(human): Write tests for ProcessData
// Test ALL branches:
// - Empty slice (error case)
// - Negative value (error case)
// - Sum > 1000 (error case)
// - Valid data (success case)

// TODO(human): Write tests for Calculate
// Test ALL operations and error cases:
// - add
// - sub
// - mul (normal and too large case)
// - div (normal and divide by zero)
// - unknown operation

// After writing initial tests:
// 1. Run: go test -cover
//    Observe coverage percentage
//
// 2. Generate profile: go test -coverprofile=coverage.out
//
// 3. View in browser: go tool cover -html=coverage.out
//    Red lines = not covered
//    Green lines = covered
//
// 4. Add tests for uncovered lines
//
// 5. Run go test -cover again
//    Aim for >90% coverage
//
// 6. View by function: go tool cover -func=coverage.out
