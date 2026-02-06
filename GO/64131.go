package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input1 := strings.TrimSpace(input)
	input2 := strings.ReplaceAll(input1, "о", "")
	input3 := strings.ReplaceAll(input2, "а", "")
	fmt.Println(input3)
}
