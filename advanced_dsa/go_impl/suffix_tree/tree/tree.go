package tree

func NewSufTree(word string) *Node {
	root := NewBranchNode(0, nil)
	for i := range len(word) {
		root.Insert(word[len(word)-i-1:])
	}
	return root
}
