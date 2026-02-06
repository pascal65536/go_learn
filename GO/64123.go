package main

import (
	"fmt"
)

func main() {
	var first string
	fmt.Scanln(&first)
	
	var second string
	fmt.Scanln(&second)

	if len(first) >= 8 && len(second) >= 8 {
		fmt.Println("Оба пароля надёжные")
	} else if len(first) < 8 && len(second) < 8 {
			fmt.Println("Оба пароля ненадёжные")
	} else if len(first) >= 8 && len(second) < 8 {
		fmt.Println("Только первый пароль надёжный")
	} else if len(first) < 8 && len(second) >= 8 {
		fmt.Println("Только второй пароль надёжный")
	} else {
		fmt.Println("Упс")
	}
}
