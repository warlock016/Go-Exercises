package types

type WeatherData struct {
	Latitude    float64           `json:"latitude"`
	Longitude   float64           `json:"longitude"`
	Elevation   float64           `json:"elevation"`
	Timezone    string            `json:"timezone"`
	Offset      int               `json:"utc_offset_seconds"`
	Variables   []string          // ordered set of variables for ordering data
	HourlyUnits map[string]string `json:"hourly_units"`
	Hourly      map[string][]any  `json:"hourly"`
}

type GeoData struct {
	DisplayName string  `json:"display_name"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
