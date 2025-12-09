package main

import (
	"context"
	"errors"
	"flag"
	"fmt"

	// "go/types"
	"log"
	"time"

	"github.com/warlock016/weather_cli/client"
	"github.com/warlock016/weather_cli/config"
	"github.com/warlock016/weather_cli/formatter"
	"github.com/warlock016/weather_cli/validation"
)

func main() {

	rawInput := validation.CLIInput{}
	flag.StringVar(&rawInput.Latitude, "lat", "52.51", "Latitude")
	flag.StringVar(&rawInput.Longitude, "long", "13.41", "Longitude")
	flag.StringVar(&rawInput.StartDate, "start_date", "2025-11-01", "Start date (YYYY-MM-DD)")
	flag.StringVar(&rawInput.EndDate, "end_date", "2025-11-30", "End date (YYYY-MM-DD)")
	flag.StringVar(&rawInput.Variables, "hourly", "temperature_2m", "Variables to fetch (use `temperature_2m,...`")
	flag.StringVar(&rawInput.Timezone, "timezone", "GMT", "Timezone (use `GMT`)")
	flag.StringVar(&rawInput.Format, "format", "table", "Output modes: JSON, CLI Table")
	flag.Parse()

	validatedInput, err := validation.Validate(rawInput)
	if err != nil {
		log.Fatalf("failed to validate input, error: %v", err)
	} else if validatedInput == nil {
		log.Fatalf("unexpected nil result && nil error")
	}

	newConfig, err := config.Load()
	if err != nil {
		log.Fatalf("unexpected error: %+v", err)
	} else if newConfig == nil {
		log.Fatalf("unexpected nil config")
	}

	ctx := context.Background()
	weatherRequest := client.WeatherRequest{
		Latitude:  validatedInput.Latitude,
		Longitude: validatedInput.Longitude,
		Timezone:  validatedInput.Timezone,
		StartDate: validatedInput.StartDate,
		EndDate:   validatedInput.EndDate,
		Variables: validatedInput.Variables,
	}

	weatherClient := client.NewWeatherClient(newConfig.OpenMeteoURL, time.Second*15)
	if weatherClient == nil {
		log.Fatal("unexpected nil weather client")
	}

	weatherData, err := weatherClient.FetchWeather(ctx, weatherRequest)
	if err != nil {
		log.Fatalf("unexpected err, got %+v", err)
	} else if weatherData == nil {
		log.Fatal("unexpected weather api nil result")
	}

	locationRequest := client.GeoRequest{
		Latitude:  validatedInput.Latitude,
		Longitude: validatedInput.Longitude,
	}

	locationClient := client.NewGeoClient(newConfig.GeocodeURL, newConfig.GeocodeAPIKey, time.Second*15)
	if locationClient == nil {
		log.Fatal("unexpected nil geoapi client")
	}

	locationData, err := locationClient.FetchGeoData(ctx, locationRequest)
	if err != nil {
		log.Fatalf("unexpected err, got %v: %+v", err, errors.Unwrap(err))
	} else if locationData == nil {
		log.Fatal("unexpected geo api nil result")
	}

	newFormat := formatter.FormatterInput{
		Weather:  weatherData,
		Location: locationData,
		Options: formatter.FormatOptions{
			Format:   validatedInput.Format,
			Timezone: validatedInput.Timezone,
		},
	}

	output, err := formatter.Format(newFormat)
	if err != nil {
		log.Fatalf("unexpected error %+v", err)
	}
	if output == nil {
		log.Fatal("unexpected nil Formatter result")
	}

	fmt.Println(output.Content)
}
