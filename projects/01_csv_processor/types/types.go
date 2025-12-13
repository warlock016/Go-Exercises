package types

import (
	"time"
)

// Shared Data Structures
type RawData struct {
	Header      [][]string
	HeaderWidth int
	Body        [][]string
	BodyWidth   int
	Source      string
}

type ParsedData struct {
	Time       []time.Time          // timestamps
	Timezone   string               // time series timezone
	Labels     []string             // time series names
	Datapoints map[string][]float64 // numeric time series
	Metadata   map[string][]string  // string time series
}

type FormattedOutput struct {
	Content string
	Type    string
}
