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
	a, err := strconv.ParseFloat(arr[0], 64)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}	
	b, err := strconv.ParseFloat(arr[1], 64)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}	
	c, err := strconv.ParseFloat(arr[2], 64)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}		
	if a == b && b == c {
		fmt.Println("Максимальное равенство")	
	} else {
		fmt.Println("Не равны")
	}
}
