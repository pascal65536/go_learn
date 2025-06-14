package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

// go get github.com/joho/godotenv

func main() {
    // Загружаем переменные окружения из файла .env
    err := godotenv.Load()
    if err != nil {
        log.Println("Файл .env не найден, переменные окружения не загружены")
    }

    // Получаем значение переменной TOKEN
    token := os.Getenv("TOKEN")
    if token == "" {
        log.Fatal("Токен не найден в переменных окружения")
    }

    fmt.Println("Токен из .env:", token)
}
