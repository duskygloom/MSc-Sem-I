package node

import (
	"fmt"
	"strconv"
)

type Node struct {
	Parent       *Node
	Children     []*Node
	Order, Value int
}

func NewNode(order int, parent *Node) *Node {
	n := Node{Order: order, Parent: parent}
	for i := range order {
		n.Children = append(n.Children, NewNode(i, &n))
	}
	return &n
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	} else if n.IsLeafNode() {
		return strconv.Itoa(n.Value)
	} else {
		s := fmt.Sprintf("%d -> [", n.Value)
		prefix := ""
		for _, child := range n.Children {
			s += prefix + child.String()
			prefix = ", "
		}
		s += "]"
		return s
	}
}

func (n *Node) IsLeafNode() bool {
	return n != nil && len(n.Children) == 0
}

func (n *Node) IsParentNode() bool {
	return n != nil && n.Parent == nil
}

// Note: [child] becomes a child of [self]. Both should be of the same [order].
// Return: Returns true if merged successfully, else false.
func (n *Node) Merge(child *Node) bool {
	if n == nil || child == nil || n.Order != child.Order {
		return false
	}
	n.Children = append(n.Children, child)
	n.Order++
	return true
}

func (n *Node) Update(value int) bool {
	ni := n
	if ni == nil {
		return false
	}
	ni.Value = value
	for ni != nil && ni.Parent != nil {
		if ni.Value < ni.Parent.Value {
			ni.Value, ni.Parent.Value = ni.Parent.Value, ni.Value
		}
		ni = ni.Parent
	}
	return true
}
