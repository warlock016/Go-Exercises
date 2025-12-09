package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/go-envparse"
)

type Config struct {
	GeocodeURL    string
	GeocodeAPIKey string

	// no api key required
	OpenMeteoURL string
	// OpenMeteoAPIKey string
}

func Load() (*Config, error) {
	result := &Config{}

	if err := LoadFromEnvFile(".env"); err != nil {
		return nil, fmt.Errorf("missing .env file in path: %w", err)
	}

	geourl := strings.TrimSpace(os.Getenv("GEOCODE_API_URL"))
	geocode := strings.TrimSpace(os.Getenv("GEOCODE_API_KEY"))
	meteourl := strings.TrimSpace(os.Getenv("OPENMETEO_API_URL"))

	if geourl == "" {
		return nil, fmt.Errorf("missing Geocode API URL in .env")
	} else if geocode == "" {
		return nil, fmt.Errorf("missing Geocode API Key in .env")
	} else if meteourl == "" {
		return nil, fmt.Errorf("missing OpenMeteo API URL in .env")
	}

	// Set variables to Struct
	result.GeocodeURL = geourl
	result.GeocodeAPIKey = geocode
	result.OpenMeteoURL = meteourl

	return result, nil
}
func LoadFromEnvFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("failed to find file: %w", err)
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open .env file: %w", err)
	}
	defer file.Close()

	res, err := envparse.Parse(file)
	if err != nil {
		return fmt.Errorf("failed to parse .env file: %w", err)
	}

	for k, v := range res {
		err := os.Setenv(k, v)
		if err != nil {
			return fmt.Errorf("failed to set .env key: %s, reason: %w", k, err)
		}
	}

	return nil
}
