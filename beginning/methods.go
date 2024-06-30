package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Post struct {
	Id        int
	Title     string
	CreatedOn time.Time
}

func (post *Post) toJson() string {
	post.Title = "Post is updated"
	js, _ := json.Marshal(post)
	return string(js)
}

func TestMethod() {
	post := Post{1, "Test", time.Now()}
	fmt.Println(post.toJson())
	fmt.Println(post.Title)
}
