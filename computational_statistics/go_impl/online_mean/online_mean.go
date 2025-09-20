package main

import (
	"fmt"
	"math/rand"
	"mcsc13/statistics"
)

func main() {
	numData := 40
	X := make([]float64, numData)
	for index := range numData {
		X[index] = rand.Float64() * 20
	}

	// online meanX
	meanX := 0.0
	for index, value := range X {
		meanX = statistics.OnlineMean(meanX, value, index)
	}
	fmt.Printf("Online mean: %f\n", meanX)

	// offline mean
	fmt.Printf("Offline mean: %f\n", statistics.Mean(X))
}
