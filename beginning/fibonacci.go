package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	previousValue := 0
	current := 0
	return func() int {
		result := previousValue + current
		previousValue = current
		current = result
        if current == 0{
            current = 1
        }
		return result
	}
}

func FibTest() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
