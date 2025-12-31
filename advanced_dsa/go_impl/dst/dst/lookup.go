package dst

func (t *DST) Lookup(key int) *Node {
	if t == nil || t.root == nil {
		return nil
	}
	parent := t.root
	for i := range t.maxBits {
		if parent == nil || parent.key == key {
			return parent
		}
		path := getIthBit(key, i, t.maxBits)
		if path == 0 {
			parent = parent.left
		} else {
			parent = parent.right
		}

	}
	if parent != nil && parent.key == key {
		return parent
	}
	return nil
}
