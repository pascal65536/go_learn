package main

import (
	"fmt"
)

func main() {
	var counter int
	fmt.Scanln(&counter)

	var tax int
	fmt.Scanln(&tax)

	i := 0
    var summa float64
    summa = 0
	for i < counter {
		var first float64
		fmt.Scanln(&first)
        summa = summa + (first - float64(first * (float64(tax) / 100)))
		i++
	}

    fmt.Println(summa)
}
