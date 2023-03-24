package client

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"strings"
	"testing"
	"weather/parser"
)

func TestWeatherAppResponseFormat(t *testing.T) {
	data, err := os.ReadFile("../testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}

	parsedAPIResponse, err := parser.ParseResponse(data)
	if err != nil {
		t.Fatalf("Parse response threw an error, but it wasn't supposed to: %s", data)
	}

	got, _ := MapResponse(parsedAPIResponse)

	want := GetWeatherResponse{
		Town:                   "Bansko",
		Country:                "BG",
		TempInKelvin:           272.09,
		FeelsLikeTempInKelvin:  272.09,
		TempInCelsius:          0,
		FeelsLikeTempInCelsius: 0,
		Description:            "few clouds",
		Humidity:               91,
		WindSpeed:              0.96,
	}

	if cmp.Equal(got, want) != true {
		t.Fatalf("Wanted %+v, but got %+v instead", want, got)
	}
}

func TestURLConstruction(t *testing.T) {
	testTown := "New York"
	testCountryCode := "US"

	apiKey := "DUMMY_KEY"
	t.Setenv("OPENWEATHERMAP_API_KEY", apiKey)

	got := ConstructWeatherAPIURL(testTown, testCountryCode, apiKey)

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
