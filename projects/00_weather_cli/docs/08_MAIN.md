# Module: Main Orchestration

## Purpose

The main package serves as the entry point and orchestration layer. It parses CLI flags, coordinates calls to all other modules, handles top-level errors with user-friendly messages, and manages exit codes.

---

## Package Location

```
main.go
```

---

## Dependencies

| Package | Purpose |
|---------|---------|
| `config` (internal) | Load configuration |
| `validation` (internal) | Validate CLI inputs |
| `client` (internal) | Weather and Geocoder clients |
| `formatter` (internal) | Output formatting |
| `errors` (internal) | Error types for classification |
| `types` (internal) | Shared data structures |

---

## CLI Interface

### Flags

| Flag | Type | Required | Default | Description |
|------|------|----------|---------|-------------|
| `-lat` | string | Yes | - | Latitude (-90 to 90) |
| `-long` | string | Yes | - | Longitude (-180 to 180) |
| `-start` | string | Yes | - | Start date (YYYY-MM-DD) |
| `-end` | string | Yes | - | End date (YYYY-MM-DD) |
| `-format` | string | No | `table` | Output format: `table` or `json` |

### Usage

```bash
# Basic usage
weather-cli -lat 52.52 -long 13.41 -start 2024-01-01 -end 2024-01-07

# JSON output
weather-cli -lat 52.52 -long 13.41 -start 2024-01-01 -end 2024-01-07 -format json

# Help
weather-cli -h
```

---

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | Error (any type) |

---

## Orchestration Flow

```
main()
   │
   ├──► Parse CLI flags
   │         └── flag.Parse()
   │
   ├──► Load configuration
   │         ├── cfg, err := config.Load()
   │         └── if err → handleError(err) → os.Exit(1)
   │
   ├──► Validate inputs
   │         ├── validated, err := validation.Validate(input)
   │         └── if err → handleError(err) → os.Exit(1)
   │
   ├──► Create clients
   │         ├── weatherClient := client.NewWeatherClient(timeout)
   │         └── geocoderClient := client.NewGeocoderClient(cfg.GeocodeAPIKey, timeout)
   │
   ├──► Fetch data (can run in parallel or sequential)
   │         │
   │         ├──► weatherData, err := weatherClient.FetchWeather(ctx, req)
   │         │         └── if err → handleError(err) → os.Exit(1)
   │         │
   │         └──► locationData, err := geocoderClient.ReverseGeocode(ctx, req)
   │                   └── if err → log warning, continue with nil location
   │
   ├──► Format output
   │         ├── output, err := formatter.Format(input)
   │         └── if err → handleError(err) → os.Exit(1)
   │
   ├──► Print output
   │         └── fmt.Print(output.Content)
   │
   └──► os.Exit(0)
```

---

## Error Handling Strategy

### Error Classification and Messages

```go
func handleError(err error) {
    var msg string

    // Check for validation errors (aggregate)
    var validationErrs *errors.ValidationErrors
    if errors.As(err, &validationErrs) {
        fmt.Fprintln(os.Stderr, validationErrs.Error())
        return
    }

    // Check for specific sentinel errors
    switch {
    case errors.Is(err, errors.ErrNotFound):
        msg = "Location not found. Please check the coordinates."

    case errors.Is(err, errors.ErrRateLimited):
        msg = "Rate limit exceeded. Please wait and try again."

    case errors.Is(err, errors.ErrUnauthorized):
        msg = "Invalid or missing API key. Check your .env file."

    case errors.Is(err, errors.ErrTimeout):
        msg = "Request timed out. Check your internet connection."

    case errors.Is(err, errors.ErrNetwork):
        msg = "Network error. Check your internet connection."

    default:
        // For unexpected errors, show the error message
        msg = fmt.Sprintf("Error: %v", err)
    }

    fmt.Fprintln(os.Stderr, msg)
}
```

### Error Output Rules

1. **All errors go to stderr:** `fmt.Fprintln(os.Stderr, ...)`
2. **User-friendly messages:** No stack traces or technical details
3. **Validation errors:** Show all errors at once
4. **Exit code 1:** Always exit with 1 on any error

---

## Behavior Specification

### main()

**Behavior:**

1. **Flag Parsing:**
   ```go
   var lat, long, start, end, format string

   flag.StringVar(&lat, "lat", "", "Latitude (-90 to 90)")
   flag.StringVar(&long, "long", "", "Longitude (-180 to 180)")
   flag.StringVar(&start, "start", "", "Start date (YYYY-MM-DD)")
   flag.StringVar(&end, "end", "", "End date (YYYY-MM-DD)")
   flag.StringVar(&format, "format", "table", "Output format: table or json")

   flag.Parse()
   ```

2. **Configuration Loading:**
   ```go
   cfg, err := config.Load()
   if err != nil {
       handleError(err)
       os.Exit(1)
   }
   ```

3. **Input Validation:**
   ```go
   input := validation.CLIInput{
       Latitude:  lat,
       Longitude: long,
       StartDate: start,
       EndDate:   end,
       Format:    format,
   }

   validated, err := validation.Validate(input)
   if err != nil {
       handleError(err)
       os.Exit(1)
   }
   ```

4. **Client Creation:**
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
   defer cancel()

   weatherClient := client.NewWeatherClient(10 * time.Second)
   geocoderClient := client.NewGeocoderClient(cfg.GeocodeAPIKey, 10*time.Second)
   ```

5. **Data Fetching:**
   ```go
   weatherReq := client.WeatherRequest{
       Latitude:  validated.Latitude,
       Longitude: validated.Longitude,
       StartDate: validated.StartDate,
       EndDate:   validated.EndDate,
   }

   weatherData, err := weatherClient.FetchWeather(ctx, weatherReq)
   if err != nil {
       handleError(err)
       os.Exit(1)
   }

   // Geocoding is optional - failure doesn't stop the program
   geocodeReq := client.GeocoderRequest{
       Latitude:  validated.Latitude,
       Longitude: validated.Longitude,
   }

   locationData, err := geocoderClient.ReverseGeocode(ctx, geocodeReq)
   if err != nil {
       // Log warning but continue
       fmt.Fprintf(os.Stderr, "Warning: Could not fetch location name: %v\n", err)
       locationData = nil
   }
   ```

6. **Formatting and Output:**
   ```go
   formatInput := formatter.FormatterInput{
       Weather:  weatherData,
       Location: locationData,
       Options: formatter.FormatOptions{
           Format: validated.Format,
       },
   }

   output, err := formatter.Format(formatInput)
   if err != nil {
       handleError(err)
       os.Exit(1)
   }

   fmt.Print(output.Content)
   ```

---

## Edge Cases

| Scenario | Behavior |
|----------|----------|
| No flags provided | Validation errors for all required fields |
| Invalid lat/long | Validation error |
| API failure | Error message to stderr, exit 1 |
| Geocoder failure | Warning to stderr, continue with "Unknown Location" |
| Empty date range | Valid if start == end |
| Keyboard interrupt | Clean exit (Go handles this) |

---

## Example Error Messages

```bash
# Validation errors (aggregated)
$ weather-cli -lat abc -long xyz -start bad
validation errors:
  - latitude: must be a valid number
  - longitude: must be a valid number
  - start_date: must be in YYYY-MM-DD format
  - end_date: is required

# Not found
$ weather-cli -lat 0 -long 0 -start 2024-01-01 -end 2024-01-02
Location not found. Please check the coordinates.

# Missing API key
$ weather-cli -lat 52.52 -long 13.41 -start 2024-01-01 -end 2024-01-02
Invalid or missing API key. Check your .env file.

# Network error
$ weather-cli -lat 52.52 -long 13.41 -start 2024-01-01 -end 2024-01-02
Network error. Check your internet connection.
```

---

## Useful Packages

| Package | Why Useful |
|---------|------------|
| `flag` | CLI flag parsing |
| `os` | `os.Exit()`, `os.Stderr` |
| `fmt` | Output formatting |
| `context` | Request timeout and cancellation |
| `time` | Timeout duration |
| `errors` | Error inspection with Is/As |

---

## Test Cases

Testing `main()` directly is difficult. Instead, test the components and the orchestration logic:

| Scenario | Test Approach |
|----------|---------------|
| Flag parsing | Test with `os.Args` manipulation or extract logic |
| Error handling | Test `handleError()` function separately |
| Happy path | Integration test with mocked clients |
| Validation failure | Unit test validation module |
| API failure | Unit test clients with mock servers |

### Testing handleError

```go
func TestHandleError_ValidationErrors(t *testing.T) {
    // Capture stderr
    oldStderr := os.Stderr
    r, w, _ := os.Pipe()
    os.Stderr = w

    verrs := &errors.ValidationErrors{}
    verrs.Add("lat", "must be a number")
    verrs.Add("long", "must be a number")

    handleError(verrs)

    w.Close()
    os.Stderr = oldStderr

    var buf bytes.Buffer
    io.Copy(&buf, r)
    output := buf.String()

    if !strings.Contains(output, "lat") {
        t.Error("expected output to contain 'lat'")
    }
    if !strings.Contains(output, "long") {
        t.Error("expected output to contain 'long'")
    }
}
```

---

## Implementation Hints

1. **Keep main() thin:** Extract logic into functions:
   ```go
   func main() {
       if err := run(); err != nil {
           handleError(err)
           os.Exit(1)
       }
   }

   func run() error {
       // All the actual logic
   }
   ```

2. **Use context for timeout:**
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
   defer cancel()
   ```

3. **Geocoder is optional:** Don't fail the whole program if geocoding fails:
   ```go
   location, err := geocoderClient.ReverseGeocode(ctx, req)
   if err != nil {
       fmt.Fprintf(os.Stderr, "Warning: %v\n", err)
       location = nil // Formatter handles nil gracefully
   }
   ```

4. **Flag defaults:** Provide sensible defaults where applicable:
   ```go
   flag.StringVar(&format, "format", "table", "Output format")
   ```

5. **stderr for all non-data output:**
   ```go
   // Data goes to stdout
   fmt.Print(output.Content)

   // Everything else goes to stderr
   fmt.Fprintln(os.Stderr, "Warning: ...")
   fmt.Fprintln(os.Stderr, "Error: ...")
   ```

6. **Help text:** `flag` package auto-generates help with `-h`:
   ```bash
   $ weather-cli -h
   Usage of weather-cli:
     -lat string
           Latitude (-90 to 90)
     -long string
           Longitude (-180 to 180)
     ...
   ```

7. **Required flags:** The `flag` package doesn't support required flags natively. Validation handles this:
   ```go
   if lat == "" {
       // Caught by validation.Validate()
   }
   ```

---

## Success Criteria

A complete implementation should:

1. ✅ Accept all specified flags
2. ✅ Show aggregated validation errors (all at once)
3. ✅ Fetch weather data with retry on transient failures
4. ✅ Fetch location data (optional, warning on failure)
5. ✅ Output formatted data to stdout
6. ✅ Output errors to stderr
7. ✅ Exit with 0 on success, 1 on error
8. ✅ Handle Ctrl+C gracefully (context cancellation)
