package config_test

import (
	"fmt"
	"testing"

	"github.com/warlock016/csv_processor/config"
)

func TestLoadParserConfig(t *testing.T) {

	tests := []struct {
		Name         string
		ConfigPath   string
		ResourcePath string
		Provider     string
		Mode         string
		wantErr      bool
	}{
		{"single file", "./providers/", "../testdata/Weathercloud Pupuseria El Mirador 2025-01.csv", "weathercloud", "file", false},
		{"all valid", "./providers/", "../testdata/", "weathercloud", "folder", false},
		{"invalid resource", "./provid/src", "../testdata/", "weathercloud", "file", true},
		{"invalid provider", "./providers/", "../testdata/", "meteocontrol", "file", true},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			cliInput := config.CliConfig{
				ResourcePath: tt.ResourcePath,
				ConfigPath:   tt.ConfigPath,
				Provider:     tt.Provider,
				Mode:         tt.Mode,
			}

			res, err := config.NewFileParser(cliInput)

			switch {
			case tt.wantErr && err != nil:
				return
			case tt.wantErr && err == nil:
				t.Fatal("unexpected nil error, want err")
			case !tt.wantErr && err != nil:
				t.Fatalf("unexpected error: %v", err)
			case !tt.wantErr && res == nil:
				t.Fatalf("unexpected nil result")
			}

			fmt.Println(res)
		})
	}
}
