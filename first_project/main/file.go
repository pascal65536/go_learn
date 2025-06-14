package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Post описывает структуру одного поста
type Post struct {
	ID        int            `json:"id"`
	Date      string         `json:"date"`
	Text      string         `json:"text"`
	Views     int            `json:"views"`
	Forwards  int            `json:"forwards"`
	Reactions map[string]int `json:"reactions"`
}

// Data — это вложенная структура: пользователь -> id поста -> Post
type Data map[string]map[string]Post

// readJSON читает JSON из файла, если файла нет — создаёт пустой и возвращает пустые данные
func readJSON(filename string) (Data, error) {
	var data Data

	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		// Файл не существует — создаём пустой JSON
		data = make(Data)
		if err := saveJSON(filename, data); err != nil {
			return nil, err
		}
		return data, nil
	} else if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if len(bytes) == 0 {
		data = make(Data)
		return data, nil
	}

	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// saveJSON сохраняет данные в файл в формате JSON с отступами
func saveJSON(filename string, data Data) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}

// addPost добавляет новый пост в структуру данных
func addPost(data Data, user string, post Post) {
	if data[user] == nil {
		data[user] = make(map[string]Post)
	}
	data[user][fmt.Sprintf("%d", post.ID)] = post
}

// getPost возвращает пост по user и postID
func getPost(data Data, user string, postID string) (Post, bool) {
	posts, ok := data[user]
	if !ok {
		return Post{}, false
	}
	post, ok := posts[postID]
	return post, ok
}

// updatePost обновляет пост, если он существует
func updatePost(data Data, user string, postID string, newPost Post) bool {
	if _, ok := data[user][postID]; ok {
		data[user][postID] = newPost
		return true
	}
	return false
}

// deletePost удаляет пост, если он существует
func deletePost(data Data, user string, postID string) bool {
	if _, ok := data[user][postID]; ok {
		delete(data[user], postID)
		return true
	}
	return false
}

func main() {
	const filename = "posts.json"

	// Читаем данные из файла (или создаём пустой файл)
	data, err := readJSON(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	fmt.Println("Исходные данные:")
	fmt.Printf("%+v\n\n", data)

	// Пример: Добавление нового поста
	newPost := Post{
		ID:       14929,
		Date:     "2025-06-05T21:00:00+00:00",
		Text:     "Новый пост, добавленный программно.",
		Views:    0,
		Forwards: 0,
		Reactions: map[string]int{
			"👍": 0,
		},
	}
	addPost(data, "@baldebut", newPost)
	fmt.Println("После добавления нового поста:")
	fmt.Printf("%+v\n\n", data)

	// Пример: Получение поста
	postID := "14929"
	post, found := getPost(data, "@baldebut", postID)
	if found {
		fmt.Println("Полученный пост:", post)
	} else {
		fmt.Println("Пост не найден")
	}

	// Пример: Обновление поста
	post.Views = 10
	updated := updatePost(data, "@baldebut", postID, post)
	if updated {
		fmt.Println("Пост обновлён")
	}

	// Пример: Удаление поста
	deleted := deletePost(data, "@baldebut", "14927")
	if deleted {
		fmt.Println("Пост 14927 удалён")
	}

	// Сохраняем изменения обратно в файл
	if err := saveJSON(filename, data); err != nil {
		fmt.Println("Ошибка сохранения файла:", err)
		return
	}

	fmt.Println("Изменения сохранены в", filename)
}
