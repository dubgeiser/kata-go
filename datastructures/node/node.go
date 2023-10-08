package node

type Node struct {
	Value any
	Next  *Node
}

func New(v any) *Node {
	n := Node{v, nil}
	return &n
}
