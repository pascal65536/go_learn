
package main

import "fmt"

func main() {
    bots := []string{"DogBot", "CatBot"}
    for _, botName := range bots {
        // Создаём разных ботов
        bot := selectBot(botName)
        if bot == nil {
            fmt.Printf("Бот %s не найден\n", botName)
            continue
        }
        fmt.Printf("Бот: %s\n", botName)
        // text := "Привет от бота!"
        // Используем одинаково благодаря интерфейсу Bot
        updates, err := bot.GetUpdatesBot(0)
        if err != nil {
            fmt.Println("Ошибка:", err)
            continue
        }
        fmt.Printf("updates: %+v\n", updates)


        resp, err := bot.GetMeBot()
        if err != nil {
            fmt.Println("err:", err)
            continue
        }
        fmt.Printf("resp: %+v\n", resp)        
    }
}

func selectBot(botName string) Bot {
    var bot Bot
    if botName == "CatBot" {
        bot = CatBot{
            Name:   "Бот Анны",
            Token:  "token123",
            ChatID: 123,
        }
    } else if botName == "DogBot" {
        bot = DogBot{
            Name:   "Бот Ивана",
            Token:  "token456",
            ChatID: 123456,
        }
    }
    return bot
}