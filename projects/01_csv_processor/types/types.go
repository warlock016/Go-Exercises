package types

import (
	"time"
)

// Shared Data Structures
type RawData struct {
	Header      [][]string
	HeaderStats map[int]int
	HeaderWidth int
	Body        [][]string
	BodyStats   map[int]int
	BodyWidth   int
	Source      string
}

type ParsedData struct {
	Time       []time.Time          // timestamps
	Timezone   *time.Location       // time series timezone
	Labels     []string             // time series names
	Datapoints map[string][]float64 // numeric time series
	Metadata   map[string][]string  // string time series
}

type FormattedOutput struct {
	Content string
	Type    string
}
