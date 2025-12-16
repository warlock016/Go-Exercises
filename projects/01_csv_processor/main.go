package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/warlock016/csv_processor/config"
	"github.com/warlock016/csv_processor/parser"
	"github.com/warlock016/csv_processor/validator"
)

func main() {

	cliConfig := config.CliConfig{}
	flag.StringVar(&cliConfig.ConfigPath, "config", "./config/providers/", "provide parser config path")
	flag.StringVar(&cliConfig.ResourcePath, "file", "./testdata/sample.csv", "provide target csv path")
	flag.StringVar(&cliConfig.Provider, "provider", "weathercloud", "provide the YAML provider name")
	flag.StringVar(&cliConfig.Timezone, "timezone", "UTC", "set timezone (e.g. Europe/Berlin)")
	flag.Parse()

	config, err := config.NewFileParser(cliConfig)
	if err != nil {
		log.Fatalf("failed to parse file: %v", err)
	}

	rawResult, rawErr := validator.ValidateRawFile(config)
	// if rawErr.HasFatalErrors() || rawErr.HasWarnings() {
	// 	fmt.Printf("%d errors, %d warnings\n", len(rawErr.Errors), len(rawErr.Warnings))
	// }
	if rawResult == nil {
		log.Fatalf("unexpected nil result for %s\n", cliConfig.ResourcePath)
	}

	fmt.Println(rawResult.HeaderStats)

	fmt.Println(rawResult.BodyStats)

	fmt.Printf("%s", rawErr.Summary())

	parsed, parseErr := parser.ParseRawData(rawResult)
	if parseErr.HasFatalErrors() {
		for _, e := range parseErr.Errors {
			fmt.Printf("ERR: %s: %s: %s ln: %d col: %d\n", e.Stage, e.Message, e.Stage, e.Line, e.Column)
		}
	}
	if parseErr.HasWarnings() {
		for _, w := range parseErr.Warnings {
			fmt.Printf("WRN: %s: %s: %s ln: %d col: %d\n", w.Stage, w.Message, w.Stage, w.Line, w.Column)

		}
	}
	if parsed == nil {
		fmt.Printf("%s", parseErr.Summary())
		log.Fatal("unexpected nil result\n")
	}

	// fmt.Println(parsed.Labels)
	// fmt.Println(parsed.Timezone.String())
	// fmt.Printf("%s\n", parseErr.Summary())

	for _, k := range parsed.Labels {
		if k == parsed.Labels[0] && strings.Contains(k, "Date") {
			continue
		}
		if len(parsed.Datapoints[k]) != len(parsed.Time) {
			fmt.Printf("\"%s\": invalid series length %d\n", k, len(parsed.Datapoints[k]))
		}
	}
	// fmt.Printf("%v\n", parsed.Time)
	// fmt.Printf("%v\n", parsed.Datapoints)
}
