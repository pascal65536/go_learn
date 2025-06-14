type MyBot struct {
    Token string
}

func (b *MyBot) HandlerMessage(upd Update) {
    fmt.Println("Новое сообщение:", upd.Message.Text)
    utils.ReplyMessage(b.Token, upd.Message.Chat.ID, "Ответ: "+upd.Message.Text, upd.Message.MessageID)
}

func (b *MyBot) HandlerMyChatMember(upd Update) {
    // обработка
}

func (b *MyBot) HandlerChannelPost(upd Update) {
    // обработка
}

func (b *MyBot) HandlerCallbackQuery(upd Update) {
    // обработка
}
