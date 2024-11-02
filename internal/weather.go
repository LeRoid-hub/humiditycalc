package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Weather(env map[string]string) (humidity float64, temperature float64) {
	// Get API key from environment variables
	apiKey := env["OPENWEATHERMAP_API_KEY"]

	// Get latitude and longitude from environment variables
	latitude := env["LATITUDE"]
	longitude := env["LONGITUDE"]

	// Check if latitude and longitude are set
	if latitude == "" || longitude == "" {
		fmt.Println("Latitude and longitude are not set")
		os.Exit(1)
	}

	// Get data from API
	data := getData(apiKey, latitude, longitude)

	// Parse data
	humidity, temperature = parseData(data)

	// Convert to Celsius
	temperature = temperature - 273.15

	return humidity, temperature

}

func getData(apiKey string, latitude string, longitude string) string {
	// Get data from API
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", latitude, longitude, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting data from API")
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body")
	}

	// Convert body to string
	return string(body)
}

func parseData(data string) (humidity float64, temperature float64) {
	// Parse data from API

	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println("Error parsing data")
	}
	n := m["main"].(map[string]interface{})

	return n["humidity"].(float64), n["temp"].(float64)
}
