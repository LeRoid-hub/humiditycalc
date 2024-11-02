package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/LeRoid-hub/humiditycalc/internal"
	"github.com/LeRoid-hub/humiditycalc/models"
)

type result struct {
	AbsoluteHumidity string
	WeatherData      WeatherData
}

type WeatherData struct {
	Temperature      string
	RelativeHumidity string
	AbsoluteHumidity string
}

// Run starts the HTTP server.
func Run(env map[string]string) {
	// Cache weather data
	var cacheWeather = models.NewWeatherCache()

	// Load HTML template
	tmpl := template.Must(template.ParseFiles("./web/templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parse query parameters
		tempCelsius, relativeHumidity := r.URL.Query().Get("temp"), r.URL.Query().Get("rh")
		if tempCelsius == "" || relativeHumidity == "" {
			tmpl.Execute(w, nil)
			return
		}

		// Convert query parameters to float64
		temp, err := strconv.ParseFloat(tempCelsius, 32)
		if err != nil {
			http.Error(w, "invalid temperature value", http.StatusBadRequest)
			return
		}
		rh, err := strconv.ParseFloat(relativeHumidity, 32)
		if err != nil {
			http.Error(w, "invalid relative humidity value", http.StatusBadRequest)
			return
		}

		// Calculate absolute humidity
		absoluteHumidity := internal.AbsoluteHumidity(temp, rh)

		var absoluteWeatherHumidity float64
		var temperature, humidity float64

		// Check if cache is expired or empty
		if cacheWeather.IsExpired() || cacheWeather.Humidity == 0 || cacheWeather.Temperature == 0 {
			// Get Weather data
			humidity, temperature := internal.Weather(env)

			// Update cache
			cacheWeather.SetData(humidity, temperature)
		}

		// Use cached data
		humidity, temperature = cacheWeather.GetData()
		absoluteWeatherHumidity = internal.AbsoluteHumidity(temperature, humidity)

		// Create response
		WData := WeatherData{Temperature: FormatFloat(temperature, 2), RelativeHumidity: FormatFloat(humidity, 4), AbsoluteHumidity: FormatFloat(absoluteWeatherHumidity, 4)}

		re := result{AbsoluteHumidity: FormatFloat(absoluteHumidity, 4), WeatherData: WData}

		// Write response
		tmpl.Execute(w, re)
	})

	http.ListenAndServe(":8080", nil)

}

func FormatFloat(num float64, prc int) string {
	var (
		zero, dot = "0", "."

		str = fmt.Sprintf("%."+strconv.Itoa(prc)+"f", num)
	)

	return strings.TrimRight(strings.TrimRight(str, zero), dot)
}
