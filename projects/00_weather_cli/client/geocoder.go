package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	apiErrors "github.com/warlock016/weather_cli/errors"
	"github.com/warlock016/weather_cli/types"
)

type GeoClient struct {
	BaseURL    string
	APIKey     string
	httpClient *http.Client
}

type GeoRequest struct {
	Latitude  float64
	Longitude float64
}

func NewGeoClient(baseURL, apiKey string, timeout time.Duration) *GeoClient {

	return &GeoClient{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		httpClient: &http.Client{Timeout: timeout},
	}
}

func (c *GeoClient) FetchGeoData(ctx context.Context, req GeoRequest) (*types.GeoData, error) {

	if ctx == nil {
		return nil, errors.New("nil context")
	}

	base, _ := url.Parse(c.BaseURL)
	q := base.Query()
	q.Set("lat", strconv.FormatFloat(req.Latitude, 'f', 9, 64))
	q.Set("lon", strconv.FormatFloat(req.Longitude, 'f', 9, 64))
	q.Set("api_key", c.APIKey)
	base.RawQuery = q.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), nil)
	if err != nil {
		return nil, &apiErrors.GeocodeError{
			StatusCode: 0,
			Message:    "failed to create request",
			Query:      base.String(),
			ErrorType:  fmt.Errorf("%w: %v", apiErrors.ErrNetwork, err),
		}
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, &apiErrors.GeocodeError{
			StatusCode: 0,
			Message:    "request failed",
			Query:      base.String(),
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

		return nil, &apiErrors.GeocodeError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("Error %d, got %s", resp.StatusCode, string(b)),
			Query:      base.String(),
			ErrorType:  err,
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &apiErrors.GeocodeError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("invalid input: got %s", string(body)),
			Query:      base.String(),
			ErrorType:  fmt.Errorf("%w, %v", apiErrors.ErrInvalidInput, err),
		}
	}

	data := struct {
		PlaceId     int               `json:"place_id"`
		License     string            `json:"licence"`
		OsmType     string            `json:"osm_type"`
		OsmID       int               `json:"osm_id"`
		Latitude    string            `json:"lat"`
		Longitude   string            `json:"lon"`
		Class       string            `json:"class"`
		Type        string            `json:"type"`
		PlaceRank   int               `json:"place_rank"`
		Importance  float64           `json:"importance"`
		AddressType string            `json:"addresstype"`
		Name        string            `json:"name"`
		DisplayName string            `json:"display_name"`
		Address     map[string]string `json:"address"`
		ExtraTags   map[string]string `json:"extratags"`
		NameDetails struct {
			Name string `json:"name"`
		} `json:"namedetails"`
		BoundingBox []string `json:"boundingbox"`
	}{}
	geoErr := apiErrors.GeocodeError{
		StatusCode: resp.StatusCode,
		Message:    "",
		Query:      base.String(),
		ErrorType:  apiErrors.ErrInvalidInput, // all errors from here are of this type
	}

	if err := os.WriteFile("./weatherCache.json", body, 0o644); err != nil {
		log.Fatalf("error writing to weather cache file: %+v", err)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		geoErr.Message = "unmarshalling error"
		return nil, &geoErr
	}

	lat, err := strconv.ParseFloat(data.Latitude, 64)
	if err != nil {
		geoErr.Message = fmt.Sprintf("invalid latitude %v", lat)
		return nil, &geoErr
	}
	long, err := strconv.ParseFloat(data.Longitude, 64)
	if err != nil {
		geoErr.Message = fmt.Sprintf("invalid longitude %v", long)
		return nil, &geoErr
	}

	if data.DisplayName == "" {
		geoErr.Message = fmt.Sprintf("invalid name %s", data.DisplayName)
		return nil, &geoErr
	}

	return &types.GeoData{
		DisplayName: data.DisplayName,
		City:        data.Address["city"],
		Country:     data.Address["country"],
		CountryCode: data.Address["country_code"],
		Latitude:    lat,
		Longitude:   long,
	}, nil
}
