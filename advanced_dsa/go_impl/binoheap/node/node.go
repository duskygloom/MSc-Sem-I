package node

import (
	"fmt"
)

type Node struct {
	parent       *Node
	children     []*Node
	order, value int
}

func NewNode(order int, parent *Node) *Node {
	n := Node{order: order, parent: parent}
	for i := range order {
		n.children = append(n.children, NewNode(i, &n))
	}
	return &n
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	s := fmt.Sprintf("%d -> [", n.value)
	prefix := ""
	for _, child := range n.children {
		s += prefix + child.String()
		prefix = ", "
	}
	s += "]"
	return s
}

// Note: [child] becomes a child of [self]. Both should be of the same [order].
// Return: Returns true if merged successfully, else false.
func (n *Node) Merge(child *Node) bool {
	if n.order != child.order {
		return false
	}
	n.children = append(n.children, child)
	n.order++
	return true
}
