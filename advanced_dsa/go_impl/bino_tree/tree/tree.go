package tree

import "binotree/node"

type Tree struct {
	order int
	root  *node.Node
}

func NewTree(order int) *Tree {
	root := node.NewNode(0, order)
	return &Tree{order: order, root: root}
}

func (t *Tree) String() string {
	return t.root.String()
}

// Merges a and b with t.
func Merge(a, b *Tree) *Tree {
	if a.order != b.order {
		panic("BOTH TREES SHOULD BE OF THE SAME ORDER")
	}
	t := NewTree(a.order + 1)
	t.root.Copy()
	return t
}
