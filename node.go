package gobinarytree

import (
	"golang.org/x/exp/constraints"
)

type Node[O constraints.Ordered] struct {
	Value O
	Left  *Node[O]
	Right *Node[O]
}

func (n *Node[O]) search(value O) *Node[O] {
	if n == nil || (*n).Value == value {
		return n
	}

	if value < n.Value {
		return n.Left.search(value)
	} else {
		return n.Right.search(value)
	}
}

func (n *Node[O]) insert(value O) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node[O]{Value: value}
		} else {
			n.Left.insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node[O]{Value: value}
		} else {
			n.Right.insert(value)
		}
	}
}

func lift[O constraints.Ordered](node, nodeToDelete *Node[O]) *Node[O] {
	if node.Left != nil {
		node.Left = lift(node.Left, nodeToDelete)
		return node
	}
	nodeToDelete.Value = node.Value
	return node.Right
}

func (n *Node[O]) remove(value O) *Node[O] {
	if n != nil {
		if value < n.Value {
			n.Left = n.Left.remove(value)
			return n
		} else if value > n.Value {
			n.Right = n.Right.remove(value)
			return n
		} else {
			if n.Left == nil {
				return n.Right
			}
			if n.Right == nil {
				return n.Left
			}
			n.Right = lift(n.Right, n)
			return n
		}
	}
	return nil
}

func (n *Node[O]) recPrint() {
	if n == nil {
		return
	}
	n.Left.recPrint()
	println(n)
	n.Right.recPrint()
}
