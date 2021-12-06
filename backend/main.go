package main

import (
	db "backend/db"
	f "fmt"
)

func main() {
	posts := db.DBRead()

	newPost := db.Post{
		ID:    "2",
		Title: "new title",
		Body:  "new body",
	}

	f.Printf("DB length before append: %d\n", len(posts))

	f.Printf("\nAppending to DB:\n%#v\n\n", newPost)
	db.DBAppendPost(&posts, newPost)

	f.Printf("DB length after append: %d\n", len(posts))
}
