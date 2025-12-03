package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

/*
{
"latitude":52.54833,
"longitude":13.407822,
"generationtime_ms":0.22292137145996094,
"utc_offset_seconds":0,
"timezone":"GMT",
"timezone_abbreviation":"GMT",
"elevation":38.0,
"hourly_units":{
	"time":"iso8601",
	"temperature_2m":"°C"
	},
"hourly":{
	"time":["2025-11-17T00:00","2025-11-17T01:00","2025-11-17T02:00","2025-11-17T03:00","2025-11-17T04:00","2025-11-17T05:00","2025-11-17T06:00","2025-11-17T07:00","2025-11-17T08:00","2025-11-17T09:00","2025-11-17T10:00","2025-11-17T11:00","2025-11-17T12:00","2025-11-17T13:00","2025-11-17T14:00","2025-11-17T15:00","2025-11-17T16:00","2025-11-17T17:00","2025-11-17T18:00","2025-11-17T19:00","2025-11-17T20:00","2025-11-17T21:00","2025-11-17T22:00","2025-11-17T23:00"],
	"temperature_2m":[3.8,3.4,2.9,2.8,2.5,2.4,3.1,3.5,3.3,3.7,4.4,4.8,5.5,5.6,5.3,2.8,2.6,2.5,2.1,1.7,1.2,0.9,0.7,0.5]
	}
}
*/

type OpenMeteoResponse struct {
	Latitude    float64        `json:"latitude"`
	Longitude   float64        `json:"longitude"`
	Timezone    string         `json:"timezone"`           // GMT, Europe/Berlin, etc...
	Offset      float64        `json:"utc_offset_seconds"` // 3600 == +1; 7200 == +2; etc...
	Elevation   float64        `json:"elevation"`
	HourlyUnits map[string]any `json:"hourly_units"`
	Hourly      map[string]any `json:"hourly"`
}

func FetchData(lat, long, start, end, hourly, tz string) (OpenMeteoResponse, error) {
	base, _ := url.Parse("https://archive-api.open-meteo.com/v1/archive")
	q := base.Query()
	q.Set("latitude", lat)
	q.Set("longitude", long)
	q.Set("start_date", start)
	q.Set("end_date", end)
	q.Set("hourly", hourly)

	if len(tz) != 0 {
		q.Set("timezone", tz)
	}

	base.RawQuery = q.Encode()

	resp, err := http.Get(base.String())

	result := OpenMeteoResponse{}

	if err != nil {
		return result, fmt.Errorf("Failed to query data: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return result, fmt.Errorf("HTTP error: %d: %s", resp.StatusCode, string(b))
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return result, fmt.Errorf("IO read error: %v", err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return result, fmt.Errorf("JSON unmarshal error: %v", err)
	}

	return result, nil
}
