package tree

type Node struct {
	parent   *Node
	children [26]*Node
	data     string
	isData   bool
	level    int
}

func NewDataNode(data string, level int, parent *Node) *Node {
	return &Node{data: data, isData: true, level: level, parent: parent}
}

func NewBranchNode(level int, parent *Node) *Node {
	return &Node{level: level, parent: parent}
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	s := n.data
	cs := ""
	prefix := ""
	for _, child := range n.children {
		if child != nil {
			cs += prefix + child.String()
			prefix = ", "
		}
	}
	if n.isData {
		if cs == "" {
			return s
		}
		return s + " -> [" + cs + "]"
	}
	return s + "* -> [" + cs + "]"
}
