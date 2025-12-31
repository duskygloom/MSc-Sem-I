package trie

func (root *Node) insert(data, maxBits int) {
	if root == nil {
		return
	}
	if root.isElement {
		if root.data == data {
			return // if same data exists, then return
		}
		// make [root] a branch node and move [root.data] down
		// this is the process where branch nodes are created
		root.isElement = false
		if getIthBit(root.data, root.level, maxBits) == 0 {
			root.left = NewElementNode(root.data, root.level+1, root)
		} else {
			root.right = NewElementNode(root.data, root.level+1, root)
			root.right.level = root.level + 1
		}
		root.data = data >> (maxBits - root.level)
		root.insert(data, maxBits)
	} else if getIthBit(data, root.level, maxBits) == 0 && root.left == nil {
		root.left = NewElementNode(data, root.level+1, root)
		root.left.parent = root
	} else if getIthBit(data, root.level, maxBits) == 0 {
		root.left.insert(data, maxBits)
	} else if getIthBit(data, root.level, maxBits) == 1 && root.right == nil {
		root.right = NewElementNode(data, root.level+1, root)
		root.right.parent = root
	} else {
		root.right.insert(data, maxBits)
	}
}

func (bt *BinTrie) Insert(data int) {
	if bt.root == nil {
		bt.root = NewElementNode(bt.mod(data), 0, nil)
	} else {
		bt.root.insert(bt.mod(data), bt.maxBits)
	}
}
