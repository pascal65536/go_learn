package main

import (
    "fmt"
    "first_project/dirty"
    "first_project/hello"
)


func printFileList(file_lst []string) {
    fmt.Println("Список файлов:")
    for _, file := range file_lst {
        hash, err := dirty.СalculateSHA256(file)
        if err != nil {
            fmt.Println("Ошибка при СalculateSHA256:", err)
            return
        }
        fmt.Println(hash, file)
    }    
}


func main() {
    hello.Hello("World")

    dirty.CreateDir("test")
    err := dirty.CopyFile("go.mod", "test/go.mod")
    if err != nil {
        fmt.Println("Ошибка при CopyFile:", err)
        return
    }

    file_lst, err := dirty.ListFiles(".")
    if err != nil {
        fmt.Println("Ошибка при ListFiles:", err)
        return
    }
    printFileList(file_lst)


}
