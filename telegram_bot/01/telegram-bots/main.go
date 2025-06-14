package main

import (
    "fmt"
    "log"
    "strconv"
    "os"
    "os/signal"
    "syscall"
    "time"

    "telegram-bots/utils"
)


func main() {
    me, err := utils.GetMe()
    if err != nil {
        log.Fatalf("Ошибка GetMe: %v", err)
    }
    fmt.Printf("Бот: @%s (ID: %d)\n", me.Result.Username, me.Result.ID)

    botIDStr := strconv.FormatInt(me.Result.ID, 10)
    folderName := "data"
    fileName := "update.json"

    processedIDs, data := loadProcessedUpdates(folderName, fileName, botIDStr)

    stopChan := setupSignalHandler()

    fmt.Println("Начинаем опрос обновлений. Нажмите Ctrl+C для остановки.")

    runUpdateLoop(botIDStr, folderName, fileName, processedIDs, data, stopChan)

    fmt.Println("Бот остановлен.")
}


func handler_message(upd utils.Update) {
    fmt.Printf("handler_message \n")
    fmt.Printf("upd.UpdateID = %v\n", upd.UpdateID)
    fmt.Printf("upd.Message.MessageID = %v\n", upd.Message.MessageID)
    fmt.Printf("upd.Message.Date = %v\n", upd.Message.Date)
    fmt.Printf("upd.Message.From = %v\n", upd.Message.From)
    fmt.Printf("upd.Message.Chat = %v\n", upd.Message.Chat)
    fmt.Printf("upd.Message.Text = %v\n", upd.Message.Text)
    fmt.Printf("upd.Message.Photo = %v\n", upd.Message.Photo)
    fmt.Printf("upd.Message.Video = %v\n", upd.Message.Video)
    fmt.Printf("upd.Message.Document = %v\n", upd.Message.Document)
    fmt.Printf("upd.Message.Caption = %v\n", upd.Message.Caption)
    fmt.Printf("upd.Message.Location = %v\n", upd.Message.Location)
    fmt.Printf("upd.Message.Venue = %v\n", upd.Message.Venue)
    fmt.Printf("upd.Message.Voice = %v\n", upd.Message.Voice)
    fmt.Printf("upd.Message.Sticker = %v\n", upd.Message.Sticker)
    fmt.Printf("upd.Message.Dice = %v\n", upd.Message.Dice)
    fmt.Printf("upd.Message.VideoNote = %v\n", upd.Message.VideoNote)
    fmt.Printf("upd.Message.PinnedMessage = %v\n", upd.Message.PinnedMessage)
    fmt.Printf("\n")    

    var text string
    var replyMarkup interface{}

    if upd.Message.Text == "/start" {
        text = "Привет, " + upd.Message.From.FirstName + "!\n\nЭтот бот умеет:\n/start - выводит это сообщение\n"
        // Удаляем клавиатуру
        replyMarkup = map[string]bool{
            "remove_keyboard": true,
        }
    } else if upd.Message.Text == "Кнопка 1" || upd.Message.Text == "Кнопка 2" || upd.Message.Text == "Кнопка 3" {
        text = upd.Message.Text + " ... Кнопка нажата"
        // Инлайновые кнопки
        replyMarkup = map[string]interface{}{
            "inline_keyboard": [][]map[string]string{
                {
                    {"text": "Кнопка 1", "callback_data": "btn1"},
                    {"text": "Кнопка 2", "callback_data": "btn2"},
                },
                {
                    {"text": "Кнопка 3", "callback_data": "btn3"},
                },
            },
        }        
    } else {
        text = upd.Message.Text
        // Формируем клавиатуру
        replyMarkup = map[string]interface{}{
            "keyboard": [][]map[string]string{
                {
                    {"text": "Кнопка 1"},
                    {"text": "Кнопка 2"},
                },
                {
                    {"text": "Кнопка 3"},
                },
            },
            "resize_keyboard":  true,
            "one_time_keyboard": true,
        }
    }

    _, err := utils.SendMessage(upd.Message.Chat.ID, text, replyMarkup)
    if err != nil {
        fmt.Printf("Ошибка SendMessage: %v\n", err)
    }
}

func handler_my_chat_member(upd utils.Update) {
    fmt.Printf("handler_channel_post \n")
    fmt.Printf("upd.UpdateID = %v\n", upd.UpdateID)
    fmt.Printf("upd.MyChatMember.Chat.ID = %v\n", upd.MyChatMember.Chat.ID)
    fmt.Printf("upd.MyChatMember = %v\n", upd.MyChatMember)
    fmt.Printf("\n")    
    fmt.Printf("\n")    
}


func handler_channel_post(upd utils.Update) {
    fmt.Printf("handler_channel_post \n")
    fmt.Printf("upd.UpdateID = %v\n", upd.UpdateID)
    fmt.Printf("upd.ChannelPost.Chat.ID = %v\n", upd.ChannelPost.Chat.ID)
    fmt.Printf("upd.ChannelPost.MessageID = %v\n", upd.ChannelPost.MessageID)
    fmt.Printf("upd.ChannelPost.Date = %v\n", upd.ChannelPost.Date)
    fmt.Printf("upd.ChannelPost.From = %v\n", upd.ChannelPost.From)
    fmt.Printf("upd.ChannelPost.Chat = %v\n", upd.ChannelPost.Chat)
    fmt.Printf("upd.ChannelPost.Text = %v\n", upd.ChannelPost.Text)
    fmt.Printf("upd.ChannelPost.Photo = %v\n", upd.ChannelPost.Photo)
    fmt.Printf("upd.ChannelPost.Video = %v\n", upd.ChannelPost.Video)
    fmt.Printf("upd.ChannelPost.Document = %v\n", upd.ChannelPost.Document)
    fmt.Printf("upd.ChannelPost.Caption = %v\n", upd.ChannelPost.Caption)
    fmt.Printf("upd.ChannelPost.Location = %v\n", upd.ChannelPost.Location)
    fmt.Printf("upd.ChannelPost.Venue = %v\n", upd.ChannelPost.Venue)
    fmt.Printf("upd.ChannelPost.Voice = %v\n", upd.ChannelPost.Voice)
    fmt.Printf("upd.ChannelPost.Sticker = %v\n", upd.ChannelPost.Sticker)
    fmt.Printf("upd.ChannelPost.Dice = %v\n", upd.ChannelPost.Dice)
    fmt.Printf("upd.ChannelPost.VideoNote = %v\n", upd.ChannelPost.VideoNote)
    fmt.Printf("upd.ChannelPost.PinnedMessage = %v\n", upd.ChannelPost.PinnedMessage)
    fmt.Printf("\n")
}

func handler_callback_query(upd utils.Update) {
    fmt.Printf("handler_callback_query \n")
    fmt.Printf("upd.UpdateID = %v\n", upd.UpdateID)
    fmt.Printf("upd.CallbackQuery.ID = %v\n", upd.CallbackQuery.ID)
    fmt.Printf("upd.CallbackQuery.From = %v\n", upd.CallbackQuery.From)
    fmt.Printf("upd.CallbackQuery.Message.Text = %v\n", upd.CallbackQuery.Message.Text)
    fmt.Printf("upd.CallbackQuery.Message.Chat = %v\n", upd.CallbackQuery.Message.Chat)
    fmt.Printf("upd.CallbackQuery.ChatInstance = %v\n", upd.CallbackQuery.ChatInstance)
    fmt.Printf("upd.CallbackQuery.Data = %v\n", upd.CallbackQuery.Data)
    fmt.Printf("\n")
}

// Загрузка обработанных update_id из файла
func loadProcessedUpdates(folderName, fileName, botIDStr string) (map[int64]bool, map[string]interface{}) {
    data := utils.LoadJSON(folderName, fileName)

    processedIDsRaw, ok := data[botIDStr]
    if !ok {
        processedIDsRaw = []interface{}{}
    }
    processedIDsSlice, ok := processedIDsRaw.([]interface{})
    if !ok {
        processedIDsSlice = []interface{}{}
    }

    processedIDs := make(map[int64]bool)
    for _, v := range processedIDsSlice {
        if idFloat, ok := v.(float64); ok {
            processedIDs[int64(idFloat)] = true
        }
    }
    return processedIDs, data
}

// Настройка перехвата сигнала остановки (Ctrl+C)
func setupSignalHandler() chan os.Signal {
    stopChan := make(chan os.Signal, 1)
    signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
    return stopChan
}

// Основной цикл обработки обновлений
func runUpdateLoop(botIDStr, folderName, fileName string, processedIDs map[int64]bool, data map[string]interface{}, stopChan chan os.Signal) {
loop:
    for {
        select {
        case <-stopChan:
            fmt.Println("\nПолучен сигнал остановки. Завершаем работу...")
            break loop
        default:
            // Получаем максимальный offset из data
            maxOffset := int64(0)
            if val, ok := data[botIDStr]; ok {
                if arr, ok := val.([]interface{}); ok && len(arr) > 0 {
                    for _, v := range arr {
                        if f, ok := v.(float64); ok { // JSON числа приходят как float64
                            id := int64(f)
                            if id > maxOffset {
                                maxOffset = id
                            }
                        }
                    }
                }
            }

            updates, err := utils.GetUpdates(maxOffset + 1)
            if err != nil {
                log.Printf("Ошибка GetUpdates: %v", err)
                time.Sleep(5 * time.Second)
                continue
            }

            processUpdates(updates.Result, processedIDs)

            saveProcessedUpdates(botIDStr, folderName, fileName, processedIDs, data)

            time.Sleep(3 * time.Second)
        }
    }
}


// Обработка полученных обновлений
func processUpdates(updates []utils.Update, processedIDs map[int64]bool) {
    for _, upd := range updates {
        if processedIDs[upd.UpdateID] {
            continue // уже обработан
        }

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

        processedIDs[upd.UpdateID] = true
    }
}

// Сохранение обновлённого списка обработанных update_id
func saveProcessedUpdates(botIDStr, folderName, fileName string, processedIDs map[int64]bool, data map[string]interface{}) {
    allIDs := make([]int64, 0, len(processedIDs))
    for id := range processedIDs {
        allIDs = append(allIDs, id)
    }
    data[botIDStr] = allIDs

    err := utils.SaveJSON(folderName, fileName, data)
    if err != nil {
        log.Printf("Ошибка сохранения JSON: %v", err)
    } else {
        fmt.Printf("Сохранено %d обработанных update_id\n", len(allIDs))
    }
}
