package printer

import (
	"fmt"
	"math"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"

	"weather/client"
)

type FormattedWeatherConditions struct {
	TimeNow                string
	Town                   string
	Country                string
	TempInKelvin           string
	TempInCelsius          string
	FeelsLikeTempInKelvin  string
	FeelsLikeTempInCelsius string
	IsCelsius              string
	Description            string
	Humidity               string
	WindSpeed              string
}

// PrintPrettyResponse returns a pretty formatted response
func PrintPrettyResponse(conditions client.GetWeatherResponse) {
	fmt.Println(text.BgCyan.Sprintf("The weather in %v:", text.Bold.Sprintf("%v, %v", conditions.Town, conditions.Country)))
	fmt.Printf("%v\t%v\n", text.FgCyan.Sprintf("Time now:"), getStringTimeNow())
	fmt.Printf("%v\t%v °C\n", text.FgCyan.Sprintf("Temperature:"), convertFromKelvinToCelsius(conditions.TempInKelvin))
	fmt.Printf("%v\t%v °C\n", text.FgCyan.Sprintf("Feels like:"), convertFromKelvinToCelsius(conditions.FeelsLikeTempInKelvin))
	fmt.Printf("%v\t%v\n", text.FgCyan.Sprintf("What to expect:"), conditions.Description)
	fmt.Printf("%v\t%v %% \n", text.FgCyan.Sprintf("Humidity:"), conditions.Humidity)
	fmt.Printf("%v\t%v Km/h \n", text.FgCyan.Sprintf("Wind speed:"), convertMphToKmph(conditions.WindSpeed))
}

func convertMphToKmph(speedInMph float64) float64 {
	return math.Round(speedInMph * 1.609)
}

func convertFromKelvinToCelsius(tempInKelvin float64) float64 {
	return math.Round(tempInKelvin - 273.15)
}

func getStringTimeNow() string {
	return time.Now().Format("Monday 15:04:05 2006-01-02")
}
