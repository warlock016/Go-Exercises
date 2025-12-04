package fixtures

import (
	"os"
	"testing"
)

// TODO(human): Test ProcessMarkdown using testdata/input.md and testdata/expected.html
// 1. Read testdata/input.md
// 2. Call ProcessMarkdown
// 3. Read testdata/expected.html
// 4. Compare the results
//
// func TestProcessMarkdown(t *testing.T) {
//     input, err := os.ReadFile("testdata/input.md")
//     if err != nil {
//         t.Fatalf("failed to read input: %v", err)
//     }
//
//     got := ProcessMarkdown(string(input))
//
//     golden, err := os.ReadFile("testdata/expected.html")
//     if err != nil {
//         t.Fatalf("failed to read golden file: %v", err)
//     }
//
//     if got != string(golden) {
//         t.Errorf("output mismatch\ngot:\n%s\nwant:\n%s", got, string(golden))
//     }
// }

// TODO(human): Test GenerateReport using golden file pattern
// 1. Create test data: []Record with a few records
// 2. Call GenerateReport
// 3. Read testdata/expected_report.txt
// 4. Compare output
//
// Or: You can manually write expected output in the test itself for this one
// since the data is programmatically generated

// Note: You'll need to create the testdata/ directory and files!
// Create: testdata/input.md with some markdown
// Create: testdata/expected.html with expected output
// Create: testdata/expected_report.txt with expected report format
