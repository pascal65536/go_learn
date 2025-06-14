package handlers

import "fmt"

type CatBot struct{}

func (b CatBot) HandleMessage(ctx BotContext) {
    text := ctx.Update["text"].(string)
    fmt.Printf("[CatBot] Получено: %s\n", text)
}