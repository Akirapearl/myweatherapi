package main

import (
	"encoding/json"
	"fmt"
	"io" // Import the internal package
	"myweatherapi/internal"
	"net/http"
)

// Idea is to get specific "core" details about the weather: Location (City), Temperature (Celsius), Sky detail (i.e Cloudy, Sunny, Rainy...)
// Upgrade point: call /astronomy.json with a new struct - get moon phase, rise and set data
type Weatherdata struct {
	Location struct {
		City string `json:"name"`
	} `json:"location"`
	/*
		Expected response
		{"location":{"name":"Barcelona","region":"Catalonia","country":"Spain","lat":41.3833,"lon":2.1833,"tz_id":"Europe/Madrid","localtime_epoch":1734278478,"localtime":"2024-12-15 17:01"}
	*/

	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`

	/*
			Expected response
		"current":{"last_updated_epoch":1734278400,"last_updated":"2024-12-15 17:00","temp_c":13.3,"temp_f":55.9,"is_day":1,"condition":{"text":"Partly cloudy","icon":"//cdn.weatherapi.com/weather/64x64/day/116.png","code":1003}
	*/
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	// Create a new HTTP client
	client := &http.Client{}
	// GetAPIKey - internal package
	apiKey := internal.GetAPIKey()
	// Openweather
	//url := fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=33.44&lon=-94.04&appid=%s&units=metric", apiKey)
	// weatherapi
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=Barcelona&aqi=no", apiKey)
	//url := fmt.Sprintf("http://127.0.0.1:8080/albums")
	//resp, err := http.Get(url)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error in GET request:", err)
		return
	}
	defer resp.Body.Close() // Ensure response body is closed

	if resp.StatusCode != 200 {
		fmt.Printf("Weather API not available - Error Code: %v", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body) //ignoring error and checking content of response
		fmt.Println("Response", string(body))
	}
	body, err := io.ReadAll(resp.Body)
	/*
		if err != nil {
			panic(err)

		}*/
	checkError(err)
	/*
		-- Basic API CALL - just returns the entire body of the JSON output --
		fmt.Println(string(body))
	*/

	var tiempo Weatherdata
	err = json.Unmarshal(body, &tiempo)
	checkError(err)

	city, weather, sky := tiempo.Location, tiempo.Current.TempC, tiempo.Current.Condition

	fmt.Printf("%v %v°C %v", city.City, weather, sky.Text)
}
