package stack

import (
	"kata/datastructures/node"
)

type Stack struct {
	length int
	head   *node.Node
}

func New() *Stack {
	s := Stack{0, nil}
	return &s
}

func (s *Stack) Push(v any) {
	n := node.New(v)
	t := s.head
	s.head = n
	s.head.Next = t
	s.length++
}

func (s *Stack) Pop() any {
	popped := s.head
	s.head = s.head.Next
	popped.Next = nil
	s.length--
	return popped.Value
}

func (s *Stack) Peek() any {
	return s.head.Value
}
