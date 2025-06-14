package handlers

import "fmt"

type DogBot struct{}

func (b DogBot) HandleMessage(ctx BotContext) {
    text := ctx.Update["text"].(string)
    fmt.Printf("[DogBot] Получено: %s\n", text)
}