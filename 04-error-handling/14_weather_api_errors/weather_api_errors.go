package weather_api_errors

// This exercise is about refactoring your Weather CLI project
// located in /Users/mode/Documents/Code/Go Exercises/projects/00_weather_cli/

// TODO(human): Review the Weather CLI code and apply error handling patterns:
// 1. Create custom error types for weather API errors
// 2. Wrap errors with context using fmt.Errorf with %w
// 3. Add structured error logging
// 4. Provide user-friendly error messages
// 5. Handle different error scenarios (network, API, invalid input)

// Example custom error types you might create:

// WeatherAPIError represents an error from the weather API
type WeatherAPIError struct {
	StatusCode int
	Message    string
	City       string
}

func (e *WeatherAPIError) Error() string {
	// TODO(human): Implement
	return ""
}

// GeocodeError represents a geocoding error
type GeocodeError struct {
	City   string
	Reason string
}

func (e *GeocodeError) Error() string {
	// TODO(human): Implement
	return ""
}

// TODO(human): Apply these patterns to your Weather CLI project
