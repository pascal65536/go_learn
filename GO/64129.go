package main

import (
 "bufio"
 "fmt"
 "os"
 "strings"
)

func contains(slice []string, str string) bool {
 for _, v := range slice {
  if v == str {
   return true
  }
 }
 return false
}

func main() {
 valid := []string{"да", "нет", "чёрный", "белый"}
 reader := bufio.NewReader(os.Stdin)

 for {
  counter, err := reader.ReadString('\n')
  if err != nil {
   fmt.Println("Поражение")
   break
  }
  counter = strings.TrimSpace(counter)

  if contains(valid, counter) {
	fmt.Println("Поражение")
	break 
  } else {
	fmt.Println("Игра продолжается")	
  }
 }
}
