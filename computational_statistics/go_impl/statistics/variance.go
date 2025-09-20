package statistics

func Variance(X []float64) float64 {
	squareSum := 0.0
	for _, val := range X {
		squareSum += val * val
	}
	mean := Mean(X)
	return squareSum/float64(len(X)) - (mean * mean)
}

// @returns new variance and new mean
func OnlineVariance(variance, mean float64, x float64, i int) (float64, float64) {
	if i == 0 {
		return 0.0, float64(x)
	}
	firstTerm := float64(i) * float64(i+1) * variance
	secondTerm := float64(i) * (x - mean) * (x - mean)
	numerator := firstTerm + secondTerm
	denominator := float64((i + 1) * (i + 1))
	variance_ := numerator / denominator
	mean_ := OnlineMean(mean, x, i)
	return variance_, mean_
}
