# Exercise 14: Test Weather Client (🌐 HTTP)

**Learning Goal:** Apply all testing skills to comprehensively test the Weather CLI client - the same code from your project!

---

## Problem Description

This is a real-world exercise using code from your Weather CLI project (`projects/00_weather_cli`). You'll write comprehensive tests for the weather client, geocoding client, and formatter.

This exercise combines:
- HTTP client testing with httptest.Server
- JSON API testing
- Error handling testing
- Table-driven tests
- Test coverage analysis

---

## Components to Test

```go
// Weather API client
type WeatherClient struct { ... }
func (c *WeatherClient) GetWeather(lat, lon float64) (*WeatherData, error)

// Geocoding client
type GeoClient struct { ... }
func (g *GeoClient) Geocode(city string) (*GeoData, error)

// Formatter
func FormatWeather(data *WeatherData) string
```

---

## Your Task

Write comprehensive tests for:

1. **WeatherClient.GetWeather**
   - Successful API response
   - HTTP errors (404, 500)
   - Invalid JSON response
   - Network errors

2. **GeoClient.Geocode**
   - Successful geocoding
   - City not found
   - Multiple results
   - Invalid response

3. **FormatWeather**
   - Different weather conditions
   - Temperature formatting
   - Units (Celsius/Fahrenheit)

---

## Testing Strategy

### Mock the External APIs

```go
func TestWeatherClient(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check API key is sent
        apiKey := r.URL.Query().Get("appid")
        if apiKey == "" {
            http.Error(w, "Missing API key", http.StatusUnauthorized)
            return
        }

        // Return mock weather data
        json.NewEncoder(w).Encode(WeatherData{
            Temp: 15.5,
            Description: "Cloudy",
            City: "London",
        })
    }))
    defer server.Close()

    client := &WeatherClient{
        BaseURL: server.URL,
        APIKey:  "test-key",
    }

    // Test...
}
```

---

## Test Checklist

**WeatherClient:**
- [ ] Successful weather fetch
- [ ] Missing API key error
- [ ] Invalid coordinates
- [ ] 404 Not Found
- [ ] 500 Server Error
- [ ] Invalid JSON response
- [ ] Network timeout (use context)

**GeoClient:**
- [ ] Successful geocoding
- [ ] City not found
- [ ] Empty city name
- [ ] Invalid API response
- [ ] Multiple results (if supported)

**FormatWeather:**
- [ ] Sunny weather
- [ ] Rainy weather
- [ ] Different temperatures (negative, zero, positive)
- [ ] Temperature unit conversion
- [ ] Nil data handling

---

## Running Tests

```bash
# Run all tests
go test -v

# With coverage
go test -cover

# Generate coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run specific test
go test -v -run TestWeatherClient
```

---

## Instructions

1. Study the weather client code structure
2. Create simplified versions of the types for testing
3. Write tests following the checklist
4. Use table-driven tests with subtests
5. Mock all HTTP calls with httptest.Server
6. Aim for >85% coverage

---

## Success Criteria

- [ ] All tests pass
- [ ] Coverage >85%
- [ ] Tests run fast (<100ms total)
- [ ] No network calls to real APIs
- [ ] Clear test names and error messages

---

## Think About

1. How would you test rate limiting?
2. How would you test retry logic?
3. Should you test with real API keys? Why or why not?
4. What happens if the API changes its response format?

---

**Next Exercise:** `15_test_driven_feature` - Build a feature using TDD
