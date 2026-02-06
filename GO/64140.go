package main

import "fmt"

// Функция для проверки корректности данных
func Scan(age int, name string, passportSeriesAndNumber string) (bool, string) {
	// Проверка валидности возраста
	if age < 14 || age > 150 {
		return false, "Ошибка: невалидный возраст"
	}
	// Проверка длины имени
	if len(name) < 2 {
		return false, "Ошибка: невалидное имя"
	}
	// Проверка длины серии и номера паспорта
	if len(passportSeriesAndNumber) != 10 {
		return false, "Ошибка: невалидная серия и номер паспорта"
	}
	return true, ""
}

func main() {
	var age int
	var name string
	var passportSeriesAndNumber string

	// Чтение данных пользователя
	n, err := fmt.Scanln(&age, &name, &passportSeriesAndNumber)
	if err != nil || n != 3 {
		fmt.Println("Ошибка: некорректный ввод")
		return
	}

	// Проверка данных через функцию Scan
	valid, errMsg := Scan(age, name, passportSeriesAndNumber)
	if !valid {
		fmt.Println(errMsg)
		return
	}

	// Если все проверки прошли успешно — выводим информацию
	fmt.Printf("Имя: %s. Возраст: %d. Серия и номер паспорта: %s\n", name, age, passportSeriesAndNumber)
}