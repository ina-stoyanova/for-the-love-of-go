package weather

import (
	"os"
	"testing"
)

// By adding _test to the package, we limit the usage of private funcs and structs
// by testing the package from outside, we're imitating the user - if anything is awakward for users, we should be the first ones to find out
// always use the same "want" and "got" or whatever else I choose to make it easier for others and myself when reading my code --> "err"

// We will test the format output of the app
func TestWeatherAppResponseFormat(t *testing.T) {
	// what is the specific behaviour I want to test?
	// i want a new behaviour in my program & how would I express it
	// iterate on the test until it gets so specific it's almost code

	// given the test data about Bansko
	// relatiuve paths here
	data, err := os.ReadFile("../testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}

	// we have tp get it in the right starting format
	parsedAPIResponse, err := ParseResponse(data)
	if err != nil {
		t.Fatalf("Parse response threw an error, but it wasn't supposed to: %s", data)
	}

	// we get back the result in the format defined
	// we actually have very little error margin here, so no need for an err
	//design errors out of existence
	got, _ := MapResponse(parsedAPIResponse)

	// this is my test data, but this is the shortest path to get the test working
	// the most reliable way to generate test data is by hand
	// easy for readers of the code too
	want := Conditions{
		Temp:        272.09,
		Town:        "Bansko",
		Country:     "BG",
		Description: "few clouds",
	}

	// we get back information that is the same in the input data as well as the output data
	if got != want {
		t.Fatalf("Wanted %+v, but got %+v instead", want, got)
	}

}

// if I call mapresponse and something is nil or empty, then it should not panic, but return an error & show to user
//func TestMapResponseReturnsError() {
//
//}

// is it working?
// is it delightful? how could it be better? ok turn it into a test
// it feels slow? but am I measuring it the right way? What is actually important here
