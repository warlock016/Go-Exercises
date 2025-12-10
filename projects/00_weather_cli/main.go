package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	// "go/types"
	"log"
	"time"

	"github.com/warlock016/weather_cli/client"
	"github.com/warlock016/weather_cli/config"
	apiErrors "github.com/warlock016/weather_cli/errors"
	"github.com/warlock016/weather_cli/formatter"
	"github.com/warlock016/weather_cli/validation"
)

func main() {

	rawInput := validation.CLIInput{}
	flag.StringVar(&rawInput.Latitude, "lat", "52.51", "Latitude (-90 to 90). Use lat=-74.01 for negative values")
	flag.StringVar(&rawInput.Longitude, "long", "13.41", "Longitude (-180 to 180). Use long=-84.23 for negative values")
	flag.StringVar(&rawInput.StartDate, "start_date", "2025-11-01", "Start date (YYYY-MM-DD) accepted")
	flag.StringVar(&rawInput.EndDate, "end_date", "2025-11-30", "End date (YYYY-MM-DD) HH accepted")
	flag.StringVar(&rawInput.Variables, "hourly", "temperature_2m", "Variables to fetch (use temperature_2m, ...)")
	flag.StringVar(&rawInput.Timezone, "timezone", "GMT", "Timezone (use `GMT`)")
	flag.StringVar(&rawInput.Format, "format", "table", "Output modes: JSON, CLI Table")
	flag.Parse()

	validatedInput, err := validation.Validate(rawInput)
	if err != nil {

		var valErrs *apiErrors.ValidationErrors

		if errors.As(err, &valErrs) {
			fmt.Fprintf(os.Stderr, "Validation errors:")
			for _, e := range valErrs.Errors {
				fmt.Fprintf(os.Stderr, " - %s: %s\n", e.Field, e.Message)
			}
			os.Exit(1)
		}

		if errors.Is(err, apiErrors.ErrNotFound) {
			fmt.Fprintln(os.Stderr, "Location not found")
			os.Exit(1)
		}

		log.Fatalf("failed to validate input, error: %v", err)

	} else if validatedInput == nil {
		log.Fatalf("unexpected nil result && nil error")
	}

	newConfig, err := config.Load()
	if err != nil {
		if errors.Is(err, apiErrors.ErrMissingConfig) {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "unknown error: %v\n", err)
		}
		os.Exit(1)
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
		var weatherErrs *apiErrors.WeatherAPIError
		if errors.As(err, &weatherErrs) {
			fmt.Fprintf(os.Stderr, "OpenMeteo API error: %v", err)
		}

		switch {
		case errors.Is(err, apiErrors.ErrRateLimited):
			fmt.Fprintln(os.Stderr, "Too many requests. Please wait and try again.")
		case errors.Is(err, apiErrors.ErrNotFound):
			fmt.Fprintln(os.Stderr, "Location or resource not found.")
		case errors.Is(err, apiErrors.ErrUnauthorized):
			fmt.Fprintln(os.Stderr, "Authentication failed. Check your API key.")
		case errors.Is(err, apiErrors.ErrNetwork):
			fmt.Fprintln(os.Stderr, "Network error. Check your connection and try again")
		case errors.Is(err, apiErrors.ErrInvalidInput):
			fmt.Fprintln(os.Stderr, "Invalid request parameters.")
		default:
			fmt.Fprintf(os.Stderr, "Unexpected error :%v\n", err)
		}

		os.Exit(1)

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
		var locErrs *apiErrors.GeocodeError
		if errors.As(err, &locErrs) {
			fmt.Fprintf(os.Stderr, "Geocode API error: %v", err)
		}

		switch {
		case errors.Is(err, apiErrors.ErrRateLimited):
			fmt.Fprintln(os.Stderr, "Too many requests. Please wait and try again.")
		case errors.Is(err, apiErrors.ErrNotFound):
			fmt.Fprintln(os.Stderr, "Location or resource not found.")
		case errors.Is(err, apiErrors.ErrUnauthorized):
			fmt.Fprintln(os.Stderr, "Authentication failed. Check your API key.")
		case errors.Is(err, apiErrors.ErrNetwork):
			fmt.Fprintln(os.Stderr, "Network error. Check your connection and try again")
		case errors.Is(err, apiErrors.ErrInvalidInput):
			fmt.Fprintln(os.Stderr, "Invalid request parameters.")
		default:
			fmt.Fprintf(os.Stderr, "Unexpected error :%v\n", err)
		}

		os.Exit(1)
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
	if output.Content == "" {
		log.Fatal("unexpected empty string")
	}

	fmt.Println(output.Content)
}
