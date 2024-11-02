package server

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/LeRoid-hub/humiditycalc/internal"
)

type result struct {
	AbsoluteHumidity float64
}

// Run starts the HTTP server.
func Run() {
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

		// Write response
		tmpl.Execute(w, result{AbsoluteHumidity: absoluteHumidity})
	})

	http.ListenAndServe(":8080", nil)

}
