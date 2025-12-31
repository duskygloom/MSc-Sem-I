package dst

type DST struct {
	root    *Node
	maxBits int
}

func NewDST(maxBits int) *DST {
	return &DST{maxBits: maxBits}
}

func (t *DST) String() string {
	if t == nil {
		return "nil"
	}
	if t.root == nil {
		return "empty"
	}
	return t.root.String(t.maxBits)
}

func (t *DST) mod(key int) int {
	return key % (1 << t.maxBits)
}
