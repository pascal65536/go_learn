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
	if a > b {
		fmt.Println("Первое число больше второго")	
	} else if a < b {
		fmt.Println("Второе число больше первого")
	} else {
		fmt.Println("Числа равны")
	}
}
