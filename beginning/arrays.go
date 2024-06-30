package main

import "fmt"

func arrays() {
	var a = []int{1, 2, 4, 8, 16, 32, 64, 128}

	fmt.Printf("%T \n", &a[0])
	for i, v := range a {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
