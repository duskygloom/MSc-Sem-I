package heap

import (
	"binoheap/node"
	"math"
)

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

func (h *Heap) IsEmpty() bool {
	return h == nil || h.roots == nil || len(h.roots) == 0
}

// consolidate once and return if consolidation occurred
// this can be later used to redo consolidation until the
// heap cannot be further consolidated
func (h *Heap) consolidate() bool {
	for i, r1 := range h.roots {
		for j, r2 := range h.roots {
			if r1 != r2 && r1.Order == r2.Order {
				if r1.Value < r2.Value {
					if !r1.Merge(r2) {
						return false
					}
					r2.Parent = r1
					h.roots = append(h.roots[:j], h.roots[j+1:]...)
				} else {
					if !r2.Merge(r1) {
						return false
					}
					r1.Parent = r2
					h.roots = append(h.roots[:i], h.roots[i+1:]...)
				}
				return true
			}
		}
	}
	return false
}

func (h *Heap) insertNode(n *node.Node) {
	h.roots = append(h.roots, n)
	for h.consolidate() {
	}
}

func (h *Heap) Insert(value int) {
	n := node.NewNode(0, nil)
	n.Value = value
	h.insertNode(n)
}

func (h *Heap) minRoot() int {
	minValue := math.MaxInt
	minIndex := -1
	for index, r := range h.roots {
		if r.Value < minValue {
			minValue = r.Value
			minIndex = index
		}
	}
	return minIndex
}

func (h *Heap) Minimum() int {
	minIndex := h.minRoot()
	if minIndex == -1 {
		return 0
	}
	return h.roots[minIndex].Value
}

func (h *Heap) DeleteMin() bool {
	minIndex := h.minRoot()
	minNode := h.roots[minIndex]
	h.roots = append(h.roots[:minIndex], h.roots[minIndex+1:]...)
	for _, c := range minNode.Children {
		h.insertNode(c)
	}
	return true
}

func searchAtNode(n *node.Node, value int) *node.Node {
	if value == n.Value {
		return n
	} else if value > n.Value {
		for _, c := range n.Children {
			foundNode := searchAtNode(c, value)
			if foundNode != nil {
				return foundNode
			}
		}
	}
	return nil
}

func (h *Heap) find(value int) *node.Node {
	for _, r := range h.roots {
		n := searchAtNode(r, value)
		if n != nil {
			return n
		}
	}
	return nil
}

func (h *Heap) Delete(value int) bool {
	n := h.find(value)
	if n == nil || !n.Update(math.MinInt) || !h.DeleteMin() {
		return false
	}
	return true
}
