package hashing

func baseFoldingMethod(key, limit, M int) int {
	hash := 0
	for key > 0 {
		hash += key % M
		key /= M
	}
	return hash % limit
}

func FoldingMethod(key, limit int) int {
	return baseFoldingMethod(key, limit, 100)
}
