package main

import "fmt"
import "math"


type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	c := Circle{radius: 5}
	r := Rectangle{width: 4, height: 3}

	fmt.Printf("Площадь круга: %.2f\n", c.Area())
	fmt.Printf("Площадь прямоугольника: %.2f\n", r.Area())
}

/*
// Определяем интерфейс Human с методом SayHello
type Human interface {
    SayHello()
}

// Структура Person
type Person struct {
    Name string
    Age  int
}

// Реализация метода SayHello для Person
func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.Name, p.Age)
}

func main() {
    // Создаём переменную интерфейсного типа Human и присваиваем ей Person
    var h Human = Person{Name: "John", Age: 25}
    h.SayHello()

    // Можно также использовать напрямую Person
    student := Person{Name: "Gosha", Age: 30}
    student.SayHello()
}
*/