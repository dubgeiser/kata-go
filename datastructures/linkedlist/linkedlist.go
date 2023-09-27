package linkedlist

import ()

type Node struct {
	Value    any
	Next     *Node
	Previous *Node
}

func New(v any) *Node {
	n := Node{v, nil, nil}
	return &n
}
