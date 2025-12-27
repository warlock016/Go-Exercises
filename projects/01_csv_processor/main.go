package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"sync"

	"github.com/warlock016/csv_processor/config"
	"github.com/warlock016/csv_processor/errors"
	"github.com/warlock016/csv_processor/formatter"
	"github.com/warlock016/csv_processor/input"
	"github.com/warlock016/csv_processor/output"
	apiParser "github.com/warlock016/csv_processor/parser"
	"github.com/warlock016/csv_processor/types"
	"github.com/warlock016/csv_processor/validator"
)

func processPipeline(cli config.CliConfig) types.PipelineResult {
	result := types.PipelineResult{
		Source: cli.ResourcePath,
		Errors: &errors.ProcessingErrors{
			Errors:   make([]errors.FieldError, 0),
			Warnings: make([]errors.FieldError, 0),
		},
	}

	parser, errs := config.NewFileParser(cli)
	if errs.HasFatalErrors() {
		result.Errors = errs
	}

	raw, errs := validator.ValidateRawFile(parser)
	if !errs.HasFatalErrors() {
		result.Errors.Errors = append(result.Errors.Errors, errs.Errors...)
		result.Errors.Warnings = append(result.Errors.Warnings, errs.Warnings...)
	}

	parsed, errs := apiParser.ParseRawData(raw)
	if !errs.HasFatalErrors() {
		result.Errors.Errors = append(result.Errors.Errors, errs.Errors...)
		result.Errors.Warnings = append(result.Errors.Warnings, errs.Warnings...)
	}

	frmt, err := formatter.FormatData(parsed, cli.OutputFormat)
	if err != nil {
		result.Errors.Errors = append(result.Errors.Errors, errors.FieldError{
			Stage:   "Formatting",
			Message: fmt.Sprintf("failed to format data: %v", err),
		})
	}

	result.Output = frmt
	return result
}

func main() {
	cliConfig := config.CliConfig{}
	flag.StringVar(&cliConfig.ConfigPath, "c", "./config/providers/", "provide parser config path")
	flag.StringVar(&cliConfig.ResourcePath, "f", "./testdata/sample.csv", "provide target csv (folder) path")
	flag.StringVar(&cliConfig.Provider, "p", "weathercloud", "provide the YAML provider name")
	flag.StringVar(&cliConfig.Timezone, "tz", "UTC", "set timezone (e.g. Europe/Berlin)")
	flag.StringVar(&cliConfig.Inputformat, "i", "csv", "define input format: csv, json, xml, ...")
	flag.StringVar(&cliConfig.OutputFormat, "o", "csv", "define output format: csv, json-row, json-col, ...")
	flag.Parse()

	paths, err := input.FindFiles(cliConfig.ResourcePath, cliConfig.Inputformat)
	if err != nil {
		log.Fatalf("failed to find valid %s files in %s: %v", cliConfig.Inputformat, cliConfig.ResourcePath, err)
	}
	if len(*paths) == 0 {
		log.Fatalf("no valid %s files in %s", cliConfig.Inputformat, cliConfig.ResourcePath)
	}

	resultsCh := make(chan types.PipelineResult)
	var wg sync.WaitGroup

	for _, p := range *paths {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			cliConfig.ResourcePath = s
			result := processPipeline(cliConfig)
			resultsCh <- result
		}(p)
	}

	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	baseDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get $cwd: %v", err)
	}

	results := make([]types.PipelineResult, 0, len(*paths))
	i := 1
	for result := range resultsCh {
		results = append(results, result)
		name := path.Join(baseDir, "temp", fmt.Sprintf("output_%d.csv", i))
		fmt.Printf("LOG: %s\n", name)
		err := output.DumpToCache(name, []byte(result.Output.Content))
		if err != nil {
			fmt.Printf("failed to dump %s to temp: %v", name, err)
		}
		i++
	}

	fmt.Printf("%d results... \n", len(results))

}
