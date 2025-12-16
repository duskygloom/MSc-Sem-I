package main

import (
	"binoheap/heap"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func chooseOperation() int {
	fmt.Println("1. Insert")
	fmt.Println("2. Minimum")
	fmt.Println("3. Delete minimum")
	fmt.Println("4. Delete")
	fmt.Println("5. Print")
	fmt.Println("0. Quit")
	fmt.Print("Choose operation: ")
	var option int
	fmt.Scanf("%d", &option)
	return option
}

func main() {
	h := heap.NewHeap()
	proceed := true
	for proceed {
		opChoice := chooseOperation()
		switch opChoice {
		case 0:
			proceed = false
		case 1:
			fmt.Print("Enter elements: ")
			scanner := bufio.NewScanner(os.Stdin)
			if !scanner.Scan() {
				fmt.Println("Failed to scan input.")
				return
			}
			for s := range strings.SplitSeq(scanner.Text(), " ") {
				num, err := strconv.ParseInt(s, 10, 32)
				if err != nil {
					fmt.Printf("%s is not an integer.\n", s)
					return
				}
				h.Insert(int(num))
			}
		case 2:
			if h.IsEmpty() {
				fmt.Println("Heap is empty.")
			} else {
				fmt.Print("Minimum element: ")
				fmt.Println(h.Minimum())
			}
		case 3:
			h.DeleteMin()
		case 4:
			fmt.Print("Enter elements: ")
			scanner := bufio.NewScanner(os.Stdin)
			if !scanner.Scan() {
				fmt.Println("Failed to scan input.")
				return
			}
			for s := range strings.SplitSeq(scanner.Text(), " ") {
				num, err := strconv.ParseInt(s, 10, 32)
				if err != nil {
					fmt.Printf("%s is not an integer.\n", s)
					return
				}
				h.Delete(int(num))
			}
		case 5:
			fmt.Println(h.String())
		default:
			fmt.Println("Invalid option.")
		}
		fmt.Println()
	}
}
