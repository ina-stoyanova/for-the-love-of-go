package main

import (
	"api-client/weather"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// we could change this in the test
// local test server - as simple as possible, always return the same data (e..g take it from the real server)
// 2 parts to test this: the URL construction; given the correct input (e.g. json), what's the output of my program
const baseURL = "https://api.openweathermap.org"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("USAGE: weather <TOWN> <COUNTRY_CODE>")
	}

	town := os.Args[1]
	countryCode := os.Args[2]

	weatherResponse := getWeather(town, countryCode)

	parsed, err := weather.ParseResponse(weatherResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(parsed)
}

func getWeather(town, countryCode string) []byte {
	URL := constructWeatherAPIURL(town, countryCode)

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	dataStream, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return dataStream
}

func constructWeatherAPIURL(town, countryCode string) string {
	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		log.Fatal("Please set the ENV variable OPENWEATHERMAP_API_KEY.")

		return ""
	}

	return fmt.Sprintf("%s/data/2.5/weather?q=%s,%s&appid=%s", baseURL, town, countryCode, key)
}
