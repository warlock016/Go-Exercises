# Module: Errors

## Purpose

Define all custom error types and sentinel errors for the weather CLI. This module provides a consistent error vocabulary that all other packages use, enabling proper error classification, inspection, and user-friendly messaging.

---

## Package Location

```
errors/errors.go
```

---

## Dependencies

**None** - This is a leaf package with no internal dependencies.

---

## Public API

### Sentinel Errors

```go
var (
    // Network/Transport errors
    ErrTimeout     error  // Request timed out
    ErrNetwork     error  // General network failure (connection refused, DNS, etc.)

    // API-specific errors
    ErrNotFound    error  // Resource not found (HTTP 404)
    ErrRateLimited error  // Too many requests (HTTP 429)
    ErrUnauthorized error // Invalid/missing API key (HTTP 401)

    // Input errors
    ErrInvalidInput error // Generic invalid input (fallback)
)
```

### Custom Error Types

#### WeatherAPIError

```go
// WeatherAPIError represents an error from the Open-Meteo API
type WeatherAPIError struct {
    StatusCode int    // HTTP status code
    Message    string // Error message from API or description
    Endpoint   string // The API endpoint that failed
    Latitude   float64
    Longitude  float64
}

// Error implements the error interface
func (e *WeatherAPIError) Error() string

// Unwrap returns the underlying sentinel error based on status code
func (e *WeatherAPIError) Unwrap() error
```

#### GeocodeError

```go
// GeocodeError represents an error from the Nominatim geocoding API
type GeocodeError struct {
    StatusCode int    // HTTP status code
    Message    string // Error message
    Query      string // The location query that failed
}

// Error implements the error interface
func (e *GeocodeError) Error() string

// Unwrap returns the underlying sentinel error based on status code
func (e *GeocodeError) Unwrap() error
```

#### ValidationErrors

```go
// FieldError represents a single validation error
type FieldError struct {
    Field   string // Field name (e.g., "latitude", "start_date")
    Message string // Human-readable error message
}

// ValidationErrors aggregates multiple validation errors
type ValidationErrors struct {
    Errors []FieldError
}

// Error implements the error interface
func (e *ValidationErrors) Error() string

// Add appends a new field error
func (e *ValidationErrors) Add(field, message string)

// HasErrors returns true if there are any errors
func (e *ValidationErrors) HasErrors() bool
```

---

## Behavior Specification

### WeatherAPIError.Error()

**Input:** None (method on struct)

**Output:** Formatted error string

**Behavior:**
1. Return a string in the format: `"weather API error (HTTP {StatusCode}): {Message}"`
2. Include endpoint if available: `"weather API error at {Endpoint} (HTTP {StatusCode}): {Message}"`

**Example Output:**
```
weather API error (HTTP 404): Location not found for coordinates 999.0, 999.0
weather API error at /v1/forecast (HTTP 500): Internal server error
```

---

### WeatherAPIError.Unwrap()

**Input:** None (method on struct)

**Output:** Underlying sentinel error

**Behavior:**
1. Map HTTP status codes to sentinel errors:
   - 400 → `ErrInvalidInput`
   - 401 → `ErrUnauthorized`
   - 404 → `ErrNotFound`
   - 429 → `ErrRateLimited`
   - 5xx → `nil` (no sentinel, but retryable)
2. Return `nil` for unmapped status codes

**Why Unwrap?**
Enables `errors.Is(err, ErrNotFound)` to work even when the actual error is `*WeatherAPIError`.

---

### GeocodeError.Error()

**Input:** None (method on struct)

**Output:** Formatted error string

**Behavior:**
1. Return: `"geocoding error for '{Query}' (HTTP {StatusCode}): {Message}"`

**Example Output:**
```
geocoding error for 'Atlantis' (HTTP 404): Location not found
geocoding error for 'Berlin' (HTTP 429): Rate limit exceeded
```

---

### GeocodeError.Unwrap()

**Input:** None (method on struct)

**Output:** Underlying sentinel error

**Behavior:**
Same mapping as WeatherAPIError.Unwrap().

---

### ValidationErrors.Error()

**Input:** None (method on struct)

**Output:** Formatted string listing all validation errors

**Behavior:**
1. If single error: `"validation error: {field}: {message}"`
2. If multiple errors:
   ```
   validation errors:
     - {field1}: {message1}
     - {field2}: {message2}
   ```

**Example Output:**
```
validation error: latitude: must be between -90 and 90

validation errors:
  - latitude: must be between -90 and 90
  - start_date: must be in YYYY-MM-DD format
  - end_date: must be after start_date
```

---

### ValidationErrors.Add()

**Input:**
- `field string` - Name of the field with error
- `message string` - Human-readable error description

**Output:** None (modifies struct)

**Behavior:**
1. Append a new `FieldError{Field: field, Message: message}` to the Errors slice

---

### ValidationErrors.HasErrors()

**Input:** None

**Output:** `bool`

**Behavior:**
1. Return `len(e.Errors) > 0`

---

## Error Classification Matrix

| HTTP Status | Sentinel Error   | Retryable? | User Message                          |
|-------------|------------------|------------|---------------------------------------|
| 400         | ErrInvalidInput  | No         | "Invalid request parameters"          |
| 401         | ErrUnauthorized  | No         | "Invalid or missing API key"          |
| 404         | ErrNotFound      | No         | "Location not found"                  |
| 429         | ErrRateLimited   | Yes        | "Too many requests, try again later"  |
| 500         | (none)           | Yes        | "Server error, retrying..."           |
| 502         | (none)           | Yes        | "Service temporarily unavailable"     |
| 503         | (none)           | Yes        | "Service temporarily unavailable"     |
| 504         | ErrTimeout       | Yes        | "Request timed out, retrying..."      |
| Timeout     | ErrTimeout       | Yes        | "Request timed out"                   |
| Network     | ErrNetwork       | Yes        | "Network error, check connection"     |

---

## Useful Packages

| Package | Why Useful |
|---------|------------|
| `errors` | `errors.New()` for sentinels, `errors.Is()` and `errors.As()` for inspection |
| `fmt` | `fmt.Sprintf()` for formatting error messages |
| `strings` | `strings.Builder` for efficient multi-error formatting |

---

## Test Cases

| Scenario | Input | Expected Output |
|----------|-------|-----------------|
| WeatherAPIError.Error() | `{StatusCode: 404, Message: "not found", Latitude: 52.5, Longitude: 13.4}` | Contains "404", "not found" |
| WeatherAPIError.Unwrap() with 404 | `{StatusCode: 404}` | `ErrNotFound` |
| WeatherAPIError.Unwrap() with 500 | `{StatusCode: 500}` | `nil` |
| GeocodeError.Error() | `{Query: "Berlin", StatusCode: 429, Message: "rate limited"}` | Contains "Berlin", "429" |
| ValidationErrors single | Add("lat", "invalid") | `"validation error: lat: invalid"` |
| ValidationErrors multiple | Add 3 errors | Multi-line output with all 3 |
| ValidationErrors.HasErrors() empty | New struct | `false` |
| ValidationErrors.HasErrors() with errors | After Add() | `true` |
| errors.Is(weatherErr, ErrNotFound) | 404 error | `true` |
| errors.Is(weatherErr, ErrTimeout) | 404 error | `false` |

---

## Implementation Hints

1. **Sentinel errors:** Use `errors.New("descriptive message")` at package level as `var`.

2. **Unwrap pattern:** The `Unwrap()` method is what makes `errors.Is()` work through the chain. Without it, `errors.Is(weatherAPIError, ErrNotFound)` would return `false`.

3. **ValidationErrors as pointer:** When using `errors.As()`, the target must be a pointer. Design your validation flow so `ValidationErrors` is returned as `*ValidationErrors`.

4. **Non-nil empty struct trap:** Remember from Exercise 07 - a pointer to an empty `ValidationErrors{}` is NOT nil. Always check `HasErrors()` before returning the error.

5. **Multi-line formatting:** For `ValidationErrors.Error()` with multiple errors, consider using `strings.Builder` for efficiency:
   ```go
   var b strings.Builder
   b.WriteString("validation errors:\n")
   for _, e := range v.Errors {
       fmt.Fprintf(&b, "  - %s: %s\n", e.Field, e.Message)
   }
   return b.String()
   ```

6. **Status code ranges:** For 5xx errors, check `statusCode >= 500 && statusCode < 600` rather than listing each code.
