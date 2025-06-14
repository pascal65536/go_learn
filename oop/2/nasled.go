/*
Вот простой пример "наследования" через встраивание структур
*/

package main

import "fmt"

type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println("...")
}

type Dog struct {
    Animal // Встраивание
}

func (d Dog) Speak() {
    fmt.Println("Гав!")
}


type Cat struct {
    Animal // Встраивание
}

func (d Cat) Speak() {
    fmt.Println("Мяу!")
}


func main() {
    d := Cat{}
    d.Name = "Бобик"
    d.Speak()
	fmt.Printf("\n%v\n", d)
}
