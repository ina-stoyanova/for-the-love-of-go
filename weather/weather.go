package weather

import (
	"encoding/json"
	"fmt"
)

// APIResponse "adapter" struct
type APIResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func ParseResponse(data []byte) (APIResponse, error) {
	parsedResponse := APIResponse{}

	if len(data) != 0 {
		err := json.Unmarshal(data, &parsedResponse)
		if err != nil {
			return APIResponse{}, fmt.Errorf("response could not be parsed %v", parsedResponse)
		}
	}

	return parsedResponse, nil
}

// Conditions struct to hold the data we're interested in from our weather client
type Conditions struct {
	Temp        float64
	Town        string
	Country     string
	Description string
}

func MapResponse(response APIResponse) (Conditions, error) {
	if len(response.Weather) < 1 {
		return Conditions{}, fmt.Errorf("response was empty, there were no available")
	}

	return Conditions{
		Temp:        response.Main.Temp,
		Town:        response.Name,
		Country:     response.Sys.Country,
		Description: response.Weather[0].Description,
		// the weather app should handle that weather[0] is always there
	}, nil
}
