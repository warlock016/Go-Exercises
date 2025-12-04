# Exercise 10: Mocking HTTP Clients (🌐 HTTP)

**Learning Goal:** Use httptest.Server to mock external APIs and test HTTP clients.

---

## Problem Description

When testing code that calls external APIs, you don't want to make real network calls. Use `httptest.Server` to create a local test server that mimics the external API.

---

## Client to Test

```go
type WeatherClient struct {
    BaseURL string
}

func (c *WeatherClient) GetWeather(city string) (*Weather, error)
```

---

## Testing with httptest.Server

```go
func TestGetWeather(t *testing.T) {
    // Create mock server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(Weather{City: "London", Temp: 15.5})
    }))
    defer server.Close()

    // Create client pointing to mock server
    client := &WeatherClient{BaseURL: server.URL}

    weather, err := client.GetWeather("London")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if weather.City != "London" {
        t.Errorf("city = %q, want %q", weather.City, "London")
    }
}
```

---

## Your Task

Test the WeatherClient by mocking the API with httptest.Server.

**Next Exercise:** `11_test_coverage` - Measure and improve test coverage
