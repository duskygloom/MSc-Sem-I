package statistics

import "math"

func Correlation(X, Y []float64) float64 {
	cov := Covariance(X, Y)
	varX := Variance(X)
	varY := Variance(Y)
	return cov / math.Sqrt(varX*varY)
}

// @returns correlation, covariance, variance x, variance y, mean x, mean y
func OnlineCorrelation(cov, varX, varY, meanX, meanY float64, x, y float64, i int) (float64, float64, float64, float64, float64, float64) {
	if i == 0 {
		return 1.0, 0.0, 0.0, 0.0, x, y
	}
	cov_, meanX_, meanY_ := OnlineCovariance(cov, meanX, meanY, x, y, i)
	varX_, _ := OnlineVariance(varX, meanX, x, i)
	varY_, _ := OnlineVariance(varY, meanY, y, i)
	cor_ := cov_ / math.Sqrt(varX_*varY_)
	return cor_, cov_, varX_, varY_, meanX_, meanY_
}
