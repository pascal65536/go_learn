package main

import "fmt"

// Типы ботов (можно расширять)
type CatBot struct {
    Token   string
    ChatID  int64
    Name    string
}

type DogBot struct {
    Token   string
    ChatID  int64
    Name    string
}

// Общий интерфейс для всех ботов
type Bot interface {
    SendMessage(text string)
}

// Реализация SendMessage для CatBot
func (b CatBot) SendMessage(text string) {
    sendTelegramMessage(b.Token, b.ChatID, text)
}

// Реализация SendMessage для DogBot
func (b DogBot) SendMessage(text string) {
    sendTelegramMessage(b.Token, b.ChatID, text)
}

// Вспомогательная функция для отправки сообщения
func sendTelegramMessage(token string, chatID int64, text string) {
    fmt.Printf("Отправка сообщения: %s %d %s\n\n\n", text, chatID, token)
}


func select_bot(bot_name string) Bot{
    bot_obj := Bot(nil)
    if bot_name == "CatBot" {
        fmt.Println("CatBot")
        bot_obj = CatBot{
            Name:   "Бот Анны",
            Token:  "token123",
            ChatID: 123,
        }
    } else if bot_name == "DogBot" {
        fmt.Println("DogBot")
        bot_obj = DogBot{
            Name:   "Бот Ивана",
            Token:  "token456",
            ChatID: 123456,
        }
    }
    return bot_obj
}

func main() {

    bots := []string{"DogBot", "CatBot"}
    for _, bot_name := range bots {
        // Создаём разных ботов
        bot_obj := select_bot(bot_name)
        fmt.Printf("Бот %v\n", bot_obj)

        text := "..."

        // Используем одинаково благодаря интерфейсу Bot
        bot_obj.SendMessage(text)
    }
}