package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeRoid-hub/humiditycalc/internal"
)

// StartServer starts the HTTP server.
func Run() {
	http.HandleFunc("/", HandleIndex)

	http.HandleFunc("/absolute-humidity", HandleAbsoluteHumidity)

	http.ListenAndServe(":8080", nil)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	html := `<html>
	<head>
		<title>Humidity Calculator</title>
	</head>
	<body>
		<h1>Humidity Calculator</h1>
		<p>Calculate absolute humidity in g/m³ given temperature in Celsius and relative humidity percentage.</p>
		<form action="/absolute-humidity">
			<label for="temp">Temperature (°C):</label>
			<input type="text" id="temp" name="temp" required>
			<br>
			<label for="rh">Relative Humidity (%):</label>
			<input type="text" id="rh" name="rh" required>
			<br>
			<button type="submit">Calculate</button>
		</form>
	</body>
</html>`
	w.Write([]byte(html))
}

func HandleAbsoluteHumidity(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	tempCelsius, relativeHumidity := r.URL.Query().Get("temp"), r.URL.Query().Get("rh")
	if tempCelsius == "" || relativeHumidity == "" {
		http.Error(w, "missing query parameters", http.StatusBadRequest)
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

	// Write response
	w.Write([]byte(fmt.Sprintf("Absolute humidity: %.2f g/m³\n", absoluteHumidity)))
}
