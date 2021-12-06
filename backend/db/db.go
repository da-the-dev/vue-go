package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func DBRead() []Post {
	buffer, err := ioutil.ReadFile("../db/posts.json")
	if err != nil {
		log.Fatalln("Bad DB file read, verify. Error\n", err)
	}

	// println("Buffer:", buffer)

	var posts []Post

	err = json.Unmarshal(buffer, &posts)
	if err != nil {
		log.Fatalln("Bad DB Unmarshal, verify.", err)
	}

	return posts
}

func DBAppendPost(posts *[]Post, post Post) {
	readPosts := DBRead()

	// Assing a new ID to a post
	lastID := readPosts[len(readPosts)-1].ID
	oldID, _ := strconv.Atoi(lastID)
	newID := strconv.Itoa(oldID + 1)
	post.ID = newID

	readPosts = append(readPosts, post)

	*posts = readPosts

	buffer, err := json.MarshalIndent(&posts, " ", "    ")
	if err != nil {
		log.Fatalln("Bad DB MarshalIndent, verify.")
	}

	err = ioutil.WriteFile("../db/posts.json", buffer, os.ModePerm)
	if err != nil {
		log.Fatalln("Bad DB write, verify.", err)
	}
}
