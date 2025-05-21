package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    done := make(chan bool)

    // Горутина, которая выполняет основную работу
    go func() {
        for {
            fmt.Println("Работаю...")
            time.Sleep(500 * time.Millisecond)
        }
    }()

    // Обработка сигналов
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    <-sigChan
    fmt.Println("Получен сигнал завершения")
    // Здесь можно остановить горутины, закрыть файлы/сокеты и т.п.
    close(done)
}
