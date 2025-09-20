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

	// online correlation
	meanX, meanY := 0.0, 0.0
	varX, varY := 0.0, 0.0
	covXY := 0.0
	corXY := 0.0
	for index, value := range X {
		corXY, covXY, varX, varY, meanX, meanY = statistics.OnlineCorrelation(covXY, varX, varY, meanX, meanY, value, Y[index], index)
	}
	fmt.Printf("Online correlation: %f\n", corXY)

	// offline covariance
	fmt.Printf("Offline correlation: %f\n", statistics.Correlation(X, Y))
}
