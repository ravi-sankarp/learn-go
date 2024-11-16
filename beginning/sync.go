package main

import (
	"fmt"
	"sync"
	"time"
)

func LoopUntil(initNo int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	timeStart := time.Now()
	for i := 0; i < n; i++ {
	}
	execTime := time.Now().Sub(timeStart)
	fmt.Println(initNo, " Finished execution in ", execTime)
}

func main() {
	timeStart := time.Now()
	var wg sync.WaitGroup
	wg.Add(3)
	go LoopUntil(1, 4000000000, &wg)
	go LoopUntil(2, 4000000000, &wg)
	go LoopUntil(3, 4000000000, &wg)
	wg.Wait()
	execTime := time.Now().Sub(timeStart)
	fmt.Println(" Finished execution in ", execTime)
}
