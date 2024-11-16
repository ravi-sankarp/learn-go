package main

import "fmt"

func FindIndex[T comparable](arr []T, searchKey T) int {
	for i, v := range arr {
		if v == searchKey {
			return i
		}
	}
	return -1
}
func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(FindIndex(arr, 1))
}
