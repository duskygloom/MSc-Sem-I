package tree

func (root *Node) Lookup(data string) *Node {
	parent := root
	for _, ch := range data {
		if parent == nil {
			return nil
		}
		path := charToInt(byte(ch))
		child := parent.children[path]
		if child != nil && child.data == data && child.isData {
			return child
		}
		parent = child
	}
	return nil
}
