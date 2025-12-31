package tree

import "strings"

func incrementSuccessorLevel(root *Node) {
	if root == nil {
		return
	}
	root.level++
	for _, child := range root.children {
		incrementSuccessorLevel(child)
	}
}

func (root *Node) Insert(data string) {
	data = strings.ToLower(data)
	if root == nil {
		return
	}
	path := charToInt(data[root.level])
	child := root.children[path]
	if child == nil {
		// if no child found at path, simply insert
		root.children[path] = NewDataNode(data, root.level+1, root)
	} else if child.data == data {
		// if the child already contains data, make
		// it data node
		child.isData = true
	} else if strings.Index(child.data, data) == 0 {
		// if data is a prefix of child data, make
		// another parent for the child with data
		parent := NewDataNode(data, child.level, root)
		childPath := charToInt(child.data[child.level])
		root.children[path] = parent
		parent.children[childPath] = child
		child.parent = parent
		incrementSuccessorLevel(child)
	} else if len(child.data) == child.level {
		// if the child cannot be pushed further down
		// (due to insufficient length), insert data
		// inside the child
		child.Insert(data)
	} else {
		// make child a branch, and insert child data
		// and data inside the child
		child.isData = false
		childData := child.data
		child.data = childData[:child.level]
		child.Insert(childData)
		child.Insert(data)
	}
}
