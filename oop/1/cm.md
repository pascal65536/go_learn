Отлично! Давай создадим **три файла**:

- `main.go` — основная программа:
  - Принимает HTTP-запросы
  - Парсит данные
  - Определяет, какой бот вызван
  - Инициализирует нужный класс (bot1 или bot2)
  - Вызывает обработчик
  - Если метода нет в классе — вызывает дефолтный из main

- `bot1.go` — поведение первого бота (например, котик)

- `bot2.go` — поведение второго бота (например, собачка)

---

### 📁 Структура проекта:

```
telegram-bot/
├── main.go      ← точка входа, сервер, маршруты, дефолты
├── bot1.go      ← логика для Bot1
└── bot2.go      ← логика для Bot2
```

---

## ✅ Что будет делать каждая часть:

### 🔹 `main.go`:
- Запускает HTTP-сервер
- Обрабатывает `/callback` от Telegram
- В запросе определяет `bot_id` и тип обновления (`message`, `callback_query`)
- По `bot_id` выбирает конфиг бота
- Создаёт объект бота с нужным классом
- Вызывает у него `.HandleMessage()` или другой хендлер
- Если метод не реализован — вызывает дефолтный хендлер

---

### 🔹 `bot1.go`:
- Реализует структуру `Bot1`
- Может иметь методы:  
  ```go
  func (b Bot1) OnMessage(text string)
  func (b Bot1) OnStartCommand()
  ```
- Эти методы содержат специфическую логику для этого типа ботов

---

### 🔹 `bot2.go`:
- То же самое, но для другого типа бота
- Например, другой набор команд или реакций на сообщения

---

### 🧩 Как будет работать вызов:

```text
Telegram API → POST /callback → main.go:
   - Получил bot_id
   - Нашёл токен и тип бота (bot1 / bot2)
   - Вызвал handler:
       если bot1.OnMessage -> вызвать его
       иначе -> вызвать DefaultOnMessage из main.go
```

---

Теперь давай напишем код. Я разобью его по файлам.

---

## 🟦 1. main.go

```go
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
```

---

## 🟨 2. bot1.go

```go
package main

type Bot1 struct{}

func (b *Bot1) OnMessage(text string) {
	fmt.Printf("[CatBot] Получено сообщение: %s\n", text)
	// здесь можно отправить ответ через Telegram API
}
```

---

## 🟩 3. bot2.go

```go
package main

type Bot2 struct{}

func (b *Bot2) OnMessage(text string) {
	fmt.Printf("[DogBot] Получено сообщение: %s\n", text)
	// здесь можно выполнить команду или отправить ответ
}
```

---

## ✅ Как запустить:

1. Убедись, что все три файла находятся в одной директории
2. Выполни:
   ```bash
   go run main.go bot1.go bot2.go
   ```

3. Отправь тестовый запрос:
   ```bash
   curl -X POST http://localhost:8080/callback \
     -H "Content-Type: application/json" \
     -d '{"bot_id":"cat123","text":"Привет","type":"message"}'
   ```

   > В консоли должно появиться: `[CatBot] Получено сообщение: Привет`

   Попробуй заменить `"bot_id":"dog456"` — должен отработать второй бот.

---

## 🚀 Возможности расширения:

- Добавь новые типы обновлений: `command`, `callback_query`
- Реализуй методы `OnStartCommand()` в каждом боте
- Добавь HTTP клиент для реальной отправки сообщений в Telegram
- Сохрани ботов в базу или JSON-файл
- Подключи `.env` для токенов
- Реализуй middleware: логирование, авторизация, рейт-лимит

---

Если хочешь продолжить:
- Добавим команды
- Настроим реальную отправку сообщений
- Сделаем систему состояний
- Подключим .env
- Добавим мидлвари

Просто скажи, куда двигаться дальше 👇
