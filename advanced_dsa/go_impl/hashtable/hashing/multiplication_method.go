package hashing

import "math"

func baseMultiplicationMethod(key, limit int, A float64) int {
	prod := float64(key) * A
	trunc := prod - math.Trunc(prod)
	return int(math.Floor(float64(limit) * trunc))
}

func MultiplicationMethod(key, limit int) int {
	r := (math.Sqrt(5) - 1) / 2
	return baseMultiplicationMethod(key, limit, r)
}
