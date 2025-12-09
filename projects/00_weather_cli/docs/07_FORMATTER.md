# Module: Formatter

## Purpose

Format combined weather and location data for output. This module transforms internal data structures into human-readable table format or machine-readable JSON, providing a clean separation between data processing and presentation.

---

## Package Location

```
formatter/formatter.go
```

---

## Dependencies

| Package | Purpose |
|---------|---------|
| `types` (internal) | `WeatherData`, `Location` types |

---

## Public API

### Types

```go
// FormatOptions configures output formatting
type FormatOptions struct {
    Format   string // "table" or "json"
    Timezone string // Timezone for displaying times (e.g., "UTC", "Local")
}

// FormatterInput combines all data for formatting
type FormatterInput struct {
    Weather  *types.WeatherData
    Location *types.Location
    Options  FormatOptions
}

// FormattedOutput contains the result
type FormattedOutput struct {
    Content    string // The formatted string
    DataPoints int    // Number of hourly data points
}
```

### Functions

```go
// Format produces formatted output from weather and location data
func Format(input FormatterInput) (*FormattedOutput, error)

// FormatTable produces ASCII table output
func FormatTable(input FormatterInput) (string, error)

// FormatJSON produces JSON output
func FormatJSON(input FormatterInput) (string, error)
```

---

## Output Formats

### Table Format

```
┌─────────────────────────────────────────────────────────────────────┐
│  Weather Report: Berlin, Germany                                    │
│  Coordinates: 52.5200°N, 13.4050°E                                  │
│  Period: 2024-01-01 to 2024-01-03                                   │
├─────────────────────────────────────────────────────────────────────┤
│  Time              │ Temp (°C) │ Humidity │ Precip (mm) │ Wind (km/h) │
├─────────────────────────────────────────────────────────────────────┤
│  2024-01-01 00:00  │     2.3   │    85%   │     0.0     │     12.5    │
│  2024-01-01 01:00  │     2.1   │    87%   │     0.0     │     11.2    │
│  2024-01-01 02:00  │     1.9   │    88%   │     0.1     │     10.8    │
│  ...               │    ...    │   ...    │    ...      │    ...      │
├─────────────────────────────────────────────────────────────────────┤
│  Summary: 72 hourly data points                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Requirements:**
- Header with location name and coordinates
- Date range display
- Column headers with units
- Right-aligned numeric values
- Summary with data point count
- Unicode box-drawing characters for borders

### JSON Format

```json
{
  "location": {
    "name": "Berlin, Germany",
    "latitude": 52.52,
    "longitude": 13.405
  },
  "period": {
    "start": "2024-01-01",
    "end": "2024-01-03"
  },
  "units": {
    "temperature": "°C",
    "humidity": "%",
    "precipitation": "mm",
    "wind_speed": "km/h"
  },
  "hourly": [
    {
      "time": "2024-01-01T00:00:00Z",
      "temperature": 2.3,
      "humidity": 85,
      "precipitation": 0.0,
      "wind_speed": 12.5
    },
    ...
  ],
  "summary": {
    "data_points": 72
  }
}
```

**Requirements:**
- Structured JSON with clear sections
- ISO 8601 timestamps
- Consistent field naming (snake_case)
- Pretty-printed with indentation

---

## Behavior Specification

### Format(input FormatterInput)

**Input:** `FormatterInput` with weather data, location, and options

**Output:** `(*FormattedOutput, error)`

**Behavior:**
1. Validate input (weather data not nil)
2. Based on `Options.Format`, call `FormatTable` or `FormatJSON`
3. Count data points
4. Return `FormattedOutput` with content and count

```go
func Format(input FormatterInput) (*FormattedOutput, error) {
    if input.Weather == nil {
        return nil, errors.New("weather data is required")
    }

    var content string
    var err error

    switch strings.ToLower(input.Options.Format) {
    case "table":
        content, err = FormatTable(input)
    case "json":
        content, err = FormatJSON(input)
    default:
        return nil, fmt.Errorf("unknown format: %s", input.Options.Format)
    }

    if err != nil {
        return nil, err
    }

    return &FormattedOutput{
        Content:    content,
        DataPoints: len(input.Weather.Hourly.Time),
    }, nil
}
```

---

### FormatTable(input FormatterInput)

**Input:** `FormatterInput`

**Output:** `(string, error)`

**Behavior:**
1. Build header section with location and coordinates
2. Build column headers with units
3. For each hourly data point:
   - Format timestamp
   - Format numeric values with appropriate precision
   - Right-align in columns
4. Add summary footer
5. Return complete table string

**Column Specifications:**

| Column | Width | Alignment | Format |
|--------|-------|-----------|--------|
| Time | 18 | Left | `2006-01-02 15:04` |
| Temperature | 10 | Right | `%.1f` |
| Humidity | 8 | Right | `%.0f%%` |
| Precipitation | 12 | Right | `%.1f` |
| Wind Speed | 12 | Right | `%.1f` |

---

### FormatJSON(input FormatterInput)

**Input:** `FormatterInput`

**Output:** `(string, error)`

**Behavior:**
1. Create output structure matching JSON schema
2. Populate location section
3. Populate period from first/last timestamps
4. Copy units from weather data
5. Build hourly array with formatted entries
6. Marshal with indentation
7. Return JSON string

---

## Edge Cases

| Scenario | Behavior |
|----------|----------|
| `Weather` is nil | Return error "weather data is required" |
| `Location` is nil | Display "Unknown Location" in output |
| Empty hourly data | Valid output with 0 data points |
| Negative temperatures | Display with minus sign: `-2.3` |
| Very long location name | Truncate or wrap in table format |
| Invalid format option | Return error |
| Missing humidity values | Display `N/A` or skip |

---

## Useful Packages

| Package | Why Useful |
|---------|------------|
| `strings` | `strings.Builder` for efficient string building |
| `fmt` | `fmt.Sprintf()` for formatting values |
| `encoding/json` | `json.MarshalIndent()` for JSON output |
| `text/tabwriter` | Alternative for aligned columns |
| `time` | Time formatting |

---

## Test Cases

| Scenario | Input | Expected |
|----------|-------|----------|
| Table format basic | Valid weather + location | Table with header, rows, footer |
| JSON format basic | Valid weather + location | Valid JSON with all sections |
| Unknown location | Weather only, no location | "Unknown Location" in output |
| Empty data | WeatherData with empty Hourly | Valid output, 0 data points |
| Invalid format | Format: "xml" | Error returned |
| Nil weather | Weather: nil | Error returned |
| Negative temps | Temperature: -5.2 | Displayed as "-5.2" |
| High precision | Temp: 2.3456789 | Displayed as "2.3" |

### Testing Output Format

```go
func TestFormatTable_Basic(t *testing.T) {
    input := FormatterInput{
        Weather: &types.WeatherData{
            Hourly: types.HourlyData{
                Time:        []time.Time{time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
                Temperature: []float64{2.3},
                Humidity:    []float64{85},
                // ...
            },
        },
        Location: &types.Location{
            DisplayName: "Berlin, Germany",
            Latitude:    52.52,
            Longitude:   13.41,
        },
        Options: FormatOptions{Format: "table"},
    }

    output, err := Format(input)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    // Check content contains expected elements
    if !strings.Contains(output.Content, "Berlin") {
        t.Error("expected output to contain location name")
    }
    if !strings.Contains(output.Content, "2.3") {
        t.Error("expected output to contain temperature")
    }
    if output.DataPoints != 1 {
        t.Errorf("expected 1 data point, got %d", output.DataPoints)
    }
}

func TestFormatJSON_Valid(t *testing.T) {
    // ... setup input ...

    output, err := Format(input)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    // Verify it's valid JSON
    var parsed map[string]interface{}
    if err := json.Unmarshal([]byte(output.Content), &parsed); err != nil {
        t.Errorf("invalid JSON: %v", err)
    }

    // Verify structure
    if _, ok := parsed["location"]; !ok {
        t.Error("expected 'location' field in JSON")
    }
}
```

---

## Implementation Hints

1. **Use strings.Builder for tables:** More efficient than string concatenation:
   ```go
   var b strings.Builder
   b.WriteString("┌────────────────────┐\n")
   b.WriteString(fmt.Sprintf("│ %-18s │\n", header))
   // ...
   return b.String(), nil
   ```

2. **Unicode box characters:**
   ```go
   const (
       topLeft     = "┌"
       topRight    = "┐"
       bottomLeft  = "└"
       bottomRight = "┘"
       horizontal  = "─"
       vertical    = "│"
       cross       = "┼"
       teeDown     = "┬"
       teeUp       = "┴"
       teeRight    = "├"
       teeLeft     = "┤"
   )
   ```

3. **Column width calculation:** For dynamic widths, find the max length:
   ```go
   maxCityLen := len("Unknown Location")
   if input.Location != nil && len(input.Location.DisplayName) > maxCityLen {
       maxCityLen = len(input.Location.DisplayName)
   }
   ```

4. **JSON structure with anonymous structs:**
   ```go
   output := struct {
       Location struct {
           Name      string  `json:"name"`
           Latitude  float64 `json:"latitude"`
           Longitude float64 `json:"longitude"`
       } `json:"location"`
       Hourly []struct {
           Time        string  `json:"time"`
           Temperature float64 `json:"temperature"`
           // ...
       } `json:"hourly"`
   }{}
   ```

5. **Handle nil Location gracefully:**
   ```go
   locationName := "Unknown Location"
   if input.Location != nil {
       locationName = input.Location.DisplayName
   }
   ```

6. **Time formatting for display:**
   ```go
   t.Format("2006-01-02 15:04")      // Table: "2024-01-01 00:00"
   t.Format(time.RFC3339)            // JSON: "2024-01-01T00:00:00Z"
   ```

7. **Iterate parallel slices safely:**
   ```go
   for i := 0; i < len(hourly.Time); i++ {
       temp := "N/A"
       if i < len(hourly.Temperature) {
           temp = fmt.Sprintf("%.1f", hourly.Temperature[i])
       }
       // ...
   }
   ```

---

## Example Usage

```go
// In main.go
input := formatter.FormatterInput{
    Weather:  weatherData,
    Location: locationData,
    Options: formatter.FormatOptions{
        Format: "table",
    },
}

output, err := formatter.Format(input)
if err != nil {
    fmt.Fprintf(os.Stderr, "Error: %v\n", err)
    os.Exit(1)
}

fmt.Print(output.Content)
fmt.Fprintf(os.Stderr, "Displayed %d hourly data points\n", output.DataPoints)
```
