package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "runtime"
    "sort"
)

const outputFileName = "prj.md"

var extensions = map[string]bool{
    ".py":   true,
    ".go":   true,
    ".yaml": true,
}

func collectFiles(node *Node) []*Node {
    var files []*Node
    for _, child := range node.Files {
        if !child.IsDir {
            files = append(files, child)
        } else {
            files = append(files, collectFiles(child)...)
        }
    }
    return files
}


func main() {
    // 1. Указываем корневую папку для обхода
    root := "/home/pacal65536/git/go_learn/telegram_bot/02/"

    // 2. Получаем путь к текущему исходному файлу (main.go)
    _, thisFile, _, _ := runtime.Caller(0)
    srcDir := filepath.Dir(thisFile)

    // 3. Путь к выходному файлу рядом с main.go
    outputFilePath := filepath.Join(srcDir, outputFileName)

    // 4. Открываем выходной файл
    outputFile, err := os.Create(outputFilePath)
    if err != nil {
        fmt.Println("Ошибка создания файла:", err)
        return
    }
    defer outputFile.Close()

    // ✅ Выводим путь к файлу в терминал
    fmt.Printf("Файл успешно создан: %s\n", outputFilePath)

    // 5. Строим дерево структуры проекта
    tree := buildTree(root)
    writeTreeToFile(outputFile, tree)

    // 6. Собираем все файлы рекурсивно и пишем их содержимое
    files := collectFiles(tree)
    for _, fileNode := range files {
        writeFileContent(outputFile, fileNode)
    }
}

type Node struct {
    Name  string
    Path  string
    IsDir bool
    Files []*Node
}

func buildTree(path string) *Node {
    node := &Node{Name: filepath.Base(path), Path: path, IsDir: true, Files: []*Node{}}

    files, err := ioutil.ReadDir(path)
    if err != nil {
        fmt.Println("Ошибка чтения папки:", path, err)
        return node
    }

    var items []os.FileInfo
    for _, f := range files {
        if f.IsDir() || extensions[filepath.Ext(f.Name())] {
            items = append(items, f)
        }
    }

    sort.Slice(items, func(i, j int) bool {
        return items[i].Name() < items[j].Name()
    })

    for _, f := range items {
        childPath := filepath.Join(path, f.Name())
        if f.IsDir() {
            node.Files = append(node.Files, buildTree(childPath))
        } else {
            node.Files = append(node.Files, &Node{Name: f.Name(), Path: childPath, IsDir: false})
        }
    }

    return node
}

func writeTreeToFile(file *os.File, root *Node) {
    file.WriteString(fmt.Sprintf("%s/\n", root.Name))
    for i, child := range root.Files {
        writeNode(file, child, "", i == len(root.Files)-1, true)
    }
    file.WriteString("\n")
}

func writeNode(file *os.File, node *Node, prefix string, isLast bool, isRoot bool) {
    // Определяем иконку для подчинённого элемента
    icon := "├── "
    if isLast {
        icon = "└── "
    }

    // Формируем имя с описанием
    name := node.Name

    // Записываем строку узла
    if isRoot {
        file.WriteString(name + "\n")
    } else {
        file.WriteString(prefix + icon + name + "\n")
    }

    // Обновляем префикс для следующего уровня
    newPrefix := prefix
    if !isRoot {
        if isLast {
            newPrefix += "    "
        } else {
            newPrefix += "│   "
        }
    }

    // Рекурсивно обрабатываем дочерние узлы
    for i, child := range node.Files {
        writeNode(file, child, newPrefix, i == len(node.Files)-1, false)
    }
}

func writeFileContent(file *os.File, node *Node) {
    content, err := ioutil.ReadFile(node.Path)
    if err != nil {
        fmt.Println("Ошибка чтения файла:", node.Path, err)
        return
    }

    file.WriteString("```\n")
    file.WriteString("// " + node.Name + "\n\n")
	file.WriteString(string(content))
    file.WriteString("\n```\n\n")
}