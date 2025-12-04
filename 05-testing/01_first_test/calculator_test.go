package calculator

import "testing"

// TODO(human): Write tests for the Add function
// Test at least 3 cases: positive numbers, negative numbers, and zero
// Example:
//   func TestAdd(t *testing.T) {
//       result := Add(2, 3)
//       if result != 5 {
//           t.Errorf("Add(2, 3) = %d, want 5", result)
//       }
//   }

// TODO(human): Write tests for the Subtract function
// Test various cases including results that are negative

// TODO(human): Write tests for the Multiply function
// Include test cases for zero, positive, and negative numbers

// TODO(human): Write tests for the Divide function
// Test normal division (when b != 0)
// Remember: Divide returns (int, error), so check both values
// Use t.Fatal() if error is unexpected (prevents nil pointer access)

// TODO(human): Write a separate test for Divide by zero
// This should test that an error IS returned when b == 0
// Example:
//   func TestDivideByZero(t *testing.T) {
//       _, err := Divide(10, 0)
//       if err == nil {
//           t.Error("Divide(10, 0) expected error, got nil")
//       }
//   }

// TODO(human): Write tests for the IsEven function
// Test even numbers, odd numbers, zero, and negative numbers
// You can test multiple cases in one test function
