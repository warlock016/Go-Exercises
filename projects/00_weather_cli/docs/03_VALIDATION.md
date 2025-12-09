# Module: Validation

## Purpose

Validate all CLI input parameters before making any API calls. This module implements the aggregated validation pattern - collecting ALL validation errors and reporting them together, rather than stopping at the first error.

---

## Package Location

```
validation/validation.go
```

---

## Dependencies

| Package | Purpose |
|---------|---------|
| `errors` (internal) | `ValidationErrors`, `FieldError` types |
| `types` (internal) | `CLIInput` struct |

---

## Public API

### Types

```go
// CLIInput represents the raw input from CLI flags
type CLIInput struct {
    Latitude  string // Raw latitude string from flag
    Longitude string // Raw longitude string from flag
    StartDate string // Raw start date string (YYYY-MM-DD)
    EndDate   string // Raw end date string (YYYY-MM-DD)
    Format    string // Output format: "table" or "json"
}

// ValidatedInput represents validated and parsed input
type ValidatedInput struct {
    Latitude  float64   // Parsed latitude
    Longitude float64   // Parsed longitude
    StartDate time.Time // Parsed start date
    EndDate   time.Time // Parsed end date
    Format    string    // Validated format
}
```

### Functions

```go
// Validate checks all input fields and returns aggregated errors
// Returns ValidatedInput if all validations pass, or *ValidationErrors if any fail
func Validate(input CLIInput) (*ValidatedInput, error)

// Individual validators (internal, but documented for clarity)
func validateLatitude(lat string) (float64, *errors.FieldError)
func validateLongitude(long string) (float64, *errors.FieldError)
func validateDate(date, fieldName string) (time.Time, *errors.FieldError)
func validateDateRange(start, end time.Time) *errors.FieldError
func validateFormat(format string) *errors.FieldError
```

---

## Validation Rules

### Latitude

| Rule | Constraint | Error Message |
|------|------------|---------------|
| Required | Must not be empty | "latitude is required" |
| Numeric | Must parse as float64 | "latitude must be a valid number" |
| Range | Must be between -90 and 90 (inclusive) | "latitude must be between -90 and 90" |

### Longitude

| Rule | Constraint | Error Message |
|------|------------|---------------|
| Required | Must not be empty | "longitude is required" |
| Numeric | Must parse as float64 | "longitude must be a valid number" |
| Range | Must be between -180 and 180 (inclusive) | "longitude must be between -180 and 180" |

### StartDate

| Rule | Constraint | Error Message |
|------|------------|---------------|
| Required | Must not be empty | "start_date is required" |
| Format | Must be YYYY-MM-DD | "start_date must be in YYYY-MM-DD format" |
| Valid | Must be a real date (no Feb 30) | "start_date is not a valid date" |

### EndDate

| Rule | Constraint | Error Message |
|------|------------|---------------|
| Required | Must not be empty | "end_date is required" |
| Format | Must be YYYY-MM-DD | "end_date must be in YYYY-MM-DD format" |
| Valid | Must be a real date | "end_date is not a valid date" |
| Order | Must be >= StartDate | "end_date must be on or after start_date" |

### Format

| Rule | Constraint | Error Message |
|------|------------|---------------|
| Required | Must not be empty | "format is required" |
| Allowed | Must be "table" or "json" (case-insensitive) | "format must be 'table' or 'json'" |

---

## Behavior Specification

### Validate(input CLIInput)

**Input:** `CLIInput` with raw string values from CLI flags

**Output:** `(*ValidatedInput, error)`
- On success: `(*ValidatedInput, nil)`
- On failure: `(nil, *ValidationErrors)`

**Behavior:**
1. Create empty `ValidationErrors` struct
2. Validate latitude → if error, add to ValidationErrors
3. Validate longitude → if error, add to ValidationErrors
4. Validate start_date → if error, add to ValidationErrors
5. Validate end_date → if error, add to ValidationErrors
6. If both dates valid, validate date range → if error, add to ValidationErrors
7. Validate format → if error, add to ValidationErrors
8. If `ValidationErrors.HasErrors()` → return `(nil, validationErrors)`
9. Return `(validatedInput, nil)`

**Critical Pattern:**
```go
func Validate(input CLIInput) (*ValidatedInput, error) {
    errs := &errors.ValidationErrors{}
    var result ValidatedInput

    // Validate each field, collecting ALL errors
    lat, latErr := validateLatitude(input.Latitude)
    if latErr != nil {
        errs.Add(latErr.Field, latErr.Message)
    } else {
        result.Latitude = lat
    }

    // ... more validations ...

    // Only return error if there ARE errors
    if errs.HasErrors() {
        return nil, errs
    }
    return &result, nil
}
```

---

## Edge Cases

| Input | Expected Behavior |
|-------|-------------------|
| `lat: ""` | Error: "latitude is required" |
| `lat: "abc"` | Error: "latitude must be a valid number" |
| `lat: "91"` | Error: "latitude must be between -90 and 90" |
| `lat: "-90"` | Valid (boundary) |
| `lat: "90"` | Valid (boundary) |
| `lat: "52.5"` | Valid |
| `lat: "  52.5  "` | Valid (trim whitespace) |
| `long: "181"` | Error: longitude out of range |
| `date: "2024-13-01"` | Error: invalid date (month 13) |
| `date: "2024-02-30"` | Error: invalid date (Feb 30) |
| `date: "01-01-2024"` | Error: wrong format (DD-MM-YYYY) |
| `start: "2024-01-31", end: "2024-01-01"` | Error: end before start |
| `start: "2024-01-15", end: "2024-01-15"` | Valid (same day) |
| `format: "TABLE"` | Valid (case-insensitive) |
| `format: "xml"` | Error: invalid format |

---

## Example Error Output

**Single error:**
```
validation error: latitude: must be between -90 and 90
```

**Multiple errors (the goal of aggregation):**
```
validation errors:
  - latitude: must be a valid number
  - start_date: must be in YYYY-MM-DD format
  - format: must be 'table' or 'json'
```

---

## Useful Packages

| Package | Why Useful |
|---------|------------|
| `strconv` | `strconv.ParseFloat()` for lat/long parsing |
| `time` | `time.Parse()` with layout `"2006-01-02"` for date parsing |
| `strings` | `strings.TrimSpace()`, `strings.ToLower()` |

### Date Parsing in Go

Go uses a reference date for parsing: `Mon Jan 2 15:04:05 MST 2006`

For `YYYY-MM-DD`, the layout is: `"2006-01-02"`

```go
t, err := time.Parse("2006-01-02", "2024-03-15")
```

---

## Test Cases

| Scenario | Input | Expected |
|----------|-------|----------|
| All valid | `{Lat: "52.5", Long: "13.4", Start: "2024-01-01", End: "2024-01-31", Format: "table"}` | ValidatedInput, nil |
| Empty latitude | `{Lat: ""}` | ValidationErrors with "latitude is required" |
| Invalid latitude | `{Lat: "abc"}` | ValidationErrors with "must be a valid number" |
| Latitude out of range | `{Lat: "100"}` | ValidationErrors with "must be between -90 and 90" |
| Multiple errors | `{Lat: "abc", Long: "xyz", Start: "bad", End: "bad", Format: "xml"}` | ValidationErrors with 5 errors |
| Boundary latitude | `{Lat: "90"}` | Valid |
| Negative longitude | `{Long: "-122.4"}` | Valid |
| End before start | `{Start: "2024-12-31", End: "2024-01-01"}` | ValidationErrors with date range error |
| Same start/end | `{Start: "2024-01-01", End: "2024-01-01"}` | Valid (same day is allowed) |
| Case insensitive format | `{Format: "JSON"}` | Valid, normalized to "json" |
| Whitespace in values | `{Lat: "  52.5  "}` | Valid after trimming |

---

## Implementation Hints

1. **Aggregation pattern:** Always validate ALL fields. Don't return early on first error:
   ```go
   // WRONG - stops at first error
   if latErr != nil {
       return nil, latErr
   }

   // RIGHT - collects all errors
   if latErr != nil {
       errs.Add("latitude", latErr.Message)
   }
   ```

2. **Dependent validation:** Date range validation depends on both dates being valid. Only run it if both parsed successfully:
   ```go
   var startDate, endDate time.Time
   var startOK, endOK bool

   if start, err := validateDate(...); err != nil {
       errs.Add(...)
   } else {
       startDate, startOK = start, true
   }
   // Similar for end date

   if startOK && endOK {
       if rangeErr := validateDateRange(startDate, endDate); rangeErr != nil {
           errs.Add(...)
       }
   }
   ```

3. **Normalize format:** Convert to lowercase for comparison and storage:
   ```go
   format := strings.ToLower(strings.TrimSpace(input.Format))
   if format != "table" && format != "json" {
       // error
   }
   result.Format = format // Store normalized version
   ```

4. **Trim all inputs:** Whitespace-only inputs should be treated as empty:
   ```go
   lat := strings.TrimSpace(input.Latitude)
   if lat == "" {
       return 0, &FieldError{Field: "latitude", Message: "latitude is required"}
   }
   ```

5. **Testing aggregation:** Write tests that verify ALL errors are returned, not just the first:
   ```go
   func TestValidate_MultipleErrors(t *testing.T) {
       input := CLIInput{Latitude: "abc", Longitude: "xyz", ...}
       _, err := Validate(input)

       var verrs *errors.ValidationErrors
       if !errors.As(err, &verrs) {
           t.Fatal("expected ValidationErrors")
       }

       if len(verrs.Errors) < 2 {
           t.Error("expected multiple errors to be aggregated")
       }
   }
   ```
