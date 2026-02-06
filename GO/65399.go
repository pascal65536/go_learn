package main

import "fmt"


type Animal interface {
	MakeSound() string
}

type Dog struct {
}


func (dog Dog) MakeSound() {
	fmt.Printf("Гав!\n")
}

type Cat struct {
}

func (cat Cat) MakeSound() {
	fmt.Printf("Мяу!\n")
}


func main() {
	cat1 := Cat{}
	cat1.MakeSound()
	dog1 := Dog{}
	dog1.MakeSound()
}

// func main() {
// 	dog := Dog{}
// 	cat := Cat{}

// 	fmt.Printf(Animal(dog).MakeSound(), "\n")
// 	fmt.Printf(Animal(cat).MakeSound(), "\n")
// }
