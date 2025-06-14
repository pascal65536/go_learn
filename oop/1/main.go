package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// === Типы данных ===

type Update struct {
    BotID string `json:"bot_id"`
    Text  string `json:"text"`
    Type  string `json:"type"` // message, command, callback_query и т.д.
}

type BotConfig struct {
    ID    string
    Token string
    Name  string
    Class BotClass
}

type BotClass interface {
    OnMessage(text string)
}

// === Глобальные переменные ===

var bots = map[string]BotConfig{
    "cat123": {ID: "cat123", Token: "token_cat", Name: "Котя", Class: &Bot1{}},
    "dog456": {ID: "dog456", Token: "token_dog", Name: "Плюш", Class: &Bot2{}},
}

// === Дефолтные функции ===

func DefaultOnMessage(text string) {
    fmt.Println("Default handler: Неизвестное сообщение:", text)
}

// === HTTP сервер ===

func callbackHandler(w http.ResponseWriter, r *http.Request) {
    var update Update
    if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    botConfig, ok := bots[update.BotID]
    if !ok {
        http.Error(w, "Bot not found", http.StatusNotFound)
        return
    }

    switch update.Type {
    case "message":
        if handler, ok := botConfig.Class.(interface{ OnMessage(text string) }); ok {
            handler.OnMessage(update.Text)
        } else {
            DefaultOnMessage(update.Text)
        }
    default:
        fmt.Println("Unknown update type:", update.Type)
    }

    w.WriteHeader(http.StatusOK)
}

// === Main ===

func main() {
    http.HandleFunc("/callback", callbackHandler)
    fmt.Println("Server is running on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}
