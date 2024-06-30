package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	Name string
	Body string
	Time time.Time
}

func structs() {
	test := Message{"Test", "This is a test message", time.Now()}
	res, _ := json.Marshal(test)
	fmt.Println(string(res[:]))
}
