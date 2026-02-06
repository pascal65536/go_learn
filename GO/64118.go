package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	arr := strings.Fields(strings.TrimSpace(input))
	
	a, err := strconv.Atoi(arr[0])
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}	
	b, err := strconv.Atoi(arr[1])
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}	
	c, err := strconv.Atoi(arr[2])
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}			

	if (a + b/100) >= c {
		fmt.Println("Сегодня будет вкусный кофе!")
	} else {
		fmt.Println("Стоит подкопить")
	}
}
