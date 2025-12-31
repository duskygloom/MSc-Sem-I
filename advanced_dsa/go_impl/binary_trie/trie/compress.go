package trie

func (root *Node) compress() {
	if root == nil || (root.left == nil && root.right == nil) {
		return
	}
	if root.parent != nil && root.left == nil {
		if root.parent.left == root {
			root.parent.left = root.right
			root.right.parent = root.parent
		} else {
			root.parent.right = root.right
			root.right.parent = root.parent
		}
	} else if root.parent != nil && root.right == nil {
		if root.parent.left == root {
			root.parent.left = root.left
			root.left.parent = root.parent
		} else {
			root.parent.right = root.left
			root.left.parent = root.parent
		}
	}
	root.left.compress()
	root.right.compress()
}

func (bt *BinTrie) Compress() {
	if bt == nil || bt.root == nil {
		return
	}
	bt.root.compress()
	if bt.root.left == nil && !bt.root.isElement {
		bt.root = bt.root.right
		bt.root.parent = nil
	} else if bt.root.right == nil && !bt.root.isElement {
		bt.root = bt.root.left
		bt.root.parent = nil
	}
}
