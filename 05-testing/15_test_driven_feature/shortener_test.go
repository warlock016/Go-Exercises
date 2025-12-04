package shortener

import "testing"

// TODO(human): Follow TDD workflow for each step
//
// STEP 1: Shorten a URL
// RED: Write test for Shorten("https://example.com")
//      Expect: non-empty code, no error
//      Run: go test (should FAIL)
// GREEN: Create URLShortener with Shorten method
//        Return hardcoded "abc123"
//        Run: go test (should PASS)
// REFACTOR: Nothing to refactor yet
//
// STEP 2: Expand a URL
// RED: Test Expand(code) returns original URL
//      Run: go test (should FAIL)
// GREEN: Store URLs in map[string]string
//        Generate random code, store URL
//        Expand returns from map
//        Run: go test (should PASS)
// REFACTOR: Maybe extract map initialization
//
// STEP 3: Handle missing codes
// RED: Test Expand("nonexistent") returns error
//      Run: go test (should FAIL)
// GREEN: Check if code exists in map
//        Return error if not found
//        Run: go test (should PASS)
// REFACTOR: Define error constants
//
// STEP 4: Validate URLs
// RED: Test Shorten("") returns error
//      Test Shorten("not-a-url") returns error
//      Run: go test (should FAIL)
// GREEN: Add URL validation in Shorten
//        Return error if empty or invalid
//        Run: go test (should PASS)
// REFACTOR: Extract validateURL helper
//
// STEP 5: Track Statistics
// RED: Test Stats(code) returns access count
//      Initially 0, increment on each Expand
//      Run: go test (should FAIL)
// GREEN: Add accessCounts map[string]int
//        Increment on Expand
//        Stats returns count
//        Run: go test (should PASS)
// REFACTOR: Create Stats struct
//
// STEP 6: Handle duplicate URLs
// RED: Test shortening same URL twice returns same code
//      Run: go test (should FAIL)
// GREEN: Add reverse map: map[string]string (url -> code)
//        Check before generating new code
//        Run: go test (should PASS)
// REFACTOR: Clean up map handling
//
// Remember:
// - Write ONE test at a time
// - See it FAIL
// - Make it PASS with minimal code
// - Refactor
// - Repeat!
