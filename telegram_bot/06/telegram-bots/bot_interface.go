package main

type Bot interface {
    SendMessageBot(text string) (map[string]interface{}, error)
    GetWebhookInfoBot()
    DeleteWebhookBot()
    SetWebhookBot(webhookUrl string)
    GetMeBot() (*GetMeResponse, error)
    GetUpdatesBot(offset int64) (*GetUpdatesResponse, error)
}
