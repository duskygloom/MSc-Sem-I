package heap

import "binoheap/node"

type Heap struct {
	roots []*node.Node
}

func NewHeap() *Heap {
	return &Heap{roots: make([]*node.Node, 0)}
}

func (h *Heap) String() string {
	if h == nil {
		return "nil"
	}
	if len(h.roots) == 0 {
		return "empty"
	}
	s := h.roots[0].String()
	for _, root := range h.roots[1:] {
		s += "\n" + root.String()
	}
	return s
}
