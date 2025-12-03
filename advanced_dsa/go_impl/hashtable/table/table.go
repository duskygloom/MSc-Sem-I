package table

import "hashtable/item"

// table interface

type BaseTable interface {
	String() string
	NilString() string

	Insert(key int) bool
	Delete(key int)
	Contains(key int) bool
	Rehash()

	NthItem(n int) *item.Item
}
