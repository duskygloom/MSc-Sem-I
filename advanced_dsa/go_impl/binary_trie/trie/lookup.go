package trie

func (root *Node) lookup(data, maxBits int) *Node {
	if root == nil {
		return nil
	}
	if root.data == data {
		return root
	}
	if getIthBit(data, root.level, maxBits) == 0 {
		return root.left.lookup(data, maxBits)
	}
	return root.right.lookup(data, maxBits)
}

func (bt *BinTrie) Lookup(data int) *Node {
	if bt == nil || bt.root == nil {
		return nil
	}
	return bt.root.lookup(bt.mod(data), bt.maxBits)
}
