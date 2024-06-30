package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loopTillLimit() {
	fmt.Println("Input the limit")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
	limit, _ := strconv.Atoi(input)
	fmt.Println("Limit is ", limit)
	for i := 0; i <= limit; i++ {
		fmt.Println(i)
	}
}
