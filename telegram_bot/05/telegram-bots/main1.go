package main

import (
    "encoding/json"
    "log"
    "net/http"
    "path/filepath"

    "telegram-bots/handlers"
    "telegram-bots/types"
    "telegram-bots/utils"
)

var botHandlers = map[string]handlers.BotHandler{}
var botConfigs = map[string]types.BotConfig{}

func init() {
    cfgPath := filepath.Join("config", "bots.yaml")
    botConfigs = utils.LoadBotConfigs(cfgPath)
    handlers.RegisterHandlers(botHandlers)
}

func main() {
    http.HandleFunc("/bot/", func(w http.ResponseWriter, r *http.Request) {
        nameBot := r.URL.Path[len("/bot/"):]

        cfg, ok := botConfigs[nameBot]
        if !ok {
            http.Error(w, "Bot not found", http.StatusNotFound)
            return
        }

        handler, ok := botHandlers[cfg.Class]
        if !ok {
            http.Error(w, "Handler not found", http.StatusInternalServerError)
            return
        }

        var update map[string]interface{}
        if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        ctx := handlers.BotContext{
            Name:     nameBot,
            Token:    cfg.Token,
            Settings: cfg.Settings,
            Update:   update,
            Handler:  handler,
        }

        handler.HandleMessage(ctx)
        utils.SendTelegramMessage(cfg.Token, "123456", "Получено: "+getMessage(update))
        w.WriteHeader(http.StatusOK)
    })

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getMessage(update map[string]interface{}) string {
    if text, ok := update["text"].(string); ok {
        return text
    }
    return "[unknown]"
}