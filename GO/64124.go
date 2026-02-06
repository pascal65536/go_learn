package main

import (
	"fmt"
)

func main() {
	var first string
	fmt.Scanln(&first)
	
	var second string
	fmt.Scanln(&second)

	if first == second {
		fmt.Println("Ничья")
	} else if first == "камень" && second == "ножницы" || first == "ножницы" && second == "бумага" || first == "бумага" && second == "камень" {
		fmt.Println("Первый игрок победил")
	} else if second == "камень" && first == "ножницы" || second == "ножницы" && first == "бумага" || second == "бумага" && first == "камень" {
		fmt.Println("Второй игрок победил")
	} else {
		fmt.Println("Упс")
	}
}
