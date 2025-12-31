package quadprobing

import (
	"fmt"
	"hashtable/hashing"
	"hashtable/item"
)

type Table struct {
	size   int
	buffer []*item.Item
	hf     hashing.HashFunction
}

func NewTable(size int, hf hashing.HashFunction) *Table {
	if size < 0 {
		size = 0
	}
	return &Table{size: size, buffer: make([]*item.Item, size), hf: hf}
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
	for i := 0; i < ht.size; i++ {
		index := (hash + i*i) % ht.size
		if ht.buffer[index] == nil {
			ht.buffer[index] = item.New(key)
			return true
		} else {
			fmt.Printf("Collision at %d: %d -> %d\n", index, key, ht.buffer[index].Key())
		}
	}
	return false
}

func (ht *Table) Delete(key int) {
	hash := ht.hf(key, ht.size)
	for i := 0; i < ht.size; i++ {
		index := (hash + i*i) % ht.size
		if ht.buffer[index] != nil && ht.buffer[index].Key() == key {
			ht.buffer[index] = nil
			break
		}
	}
}

func (ht *Table) Contains(key int) bool {
	hash := ht.hf(key, ht.size)
	for i := 0; i < ht.size; i++ {
		index := (hash + i*i) % ht.size
		if ht.buffer[index] != nil && ht.buffer[index].Key() == key {
			return true
		}
	}
	return false
}

func (ht *Table) Rehash() {
	newTable := NewTable(ht.size*2, ht.hf)
	for _, v := range ht.buffer {
		if v != nil {
			newTable.Insert(v.Key())
		}
	}
	ht.size *= 2
	ht.buffer = newTable.buffer
}

func (ht *Table) NthItem(n int) *item.Item {
	return ht.buffer[n]
}
