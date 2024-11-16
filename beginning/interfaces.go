package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type I interface {
	ToJson() string
}

type Article struct {
	Id        int
	Title     string
	CreatedOn time.Time
}

func (article *Article) ToJson() string {
	js, _ := json.Marshal(article)
	return string(js)
}

func Interfaces() {
	var updater I = &Article{1, "Test", time.Now()}
    
    fmt.Println(updater.ToJson())
}
