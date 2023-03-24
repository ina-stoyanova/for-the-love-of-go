package weather

import (
	"log"
	"os"
	"weather/client"
	"weather/printer"
)

func RunApp() {
	if len(os.Args) < 2 {
		log.Fatal("USAGE: weather <TOWN> <COUNTRY_CODE>")
	}

	town := os.Args[1]
	countryCode := os.Args[2]

	weatherResponse := client.GetWeather(town, countryCode)

	printer.PrintPrettyResponse(weatherResponse)
}
