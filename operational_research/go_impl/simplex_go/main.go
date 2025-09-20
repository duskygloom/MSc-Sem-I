package main

import (
	"bufio"
	"fmt"
	"os"
	"simplex/slackform"
	"simplex/stdform"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(input)

	c := []float64{}
	scanner.Scan()
	for i := range strings.SplitSeq(scanner.Text(), " ") {
		value, err := strconv.ParseFloat(i, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		c = append(c, value)
	}
	C := mat.NewDense(1, len(c), c)

	b := []float64{}
	scanner.Scan()
	for i := range strings.SplitSeq(scanner.Text(), " ") {
		value, err := strconv.ParseFloat(i, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		b = append(b, value)
	}
	B := mat.NewDense(1, len(b), b)

	a := []float64{}
	for range len(b) {
		scanner.Scan()
		for i := range strings.SplitSeq(scanner.Text(), " ") {
			value, err := strconv.ParseFloat(i, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			a = append(a, value)
		}
	}
	A := mat.NewDense(len(b), len(c), a)

	std, err := stdform.NewStd(A, B, C)
	if err != nil {
		fmt.Println(err)
		return
	}

	// stdFmt, err := std.Format()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(stdFmt)

	slack, err := slackform.FromStd(std)
	if err != nil {
		fmt.Println(err)
		return
	}

	slackFmt, err := slack.Format()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(slackFmt)

	for slack.Incoming() {

	}
}
