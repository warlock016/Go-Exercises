# Weather CLI - Project Specification

A command-line weather tool that fetches historical weather data from the Open-Meteo API.

**Estimated Time:** 2-4 hours
**Difficulty:** Beginner-Intermediate
**Prerequisites:** Your working code in `playground/main.go`

---

## Learning Goals

This project practices:
- **CLI flags** - Using the `flag` package for command-line arguments (NEW)
- **Code organization** - Separating concerns into multiple files (NEW)
- **Error handling** - Returning errors instead of `log.Fatal` (NEW pattern)
- **HTTP client** - Making API requests with `net/http` (you know this)
- **JSON parsing** - Unmarshaling API responses into structs (you know this)

---

## Requirements

### Functional Requirements

1. **Fetch weather data** from Open-Meteo Archive API
2. **Accept command-line flags** for customization:
   - `-lat` (float64, default: 52.51) - Latitude
   - `-lon` (float64, default: 13.41) - Longitude
   - `-city` (string, default: "Berlin") - City name for display
   - `-start` (string, default: yesterday) - Start date (YYYY-MM-DD)
   - `-end` (string, default: yesterday) - End date (YYYY-MM-DD)
   - `-format` (string, default: "table") - Output format: "table" or "json"
3. **Display output** in either table or JSON format
4. **Handle errors gracefully** - print to stderr, exit with code 1

### Non-Functional Requirements

- Code should be organized into multiple files (not everything in main.go)
- API client function should return `(data, error)` not call `log.Fatal`
- Invalid flag values should produce helpful error messages

---

## API Reference

**Base URL:** `https://archive-api.open-meteo.com/v1/archive`

**Query Parameters:**
- `latitude` - Latitude coordinate
- `longitude` - Longitude coordinate
- `start_date` - Start date (YYYY-MM-DD)
- `end_date` - End date (YYYY-MM-DD)
- `hourly` - Variables to fetch (use `temperature_2m`)
- `timezone` - Timezone (use `GMT`)

**Example Response:**
```json
{
  "latitude": 52.52,
  "longitude": 13.419,
  "hourly": {
    "time": ["2025-11-01T00:00", "2025-11-01T01:00", ...],
    "temperature_2m": [8.5, 7.2, ...]
  }
}
```

You already have working code for this in `playground/main.go` - adapt it!

---

## Expected Output

### Table Format (default)
```
Weather for Berlin (52.52, 13.42)
========================================
2025-11-01T00:00  |    8.5°C
2025-11-01T01:00  |    7.2°C
2025-11-01T02:00  |    6.8°C
...
========================================
Total data points: 24
```

### JSON Format (-format json)
```json
{
  "latitude": 52.52,
  "longitude": 13.419,
  "hourly": {
    "time": [...],
    "temperature_2m": [...]
  }
}
```

### Error Cases
```
# Invalid format flag
Error: invalid format "xml", must be "table" or "json"

# Network error
Error: failed to fetch weather data: <details>

# Invalid date
Error: API error (status 400): <API message>
```

---

## Suggested File Structure

```
00_weather_cli/
├── main.go           # Entry point: flag parsing, orchestration
├── weather.go        # API client: FetchWeather function
├── formatter.go      # Output: PrintTable, PrintJSON functions
└── README.md         # This spec
```

You decide the exact function signatures and struct definitions.

---

## Hints (Read Only If Stuck)

<details>
<summary>Hint 1: Flag package basics</summary>

```go
import "flag"

// flag.Type returns a POINTER
lat := flag.Float64("lat", 52.51, "Latitude")
flag.Parse()
fmt.Println(*lat)  // Dereference to get value
```
</details>

<details>
<summary>Hint 2: Getting yesterday's date</summary>

```go
import "time"

yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
```
</details>

<details>
<summary>Hint 3: Writing to stderr</summary>

```go
import "os"

fmt.Fprintf(os.Stderr, "Error: %v\n", err)
os.Exit(1)
```
</details>

<details>
<summary>Hint 4: Function signature for API client</summary>

Consider returning a pointer and error:
```go
func FetchWeather(...) (*WeatherData, error)
```
This allows the caller to decide how to handle errors.
</details>

---

## Success Criteria

- [ ] `go run . -help` shows all flags with descriptions
- [ ] `go run .` fetches and displays Berlin weather for yesterday
- [ ] `go run . -lat 40.71 -lon -74.01 -city "New York"` works
- [ ] `go run . -format json` outputs valid JSON
- [ ] `go run . -format invalid` prints error to stderr and exits with code 1
- [ ] Code is split across at least 2-3 files
- [ ] No `log.Fatal` in the API client function

---

## Getting Started

1. Copy relevant parts from your `playground/main.go` as a starting point
2. Start with `main.go` - get flag parsing working first
3. Move API logic to `weather.go`
4. Add formatting to `formatter.go`
5. Test each piece as you go

Ask for help if you're stuck on a specific concept, but try to implement each part yourself first!
