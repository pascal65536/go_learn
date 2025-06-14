package handlers

import "log"

type AdminBot struct {
    BaseBot
}

func (AdminBot) HandleMessage(ctx BotContext) {
    if ctx.Update.Message != nil {
        text := ctx.Update.Message.Text
        log.Println("[AdminBot] Text:", text)
    }
}

func (AdminBot) HandleMyChatMember(ctx BotContext) {
    log.Println("[AdminBot] MyChatMember event %v", ctx)
}

func (AdminBot) HandleChannelPost(ctx BotContext) {
    log.Println("[AdminBot] ChannelPost event %v", ctx)
}

func (AdminBot) HandleCallbackQuery(ctx BotContext) {
    log.Println("[AdminBot] CallbackQuery event %v", ctx)
}

func (h AdminBot) HandleUpdate(ctx BotContext) {
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
