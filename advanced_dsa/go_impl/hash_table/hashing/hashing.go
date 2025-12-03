package hashing

type HashFunction func(int, int) int

func DivisionMethod(key, limit int) int {
	return key % limit
}
