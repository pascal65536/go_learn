package main

import 	"strings"
import "fmt"

func CountLengthAndBytes(first, second string) string {
    words := []string{first, second}
    joined := strings.Join(words, "")
	i := 0
	for range []rune(joined) {
		i++
	}
	length := len(joined)
	resultMessage := fmt.Sprintf("Объединённая строка: %s. Количество байт: %d. Количество символов: %d.", joined, length, i)
	return resultMessage
}





func main() {
	fmt.Println(CountLengthAndBytes("Привет,", " мир!"))
	fmt.Println(CountLengthAndBytes("I love ", " Yandex!"))
	fmt.Println(CountLengthAndBytes("你好", "不好意思"))
}

