package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		lat    string
		long   string
		start  string
		end    string
		hourly string // comma-separated list for multiple variables
		tz     string
		format string
	)

	timestamps := make([]string, 0)
	variables := make(map[string][]float64)
	units := make(map[string]string)

	flag.StringVar(&lat, "lat", "52.51", "Latitude")
	flag.StringVar(&long, "long", "13.41", "Longitude")
	flag.StringVar(&start, "start_date", "2025-11-01", "Start date (YYYY-MM-DD)")
	flag.StringVar(&end, "end_date", "2025-11-30", "End date (YYYY-MM-DD)")
	flag.StringVar(&hourly, "hourly", "temperature_2m", "Variables to fetch (use `temperature_2m,...`")
	flag.StringVar(&tz, "timezone", "GMT", "Timezone (use `GMT`)")
	flag.StringVar(&format, "format", "table", "Output modes: JSON, CLI Table")
	flag.Parse()

	// fmt.Println(lat, long, start, end, tz, hourly)
	openMeteo, err := FetchData(lat, long, start, end, hourly, tz)

	if err != nil {
		fmt.Printf("%v", err)
	}

	// Unpacking data for tabular representation
	// data
	for key, value := range openMeteo.Hourly {
		if rawSlice, ok := value.([]any); ok {
			for _, v := range rawSlice {
				switch key {
				case "time":
					if r, ok := v.(string); ok {
						timestamps = append(timestamps, r)
					}
				default:
					if r, ok := v.(float64); ok {
						variables[key] = append(variables[key], r)
					}
				}
			}
		}
	}
	// units
	for key, value := range openMeteo.HourlyUnits {
		units[key] = value.(string)
	}

	byteRep, err := json.MarshalIndent(openMeteo, "", "  ")

	err = os.WriteFile("./cache.json", byteRep, 0o644)

	if err != nil {
		fmt.Printf("Error saving JSON: %v\n", err)
	}

	geoLoc, err := FetchLocation(lat, long)

	if err != nil {
		fmt.Printf("Error fetching geolocation: %v", err)
	}

	result, err := Formatter(&openMeteo, &geoLoc, timestamps, variables, units, format)

	if err != nil {
		fmt.Printf("Formatter error: %v", err)
	}

	fmt.Println(result)
}
