package tableCH

import (
	"hashtable/item"
)

type Node struct {
	prev, next *Node
	value      *item.Item
}

func NewNode(value *item.Item) *Node {
	return &Node{prev: nil, next: nil, value: value}
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	s := n.value.String()
	node := n.next
	for node != nil {
		s += " -> " + node.value.String()
		node = node.next
	}
	return s
}

func (n *Node) Append(value *item.Item) bool {
	if n == nil {
		return false
	}
	node := n
	for node.next != nil {
		node = node.next
	}
	node.next = NewNode(value)
	node.next.prev = node
	return true
}

func (n *Node) Remove() bool {
	if n == nil || n.prev == nil {
		return false
	} else {
		n.prev.next = n.next
		if n.next != nil {
			n.next.prev = nil
		}
		return true
	}
}
