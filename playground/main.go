package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type OpenMeteoResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Hourly    struct {
		Time          []string  `json:"time"`
		Temperature2m []float64 `json:"temperature_2m"`
	} `json:"hourly"`
}

func main() {

	base, _ := url.Parse("https://archive-api.open-meteo.com/v1/archive")
	q := base.Query()
	q.Set("latitude", "52.51")
	q.Set("longitude", "13.41")
	q.Set("start_date", "2025-11-01")
	q.Set("end_date", "2025-11-01")
	q.Set("hourly", "temperature_2m")
	q.Set("timezone", "GMT")
	base.RawQuery = q.Encode()

	response, err := http.Get(base.String())
	if err != nil {
		log.Fatal("Failed to query data", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(response.Body)
		log.Fatalf("HTTP error %d: %s", response.StatusCode, string(b))
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("IO read error", err)
	}

	var result OpenMeteoResponse

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal("JSON unmarshal error: ", err)
	}

	if len(result.Hourly.Temperature2m) != len(result.Hourly.Time) {
		log.Fatal("timestamps and values not matching")
	}

}
