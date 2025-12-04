package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO(human): Test GetWeather with successful response
// 1. Create httptest.Server that returns mock weather data
// 2. Create WeatherClient with server.URL as BaseURL
// 3. Call GetWeather
// 4. Verify weather data matches mock response
// 5. Don't forget: defer server.Close()
//
// Example structure:
// func TestGetWeather(t *testing.T) {
//     server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Check request URL contains city parameter
//         city := r.URL.Query().Get("city")
//         if city != "London" {
//             t.Errorf("city = %q, want London", city)
//         }
//
//         // Return mock response
//         w.Header().Set("Content-Type", "application/json")
//         json.NewEncoder(w).Encode(Weather{
//             City: "London",
//             Temp: 15.5,
//             Desc: "Cloudy",
//         })
//     }))
//     defer server.Close()
//
//     client := &WeatherClient{BaseURL: server.URL}
//     weather, err := client.GetWeather("London")
//     // ... test assertions
// }

// TODO(human): Test GetWeather with 404 error
// Mock server should return 404 status
// Verify client returns error

// TODO(human): Test GetWeather with invalid JSON
// Mock server returns invalid JSON
// Verify client returns error
