package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func generateNos(n int, wg *sync.WaitGroup, noChan chan<- int) {
	defer wg.Done()
	defer close(noChan)
	for i := 0; i < n; i++ {
		no := rand.Int()
		noChan <- no
		fmt.Println("Iteration No ", i, " Generated and send ", no)
	}

}
func printNos(idx int, wg *sync.WaitGroup, noChan <-chan int) {
	defer wg.Done()
	for no := range noChan {
		fmt.Println(idx, " Read ", no, "from channel")
	}

}

func main() {
	var wg sync.WaitGroup
	noChan := make(chan int)
	wg.Add(3)
	go generateNos(10, &wg, noChan)
	go printNos(1, &wg, noChan)
	go printNos(2, &wg, noChan)
	wg.Wait()
}
