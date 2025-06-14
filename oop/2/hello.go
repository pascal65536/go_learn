package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Создаем сканер для чтения ввода с консоли
    reader := bufio.NewReader(os.Stdin)

    // Запрашиваем имя у пользователя
    fmt.Print("Введите ваше имя: ")
    name, _ := reader.ReadString('\n')

    // Убираем лишние пробелы и символы новой строки
    name = strings.TrimSpace(name)

    // Выводим приветствие
    fmt.Printf("Привет, %s! Добро пожаловать в мир Go!\n", name)
}
