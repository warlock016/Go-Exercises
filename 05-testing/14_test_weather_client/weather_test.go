package weather

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO(human): Test WeatherClient.GetWeather with successful response
// Use httptest.NewServer to mock the API
// Return mock WeatherData as JSON
// Verify client parses it correctly

// TODO(human): Test WeatherClient.GetWeather with missing API key
// Create client with empty APIKey
// Should return error before making request

// TODO(human): Test WeatherClient.GetWeather with 404 error
// Mock server returns 404
// Verify client returns error with status code

// TODO(human): Test WeatherClient.GetWeather with invalid JSON
// Mock server returns invalid JSON
// Verify client returns parsing error

// TODO(human): Test GeoClient.Geocode with successful response
// Mock server returns []GeoData with one result
// Verify client returns first result

// TODO(human): Test GeoClient.Geocode with empty city
// Should return error without making request

// TODO(human): Test GeoClient.Geocode with empty results
// Mock server returns empty array []
// Should return "city not found" error

// TODO(human): Test FormatWeather with valid data
// Create WeatherData with test values
// Verify output format is correct

// TODO(human): Test FormatWeather with nil data
// Should return "No weather data available"

// Tips:
// - Use table-driven tests with subtests
// - Remember to defer server.Close()
// - Check both error AND result in each test
// - Run with: go test -v -cover
// - Aim for >85% coverage
