package main

import "fmt"


type Bot1 struct{}

func (b *Bot1) OnMessage(text string) {
    fmt.Printf("[CatBot] Получено сообщение: %s\n", text)
    // здесь можно отправить ответ через Telegram API
}
