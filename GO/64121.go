package main

import "fmt"
import "math"

func main() {
	var a float64
	fmt.Scanln(&a)
	if a == 0 {
		fmt.Println("Число 0")		
	} else if a > -10 && a < 10 {
		fmt.Println("Число однозначное")
	} else if math.Floor(a / 2) == a / 2 {
		fmt.Println("Число чётное")
	} else if a > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число красивое")
	}
}