package main

import (
	"binarytrie/trie"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getOption(compressed bool) int {
	fmt.Println("1. Print")
	if !compressed {
		fmt.Println("2. Insert")
		fmt.Println("3. Delete")
		fmt.Println("4. Lookup")
		fmt.Println("5. Compress")
	} else {
		fmt.Println("6. Uncompress")
	}
	fmt.Println("0. Exit")
	fmt.Print("Choose an operation: ")
	var option int
	fmt.Scanf("%d", &option)
	if compressed {
		switch option {
		case 0, 6:
			return option
		default:
			return -1
		}
	}
	if option >= 0 && option <= 5 {
		return option
	}
	return -1
}

func main() {
	fmt.Print("Number of bits: ")
	var maxBits int
	fmt.Scanf("%d", &maxBits)
	fmt.Printf("%d-Bit Binary Trie\n", maxBits)
	bt := trie.NewBinTrie(maxBits)
	scanner := bufio.NewScanner(os.Stdin)

	proceed := true
	compressed := false

	for proceed {
		fmt.Println()
		switch getOption(compressed) {
		case 0:
			proceed = false
		case 1:
			fmt.Println(bt.String())
		case 2:
			fmt.Print("Enter elements: ")
			if !scanner.Scan() {
				fmt.Println("Error: Failed to scan input")
				return
			}
			splitInput := strings.Split(scanner.Text(), " ")
			nums := make([]int, len(splitInput))
			for i, val := range splitInput {
				num, err := strconv.Atoi(val)
				if err == nil {
					nums[i] = num
				}
			}
			for _, val := range nums {
				bt.Insert(val)
			}
			fmt.Println(bt.String())
		case 3:
			fmt.Print("Enter elements: ")
			if !scanner.Scan() {
				fmt.Println("Error: Failed to scan input")
				return
			}
			splitInput := strings.Split(scanner.Text(), " ")
			nums := make([]int, len(splitInput))
			for i, val := range splitInput {
				num, err := strconv.Atoi(val)
				if err == nil {
					nums[i] = num
				}
			}
			for _, val := range nums {
				bt.Delete(val)
			}
			fmt.Println(bt.String())
		case 4:
			fmt.Print("Enter elements: ")
			if !scanner.Scan() {
				fmt.Println("Error: Failed to scan input")
				return
			}
			splitInput := strings.Split(scanner.Text(), " ")
			nums := make([]int, len(splitInput))
			for i, val := range splitInput {
				num, err := strconv.Atoi(val)
				if err == nil {
					nums[i] = num
				}
			}
			fmt.Print("Found: ")
			prefix := ""
			for _, val := range nums {
				n := bt.Lookup(val)
				if n != nil {
					fmt.Print(prefix + strconv.Itoa(val))
					prefix = ", "
				}
			}
			fmt.Println()
		case 5:
			bt.Compress()
			compressed = true
			fmt.Println(bt.String())
		case 6:
			bt.Uncompress()
			compressed = false
			fmt.Println(bt.String())
		}
	}
}
