package utils

import (
    "fmt"
    "net/url"

    "telegram-bots/types"
)

func LoadBotConfigs(path string) map[string]types.BotConfig {
    return map[string]types.BotConfig{
        "bot_cat": {
            Token: "cat_token_123",
            Class: "CatBot",
            Settings: map[string]interface{}{
                "greeting": "Мяу! Привет!",
            },
        },
        "bot_dog": {
            Token: "dog_token_456",
            Class: "DogBot",
            Settings: map[string]interface{}{
                "greeting": "Гав! Привет!",
            },
        },
        "bot_admin": {
            Token: "admin_token_789",
            Class: "AdminBot",
            Settings: map[string]interface{}{
                "access_level": "high",
            },
        },
    }
}

func SendTelegramMessage(token, chatID, text string) {
    baseURL := "https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s"
    escapedText := url.QueryEscape(text)
    reqURL := fmt.Sprintf(baseURL, token, chatID, escapedText)
    fmt.Println("Отправка:", reqURL)
}