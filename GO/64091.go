package main

import "fmt"

func main() {
	var name string
	fmt.Scanln(&name)
	
	var flat float64
	fmt.Scanln(&flat)
	
	var password float64
	fmt.Scanln(&password)
	
	var time float64
	fmt.Scanln(&time)

	resultMessage := fmt.Sprintf("Привет, %s! Хочу пригласить тебя на соревнование по программированию, которое пройдет, как всегда, в квартире %.0f. Оно будет длиться примерно %.1f часа. Не забудь секретный пароль для входа: %.0f.", name, flat, time, password)
	fmt.Println(resultMessage)
}
