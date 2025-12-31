package main

import (
	"bufio"
	"fmt"
	"os"
	"prefixtree/tree"
	"strings"
)

func getOption() int {
	fmt.Println("1. Insert")
	fmt.Println("2. Delete")
	fmt.Println("3. Lookup")
	fmt.Println("4. Print")
	fmt.Println("0. Quit")
	fmt.Print("Choose an operation: ")
	var option int
	fmt.Scanf("%d", &option)
	return option
}

func main() {
	t := tree.NewPreTree()
	scanner := bufio.NewScanner(os.Stdin)

	proceed := true
	for proceed {
		switch getOption() {
		case 0:
			proceed = false
		case 1:
			fmt.Print("Enter elements: ")
			if !scanner.Scan() {
				fmt.Println("Error: Failed to scan input")
				return
			}
			values := strings.Split(scanner.Text(), " ")
			for _, val := range values {
				t.Insert(val)
			}
			fmt.Println(t)
		case 2:
			fmt.Print("Enter elements: ")
			if !scanner.Scan() {
				fmt.Println("Error: Failed to scan input")
				return
			}
			values := strings.Split(scanner.Text(), " ")
			for _, val := range values {
				t.Delete(val)
			}
			fmt.Println(t)
		case 3:
			fmt.Print("Enter elements: ")
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
		case 4:
			fmt.Println(t)
		}
		fmt.Println()
	}
}
