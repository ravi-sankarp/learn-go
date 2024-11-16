package main

import (
	"fmt"
	"time"
)

func say(s string, isThread bool) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, " thread ", isThread)
	}
}

func main() {
	go say("world", true)
	say("hello", false)
}
