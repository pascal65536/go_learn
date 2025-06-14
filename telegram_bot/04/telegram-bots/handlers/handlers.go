package handlers

type BotHandler interface {
    HandleMessage(ctx BotContext)
}

type BotContext struct {
    Name     string
    Token    string
    Settings map[string]interface{}
    Update   map[string]interface{}
    Handler  BotHandler
}

func RegisterHandlers(handlersMap map[string]BotHandler) {
    handlersMap["CatBot"] = CatBot{}
    handlersMap["DogBot"] = DogBot{}
    handlersMap["AdminBot"] = AdminBot{}
}