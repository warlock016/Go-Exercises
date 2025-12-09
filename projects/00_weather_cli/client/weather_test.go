package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/warlock016/weather_cli/client"
)

const validWeatherResponse = `{
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
}`

func TestFetchWeather(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("lat") == "" {
			t.Error("missing latitude parameter")
		}
		if r.URL.Query().Get("lon") == "" {
			t.Error("missing longitude parameter")
		}
		if r.URL.Query().Get("start_date") == "" {
			t.Error("missing start parameter")
		}
		if r.URL.Query().Get("end_date") == "" {
			t.Error("missing end parameter")
		}
		if r.URL.Query().Get("hourly") == "" {
			t.Error("missing variable parameter")
		}
		// Return fake weather data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(validWeatherResponse)) // mock response
	}))
	defer server.Close()

	weatherClient := client.NewWeatherClient(server.URL, 3*time.Second)
	ctx := context.Background()
	req := client.WeatherRequest{
		Latitude:  52.52,
		Longitude: 13.41,
		Timezone:  "GMT",
		StartDate: time.Date(2025, 11, 17, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2025, 11, 17, 0, 0, 0, 0, time.UTC),
		Variables: []string{"temperature_2m"},
	}

	data, err := weatherClient.FetchWeather(ctx, req)
	if err != nil {
		t.Fatalf("FetchWeather failed: %v", err)
	}
	if data == nil {
		t.Fatalf("Empty WeatherData")
	}

	if _, exists := data.Hourly["time"]; !exists {
		t.Fatalf("missing timestamps in response %v", data.Hourly)
	}
	if len(data.Hourly["time"]) != 24 {
		t.Errorf("Invalid result array")
	}

	for _, v := range req.Variables {
		slice, exists := data.Hourly[v]
		if !exists {
			t.Fatalf("missing variable %s in response", v)
		}
		if len(slice) == 0 {
			t.Fatalf("no time-series data for %s in response. Got %v", v, slice)
		}

		if len(data.Hourly["time"]) != len(slice) {
			t.Errorf("variable and timestamp array do not match %d: %d", len(data.Hourly["time"]), len(slice))
		}

		for idx, val := range slice {
			if _, ok := val.(float64); !ok {
				t.Errorf("%s[%d]: expected float64, got %T (value: %v)", v, idx, val, val)
			}
		}

		if _, exists := data.HourlyUnits[v]; !exists {
			t.Fatalf("missing variable %s in response", v)
		} else {
			trimmed := strings.TrimSpace(data.HourlyUnits[v])
			if trimmed == "" {
				t.Errorf("Invalid dp %v for label %s. Trimmed: %s", data.HourlyUnits[v], v, trimmed)
			}
		}
	}
}

func TestFetchWeather_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// return server response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(strconv.Itoa(http.StatusInternalServerError))) // mock response
	}))
	defer server.Close()

	weatherClient := client.NewWeatherClient(server.URL, 3*time.Second)
	ctx := context.Background()
	req := client.WeatherRequest{
		Latitude:  52.51,
		Longitude: 13.41,
		StartDate: time.Date(2026, 12, 30, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
		Timezone:  "GMT",
		Variables: []string{"temperature_2m", "wind_speed_10m"},
	}

	resp, err := weatherClient.FetchWeather(ctx, req)
	if err == nil {
		t.Errorf("expected error for 500 response, got nil")
	}

	if resp != nil {
		t.Errorf("expected nil response for error case, got %+v", resp)
	}

}

func TestFetchWeather_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusRequestTimeout)
	}))
	defer server.Close()

	weatherClient := client.NewWeatherClient(server.URL, time.Millisecond*1000)
	ctx := context.Background()
	req := client.WeatherRequest{
		Latitude:  52.51,
		Longitude: 13.41,
		StartDate: time.Date(2026, 12, 30, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
		Timezone:  "GMT",
		Variables: []string{"temperature_2m", "wind_speed_10m"},
	}

	resp, err := weatherClient.FetchWeather(ctx, req)
	if err == nil {
		t.Error("expected error for 408 response, got nil")
	}
	if resp != nil {
		t.Errorf("unexpected nil response, got %+v", resp)
	}
}

func TestFetchWeather_ClientTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second) // this server is slow to reply!
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(validWeatherResponse))
	}))

	defer server.Close()

	weatherClient := client.NewWeatherClient(server.URL, time.Millisecond*100) // fast client timeout -> do not wait too long for slow server!
	ctx := context.Background()
	req := client.WeatherRequest{
		Latitude:  52.51,
		Longitude: 13.41,
		StartDate: time.Date(2026, 12, 30, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
		Timezone:  "GMT",
		Variables: []string{"temperature_2m", "wind_speed_10m"},
	}

	resp, err := weatherClient.FetchWeather(ctx, req)
	if err == nil {
		t.Error("unexpected nil error, wanted client timeout")
	}
	if resp != nil {
		t.Errorf("got != nil response, %+v", resp)
	}
}

// Testing if Server says OK but sends garbage data
func TestFetchWeather_MalformedJSON(t *testing.T) {
	tests := []struct {
		Name     string
		Body     string
		wantErr  bool
		wantResp bool
	}{
		{"Syntax error", "{[0,1,2}", true, false},
		{"HTML error page", "<html>Error</html>", true, false},
		{"Empty response body", "", true, false},
		{"Truncated JSON", "{", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.Body))
			}))

			defer server.Close()

			weatherClient := client.NewWeatherClient(server.URL, time.Millisecond*500)
			ctx := context.Background()
			req := client.WeatherRequest{
				Latitude:  52.51,
				Longitude: 13.41,
				StartDate: time.Date(2026, 12, 30, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
				Timezone:  "GMT",
				Variables: []string{"temperature_2m", "wind_speed_10m"},
			}

			resp, err := weatherClient.FetchWeather(ctx, req)

			if err == nil {
				t.Error("unexpected unmarshalling, expected error")
			}

			if resp != nil {
				t.Errorf("expected nil data, got %+v", resp)
			}
		})
	}
}
