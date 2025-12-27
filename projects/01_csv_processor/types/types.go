package types

import (
	"time"

	"github.com/warlock016/csv_processor/errors"
)

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
	Time       []time.Time          // timestamps
	Timezone   *time.Location       // time series timezone
	Labels     []string             // time series names
	Datapoints map[string][]float64 // numeric time series
	Metadata   map[string][]string  // string time series
	// SkippedCols   map[int]bool
	LabelColIndex map[int]int
	Source        string
}

type FormattedOutput struct {
	Content string
	Type    string
}

type JsonData struct {
	Source     string      `json:"source"`
	Timezone   string      `json:"timezone"`
	Count      int         `json:"count"`
	Time       []string    `json:"time,omitempty"`
	Labels     []string    `json:"labels,omitempty"`
	Records    []RowRecord `json:"records,omitempty"`
	Datapoints ColRecord   `json:"datapoints,omitempty"`
}

type RowRecord map[string]any // any could be string or float (json.RawMessage)

type ColRecord map[string][]any

type PipelineResult struct {
	Output *FormattedOutput
	Errors *errors.ProcessingErrors
	Source string
}
