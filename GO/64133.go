
package main

import "fmt"

func PrintFlightRow(first string) {
	if first == "Go" {
		fmt.Println("Go!")
	} else {
		fmt.Println("Я знаю только Go!")
	}
}


func main() {
	GoOrNot("12")
}