# Module: Weather Client

## Purpose

Fetch historical weather data from the Open-Meteo Archive API. This module handles HTTP communication, response parsing, error classification, and integrates with the retry module for transient failure handling.

---

## Package Location

```
client/weather.go
```

---

## Dependencies

| Package | Purpose |
|---------|---------|
| `errors` (internal) | `WeatherAPIError`, sentinel errors |
| `types` (internal) | `WeatherData`, `HourlyData` |
| `client/retry` (internal) | `RetryWithBackoff` for transient failures |

---

## Open-Meteo Archive API Specification

### Endpoint

```
https://archive-api.open-meteo.com/v1/archive
```

### Request Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `latitude` | float | Yes | Latitude (-90 to 90) |
| `longitude` | float | Yes | Longitude (-180 to 180) |
| `start_date` | string | Yes | Start date (YYYY-MM-DD) |
| `end_date` | string | Yes | End date (YYYY-MM-DD) |
| `hourly` | string | Yes | Comma-separated variables to fetch |
| `timezone` | string | No | Timezone for timestamps (default: UTC) |

### Example Request

```
GET https://archive-api.open-meteo.com/v1/archive?latitude=52.52&longitude=13.41&start_date=2024-01-01&end_date=2024-01-07&hourly=temperature_2m,relative_humidity_2m&timezone=UTC
```

### Response Structure

```json
{
  "latitude": 52.52,
  "longitude": 13.419998,
  "generationtime_ms": 0.5,
  "utc_offset_seconds": 0,
  "timezone": "UTC",
  "timezone_abbreviation": "UTC",
  "elevation": 38.0,
  "hourly_units": {
    "time": "iso8601",
    "temperature_2m": "°C",
    "relative_humidity_2m": "%"
  },
  "hourly": {
    "time": ["2024-01-01T00:00", "2024-01-01T01:00", ...],
    "temperature_2m": [2.3, 2.1, 1.9, ...],
    "relative_humidity_2m": [85, 87, 88, ...]
  }
}
```

### Error Responses

| Status | Meaning | Example Response |
|--------|---------|------------------|
| 400 | Bad request | `{"error": true, "reason": "Invalid date format"}` |
| 404 | Not found | Location outside coverage area |
| 429 | Rate limited | Too many requests |
| 5xx | Server error | Internal error |

---

## Public API

### Types

```go
// WeatherClient handles communication with Open-Meteo API
type WeatherClient struct {
    baseURL    string
    httpClient *http.Client
}

// WeatherRequest contains parameters for a weather data request
type WeatherRequest struct {
    Latitude  float64
    Longitude float64
    StartDate time.Time
    EndDate   time.Time
    Variables []string // e.g., ["temperature_2m", "relative_humidity_2m"]
}
```

### Functions

```go
// NewWeatherClient creates a new client with the given timeout
func NewWeatherClient(timeout time.Duration) *WeatherClient

// FetchWeather retrieves weather data for the given parameters
// Uses retry logic for transient failures
func (c *WeatherClient) FetchWeather(ctx context.Context, req WeatherRequest) (*types.WeatherData, error)
```

---

## Behavior Specification

### NewWeatherClient(timeout time.Duration)

**Input:** `timeout` - HTTP request timeout duration

**Output:** `*WeatherClient`

**Behavior:**
1. Create `http.Client` with specified timeout
2. Set base URL to `https://archive-api.open-meteo.com/v1/archive`
3. Return configured client

---

### FetchWeather(ctx context.Context, req WeatherRequest)

**Input:**
- `ctx` - Context for cancellation/timeout
- `req` - Request parameters

**Output:** `(*types.WeatherData, error)`

**Behavior:**

```
FetchWeather(ctx, req)
   │
   ├──► Build URL with query parameters
   │
   ├──► Create HTTP request with context
   │
   ├──► RetryWithBackoff(func() error {
   │         │
   │         ├── Execute HTTP request
   │         │
   │         ├── Check status code
   │         │     ├── 200 → Parse JSON → return nil
   │         │     ├── 400 → return WeatherAPIError (non-retryable)
   │         │     ├── 404 → return WeatherAPIError wrapping ErrNotFound
   │         │     ├── 429 → return WeatherAPIError wrapping ErrRateLimited
   │         │     └── 5xx → return WeatherAPIError (retryable)
   │         │
   │         └── Network error → return wrapped error
   │    })
   │
   └──► Return WeatherData or error
```

### URL Building

Build query string with:
```
?latitude={lat}&longitude={long}&start_date={YYYY-MM-DD}&end_date={YYYY-MM-DD}&hourly={vars}&timezone=UTC
```

Variables to request: `temperature_2m,relative_humidity_2m,precipitation,wind_speed_10m`

---

## Error Handling

### HTTP Status Code Mapping

```go
switch resp.StatusCode {
case http.StatusOK:
    // Parse response
case http.StatusBadRequest:
    return &WeatherAPIError{
        StatusCode: 400,
        Message:    parseErrorMessage(body),
        Latitude:   req.Latitude,
        Longitude:  req.Longitude,
    }
case http.StatusNotFound:
    return &WeatherAPIError{
        StatusCode: 404,
        Message:    "Location not found or outside coverage area",
        Latitude:   req.Latitude,
        Longitude:  req.Longitude,
    }
case http.StatusTooManyRequests:
    return &WeatherAPIError{
        StatusCode: 429,
        Message:    "Rate limit exceeded",
        Latitude:   req.Latitude,
        Longitude:  req.Longitude,
    }
default:
    if resp.StatusCode >= 500 {
        return &WeatherAPIError{
            StatusCode: resp.StatusCode,
            Message:    "Server error",
            // ...
        }
    }
}
```

### Network Errors

Wrap network errors with context:
```go
resp, err := c.httpClient.Do(req)
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        return fmt.Errorf("weather API timeout: %w", errors.ErrTimeout)
    }
    return fmt.Errorf("weather API network error: %w", err)
}
```

---

## Response Parsing

### Target Data Structure (from types package)

```go
// In types/types.go
type WeatherData struct {
    Latitude    float64
    Longitude   float64
    Timezone    string
    HourlyUnits map[string]string // e.g., {"temperature_2m": "°C"}
    Hourly      HourlyData
}

type HourlyData struct {
    Time             []time.Time
    Temperature      []float64 // temperature_2m
    Humidity         []float64 // relative_humidity_2m
    Precipitation    []float64 // precipitation
    WindSpeed        []float64 // wind_speed_10m
}
```

### Parsing Logic

1. Decode JSON into intermediate struct matching API response
2. Parse time strings into `time.Time` values
3. Copy data into typed `HourlyData` struct
4. Return `WeatherData`

**Important:** Handle missing data gracefully. Some hours might have `null` values in the API response.

---

## Useful Packages

| Package | Why Useful |
|---------|------------|
| `net/http` | HTTP client, request building |
| `net/url` | `url.Values` for query string building |
| `encoding/json` | JSON parsing |
| `context` | Request cancellation and timeout |
| `time` | Date formatting, parsing |
| `io` | `io.ReadAll()` for response body |
| `fmt` | Error message formatting |

---

## Test Cases

| Scenario | Input | Expected |
|----------|-------|----------|
| Successful fetch | Valid lat/long/dates | WeatherData with hourly data |
| HTTP 404 | Invalid location | WeatherAPIError with ErrNotFound |
| HTTP 429 | Rapid requests | WeatherAPIError with ErrRateLimited |
| HTTP 500 | Server error | WeatherAPIError, retried |
| Timeout | Slow server | Error wrapping ErrTimeout |
| Context canceled | Canceled context | Error from context |
| Invalid JSON | Malformed response | Parse error |
| Empty hourly data | Valid request, no data | Empty slices, no error |

### Testing HTTP Clients

Use `httptest.NewServer()` to mock API responses:

```go
func TestFetchWeather_Success(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(mockResponse)
    }))
    defer server.Close()

    client := &WeatherClient{
        baseURL:    server.URL,
        httpClient: &http.Client{Timeout: 5 * time.Second},
    }

    data, err := client.FetchWeather(context.Background(), req)
    // assertions...
}
```

---

## Implementation Hints

1. **URL building:** Use `url.Values` for proper query string encoding:
   ```go
   params := url.Values{}
   params.Set("latitude", fmt.Sprintf("%f", req.Latitude))
   params.Set("longitude", fmt.Sprintf("%f", req.Longitude))
   params.Set("start_date", req.StartDate.Format("2006-01-02"))
   params.Set("end_date", req.EndDate.Format("2006-01-02"))
   params.Set("hourly", strings.Join(defaultVariables, ","))
   params.Set("timezone", "UTC")

   fullURL := c.baseURL + "?" + params.Encode()
   ```

2. **Context propagation:** Always use `http.NewRequestWithContext()`:
   ```go
   req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
   ```

3. **Body reading:** Always close the response body:
   ```go
   defer resp.Body.Close()
   body, err := io.ReadAll(resp.Body)
   ```

4. **Intermediate struct for parsing:** The API returns `hourly.time` as strings. Parse them:
   ```go
   type apiResponse struct {
       Hourly struct {
           Time          []string  `json:"time"`
           Temperature2m []float64 `json:"temperature_2m"`
           // ...
       } `json:"hourly"`
   }

   // Then convert:
   for _, ts := range apiResp.Hourly.Time {
       t, _ := time.Parse("2006-01-02T15:04", ts)
       result.Hourly.Time = append(result.Hourly.Time, t)
   }
   ```

5. **Default variables:** Define constants for the weather variables to fetch:
   ```go
   var defaultVariables = []string{
       "temperature_2m",
       "relative_humidity_2m",
       "precipitation",
       "wind_speed_10m",
   }
   ```

6. **Retry integration:** Wrap the HTTP call in the retry function:
   ```go
   var result *types.WeatherData
   err := RetryWithBackoff(func() error {
       // ... HTTP request logic ...
       result = parsed
       return nil
   }, DefaultRetryConfig())

   return result, err
   ```
