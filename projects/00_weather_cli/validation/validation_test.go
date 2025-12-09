package validation_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/warlock016/weather_cli/validation"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   validation.CLIInput
		want    validation.ValidatedInput
		wantLen int
		wantErr bool
	}{
		{
			name: "valid input",
			input: validation.CLIInput{
				Latitude:  "52.52",
				Longitude: "13.41",
				StartDate: "2025-01-01",
				EndDate:   "2025-01-31",
				Format:    "json",
				Variables: "temperature_2m,wind_speed_10m",
			},
			want: validation.ValidatedInput{
				Latitude:  52.52,
				Longitude: 13.41,
				StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2025, 1, 31, 0, 0, 0, 0, time.UTC),
				Format:    "json",
				Variables: []string{"temperature_2m", "wind_speed_10m"},
			},
			wantLen: 2,
			wantErr: false,
		},
		{
			name: "invalid latitude",
			input: validation.CLIInput{
				Latitude:  "180",
				Longitude: "13.41",
				StartDate: "2025-01-01",
				EndDate:   "2025-01-31",
				Format:    "json",
				Variables: "temperature_2m,wind_speed_10m",
			},
			want: validation.ValidatedInput{
				Latitude:  0,
				Longitude: 13.41,
				StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2025, 1, 31, 0, 0, 0, 0, time.UTC),
				Format:    "json",
				Variables: []string{"temperature_2m", "wind_speed_10m"},
			},
			wantLen: 2,
			wantErr: true,
		},
		{
			name: "invalid longitude",
			input: validation.CLIInput{
				Latitude:  "10",
				Longitude: "-183",
				StartDate: "2025-01-01",
				EndDate:   "2025-01-31",
				Format:    "json",
				Variables: "temperature_2m,wind_speed_10m",
			},
			want: validation.ValidatedInput{
				Latitude:  10,
				Longitude: 0,
				StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2025, 1, 31, 0, 0, 0, 0, time.UTC),
				Format:    "json",
				Variables: []string{"temperature_2m", "wind_speed_10m"},
			},
			wantLen: 2,
			wantErr: true,
		},
		{
			name: "start after end",
			input: validation.CLIInput{
				Latitude:  "10",
				Longitude: "-20",
				StartDate: "2025-01-31",
				EndDate:   "2025-01-01",
				Format:    "json",
				Variables: "temperature_2m,wind_speed_10m",
			},
			want: validation.ValidatedInput{
				Latitude:  10,
				Longitude: -20,
				StartDate: time.Date(2025, 1, 31, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				Format:    "json",
				Variables: []string{"temperature_2m", "wind_speed_10m"},
			},
			wantLen: 2,
			wantErr: true,
		},
		{
			name: "missing start date",
			input: validation.CLIInput{
				Latitude:  "10",
				Longitude: "0",
				StartDate: "",
				EndDate:   "2025-01-01",
				Format:    "json",
				Variables: "temperature_2m,wind_speed_10m",
			},
			want: validation.ValidatedInput{
				Latitude:  10,
				Longitude: 0,
				StartDate: time.Time{},
				EndDate:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				Format:    "json",
				Variables: []string{"temperature_2m", "wind_speed_10m"},
			},
			wantLen: 2,
			wantErr: true,
		},
		{
			name: "missing end date",
			input: validation.CLIInput{
				Latitude:  "10",
				Longitude: "-10",
				StartDate: "2025-01-01",
				EndDate:   "",
				Format:    "json",
				Variables: "temperature_2m,wind_speed_10m",
			},
			want: validation.ValidatedInput{
				Latitude:  10,
				Longitude: -10,
				StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Time{},
				Format:    "json",
				Variables: []string{"temperature_2m", "wind_speed_10m"},
			},
			wantLen: 2,
			wantErr: true,
		},
		{
			name: "empty variable",
			input: validation.CLIInput{
				Latitude:  "0",
				Longitude: "0",
				StartDate: "2025-01-01",
				EndDate:   "2025-12-31",
				Format:    "json",
				Variables: "",
			},
			want: validation.ValidatedInput{
				Latitude:  0,
				Longitude: 0,
				StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC),
				Format:    "json",
				Variables: []string{},
			},
			wantLen: 0,
			wantErr: true,
		},
		{
			name: "single variable",
			input: validation.CLIInput{
				Latitude:  "0",
				Longitude: "0",
				StartDate: "2025-01-01",
				EndDate:   "2025-12-31",
				Format:    "json",
				Variables: "temperature_2m",
			},
			want: validation.ValidatedInput{
				Latitude:  0,
				Longitude: 0,
				StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC),
				Format:    "json",
				Variables: []string{"temperature_2m"},
			},
			wantLen: 1,
			wantErr: false,
		},
		{
			name: "single variable bad sep",
			input: validation.CLIInput{
				Latitude:  "0",
				Longitude: "0",
				StartDate: "2025-01-01",
				EndDate:   "2025-12-31",
				Format:    "json",
				Variables: "temperature_2m,",
			},
			want: validation.ValidatedInput{
				Latitude:  0,
				Longitude: 0,
				StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC),
				Format:    "json",
				Variables: []string{"temperature_2m"},
			},
			wantLen: 1,
			wantErr: false,
		},
		{
			name: "single variable bad format",
			input: validation.CLIInput{
				Latitude:  "0",
				Longitude: "0",
				StartDate: "2025-01-01",
				EndDate:   "2025-12-31",
				Format:    "csv",
				Variables: "temperature_2m",
			},
			want: validation.ValidatedInput{
				Latitude:  0,
				Longitude: 0,
				StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC),
				Format:    "",
				Variables: []string{"temperature_2m"},
			},
			wantLen: 1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validation.Validate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error %v,  wantErr: %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Validate() got %+v, want %+v", *got, tt.want)
			}
			if len(got.Variables) != tt.wantLen {
				t.Errorf("Validate() got %+v, want %+v", *got, tt.wantLen)
			}
		})
	}
}
