package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/go-envparse"
	apiErrors "github.com/warlock016/weather_cli/errors"
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
		return nil, fmt.Errorf("%w: GEOCODE_API_URL", apiErrors.ErrMissingConfig)
	} else if geocode == "" {
		return nil, fmt.Errorf("%w: GEOCODE_API_KEY", apiErrors.ErrMissingConfig)
	} else if meteourl == "" {
		return nil, fmt.Errorf("%w: OPENMETEO_URL", apiErrors.ErrMissingConfig)
	}

	// Set variables to Struct
	result.GeocodeURL = geourl
	result.GeocodeAPIKey = geocode
	result.OpenMeteoURL = meteourl

	return result, nil
}
func LoadFromEnvFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("failed to find file: %w: %w", apiErrors.ErrMissingConfig, err)
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open .env file: %w: %w", apiErrors.ErrMissingConfig, err)
	}
	defer file.Close()

	res, err := envparse.Parse(file)
	if err != nil {
		return fmt.Errorf("failed to parse .env file: %w: %w", apiErrors.ErrMissingConfig, err)
	}

	for k, v := range res {
		err := os.Setenv(k, v)
		if err != nil {
			return fmt.Errorf("failed to set .env key: %s, reason: %w: %w", k, apiErrors.ErrMissingConfig, err)
		}
	}

	return nil
}
