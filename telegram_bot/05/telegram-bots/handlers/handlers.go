package handlers

import (
    "telegram_bot/utils"
)


func RegisterHandlers(handlersMap map[string]BotHandler) {
    handlersMap["CatBot"] = CatBot{}
    handlersMap["DogBot"] = DogBot{}
    handlersMap["AdminBot"] = AdminBot{}
}


// Базовая реализация (можно не реализовывать, если не нужно)
type BaseBot struct{}
func (BaseBot) HandleMessage(ctx BotContext)         {}
func (BaseBot) HandleMyChatMember(ctx BotContext)    {}
func (BaseBot) HandleChannelPost(ctx BotContext)     {}
func (BaseBot) HandleCallbackQuery(ctx BotContext)   {}

type BotHandler interface {
    HandleUpdate(ctx BotContext)
    HandleMessage(ctx BotContext)
    HandleMyChatMember(ctx BotContext)
    HandleChannelPost(ctx BotContext)
    HandleCallbackQuery(ctx BotContext)
}

type BotContext struct {
    Name     string
    Token    string
    Settings map[string]interface{}
    Update   utils.Update
    Handler  BotHandler
}


func (h BaseBot) HandleUpdate(ctx BotContext) {
    if ctx.Update.Message != nil {
        h.HandleMessage(ctx)
    }
    if ctx.Update.MyChatMember != nil {
        h.HandleMyChatMember(ctx)
    }
    if ctx.Update.ChannelPost != nil {
        h.HandleChannelPost(ctx)
    }
    if ctx.Update.CallbackQuery != nil {
        h.HandleCallbackQuery(ctx)
    }
}
