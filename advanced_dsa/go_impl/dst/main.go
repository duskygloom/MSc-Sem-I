package main

import (
	"bufio"
	"dst/dst"
	"fmt"
	"os"
	"strconv"
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
	fmt.Print("Number of bits: ")
	var maxBits int
	fmt.Scanf("%d", &maxBits)
	fmt.Printf("%d-Bit Digital Search Tree\n", maxBits)

	t := dst.NewDST(maxBits)
	scanner := bufio.NewScanner(os.Stdin)

	proceed := true
	for proceed {
		fmt.Println()
		switch getOption() {
		case 0:
			proceed = false
		case 1:
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
				t.Insert(val)
			}
			fmt.Println(t)
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
				t.Delete(val)
			}
			fmt.Println(t)
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
			fmt.Print("Found: ")
			prefix := ""
			for _, val := range nums {
				n := t.Lookup(val)
				if n != nil {
					fmt.Print(prefix + strconv.Itoa(val))
					prefix = ", "
				}
			}
			fmt.Println()
			fmt.Println(t)
		case 4:
			fmt.Println(t)
		}
	}

}
