package types

import (
	"time"
)

// type Reading struct {
// 	Timestamp time.Time // Ideally converted datetime string + timezone to Unix
// 	Value     float64
// }

// type Metadata struct {
// 	Timestamp time.Time
// 	Value     string
// }

// Shared Data Structures
type RawData struct {
	Header      [][]string
	HeaderStats map[int]int
	HeaderWidth int
	Body        [][]string
	BodyStats   map[int]int
	BodyWidth   int
	Source      string

	// Fields from ParserConfig struct
	DatetimeIndex   int
	DatetimeFormats []string
	Timezone        *time.Location
	DigitSeparator  string
}

type ParsedData struct {
	Time        []time.Time          // timestamps
	Timezone    *time.Location       // time series timezone
	Labels      []string             // time series names
	Datapoints  map[string][]float64 // numeric time series
	Metadata    map[string][]string  // string time series
	SkippedCols map[int]bool
}

type FormattedOutput struct {
	Content string
	Type    string
}
