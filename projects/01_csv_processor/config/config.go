package config

import (
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
	Inputformat  string // input format file definition
	OutputFormat string // output format for local output dumping
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

func NewFileParser(cli CliConfig) (*ParserConfig, *apiErrors.ProcessingErrors) {

	resultErr := apiErrors.ProcessingErrors{
		Errors:   make([]apiErrors.FieldError, 0, 4),
		Warnings: make([]apiErrors.FieldError, 0, 4),
	}

	// Validate CLI flag arguments
	if cli.ConfigPath == "" {
		resultErr.AddError("ERR: config", "invalid input", "empty config path", 0, 0)
		return nil, &resultErr
	}
	if cli.ResourcePath == "" {
		resultErr.AddError("ERR: config", "invalid input", "empty resource path", 0, 0)
		return nil, &resultErr
	}
	if cli.Provider == "" {
		resultErr.AddError("ERR: config", "invalid input", "empty provider", 0, 0)
		return nil, &resultErr
	}
	if cli.Timezone == "" {
		resultErr.AddError("ERR: config", "invalid input", "empty timezone", 0, 0)
		return nil, &resultErr
	}

	// Validate YAML config directory
	configDir, err := os.ReadDir(cli.ConfigPath)
	if err != nil {
		resultErr.AddError("ERR: config", "invalid directory", cli.ConfigPath, 0, 0)
		return nil, &resultErr
	}
	if len(configDir) == 0 {
		resultErr.AddError("ERR: config", "empty directory", cli.ConfigPath, 0, 0)
		return nil, &resultErr
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
				resultErr.AddError("ERR: config", "invalid file", err.Error(), 0, 0)
				return nil, &resultErr
			}
			defer f.Close()

			res, err := io.ReadAll(f)
			if err != nil {
				resultErr.AddError("ERR: config", "invalid file", err.Error(), 0, 0)
				return nil, &resultErr
			}

			err = yaml.Unmarshal(res, &newConfig)
			if err != nil {
				resultErr.AddError("ERR: config", "invalid file", err.Error(), 0, 0)
				return nil, &resultErr
			}

			if newConfig.Name == cli.Provider {
				found = true
				break
			}
		}
	}
	if !found {
		resultErr.AddError("ERR: config", "missing config", cli.ConfigPath, 0, 0)
		return nil, &resultErr
	}

	// now we need to validate mandatory fields in newConfig
	if newConfig.Name == "" {
		resultErr.AddError("ERR: config", "invalid config", "invalid name", 0, 0)
		return nil, &resultErr
	}
	if newConfig.Encoding == "" {
		resultErr.AddError("ERR: config", "invalid config", "invalid encoding", 0, 0)
		return nil, &resultErr
	}
	if newConfig.Delimiter == "" {
		resultErr.AddError("ERR: config", "invalid config", "invalid delimiter", 0, 0)
		return nil, &resultErr
	}
	if newConfig.HeaderRows < 0 {
		resultErr.AddError("ERR: config", "invalid config", "negative header count", 0, 0)
		return nil, &resultErr
	}
	if newConfig.SkipRows < 0 {
		resultErr.AddError("ERR: config", "invalid config", "negative skiprow count", 0, 0)
		return nil, &resultErr
	}

	if newConfig.DateConfig.Detection == "" {
		resultErr.AddError("ERR: config", "invalid config", "invalid datetime detection strategy", 0, 0)
		return nil, &resultErr
	}
	if newConfig.DateConfig.Index < 0 {
		resultErr.AddError("ERR: config", "invalid config", "invalid datetime col index", 0, 0)
		return nil, &resultErr
	}
	if len(newConfig.DateConfig.Formats) == 0 {
		resultErr.AddError("ERR: config", "invalid config", "missing datetime layouts", 0, 0)
		return nil, &resultErr
	}
	if cli.Timezone == "" {
		resultErr.AddError("ERR: config", "invalid config", "invalid timezone", 0, 0)
		return nil, &resultErr
	}
	loc, err := time.LoadLocation(cli.Timezone)
	if err != nil {
		resultErr.AddError("ERR: config", "invalid config", "invalid timezone", 0, 0)
		return nil, &resultErr
	}
	newConfig.DateConfig.Timezone = loc

	if len(newConfig.ColConfig) == 0 {
		resultErr.AddError("ERR: config", "invalid config", "invalid column config", 0, 0)
		return nil, &resultErr
	}

	if newConfig.ResourcePath == "" {
		resultErr.AddError("ERR: config", "invalid config", "invalid resource path", 0, 0)
		return nil, &resultErr
	}

	return &newConfig, &resultErr
}
