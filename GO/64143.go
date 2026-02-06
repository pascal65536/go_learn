package main

import (
	"fmt"
	"strings"
)

func CheckLetters(text string) string {
	countE := strings.Count(text, "е")
	if countE == 0 {
		return "Текст готов к публикации!"
	} else {
		return fmt.Sprintf("Количество возможных ошибок: %d, перепроверьте текст.", countE)
	}
}

func main() {
	fmt.Println(CheckLetters(""))
	fmt.Println(CheckLetters("Тут точно нет сложных букв, хотя ж выглядит подозрительно"))
	fmt.Println(CheckLetters("Ох уж эти е и ё"))
	fmt.Println(CheckLetters("Е ее ее тум турун тум турун тум"))
}