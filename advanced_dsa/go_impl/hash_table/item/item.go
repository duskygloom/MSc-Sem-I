package item

import "fmt"

type Item struct {
	key int
}

func New(key int) *Item {
	return &Item{key: key}
}

func (i *Item) String() string {
	if i == nil {
		return "nil"
	}
	return fmt.Sprintf("%d", i.key)
}

func (i Item) Key() int {
	return i.key
}
