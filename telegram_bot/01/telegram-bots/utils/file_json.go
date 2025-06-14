package utils

import (
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


func SaveJSON(folderName string, fileName string, data interface{}) error {
    // Создаём папку, если она не существует
    err := os.MkdirAll(folderName, os.ModePerm)
    if err != nil {
        return err
    }

    // Полный путь к файлу
    filePath := folderName + "/" + fileName

    // Преобразуем данные в JSON
    jsonData, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return err
    }

    // Сохраняем JSON в файл
    err = os.WriteFile(filePath, jsonData, 0644)
    return err
}
