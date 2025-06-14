package utils

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "os"
    "path/filepath"

    "github.com/joho/godotenv"
)

func getToken() (string, error) {
    _ = godotenv.Load()
    token := os.Getenv("TOKEN")
    if token == "" {
        return "", fmt.Errorf("токен не найден в переменных окружения")
    }
    return token, nil
}

func postToTelegram(method string, payload interface{}, result interface{}) error {
    token, err := getToken()
    if err != nil {
        return err
    }
    url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, method)
    body, err := json.Marshal(payload)
    if err != nil {
        return err
    }
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    // Читаем тело ответа
    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    // Выводим тело ответа в консоль как строку
    // fmt.Println("Ответ Telegram API:", string(respBody))
    // Декодируем в result
    return json.Unmarshal(respBody, result)
}

func SetWebhook(webhookUrl string) (map[string]interface{}, error) {
    payload := map[string]string{"url": webhookUrl}
    var result map[string]interface{}
    err := postToTelegram("setWebhook", payload, &result)
    return result, err
}

func DeleteWebhook() (map[string]interface{}, error) {
    var result map[string]interface{}
    err := postToTelegram("deleteWebhook", map[string]interface{}{}, &result)
    return result, err
}

func GetWebhookInfo() (map[string]interface{}, error) {
    var result map[string]interface{}
    err := postToTelegram("getWebhookInfo", map[string]interface{}{}, &result)
    return result, err
}

type SendMessagePayload struct {
    ChatID      int64       `json:"chat_id"`
    Text        string      `json:"text"`
    ParseMode   string      `json:"parse_mode,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

func SendMessage(chatID int64, text string, replyMarkup interface{}) (map[string]interface{}, error) {
    payload := SendMessagePayload{
        ChatID:    chatID,
        Text:      text,
        ParseMode: "html",
    }
    if replyMarkup != nil {
        payload.ReplyMarkup = replyMarkup
    }
    var result map[string]interface{}
    err := postToTelegram("sendMessage", payload, &result)
    return result, err
}

type GetMeResponse struct {
    Ok     bool `json:"ok"`
    Result struct {
        ID                      int64  `json:"id"`
        IsBot                   bool   `json:"is_bot"`
        FirstName               string `json:"first_name"`
        Username                string `json:"username"`
        CanJoinGroups           bool   `json:"can_join_groups"`
        CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
        SupportsInlineQueries   bool   `json:"supports_inline_queries"`
    } `json:"result"`
}

func GetMe() (*GetMeResponse, error) {
    var resp GetMeResponse
    err := postToTelegram("getMe", map[string]interface{}{}, &resp)
    return &resp, err
}

type GetUpdatesResponse struct {
    Ok     bool     `json:"ok"`
    Result []Update `json:"result"`
}

type Update struct {
    UpdateID      int64           `json:"update_id"`
    Message       *Message        `json:"message,omitempty"`
    MyChatMember  *MyChatMember   `json:"my_chat_member,omitempty"`
    ChannelPost   *Message        `json:"channel_post,omitempty"`
    CallbackQuery *CallbackQuery  `json:"callback_query,omitempty"`
}

type CallbackQuery struct {
    ID              string  `json:"id"`
    From            User    `json:"from"`
    Message         Message `json:"message"`
    ChatInstance    string  `json:"chat_instance"`
    Data            string  `json:"data"`
}

type User struct {
    ID           int64  `json:"id"`
    IsBot        bool   `json:"is_bot"`
    FirstName    string `json:"first_name"`
    Username     string `json:"username"`
    LanguageCode string `json:"language_code"`
    IsPremium    bool   `json:"is_premium"`
}

// PhotoSize описывает размер фото или миниатюры (thumbnail)
type PhotoSize struct {
    FileID       string `json:"file_id"`
    FileUniqueID string `json:"file_unique_id"`
    Width        int    `json:"width"`
    Height       int    `json:"height"`
    FileSize     int    `json:"file_size,omitempty"`
}

// Video описывает видеофайл
type Video struct {
    FileID       string     `json:"file_id"`
    FileUniqueID string     `json:"file_unique_id"`
    Width        int        `json:"width"`
    Height       int        `json:"height"`
    Duration     int        `json:"duration"`
    Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // иногда thumb
    FileName     string     `json:"file_name,omitempty"`
    MimeType     string     `json:"mime_type,omitempty"`
    FileSize     int        `json:"file_size,omitempty"`
}

// Document описывает документ
type Document struct {
    FileID       string     `json:"file_id"`
    FileUniqueID string     `json:"file_unique_id"`
    FileName     string     `json:"file_name,omitempty"`
    MimeType     string     `json:"mime_type,omitempty"`
    FileSize     int        `json:"file_size,omitempty"`
}

type Location struct {
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}

type Venue struct {
    Location     Location `json:"location"`
    Title        string   `json:"title"`
    Address      string   `json:"address"`
    FoursquareID string   `json:"foursquare_id,omitempty"`
}

// Voice описывает голосовое сообщение
type Voice struct {
    FileID       string `json:"file_id"`
    FileUniqueID string `json:"file_unique_id"`
    Duration     int    `json:"duration"`               // длительность в секундах
    MimeType     string `json:"mime_type,omitempty"`    // MIME-тип, например "audio/ogg"
    FileSize     int    `json:"file_size,omitempty"`    // размер файла в байтах
}

type Sticker struct {
    Width      int    `json:"width"`
    Height     int    `json:"height"`
    Emoji      string `json:"emoji"`
    SetName    string `json:"set_name,omitempty"`
    IsAnimated bool   `json:"is_animated"`
    IsVideo    bool   `json:"is_video"`
    Type       string `json:"type"`
    Thumbnail  *PhotoSize `json:"thumbnail,omitempty"`
    Thumb      *PhotoSize `json:"thumb,omitempty"`
    FileID     string `json:"file_id"`
    FileUniqueID string `json:"file_unique_id"`
    FileSize   int    `json:"file_size,omitempty"`
}

type Dice struct {
    Emoji string `json:"emoji"` // Эмодзи кости, например 🎲
    Value int    `json:"value"` // Значение броска (1-6 или другое в зависимости от типа кости)
}

type VideoNote struct {
    Duration  int        `json:"duration"`  // длительность в секундах
    Length    int        `json:"length"`    // диаметр видео в пикселях (квадрат)
    Thumbnail *PhotoSize `json:"thumbnail,omitempty"` // миниатюра (thumbnail)
    Thumb     *PhotoSize `json:"thumb,omitempty"`     // синоним thumbnail (Telegram иногда использует оба)
    FileID    string     `json:"file_id"`
    FileUniqueID string  `json:"file_unique_id"`
    FileSize  int        `json:"file_size,omitempty"`
}

type Message struct {
    MessageID    int64       `json:"message_id"`
    From         User        `json:"from"`
    Chat         struct {
        ID        int64  `json:"id"`
        FirstName string `json:"first_name"`
        Username  string `json:"username"`
        Type      string `json:"type"`
    } `json:"chat"`
    Date         int64       `json:"date"`
    Text         string      `json:"text"`
    Photo        []PhotoSize `json:"photo,omitempty"`
    Video        *Video      `json:"video,omitempty"`
    Document     *Document   `json:"document,omitempty"`
    Caption      string      `json:"caption"`
    Location     *Location   `json:"location,omitempty"`
    Venue        *Venue      `json:"venue,omitempty"`
    Voice        *Voice      `json:"voice,omitempty"`
    Sticker      *Sticker    `json:"sticker,omitempty"`
    Dice         *Dice       `json:"dice,omitempty"`
    VideoNote    *VideoNote  `json:"video_note,omitempty"`
    PinnedMessage *Message   `json:"pinned_message,omitempty"` // добавлено поле pinned_message
}

type MyChatMember struct {
    Chat struct {
        ID    int64  `json:"id"`
        Title string `json:"title"`
        Type  string `json:"type"`
    } `json:"chat"`
    From struct {
        ID           int64  `json:"id"`
        IsBot        bool   `json:"is_bot"`
        FirstName    string `json:"first_name"`
        Username     string `json:"username"`
        LanguageCode string `json:"language_code"`
        IsPremium    bool   `json:"is_premium"`
    } `json:"from"`
    Date           int64 `json:"date"`
    OldChatMember  struct {
        User struct {
            ID        int64  `json:"id"`
            IsBot     bool   `json:"is_bot"`
            FirstName string `json:"first_name"`
            Username  string `json:"username"`
        } `json:"user"`
        Status string `json:"status"`
    } `json:"old_chat_member"`
    NewChatMember  struct {
        User struct {
            ID        int64  `json:"id"`
            IsBot     bool   `json:"is_bot"`
            FirstName string `json:"first_name"`
            Username  string `json:"username"`
        } `json:"user"`
        Status               string `json:"status"`
        CanBeEdited          bool   `json:"can_be_edited"`
        CanManageChat        bool   `json:"can_manage_chat"`
        CanChangeInfo        bool   `json:"can_change_info"`
        CanPostMessages      bool   `json:"can_post_messages"`
        CanEditMessages      bool   `json:"can_edit_messages"`
        CanDeleteMessages    bool   `json:"can_delete_messages"`
        CanInviteUsers       bool   `json:"can_invite_users"`
        CanRestrictMembers   bool   `json:"can_restrict_members"`
        CanPromoteMembers    bool   `json:"can_promote_members"`
        CanManageVideoChats  bool   `json:"can_manage_video_chats"`
        CanPostStories       bool   `json:"can_post_stories"`
        CanEditStories       bool   `json:"can_edit_stories"`
        CanDeleteStories     bool   `json:"can_delete_stories"`
        IsAnonymous          bool   `json:"is_anonymous"`
        CanManageVoiceChats  bool   `json:"can_manage_voice_chats"`
    } `json:"new_chat_member"`
}

func GetUpdates(offset int64) (*GetUpdatesResponse, error) {
    _ = godotenv.Load()
    token := os.Getenv("TOKEN")
    if token == "" {
        return nil, fmt.Errorf("токен не найден в переменных окружения")
    }
    // Формируем URL с параметром offset
    url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d", token, offset)
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("ошибка запроса: %w", err)
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("ошибка чтения ответа: %w", err)
    }
    var updates GetUpdatesResponse
    err = json.Unmarshal(body, &updates)
    if err != nil {
        return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
    }
    return &updates, nil
}

func SendDocument(chatID int64, filename, caption string) (map[string]interface{}, error) {
    token, err := getToken()
    if err != nil {
        return nil, err
    }
    url := fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", token)

    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    _ = writer.WriteField("chat_id", fmt.Sprintf("%d", chatID))
    _ = writer.WriteField("caption", caption)
    part, err := writer.CreateFormFile("document", filepath.Base(filename))
    if err != nil {
        return nil, err
    }
    _, err = io.Copy(part, file)
    if err != nil {
        return nil, err
    }
    writer.Close()
    req, err := http.NewRequest("POST", url, body)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    var result map[string]interface{}
    err = json.NewDecoder(resp.Body).Decode(&result)
    return result, err
}

type GetFileResponse struct {
    Ok     bool `json:"ok"`
    Result struct {
        FileID   string `json:"file_id"`
        FilePath string `json:"file_path"`
    } `json:"result"`
}

func GetFile(fileID string) (*GetFileResponse, error) {
    payload := map[string]string{"file_id": fileID}
    var resp GetFileResponse
    err := postToTelegram("getFile", payload, &resp)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

func DownloadFile(filePath, saveAs string) error {
    token, err := getToken()
    if err != nil {
        return err
    }
    url := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", token, filePath)
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    out, err := os.Create(saveAs)
    if err != nil {
        return err
    }
    defer out.Close()
    _, err = io.Copy(out, resp.Body)
    return err
}

func MassSend(chatIDs []int64, text string, blackList map[int64]bool) (map[string]interface{}, error) {
    count := 0
    var lastResp map[string]interface{}

    for _, chatID := range chatIDs {
        if blackList != nil && blackList[chatID] {
            continue
        }

        resp, err := SendMessage(chatID, text, nil)
        if err != nil {
            // Можно логировать ошибку и продолжить, или вернуть ошибку
            // Здесь продолжим отправку другим чатам
            continue
        }

        lastResp = resp

        if ok, exists := resp["ok"].(bool); exists && ok {
            count++
        }

        // При необходимости можно добавить задержку
        // time.Sleep(500 * time.Millisecond)
    }

    if lastResp == nil {
        lastResp = map[string]interface{}{"ok": false, "description": "no messages sent"}
    }

    lastResp["sent_count"] = count
    return lastResp, nil
}

type ForwardMessagePayload struct {
    ChatID     int64 `json:"chat_id"`      // Куда переслать
    FromChatID int64 `json:"from_chat_id"` // Откуда переслать
    MessageID  int64 `json:"message_id"`   // Какое сообщение переслать
}

func ForwardMessage(chatID, fromChatID, messageID int64) (map[string]interface{}, error) {
    payload := ForwardMessagePayload{
        ChatID:     chatID,
        FromChatID: fromChatID,
        MessageID:  messageID,
    }

    var result map[string]interface{}
    err := postToTelegram("forwardMessage", payload, &result)
    return result, err
}

type ReplyMessagePayload struct {
    ChatID           int64       `json:"chat_id"`
    Text             string      `json:"text"`
    ParseMode        string      `json:"parse_mode,omitempty"`
    ReplyMarkup      interface{} `json:"reply_markup,omitempty"`
    ReplyToMessageID int64       `json:"reply_to_message_id,omitempty"`
}

func ReplyMessage(chatID int64, text string, replyToMessageID int64, replyMarkup interface{}) (map[string]interface{}, error) {
    payload := ReplyMessagePayload{
        ChatID:           chatID,
        Text:             text,
        ParseMode:        "html",
        ReplyToMessageID: replyToMessageID,
        ReplyMarkup:      replyMarkup,
    }

    var result map[string]interface{}
    err := postToTelegram("sendMessage", payload, &result)
    return result, err
}
