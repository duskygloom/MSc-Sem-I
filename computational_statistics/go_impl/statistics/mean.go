package statistics

func Mean(X []float64) float64 {
	sum := 0.0
	for _, val := range X {
		sum += val
	}
	return sum / float64(len(X))
}

func OnlineMean(mean float64, x float64, i int) float64 {
	if i == 0 {
		return x
	}
	numerator := float64(i)*mean + x
	denominator := float64(i) + 1
	return numerator / denominator
}
