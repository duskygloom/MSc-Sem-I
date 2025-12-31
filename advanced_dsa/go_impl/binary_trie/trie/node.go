package trie

import (
	"fmt"
)

type Node struct {
	left, right *Node
	parent      *Node
	data, level int
	isElement   bool
}

func NewElementNode(data, level int, parent *Node) *Node {
	return &Node{data: data, level: level, parent: parent, isElement: true}
}

func (n *Node) String(maxBits int) string {
	if n == nil {
		return "nil"
	}
	if n.isElement {
		return fmt.Sprintf("%d (%s)", n.data, getBitString(n.data, maxBits))
	}
	s := ""
	if n.level > 0 {
		s = getBitString(n.data, n.level)
	}
	return fmt.Sprintf("%s* -> [%s, %s]", s, n.left.String(maxBits), n.right.String(maxBits))
}
