package queue

import (
	"kata/datastructures/node"
)

type Queue struct {
	head   *node.Node
	tail   *node.Node
	length int
}

func New() *Queue {
	q := Queue{nil, nil, 0}
	return &q
}

func (q *Queue) Enqueue(v any) {
	n := node.New(v)
	if q.head == nil {
		q.head = n
	}
	if q.tail != nil {
		q.tail.Next = n
	}
	q.tail = n
	q.length++
}

func (q *Queue) Deque() any {
	dq := q.head
	q.head = q.head.Next
	dq.Next = nil
	q.length--
	return dq.Value
}

func (q *Queue) Peek() any {
	return q.head.Value
}
