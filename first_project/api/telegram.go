package api

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
 "os"
)

// Структура для данных из Telegram API
type Update struct {
 UpdateID int `json:"update_id"`
 Message  struct {
  MessageID int `json:"message_id"`
  Chat      struct {
   ID   int    `json:"id"`
   Type string `json:"type"`
  } `json:"chat"`
  Text string `json:"text"`
 } `json:"message"`
}

// Запись JSON в файл
func writeJSON(filename string, data interface{}) error {
 file, err := os.Create(filename)
 if err != nil {
  return fmt.Errorf("ошибка создания файла: %v", err)
 }
 defer file.Close()

 encoder := json.NewEncoder(file)
 encoder.SetIndent("", "  ")
 if err := encoder.Encode(data); err != nil {
  return fmt.Errorf("ошибка записи JSON: %v", err)
 }
 return nil
}

// Чтение JSON из файла
func readJSON(filename string, v interface{}) error {
 data, err := ioutil.ReadFile(filename)
 if err != nil {
  return fmt.Errorf("ошибка чтения файла: %v", err)
 }

 if err := json.Unmarshal(data, v); err != nil {
  return fmt.Errorf("ошибка парсинга JSON: %v", err)
 }
 return nil
}

func main() {
 // Замените YOUR_BOT_TOKEN на токен вашего бота
 token := "YOUR_BOT_TOKEN"
 url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", token)

 // Запрос к Telegram API
 resp, err := http.Get(url)
 if err != nil {
  fmt.Println("Ошибка API:", err)
  return
 }
 defer resp.Body.Close()

 // Чтение ответа
 body, err := ioutil.ReadAll(resp.Body)
 if err != nil {
  fmt.Println("Ошибка чтения ответа:", err)
  return
 }

 // Парсинг JSON
 var result struct {
  Ok     bool      `json:"ok"`
  Result []Update  `json:"result"`
 }
 if err := json.Unmarshal(body, &result); err != nil {
  fmt.Println("Ошибка парсинга JSON:", err)
  return
 }

 if !result.Ok {
  fmt.Println("API вернул ошибку")
  return
 }

 // Сохранение в файл
 if err := writeJSON("updates.json", result.Result); err != nil {
  fmt.Println(err)
  return
 }
 fmt.Println("Данные сохранены в updates.json")

 // Чтение из файла
 var loadedUpdates []Update
 if err := readJSON("updates.json", &loadedUpdates); err != nil {
  fmt.Println(err)
  return
 }

 // Вывод первых 2 обновлений
 for i, update := range loadedUpdates {
  if i >= 2 {
   break
  }
  fmt.Printf("Update ID: %d, Chat ID: %d, Text: %s\n",
   update.UpdateID, update.Message.Chat.ID, update.Message.Text)
 }
}
