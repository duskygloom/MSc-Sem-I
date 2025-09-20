package main

import (
	"binotree/tree"
	"fmt"
)

func main() {
	n := tree.NewTree(2)
	m := tree.NewTree(2)
	fmt.Println(n.String())
	fmt.Println(tree.Merge(n, m))
}
