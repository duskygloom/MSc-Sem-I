package trie

import "fmt"

func getIthBit(num, i, maxBits int) int {
	if i >= 0 && i <= maxBits-1 {
		return (num >> (maxBits - i - 1)) % 2
	} else {
		return 0
	}
}

func getBitString(num, maxBits int) string {
	return fmt.Sprintf("%0*b", maxBits, num%(1<<maxBits))
}
