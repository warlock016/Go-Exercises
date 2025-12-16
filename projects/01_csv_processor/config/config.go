package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	apiErrors "github.com/warlock016/csv_processor/errors"
	"gopkg.in/yaml.v3"
)

type CliConfig struct {
	ResourcePath string // folder path to csv files
	ConfigPath   string // folder path to yaml configs
	Provider     string // provider name should match provider_name in yaml file
	Timezone     string // optional timezone override
}

type DateTimeConfig struct {
	Detection string   `yaml:"detection"`
	Index     int      `yaml:"index"`
	Formats   []string `yaml:"formats"` // holds the datetime string layouts (e.g. "2006-01-02 15:04:15")
	Timezone  *time.Location
}

type ColumnConfig struct {
	Name          string   `yaml:"name"`
	Type          string   `yaml:"type"`
	Aliases       []string `yaml:"aliases"`
	Required      bool     `yaml:"required"`
	AvailableFrom string   `yaml:"available_from"` // optional
}

type ParserConfig struct {
	Name           string         `yaml:"name"`      // mandatory
	Encoding       string         `yaml:"encoding"`  // mandatory
	Delimiter      string         `yaml:"delimiter"` // mandatory
	Comment        string         `yaml:"comment"`   // optional
	SkipRows       int            `yaml:"skip_rows"`
	HeaderRows     int            `yaml:"header_rows"`         // optional
	DigitSeparator string         `yaml:"thousands_separator"` //optional
	DateConfig     DateTimeConfig `yaml:"datetime"`            // mandatory
	ColConfig      []ColumnConfig `yaml:"columns"`             // mandatory
	ResourcePath   string         // file/folder path, based on value of ParserMode
	ConfigPath     string
	Provider       string
	Timezone       string
}

func NewFileParser(cli CliConfig) (*ParserConfig, error) {

	// CLI flag arguments
	if cli.ConfigPath == "" {
		return nil, fmt.Errorf("empty YAML config path: %w", apiErrors.ErrInvalidInput)
	}
	if cli.ResourcePath == "" {
		return nil, fmt.Errorf("empty resource path: %w", apiErrors.ErrInvalidInput)
	}
	if cli.Provider == "" {
		return nil, fmt.Errorf("empty provider name: %w", apiErrors.ErrInvalidInput)
	}

	// Validate YAML config directory
	configDir, err := os.ReadDir(cli.ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML config directory: %w", err)
	}
	if len(configDir) == 0 {
		return nil, fmt.Errorf("empty YAML config directory: %w", apiErrors.ErrNotFound)
	}

	newConfig := ParserConfig{
		ResourcePath: cli.ResourcePath,
		ConfigPath:   cli.ConfigPath,
		Provider:     cli.Provider,
		Timezone:     cli.Timezone,
	}

	// Find and read target YAML config file
	found := false
	for _, file := range configDir {
		if !file.IsDir() {
			target := filepath.Join(cli.ConfigPath, file.Name())
			f, err := os.Open(target)
			if err != nil {
				return nil, fmt.Errorf("%s not found: %w", target, err)
			}
			defer f.Close()

			res, err := io.ReadAll(f)
			if err != nil {
				return nil, fmt.Errorf("failed to read %s: %w", target, err)
			}

			err = yaml.Unmarshal(res, &newConfig)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal YAML: %s %w", target, err)
			}

			if newConfig.Name == cli.Provider {
				found = true
				break
			}
		}
	}
	if !found {
		return nil, fmt.Errorf("YAML config not found: %w", apiErrors.ErrNotFound)
	}

	// now we need to validate mandatory fields in newConfig
	if newConfig.Name == "" {
		return nil, fmt.Errorf("unexpected empty provider name in YAML config: %w", apiErrors.ErrInvalidInput)
	}
	if newConfig.Encoding == "" {
		return nil, fmt.Errorf("empty encoding in YAML config: %w", apiErrors.ErrInvalidInput)
	}
	if newConfig.Delimiter == "" {
		return nil, fmt.Errorf("empty delimiter in YAML config: %w", apiErrors.ErrInvalidInput)
	}
	if newConfig.HeaderRows < 0 {
		return nil, fmt.Errorf("invalid header rows in YAML config: %w", apiErrors.ErrInvalidInput)
	}
	if newConfig.SkipRows < 0 {
		return nil, fmt.Errorf("invalid skip rows in YAML config: %w", apiErrors.ErrInvalidInput)
	}

	if newConfig.DateConfig.Detection == "" {
		return nil, fmt.Errorf("empty datetime detection in YAML config: %w", apiErrors.ErrInvalidInput)
	}
	if newConfig.DateConfig.Index < 0 {
		return nil, fmt.Errorf("invalid datetime index in YAML config: %w", apiErrors.ErrInvalidInput)
	}
	if len(newConfig.DateConfig.Formats) == 0 {
		return nil, fmt.Errorf("empty datetime formats in YAML config: %w", apiErrors.ErrInvalidInput)
	}
	if cli.Timezone == "" {
		return nil, fmt.Errorf("empty timezone override: %w", apiErrors.ErrInvalidInput)
	}
	if cli.Timezone != "" {
		loc, err := time.LoadLocation(cli.Timezone)
		if err != nil {
			return nil, fmt.Errorf("invalid timezone override: %w", err)
		}
		newConfig.DateConfig.Timezone = loc
	}

	if len(newConfig.ColConfig) == 0 {
		return nil, fmt.Errorf("empty columns configuration in YAML config: %w", apiErrors.ErrInvalidInput)
	}

	if newConfig.ResourcePath == "" {
		return nil, fmt.Errorf("unexpected empty resource path in YAML config: %w", apiErrors.ErrInvalidInput)
	}

	return &newConfig, nil
}
