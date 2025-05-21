package main

import (
    "bufio"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
    "os"
    "path/filepath"

    "github.com/fsnotify/fsnotify"
)

// loadKey загружает ключ из файла
func loadKey(keyPath string) ([]byte, error) {
    key, err := os.ReadFile(keyPath)
    if err != nil {
        return nil, err
    }
    if len(key) != 32 { // AES-256 требует 32 байта
        return nil, fmt.Errorf("ключ должен быть длиной 32 байта")
    }
    return key, nil
}

// encryptFile шифрует файл
func encryptFile(key []byte, srcPath, dstPath string) error {
    plaintext, err := os.ReadFile(srcPath)
    if err != nil {
        return err
    }

    block, _ := aes.NewCipher(key)
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return err
    }

    ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

    return os.WriteFile(dstPath, ciphertext, 0644)
}

// watchDirectory наблюдает за изменениями в директории
func watchDirectory(path string, callback func(string)) error {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return err
    }

    err = watcher.Add(path)
    if err != nil {
        return err
    }

    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
                    callback(event.Name)
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                fmt.Println("Ошибка наблюдателя:", err)
            }
        }
    }()

    return nil
}

func main() {
    plainDir := "plain"
    encryptedDir := "encrypted"
    keyPath := "key.bin"

    // Загрузка ключа
    key, err := loadKey(keyPath)
    if err != nil {
        fmt.Println("Ошибка загрузки ключа:", err)
        os.Exit(1)
    }

    fmt.Println("Начинаем наблюдение за папкой:", plainDir)

    // Запуск наблюдателя
    err = watchDirectory(plainDir, func(filePath string) {
        relPath, _ := filepath.Rel(plainDir, filePath)
        encryptedPath := filepath.Join(encryptedDir, relPath)

        fmt.Printf("Файл изменён: %s\n", filePath)
        fmt.Printf("Шифруем и сохраняем как: %s\n", encryptedPath)

        // Создаем подкаталоги при необходимости
        os.MkdirAll(filepath.Dir(encryptedPath), 0755)

        // Шифруем
        if err := encryptFile(key, filePath, encryptedPath); err != nil {
            fmt.Println("Ошибка шифрования:", err)
        } else {
            fmt.Println("Файл успешно зашифрован.")
        }
    })
    if err != nil {
        fmt.Println("Ошибка запуска наблюдателя:", err)
        os.Exit(1)
    }

    // Бесконечный цикл, чтобы программа не завершалась
    fmt.Println("Программа работает... Для выхода нажмите Ctrl+C")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {}
}
