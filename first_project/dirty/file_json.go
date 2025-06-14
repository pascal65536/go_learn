package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "os"
)

func LoadJSON(folderName string, fileName string) map[string]interface{} {
    // Проверяем и создаём папку, если она не существует
    os.MkdirAll(folderName, os.ModePerm)

    // Полный путь к файлу
    filePath := folderName + "/" + fileName

    // Проверяем, существует ли файл
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        // Создаём пустой JSON файл, если он не существует
        emptyJSON := make(map[string]interface{})
        jsonData, _ := json.Marshal(emptyJSON)
        ioutil.WriteFile(filePath, jsonData, 0644)
    }

    // Читаем содержимое JSON файла
    fileContent, _ := ioutil.ReadFile(filePath)

    // Декодируем JSON в структуру
    var data map[string]interface{}
    json.Unmarshal(fileContent, &data)

    return data
}


func SaveJSON(folderName string, fileName string, data map[string]interface{}) {
    // Создаём папку, если она не существует
    os.MkdirAll(folderName, os.ModePerm)

    // Полный путь к файлу
    filePath := folderName + "/" + fileName

    // Преобразуем данные в JSON
    jsonData, _ := json.MarshalIndent(data, "", "  ")

    // Сохраняем JSON в файл
    ioutil.WriteFile(filePath, jsonData, 0644)
}


func main() {
    folderName := "data"
    fileName := "sample.json"

    // Выводим содержимое JSON файла
    fmt.Println("Содержимое JSON файла:", LoadJSON(folderName, fileName))
  
    data := map[string]interface{}{
        "name": "John Doe",
        "age":  30,
        "city": "New York",
    }

    // Вызываем функцию для сохранения данных
    SaveJSON(folderName, fileName, data)    

    // Выводим содержимое JSON файла
    fmt.Println("Содержимое JSON файла:", LoadJSON(folderName, fileName))
    
}
