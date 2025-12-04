# Exercise 14: Weather API Errors

## 🎯 Learning Goal
Apply error handling patterns learned in this module to refactor the Weather CLI project with proper error wrapping, logging, and user-friendly error messages.

## 📝 Problem Description

This is a capstone exercise that brings together all error handling concepts:
- Wrap errors from the weather API
- Create custom error types for API failures
- Log errors with context
- Provide helpful error messages to users

## 📋 Instructions

1. Review your Weather CLI project in `/projects/00_weather_cli/`
2. Identify all error handling locations
3. Apply patterns from exercises 01-13:
   - Wrap errors with context
   - Create custom error types for API errors
   - Map HTTP status codes to user-friendly messages
   - Add structured error logging
4. Test error scenarios:
   - Invalid city names
   - Network failures
   - API rate limits
   - Invalid API keys

## 🔧 Suggested Improvements

```go
// Custom error types
type WeatherAPIError struct {
    StatusCode int
    Message    string
    City       string
}

type GeocodeError struct {
    City   string
    Reason string
}

// Error wrapping
func FetchWeather(city string) (*WeatherData, error) {
    data, err := callAPI(city)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch weather for %s: %w", city, err)
    }
    return data, nil
}

// User-friendly error messages
func HandleWeatherError(err error) {
    if errors.Is(err, ErrCityNotFound) {
        fmt.Println("City not found. Please check the spelling and try again.")
        return
    }
    // ... handle other errors
}
```

## 🎓 What This Teaches

- **Real-world application** - Applying error patterns to production code
- **Error design** - Thinking about error UX
- **Refactoring** - Improving existing error handling
- **Integration** - Combining multiple error patterns

---

**Congratulations!** You've completed Module 04: Error Handling. You now understand Go's error handling philosophy and can build robust, production-ready error handling systems.
