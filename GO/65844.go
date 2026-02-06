package main

import "fmt"

const (
	minPrice = 99.0
	maxPrice = 20000.0
)

func ApplyPriceLimits(price float64) float64 {
	if price < minPrice {
		return minPrice
	}
	if price > maxPrice {
		return maxPrice
	}
	return price
}

func main() {
	testPrices := []float64{50, 99, 5000, 20000, 25000}
	for _, p := range testPrices {
		limited := ApplyPriceLimits(p)
		fmt.Printf("Original: %.2f, Limited: %.2f\n", p, limited)
	}
}
