package trie

func (n *Node) delete() {
	if n == nil {
		return
	}
	if n.parent.left == n {
		n.parent.left = nil
	} else {
		n.parent.right = nil
	}
	// delete parent if it becomes a leaf branch node
	if n.parent != nil && n.parent.left == nil && n.parent.right == nil && !n.parent.isElement {
		n.parent.delete()
	}
	n.parent = nil
}

func (bt *BinTrie) Delete(data int) {
	if bt == nil || bt.root == nil {
		return
	}
	if bt.root.data == data {
		// if root is data, then simply delete root
		bt.root = nil
	} else {
		n := bt.Lookup(bt.mod(data))
		n.delete()
	}
}
