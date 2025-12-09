package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/warlock016/weather_cli/client"
	"github.com/warlock016/weather_cli/config"
	"github.com/warlock016/weather_cli/types"
)

const validGeoResponse string = `
{"place_id":65488016,
"licence":"Data © OpenStreetMap contributors, ODbL 1.0. http://osm.org/copyright",
"osm_type":"relation",
"osm_id":3722465,
"lat":"10.0110603",
"lon":"9.9623522",
"class":"boundary",
"type":"administrative",
"place_rank":16,
"importance":0.18671687418936442,
"addresstype":"city_district",
"name":"Mun-Munsal",
"display_name":"Mun-Munsal, Bauchi, Nigeria",
"address":{
	"city_district":"Mun-Munsal",
	"town":"Bauchi",
	"state":"Bauchi",
	"ISO3166-2-lvl4":"NG-BA",
	"country":"Nigeria",
	"country_code":"ng"
	},
"extratags":null,
"namedetails":{
	"name": "Mun-Munsal"
	},
"boundingbox":["9.8953012","10.1286713","9.8773655","10.0987622"]}
`

func setupConfig(t *testing.T) *config.Config {
	t.Helper()
	test, err := config.Load()
	assertNoError(t, err)
	return test
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func assertResponse(t *testing.T, response *types.GeoData) {
	t.Helper()
	if response == nil {
		t.Fatal("expected response, got nil")
	}
}

func assertNoResponse(t *testing.T, response *types.GeoData) {
	t.Helper()
	if response != nil {
		t.Fatalf("expected nil, got %+v", response)
	}
}

func TestFetchGeoData(t *testing.T) {

	testConfig := setupConfig(t)

	if testConfig == nil {
		t.Fatal("failed to setup env variables")
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("lat") == "" {
			t.Error("missing latitude parameter")
		}
		if r.URL.Query().Get("lon") == "" {
			t.Error("missing longitude parameter")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(validGeoResponse))
	}))

	defer server.Close()

	geoClient := client.NewGeoClient(server.URL, testConfig.GeocodeAPIKey, 100*time.Millisecond)
	ctx := context.Background()
	req := client.GeoRequest{
		Latitude:  50,
		Longitude: 25,
	}

	resp, err := geoClient.FetchGeoData(ctx, req)
	assertNoError(t, err)
	assertResponse(t, resp)

	if resp.Country != "Nigeria" {
		t.Errorf("expected country 'Nigeria', got %q", resp.Country)
	}

	if resp.DisplayName != "Mun-Munsal, Bauchi, Nigeria" {
		t.Errorf("got %q, want: 'Mun-Munsal, Bauchi, Nigeria'", resp.DisplayName)
	}
}

func TestFetchGeoData_ServerError(t *testing.T) {

	testConfig := setupConfig(t)
	if testConfig == nil {
		t.Fatal("failed to setup env variables")
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// return server response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(strconv.Itoa(http.StatusInternalServerError))) // mock response
	}))
	defer server.Close()

	geoClient := client.NewGeoClient(server.URL, testConfig.GeocodeAPIKey, 3*time.Second)
	ctx := context.Background()
	req := client.GeoRequest{
		Latitude:  52.51,
		Longitude: 13.41,
	}

	resp, err := geoClient.FetchGeoData(ctx, req)
	assertError(t, err)
	assertNoResponse(t, resp)
}

func TestFetchGeoData_Timeout(t *testing.T) {

	testConfig := setupConfig(t)
	if testConfig == nil {
		t.Fatal("failed to setup config")
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusRequestTimeout)
	}))
	defer server.Close()

	geoClient := client.NewGeoClient(server.URL, testConfig.GeocodeAPIKey, time.Millisecond*1000)
	ctx := context.Background()
	req := client.GeoRequest{
		Latitude:  52.51,
		Longitude: 13.41,
	}

	resp, err := geoClient.FetchGeoData(ctx, req)
	assertError(t, err)
	assertNoResponse(t, resp)
}

func TestFetchGeoData_ClientTimeout(t *testing.T) {

	testConfig := setupConfig(t)
	if testConfig == nil {
		t.Fatal("failed to setup config")
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second) // this server is slow to reply!
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(validGeoResponse))
	}))

	defer server.Close()

	geoClient := client.NewGeoClient(server.URL, testConfig.GeocodeAPIKey, time.Millisecond*100) // fast client timeout -> do not wait too long for slow server!
	ctx := context.Background()
	req := client.GeoRequest{
		Latitude:  52.51,
		Longitude: 13.41,
	}

	resp, err := geoClient.FetchGeoData(ctx, req)
	assertError(t, err)
	assertNoResponse(t, resp)
}

// Testing if Server says OK but sends garbage data
func TestFetchGeoData_MalformedJSON(t *testing.T) {
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

			testConfig := setupConfig(t)
			if testConfig == nil {
				t.Fatal("failed to setup config")
			}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.Body))
			}))

			defer server.Close()

			geoClient := client.NewGeoClient(server.URL, testConfig.GeocodeAPIKey, time.Millisecond*500)
			ctx := context.Background()
			req := client.GeoRequest{
				Latitude:  52.51,
				Longitude: 13.41,
			}

			resp, err := geoClient.FetchGeoData(ctx, req)
			assertError(t, err)
			assertNoResponse(t, resp)
		})
	}
}
