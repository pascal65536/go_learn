package main

import "fmt"

type Bot2 struct{}

func (b *Bot2) OnMessage(text string) {
    fmt.Printf("[DogBot] Получено сообщение: %s\n", text)
    // здесь можно выполнить команду или отправить ответ
}