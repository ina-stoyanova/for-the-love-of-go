package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"weather/parser"
)

const baseURL = "https://api.openweathermap.org"

type WeatherClient struct {
	APIKey string
}

type GetWeatherResponse struct {
	Town                   string
	Country                string
	TempInKelvin           float64
	TempInCelsius          float64
	FeelsLikeTempInKelvin  float64
	FeelsLikeTempInCelsius float64
	Description            string
	Humidity               int
	WindSpeed              float64
}

// ConstructWeatherAPIURL constructs the URL our client will make a request to
func ConstructWeatherAPIURL(town, countryCode, apiKey string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s,%s&appid=%s", baseURL, town, countryCode, apiKey)
}

func GetWeather(town, countryCode string) GetWeatherResponse {
	client := NewWeatherClient()

	URL := ConstructWeatherAPIURL(town, countryCode, client.APIKey)

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	dataStream, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	parsedAPIResponse, err := parser.ParseResponse(dataStream)
	if err != nil {
		fmt.Printf("Parse response threw an error, but it wasn't supposed to: %s", dataStream)
		os.Exit(0)
	}

	response, err := MapResponse(parsedAPIResponse)
	if err != nil {
		fmt.Printf("Parser threw an error while parsing the API response: %s", err)
		os.Exit(0)
	}

	return GetWeatherResponse{
		Town:                  response.Town,
		Country:               response.Country,
		TempInKelvin:          response.TempInKelvin,
		FeelsLikeTempInKelvin: response.FeelsLikeTempInKelvin,
		Description:           response.Description,
		Humidity:              response.Humidity,
		WindSpeed:             response.WindSpeed,
	}
}

func NewWeatherClient() WeatherClient {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set the ENV variable OPENWEATHERMAP_API_KEY.")

		return WeatherClient{}
	}

	return WeatherClient{
		APIKey: apiKey,
	}
}

func MapResponse(response parser.APIResponse) (GetWeatherResponse, error) {
	if len(response.Weather) < 1 {
		return GetWeatherResponse{}, fmt.Errorf("response was empty, there was no available weather data")
	}

	return GetWeatherResponse{
		Town:                  response.Name,
		Country:               response.Sys.Country,
		TempInKelvin:          response.Main.Temp,
		FeelsLikeTempInKelvin: response.Main.FeelsLike,
		Description:           response.Weather[0].Description,
		Humidity:              response.Main.Humidity,
		WindSpeed:             response.Wind.Speed,
	}, nil
}
