package handlers

import "fmt"

type AdminBot struct{}

func (b AdminBot) HandleMessage(ctx BotContext) {
    text := ctx.Update["text"].(string)
    fmt.Printf("[AdminBot] Получено: %s\n", text)
}