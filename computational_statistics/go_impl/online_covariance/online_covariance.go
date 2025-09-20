package main

import (
	"fmt"
	"math/rand"
	"mcsc13/statistics"
)

func main() {
	numData := 40
	X := make([]float64, numData)
	Y := make([]float64, numData)
	for index := range numData {
		X[index] = rand.Float64() * 20
		Y[index] = rand.Float64() * 20
	}

	// online covariance
	meanX, meanY := 0.0, 0.0
	covXY := 0.0
	for index, value := range X {
		covXY, meanX, meanY = statistics.OnlineCovariance(covXY, meanX, meanY, value, Y[index], index)
	}
	fmt.Printf("Online covariance: %f\n", covXY)

	// offline covariance
	fmt.Printf("Offline covariance: %f\n", statistics.Covariance(X, Y))
}
