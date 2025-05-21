package dirty

import (
    "io"
    "os"
    "path/filepath"
    "crypto/sha256"
    "fmt"
)

func СalculateSHA256(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    hash := sha256.New()
    if _, err := io.Copy(hash, file); err != nil {
        return "", err
    }
    
    checksum := hash.Sum(nil)
    return fmt.Sprintf("%x", checksum), nil
}


// Функция для проверки существования директории
func CheckDirExists(dir string) bool {
    info, err := os.Stat(dir)
    return !os.IsNotExist(err) && info.IsDir()
}

// Функция для проверки существования файла
func CheckFileExists(file string) bool {
    _, err := os.Stat(file)
    return !os.IsNotExist(err)
}

// Функция для создания директории
func CreateDir(dir string) error {
    return os.MkdirAll(dir, os.ModePerm)
}

// Функция для копирования файла
func CopyFile(src string, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()

    destinationFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destinationFile.Close()

    _, err = io.Copy(destinationFile, sourceFile)
    return err
}

// Функция для перемещения файла
func MoveFile(src string, dst string) error {
    return os.Rename(src, dst)
}

func ListFiles(dir string) ([]string, error) {
    var files []string
    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            files = append(files, path)
        }
        return nil
    })
    return files, err
}
