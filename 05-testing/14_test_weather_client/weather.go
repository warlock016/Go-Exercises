package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// WeatherData represents weather information
type WeatherData struct {
	City        string  `json:"name"`
	Temp        float64 `json:"temp"`
	Description string  `json:"description"`
	Humidity    int     `json:"humidity"`
}

// WeatherClient fetches weather data
type WeatherClient struct {
	BaseURL string
	APIKey  string
}

// GetWeather fetches weather for given coordinates
func (c *WeatherClient) GetWeather(lat, lon float64) (*WeatherData, error) {
	if c.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	url := fmt.Sprintf("%s/weather?lat=%.2f&lon=%.2f&appid=%s",
		c.BaseURL, lat, lon, c.APIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: status %d", resp.StatusCode)
	}

	var data WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return &data, nil
}

// GeoData represents geocoding result
type GeoData struct {
	City string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

// GeoClient handles geocoding
type GeoClient struct {
	BaseURL string
	APIKey  string
}

// Geocode converts city name to coordinates
func (g *GeoClient) Geocode(city string) (*GeoData, error) {
	if city == "" {
		return nil, fmt.Errorf("city name is required")
	}

	url := fmt.Sprintf("%s/geo?q=%s&appid=%s",
		g.BaseURL, url.QueryEscape(city), g.APIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("geocoding failed: status %d", resp.StatusCode)
	}

	var results []GeoData
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("city not found: %s", city)
	}

	return &results[0], nil
}

// FormatWeather formats weather data for display
func FormatWeather(data *WeatherData) string {
	if data == nil {
		return "No weather data available"
	}

	return fmt.Sprintf("Weather in %s\nTemperature: %.1f°C\nConditions: %s\nHumidity: %d%%",
		data.City, data.Temp, data.Description, data.Humidity)
}
