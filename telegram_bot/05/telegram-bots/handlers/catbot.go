package handlers

import "log"

type CatBot struct {
    BaseBot
}

func (CatBot) HandleMessage(ctx BotContext) {
    if ctx.Update.Message != nil {
        text := ctx.Update.Message.Text
        log.Println("[CatBot] Text:", text)
    }
}

func (CatBot) HandleMyChatMember(ctx BotContext) {
    log.Println("[CatBot] MyChatMember event %v", ctx)
}

func (CatBot) HandleChannelPost(ctx BotContext) {
    log.Println("[CatBot] ChannelPost event %v", ctx)
}

func (CatBot) HandleCallbackQuery(ctx BotContext) {
    log.Println("[CatBot] CallbackQuery event %v", ctx)
}

func (h CatBot) HandleUpdate(ctx BotContext) {
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
