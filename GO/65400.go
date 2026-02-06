/*

Создайте интерфейс Logger с методом Log(message string), который будет записывать сообщение в журнал. 
Создайте пользовательский тип LogLevel типа string, сделайте константные значения типа LogLevel со значениями Error и Info. 
Создайте структуру Log с полем Level типа LogLevel. Реализуйте метод Log с параметром типа string (текст ошибки), 
который в зависимости от текущего LogLevel выводит сообщение ERROR: *текст ошибки* или INFO: *текст ошибки*.

Интерфейс Logger с методом Log(message string), пользовательский тип LogLevel типа string с значениями Error и Info, 
структура Log с полем Level типа LogLevel, реализующая этот интерфейс

*/



package main

import "fmt"

type Logger interface {
	Log(message string)
}

type LogLevel string

const (
	Error LogLevel = "ERROR"
	Info  LogLevel = "INFO"
)

type Log struct {
	Level LogLevel
}

func (l Log) Log(message string) {
	switch l.Level {
	case Error:
		fmt.Printf("ERROR: %s\n", message)
	case Info:
		fmt.Printf("INFO: %s\n", message)
	default:
		fmt.Printf("LOG: %s\n", message)
	}
}

func main() {
	errorLogger := Log{Level: Error}
	infoLogger := Log{Level: Info}

	errorLogger.Log("Something went wrong")
	infoLogger.Log("Application started")
}
