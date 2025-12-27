package config_test

import (
	"testing"

	"github.com/warlock016/csv_processor/config"
)

func TestLoadParserConfig(t *testing.T) {

	tests := []struct {
		Name         string
		ConfigPath   string
		ResourcePath string
		Provider     string
		Timezone     string
		wantErr      bool
	}{
		{"local tz file", "./providers/", "../testdata/Weathercloud Pupuseria El Mirador 2025-01.csv", "weathercloud", "America/Costa_Rica", false},
		{"utc tz file", "./providers/", "../testdata/Weathercloud Pupuseria El Mirador 2025-01.csv", "weathercloud", "UTC", false},
		{"invalid tz file", "./providers/", "../testdata/Weathercloud Pupuseria El Mirador 2025-01.csv", "weathercloud", "", true},
		{"all valid", "./providers/", "../testdata/", "weathercloud", "America/Costa_Rica", false},
		{"invalid resource", "./provid/src", "../testdata/", "weathercloud", "America/Costa_Rica", true},
		{"invalid provider", "./providers/", "../testdata/", "meteocontrol", "America/Costa_Rica", true},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			cliInput := config.CliConfig{
				ResourcePath: tt.ResourcePath,
				ConfigPath:   tt.ConfigPath,
				Provider:     tt.Provider,
				Timezone:     tt.Timezone,
			}

			res, err := config.NewFileParser(cliInput)

			switch {
			case tt.wantErr && err.HasFatalErrors():
				return
			case tt.wantErr && !err.HasFatalErrors():
				t.Fatal("unexpected nil error, want err")
			case !tt.wantErr && err.HasFatalErrors():
				t.Fatalf("unexpected error: %v", err)
			case !tt.wantErr && res == nil:
				t.Fatalf("unexpected nil result")
			}

			if res.DateConfig.Timezone.String() != tt.Timezone {
				t.Fatalf("unexpected timezone, got %s, want %s", res.DateConfig.Timezone.String(), tt.Timezone)
			}
		})
	}
}
