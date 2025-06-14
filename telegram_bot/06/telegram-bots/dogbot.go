package main

type DogBot struct {
    Name   string
    Token  string
    ChatID int64
}

func (b DogBot) SendMessageBot(text string) (map[string]interface{}, error) {
    result, err := SendMessage(b.Token, b.ChatID, text, nil)
    return result, err
}

func (b DogBot) GetWebhookInfoBot() {
    GetWebhookInfo(b.Token)
}

func (b DogBot) DeleteWebhookBot() {
    DeleteWebhook(b.Token)
}

func (b DogBot) SetWebhookBot(webhookUrl string) {
    SetWebhook(b.Token, webhookUrl)
}

func (b DogBot) GetMeBot() (*GetMeResponse, error) {
    resp, err := GetMe(b.Token)
    return resp, err
}

func (b DogBot) GetUpdatesBot(offset int64) (*GetUpdatesResponse, error) {
    updates, err := GetUpdates(b.Token, offset)
    return updates, err
}