package main

import "fmt"

type MyString string

func (s MyString) IsUpperCase() bool {
    for _, c := range s {
        if c >= 'a' && c <= 'z' {
            return false
        }
    }
    return true
}

func main() {
    i := MyString("A sSD").IsUpperCase()
    fmt.Println(i)
}
