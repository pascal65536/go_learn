/*
Вот пример полиморфизма через композицию в Go
*/

package main

import "fmt"

type Speaker interface {
    Speak()
}

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Гав!")
}

type Cat struct{}

func (c Cat) Speak() {
    fmt.Println("Мяу!")
}

type Person struct {
    Name   string
    Pet    Speaker
}

func (p Person) Talk() {
    fmt.Printf("%s говорит: ", p.Name)
    p.Pet.Speak()
}

func main() {
    dog := Dog{}
    cat := Cat{}

    person1 := Person{Name: "Анна", Pet: cat}
    person2 := Person{Name: "Иван", Pet: dog}

    person1.Talk()
    person2.Talk()
}