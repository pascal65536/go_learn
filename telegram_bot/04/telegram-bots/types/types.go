package types

type BotConfig struct {
    Token    string                 `json:"token"`
    Class    string                 `json:"class"`
    Settings map[string]interface{} `json:"settings"`
}