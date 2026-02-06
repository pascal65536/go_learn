package main

import "fmt"

func PrintFlightRow(rase string, from string, to string, fly float64, stand int, gate int, isDirect bool) {
	if isDirect {
		result := fmt.Sprintf("| %s | %s—%s | регистрация закончилась, проходите к гейту: %d | длительность полёта %.1f часа |", rase, from, to, gate, fly)
		fmt.Println(result)		
	} else {
		result := fmt.Sprintf("| %s | %s—%s | %d регистрация продолжается |", rase, from, to, stand)
		fmt.Println(result)	
	}
}


func main() {
	PrintFlightRow("117B", "Москва", "Казань", 2, 115, 3, false)
	PrintFlightRow("117B", "Москва", "Казань", 2, 115, 3, true)
}


/*
| 117B | Москва—Казань | 115 регистрация продолжается |
| 117B | Москва—Казань | регистрация закончилась, проходите к гейту: 3 | длительность полёта 2.0 часа |
*/