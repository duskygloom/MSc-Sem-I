package node

import "fmt"

type Node struct {
	value, order int
	children     []*Node
}

func NewNode(value, order int) *Node {
	if order == 0 {
		return nil
	}
	children := make([]*Node, order)
	for index := range children {
		children[index] = NewNode(0, index)
	}
	return &Node{value: value, order: order, children: children}
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	s := fmt.Sprintf("(%d,%d)[", n.value, n.order)
	if n.order >= 1 {
		s += n.children[0].String()
	}
	for i := 1; i < n.order; i++ {
		s += "," + n.children[i].String()
	}
	s += "]"
	return s
}

func (n *Node) Copy() *Node {
	return &Node{value: n.value, order: n.order, children: n.children}
}

func (n *Node) AppendChild(child *Node) {
	n.children = append(n.children, child)
	n.order++
}
