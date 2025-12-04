package math

import "testing"

// TODO(human): Create a table-driven test with subtests for Abs
// Use t.Run(tt.name, func(t *testing.T) { ... }) to wrap each test case
// Structure:
//   for _, tt := range tests {
//       t.Run(tt.name, func(t *testing.T) {
//           got := Abs(tt.input)
//           if got != tt.want {
//               t.Errorf("Abs(%d) = %d, want %d", tt.input, got, tt.want)
//           }
//       })
//   }
//
// Test cases: positive, negative, zero, large numbers

// TODO(human): Create a table-driven test with subtests for Max
// Test cases: a > b, b > a, a == b, negative numbers, zero

// TODO(human): Create a table-driven test with subtests for Clamp
// Test cases:
// - Value within range (should return value)
// - Value below min (should return min)
// - Value above max (should return max)
// - Value equals min/max
// - Min equals max (edge case)
//
// After writing tests, experiment with:
//   go test -v
//   go test -v -run TestClamp/below
//   go test -v -run TestMax/greater
