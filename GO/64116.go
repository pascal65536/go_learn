package main

import "fmt"
import "time"


func main() {
	var end string
	fmt.Scanln(&end)

    timeEnd, err := time.Parse("02.01.2006", end)
    if err != nil {
        fmt.Println("Ошибка парсинга времени:", err)
        return
    }

	timeOff := timeEnd.Add(time.Duration(15 * 24 * time.Hour)) // Добавляем 15 дней


	var firstname string
	fmt.Scanln(&firstname)
	
	var surname string
	fmt.Scanln(&surname)
	
	var lasname string
	fmt.Scanln(&lasname)

	var money1 float64
	fmt.Scanln(&money1)
	
	var money2 float64
	fmt.Scanln(&money2)
	
	var money3 float64
	fmt.Scanln(&money3)

	total := money1 + money2 + money3
	rubles := int(total)
	kop := int((total - float64(rubles)) * 100)

	resultMessage := fmt.Sprintf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\nДата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\nОбщая сумма выплат составит %d руб. %d коп.\n\nС уважением,\nГл. бух. Иванов А.Е.", surname, firstname, lasname, timeOff.Format("02.01.2006"), rubles, kop)
	fmt.Println(resultMessage)
}

/*
10.04.2005
Андрей
Иванов
Валерьевич
15000
19999
123
*/