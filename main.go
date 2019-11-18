// from https://learning.oreilly.com/library/view/go-web-programming/9781617292569/kindle_split_018.html

package main

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"os"
)

type Post struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}
	return
}

func unmarshal(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("Error reading JSON data:", err)
		return
	}
	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return
	}
	return
}

func main() {
	_, err := decode("post.json")
	if err != nil {
		log.Println("Error:", err)
	}
	_, err = unmarshal("invalid.json")
	if err != nil {
		log.Println("Error:", err)
	}
}
