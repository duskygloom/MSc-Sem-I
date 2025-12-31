package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"suffixtree/tree"
)

func getOption() int {
	fmt.Println("1. Create")
	fmt.Println("2. Lookup")
	fmt.Println("3. Print")
	fmt.Println("0. Quit")
	fmt.Print("Choose an operation: ")
	var option int
	fmt.Scanf("%d", &option)
	return option
}

func main() {
	var t *tree.Node
	scanner := bufio.NewScanner(os.Stdin)

	proceed := true
	for proceed {
		switch getOption() {
		case 0:
			proceed = false
		case 1:
			fmt.Print("Enter word: ")
			if !scanner.Scan() {
				fmt.Println("Error: Failed to scan input")
				return
			}
			t = tree.NewSufTree(scanner.Text())
			fmt.Println(t)
		case 2:
			fmt.Print("Enter suffix: ")
			if !scanner.Scan() {
				fmt.Println("Error: Failed to scan input")
				return
			}
			values := strings.Split(scanner.Text(), " ")
			fmt.Print("Found: ")
			prefix := ""
			for _, val := range values {
				n := t.Lookup(val)
				if n != nil {
					fmt.Print(prefix + val)
					prefix = ", "
				}
			}
			fmt.Println()
			fmt.Println(t)
		case 3:
			fmt.Println(t)
		}
		fmt.Println()
	}
}
