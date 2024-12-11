package main

import (
	"fmt"
	"io" // Import the internal package
	"myweatherapi/internal"
	"net/http"
)

func main() {
	fmt.Println("Hi")

	// Create a new HTTP client
	client := &http.Client{}
	// GetAPIKey - internal package
	apiKey := internal.GetAPIKey()
	url := fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=33.44&lon=-94.04&appid=%s&units=metric", apiKey)

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error in GET request:", err)
		return
	}
	defer resp.Body.Close() // Ensure response body is closed

	//fmt.Println("Request successful!")

	if resp.StatusCode != 200 {
		fmt.Printf("Weather API not available - Error Code: %v", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body) //ignoring error and checking content of response
		fmt.Println("Response", string(body))
	}

}
