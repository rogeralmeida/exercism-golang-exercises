// Package weather This package offers functions to work with weather.
package weather

// CurrentCondition this holds the current condition data.
var CurrentCondition string

// CurrentLocation this holds the current location data.
var CurrentLocation string

// Forecast given a city and condition this predicts the weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
