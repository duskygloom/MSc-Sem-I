package trie

func (root *Node) uncompress() {
	if root == nil {
		return
	}
	// if [root.level] is not following [root.parent.level]
	if root.parent != nil && root.level != root.parent.level+1 {
		parent := &Node{level: root.level - 1, parent: root.parent, data: root.data >> 1}
		path := (root.data >> (root.level - 1)) % 2
		// connect [root.parent] to [parent]
		if root.parent.left == root {
			root.parent.left = parent
		} else {
			root.parent.right = parent
		}
		// connect [parent] to [root]
		if path == 0 {
			parent.left = root
		} else {
			parent.right = root
		}
		root.parent = parent
		// uncompress parent
		parent.uncompress()
	}
	root.left.uncompress()
	root.right.uncompress()
}

func (bt *BinTrie) Uncompress() {
	if bt == nil || bt.root == nil {
		return
	}
	if bt.root.level != 0 {
		parent := &Node{}
		path := (bt.root.data >> (bt.root.level - 1)) % 2
		if path == 0 {
			parent.left = bt.root
		} else {
			parent.right = bt.root
		}
		bt.root.parent = parent
		bt.root = parent
	}
	bt.root.uncompress()
}
