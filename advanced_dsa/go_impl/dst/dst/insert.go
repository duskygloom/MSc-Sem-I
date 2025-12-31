package dst

func (t *DST) Insert(key int) {
	key = t.mod(key)
	if t == nil {
		return
	}
	if t.root == nil {
		t.root = NewNode(key, nil)
	} else {
		parent := t.root
		for i := range t.maxBits {
			if parent.key == key {
				break
			}
			path := getIthBit(key, i, t.maxBits)
			if path == 0 && parent.left == nil {
				parent.left = NewNode(key, parent)
				break
			} else if path == 0 {
				parent = parent.left
			} else if parent.right == nil {
				parent.right = NewNode(key, parent)
				break
			} else {
				parent = parent.right
			}
		}
	}
}
