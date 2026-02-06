package main

import "fmt"

const (
	pricePerKm      = 10.0
	pricePerMinute  = 2.0
)

type TripParameters struct {
	Distance float64
	Duration float64
}

func CalculateBasePrice(params TripParameters) float64 {
	return params.Distance*pricePerKm + params.Duration*pricePerMinute
}

func main() {
	trip := TripParameters{
		Distance: 15.5,
		Duration: 30.0,
	}
	price := CalculateBasePrice(trip)
	fmt.Printf("Базовая цена поездки: %.2f\n", price)
}
