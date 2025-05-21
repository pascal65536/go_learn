package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    // Создаем канал для получения сигналов
    sigChan := make(chan os.Signal, 1)

    // signal.Notify связывает сигналы с каналом
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    fmt.Println("Программа запущена. Нажмите Ctrl+C для выхода.")

    // Ждем сигнал
    receivedSignal := <-sigChan
    fmt.Printf("\nПолучен сигнал: %v\n", receivedSignal)

    // Здесь можно выполнить очистку
    fmt.Println("Выполняю очистку перед выходом...")
    time.Sleep(time.Second)
    fmt.Println("Завершаю работу.")
}
