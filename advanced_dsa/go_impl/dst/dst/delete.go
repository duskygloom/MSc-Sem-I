package dst

func findNextLeaf(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.left != nil {
		return findNextLeaf(node.left)
	}
	if node.right != nil {
		return findNextLeaf(node.right)
	}
	return node
}

func (t *DST) Delete(key int) {
	if t == nil {
		return
	}
	n := t.Lookup(key)
	if n == nil {
		return
	}
	if n.isLeaf() {
		// if leaf node, simply delete it
		n.delete()
	} else {
		// replace n with next leaf
		leaf := findNextLeaf(n)
		if leaf == nil {
			panic("Next leaf node is nil") // how can this even happen?
		}
		n.key, leaf.key = leaf.key, n.key
		// delete the leaf
		leaf.delete()
	}
}
