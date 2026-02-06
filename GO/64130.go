package main

import (
 "bufio"
 "fmt"
 "os"
 "strings"
)

func contains(slice string, str string, count int) int {
 words := strings.Fields(slice)
 i := 0
 for j, v := range words {
  if j >= count {
   break
  }
  if strings.ToLower(v) == strings.ToLower(str) {
   i++
  }
 }
 return i
}

func main() {
 var counter int
 fmt.Scanln(&counter)

 var input2 string
 fmt.Scanln(&input2)

 reader3 := bufio.NewReader(os.Stdin)
 input3, _ := reader3.ReadString('\n')
//  input3 = strings.TrimSpace(input3)

 fmt.Println(contains(input3, input2, counter))
}
