package main

import (
	"fmt"
)

func main() {
	var counter int
	fmt.Scanln(&counter)

	i := 0
	for i < counter {
		var first float64
		fmt.Scanln(&first)

		if first <= 100 && first >= 90 {
			fmt.Println(5)
		} else if first < 90 && first >= 75 {
			fmt.Println(4)
		} else if first < 75 && first >= 50 {
			fmt.Println(3)
		} else if first < 50 && first >= 0 {
			fmt.Println(2)			
		} else {
			fmt.Println("Неверный балл")
		}

		i++
	}
}
