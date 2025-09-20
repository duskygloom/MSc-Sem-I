package statistics

import (
	"fmt"
	"math"
)

func Covariance(X, Y []float64) float64 {
	if len(X) != len(Y) {
		fmt.Println("Size of both datasets should be equal.")
		return math.NaN()
	}
	meanX, meanY := Mean(X), Mean(Y)
	prodSum := 0.0
	for index, value := range X {
		prodSum += (value - meanX) * (Y[index] - meanY)
	}
	return prodSum / float64(len(X))
}

// @returns covariance, mean1 and mean2
func OnlineCovariance(cov, meanX, meanY float64, x, y float64, i int) (float64, float64, float64) {
	if i == 0 {
		return 0.0, x, y
	}
	term1 := float64(i) * float64(i+1) * cov
	term2 := float64(i) * (x - meanX) * (y - meanY)
	numerator := term1 + term2
	denominator := float64(i+1) * float64(i+1)
	cov_ := numerator / denominator
	meanX_ := OnlineMean(meanX, x, i)
	meanY_ := OnlineMean(meanY, y, i)
	return cov_, meanX_, meanY_
}
