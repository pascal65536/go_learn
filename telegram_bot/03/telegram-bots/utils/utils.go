package utils

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type Update struct {
    UpdateID      int64        `json:"update_id"`
    Message       *Message     `json:"message,omitempty"`
    MyChatMember  *MyChatMember `json:"my_chat_member,omitempty"`
    ChannelPost   *Message     `json:"channel_post,omitempty"`
    CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
}

type Message struct {
    MessageID int64  `json:"message_id"`
    Chat      Chat   `json:"chat"`
    Text      string `json:"text"`
}

type Chat struct {
    ID int64 `json:"id"`
}

type MyChatMember struct {
    NewChatMember ChatMember `json:"new_chat_member"`
}

type ChatMember struct {
    Status string `json:"status"`
}

type CallbackQuery struct {
    Data string `json:"data"`
}

type GetUpdatesResponse struct {
    Ok     bool     `json:"ok"`
    Result []Update `json:"result"`
}

func GetUpdates(token string, offset int64) (*GetUpdatesResponse, error) {
    url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d&timeout=30", token, offset)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var updatesResp GetUpdatesResponse
    err = json.Unmarshal(body, &updatesResp)
    if err != nil {
        return nil, err
    }

    if !updatesResp.Ok {
        return nil, fmt.Errorf("telegram API returned not ok")
    }

    return &updatesResp, nil
}

type SendMessagePayload struct {
    ChatID           int64       `json:"chat_id"`
    Text             string      `json:"text"`
    ParseMode        string      `json:"parse_mode,omitempty"`
    ReplyToMessageID int64       `json:"reply_to_message_id,omitempty"`
    ReplyMarkup      interface{} `json:"reply_markup,omitempty"`
}

func ReplyMessage(token string, chatID int64, text string, replyToMessageID int64) error {
    payload := SendMessagePayload{
        ChatID:           chatID,
        Text:             text,
        ParseMode:        "HTML",
        ReplyToMessageID: replyToMessageID,
    }

    url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
    bodyBytes, err := json.Marshal(payload)
    if err != nil {
        return err
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    var respData map[string]interface{}
    err = json.Unmarshal(respBody, &respData)
    if err != nil {
        return err
    }

    if ok, exists := respData["ok"].(bool); !exists || !ok {
        return fmt.Errorf("failed to send message: %v", respData)
    }

    return nil
}
