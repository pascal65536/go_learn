package main

import "fmt"

type WeatherCondition int

const (
	Clear WeatherCondition = iota
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	multiplier := 1.0

	switch weather.Condition {
	case HeavyRain:
		multiplier += 0.2
	case Snow:
		multiplier += 0.15
	case Rain:
		multiplier += 0.125
	}
	if weather.WindSpeed > 15 {
		multiplier += 0.1
	}
	return multiplier
}

func main() {
	tests := []WeatherData{
		{Condition: Clear, WindSpeed: 10},
		{Condition: Rain, WindSpeed: 10},
		{Condition: HeavyRain, WindSpeed: 10},
		{Condition: Snow, WindSpeed: 10},
		{Condition: Rain, WindSpeed: 20},
		{Condition: HeavyRain, WindSpeed: 20},
		{Condition: Snow, WindSpeed: 20},
	}

	for _, w := range tests {
		fmt.Printf("Condition: %v, WindSpeed: %d -> Multiplier: %.3f\n", w.Condition, w.WindSpeed, GetWeatherMultiplier(w))
	}
}
