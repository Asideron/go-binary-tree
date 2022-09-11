package gobinarytree

import (
	"golang.org/x/exp/constraints"
)

type BST[O constraints.Ordered] struct {
	root *Node[O]
}

func NewBST[O constraints.Ordered]() *BST[O] {
	return &BST[O]{}
}

func (t *BST[O]) Get(value O) *Node[O] {
	return t.root.search(value)
}

func (t *BST[O]) Add(value O) {
	if t.root == nil {
		t.root = &Node[O]{Value: value}
	} else {
		t.root.insert(value)
	}
}

func (t *BST[O]) Del(value O) bool {
	if n := t.Get(value); n != nil {
		n.remove(value)
		return true
	}
	return false
}

func (t *BST[O]) Print() {
	t.root.recPrint()
}
