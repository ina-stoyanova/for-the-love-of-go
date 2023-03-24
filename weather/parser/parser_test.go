package parser

import (
	"os"
	"testing"
)

func TestParseResponse(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile("../testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}

	parsedResponse, err := ParseResponse(data)
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
