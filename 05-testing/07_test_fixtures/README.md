# Exercise 07: Test Fixtures

**Learning Goal:** Use testdata/ directory for test fixtures and implement golden file testing pattern.

---

## Problem Description

Many tests need input files or expected output files. Go convention is to use a `testdata/` directory which is ignored by `go build` but accessible to tests. Golden file testing compares actual output to expected output stored in files.

---

## Functions to Test

```go
func ProcessMarkdown(input string) string
func GenerateReport(data []Record) string
```

---

## The testdata/ Convention

```
07_test_fixtures/
├── fixtures.go
├── fixtures_test.go
└── testdata/
    ├── input1.md
    ├── expected1.html
    ├── input2.md
    └── expected2.html
```

**Key facts:**
- `go build` ignores testdata/ directories
- Tests can read files from testdata/
- Use relative path: `"testdata/file.txt"`
- Store both inputs and expected outputs

---

## Golden File Pattern

```go
func TestProcessMarkdown(t *testing.T) {
    input, err := os.ReadFile("testdata/input.md")
    if err != nil {
        t.Fatal(err)
    }

    got := ProcessMarkdown(string(input))

    golden, err := os.ReadFile("testdata/expected.html")
    if err != nil {
        t.Fatal(err)
    }

    if got != string(golden) {
        t.Errorf("output mismatch\ngot:\n%s\nwant:\n%s", got, string(golden))
    }
}
```

---

## Your Task

1. Create testdata/ directory with test files
2. Implement tests using golden files
3. Compare actual output to expected output

---

## Instructions

The testdata/ directory and files are already created. Write tests that:
1. Read input from testdata/
2. Process it
3. Compare to expected output in testdata/

---

**Next Exercise:** `08_http_handler_tests` - Testing HTTP handlers
