package tableCH

// table chaining

import (
	"hashtable/hashing"
	"hashtable/item"
)

type Table struct {
	size   int
	buffer []*Node
	hf     hashing.HashFunction
}

func New(size int, hf hashing.HashFunction) *Table {
	if size < 0 {
		size = 0
	}
	return &Table{size: size, buffer: make([]*Node, size), hf: hf}
}

func (ht *Table) String() string {
	if ht == nil {
		return "nil"
	}

	s := "["
	prefix := ""
	for i := 0; i < ht.size; i++ {
		if ht.buffer[i] != nil {
			s += prefix + ht.buffer[i].String()
			prefix = ", "
		}
	}
	s += "]"
	return s
}

func (ht *Table) NilString() string {
	if ht == nil {
		return "nil"
	}

	s := "["
	prefix := ""
	for i := 0; i < ht.size; i++ {
		s += prefix + ht.buffer[i].String()
		prefix = ", "
	}
	s += "]"
	return s
}

func (ht *Table) Insert(key int) bool {
	hash := ht.hf(key, ht.size)
	if ht.buffer[hash] == nil {
		ht.buffer[hash] = NewNode(item.New(key))
		return true
	} else {
		return ht.buffer[hash].Append(item.New(key))
	}
}

func (ht *Table) Delete(key int) {
	hash := ht.hf(key, ht.size)
	if ht.buffer[hash] != nil {
		node := ht.buffer[hash]
		if node.value.Key() == key {
			// root has to be removed
			ht.buffer[hash] = ht.buffer[hash].next
			return
		}
		for node != nil {
			if node.value.Key() == key {
				node.Remove()
				break
			}
			node = node.next
		}
	}
}

func (ht *Table) Contains(key int) bool {
	hash := ht.hf(key, ht.size)
	if ht.buffer[hash] != nil {
		node := ht.buffer[hash]
		for node != nil {
			if node.value.Key() == key {
				return true
			}
			node = node.next
		}
	}
	return false
}

func (ht *Table) Rehash() {
	newTable := New(ht.size*2, ht.hf)
	for _, n := range ht.buffer {
		if n != nil {
			node := n
			for node != nil {
				newTable.Insert(node.value.Key())
				node = node.next
			}
		}
	}
	ht.size *= 2
	ht.buffer = newTable.buffer
}

func (ht *Table) NthItem(n int) *item.Item {
	return ht.buffer[n].value
}
