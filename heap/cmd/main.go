package main

import (
	"fmt"

	"data-structures/heap"
)

func main() {
	tree := heap.NewHeapTree(100)

	tree.Push(70)
	tree.Push(50)
	tree.Push(120)
	tree.Push(20)
	tree.Push(60)

	tree.Print()
	fmt.Printf("%+v\n", tree)
}
