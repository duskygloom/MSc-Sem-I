package dst

import (
	"fmt"
)

type Node struct {
	key         int
	parent      *Node
	left, right *Node
}

func NewNode(key int, parent *Node) *Node {
	return &Node{key: key, parent: parent}
}

func (n *Node) String(maxBits int) string {
	if n == nil {
		return "nil"
	}
	if n.left == nil && n.right == nil {
		return fmt.Sprintf("%d (%s)", n.key, getBitString(n.key, maxBits))
	}
	return fmt.Sprintf("%d (%s) -> [%s, %s]", n.key, getBitString(n.key, maxBits), n.left.String(maxBits), n.right.String(maxBits))
}

func (n *Node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *Node) delete() {
	if n == n.parent.left {
		n.parent.left = nil
	} else {
		n.parent.right = nil
	}
	n.parent = nil
}
