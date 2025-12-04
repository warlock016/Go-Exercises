package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Weather struct {
	City string  `json:"city"`
	Temp float64 `json:"temp"`
	Desc string  `json:"description"`
}

type WeatherClient struct {
	BaseURL string
}

func (c *WeatherClient) GetWeather(city string) (*Weather, error) {
	url := fmt.Sprintf("%s/weather?city=%s", c.BaseURL, url.QueryEscape(city))

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var weather Weather
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &weather, nil
}
