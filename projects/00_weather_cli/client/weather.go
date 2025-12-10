package client

import (
	"context"
	"encoding/json"
	"fmt"

	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	apiErrors "github.com/warlock016/weather_cli/errors"
	"github.com/warlock016/weather_cli/types"
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

type WeatherClient struct {
	BaseURL    string
	httpClient *http.Client
}

type WeatherRequest struct {
	Latitude  float64
	Longitude float64
	Timezone  string
	StartDate time.Time
	EndDate   time.Time
	Variables []string
}

func NewWeatherClient(baseURL string, timeout time.Duration) *WeatherClient {
	return &WeatherClient{
		BaseURL:    baseURL,
		httpClient: &http.Client{Timeout: timeout},
	}
}

func (c *WeatherClient) FetchWeather(ctx context.Context, req WeatherRequest) (*types.WeatherData, error) {

	if ctx == nil {
		return nil, errors.New("nil context")
	}

	base, _ := url.Parse(c.BaseURL)
	q := base.Query()
	q.Set("latitude", strconv.FormatFloat(req.Latitude, 'f', 9, 64))
	q.Set("longitude", strconv.FormatFloat(req.Longitude, 'f', 9, 64))
	q.Set("start_date", req.StartDate.Format("2006-01-02"))
	q.Set("end_date", req.EndDate.Format("2006-01-02"))
	q.Set("hourly", strings.Join(req.Variables, ","))
	if req.Timezone != "" {
		q.Set("timezone", req.Timezone)
	}
	base.RawQuery = q.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), nil)
	if err != nil {
		return nil, &apiErrors.WeatherAPIError{
			StatusCode: 0,
			Message:    "failed to create request",
			Endpoint:   base.String(),
			Latitude:   req.Latitude,
			Longitude:  req.Longitude,
			ErrorType:  fmt.Errorf("%w: %v", apiErrors.ErrNetwork, err),
		}
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, &apiErrors.WeatherAPIError{
			StatusCode: 0,
			Message:    "request failed",
			Endpoint:   base.String(),
			Latitude:   req.Latitude,
			Longitude:  req.Longitude,
			ErrorType:  fmt.Errorf("%w: %v", apiErrors.ErrNetwork, err),
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		switch resp.StatusCode {

		// Client-Side Errors
		case 401, 403:
			err = fmt.Errorf("%w: HTTP %d from %s", apiErrors.ErrUnauthorized, resp.StatusCode, base.String())
		case 404:
			err = fmt.Errorf("%w: HTTP %d from %s", apiErrors.ErrNotFound, resp.StatusCode, base.String())
		case 400, 422, 405, 409, 415:
			err = fmt.Errorf("%w: HTTP %d from %s", apiErrors.ErrInvalidInput, resp.StatusCode, base.String())
		case 429:
			err = fmt.Errorf("%w: HTTP %d from %s", apiErrors.ErrRateLimited, resp.StatusCode, base.String())

		// Server-Side Errors
		case 500, 502, 503, 504:
			err = fmt.Errorf("%w: HTTP %d from %s", apiErrors.ErrNetwork, resp.StatusCode, base.String())
		case 501:
			err = fmt.Errorf("%w: HTTP %d from %s", apiErrors.ErrNotFound, resp.StatusCode, base.String())
		default:
			err = fmt.Errorf("%w: HTTP %d from %s", apiErrors.ErrNetwork, resp.StatusCode, base.String())
		}

		return nil, &apiErrors.WeatherAPIError{
			StatusCode: resp.StatusCode,
			Message:    string(b),
			Endpoint:   base.String(),
			Latitude:   req.Latitude,
			Longitude:  req.Longitude,
			ErrorType:  err,
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &apiErrors.WeatherAPIError{
			StatusCode: resp.StatusCode,
			Message:    "parsed invalid response body",
			Endpoint:   base.String(),
			Latitude:   req.Latitude,
			Longitude:  req.Longitude,
			ErrorType:  fmt.Errorf("%w, %v", apiErrors.ErrInvalidInput, err),
		}
	}

	data := types.WeatherData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		preview := string(body)
		if len(preview) > 12 {
			preview = preview[:12]
		}
		return nil, &apiErrors.WeatherAPIError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("failed to unmarshal response %s, err %v", preview, err),
			Endpoint:   base.String(),
			Latitude:   req.Latitude,
			Longitude:  req.Longitude,
			ErrorType:  fmt.Errorf("%w: %v", apiErrors.ErrInvalidInput, err),
		}
	}

	data.Variables = req.Variables

	return &data, nil
}
