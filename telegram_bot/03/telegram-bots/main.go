package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/joho/godotenv"

    "telegram-bots/utils"

)

func getToken() (string, error) {
    _ = godotenv.Load()
    token := os.Getenv("TOKEN")
    if token == "" {
        return "", fmt.Errorf("токен не найден в переменных окружения")
    }
    return token, nil
}

// Интерфейс для хендлеров бота
type BotHandler interface {
    HandlerMessage(upd utils.Update)
    HandlerMyChatMember(upd utils.Update)
    HandlerChannelPost(upd utils.Update)
    HandlerCallbackQuery(upd utils.Update)
}

// Минималистичный бот с токеном и хендлерами
type Bot struct {
    Token   string
    Handler BotHandler
}

func main() {

    token, _ := getToken()

    bot := Bot{
        Token:   token,
        Handler: &MyBotHandlers{Token: token},
    }

    stopChan := make(chan os.Signal, 1)
    signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

    fmt.Println("Бот запущен. Нажмите Ctrl+C для остановки.")

    go RunBot(&bot, stopChan)

    <-stopChan
    fmt.Println("Бот остановлен.")
}

func RunBot(bot *Bot, stopChan chan os.Signal) {
    offset := int64(0)

    for {
        select {
        case <-stopChan:
            return
        default:
            updatesResp, err := utils.GetUpdates(bot.Token, offset+1)
            if err != nil {
                log.Printf("Ошибка GetUpdates: %v", err)
                time.Sleep(5 * time.Second)
                continue
            }

            for _, upd := range updatesResp.Result {
                offset = upd.UpdateID

                if upd.Message != nil {
                    bot.Handler.HandlerMessage(upd)
                }
                if upd.MyChatMember != nil {
                    bot.Handler.HandlerMyChatMember(upd)
                }
                if upd.ChannelPost != nil {
                    bot.Handler.HandlerChannelPost(upd)
                }
                if upd.CallbackQuery != nil {
                    bot.Handler.HandlerCallbackQuery(upd)
                }
            }
        }
    }
}

// Пример реализации хендлеров бота
type MyBotHandlers struct {
    Token string
}

func (h *MyBotHandlers) HandlerMessage(upd utils.Update) {
    fmt.Println("Новое сообщение:", upd.Message.Text)
    err := utils.ReplyMessage(h.Token, upd.Message.Chat.ID, "Вы написали: "+upd.Message.Text, upd.Message.MessageID)
    if err != nil {
        log.Printf("Ошибка отправки ответа: %v", err)
    }
}

func (h *MyBotHandlers) HandlerMyChatMember(upd utils.Update) {
    fmt.Println("Изменение статуса в чате:", upd.MyChatMember.NewChatMember.Status)
}

func (h *MyBotHandlers) HandlerChannelPost(upd utils.Update) {
    fmt.Println("Пост в канале:", upd.ChannelPost.Text)
}

func (h *MyBotHandlers) HandlerCallbackQuery(upd utils.Update) {
    fmt.Println("Callback query:", upd.CallbackQuery.Data)
}
