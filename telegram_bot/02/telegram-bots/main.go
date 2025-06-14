package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "strconv"
    "syscall"
    "time"

    "telegram-bots/utils"
)

type ProcessedData struct {
    BotUpdates map[string]int64
}

func NewProcessedData() *ProcessedData {
    return &ProcessedData{
        BotUpdates: make(map[string]int64),
    }
}

func main() {
    me, err := utils.GetMe()
    if err != nil {
        log.Fatalf("Ошибка GetMe: %v", err)
    }
    fmt.Printf("Бот: @%s (ID: %d)\n", me.Result.Username, me.Result.ID)

    botIDStr := strconv.FormatInt(me.Result.ID, 10)
    folderName := "data"
    fileName := "update.json"

    data := NewProcessedData()
    data.BotUpdates[botIDStr] = loadProcessedUpdates(folderName, fileName, botIDStr)
    
    stopChan := setupSignalHandler()
    fmt.Println("Начинаем опрос обновлений. Нажмите Ctrl+C для остановки.")

    runUpdateLoop(botIDStr, folderName, fileName, data, stopChan)
    fmt.Println("Бот остановлен.")
}

func loadProcessedUpdates(folderName, fileName, botIDStr string) int64 {
    rawData := utils.LoadJSON(folderName, fileName)
    maxID := int64(0)

    // Получаем данные для текущего бота
    if botData, ok := rawData[botIDStr]; ok {
        if idsSlice, ok := botData.([]interface{}); ok {
            for _, v := range idsSlice {
                if id, ok := v.(float64); ok {
                    if int64(id) > maxID {
                        maxID = int64(id)
                    }
                }
            }
        }
    }
    
    return maxID
}

func saveProcessedUpdates(botIDStr, folderName, fileName string, data *ProcessedData) {
    // Конвертируем в формат для сохранения
    rawData := make(map[string]interface{})
    for botID, lastID := range data.BotUpdates {
        rawData[botID] = []int64{lastID}
    }

    err := utils.SaveJSON(folderName, fileName, rawData)
    if err != nil {
        log.Printf("Ошибка сохранения JSON: %v", err)
    }
}

func setupSignalHandler() chan os.Signal {
    stopChan := make(chan os.Signal, 1)
    signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
    return stopChan
}

func runUpdateLoop(botIDStr, folderName, fileName string, data *ProcessedData, stopChan chan os.Signal) {
loop:
    for {
        select {
        case <-stopChan:
            fmt.Println("\nПолучен сигнал остановки. Завершаем работу...")
            break loop
        default:
            // Получаем текущий offset
            currentOffset := data.BotUpdates[botIDStr] + 1
            
            updates, err := utils.GetUpdates(currentOffset)
            if err != nil {
                log.Printf("Ошибка GetUpdates: %v", err)
                time.Sleep(5 * time.Second)
                continue
            }

            // Обработка обновлений и обновление максимального ID
            newMaxID := processUpdates(updates.Result, data.BotUpdates[botIDStr])
            if newMaxID > data.BotUpdates[botIDStr] {
                data.BotUpdates[botIDStr] = newMaxID
                saveProcessedUpdates(botIDStr, folderName, fileName, data)
            }

            time.Sleep(3 * time.Second)
        }
    }
}

func processUpdates(updates []utils.Update, currentMaxID int64) int64 {
    newMaxID := currentMaxID
    for _, upd := range updates {
        if upd.UpdateID <= currentMaxID {
            continue
        }

        // Обработка разных типов сообщений
        if upd.Message != nil {
            handler_message(upd)
        }
        if upd.MyChatMember != nil {
            handler_my_chat_member(upd)
        }
        if upd.ChannelPost != nil {
            handler_channel_post(upd)
        }
        if upd.CallbackQuery != nil {
            handler_callback_query(upd)
        }

        // Обновляем максимальный ID
        if upd.UpdateID > newMaxID {
            newMaxID = upd.UpdateID
        }
    }
    return newMaxID
}

func handler_message(upd utils.Update) {
    fmt.Printf("Новое сообщение: %s\n", upd.Message.Text)
    utils.ReplyMessage(upd.Message.Chat.ID, upd.Message.Text, upd.Message.MessageID, nil)
}

func handler_my_chat_member(upd utils.Update) {
    fmt.Printf("Изменение статуса в чате: %s\n", upd.MyChatMember.NewChatMember.Status)
}

func handler_channel_post(upd utils.Update) {
    fmt.Printf("Пост в канале: %s\n", upd.ChannelPost.Text)
}

func handler_callback_query(upd utils.Update) {
    fmt.Printf("Callback query: %s\n", upd.CallbackQuery.Data)
}
