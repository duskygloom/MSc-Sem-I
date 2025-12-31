package main

import (
	"bufio"
	"fmt"
	"hashtable/chaining"
	"hashtable/doublehashing"
	"hashtable/hashing"
	"hashtable/linprobing"
	"hashtable/quadprobing"
	"hashtable/table"
	"os"
	"strconv"
	"strings"
)

func chooseHashFunction() int {
	fmt.Println("1. Division method")
	fmt.Println("2. Multiplication method")
	fmt.Println("3. Folding method")
	fmt.Println("4. Mid-square method")
	fmt.Print("Choose hash function: ")
	var option int
	fmt.Scanf("%d", &option)
	return option
}

func chooseCollisionResolutionTechnique() int {
	fmt.Println("1. Linear probing")
	fmt.Println("2. Quadratic probing")
	fmt.Println("3. Double hashing")
	fmt.Println("4. Chaining")
	fmt.Print("Choose collision resolution technique: ")
	var option int
	fmt.Scanf("%d", &option)
	return option
}

func chooseOperation() int {
	fmt.Println("1. Insert")
	fmt.Println("2. Delete")
	fmt.Println("3. Search")
	fmt.Println("4. Rehash")
	fmt.Println("5. Print")
	fmt.Println("0. Quit")
	fmt.Print("Choose operation: ")
	var option int
	fmt.Scanf("%d", &option)
	return option
}

func main() {
	var hf hashing.HashFunction
	hfChoice := chooseHashFunction()
	switch hfChoice {
	case 1:
		hf = hashing.DivisionMethod
	case 2:
		hf = hashing.MultiplicationMethod
	case 3:
		hf = hashing.FoldingMethod
	case 4:
		hf = hashing.MidSquareMethod
	default:
		fmt.Println("Invalid choice of hash function.")
		return
	}
	fmt.Println()

	fmt.Print("Size of hash table: ")
	var size int
	fmt.Scanf("%d", &size)
	if size <= 0 {
		fmt.Println("Size should be greater than zero.")
		return
	}
	fmt.Println()

	var ht table.BaseTable
	htChoice := chooseCollisionResolutionTechnique()
	switch htChoice {
	case 1:
		ht = linprobing.NewTable(size, hf)
	case 2:
		ht = quadprobing.NewTable(size, hf)
	case 3:
		fmt.Println()
		var hf2 hashing.HashFunction
		hf2Choice := chooseHashFunction()
		switch hf2Choice {
		case 1:
			hf2 = hashing.DivisionMethod
		case 2:
			hf2 = hashing.MultiplicationMethod
		case 3:
			hf2 = hashing.FoldingMethod
		case 4:
			hf2 = hashing.MidSquareMethod
		default:
			fmt.Println("Invalid choice of hash function.")
			return
		}
		ht = doublehashing.NewTable(size, hf, hf2)
	case 4:
		ht = chaining.NewTable(size, hf)
	default:
		fmt.Println("Invalid choice of collision resolution technique.")
		return
	}

	proceed := true
	for proceed {
		fmt.Println()
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
				ok := ht.Insert(int(num))
				if !ok {
					fmt.Printf("Failed to insert %d.\n", num)
				}
			}
			fmt.Println(ht.NilString())
		case 2:
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
				ht.Delete(int(num))
			}
			fmt.Println(ht.NilString())
		case 3:
			fmt.Print("Enter elements: ")
			scanner := bufio.NewScanner(os.Stdin)
			if !scanner.Scan() {
				fmt.Println("Failed to scan input.")
				return
			}
			fmt.Print("Found: ")
			prefix := ""
			for s := range strings.SplitSeq(scanner.Text(), " ") {
				num, err := strconv.ParseInt(s, 10, 32)
				if err != nil {
					fmt.Printf("%s is not an integer.\n", s)
					return
				}
				if ht.Contains(int(num)) {
					fmt.Print(prefix + strconv.Itoa(int(num)))
					prefix = ", "
				}
			}
			fmt.Println()
		case 4:
			ht.Rehash()
			fmt.Println(ht.NilString())
		case 5:
			fmt.Println(ht.String())
		}
	}
}
