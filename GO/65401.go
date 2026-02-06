package main

import "fmt"

type Person struct {
	Name       string
	MiddleName string
	Age        int
}

// Конструктор в Go обычно возвращает новый экземпляр структуры
func NewPerson(name, middleName string) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("имя не может быть пустым")
	}

	if middleName == "" {
		middleName = "Отчество отсутствует"
	}

	return &Person{Name: name, MiddleName: middleName, Age: 40}, nil
}


type Book struct {
	Title string
	Author string
	Year int
	Genre string
}


func NewBook(title, author string, year int, genre string) *Book {
	return &Book{
		Title: title,
		Author: author,
		Year: year,
		Genre: genre,
	}
}


type User struct {
	Name     string
	Age      int
	IsActive bool
}

func NewUser(name string, age int) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is empty for user")
	}
	if age == 0 {
		age = 18
	}
	return &User{
		Name:     name,
		Age:      age,
		IsActive: true,
	}, nil
}


func main() {
	// Используем конструктор
	person, err := NewPerson("Alice", "")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Имя: %s, Возраст: %d\n", person.Name, person.Age)
}
