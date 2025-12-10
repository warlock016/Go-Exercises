package formatter_test

import (
	"encoding/json"
	"testing"

	"github.com/warlock016/weather_cli/formatter"
	"github.com/warlock016/weather_cli/types"
)

func mockWeatherData(t *testing.T) *types.WeatherData {
	t.Helper()
	return &types.WeatherData{
		Latitude:  52.52,
		Longitude: 13.41,
		Elevation: 100,
		Timezone:  "Europe/Berlin",
		HourlyUnits: map[string]string{
			"time":           "iso8601",
			"temperature_2m": "°C",
			"wind_speed_10m": "km/h",
		},
		Hourly: map[string][]any{
			"time":           {"2024-01-15T00:00", "2024-01-15T01:00", "2024-01-15T02:00"},
			"temperature":    {5.2, 4.8, 4.5},
			"wind_speed_10m": {2.0, 1.3, 0.8},
		},
	}
}

func mockGeoData(t *testing.T) *types.GeoData {
	t.Helper()
	return &types.GeoData{
		DisplayName: "Berlin, Germany",
		City:        "Berlin",
		Country:     "Germany",
		CountryCode: "DE",
		Latitude:    52.52,
		Longitude:   13.41,
	}
}

func TestFormat_NilInputs(t *testing.T) {
	tests := []struct {
		name     string
		weather  *types.WeatherData
		location *types.GeoData
		wantErr  bool
	}{
		{"nil weather", nil, mockGeoData(t), true},
		{"nil location", mockWeatherData(t), nil, true},
		{"valid location", mockWeatherData(t), mockGeoData(t), false},
		{"both nil", nil, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := formatter.FormatterInput{
				Weather:  tt.weather,
				Location: tt.location,
				Options:  formatter.FormatOptions{Format: "table"},
			}

			result, err := formatter.Format(input)

			if tt.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error, got %v, want nil", err)
			}

			if !tt.wantErr && result.Content == "" {
				t.Error("unexpected empty string")
			}
		})
	}
}

func TestFormat_InvalidFormat(t *testing.T) {
	input := formatter.FormatterInput{
		Weather:  mockWeatherData(t),
		Location: mockGeoData(t),
		Options: formatter.FormatOptions{
			Format: "xml",
		},
	}

	res, err := formatter.Format(input)
	if err == nil {
		t.Fatalf("expected error for invalid format %s", input.Options.Format)
	}

	if res != nil {
		t.Fatalf("received unexpected !nil value %v", res)
	}
}

func TestFormatTable_Structure(t *testing.T) {
	input := formatter.FormatterInput{
		Weather:  mockWeatherData(t),
		Location: mockGeoData(t),
		Options: formatter.FormatOptions{
			Format:   "table",
			Timezone: "Europe/Berlin",
		},
	}

	res, err := formatter.Format(input)
	if err != nil {
		t.Fatalf("unexpected error during conversion: %+v", err)
	}
	if res == nil {
		t.Fatal("unexpected nil response")
	}
}

func TestFormatJSON_ValidOutput(t *testing.T) {
	// t.Skip()
	input := formatter.FormatterInput{
		Weather:  mockWeatherData(t),
		Location: mockGeoData(t),
		Options: formatter.FormatOptions{
			Format:   "json",
			Timezone: "Europe/Berlin",
		},
	}

	res, err := formatter.Format(input)
	if err != nil {
		t.Fatalf("unexpected error during conversion: %+v", err)
	}
	if res == nil {
		t.Fatal("unexpected nil response")
	}

	if !json.Valid([]byte(res.Content)) {
		t.Error("FormatJSON returned invalid JSON")
	}
}

func TestFormat_DataPointsCount(t *testing.T) {

	input := formatter.FormatterInput{
		Weather:  mockWeatherData(t),
		Location: mockGeoData(t),
		Options: formatter.FormatOptions{
			Format:   "table",
			Timezone: "Europe/Berlin",
		},
	}

	res, err := formatter.Format(input)
	if err != nil {
		t.Fatalf("unexpected error %+v", err)
	}

	if res == nil {
		t.Fatal("unexpected nil result, with nil error")
	}

	if res.DataPoints != len(input.Weather.Hourly["time"]) {
		t.Errorf("unexpected datapoint count: got %d, want %d", res.DataPoints, 3)
	}
}
