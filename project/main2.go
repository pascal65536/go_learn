package main

import (
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

    // Создаем watcher
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        fmt.Println("Ошибка создания watcher:", err)
        os.Exit(1)
    }
    defer watcher.Close()

    err = watcher.Add(plainDir)
    if err != nil {
        fmt.Println("Ошибка добавления директории в watcher:", err)
        os.Exit(1)
    }

    fmt.Printf("Слежу за изменениями в папке: %s\n", plainDir)

    for {
        select {
        case event, ok := <-watcher.Events:
            if !ok {
                return
            }
            if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
                fmt.Printf("Файл изменён: %q\n", event.Name)

                relPath, _ := filepath.Rel(plainDir, event.Name)
                encryptedPath := filepath.Join(encryptedDir, relPath)

                // Создаем подкаталоги
                os.MkdirAll(filepath.Dir(encryptedPath), 0755)

                // Шифруем
                if err := encryptFile(key, event.Name, encryptedPath); err != nil {
                    fmt.Println("Ошибка шифрования:", err)
                } else {
                    fmt.Printf("Файл зашифрован: %s → %s\n", event.Name, encryptedPath)
                }
            }
        case err, ok := <-watcher.Errors:
            if !ok {
                return
            }
            fmt.Println("Ошибка watcher:", err)
        }
    }
}
