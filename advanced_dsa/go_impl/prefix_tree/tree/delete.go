package tree

func deleteNode(n *Node) {
	if n == nil || n.parent == nil {
		return // don't delete the root
	}
	for _, child := range n.children {
		// if n contains any child, we cannot delete
		// the node, so we just make it a branch node
		if child != nil {
			n.isData = false
			return
		}
	}
	// can safely delete n
	path := charToInt(n.data[n.level-1])
	parent := n.parent
	n.parent = nil
	parent.children[path] = nil
	// check if parent can be deleted
	if parent != nil && !parent.isData {
		deleteNode(parent)
	}
}

func (root *Node) Delete(data string) {
	n := root.Lookup(data)
	if n == nil {
		return
	}
	deleteNode(n)
}
