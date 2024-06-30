package main

import "fmt"

type I interface {
	Print()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) Print() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.Print()
}
