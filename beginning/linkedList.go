package main

import (
	"fmt"
)

type Node[T any] struct {
	next *Node[T]
	val  T
}

func main() {
	fmt.Println("Input the limit")
	length := 0
	fmt.Scanln(&length)
	var startNode Node[int]
	var currentNode *Node[int]
	for i := 0; i < length; i++ {
		fmt.Println("Enter", i, " st value")
		var value int
		fmt.Scanln(&value)
		node := Node[int]{nil, value}
		if i == 0 {
			startNode = node
			currentNode = &startNode

		} else {
			currentNode.next = &node
			currentNode = currentNode.next
		}
        fmt.Println(currentNode)
	}
	currentNode = &startNode
	for currentNode != nil {
		fmt.Println(currentNode)
		currentNode = currentNode.next
	}
}
