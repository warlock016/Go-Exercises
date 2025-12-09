package main

import "time"

type Datapoint struct {
	Value     float64
	Timestamp time.Time // unix timestamp
}

type Timeseries struct {
	Name string      // time series name ("temperature","irradiation",...)
	Data []Datapoint // sorted collection of timestamp/value Datapoint structs
}

type Collection struct {
	Name        string // collection name
	Location    string
	Description string
	Latitude    float64
	Longitude   float64
	Timezone    string
	Data        map[string]Timeseries // unsorted collection of Timeseries accessed by key
}
