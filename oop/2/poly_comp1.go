package main

import (
    "fmt"
    "net/url"
)

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
    baseURL := "https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s"
    escapedText := url.QueryEscape(text)
    url := fmt.Sprintf(baseURL, token, chatID, escapedText)
    fmt.Println("Отправка сообщения:", url)
}

func main() {
    
    // Создаём разных ботов
    bot1 := CatBot{
        Name:   "Бот Анны",
        Token:  "token123",
        ChatID: 123,
    }
    bot2 := DogBot{
        Name:   "Бот Ивана",
        Token:  "token456",
        ChatID: 123456,
    }

    // Используем одинаково благодаря интерфейсу Bot
    bot1.SendMessage("Мяу!")
    bot2.SendMessage("Гав!")
}