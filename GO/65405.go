package main

import (
	"fmt"
	"strings"
)

type Writer interface {
	Write(p []byte) int
}

type Reader interface {
	Read() []byte
}

type ReaderWriter interface {
	Writer
	Reader
}

type UpperReaderWriter struct {
	UpperString string
}

func (u *UpperReaderWriter) Write(p []byte) int {
	u.UpperString = strings.ToUpper(string(p))
	return len(p)
}

func (u *UpperReaderWriter) Read() []byte {
	return []byte(u.UpperString)
}

func main() {
	var rw ReaderWriter = &UpperReaderWriter{}

	data := []byte("Hello, Арсений!")
	n := rw.Write(data)
	fmt.Printf("Записано байт: %d\n", n)

	readData := rw.Read()
	fmt.Printf("Прочитано: %s\n", string(readData))
}
