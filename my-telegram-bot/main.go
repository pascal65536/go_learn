package main

import (
    "log"
    "os"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
    // Получаем токен из переменной окружения или напрямую
    token := os.Getenv("YOUR_BOT_TOKEN")
    if token == "" {
        log.Fatal("TELEGRAM_BOT_TOKEN is not set")
    }

    // Создаем новый бот
    bot, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Authorized on account %s", bot.Self.UserName)

    // Устанавливаем обновления
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)

    // Обрабатываем обновления
    for update := range updates {
        if update.Message == nil { // ignore non-message updates
            continue
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        // Отправляем ответ
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы написали: "+update.Message.Text)
        if _, err := bot.Send(msg); err != nil {
            log.Println(err)
        }
    }
}

