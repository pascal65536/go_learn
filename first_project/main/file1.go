package main

import (
	"encoding/json"
	"os"
)

type Post struct {
	ID        int               `json:"id"`
	Date      string            `json:"date"`
	Text      string            `json:"text"`
	Views     int               `json:"views"`
	Forwards  int               `json:"forwards"`
	Reactions map[string]int    `json:"reactions"`
}

type Data map[string]map[string]Post

func main() {
	data, err := readJSON("posts.json")
	if err != nil {
		panic(err)
	}

	data["@baldebut"]["14929"] = Post{
		ID:       14929,
		Date:     "2025-06-01T10:00:00+00:00",
		Text:     "Новый пост для примера",
		Views:    1000,
		Forwards: 10,
		Reactions: map[string]int{
			"😊": 20,
			"👍": 15,
		},
	}

	err = saveJSON("data.json", data)
	if err != nil {
		panic(err)
	}
}

func readJSON(filename string) (Data, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data Data
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func saveJSON(filename string, data Data) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(data)
}
