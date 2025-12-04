package weather_api_errors

import (
	"strings"
	"testing"
)

func TestWeatherAPIError(t *testing.T) {
	err := &WeatherAPIError{
		StatusCode: 404,
		Message:    "City not found",
		City:       "InvalidCity",
	}

	errStr := err.Error()
	if errStr == "" {
		t.Error("WeatherAPIError.Error() returned empty string")
	}
	if !strings.Contains(errStr, "InvalidCity") {
		t.Errorf("WeatherAPIError.Error() should contain city name, got %q", errStr)
	}
}

func TestGeocodeError(t *testing.T) {
	err := &GeocodeError{
		City:   "UnknownPlace",
		Reason: "No results found",
	}

	errStr := err.Error()
	if errStr == "" {
		t.Error("GeocodeError.Error() returned empty string")
	}
	if !strings.Contains(errStr, "UnknownPlace") {
		t.Errorf("GeocodeError.Error() should contain city name, got %q", errStr)
	}
}

// TODO(human): After refactoring your Weather CLI, add integration tests here
// that test the error handling in real scenarios:
// - Invalid city names
// - Network failures (mock the HTTP client)
// - API rate limits
// - Invalid API responses
