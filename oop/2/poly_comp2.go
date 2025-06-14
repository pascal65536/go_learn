package main

import "fmt"

type Handler interface {
    SendMessage(TelegramBot)
}

type Dog struct{}

func (d Dog) SendMessage(s TelegramBot) {
    method := "sendMessage"
    url := fmt.Sprintf("https://api.telegram.org/bot%s/%s?chat_id=%d&text=%s", s.Token, method, s.ChatID, "Гав!")
    fmt.Println(url)
}

type Cat struct{}

func (c Cat) SendMessage(s TelegramBot) {
    method := "sendMessage"
    url := fmt.Sprintf("https://api.telegram.org/bot%s/%s?chat_id=%d&text=%s", s.Token, method, s.ChatID, "Мяу!")
    fmt.Println(url)
}

type TelegramBot struct {
    Name    string
    Token   string  // Токен бота из .env
    ChatID  int64   // ID чата из .env
    Handler Handler
}

func (p TelegramBot) Run() {
    p.Handler.SendMessage(p)
}

func main() {
    dog := Dog{}
    cat := Cat{}

    person1 := TelegramBot{Name: "Анна", Handler: cat, Token: "token123", ChatID: 123}
    person2 := TelegramBot{Name: "Иван", Handler: dog, Token: "token456", ChatID: 123456}

    person1.Run()
    person2.Run()
}
