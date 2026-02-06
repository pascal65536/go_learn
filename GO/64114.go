
package main

import "fmt"
import "time"	

func main() {
	var input string
	fmt.Scanln(&input)

    timeGo, err := time.Parse("2006-01-02/15:04:05", input)
    if err != nil {
        fmt.Println("Ошибка парсинга времени:", err)
        return
    }

	resultMessage := fmt.Sprintf("Текущее время %d часов, %d минут. Ты точно не забыл про важные дела на сегодня?", timeGo.Hour(), timeGo.Minute())
	fmt.Println(resultMessage)	
}