package main

import "errors"

var DivisionByZeroError = errors.New("division by zero is not allowed")

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, DivisionByZeroError
	}
	return float64(a) / float64(b), nil
}



import "fmt"

func main() {
	// Пример использования функции
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}

	result, err = Divide(10, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}