package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	WordCountMap := make(map[string]int)
	WordsInString := strings.Fields(s)
    fmt.Println(WordsInString)
	for _, char := range WordsInString {
		current, ok := WordCountMap[string(char)]
		if ok {
			WordCountMap[string(char)] = current + 1
		} else {
            WordCountMap[string(char)] = 1
        }
	}
    fmt.Print(WordCountMap)
	return WordCountMap
}

func Run() {
	wc.Test(WordCount)
}
