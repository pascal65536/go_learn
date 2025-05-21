package main

import (
    "fmt"

    "myproject/dirty"
    "myproject/hello"
)

func main() {
    hello.Hello("World")
    fmt.Println("!!!!")
    
    // Измените вызов функции на Exported version
    files, err := dirty.ListFiles(".") // Изменено на ListFiles
    if err != nil {
        fmt.Println("Ошибка при получении списка файлов:", err)
        return
    }
    
    fmt.Println("Список файлов:")
    for _, file := range files {
        fmt.Println(file)
    }
}
