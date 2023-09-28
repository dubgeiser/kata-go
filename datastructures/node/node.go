package node

type Node struct {
	Value any
	Next  *Node
}

func New(v any) *Node {
	n := Node{v, nil}
	return &n
}

// I want to call this 'Node' too :-(
type DNode struct {
	Value    any
	Next     *Node
	Previous *Node
}

func DNew(v any) *DNode {
	n := DNode{v, nil, nil}
	return &n
}
