package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 5 {
		var first string
		fmt.Scanln(&first)
		if first == "Go" {
			fmt.Println("Go!")
		} else {
			fmt.Println("Я знаю только Go!")
		}
		i++
	}
}
