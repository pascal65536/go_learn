package main

import (
    "log"
    "os"
)

func main() {
    token := "ваш_секретный_токен123"
    f, err := os.Create(".env")
    if err != nil {
        log.Fatalf("Ошибка создания .env: %v", err)
    }
    defer f.Close()
    _, err = f.WriteString("TOKEN=" + token + "\n")
    if err != nil {
        log.Fatalf("Ошибка записи в .env: %v", err)
    }
    log.Println("Токен сохранён в .env")
}