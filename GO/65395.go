package main

import "fmt"


type Person struct {
    name string
    age  int
    address string
}

func (person Person) PrettyPrint() {
    fmt.Println("Name:", person.name)
    fmt.Println("Age:", person.age)
    fmt.Println("Address:", person.address)
}


type Employee struct {
    name string
    position string
    salary float64 
    bonus float64
}

func (employee Employee) CalculateTotalSalary() {
	resultMessage := fmt.Sprintf(
        "Employee: %s\nPosition: %s\nTotal Salary: %.2f", 
        employee.name, employee.position, employee.salary + employee.bonus,
    )
	fmt.Println(resultMessage)
}


type Student struct {
    name string
    solvedProblems int
    scoreForOneTask float64
    passingScore float64
}

func (student Student) IsExcellentStudent() bool {
    return (student.scoreForOneTask * float64(student.solvedProblems)) >= student.passingScore
}



func main() {
	student := Student{name: "Gosha", solvedProblems: 30, scoreForOneTask: 10.0, passingScore: 290.0}
	fmt.Print(student.IsExcellentStudent(), "\n")
}