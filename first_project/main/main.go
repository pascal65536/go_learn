package main

import "os"
import "fmt"
import "io/ioutil"
import "encoding/json"


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

func saveJSON(filename string, data Data) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}

// readJSON читает JSON из файла, если файла нет — создаёт пустой и возвращает пустые данные
func readJSON(filename string) (Data, error) {
	var data Data

	file, err := os.Open(filename)
	if err != nil {
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

func main1() {
	fmt.Println("Hello, World!")

	data := Data{
		"abbey_road": {		
			"123": Post{
				ID: 123,
				Reactions: map[string]int{
					"001": 1,
				},
			},
		},
	}

	data["abbey_road"]["123"].Reactions["010"] = 2
	data["abbey_road"]["123"].Reactions["000"] = 0

	data["pascal65536"] = make(map[string]Post)
	post := Post{
		ID: 123,
		Reactions: map[string]int{
			"1": 11,
		},
	}
	data["pascal65536"]["Bob"] = post

	filename := "posts.json"
	saveJSON(filename, data)
}

func main() {
	fmt.Println("Hello, World!")

	filename := "posts.json"
	data, err := readJSON(filename)
	if err != nil {
		panic(err)
	}

	var key string
	key = "101"
	post := data["abbey_road"][key]
	fmt.Println(post)
	reactions := post.Reactions
	if reactions == nil {
		reactions = make(map[string]int)
	}
	fmt.Println(reactions)
	reactions["000"] = 1
	fmt.Println(reactions)
	post.Reactions = reactions
	data["abbey_road"][key] = post
	
	saveJSON(filename, data)
}