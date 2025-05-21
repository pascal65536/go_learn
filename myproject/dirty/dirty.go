package dirty

import (
 "io"
 "os"
 "path/filepath"
)

// Функция для проверки существования директории
func checkDirExists(dir string) bool {
 info, err := os.Stat(dir)
 return !os.IsNotExist(err) && info.IsDir()
}

// Функция для проверки существования файла
func checkFileExists(file string) bool {
 _, err := os.Stat(file)
 return !os.IsNotExist(err)
}

// Функция для создания директории
func createDir(dir string) error {
 return os.MkdirAll(dir, os.ModePerm)
}

// Функция для копирования файла
func copyFile(src string, dst string) error {
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
func moveFile(src string, dst string) error {
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
