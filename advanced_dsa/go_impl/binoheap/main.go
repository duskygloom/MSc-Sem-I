package main

import (
	"binoheap/node"
	"fmt"
)

func main() {
	t := node.NewNode(4, nil)
	fmt.Println(t.String())
}
