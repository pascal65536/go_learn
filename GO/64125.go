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
	tmpr, _ := strconv.ParseFloat(strings.ReplaceAll(strings.TrimSpace(input), " ", ""), 64)
	if tmpr > 20 {
		fmt.Println("Стоит надеть майку и шорты")
	} else if tmpr > 10 && tmpr <= 20 {
		fmt.Println("Стоит надеть штаны и кофту")
	} else if tmpr > -5 && tmpr <= 9 {
		fmt.Println("Стоит надеть куртку")
	} else {
		fmt.Println("Стоит надеть зимнюю куртку")
	}	
}
