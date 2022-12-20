package facade

import (
	"fmt"
	"testing"
)

// test with mock data
func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{APIkey: ""}

	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}
	if weather.ID != 3117735 {
		t.Errorf("Madrid id is 3117735, not %d\n", weather.ID)
	}
}

// if there is api then use this test
func TestWithApi(t *testing.T) {
	weatherMap := CurrentWeatherData{"*Apikey"}
	weather, err := weatherMap.GetByCityAndCountryCode("Madrid", "ES")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Temperature in Madrid is %f celsius\n", weather.Main.Temp-273.15)
}
