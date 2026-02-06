package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func show(m map[int]string) {
	for i := 1; i <= 5; i++ {
		if value, ok := m[i]; ok {
			fmt.Printf("%d. %s\n", i, value)
		} else {
			fmt.Printf("%d. -\n", i)
		}
	}
}

func calc(m map[int]string) {
	empty := 0
	study := 0
	for i := 1; i <= 5; i++ {
		if _, ok := m[i]; ok {
			study++
		} else {
			empty++
		}
	}
	fmt.Printf("Осталось свободных мест: %d\nВсего человек в очереди: %d\n", empty, study)
}

func main() {
	m := make(map[int]string)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)

		if len(parts) == 1 {
			cmd := parts[0]
			switch cmd {
			case "очередь":
				show(m)
			case "количество":
				calc(m)
			case "конец":
				show(m)
				return
			default:
				fmt.Println("некорректный ввод")
			}
		} else if len(parts) == 2 {
			name := parts[0]
			num, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("некорректный ввод")
				continue
			}
			if num < 1 || num > 5 {
				fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", num)
				continue
			}

			if len(m) >= 5 {
				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", num)
				continue
			}

			if _, ok := m[num]; ok {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
			} else {
				m[num] = name
			}

		} else {
			fmt.Println("некорректный ввод")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
	}
}
