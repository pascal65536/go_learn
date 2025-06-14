package main

type CatBot struct {
    Name   string
    Token  string
    ChatID int64
}

func (b CatBot) SendMessageBot(text string) (map[string]interface{}, error) {
    result, err := SendMessage(b.Token, b.ChatID, text, nil)
    return result, err
}

func (b CatBot) GetWebhookInfoBot() {
    GetWebhookInfo(b.Token)
}

func (b CatBot) DeleteWebhookBot() {
    DeleteWebhook(b.Token)
}

func (b CatBot) SetWebhookBot(webhookUrl string) {
    SetWebhook(b.Token, webhookUrl)
}

func (b CatBot) GetMeBot() (*GetMeResponse, error) {
    resp, err := GetMe(b.Token)
    return resp, err
}

func (b CatBot) GetUpdatesBot(offset int64) (*GetUpdatesResponse, error) {
    updates, err := GetUpdates(b.Token, offset)
    return updates, err
}
