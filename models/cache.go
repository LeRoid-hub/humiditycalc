package models

import (
	"time"
)

type WeatherCache struct {
	Humidity    float64   // relative humidity in percentage
	Temperature float64   // temperature in Celsius
	timestamp   time.Time // timestamp of last update
	duration    int16     // duration in seconds
}

func (w *WeatherCache) IsExpired() bool {
	return time.Since(w.timestamp) > time.Duration(w.duration)*time.Second
}

func (w *WeatherCache) SetData(humidity float64, temperature float64) {
	w.Humidity = humidity
	w.Temperature = temperature
	w.timestamp = time.Now()
	w.duration = 60
}

func (w *WeatherCache) GetData() (float64, float64) {
	if w.IsExpired() {
		return 0, 0
	}
	return w.Humidity, w.Temperature
}

func NewWeatherCache() *WeatherCache {
	return &WeatherCache{}
}
