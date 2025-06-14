Главная функция получает команду по Callback API.
Роут, на который приходт Callback имеет вид https://mysite.com/bot/<name_bot>
Главная функция инициирует бота по его name_bot. 
Бот - это, name_bot, token_bot, class_bot. 
Класс - это набор функций, которые выполняет бот этого класса. Некоторые параметры функции, свойственные конкретному боту, хранятся в настройках бота. 
При инициализации главная функция собирает настройки бота из отдельного файла, где они хранятся в струкрутрированном виде. 



go mod init telegram-bots
go mod tidy
go get github.com/joho/godotenv
go run main.go handlers/. utils/.


curl -X POST http://localhost:8080/bot/bot_cat -H "Content-Type: application/json" -d '{"text":"Привет от кота!"}'
