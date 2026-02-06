package main

import (
	"fmt"
	"strings"
	"unicode"
)

func NumbersToLetters(input string) string {
	m := map[rune]string{
		'0': "ноль",
		'1': "один",
		'2': "два",
		'3': "три",
		'4': "четыре",
		'5': "пять",
		'6': "шесть",
		'7': "семь",
		'8': "восемь",
		'9': "девять",
		'+': "плюс",
		'-': "минус",
		'*': "умножить на",
		'/': "разделить на",
		'(': "(",
		')': ")",
		' ': " ",
	}

	var total []string

	for _, r := range input {
		if val, ok := m[r]; ok {
			total = append(total, val)
		} else {
			if unicode.IsSpace(r) {
				total = append(total, " ")
			} else {
				total = append(total, string(r))
			}
		}
	}

	joined := strings.Join(total, "")
	joined = strings.Join(strings.Fields(joined), " ")

	return joined
}

func main() {
	fmt.Println(NumbersToLetters("(1 + 2) * 3 / 8"))
	fmt.Println(NumbersToLetters("2 + 2 * 2"))
	fmt.Println(NumbersToLetters("один - 2 / 2"))
	fmt.Println(NumbersToLetters("Хочу попасть на стажировку в Яндекс!"))
	fmt.Println(NumbersToLetters("1 2 3 4 5 6 7 8 9 + + - * /"))
}
