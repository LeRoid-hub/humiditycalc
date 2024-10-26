package internal

import (
	"math"
)

// AbsoluteHumidity calculates absolute humidity in g/m³ given temperature in Celsius and relative humidity percentage.
func AbsoluteHumidity(tempCelsius, relativeHumidity float64) float64 {
	// Constants
	const Mw = 18.016 // Molar mass of water vapor in g/mol
	const R = 8.314   // Universal gas constant in J/(mol·K)
	const A = 6.112   // Constant for saturation vapor pressure in hPa
	const B = 17.67   // Constant for saturation vapor pressure
	const C = 243.5   // Constant for saturation vapor pressure in Celsius

	// Convert temperature from Celsius to Kelvin
	tempKelvin := tempCelsius + 273.15

	// Calculate saturation vapor pressure (in hPa) using temperature in Celsius
	saturationVaporPressure := A * math.Exp(B*tempCelsius/(tempCelsius+C))

	// Calculate absolute humidity
	absoluteHumidity := (relativeHumidity / 100) * saturationVaporPressure * Mw / (R * tempKelvin)

	return absoluteHumidity
}
