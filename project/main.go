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
    "syscall"

    "github.com/fsnotify/fsnotify"
    "os/signal"
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

// decryptFile расшифровывает файл
func decryptFile(key []byte, srcPath, dstPath string) error {
    ciphertext, err := os.ReadFile(srcPath)
    if err != nil {
        return err
    }

    block, _ := aes.NewCipher(key)
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return err
    }

    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return fmt.Errorf("данные слишком маленькие для nonce")
    }

    nonce, cipherdata := ciphertext[:nonceSize], ciphertext[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, cipherdata, nil)
    if err != nil {
        return err
    }

    return os.WriteFile(dstPath, plaintext, 0644)
}

// initPlainFolder создаёт папку plain и расшифровывает туда файлы из encrypted
func initPlainFolder(encryptedDir, plainDir string, key []byte) error {
    // Удаляем, если уже существует (например, после некорректного завершения)
    os.RemoveAll(plainDir)
    if err := os.Mkdir(plainDir, 0755); err != nil {
        return err
    }

    err := filepath.Walk(encryptedDir, func(path string, info os.FileInfo, err error) error {
        if info.IsDir() {
            relPath, err := filepath.Rel(encryptedDir, path)
            if err != nil {
                return err
            }
            os.MkdirAll(filepath.Join(plainDir, relPath), 0755)
            return nil
        }

        relPath, err := filepath.Rel(encryptedDir, path)
        if err != nil {
            fmt.Printf("Ошибка вычисления относительного пути для %s: %v\n", path, err)
            return nil
        }
        rel := filepath.Join(plainDir, relPath)

        if err := decryptFile(key, path, rel); err != nil {
            fmt.Printf("Ошибка расшифровки %s: %v\n", path, err)
        } else {
            fmt.Printf("Расшифрован: %s → %s\n", path, rel)
        }
        return nil
    })

    return err
}

// watchDirectory наблюдает за изменениями в директории
func watchDirectory(plainDir, encryptedDir string, key []byte) error {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return err
    }

    err = watcher.Add(plainDir)
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
                    relPath, err := filepath.Rel(plainDir, event.Name)
                    if err != nil {
                        fmt.Println("Ошибка получения относительного пути:", err)
                        continue
                    }
                    encryptedPath := filepath.Join(encryptedDir, relPath)

                    fmt.Printf("Файл изменён: %s\n", event.Name)
                    fmt.Printf("Шифруем и сохраняем как: %s\n", encryptedPath)

                    os.MkdirAll(filepath.Dir(encryptedPath), 0755)

                    if err := encryptFile(key, event.Name, encryptedPath); err != nil {
                        fmt.Println("Ошибка шифрования:", err)
                    } else {
                        fmt.Println("Файл успешно зашифрован.")
                    }
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

    // Создание plain и расшифровка
    fmt.Println("Инициализация папки plain...")
    if err := initPlainFolder(encryptedDir, plainDir, key); err != nil {
        fmt.Println("Ошибка инициализации plain:", err)
        os.Exit(1)
    }

    // Настройка выхода по Ctrl+C
    cleanup := func() {
        fmt.Println("\nОчистка: удаление папки plain...")
        if err := os.RemoveAll(plainDir); err != nil {
            fmt.Println("Ошибка удаления папки:", err)
        }
        os.Exit(0)
    }

    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-signalChan
        cleanup()
    }()

    // Запуск наблюдателя
    fmt.Println("Начинаем наблюдение за папкой:", plainDir)
    if err := watchDirectory(plainDir, encryptedDir, key); err != nil {
        fmt.Println("Ошибка запуска наблюдателя:", err)
        cleanup()
    }

    // Бесконечный цикл, чтобы программа не завершалась
    fmt.Println("Программа работает... Для выхода нажмите Ctrl+C")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {}
}