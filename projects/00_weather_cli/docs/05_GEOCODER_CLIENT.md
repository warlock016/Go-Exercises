# Module: Geocoder Client

## Purpose

Perform reverse geocoding to convert latitude/longitude coordinates into human-readable location names using the Nominatim API. This module provides location context (city, country) to display alongside weather data.

---

## Package Location

```
client/geocoder.go
```

---

## Dependencies

| Package | Purpose |
|---------|---------|
| `errors` (internal) | `GeocodeError`, sentinel errors |
| `types` (internal) | `Location` struct |
| `client/retry` (internal) | `RetryWithBackoff` for transient failures |

---

## Nominatim API Specification

### Endpoint (Reverse Geocoding)

```
https://nominatim.openstreetmap.org/reverse
```

**Note:** For production use with higher rate limits, you can use a commercial Nominatim provider that requires an API key. The implementation should support both authenticated and unauthenticated requests.

### Request Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `lat` | float | Yes | Latitude |
| `lon` | float | Yes | Longitude |
| `format` | string | Yes | Response format (use "json") |
| `zoom` | int | No | Detail level (10 = city level) |
| `addressdetails` | int | No | Include address breakdown (1 = yes) |

### Request Headers

| Header | Value | Purpose |
|--------|-------|---------|
| `User-Agent` | `WeatherCLI/1.0` | **Required** - Nominatim blocks requests without User-Agent |
| `Accept-Language` | `en` | Preferred language for results |

### Example Request

```
GET https://nominatim.openstreetmap.org/reverse?lat=52.52&lon=13.41&format=json&zoom=10&addressdetails=1
```

### Response Structure

```json
{
  "place_id": 123456,
  "licence": "Data © OpenStreetMap contributors, ODbL 1.0.",
  "lat": "52.5170365",
  "lon": "13.3888599",
  "display_name": "Berlin, Germany",
  "address": {
    "city": "Berlin",
    "state": "Berlin",
    "country": "Germany",
    "country_code": "de"
  }
}
```

### Error Responses

| Status | Meaning |
|--------|---------|
| 400 | Invalid parameters |
| 403 | Blocked (missing User-Agent) |
| 429 | Rate limited (max 1 req/sec for free tier) |
| 5xx | Server error |

### Rate Limiting

**Free tier limits:**
- Maximum 1 request per second
- No bulk geocoding

**Best practice:** Add a small delay between requests if making multiple calls.

---

## Public API

### Types

```go
// GeocoderClient handles communication with Nominatim API
type GeocoderClient struct {
    baseURL    string
    httpClient *http.Client
    apiKey     string // Optional, for commercial providers
    userAgent  string
}

// GeocoderRequest contains parameters for a geocoding request
type GeocoderRequest struct {
    Latitude  float64
    Longitude float64
}
```

### Functions

```go
// NewGeocoderClient creates a new client
// apiKey can be empty for free Nominatim tier
func NewGeocoderClient(apiKey string, timeout time.Duration) *GeocoderClient

// ReverseGeocode converts coordinates to a location name
func (c *GeocoderClient) ReverseGeocode(ctx context.Context, req GeocoderRequest) (*types.Location, error)
```

---

## Behavior Specification

### NewGeocoderClient(apiKey string, timeout time.Duration)

**Input:**
- `apiKey` - API key (empty string for free tier)
- `timeout` - HTTP request timeout

**Output:** `*GeocoderClient`

**Behavior:**
1. Create `http.Client` with specified timeout
2. Set base URL to Nominatim endpoint
3. Set User-Agent to `WeatherCLI/1.0`
4. Store API key (may be empty)
5. Return configured client

---

### ReverseGeocode(ctx context.Context, req GeocoderRequest)

**Input:**
- `ctx` - Context for cancellation/timeout
- `req` - Coordinates to geocode

**Output:** `(*types.Location, error)`

**Behavior:**

```
ReverseGeocode(ctx, req)
   │
   ├──► Build URL with query parameters
   │     • lat, lon, format=json, zoom=10, addressdetails=1
   │     • Add api_key if configured
   │
   ├──► Create HTTP request with context
   │     • Set User-Agent header
   │     • Set Accept-Language: en
   │
   ├──► RetryWithBackoff(func() error {
   │         │
   │         ├── Execute HTTP request
   │         │
   │         ├── Check status code
   │         │     ├── 200 → Parse JSON → return nil
   │         │     ├── 400 → return GeocodeError
   │         │     ├── 403 → return GeocodeError (User-Agent issue)
   │         │     ├── 429 → return GeocodeError wrapping ErrRateLimited
   │         │     └── 5xx → return GeocodeError (retryable)
   │         │
   │         └── Network error → return wrapped error
   │    })
   │
   ├──► Check for "error" field in response
   │     • Nominatim returns 200 with error for invalid coords
   │
   └──► Return Location or error
```

---

## Error Handling

### HTTP Status Code Mapping

```go
switch resp.StatusCode {
case http.StatusOK:
    // Parse response, but check for error field
case http.StatusBadRequest:
    return &GeocodeError{
        StatusCode: 400,
        Query:      fmt.Sprintf("%.4f, %.4f", req.Latitude, req.Longitude),
        Message:    "Invalid coordinates",
    }
case http.StatusForbidden:
    return &GeocodeError{
        StatusCode: 403,
        Query:      fmt.Sprintf("%.4f, %.4f", req.Latitude, req.Longitude),
        Message:    "Request blocked - check User-Agent header",
    }
case http.StatusTooManyRequests:
    return &GeocodeError{
        StatusCode: 429,
        Query:      fmt.Sprintf("%.4f, %.4f", req.Latitude, req.Longitude),
        Message:    "Rate limit exceeded - max 1 request/second",
    }
default:
    // Handle 5xx
}
```

### Special Case: 200 with Error

Nominatim may return HTTP 200 but include an error in the body:

```json
{
  "error": "Unable to geocode"
}
```

Check for this:
```go
if apiResp.Error != "" {
    return &GeocodeError{
        StatusCode: 200,
        Query:      fmt.Sprintf("%.4f, %.4f", req.Latitude, req.Longitude),
        Message:    apiResp.Error,
    }
}
```

---

## Response Parsing

### Target Data Structure (from types package)

```go
// In types/types.go
type Location struct {
    DisplayName string  // Full formatted name
    City        string  // City name
    State       string  // State/region
    Country     string  // Country name
    CountryCode string  // ISO country code
    Latitude    float64 // Actual coordinates returned
    Longitude   float64
}
```

### Parsing Logic

1. Decode JSON response
2. Check for error field
3. Extract address components
4. Parse lat/lon strings to float64
5. Build Location struct

**Handling missing fields:** Not all locations have all address components. Use empty strings for missing fields.

---

## Useful Packages

| Package | Why Useful |
|---------|------------|
| `net/http` | HTTP client, request building, setting headers |
| `net/url` | Query string building |
| `encoding/json` | JSON parsing |
| `context` | Request cancellation |
| `strconv` | `strconv.ParseFloat()` for lat/lon strings |
| `fmt` | Error formatting, coordinate formatting |

---

## Test Cases

| Scenario | Input | Expected |
|----------|-------|----------|
| Valid coordinates | Berlin (52.52, 13.41) | Location with "Berlin", "Germany" |
| Ocean coordinates | (0, 0) | GeocodeError (no land) |
| HTTP 429 | Rapid requests | GeocodeError with ErrRateLimited |
| HTTP 403 | Missing User-Agent | GeocodeError with appropriate message |
| Timeout | Slow server | Error wrapping ErrTimeout |
| 200 with error field | Invalid coords | GeocodeError from response |
| Missing city field | Rural location | Location with empty City, valid Country |
| Context canceled | Canceled context | Context error |

### Testing Tips

```go
func TestReverseGeocode_Success(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Verify User-Agent header is set
        if r.Header.Get("User-Agent") == "" {
            t.Error("User-Agent header not set")
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(mockNominatimResponse)
    }))
    defer server.Close()

    // Test with mock server
}
```

---

## Implementation Hints

1. **User-Agent is mandatory:** Nominatim will block requests without it:
   ```go
   req.Header.Set("User-Agent", "WeatherCLI/1.0 (contact@example.com)")
   ```

2. **Coordinate formatting:** Format coordinates with limited precision:
   ```go
   params.Set("lat", fmt.Sprintf("%.6f", req.Latitude))
   params.Set("lon", fmt.Sprintf("%.6f", req.Longitude))
   ```

3. **Zoom level:** Use zoom=10 for city-level detail. Higher values give more specific results (street level), lower values give country/continent.

4. **Parse lat/lon from strings:** Nominatim returns coordinates as strings:
   ```go
   type apiResponse struct {
       Lat string `json:"lat"`
       Lon string `json:"lon"`
       // ...
   }

   lat, _ := strconv.ParseFloat(apiResp.Lat, 64)
   lon, _ := strconv.ParseFloat(apiResp.Lon, 64)
   ```

5. **Address fallbacks:** Some fields may be missing. Handle gracefully:
   ```go
   type addressDetail struct {
       City        string `json:"city"`
       Town        string `json:"town"`
       Village     string `json:"village"`
       Country     string `json:"country"`
       CountryCode string `json:"country_code"`
   }

   // City might be empty, but town or village might exist
   city := addr.City
   if city == "" {
       city = addr.Town
   }
   if city == "" {
       city = addr.Village
   }
   ```

6. **Rate limit awareness:** For the free tier, add a small delay if you expect multiple calls:
   ```go
   const nominatimRateLimit = time.Second // 1 request per second
   ```

7. **API key handling:** If API key is provided, add it as a query parameter:
   ```go
   if c.apiKey != "" {
       params.Set("key", c.apiKey)
   }
   ```

---

## Example Output

For coordinates (52.52, 13.41):

```go
Location{
    DisplayName: "Berlin, Germany",
    City:        "Berlin",
    State:       "Berlin",
    Country:     "Germany",
    CountryCode: "de",
    Latitude:    52.5170365,
    Longitude:   13.3888599,
}
```
