package main

import "fmt"
import "time"

func main() {
	var end string
	fmt.Scanln(&end)
	var begin string
	fmt.Scanln(&begin)

    timeEnd, err := time.Parse("2006-01-02", end)
    if err != nil {
        fmt.Println("Ошибка парсинга времени:", err)
        return
    }
    timeBegin, err := time.Parse("2006-01-02", begin)
    if err != nil {
        fmt.Println("Ошибка парсинга времени:", err)
        return
    }	

	resultMessage := fmt.Sprintf("%d year ago", timeEnd.Year() - timeBegin.Year())
	fmt.Println(resultMessage)
}