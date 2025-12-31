package trie

type BinTrie struct {
	maxBits int
	root    *Node
}

func NewBinTrie(maxBits int) *BinTrie {
	return &BinTrie{maxBits: maxBits}
}

// Returns the last [maxBits] bits of [num].
func (bt *BinTrie) mod(num int) int {
	return num % (1 << bt.maxBits)
}

func (bt *BinTrie) String() string {
	if bt == nil {
		return "nil"
	}
	if bt.root == nil {
		return "empty"
	}
	return bt.root.String(bt.maxBits)
}
