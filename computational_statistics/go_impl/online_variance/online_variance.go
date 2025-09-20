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

	// online variance
	meanX := 0.0
	varX := 0.0
	for index, value := range X {
		varX, meanX = statistics.OnlineVariance(varX, meanX, value, index)
	}
	fmt.Printf("Online variance: %f\n", varX)

	// offline variance
	fmt.Printf("Offline variance: %f\n", statistics.Variance(X))
}
