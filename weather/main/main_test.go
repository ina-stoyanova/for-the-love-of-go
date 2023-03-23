package main

import (
	weather "api-client"
	"os"
	"strings"
	"testing"
)

func TestURLConstruction(t *testing.T) {
	testTown := "New York"
	testCountryCode := "US"

	// we want to use a test API key here
	// we shouldn't predict too much how
	// it's useful if the user knows which inputs matter and which don't, so they can just focus on the test and not worry about a valid key
	// testing book
	apiKey := "DUMMY_KEY"
	t.Setenv("OPENWEATHERMAP_API_KEY", apiKey)

	got := ConstructWeatherAPIURL(testTown, testCountryCode)

	if got == "" {
		t.Fatal("was expecting a valid URL, but got empty string instead")
	}

	if !strings.Contains(got, testTown) {
		t.Fatalf("was expecting the URL to contain the town name %s, but it did not", testTown)
	}

	if !strings.Contains(got, testCountryCode) {
		t.Fatalf("was expecting the URL to contain the country code %s, but it did not", testCountryCode)
	}

	if !strings.Contains(got, apiKey) {
		t.Fatalf("was expecting the URL to contain the API key %s, but it did not", apiKey)
	}
}

func TestParseResponse(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}

	parsedResponse, err := weather.ParseResponse(data)
	if err != nil {
		t.Fatalf("Parse response threw an error, but it wasn't supposed to: %s", data)
	}

	if parsedResponse.Name == "" {
		t.Fatalf("API response empty: \n It needs to contain data about one town location under field `Name`: \n Got %s", parsedResponse.Name)
	}

	if len(parsedResponse.Weather) == 0 {
		t.Fatalf("API response not valid: \n It needs to contain at least 1 element with name `Weather`")
	}

	if parsedResponse.Name != "Bansko" {
		t.Fatalf("Data response was not as expected. Wanted `Bansko`, but got %s", parsedResponse.Name)
	}
}
