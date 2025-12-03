# CSV Data Processor - Project Specification

A command-line tool for reading, transforming, and outputting CSV data.

**Estimated Time:** 4-6 hours
**Difficulty:** Intermediate
**Prerequisites:** Project 0 (Weather CLI) completed, basic file I/O understanding

---

## Learning Goals

This project practices:
- **File I/O** - Opening, reading, and writing files (NEW)
- **CSV parsing** - Using `encoding/csv` package (NEW)
- **Interfaces** - First exposure to `io.Reader` and `io.Writer` (NEW)
- **Error handling** - Handling malformed data gracefully
- **Data transformations** - Filter, aggregate, compute operations
- **Testing** - Writing tests for data transformations

---

## Requirements

### Functional Requirements

1. **Read CSV files** from disk or stdin
2. **Apply transformations** via command-line flags:
   - Filter rows by column value
   - Select specific columns
   - Sort by column
   - Compute aggregates (sum, avg, count, min, max)
3. **Output results** in CSV or JSON format
4. **Handle errors** gracefully:
   - Missing files
   - Malformed CSV rows
   - Invalid column names

### Command-Line Interface

```bash
# Basic: read and output CSV
csvproc input.csv

# Filter rows where column "status" equals "active"
csvproc -filter "status=active" input.csv

# Select only specific columns
csvproc -columns "name,email,age" input.csv

# Sort by column (ascending)
csvproc -sort "age" input.csv

# Aggregate: compute sum of "amount" column
csvproc -aggregate "sum:amount" input.csv

# Output as JSON
csvproc -format json input.csv

# Read from stdin
cat input.csv | csvproc -filter "status=active"

# Combine operations
csvproc -filter "status=active" -columns "name,amount" -sort "amount" -format json input.csv
```

### Non-Functional Requirements

- Code organized into multiple files
- Transformation functions should be testable (no direct file I/O in transform logic)
- Support both file input and stdin
- Meaningful error messages with line numbers for malformed data

---

## Input/Output Examples

### Input CSV (sales.csv)
```csv
name,amount,status,date
Alice,150.00,active,2025-01-15
Bob,75.50,inactive,2025-01-16
Carol,200.00,active,2025-01-17
Dave,50.00,active,2025-01-18
```

### Filter Example
```bash
$ csvproc -filter "status=active" sales.csv
name,amount,status,date
Alice,150.00,active,2025-01-15
Carol,200.00,active,2025-01-17
Dave,50.00,active,2025-01-18
```

### Columns Example
```bash
$ csvproc -columns "name,amount" sales.csv
name,amount
Alice,150.00
Bob,75.50
Carol,200.00
Dave,50.00
```

### Aggregate Example
```bash
$ csvproc -aggregate "sum:amount" sales.csv
sum_amount
475.50

$ csvproc -aggregate "count:name" -filter "status=active" sales.csv
count_name
3
```

### JSON Output Example
```bash
$ csvproc -format json -columns "name,amount" sales.csv
[
  {"name": "Alice", "amount": "150.00"},
  {"name": "Bob", "amount": "75.50"},
  {"name": "Carol", "amount": "200.00"},
  {"name": "Dave", "amount": "50.00"}
]
```

### Error Examples
```bash
$ csvproc nonexistent.csv
Error: cannot open file: nonexistent.csv

$ csvproc -filter "invalid" sales.csv
Error: invalid filter format "invalid", expected "column=value"

$ csvproc -columns "name,invalid_col" sales.csv
Error: column "invalid_col" not found in header
```

---

## Suggested File Structure

```
01_csv_processor/
├── main.go              # Entry point: flag parsing, orchestration
├── reader.go            # CSV reading: ReadCSV function
├── transformer.go       # Data transformations: Filter, Select, Sort, Aggregate
├── writer.go            # Output: WriteCSV, WriteJSON functions
├── transformer_test.go  # Tests for transformation functions
├── testdata/            # Sample CSV files for testing
│   ├── simple.csv
│   ├── sales.csv
│   └── malformed.csv
└── README.md            # This spec
```

---

## Useful Packages

| Package | Purpose | Documentation |
|---------|---------|---------------|
| `encoding/csv` | Read and write CSV files | https://pkg.go.dev/encoding/csv |
| `encoding/json` | JSON output | https://pkg.go.dev/encoding/json |
| `os` | File operations, stdin | https://pkg.go.dev/os |
| `io` | Reader/Writer interfaces | https://pkg.go.dev/io |
| `flag` | Command-line flags | https://pkg.go.dev/flag |
| `sort` | Sorting slices | https://pkg.go.dev/sort |
| `strconv` | String to number conversion | https://pkg.go.dev/strconv |
| `strings` | String manipulation | https://pkg.go.dev/strings |

---

## Key Concepts to Learn

### 1. Opening Files

```go
file, err := os.Open("filename.csv")
if err != nil {
    return err
}
defer file.Close()  // Always close files!
```

### 2. CSV Reader

```go
reader := csv.NewReader(file)
// Read all at once
records, err := reader.ReadAll()

// Or read line by line
for {
    record, err := reader.Read()
    if err == io.EOF {
        break
    }
    if err != nil {
        // Handle malformed row
    }
    // Process record ([]string)
}
```

### 3. Reading from stdin

```go
// os.Stdin implements io.Reader
reader := csv.NewReader(os.Stdin)
```

### 4. The io.Reader Interface

Both files and stdin implement `io.Reader`. Your CSV reading function can accept `io.Reader` to work with both:

```go
func ReadCSV(r io.Reader) ([][]string, error) {
    reader := csv.NewReader(r)
    return reader.ReadAll()
}
```

---

## Hints (Read Only If Stuck)

<details>
<summary>Hint 1: Data representation</summary>

Consider representing CSV data as:
```go
type Table struct {
    Headers []string
    Rows    [][]string
}
```
This makes transformations cleaner.
</details>

<details>
<summary>Hint 2: Finding column index</summary>

```go
func findColumnIndex(headers []string, name string) int {
    for i, h := range headers {
        if h == name {
            return i
        }
    }
    return -1  // Not found
}
```
</details>

<details>
<summary>Hint 3: Detecting stdin vs file</summary>

```go
// Check if we have a filename argument or should read stdin
if flag.NArg() == 0 {
    // No filename, read from stdin
    reader = os.Stdin
} else {
    // Open the file
    file, err := os.Open(flag.Arg(0))
    // ...
    reader = file
}
```
</details>

<details>
<summary>Hint 4: Sorting with custom comparator</summary>

```go
import "sort"

sort.Slice(rows, func(i, j int) bool {
    return rows[i][colIndex] < rows[j][colIndex]
})
```
</details>

<details>
<summary>Hint 5: Testing transformations</summary>

Test transformation functions with in-memory data, not files:
```go
func TestFilter(t *testing.T) {
    input := Table{
        Headers: []string{"name", "status"},
        Rows: [][]string{
            {"Alice", "active"},
            {"Bob", "inactive"},
        },
    }
    result := Filter(input, "status", "active")
    // Assert result has only Alice
}
```
</details>

---

## Success Criteria

- [ ] `csvproc testdata/simple.csv` reads and outputs CSV
- [ ] `csvproc -filter "col=val" file.csv` filters correctly
- [ ] `csvproc -columns "a,b,c" file.csv` selects columns
- [ ] `csvproc -sort "col" file.csv` sorts by column
- [ ] `csvproc -aggregate "sum:col" file.csv` computes aggregates
- [ ] `csvproc -format json file.csv` outputs valid JSON
- [ ] `cat file.csv | csvproc` reads from stdin
- [ ] Malformed CSV produces helpful error with line number
- [ ] Missing file produces clear error message
- [ ] Tests exist for transformation functions
- [ ] Code is split across multiple files

---

## Stretch Goals (Optional)

- [ ] Support numeric sorting (not just string comparison)
- [ ] Support multiple filters (`-filter "a=1" -filter "b=2"`)
- [ ] Support descending sort (`-sort "-col"` or `-sort "col:desc"`)
- [ ] Add `-limit N` flag to output only first N rows
- [ ] Add `-skip N` flag to skip first N rows
- [ ] Support TSV (tab-separated) input with `-delimiter` flag

---

## Getting Started

1. Create `testdata/` folder with sample CSV files
2. Start with `main.go` - parse flags and determine input source
3. Implement `reader.go` - get CSV reading working first
4. Add `writer.go` - output what you read (round-trip test)
5. Implement transformations one at a time in `transformer.go`
6. Write tests as you go in `transformer_test.go`

Ask for help if you're stuck on a specific concept!
