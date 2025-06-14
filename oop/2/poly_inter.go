/* 
Вот простой пример полиморфизма через интерфейсы в Go 
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

type Robot struct{}

func (r Robot) Speak() {
    fmt.Println("Привет, я робот.")
}

func SayHello(s Speaker) {
    s.Speak()
}

func main() {
    dog := Dog{}
    cat := Cat{}
    robot := Robot{}

    SayHello(dog)
    SayHello(cat)
    SayHello(robot)
}