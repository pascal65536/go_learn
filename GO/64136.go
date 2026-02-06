package main

import "fmt"

func myFunc() int {
    a := 1
    return a
}

func main() {
    result := myFunc()
    fmt.Println(result)
}
