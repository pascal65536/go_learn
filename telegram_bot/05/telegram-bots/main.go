package main

import (
    "log"
    "time"

    "telegram_bot/handlers"
    "telegram_bot/types"
    "telegram_bot/utils"
)

var botHandlers = map[string]handlers.BotHandler{}
var botConfigs = map[string]types.BotConfig{}

func init() {
    handlers.RegisterHandlers(botHandlers)
}

func main() {
    botConfigs = utils.LoadBotConfigs("config/bots.yaml")
    if len(botConfigs) == 0 {
        log.Fatal("Не найдено ни одного бота в конфиге")
    }

    log.Println("Запуск ботов...")

    for name, cfg := range botConfigs {
        handler, ok := botHandlers[cfg.Class]
        if !ok {
            log.Printf("Хендлер не найден для бота: %s (%s)", name, cfg.Class)
            continue
        }

        log.Printf("Бот '%s' (%s) запущен", name, cfg.Class)

        // Создаем копию переменной для горутины
        name := name
        cfg := cfg

        go func() {
            offset := int64(0)

            for {
                // Получаем обновления от Telegram
                updates, err := utils.GetUpdates(cfg.Token, offset)
                if err != nil {
                    log.Printf("[%s] Ошибка получения обновлений: %v", name, err)
                    time.Sleep(5 * time.Second)
                    continue
                }

                for _, update := range updates.Result {
                    offset = update.UpdateID
                    ctx := handlers.BotContext{
                        Name:     name,
                        Token:    cfg.Token,
                        Settings: cfg.Settings,
                        Update:   update, // а не ConvertUpdateToMap(update)
                        Handler:  handler,
                    }
                    // Вместо handler.HandleMessage(ctx):
                    handler.HandleUpdate(ctx)
                }

            }
        }()
    }

    select {} // держим основную горутину живой
}