package main

import "fmt"
import "math"

func main() {
	var x float64
	fmt.Scanln(&x)
	var y float64
	fmt.Scanln(&y)
	var z float64
	fmt.Scanln(&z)
	var m float64
	fmt.Scanln(&m)
	var n float64
	fmt.Scanln(&n)
	fmt.Println((5*x) * (math.Sin(2*y)) / (z + math.Pow(n, math.Log(m))))
}