package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/hashicorp/go-envparse"
)

/*
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
*/

type GeoApiResonse struct {
	PlaceId     int     `json:"place_id"`
	License     string  `json:"licence"`
	OsmType     string  `json:"osm_type"`
	OsmID       int     `json:"osm_id"`
	Latitude    string  `json:"lat"`
	Longitude   string  `json:"lon"`
	Class       string  `json:"class"`
	Type        string  `json:"type"`
	PlaceRank   int     `json:"place_rank"`
	Importance  float64 `json:"importance"`
	AddressType string  `json:"addresstype"`
	Name        string  `json:"name"`
	DisplayName string  `json:"display_name"`
	Address     struct {
		CityDistrict string `json:"city_district"`
		Town         string `json:"town"`
		State        string `json:"state"`
		IsoCode      string `json:"ISO3166-2-lvl4"`
		Country      string `json:"country"`
		CountryCode  string `json:"country_code"`
	} `json:"address"`
	ExtraTags   []string `json:"extratags"`
	NameDetails struct {
		Name string `json:"name"`
	} `json:"namedetails"`
	BoundingBox []string `json:"boundingbox"`
}

func FetchLocation(lat, long string) (string, error) {

	var result strings.Builder
	env, err := os.Open("./.env")
	if err != nil {
		return result.String(), fmt.Errorf("error opening env file: %s", err)
	}
	defer env.Close()

	envMap, err := envparse.Parse(env)
	if err != nil {
		return result.String(), fmt.Errorf("error reading env file: %s", err)
	}

	base, _ := url.Parse("https://geocode.maps.co/reverse")
	q := base.Query()
	q.Set("lat", lat)
	q.Set("lon", long)
	q.Set("api_key", envMap["GEOCODE_API"])
	base.RawQuery = q.Encode()

	// fmt.Println(base.String())

	resp, err := http.Get(base.String())

	if err != nil {
		return result.String(), fmt.Errorf("Failed to query data: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return result.String(), fmt.Errorf("HTTP error: %d, %s", resp.StatusCode, string(b))
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return result.String(), fmt.Errorf("IO read error: %v", err)
	}

	obj := GeoApiResonse{}
	err = json.Unmarshal(body, &obj)

	if err != nil {
		return result.String(), fmt.Errorf("JSON unmarshal error: %v", err)
	}

	result.WriteString(obj.Name)
	result.WriteString(", ")
	result.WriteString(obj.Address.Country)
	// result.WriteString(" [")
	// result.WriteString(strings.ToUpper(obj.Address.CountryCode))
	// result.WriteString("]")

	return result.String(), nil
}
