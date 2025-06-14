package handlers

import "log"

type DogBot struct {
    BaseBot
}

func (DogBot) HandleMessage(ctx BotContext) {
    if ctx.Update.Message != nil {
        text := ctx.Update.Message.Text
        log.Println("[DogBot] Text:", text)
    }
}

func (DogBot) HandleMyChatMember(ctx BotContext) {
    log.Println("[DogBot] MyChatMember event %v", ctx)
}

func (DogBot) HandleChannelPost(ctx BotContext) {
    log.Println("[DogBot] ChannelPost event %v", ctx)
}

func (DogBot) HandleCallbackQuery(ctx BotContext) {
    log.Println("[DogBot] CallbackQuery event %v", ctx)
}

func (h DogBot) HandleUpdate(ctx BotContext) {
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
